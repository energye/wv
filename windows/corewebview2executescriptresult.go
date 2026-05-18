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
)

// ICoreWebView2ExecuteScriptResult Parent: IObject
type ICoreWebView2ExecuteScriptResult interface {
	IObject
	// TryGetResultAsString
	//  If Succeeded is true and the result of script execution is a string, this method provides the value of the string result,
	//  and we will get the `FALSE` var value when the js result is not string type.
	//  The return value description is as follows
	//  1. S_OK: Execution succeeds.
	//  2. E_POINTER: When the `stringResult` or `value` is nullptr.
	//  NOTE: If the `value` returns `FALSE`, the `stringResult` will be set to a empty string.
	TryGetResultAsString(stringResult *string, value *bool) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ExecuteScriptResult // property BaseIntf Getter
	// Succeeded
	//  This property is true if ExecuteScriptWithResult successfully executed script with
	//  no unhandled exceptions and the result is available in the ResultAsJson property
	//  or via the TryGetResultAsString method.
	//  If it is false then the script execution had an unhandled exception which you
	//  can get via the Exception property.
	Succeeded() bool // property Succeeded_ Getter
	// ResultAsJson
	//  A function that has no explicit return value returns undefined. If the
	//  script that was run throws an unhandled exception, then the result is
	//  also "null". This method is applied asynchronously. If the method is
	//  run before `ContentLoading`, the script will not be executed
	//  and the string "null" will be returned.
	//  The return value description is as follows
	//  1. S_OK: Execution succeeds.
	//  2. E_POINTER: When the `jsonResult` is nullptr.
	ResultAsJson() string // property ResultAsJson Getter
	// Exception
	//  If Succeeded is false, you can use this property to get the unhandled exception thrown by script execution
	//  Note that due to the compatibility of the WinRT/.NET interface,
	//  S_OK will be returned even if the acquisition fails.
	//  We can determine whether the acquisition is successful by judging whether the `exception` is nullptr.
	Exception() ICoreWebView2ScriptException // property Exception Getter
}

type TCoreWebView2ExecuteScriptResult struct {
	TObject
}

func (m *TCoreWebView2ExecuteScriptResult) TryGetResultAsString(stringResult *string, value *bool) bool {
	if !m.IsValid() {
		return false
	}
	stringResultPtr := api.PasStr(*stringResult)
	r := coreWebView2ExecuteScriptResultAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&stringResultPtr)), uintptr(base.UnsafePointer(value)))
	*stringResult = api.GoStr(stringResultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2ExecuteScriptResult) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ExecuteScriptResultAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ExecuteScriptResult) BaseIntf() (result ICoreWebView2ExecuteScriptResult) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ExecuteScriptResultAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ExecuteScriptResult(resultPtr)
	return
}

func (m *TCoreWebView2ExecuteScriptResult) Succeeded() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ExecuteScriptResultAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ExecuteScriptResult) ResultAsJson() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ExecuteScriptResultAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ExecuteScriptResult) Exception() (result ICoreWebView2ScriptException) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ExecuteScriptResultAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ScriptException(resultPtr)
	return
}

// NewCoreWebView2ExecuteScriptResult class constructor
func NewCoreWebView2ExecuteScriptResult(baseIntf ICoreWebView2ExecuteScriptResult) ICoreWebView2ExecuteScriptResult {
	r := coreWebView2ExecuteScriptResultAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ExecuteScriptResult(r)
}

var (
	coreWebView2ExecuteScriptResultOnce   base.Once
	coreWebView2ExecuteScriptResultImport *imports.Imports = nil
)

func coreWebView2ExecuteScriptResultAPI() *imports.Imports {
	coreWebView2ExecuteScriptResultOnce.Do(func() {
		coreWebView2ExecuteScriptResultImport = api.NewDefaultImports()
		coreWebView2ExecuteScriptResultImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_Create", 0), // constructor NewCoreWebView2ExecuteScriptResult
			/* 1 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_TryGetResultAsString", 0), // function TryGetResultAsString
			/* 2 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_Succeeded_", 0), // property Succeeded
			/* 5 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_ResultAsJson", 0), // property ResultAsJson
			/* 6 */ imports.NewTable("TCoreWebView2ExecuteScriptResult_Exception", 0), // property Exception
		}
	})
	return coreWebView2ExecuteScriptResultImport
}
