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
//
//	An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
//	https://developer.apple.com/documentation/webkit/wknavigationresponse?language=objc
type IWKNavigationResponse interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKNavigationResponse // function
	// IsForMainFrame
	//  A Boolean value that indicates whether the response targets the web view’s main frame.
	IsForMainFrame() bool // function
	// Response
	//  The frame’s response.
	Response() NSURLResponse // function
	// CanShowMIMEType
	//  A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
	CanShowMIMEType() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKNavigationResponse Root Object
//
//	An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
//	https://developer.apple.com/documentation/webkit/wknavigationresponse?language=objc
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

// New
//
//	Creates and returns an WKNavigationResponse object.
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
