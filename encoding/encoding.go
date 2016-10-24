package encoding

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func AppendArrayQuotedBytes(b, v []byte) []byte {
	b = append(b, '"')
	for {
		i := bytes.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}

func SplitBytes(src []byte) ([][]byte, error) {
	if len(src) == 0 {
		return nil, nil
	}
	if src[0] != '(' {
		return nil, fmt.Errorf(
			"unable to parse type, expected %q at offset %d", '(', 0)
	}

	buf := &scanner{bytes.NewReader(src[1:])}
	chunks := make([][]byte, 0)
	var (
		depth              = 1
		elem               []byte
		quoted, aQuoteSeen bool
		prev               rune
	)
	for {
		b, err := buf.Next()
		if err != nil && err != io.EOF {
			// I can't think of why this would happen
			return nil, err
		}
		switch b {
		case '(':
			elem = append(elem, b)
			if escaped(prev) {
				break
			}
			depth++
		case ')':
			if escaped(prev) {
				elem = append(elem, b)
				break
			}
			depth--
			if depth == 0 {
				chunks = append(chunks, elem)
				elem = nil
				break
			}
			elem = append(elem, b)
		case ',':
			if !quoted && !escaped(prev) {
				chunks = append(chunks, elem)
				elem = nil
				break
			}
			elem = append(elem, b)
		case '"':
			if quoted {
				p, _ := buf.Peek()
				if p == '"' {
					// probably next item is double quoted
					if !aQuoteSeen {
						aQuoteSeen = true
						break
					}
				}
				if aQuoteSeen {
					elem = append(elem, b)
					aQuoteSeen = false
					break
				}
				quoted = false
				break
			}
			if prev != '"' {
				quoted = true
				break
			}

		default:
			if escaped(rune(b)) {
				break
			}
			elem = append(elem, b)
		}
		prev = rune(b)
		if err == io.EOF {
			break
		}
	}
	return chunks, nil
}

func escaped(r rune) bool {
	return r == '\\'
}

type scanner struct {
	*bytes.Reader
}

func (s *scanner) Next() (byte, error) {
	return s.ReadByte()
}

func (s *scanner) Peek() (byte, error) {
	b, err := s.ReadByte()
	_, err = s.Seek(-1, 1)
	if err != nil {
		panic(err)
	}
	return b, err
}

func ScanLinearArray(src, del []byte, typ string) (elems [][]byte, err error) {
	dims, elems, err := parseArray(src, del)
	if err != nil {
		return nil, err
	}
	if len(dims) > 1 {
		return nil, fmt.Errorf("pq: cannot convert ARRAY%s to %s", strings.Replace(fmt.Sprint(dims), " ", "][", -1), typ)
	}
	return elems, err
}

// parseArray extracts the dimensions and elements of an array represented in
// text format. Only representations emitted by the backend are supported.
// Notably, whitespace around brackets and delimiters is significant, and NULL
// is case-sensitive.
//
// See http://www.postgresql.org/docs/current/static/arrays.html#ARRAYS-IO
func parseArray(src, del []byte) (dims []int, elems [][]byte, err error) {
	var depth, i int

	if len(src) < 1 || src[0] != '{' {
		return nil, nil, fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '{', 0)
	}

Open:
	for i < len(src) {
		switch src[i] {
		case '{':
			depth++
			i++
		case '}':
			elems = make([][]byte, 0)
			goto Close
		default:
			break Open
		}
	}
	dims = make([]int, i)

Element:
	for i < len(src) {
		switch src[i] {
		case '{':
			if depth == len(dims) {
				break Element
			}
			depth++
			dims[depth-1] = 0
			i++
		case '"':
			var elem = []byte{}
			var escape bool
			for i++; i < len(src); i++ {
				if escape {
					elem = append(elem, src[i])
					escape = false
				} else {
					switch src[i] {
					default:
						elem = append(elem, src[i])
					case '\\':
						escape = true
					case '"':
						elems = append(elems, elem)
						i++
						break Element
					}
				}
			}
		default:
			for start := i; i < len(src); i++ {
				if bytes.HasPrefix(src[i:], del) || src[i] == '}' {
					elem := src[start:i]
					if len(elem) == 0 {
						return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
					}
					if bytes.Equal(elem, []byte("NULL")) {
						elem = nil
					}
					elems = append(elems, elem)
					break Element
				}
			}
		}
	}

	for i < len(src) {
		if bytes.HasPrefix(src[i:], del) && depth > 0 {
			dims[depth-1]++
			i += len(del)
			goto Element
		} else if src[i] == '}' && depth > 0 {
			dims[depth-1]++
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}

Close:
	for i < len(src) {
		if src[i] == '}' && depth > 0 {
			depth--
			i++
		} else {
			return nil, nil, fmt.Errorf("pq: unable to parse array; unexpected %q at offset %d", src[i], i)
		}
	}
	if depth > 0 {
		err = fmt.Errorf("pq: unable to parse array; expected %q at offset %d", '}', i)
	}
	if err == nil {
		for _, d := range dims {
			if (len(elems) % d) != 0 {
				err = fmt.Errorf("pq: multidimensional arrays must have elements with matching dimensions")
			}
		}
	}
	return
}
