package example

import "time"

//go:generate generate-composite-types -structs Job -sql -array
type Job struct {
	ID     string
	Amount *Money
	Name   string
	Number Int64Alias
}

//go:generate generate-composite-types -structs Money -sql -array
type Money struct {
	Amount   int64
	Currency string
	Rounded  bool
}

//go:generate generate-composite-types -structs Complete -sql -array
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

//go:generate generate-composite-types -structs Times -sql -array
type Times struct {
	CustomTime  time.Time
	CustomTimep *time.Time
}

//go:generate generate-composite-types -alias Int64Alias
type Int64Alias int64
