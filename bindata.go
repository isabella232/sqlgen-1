// Code generated by go-bindata.
// sources:
// templates/aliastype.go.tmpl
// templates/common.go.tmpl
// templates/create-type.sql.tmpl
// templates/sqltype.go.tmpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesAliastypeGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x51\xcd\x6a\xe3\x30\x10\x3e\x5b\x4f\x31\x18\x16\xa4\x60\xec\x3d\x07\x72\xd8\x85\x5e\x4b\xa1\xa5\x97\xd2\x83\x2c\x8f\x53\x11\x5b\x4e\x47\x93\x34\x41\xe8\xdd\x8b\x2c\xa7\x34\x81\xd6\x17\x8f\x34\xdf\x7c\x3f\xa3\xa6\x81\x2d\x3a\x24\xcd\xd8\x81\x99\x3a\x14\x7b\x6d\x76\x7a\x8b\x10\x42\xfd\x90\xcb\x7b\x3d\x62\x8c\x42\xd8\x71\x3f\x11\x83\x14\x45\xd9\x69\xd6\xad\xf6\xd8\xf8\xf7\xa1\xe9\xc8\x1e\x91\x4a\x51\x94\x9e\xc9\x4c\xee\x98\xca\x7e\xe4\x52\x28\x21\xfa\x83\x33\x20\x4f\xb0\x0a\xa1\xfe\x37\x58\xed\x33\x9b\x82\x47\xa3\x9d\xf4\x64\xc0\x3a\x46\xea\xb5\xc1\x10\x15\x20\xd1\x44\x10\x44\xe1\x3f\x2c\x9b\x37\x48\x80\xf5\x26\xfd\x6a\xc9\xe7\x3d\xaa\xd4\x33\xda\x23\xbc\xbc\xb6\x67\xc6\xb5\x28\x0a\x42\x3e\x90\x83\x53\xed\x8d\x76\xff\xcf\x8c\x3e\xf1\xaa\x05\xe7\x99\xac\xdb\xfe\x80\xcb\x24\x33\x5c\x89\x22\x84\x3a\xb9\x72\x48\x31\x8a\x22\x8a\xcb\x44\x3f\x72\x7d\x97\x8c\xf5\xb2\x4c\xfd\x89\x21\xc5\x44\x62\xf8\xf3\x04\x3c\xc1\x75\xb6\xb2\x82\x59\x3f\xfe\x92\xfe\xca\xea\x92\x65\x49\x1f\x04\xa4\x6f\x31\x33\x83\x62\xcc\x77\xab\x13\x6c\x6e\xc4\xa4\x55\xb9\xb7\x78\x75\x76\xf8\x2e\x7c\xab\xfb\xac\x87\x03\x4a\x05\x32\x3f\x5a\x3d\x9f\xab\x2c\x9c\x76\x3b\x53\xb5\x69\xe5\xa3\xde\xe1\xb2\x9f\xea\xaf\xfa\xf2\x34\x0f\xd0\xc5\xd0\x22\xda\x56\x59\xf7\x33\x00\x00\xff\xff\x61\x2f\x3f\xe1\x4e\x02\x00\x00")

func templatesAliastypeGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAliastypeGoTmpl,
		"templates/aliastype.go.tmpl",
	)
}

