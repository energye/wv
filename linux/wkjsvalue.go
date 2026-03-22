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

// IWkJSValue Parent: IObject
type IWkJSValue interface {
	IObject
	StringValue() string                // function
	NumberValue() float64               // function
	IntegerValue() int32                // function
	BooleanValue() bool                 // function
	ExceptionMessage() string           // function
	ValueType() wvTypes.UWJSValueJSType // property ValueType Getter
}

type TWkJSValue struct {
	TObject
}

func (m *TWkJSValue) StringValue() string {
	if !m.IsValid() {
		return ""
	}
	r := wkJSValueAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TWkJSValue) NumberValue() (result float64) {
	if !m.IsValid() {
		return
	}
	wkJSValueAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkJSValue) IntegerValue() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkJSValueAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TWkJSValue) BooleanValue() bool {
	if !m.IsValid() {
		return false
	}
	r := wkJSValueAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkJSValue) ExceptionMessage() string {
	if !m.IsValid() {
		return ""
	}
	r := wkJSValueAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TWkJSValue) ValueType() wvTypes.UWJSValueJSType {
	if !m.IsValid() {
		return 0
	}
	r := wkJSValueAPI().SysCallN(7, m.Instance())
	return wvTypes.UWJSValueJSType(r)
}

// NewJSValue class constructor
func NewJSValue(result wvTypes.WebKitJavascriptResult, isException bool, exception string) IWkJSValue {
	r := wkJSValueAPI().SysCallN(0, uintptr(result), api.PasBool(isException), api.PasStr(exception))
	return AsWkJSValue(r)
}

// NewJSValueWithWebKitJavascriptResult class constructor
func NewJSValueWithWebKitJavascriptResult(result wvTypes.WebKitJavascriptResult) IWkJSValue {
	r := wkJSValueAPI().SysCallN(1, uintptr(result))
	return AsWkJSValue(r)
}

var (
	wkJSValueOnce   base.Once
	wkJSValueImport *imports.Imports = nil
)

func wkJSValueAPI() *imports.Imports {
	wkJSValueOnce.Do(func() {
		wkJSValueImport = api.NewDefaultImports()
		wkJSValueImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkJSValue_Create", 0), // constructor NewJSValue
			/* 1 */ imports.NewTable("TWkJSValue_CreateWithWebKitJavascriptResult", 0), // constructor NewJSValueWithWebKitJavascriptResult
			/* 2 */ imports.NewTable("TWkJSValue_StringValue", 0), // function StringValue
			/* 3 */ imports.NewTable("TWkJSValue_NumberValue", 0), // function NumberValue
			/* 4 */ imports.NewTable("TWkJSValue_IntegerValue", 0), // function IntegerValue
			/* 5 */ imports.NewTable("TWkJSValue_BooleanValue", 0), // function BooleanValue
			/* 6 */ imports.NewTable("TWkJSValue_ExceptionMessage", 0), // function ExceptionMessage
			/* 7 */ imports.NewTable("TWkJSValue_ValueType", 0), // property ValueType
		}
	})
	return wkJSValueImport
}
