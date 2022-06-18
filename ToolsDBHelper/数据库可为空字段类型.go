package ToolsDBHelper

import (
	"database/sql/driver"
	"time"
)

type NullInt struct {
	Int   int
	Valid bool // Valid is true if Int is not NULL
}

// Scan 实现它的赋值方法(注意，这个方属于指针)
func (nt *NullInt) Scan(value interface{}) error {
	nt.Int, nt.Valid = value.(int)
	return nil
}

// Value 实现它的取值方式
func (nt *NullInt) Value() (driver.Value, error) {
	if !nt.Valid {
		return 0, nil
	}
	return nt.Int, nil
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan 实现它的赋值方法(注意，这个方属于指针)
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value 实现它的取值方式
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return 0, nil
	}
	return nt.Time, nil
}
