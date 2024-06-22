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

// ICoreWebView2PrintCompletedHandler Parent: IObject
//
//	Receives the result of the Print method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler">See the ICoreWebView2PrintCompletedHandler article.</a>
type ICoreWebView2PrintCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintCompletedHandler Parent: TObject
//
//	Receives the result of the Print method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printcompletedhandler">See the ICoreWebView2PrintCompletedHandler article.</a>
type TCoreWebView2PrintCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintCompletedHandler {
	r1 := coreWebView2PrintCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintCompletedHandler(r1)
}

var (
	coreWebView2PrintCompletedHandlerImport       *imports.Imports = nil
	coreWebView2PrintCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PrintCompletedHandler_Create", 0),
	}
)

func coreWebView2PrintCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2PrintCompletedHandlerImport == nil {
		coreWebView2PrintCompletedHandlerImport = NewDefaultImports()
		coreWebView2PrintCompletedHandlerImport.SetImportTable(coreWebView2PrintCompletedHandlerImportTables)
		coreWebView2PrintCompletedHandlerImportTables = nil
	}
	return coreWebView2PrintCompletedHandlerImport
}