func templatesAliastypeGoTmpl() (*asset, error) {
	bytes, err := templatesAliastypeGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/aliastype.go.tmpl", size: 590, mode: os.FileMode(420), modTime: time.Unix(1476857686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCommonGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\xdd\x6e\xdb\xc8\x0e\xbe\x96\x9e\x82\xd5\x41\x62\xab\x51\xec\xf8\xd6\xad\x0b\xe4\x1c\xe4\x60\x0b\x14\x49\xb7\xd9\xbd\x58\x38\x46\x3b\x96\x28\x7b\x10\x69\xa4\xcc\x8c\xec\x18\xae\xde\x7d\xc1\x99\x91\x2c\x3b\x4e\xb2\xdb\xdd\x05\xf6\xc6\xd6\x0f\xc9\x21\x3f\x92\x1f\xa9\xe1\x10\x16\x28\x50\x32\x8d\x09\xc4\x45\x82\x7e\xc9\xe2\x7b\xb6\x40\xd8\x6e\x07\x9f\xed\xe5\x35\xcb\xb1\xae\x7d\x9f\xe7\x65\x21\x35\xf4\x7d\x2f\x98\x6f\x34\xaa\xc0\xf7\x82\x34\xd7\xf4\xc7\x0b\xfa\x55\x5a\x72\xb1\x50\x81\x1f\xfa\x7e\x5a\x89\x18\x58\x59\xa2\x48\x2e\xa5\x64\x9b\x9f\xab\x42\x63\xf2\x5f\xd2\xeb\xcf\x23\x58\xc1\x74\x46\x46\x42\xf7\x0f\x5b\xdf\x9b\xc3\xc4\x69\x90\x44\x2f\xe8\x85\xbe\x97\x16\x92\x5e\x79\x1c\xc6\x13\x30\xa7\x0e\x3e\x8a\x04\x1f\x2f\xc5\xa6\xbf\x8a\xe0\x5b\x70\xf7\x2d\xa4\xd7\x29\x70\x78\x0f\x17\x46\xf6\xc0\xd0\x6a\x30\x18\x84\xe6\xb1\x44\x76\xef\x7b\x5e\xdd\x28\x7c\x78\x46\x61\x3a\xe6\x33\xa7\x44\xb2\x07\x7e\xdd\xdd\xf5\x48\x86\xcf\xe8\xfd\x0a\x26\x74\x7d\x36\x1a\xcf\x7c\x12\x96\xa8\x2b\x29\x0e\xc3\xa8\x1d\x1e\xaa\xcc\xb8\xb6\x18\x28\x19\xb7\x10\xf4\xa7\x33\x7b\x19\x01\x4a\x59\xc8\x90\xbc\xe2\x29\x64\x28\x48\x2e\x84\xf7\x30\x82\xef\xdf\x41\xc9\x78\x7a\x31\x83\x37\x13\xe8\xf5\x7b\xc6\x73\x77\x9c\xe0\x59\x04\x69\xae\x07\x57\xa4\x9e\xf6\x83\x4a\xb0\x79\x86\xa0\x0b\x28\x99\x54\x08\x7a\x53\x92\xf1\xc7\x12\x63\x4a\xf4\xc9\x03\x30\x0d\x45\x9a\x2a\xd4\x70\x92\x04\x11\x19\x8c\xe0\x22\xa4\x18\x7c\x4f\x12\xd8\xa7\x2a\x66\x42\xa0\xdc\x5a\xd4\xaf\x71\xfd\x05\x59\x82\x92\x3c\x9a\x8e\xc6\xb3\xb0\xf6\xbd\x78\x59\x89\x7b\x45\xd2\x39\xbb\xc7\x4e\x18\x64\x69\xc5\x24\xd5\x8a\x87\x19\xe6\x00\x2e\x58\xdf\xf3\x4a\x89\x2b\x00\x70\x77\x09\x96\x7a\x09\x30\x81\x91\xef\x79\x0f\xa6\x48\x60\x5e\x14\x99\xef\x75\xb3\x9f\x82\x95\x9b\x4c\x9a\x94\x0d\x87\xf0\xcb\x92\x2b\xc8\xf9\x62\xa9\x21\x47\x26\x60\x8d\xc0\x17\xa2\x90\x08\xaa\xc8\x11\xb4\x64\x3c\xe3\x62\x61\xcb\xc6\xe9\x7c\x84\xa4\x10\x3d\x0d\xf7\xa2\x58\xc3\x7a\xc9\x34\x68\x63\x05\x99\x50\x10\x57\x52\xa2\xd0\xd9\xe6\xb0\x5a\xe6\x26\x31\x14\xa7\x1c\x5c\xe3\xa3\xee\xbb\x9a\xa3\x87\x6f\x26\x04\x3f\x9c\x9e\x36\x77\xbc\x18\x5c\xdd\xfc\xdf\x7a\xd9\xcd\x0f\x4a\xe9\xcc\xa9\x35\xd7\xf1\x12\xe6\x46\x26\x66\x0a\xa1\x17\xf6\xc6\x24\xcf\x53\x78\xe3\x40\x30\xfa\x16\x9e\xf3\x73\x73\xed\xd0\x6e\x8b\xd1\xde\x47\x40\x00\x9b\x02\xaf\x01\x33\x85\x4e\xd3\xc0\xde\xca\xd2\x5d\x04\x73\x2b\xd6\x9e\x1a\xd8\x53\xe7\x9b\x4e\x7c\x9f\x11\xef\x4d\x7c\x4d\x80\x13\x1b\xa0\xb5\xca\x53\x98\x6f\xe8\x51\x2f\xe8\xb9\x47\x2f\x9c\xe4\x79\x5f\x23\xf8\x0a\x7b\xb0\x75\xa0\xb5\xae\xd8\x1f\x9e\x82\x29\x0c\x2a\xee\xbb\xbb\xc6\xb6\xc3\x62\xd2\xa0\xf2\x43\x51\x46\xcf\x60\xfb\x2a\x9e\x8d\x75\x57\xdc\xbb\xd2\xfe\x53\x4e\x24\x98\xb2\x2a\xd3\xc6\x87\x67\x25\xeb\xa6\x31\x26\x30\x77\x4d\x68\x4b\xa7\x71\x4a\xf0\x8c\x68\x84\x1a\x19\x5c\x67\x82\xd2\xb2\x8a\x35\x39\xf1\xd6\x36\xa9\xed\xd0\x96\x6e\xfa\x0a\xde\x3a\xd9\x10\x6c\x02\xa0\x7f\xc8\x33\xee\x1c\xab\x4c\xdc\xd4\x0f\x8f\x1b\xb0\x85\xf1\xd4\xc0\xae\x39\xf6\x6c\x50\xea\x4d\xf9\x80\x1a\xdc\x92\xea\xf9\x28\x82\x51\xe8\x1f\xf4\x0d\x21\x58\x32\xc1\xe3\x3e\x4a\x19\x76\x39\xd4\xda\xdd\x71\x67\xcc\xc4\x27\x2e\x90\x49\x33\x4f\x88\x86\x22\x48\x30\x83\x26\x31\x7a\x53\x82\x9d\x40\x21\x18\x68\x15\xec\xf1\x6a\xc7\xe5\x84\xe7\x2e\xcf\xaa\x75\xde\x10\xe5\xbe\xe9\xa3\xce\x3e\x69\xe9\xba\xe5\x6a\x32\x1b\xc2\x07\x18\xbd\xc8\xcd\xe5\xc3\x18\x08\xd3\x42\x43\x5c\x88\x15\x4a\x0d\x97\x5f\xbe\x5c\xfe\x76\xa2\x88\xaf\x4f\x54\x10\xb9\x30\x08\xce\x32\x63\x31\xf6\x49\xfd\xb6\x94\x5c\x68\x7b\x46\x04\x01\x04\x11\x04\xb3\x69\x10\xc1\xf9\x28\x34\xc1\xef\x81\xb7\x8b\x8d\x00\x1c\x0e\x3b\xe1\x01\x3e\x6a\xc9\x62\xad\x40\x2f\x11\x12\x9e\xa3\x50\xbc\x10\x0a\x98\x48\x8c\x1e\x0a\xad\xa0\x48\x81\x09\x60\x46\x41\x62\x29\x51\xa1\xa0\xd6\xe1\x82\xac\x69\x7c\xd4\x90\x16\x32\x67\x7a\x00\x37\x22\xeb\xc8\x30\x6d\x8c\x61\xce\xb5\xe1\xf2\x8d\x39\x66\xce\xe2\x7b\x14\x09\x30\x62\xe7\xaa\xa4\x45\x02\x93\x01\x99\xba\x2e\x34\x9b\x67\x9b\x08\xd6\x4b\xae\x51\x95\x2c\x46\x60\xb2\xa8\x44\x02\x73\x49\x5a\xda\x7a\x96\x60\xc6\x73\xae\x51\x2a\xe0\x0a\x14\x5f\x08\x9e\xf2\x98\x09\x1d\x99\xd7\xd7\xbf\x7e\xfa\x44\xe6\xb8\x02\xea\xfa\x73\x45\x51\x69\xbe\x42\x3a\x84\x5e\xdc\x22\xc2\x52\xeb\x72\x3c\x1c\xae\xd7\xeb\x41\x59\x28\xbd\x90\xa8\x1e\xb2\x41\x21\x17\xc3\xa4\x88\xd5\xd0\xf1\xff\x50\x51\x10\xf1\xd0\x04\xaf\x06\x4b\x9d\x67\xff\x31\x19\xba\x3d\xff\x78\x63\x8b\xf1\x48\xb1\xec\xa6\x39\xa5\x08\xa6\x33\x4e\xae\xbd\x5c\x87\x34\x22\x0d\xc3\x47\xc0\x81\x0b\xed\xbf\x36\xf6\xb7\x4f\xc7\xfe\xd1\xfa\x3a\x9c\xff\x26\x94\x77\x2f\x2e\x00\xdb\xdd\x02\x70\x53\xa2\x18\xdb\xe9\x4b\x5b\x55\xeb\xcf\x76\x37\xbb\xc8\x27\x3e\xeb\x0c\xb0\xad\x25\x59\x13\xcd\xd9\x99\xe1\x5b\xf3\x67\xdf\xd6\xbd\x96\xfe\x14\x1c\xdd\x16\x3c\x6f\x51\xe8\x02\xfe\x97\x15\x0a\x0f\x48\xd3\x8c\x0b\x20\xa7\x2c\x51\xd6\xb6\x81\x77\x76\x0c\xd2\x3c\xf4\xfd\x2b\x5b\xbd\x7f\xdd\x77\xb2\x3f\xb5\xb3\x77\x34\x83\x09\x5c\x1c\x06\xe4\x26\x27\x25\xd0\x71\xba\x0d\x66\x5b\xb7\x8f\x55\xcc\x4a\x74\xbb\x8c\x67\x1d\x3a\x3b\x7b\xb7\xe7\xd5\x3b\x7a\xb4\x9b\xac\x4e\xe5\x85\xb1\x6a\x5d\x77\x63\xd4\x89\x4f\x20\x65\x99\x01\xed\x60\x2c\x1d\x89\xd7\xdb\x47\xf6\x8f\x9c\x62\xe3\xbd\xbb\xeb\xb5\x1a\xcd\xb1\x5a\x56\xd8\x95\x09\x7a\x5d\xa3\x6a\xdf\xea\xde\x5c\x6d\xb0\x6c\x97\x01\x70\x99\xb3\xcf\xea\xfd\xdd\xa0\xeb\x31\xc1\xa8\x34\x93\x9a\x28\x9b\xbf\x88\xa6\x1d\x89\x3f\x31\xf5\x59\x62\xca\x1f\xcd\xde\xca\xc7\x33\xcb\xec\x4d\x5f\xf1\x99\x59\x66\xea\xfd\x65\x86\x66\x99\x8c\xa7\xe6\xa0\x31\x9f\xd9\x37\xae\x33\x4d\x14\x9d\x45\xd4\xfb\xf1\x76\xac\xc4\x0b\x0d\x69\xbd\x33\x75\xdd\x01\x65\x17\xd8\xd5\x43\xc5\x32\x97\x30\x5b\x7b\xfd\x80\x48\x30\x08\xc3\xd6\x31\x97\x5c\xda\x1d\xba\x36\x5e\x4d\xcf\xd3\x9c\xb4\xc9\xa8\xed\x76\x72\xb4\xbf\x5e\x05\x7d\xfb\xa4\xb7\x1c\x55\xc0\xd9\xc4\x4e\x50\x33\x74\x1d\x17\xec\xce\x77\x65\xcd\xd3\x63\x39\x3b\x62\xaf\xb3\x35\xdb\x42\xeb\xf6\xc5\x3f\x9f\x2d\x8b\x91\xa1\xb2\xe7\xa8\xe8\x49\x28\xa7\xa7\xee\x13\xa7\xfd\x28\xfd\x37\x04\xb1\xfb\xf2\x6a\xdc\xb2\x2b\xdd\xdf\x30\x6f\xea\x9e\x3d\xa8\xf6\x8f\x7c\x68\x10\x66\x5f\x23\x48\xcc\x27\x09\x13\x0b\xb3\xa7\x28\x1b\x3a\x4f\xa1\xdf\x34\xa2\x0a\xe1\x04\x92\x90\x66\x63\xd3\x8e\xc7\x1d\xcc\xab\x4c\xf3\x76\xd7\x61\x99\xf5\x50\x41\x5e\x29\x0d\x4b\xb6\xc2\xdd\xe6\xb3\xe6\x7a\x09\x39\xd3\xf1\x92\x3e\x1e\x77\xfb\x51\x10\xee\xf5\x80\x4b\x81\x5f\xfb\xbf\x07\x00\x00\xff\xff\xc8\x97\xf8\xa2\x44\x11\x00\x00")

func templatesCommonGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommonGoTmpl,
		"templates/common.go.tmpl",
	)
}

func templatesCommonGoTmpl() (*asset, error) {
	bytes, err := templatesCommonGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/common.go.tmpl", size: 4420, mode: os.FileMode(420), modTime: time.Unix(1476857686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCreateTypeSqlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8d\x31\x0a\xc2\x50\x10\x44\x6b\x73\x8a\x29\x52\x28\x84\x1c\x40\xb1\xb5\x12\x41\xf4\x02\x1f\xff\x44\x3e\xe8\x12\x92\x14\x86\x65\xee\x2e\x0b\x1f\xd2\xed\x0e\x6f\xde\xbc\x26\xa6\x85\x58\xd6\x91\x70\xef\x1f\xf7\xeb\x73\x1d\x79\x4b\x5f\x4a\x48\x33\xf6\x70\xc7\x94\xec\x4d\xb4\xc5\x32\x7f\x1d\xda\xa1\xf0\x93\x71\x3c\xa3\xbf\xc4\x35\x4b\x0d\x80\x00\xcb\x50\x29\x48\x9d\x3b\x2d\x4b\x91\x87\x37\x9c\x90\x76\xf5\x8d\x99\xad\x48\xcb\x90\x9a\xc3\xe9\x1f\x00\x00\xff\xff\x47\x62\x8b\xd8\x90\x00\x00\x00")

func templatesCreateTypeSqlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCreateTypeSqlTmpl,
		"templates/create-type.sql.tmpl",
	)
}

func templatesCreateTypeSqlTmpl() (*asset, error) {
	bytes, err := templatesCreateTypeSqlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/create-type.sql.tmpl", size: 144, mode: os.FileMode(420), modTime: time.Unix(1476857686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSqltypeGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x4d\x6f\xe3\x36\x10\x3d\x8b\xbf\x62\x56\xd8\xc0\x92\xa3\xca\x4e\x7b\x73\xea\x02\x2d\xd0\x9e\x8a\xa0\x45\x83\x5e\x5c\x1f\x28\x69\xe4\x10\x96\x29\x85\xa4\x1d\x1b\x5a\xfe\xf7\x62\x48\xc5\x92\x3f\xb2\x9b\xa0\x87\xe6\xe0\x48\x9c\x37\x33\x7c\xa3\x79\x43\x4e\x26\xb0\x42\x89\x8a\x1b\x2c\x20\xaf\x0b\x64\x0d\xcf\xd7\x7c\x85\xd0\xb6\xe9\x1f\xfe\xf1\x81\x6f\xd0\x5a\xc6\xc4\xa6\xa9\x95\x81\x88\x05\x61\xc1\x0d\xcf\xb8\xc6\x89\x7e\xae\xc2\xb3\xf7\x49\xa1\xc4\x0e\x15\x2d\x6b\xa3\xf2\x5a\xee\xe8\xb1\xdc\x98\x90\xc5\x8c\xb5\x2d\x88\x12\xd2\x9f\x95\xe2\x07\xb0\x96\x99\x43\xe3\x72\x3d\x1e\x9a\x2e\x91\x37\x2d\x96\x27\x8b\x8c\x4d\x26\xf0\x37\xaf\xb6\x08\x62\xd3\x54\xb8\x41\x69\x34\x98\x27\x04\x9f\x2d\x75\x36\x05\x42\x1a\x54\x25\xcf\x31\x65\xe5\x56\xe6\x10\xed\xaf\x04\x8f\x7d\xa4\x28\x86\x68\xe8\x9d\x00\x2a\x55\xab\x18\x5a\x16\x88\x12\xf6\x30\x9f\x83\x14\x15\xbd\x06\x0a\xcd\x56\x49\x7a\x4d\xe8\x87\x05\x96\x39\x90\x84\xd9\x1c\x2a\x94\xd1\x3e\xbe\x07\x09\x3f\xc1\xd4\xc1\x27\x13\x78\x7c\x42\x85\xf0\x22\xaa\x0a\x32\x04\x6e\xa0\x42\xae\x0d\x98\x97\x1a\xf2\xad\xaa\x0e\x90\x29\x9e\xaf\xd1\xe8\x04\xbe\x1f\x3f\x40\x76\x30\xa8\xa1\x2e\xe1\x79\x5b\x1b\xd4\x89\x0f\xc2\x65\x01\x0f\xdf\xdd\xf5\xd6\x02\x2b\xb1\x11\x06\x95\x4e\x59\x10\x64\x94\x7d\xc3\xd7\x18\x2d\x96\x04\x49\xe0\x2e\x81\xbb\xdb\x1f\xc6\x32\x26\xeb\x62\xba\x84\x39\x8c\xda\x11\x63\x41\xb0\xe3\xd5\xd4\x31\x24\x9f\xfd\x62\xba\x4c\xbb\x22\xb0\x80\x88\x90\xe1\x53\xcf\xf7\x84\x30\x2a\xc5\x82\xc0\xfa\x20\x09\xd4\x6b\x0a\x41\xf1\xd2\x2e\x6f\x17\xe3\x53\xbd\xbe\x74\x2e\x37\x26\xfd\x95\xea\x5a\x46\xe1\x56\xe2\xbe\xc1\x9c\x7a\x6d\x47\xc9\xc3\xb8\x8b\x9b\xc1\x1c\x78\xd3\xa0\x2c\xdc\x07\xfa\x93\x6a\x50\xfc\x42\xa4\xa3\x2c\x21\x2c\x01\xcb\x5a\x81\xa0\xd4\x77\xf7\x20\xe0\x47\x90\xf7\x20\x6e\x6f\x7d\xc6\x1d\xaf\xc4\x80\x9d\x18\xb2\xbb\x46\xef\x0a\x3f\xb7\x91\x33\x86\x62\xc8\xf0\x84\xe2\xc7\x38\xfa\xd8\x3d\x4b\x62\x35\x4a\x46\xf1\xe9\xea\x57\xb8\x53\xbb\xbd\xa6\xd4\x46\x09\xb9\x8a\x06\x91\xec\x28\x8e\x07\x7d\xd9\xe1\xc2\xd6\x86\x7e\xd5\xcb\xe7\xaf\x9c\xcb\x73\xf5\xe8\xe7\x2a\xa5\x75\x79\x5d\x3b\xe3\x6b\xe2\x21\x7c\xa4\x55\xde\x3b\xb4\x36\xf6\xda\x81\x96\x01\xfd\x51\xa9\x08\xd1\x2b\x08\xfc\x5f\x5f\x35\xda\x6a\xa0\x5f\x84\xc9\x9f\x1c\x74\x36\xa7\x7f\x69\x44\x03\xc1\x49\x30\xe7\x1a\xc1\x57\x7f\xd6\x93\xdf\xa7\x3a\xe7\xd2\x97\x47\xab\x3c\xee\x70\xbe\x26\x6f\xe0\x7c\x10\x07\x8f\x4f\x2a\x34\xfc\x6e\xcd\xf3\x0c\xa8\x10\xb5\x01\x9a\x59\xa8\x0c\xdc\x3c\x82\xa9\xaf\xcc\x8f\x7f\xe4\x8d\x0e\x13\xda\xae\xfb\x89\xa9\xbe\x5f\xad\xd8\xc9\x96\x3b\x4e\x7d\xc5\x02\xac\x70\xa3\x8f\xdd\x4b\xe0\xdf\x85\x44\xae\x9c\x77\xe4\xd2\x78\x9f\x76\x94\x8c\x6c\x02\xe1\x49\x0e\x6a\xb0\xcb\x0e\xef\x18\xba\xd6\xb6\x0e\x40\x53\xca\x65\x8a\xe9\xb3\xf8\x39\x35\xde\xc3\x1c\xa2\xf1\x3e\x5e\xcc\xa6\x4b\x16\x58\xc0\x4a\xa3\xb3\xf4\xa3\xe5\x92\x50\x32\x88\x75\x14\x66\x02\x3b\x72\x51\x5c\xae\x10\x9c\xcd\x0b\xa5\x36\x4f\xe8\x78\x9d\xc4\x69\x9d\x20\x3a\xc6\x0e\xe2\xda\x30\xda\x7d\x4b\xb0\x43\xad\x8a\x12\xb2\x85\xa0\x19\xe7\x22\xdc\xc3\x6e\x38\xb2\x83\xb7\xbe\x72\xc3\x95\x16\x72\x05\xdc\x1d\x34\xe8\xe5\x00\x42\x16\xb8\x87\x9b\xe2\xa2\x0b\x28\x9e\xa9\xbb\x0e\x0b\x13\x10\x47\x41\xdb\xd7\x0a\x66\xae\xc6\x83\xde\xb6\x74\xce\xa1\x2c\xec\x9b\x9d\xf1\x1f\x64\xd4\xe7\x81\xff\x51\x44\xd7\xaa\xfb\x6d\xfd\xbc\x5f\x3a\x1f\x55\x4d\x53\x09\x33\x64\xf4\x41\x49\x7c\x72\x0d\xfa\x9b\xc0\xaa\xd0\x5f\x2a\x94\xd6\x0e\x7d\x86\x24\x8f\xe3\xfd\xa6\x78\x6d\x1e\x0d\xab\xda\xc0\x4d\x11\x26\xe7\x41\xce\xa4\x62\x59\xd0\xb6\x9d\x46\x3e\xbb\x8e\x4b\xe0\x73\x49\x78\xe2\xd0\x79\x5a\x8f\xf2\x76\xf8\xd2\x01\x8e\x63\xba\x33\xa3\x2c\xdc\xe3\x49\xd7\x5d\xbf\xf4\xbc\xe3\xbe\x73\xe5\x26\x11\xb3\xe3\x0d\x22\x1a\xbd\x67\xdf\xdd\xce\x44\xd9\x41\xac\xbd\x3c\xf0\xda\xd6\xab\x22\x68\xdb\x73\x7e\xfe\x06\x77\xc6\xee\x2c\x42\x4c\x47\x66\xc7\x38\x7b\x3d\xdb\xfe\x0d\x00\x00\xff\xff\x44\xa9\xd0\xb7\xc4\x0a\x00\x00")

func templatesSqltypeGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSqltypeGoTmpl,
		"templates/sqltype.go.tmpl",
	)
}

func templatesSqltypeGoTmpl() (*asset, error) {
	bytes, err := templatesSqltypeGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/sqltype.go.tmpl", size: 2756, mode: os.FileMode(420), modTime: time.Unix(1476857686, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/aliastype.go.tmpl": templatesAliastypeGoTmpl,
	"templates/common.go.tmpl": templatesCommonGoTmpl,
	"templates/create-type.sql.tmpl": templatesCreateTypeSqlTmpl,
	"templates/sqltype.go.tmpl": templatesSqltypeGoTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"aliastype.go.tmpl": &bintree{templatesAliastypeGoTmpl, map[string]*bintree{}},
		"common.go.tmpl": &bintree{templatesCommonGoTmpl, map[string]*bintree{}},
		"create-type.sql.tmpl": &bintree{templatesCreateTypeSqlTmpl, map[string]*bintree{}},
		"sqltype.go.tmpl": &bintree{templatesSqltypeGoTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
