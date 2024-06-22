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

// ICoreWebView2BrowserExtensionEnableCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of setting the
//	browser Extension as enabled or disabled. If enabled, the browser Extension is
//	running in WebView instances. If disabled, the browser Extension is not running in WebView instances.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionenablecompletedhandler">See the ICoreWebView2BrowserExtensionEnableCompletedHandler article.</a>
type ICoreWebView2BrowserExtensionEnableCompletedHandler interface {
	IObject
}

// TCoreWebView2BrowserExtensionEnableCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of setting the
//	browser Extension as enabled or disabled. If enabled, the browser Extension is
//	running in WebView instances. If disabled, the browser Extension is not running in WebView instances.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionenablecompletedhandler">See the ICoreWebView2BrowserExtensionEnableCompletedHandler article.</a>
type TCoreWebView2BrowserExtensionEnableCompletedHandler struct {
	TObject
}

func NewCoreWebView2BrowserExtensionEnableCompletedHandler(aEvents IWVBrowserEvents, aExtensionID string) ICoreWebView2BrowserExtensionEnableCompletedHandler {
	r1 := coreWebView2BrowserExtensionEnableCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents), PascalStr(aExtensionID))
	return AsCoreWebView2BrowserExtensionEnableCompletedHandler(r1)
}

var (
	coreWebView2BrowserExtensionEnableCompletedHandlerImport       *imports.Imports = nil
	coreWebView2BrowserExtensionEnableCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BrowserExtensionEnableCompletedHandler_Create", 0),
	}
)

func coreWebView2BrowserExtensionEnableCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2BrowserExtensionEnableCompletedHandlerImport == nil {
		coreWebView2BrowserExtensionEnableCompletedHandlerImport = NewDefaultImports()
		coreWebView2BrowserExtensionEnableCompletedHandlerImport.SetImportTable(coreWebView2BrowserExtensionEnableCompletedHandlerImportTables)
		coreWebView2BrowserExtensionEnableCompletedHandlerImportTables = nil
	}
	return coreWebView2BrowserExtensionEnableCompletedHandlerImport
}
