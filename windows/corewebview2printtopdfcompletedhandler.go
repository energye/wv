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

// ICoreWebView2PrintToPdfCompletedHandler Parent: IObject
//
//	Receives the result of the PrintToPdf method. If the print to PDF
//	operation succeeds, isSuccessful is true. Otherwise, if the operation
//	failed, isSuccessful is set to false. An invalid path returns
//	E_INVALIDARG.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfcompletedhandler">See the ICoreWebView2PrintToPdfCompletedHandler article.</a>
type ICoreWebView2PrintToPdfCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintToPdfCompletedHandler Parent: TObject
//
//	Receives the result of the PrintToPdf method. If the print to PDF
//	operation succeeds, isSuccessful is true. Otherwise, if the operation
//	failed, isSuccessful is set to false. An invalid path returns
//	E_INVALIDARG.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfcompletedhandler">See the ICoreWebView2PrintToPdfCompletedHandler article.</a>
type TCoreWebView2PrintToPdfCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintToPdfCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintToPdfCompletedHandler {
	r1 := coreWebView2PrintToPdfCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintToPdfCompletedHandler(r1)
}

var (
	coreWebView2PrintToPdfCompletedHandlerImport       *imports.Imports = nil
	coreWebView2PrintToPdfCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PrintToPdfCompletedHandler_Create", 0),
	}
)

func coreWebView2PrintToPdfCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2PrintToPdfCompletedHandlerImport == nil {
		coreWebView2PrintToPdfCompletedHandlerImport = NewDefaultImports()
		coreWebView2PrintToPdfCompletedHandlerImport.SetImportTable(coreWebView2PrintToPdfCompletedHandlerImportTables)
		coreWebView2PrintToPdfCompletedHandlerImportTables = nil
	}
	return coreWebView2PrintToPdfCompletedHandlerImport
}
