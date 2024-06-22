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

// ICoreWebView2BrowserExtensionRemoveCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result of removing
//	the browser Extension from the Profile.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionremovecompletedhandler">See the ICoreWebView2BrowserExtensionRemoveCompletedHandler article.</a>
type ICoreWebView2BrowserExtensionRemoveCompletedHandler interface {
	IObject
}

// TCoreWebView2BrowserExtensionRemoveCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result of removing
//	the browser Extension from the Profile.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionremovecompletedhandler">See the ICoreWebView2BrowserExtensionRemoveCompletedHandler article.</a>
type TCoreWebView2BrowserExtensionRemoveCompletedHandler struct {
	TObject
}

func NewCoreWebView2BrowserExtensionRemoveCompletedHandler(aEvents IWVBrowserEvents, aExtensionID string) ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	r1 := coreWebView2BrowserExtensionRemoveCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents), PascalStr(aExtensionID))
	return AsCoreWebView2BrowserExtensionRemoveCompletedHandler(r1)
}

var (
	coreWebView2BrowserExtensionRemoveCompletedHandlerImport       *imports.Imports = nil
	coreWebView2BrowserExtensionRemoveCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BrowserExtensionRemoveCompletedHandler_Create", 0),
	}
)

func coreWebView2BrowserExtensionRemoveCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2BrowserExtensionRemoveCompletedHandlerImport == nil {
		coreWebView2BrowserExtensionRemoveCompletedHandlerImport = NewDefaultImports()
		coreWebView2BrowserExtensionRemoveCompletedHandlerImport.SetImportTable(coreWebView2BrowserExtensionRemoveCompletedHandlerImportTables)
		coreWebView2BrowserExtensionRemoveCompletedHandlerImportTables = nil
	}
	return coreWebView2BrowserExtensionRemoveCompletedHandlerImport
}
