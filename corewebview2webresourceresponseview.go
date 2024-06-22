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

// ICoreWebView2WebResourceResponseView Parent: IObject
//
//	View of the HTTP representation for a web resource response. The properties
//	of this object are not mutable. This response view is used with the
//	TWVBrowserBase.OnWebResourceResponseReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview">See the ICoreWebView2WebResourceResponseView article.</a>
type ICoreWebView2WebResourceResponseView interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponseView // property
	// StatusCode
	//  The HTTP response status code.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_statuscode">See the ICoreWebView2WebResourceResponseView article.</a>
	StatusCode() int32 // property
	// ReasonPhrase
	//  The HTTP response reason phrase.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_reasonphrase">See the ICoreWebView2WebResourceResponseView article.</a>
	ReasonPhrase() string // property
	// Headers
	//  The HTTP response headers as received.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_headers">See the ICoreWebView2WebResourceResponseView article.</a>
	Headers() ICoreWebView2HttpResponseHeaders // property
	// GetContent
	//  Get the response content asynchronously. The handler will receive the
	//  response content stream.
	//  This method returns null if content size is more than 123MB or for navigations that become downloads
	//  or if response is downloadable content type(e.g., application/octet-stream).
	//  See `add_DownloadStarting` event to handle the response.
	//  If this method is being called again before a first call has completed,
	//  the handler will be invoked at the same time the handlers from prior calls
	//  are invoked.
	//  If this method is being called after a first call has completed, the
	//  handler will be invoked immediately.
	//  TCoreWebView2WebResourceResponseView.GetContent will trigger the
	//  TWVBrowserBase.OnWebResourceResponseViewGetContentCompleted event with the resource contents.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#getcontent">See the ICoreWebView2WebResourceResponseView article.</a>
	GetContent(aHandler ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) bool // function
}

// TCoreWebView2WebResourceResponseView Parent: TObject
//
//	View of the HTTP representation for a web resource response. The properties
//	of this object are not mutable. This response view is used with the
//	TWVBrowserBase.OnWebResourceResponseReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview">See the ICoreWebView2WebResourceResponseView article.</a>
type TCoreWebView2WebResourceResponseView struct {
	TObject
}

func NewCoreWebView2WebResourceResponseView(aBaseIntf ICoreWebView2WebResourceResponseView) ICoreWebView2WebResourceResponseView {
	r1 := coreWebView2WebResourceResponseViewImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2WebResourceResponseView(r1)
}

func (m *TCoreWebView2WebResourceResponseView) Initialized() bool {
	r1 := coreWebView2WebResourceResponseViewImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebResourceResponseView) BaseIntf() ICoreWebView2WebResourceResponseView {
	var resultCoreWebView2WebResourceResponseView uintptr
	coreWebView2WebResourceResponseViewImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceResponseView)))
	return AsCoreWebView2WebResourceResponseView(resultCoreWebView2WebResourceResponseView)
}

func (m *TCoreWebView2WebResourceResponseView) StatusCode() int32 {
	r1 := coreWebView2WebResourceResponseViewImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2WebResourceResponseView) ReasonPhrase() string {
	r1 := coreWebView2WebResourceResponseViewImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2WebResourceResponseView) Headers() ICoreWebView2HttpResponseHeaders {
	var resultCoreWebView2HttpResponseHeaders uintptr
	coreWebView2WebResourceResponseViewImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpResponseHeaders)))
	return AsCoreWebView2HttpResponseHeaders(resultCoreWebView2HttpResponseHeaders)
}

func (m *TCoreWebView2WebResourceResponseView) GetContent(aHandler ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) bool {
	r1 := coreWebView2WebResourceResponseViewImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

var (
	coreWebView2WebResourceResponseViewImport       *imports.Imports = nil
	coreWebView2WebResourceResponseViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2WebResourceResponseView_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2WebResourceResponseView_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2WebResourceResponseView_GetContent", 0),
		/*3*/ imports.NewTable("CoreWebView2WebResourceResponseView_Headers", 0),
		/*4*/ imports.NewTable("CoreWebView2WebResourceResponseView_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2WebResourceResponseView_ReasonPhrase", 0),
		/*6*/ imports.NewTable("CoreWebView2WebResourceResponseView_StatusCode", 0),
	}
)

func coreWebView2WebResourceResponseViewImportAPI() *imports.Imports {
	if coreWebView2WebResourceResponseViewImport == nil {
		coreWebView2WebResourceResponseViewImport = NewDefaultImports()
		coreWebView2WebResourceResponseViewImport.SetImportTable(coreWebView2WebResourceResponseViewImportTables)
		coreWebView2WebResourceResponseViewImportTables = nil
	}
	return coreWebView2WebResourceResponseViewImport
}
