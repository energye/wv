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

// IWkURIResponse Root Interface
type IWkURIResponse interface {
	IObject
	URI() string                             // property
	Data() WebKitURIResponse                 // function
	Headers() PSoupMessageHeaders            // function
	GetStatusCode() int32                    // function
	GetContentLength() (resultUint64 uint64) // function
	GetMimeType() string                     // function
	GetSuggestedFilename() string            // function
}

// TWkURIResponse Root Object
type TWkURIResponse struct {
	TObject
}

func NewWkURIResponse(aURIResponse WebKitURIResponse) IWkURIResponse {
	r1 := wkURIResponseImportAPI().SysCallN(0, uintptr(aURIResponse))
	return AsWkURIResponse(r1)
}

func (m *TWkURIResponse) URI() string {
	r1 := wkURIResponseImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TWkURIResponse) Data() WebKitURIResponse {
	r1 := wkURIResponseImportAPI().SysCallN(1, m.Instance())
	return WebKitURIResponse(r1)
}

func (m *TWkURIResponse) Headers() PSoupMessageHeaders {
	r1 := wkURIResponseImportAPI().SysCallN(6, m.Instance())
	return PSoupMessageHeaders(r1)
}

func (m *TWkURIResponse) GetStatusCode() int32 {
	r1 := wkURIResponseImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TWkURIResponse) GetContentLength() (resultUint64 uint64) {
	wkURIResponseImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TWkURIResponse) GetMimeType() string {
	r1 := wkURIResponseImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TWkURIResponse) GetSuggestedFilename() string {
	r1 := wkURIResponseImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

var (
	wkURIResponseImport       *imports.Imports = nil
	wkURIResponseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkURIResponse_Create", 0),
		/*1*/ imports.NewTable("WkURIResponse_Data", 0),
		/*2*/ imports.NewTable("WkURIResponse_GetContentLength", 0),
		/*3*/ imports.NewTable("WkURIResponse_GetMimeType", 0),
		/*4*/ imports.NewTable("WkURIResponse_GetStatusCode", 0),
		/*5*/ imports.NewTable("WkURIResponse_GetSuggestedFilename", 0),
		/*6*/ imports.NewTable("WkURIResponse_Headers", 0),
		/*7*/ imports.NewTable("WkURIResponse_URI", 0),
	}
)

func wkURIResponseImportAPI() *imports.Imports {
	if wkURIResponseImport == nil {
		wkURIResponseImport = NewDefaultImports()
		wkURIResponseImport.SetImportTable(wkURIResponseImportTables)
		wkURIResponseImportTables = nil
	}
	return wkURIResponseImport
}
