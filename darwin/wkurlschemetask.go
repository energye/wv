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

// IWKURLSchemeTask Root Interface
type IWKURLSchemeTask interface {
	IObject
	Data() WKURLSchemeTask                  // function
	Request() NSURLRequest                  // function
	ReceiveResponse(response NSURLResponse) // procedure
	// ReceiveData
	//  NSData
	ReceiveData(aDataBytes uintptr, aLength int32) // procedure
	Finish()                                       // procedure
	// FailWithError
	//  NSError
	FailWithError(error string) // procedure
}

// TWKURLSchemeTask Root Object
type TWKURLSchemeTask struct {
	TObject
}

func NewWKURLSchemeTask(aData WKURLSchemeTask) IWKURLSchemeTask {
	r1 := wKURLSchemeTaskImportAPI().SysCallN(0, uintptr(aData))
	return AsWKURLSchemeTask(r1)
}

func (m *TWKURLSchemeTask) Data() WKURLSchemeTask {
	r1 := wKURLSchemeTaskImportAPI().SysCallN(1, m.Instance())
	return WKURLSchemeTask(r1)
}

func (m *TWKURLSchemeTask) Request() NSURLRequest {
	r1 := wKURLSchemeTaskImportAPI().SysCallN(6, m.Instance())
	return NSURLRequest(r1)
}

func (m *TWKURLSchemeTask) ReceiveResponse(response NSURLResponse) {
	wKURLSchemeTaskImportAPI().SysCallN(5, m.Instance(), uintptr(response))
}

func (m *TWKURLSchemeTask) ReceiveData(aDataBytes uintptr, aLength int32) {
	wKURLSchemeTaskImportAPI().SysCallN(4, m.Instance(), uintptr(aDataBytes), uintptr(aLength))
}

func (m *TWKURLSchemeTask) Finish() {
	wKURLSchemeTaskImportAPI().SysCallN(3, m.Instance())
}

func (m *TWKURLSchemeTask) FailWithError(error string) {
	wKURLSchemeTaskImportAPI().SysCallN(2, m.Instance(), PascalStr(error))
}

var (
	wKURLSchemeTaskImport       *imports.Imports = nil
	wKURLSchemeTaskImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKURLSchemeTask_Create", 0),
		/*1*/ imports.NewTable("WKURLSchemeTask_Data", 0),
		/*2*/ imports.NewTable("WKURLSchemeTask_FailWithError", 0),
		/*3*/ imports.NewTable("WKURLSchemeTask_Finish", 0),
		/*4*/ imports.NewTable("WKURLSchemeTask_ReceiveData", 0),
		/*5*/ imports.NewTable("WKURLSchemeTask_ReceiveResponse", 0),
		/*6*/ imports.NewTable("WKURLSchemeTask_Request", 0),
	}
)

func wKURLSchemeTaskImportAPI() *imports.Imports {
	if wKURLSchemeTaskImport == nil {
		wKURLSchemeTaskImport = NewDefaultImports()
		wKURLSchemeTaskImport.SetImportTable(wKURLSchemeTaskImportTables)
		wKURLSchemeTaskImportTables = nil
	}
	return wKURLSchemeTaskImport
}
