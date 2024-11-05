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

// ICoreWebView2GetCookiesCompletedHandler Parent: IObject
//
//	Receives the result of the GetCookies method.  The result is written to
//	the cookie list provided in the GetCookies method call.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getcookiescompletedhandler">See the ICoreWebView2GetCookiesCompletedHandler article.</a>
type ICoreWebView2GetCookiesCompletedHandler interface {
	IObject
}

// TCoreWebView2GetCookiesCompletedHandler Parent: TObject
//
//	Receives the result of the GetCookies method.  The result is written to
//	the cookie list provided in the GetCookies method call.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getcookiescompletedhandler">See the ICoreWebView2GetCookiesCompletedHandler article.</a>
type TCoreWebView2GetCookiesCompletedHandler struct {
	TObject
}

func NewCoreWebView2GetCookiesCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2GetCookiesCompletedHandler {
	r1 := coreWebView2GetCookiesCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2GetCookiesCompletedHandler(r1)
}

var (
	coreWebView2GetCookiesCompletedHandlerImport       *imports.Imports = nil
	coreWebView2GetCookiesCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2GetCookiesCompletedHandler_Create", 0),
	}
)

func coreWebView2GetCookiesCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2GetCookiesCompletedHandlerImport == nil {
		coreWebView2GetCookiesCompletedHandlerImport = NewDefaultImports()
		coreWebView2GetCookiesCompletedHandlerImport.SetImportTable(coreWebView2GetCookiesCompletedHandlerImportTables)
		coreWebView2GetCookiesCompletedHandlerImportTables = nil
	}
	return coreWebView2GetCookiesCompletedHandlerImport
}
