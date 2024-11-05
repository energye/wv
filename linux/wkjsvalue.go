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

// IWkJSValue Root Interface
type IWkJSValue interface {
	IObject
	ValueType() JSType                    // property
	StringValue() string                  // function
	NumberValue() (resultFloat64 float64) // function
	IntegerValue() int32                  // function
	BooleanValue() bool                   // function
	ExceptionMessage() string             // function
}

// TWkJSValue Root Object
type TWkJSValue struct {
	TObject
}

func NewWkJSValue(aResult WebKitJavascriptResult, aIsException bool, exception string) IWkJSValue {
	r1 := wkJSValueImportAPI().SysCallN(1, uintptr(aResult), PascalBool(aIsException), PascalStr(exception))
	return AsWkJSValue(r1)
}

func NewWkJSValue1(aResult WebKitJavascriptResult) IWkJSValue {
	r1 := wkJSValueImportAPI().SysCallN(2, uintptr(aResult))
	return AsWkJSValue(r1)
}

func (m *TWkJSValue) ValueType() JSType {
	r1 := wkJSValueImportAPI().SysCallN(7, m.Instance())
	return JSType(r1)
}

func (m *TWkJSValue) StringValue() string {
	r1 := wkJSValueImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TWkJSValue) NumberValue() (resultFloat64 float64) {
	wkJSValueImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TWkJSValue) IntegerValue() int32 {
	r1 := wkJSValueImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TWkJSValue) BooleanValue() bool {
	r1 := wkJSValueImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TWkJSValue) ExceptionMessage() string {
	r1 := wkJSValueImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

var (
	wkJSValueImport       *imports.Imports = nil
	wkJSValueImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkJSValue_BooleanValue", 0),
		/*1*/ imports.NewTable("WkJSValue_Create", 0),
		/*2*/ imports.NewTable("WkJSValue_Create1", 0),
		/*3*/ imports.NewTable("WkJSValue_ExceptionMessage", 0),
		/*4*/ imports.NewTable("WkJSValue_IntegerValue", 0),
		/*5*/ imports.NewTable("WkJSValue_NumberValue", 0),
		/*6*/ imports.NewTable("WkJSValue_StringValue", 0),
		/*7*/ imports.NewTable("WkJSValue_ValueType", 0),
	}
)

func wkJSValueImportAPI() *imports.Imports {
	if wkJSValueImport == nil {
		wkJSValueImport = NewDefaultImports()
		wkJSValueImport.SetImportTable(wkJSValueImportTables)
		wkJSValueImportTables = nil
	}
	return wkJSValueImport
}
