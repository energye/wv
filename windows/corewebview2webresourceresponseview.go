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

// ICoreWebView2WebResourceResponseView Parent: IObject
type ICoreWebView2WebResourceResponseView interface {
	IObject
	// GetContent
	//  Get the response content asynchronously. The handler will receive the
	//  response content stream.
	//  This method returns null if content size is more than 123MB or for navigations that become downloads
	//  or if response is downloadable content type (e.g., application/octet-stream).
	//  See `add_DownloadStarting` event to handle the response.
	//  If this method is being called again before a first call has completed,
	//  the handler will be invoked at the same time the handlers from prior calls
	//  are invoked.
	//  If this method is being called after a first call has completed, the
	//  handler will be invoked immediately.
	//  TCoreWebView2WebResourceResponseView.GetContent will trigger the
	//  TWVBrowserBase.OnWebResourceResponseViewGetContentCompleted event with the resource contents.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#getcontent">See the ICoreWebView2WebResourceResponseView article.</see>
	GetContent(handler ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponseView // property BaseIntf Getter
	// StatusCode
	//  The HTTP response status code.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_statuscode">See the ICoreWebView2WebResourceResponseView article.</see>
	StatusCode() int32 // property StatusCode Getter
	// ReasonPhrase
	//  The HTTP response reason phrase.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_reasonphrase">See the ICoreWebView2WebResourceResponseView article.</see>
	ReasonPhrase() string // property ReasonPhrase Getter
	// Headers
	//  The HTTP response headers as received.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#get_headers">See the ICoreWebView2WebResourceResponseView article.</see>
	Headers() ICoreWebView2HttpResponseHeaders // property Headers Getter
}

type TCoreWebView2WebResourceResponseView struct {
	TObject
}

func (m *TCoreWebView2WebResourceResponseView) GetContent(handler ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceResponseViewAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceResponseView) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceResponseViewAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceResponseView) BaseIntf() (result ICoreWebView2WebResourceResponseView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseViewAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceResponseView(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceResponseView) StatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceResponseViewAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2WebResourceResponseView) ReasonPhrase() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebResourceResponseViewAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebResourceResponseView) Headers() (result ICoreWebView2HttpResponseHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseViewAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpResponseHeaders(resultPtr)
	return
}

// NewCoreWebView2WebResourceResponseView class constructor
func NewCoreWebView2WebResourceResponseView(baseIntf ICoreWebView2WebResourceResponseView) ICoreWebView2WebResourceResponseView {
	r := coreWebView2WebResourceResponseViewAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2WebResourceResponseView(r)
}

var (
	coreWebView2WebResourceResponseViewOnce   base.Once
	coreWebView2WebResourceResponseViewImport *imports.Imports = nil
)

func coreWebView2WebResourceResponseViewAPI() *imports.Imports {
	coreWebView2WebResourceResponseViewOnce.Do(func() {
		coreWebView2WebResourceResponseViewImport = api.NewDefaultImports()
		coreWebView2WebResourceResponseViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceResponseView_Create", 0), // constructor NewCoreWebView2WebResourceResponseView
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceResponseView_GetContent", 0), // function GetContent
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceResponseView_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceResponseView_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceResponseView_StatusCode", 0), // property StatusCode
			/* 5 */ imports.NewTable("TCoreWebView2WebResourceResponseView_ReasonPhrase", 0), // property ReasonPhrase
			/* 6 */ imports.NewTable("TCoreWebView2WebResourceResponseView_Headers", 0), // property Headers
		}
	})
	return coreWebView2WebResourceResponseViewImport
}
