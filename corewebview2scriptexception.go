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

// ICoreWebView2ScriptException Parent: IObject
//
//	This interface represents a JavaScript exception.
//	If the CoreWebView2.ExecuteScriptWithResult result has Succeeded as false,
//	you can use the result's Exception property to get the script exception.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptexception">See the ICoreWebView2ScriptException article.</a>
type ICoreWebView2ScriptException interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScriptException // property
	// LineNumber
	//  The line number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.lineNumber`.
	//  Note that this position starts at 0.
	LineNumber() uint32 // property
	// ColumnNumber
	//  The column number of the source where the exception occurred.
	//  In the JSON it is `exceptionDetail.columnNumber`.
	//  Note that this position starts at 0.
	ColumnNumber() uint32 // property
	// Name
	//  The Name is the exception's class name.
	//  In the JSON it is `exceptionDetail.exception.className`.
	//  This is the empty string if the exception doesn't have a class name.
	//  This can happen if the script throws a non-Error object such as `throw "abc";`
	Name() string // property
	// Message
	//  The Message is the exception's message and potentially stack.
	//  In the JSON it is exceptionDetail.exception.description.
	//  This is the empty string if the exception doesn't have a description.
	//  This can happen if the script throws a non-Error object such as throw "abc";.
	Message() string // property
	// ToJson
	//  This will return all details of the exception as a JSON string.
	//  In the case that script has thrown a non-Error object such as `throw "abc";`
	//  or any other non-Error object, you can get object specific properties.
	ToJson() string // function
}

// TCoreWebView2ScriptException Parent: TObject
//
//	This interface represents a JavaScript exception.
//	If the CoreWebView2.ExecuteScriptWithResult result has Succeeded as false,
//	you can use the result's Exception property to get the script exception.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptexception">See the ICoreWebView2ScriptException article.</a>
type TCoreWebView2ScriptException struct {
	TObject
}

func NewCoreWebView2ScriptException(aBaseIntf ICoreWebView2ScriptException) ICoreWebView2ScriptException {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ScriptException(r1)
}

func (m *TCoreWebView2ScriptException) Initialized() bool {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ScriptException) BaseIntf() ICoreWebView2ScriptException {
	var resultCoreWebView2ScriptException uintptr
	coreWebView2ScriptExceptionImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ScriptException)))
	return AsCoreWebView2ScriptException(resultCoreWebView2ScriptException)
}

func (m *TCoreWebView2ScriptException) LineNumber() uint32 {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(4, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ScriptException) ColumnNumber() uint32 {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(1, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ScriptException) Name() string {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptException) Message() string {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptException) ToJson() string {
	r1 := coreWebView2ScriptExceptionImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

var (
	coreWebView2ScriptExceptionImport       *imports.Imports = nil
	coreWebView2ScriptExceptionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ScriptException_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ScriptException_ColumnNumber", 0),
		/*2*/ imports.NewTable("CoreWebView2ScriptException_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2ScriptException_Initialized", 0),
		/*4*/ imports.NewTable("CoreWebView2ScriptException_LineNumber", 0),
		/*5*/ imports.NewTable("CoreWebView2ScriptException_Message", 0),
		/*6*/ imports.NewTable("CoreWebView2ScriptException_Name", 0),
		/*7*/ imports.NewTable("CoreWebView2ScriptException_ToJson", 0),
	}
)

func coreWebView2ScriptExceptionImportAPI() *imports.Imports {
	if coreWebView2ScriptExceptionImport == nil {
		coreWebView2ScriptExceptionImport = NewDefaultImports()
		coreWebView2ScriptExceptionImport.SetImportTable(coreWebView2ScriptExceptionImportTables)
		coreWebView2ScriptExceptionImportTables = nil
	}
	return coreWebView2ScriptExceptionImport
}
