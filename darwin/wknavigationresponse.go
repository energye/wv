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

// IWKNavigationResponse Root Interface
type IWKNavigationResponse interface {
	IObject
	Data() WKNavigationResponse // function
	IsForMainFrame() bool       // function
	Response() NSURLResponse    // function
	CanShowMIMEType() bool      // function
	// Release
	//  Release the current object and Data pointer
	Release() // procedure
}

// TWKNavigationResponse Root Object
type TWKNavigationResponse struct {
	TObject
}

func NewWKNavigationResponse(aData WKNavigationResponse) IWKNavigationResponse {
	r1 := wKNavigationResponseImportAPI().SysCallN(1, uintptr(aData))
	return AsWKNavigationResponse(r1)
}

// WKNavigationResponseRef -> IWKNavigationResponse
var WKNavigationResponseRef wKNavigationResponse

// wKNavigationResponse TWKNavigationResponse Ref
type wKNavigationResponse uintptr

func (m *wKNavigationResponse) New() IWKNavigationResponse {
	r1 := wKNavigationResponseImportAPI().SysCallN(4)
	return AsWKNavigationResponse(r1)
}

func (m *TWKNavigationResponse) Data() WKNavigationResponse {
	r1 := wKNavigationResponseImportAPI().SysCallN(2, m.Instance())
	return WKNavigationResponse(r1)
}

func (m *TWKNavigationResponse) IsForMainFrame() bool {
	r1 := wKNavigationResponseImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TWKNavigationResponse) Response() NSURLResponse {
	r1 := wKNavigationResponseImportAPI().SysCallN(6, m.Instance())
	return NSURLResponse(r1)
}

func (m *TWKNavigationResponse) CanShowMIMEType() bool {
	r1 := wKNavigationResponseImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TWKNavigationResponse) Release() {
	wKNavigationResponseImportAPI().SysCallN(5, m.Instance())
}

var (
	wKNavigationResponseImport       *imports.Imports = nil
	wKNavigationResponseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKNavigationResponse_CanShowMIMEType", 0),
		/*1*/ imports.NewTable("WKNavigationResponse_Create", 0),
		/*2*/ imports.NewTable("WKNavigationResponse_Data", 0),
		/*3*/ imports.NewTable("WKNavigationResponse_IsForMainFrame", 0),
		/*4*/ imports.NewTable("WKNavigationResponse_New", 0),
		/*5*/ imports.NewTable("WKNavigationResponse_Release", 0),
		/*6*/ imports.NewTable("WKNavigationResponse_Response", 0),
	}
)

func wKNavigationResponseImportAPI() *imports.Imports {
	if wKNavigationResponseImport == nil {
		wKNavigationResponseImport = NewDefaultImports()
		wKNavigationResponseImport.SetImportTable(wKNavigationResponseImportTables)
		wKNavigationResponseImportTables = nil
	}
	return wKNavigationResponseImport
}
