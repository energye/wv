//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
)

// ICoreWebView2ScriptException Parent: lcl.IObject
type ICoreWebView2ScriptException interface {
	lcl.IObject
	// ToJson
	//  This will return all details of the exception as a JSON string.
	//  In the case that script has thrown a non-Error object such as `throw "abc";`
	//  or any other non-Error object, you can get object specific properties.
	ToJson() string // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScriptException // property BaseIntf Getter
	// LineNumber
	//  The line number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.lineNumber`.
	//  Note that this position starts at 0.
	LineNumber() uint32 // property LineNumber Getter
	// ColumnNumber
	//  The column number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.columnNumber`.
	//  Note that this position starts at 0.
	ColumnNumber() uint32 // property ColumnNumber Getter
	// Name
	//  The Name is the exception's class name.
	//  In the JSON it is `exceptionDetail.exception.className`.
	//  This is the empty string if the exception doesn't have a class name.
	//  This can happen if the script throws a non-Error object such as `throw "abc";`
	Name() string // property Name Getter
	// Message
	//  The Message is the exception's message and potentially stack.
	//  In the JSON it is exceptionDetail.exception.description.
	//  This is the empty string if the exception doesn't have a description.
	//  This can happen if the script throws a non-Error object such as throw "abc";.
	Message() string // property Message_ Getter
}

type TCoreWebView2ScriptException struct {
	lcl.TObject
}

func (m *TCoreWebView2ScriptException) ToJson() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptException) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScriptException) BaseIntf() (result ICoreWebView2ScriptException) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScriptExceptionAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ScriptException(resultPtr)
	return
}

func (m *TCoreWebView2ScriptException) LineNumber() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(4, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ScriptException) ColumnNumber() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(5, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ScriptException) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptException) Message() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptExceptionAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

// NewCoreWebView2ScriptException class constructor
func NewCoreWebView2ScriptException(baseIntf ICoreWebView2ScriptException) ICoreWebView2ScriptException {
	r := coreWebView2ScriptExceptionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ScriptException(r)
}

var (
	coreWebView2ScriptExceptionOnce   base.Once
	coreWebView2ScriptExceptionImport *imports.Imports = nil
)

func coreWebView2ScriptExceptionAPI() *imports.Imports {
	coreWebView2ScriptExceptionOnce.Do(func() {
		coreWebView2ScriptExceptionImport = api.NewDefaultImports()
		coreWebView2ScriptExceptionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ScriptException_Create", 0), // constructor NewCoreWebView2ScriptException
			/* 1 */ imports.NewTable("TCoreWebView2ScriptException_ToJson", 0), // function ToJson
			/* 2 */ imports.NewTable("TCoreWebView2ScriptException_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2ScriptException_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2ScriptException_LineNumber", 0), // property LineNumber
			/* 5 */ imports.NewTable("TCoreWebView2ScriptException_ColumnNumber", 0), // property ColumnNumber
			/* 6 */ imports.NewTable("TCoreWebView2ScriptException_Name", 0), // property Name
			/* 7 */ imports.NewTable("TCoreWebView2ScriptException_Message_", 0), // property Message
		}
	})
	return coreWebView2ScriptExceptionImport
}
