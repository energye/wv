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

// ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of
//	getting the browser Extensions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profilegetbrowserextensionscompletedhandler">See the ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler article.</a>
type ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler interface {
	IObject
}

// TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of
//	getting the browser Extensions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profilegetbrowserextensionscompletedhandler">See the ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler article.</a>
type TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler struct {
	TObject
}

func NewCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	r1 := coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(r1)
}

var (
	coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport       *imports.Imports = nil
	coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ProfileGetBrowserExtensionsCompletedHandler_Create", 0),
	}
)

func coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport == nil {
		coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport = NewDefaultImports()
		coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport.SetImportTable(coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImportTables)
		coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImportTables = nil
	}
	return coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport
}
