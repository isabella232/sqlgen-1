package example

import "time"

//go:generate sqlgen -structs Job -sql -array
type Job struct {
	ID     string
	Amount *Money
	Name   string
	Number Int64Alias
}

//go:generate sqlgen -structs Money -sql -array
type Money struct {
	Amount   int64
	Currency string
	Rounded  bool
}

//go:generate sqlgen -structs Complete -sql -array
type Complete struct {
	CustomString  string
	CustomInt16   int16
	CustomInt32   int32
	CustomInt     int
	CustomInt64   int64
	CustomUint16  uint16
	CustomUint32  uint32
	CustomUint    uint
	CustomUint64  uint64
	CustomFloat32 float32
	CustomFloat64 float64
	CustomBool    bool
}

//go:generate sqlgen -structs Times -sql -array
type Times struct {
	CustomTime  time.Time
	CustomTimep *time.Time
}

//go:generate sqlgen -alias Int64Alias
type Int64Alias int64
