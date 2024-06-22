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

// ICoreWebView2ProfileAddBrowserExtensionCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the result
//	of loading an browser Extension.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profileaddbrowserextensioncompletedhandler">See the ICoreWebView2ProfileAddBrowserExtensionCompletedHandler article.</a>
type ICoreWebView2ProfileAddBrowserExtensionCompletedHandler interface {
	IObject
}

// TCoreWebView2ProfileAddBrowserExtensionCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the result
//	of loading an browser Extension.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profileaddbrowserextensioncompletedhandler">See the ICoreWebView2ProfileAddBrowserExtensionCompletedHandler article.</a>
type TCoreWebView2ProfileAddBrowserExtensionCompletedHandler struct {
	TObject
}

func NewCoreWebView2ProfileAddBrowserExtensionCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	r1 := coreWebView2ProfileAddBrowserExtensionCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler(r1)
}

var (
	coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport       *imports.Imports = nil
	coreWebView2ProfileAddBrowserExtensionCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ProfileAddBrowserExtensionCompletedHandler_Create", 0),
	}
)

func coreWebView2ProfileAddBrowserExtensionCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport == nil {
		coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport = NewDefaultImports()
		coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport.SetImportTable(coreWebView2ProfileAddBrowserExtensionCompletedHandlerImportTables)
		coreWebView2ProfileAddBrowserExtensionCompletedHandlerImportTables = nil
	}
	return coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport
}
