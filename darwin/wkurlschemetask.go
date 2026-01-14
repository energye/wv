//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkURLSchemeTask Parent: lcl.IObject
type IWkURLSchemeTask interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKURLSchemeTask // function
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
	ReceiveData(dataBytes uintptr, length int32) // procedure
	// Finish
	//  Signals the successful completion of the task.
	Finish() // procedure
	// FailWithError
	//  Completes the task and reports the specified error back to WebKit.
	FailWithError(code int32, error_ string) // procedure
}

type TWkURLSchemeTask struct {
	lcl.TObject
}

func (m *TWkURLSchemeTask) Data() wvTypes.WKURLSchemeTask {
	if !m.IsValid() {
		return 0
	}
	r := wkURLSchemeTaskAPI().SysCallN(1, m.Instance())
	return wvTypes.WKURLSchemeTask(r)
}

func (m *TWkURLSchemeTask) Request() NSURLRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkURLSchemeTaskAPI().SysCallN(2, m.Instance())
	return NSURLRequest(r)
}

func (m *TWkURLSchemeTask) Release() {
	if !m.IsValid() {
		return
	}
	wkURLSchemeTaskAPI().SysCallN(3, m.Instance())
}

func (m *TWkURLSchemeTask) ReceiveResponse(response NSURLResponse) {
	if !m.IsValid() {
		return
	}
	wkURLSchemeTaskAPI().SysCallN(4, m.Instance(), uintptr(response))
}

func (m *TWkURLSchemeTask) ReceiveData(dataBytes uintptr, length int32) {
	if !m.IsValid() {
		return
	}
	wkURLSchemeTaskAPI().SysCallN(5, m.Instance(), uintptr(dataBytes), uintptr(length))
}

func (m *TWkURLSchemeTask) Finish() {
	if !m.IsValid() {
		return
	}
	wkURLSchemeTaskAPI().SysCallN(6, m.Instance())
}

func (m *TWkURLSchemeTask) FailWithError(code int32, error_ string) {
	if !m.IsValid() {
		return
	}
	wkURLSchemeTaskAPI().SysCallN(7, m.Instance(), uintptr(code), api.PasStr(error_))
}

// NewURLSchemeTask class constructor
func NewURLSchemeTask(data wvTypes.WKURLSchemeTask) IWkURLSchemeTask {
	r := wkURLSchemeTaskAPI().SysCallN(0, uintptr(data))
	return AsWkURLSchemeTask(r)
}

var (
	wkURLSchemeTaskOnce   base.Once
	wkURLSchemeTaskImport *imports.Imports = nil
)

func wkURLSchemeTaskAPI() *imports.Imports {
	wkURLSchemeTaskOnce.Do(func() {
		wkURLSchemeTaskImport = api.NewDefaultImports()
		wkURLSchemeTaskImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkURLSchemeTask_Create", 0), // constructor NewURLSchemeTask
			/* 1 */ imports.NewTable("TWkURLSchemeTask_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkURLSchemeTask_Request", 0), // function Request
			/* 3 */ imports.NewTable("TWkURLSchemeTask_Release", 0), // procedure Release
			/* 4 */ imports.NewTable("TWkURLSchemeTask_ReceiveResponse", 0), // procedure ReceiveResponse
			/* 5 */ imports.NewTable("TWkURLSchemeTask_ReceiveData", 0), // procedure ReceiveData
			/* 6 */ imports.NewTable("TWkURLSchemeTask_Finish", 0), // procedure Finish
			/* 7 */ imports.NewTable("TWkURLSchemeTask_FailWithError", 0), // procedure FailWithError
		}
	})
	return wkURLSchemeTaskImport
}
