//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IWkCookieDate Root Interface
type IWkCookieDate interface {
	IObject
	Data() PSoupDate                            // function
	Year() int32                                // function
	Month() int32                               // function
	Day() int32                                 // function
	Hour() int32                                // function
	Minute() int32                              // function
	Second() int32                              // function
	Utc() bool                                  // function
	Offset() int32                              // function
	ToStringTime(format TSoupDateFormat) string // function
	ToLongTime() (resultInt64 int64)            // function
}

// TWkCookieDate Root Object
type TWkCookieDate struct {
	TObject
}

func NewWkCookieDate(aDate PSoupDate) IWkCookieDate {
	r1 := wkCookieDateImportAPI().SysCallN(0, uintptr(aDate))
	return AsWkCookieDate(r1)
}

// WkCookieDateRef -> IWkCookieDate
var WkCookieDateRef wkCookieDate

// wkCookieDate TWkCookieDate Ref
type wkCookieDate uintptr

func (m *wkCookieDate) NewCustomCookieDate(aYear, aMonth, aDay, aHour, aMinute, aSecond int32) IWkCookieDate {
	r1 := wkCookieDateImportAPI().SysCallN(7, uintptr(aYear), uintptr(aMonth), uintptr(aDay), uintptr(aHour), uintptr(aMinute), uintptr(aSecond))
	return AsWkCookieDate(r1)
}

func (m *wkCookieDate) NewCookieDate() IWkCookieDate {
	r1 := wkCookieDateImportAPI().SysCallN(6)
	return AsWkCookieDate(r1)
}

func (m *TWkCookieDate) Data() PSoupDate {
	r1 := wkCookieDateImportAPI().SysCallN(1, m.Instance())
	return PSoupDate(r1)
}

func (m *TWkCookieDate) Year() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Month() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Day() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Hour() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Minute() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Second() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) Utc() bool {
	r1 := wkCookieDateImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TWkCookieDate) Offset() int32 {
	r1 := wkCookieDateImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func (m *TWkCookieDate) ToStringTime(format TSoupDateFormat) string {
	r1 := wkCookieDateImportAPI().SysCallN(11, m.Instance(), uintptr(format))
	return GoStr(r1)
}

func (m *TWkCookieDate) ToLongTime() (resultInt64 int64) {
	wkCookieDateImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

var (
	wkCookieDateImport       *imports.Imports = nil
	wkCookieDateImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkCookieDate_Create", 0),
		/*1*/ imports.NewTable("WkCookieDate_Data", 0),
		/*2*/ imports.NewTable("WkCookieDate_Day", 0),
		/*3*/ imports.NewTable("WkCookieDate_Hour", 0),
		/*4*/ imports.NewTable("WkCookieDate_Minute", 0),
		/*5*/ imports.NewTable("WkCookieDate_Month", 0),
		/*6*/ imports.NewTable("WkCookieDate_NewCookieDate", 0),
		/*7*/ imports.NewTable("WkCookieDate_NewCustomCookieDate", 0),
		/*8*/ imports.NewTable("WkCookieDate_Offset", 0),
		/*9*/ imports.NewTable("WkCookieDate_Second", 0),
		/*10*/ imports.NewTable("WkCookieDate_ToLongTime", 0),
		/*11*/ imports.NewTable("WkCookieDate_ToStringTime", 0),
		/*12*/ imports.NewTable("WkCookieDate_Utc", 0),
		/*13*/ imports.NewTable("WkCookieDate_Year", 0),
	}
)

func wkCookieDateImportAPI() *imports.Imports {
	if wkCookieDateImport == nil {
		wkCookieDateImport = NewDefaultImports()
		wkCookieDateImport.SetImportTable(wkCookieDateImportTables)
		wkCookieDateImportTables = nil
	}
	return wkCookieDateImport
}
