//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkCookieDate Parent: IObject
type IWkCookieDate interface {
	IObject
	Data() wvTypes.PSoupDate                            // function
	Year() int32                                        // function
	Month() int32                                       // function
	Day() int32                                         // function
	Hour() int32                                        // function
	Minute() int32                                      // function
	Second() int32                                      // function
	Utc() bool                                          // function
	Offset() int32                                      // function
	ToStringTime(format wvTypes.TSoupDateFormat) string // function
	ToLongTime() int64                                  // function
}

type TWkCookieDate struct {
	TObject
}

func (m *TWkCookieDate) Data() wvTypes.PSoupDate {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(1, m.Instance())
	return wvTypes.PSoupDate(r)
}

func (m *TWkCookieDate) Year() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Month() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Day() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Hour() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Minute() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Second() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) Utc() bool {
	if !m.IsValid() {
		return false
	}
	r := wkCookieDateAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TWkCookieDate) Offset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieDateAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TWkCookieDate) ToStringTime(format wvTypes.TSoupDateFormat) string {
	if !m.IsValid() {
		return ""
	}
	r := wkCookieDateAPI().SysCallN(10, m.Instance(), uintptr(format))
	return api.GoStr(r)
}

func (m *TWkCookieDate) ToLongTime() (result int64) {
	if !m.IsValid() {
		return
	}
	wkCookieDateAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

// CookieDate  is static instance
var CookieDate _CookieDateClass

// _CookieDateClass is class type defined by TWkCookieDate
type _CookieDateClass uintptr

func (_CookieDateClass) NewCustomCookieDate(year int32, month int32, day int32, hour int32, minute int32, second int32) IWkCookieDate {
	r := wkCookieDateAPI().SysCallN(12, uintptr(year), uintptr(month), uintptr(day), uintptr(hour), uintptr(minute), uintptr(second))
	return AsWkCookieDate(r)
}

func (_CookieDateClass) NewCookieDate() IWkCookieDate {
	r := wkCookieDateAPI().SysCallN(13)
	return AsWkCookieDate(r)
}

// NewCookieDate class constructor
func NewCookieDate(date wvTypes.PSoupDate) IWkCookieDate {
	r := wkCookieDateAPI().SysCallN(0, uintptr(date))
	return AsWkCookieDate(r)
}

var (
	wkCookieDateOnce   base.Once
	wkCookieDateImport *imports.Imports = nil
)

func wkCookieDateAPI() *imports.Imports {
	wkCookieDateOnce.Do(func() {
		wkCookieDateImport = api.NewDefaultImports()
		wkCookieDateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkCookieDate_Create", 0), // constructor NewCookieDate
			/* 1 */ imports.NewTable("TWkCookieDate_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkCookieDate_Year", 0), // function Year
			/* 3 */ imports.NewTable("TWkCookieDate_Month", 0), // function Month
			/* 4 */ imports.NewTable("TWkCookieDate_Day", 0), // function Day
			/* 5 */ imports.NewTable("TWkCookieDate_Hour", 0), // function Hour
			/* 6 */ imports.NewTable("TWkCookieDate_Minute", 0), // function Minute
			/* 7 */ imports.NewTable("TWkCookieDate_Second", 0), // function Second
			/* 8 */ imports.NewTable("TWkCookieDate_Utc", 0), // function Utc
			/* 9 */ imports.NewTable("TWkCookieDate_Offset", 0), // function Offset
			/* 10 */ imports.NewTable("TWkCookieDate_ToStringTime", 0), // function ToStringTime
			/* 11 */ imports.NewTable("TWkCookieDate_ToLongTime", 0), // function ToLongTime
			/* 12 */ imports.NewTable("TWkCookieDate_NewCustomCookieDate", 0), // static function NewCustomCookieDate
			/* 13 */ imports.NewTable("TWkCookieDate_NewCookieDate", 0), // static function NewCookieDate
		}
	})
	return wkCookieDateImport
}
