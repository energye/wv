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
//
//	An interface that WebKit uses to request custom resources from your app.
//	https://developer.apple.com/documentation/webkit/wkurlschemetask?language=objc
type IWKURLSchemeTask interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKURLSchemeTask // function
	// Request
	//  Information about the resource to load.
	Request() NSURLRequest // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// ReceiveResponse
	//  Returns a URL response to WebKit with information about the requested resource.
	ReceiveResponse(response NSURLResponse) // procedure
	// ReceiveData
	//  Sends some or all of the resource data to WebKit.
	//  NSData
	ReceiveData(aDataBytes uintptr, aLength int32) // procedure
	// Finish
	//  Signals the successful completion of the task.
	Finish() // procedure
	// FailWithError
	//  Completes the task and reports the specified error back to WebKit.
	//  NSError
	FailWithError(error string) // procedure
}

// TWKURLSchemeTask Root Object
//
//	An interface that WebKit uses to request custom resources from your app.
//	https://developer.apple.com/documentation/webkit/wkurlschemetask?language=objc
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
	r1 := wKURLSchemeTaskImportAPI().SysCallN(7, m.Instance())
	return NSURLRequest(r1)
}

func (m *TWKURLSchemeTask) Release() {
	wKURLSchemeTaskImportAPI().SysCallN(6, m.Instance())
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
		/*6*/ imports.NewTable("WKURLSchemeTask_Release", 0),
		/*7*/ imports.NewTable("WKURLSchemeTask_Request", 0),
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
