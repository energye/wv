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

// ICoreWebView2PrintToPdfStreamCompletedHandler Parent: IObject
//
//	Receives the result of the PrintToPdfStream method.
//	errorCode returns S_OK if the PrintToPdfStream operation succeeded.
//	The printable pdf data is returned in the pdfStream object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler">See the ICoreWebView2PrintToPdfStreamCompletedHandler article.</a>
type ICoreWebView2PrintToPdfStreamCompletedHandler interface {
	IObject
}

// TCoreWebView2PrintToPdfStreamCompletedHandler Parent: TObject
//
//	Receives the result of the PrintToPdfStream method.
//	errorCode returns S_OK if the PrintToPdfStream operation succeeded.
//	The printable pdf data is returned in the pdfStream object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printtopdfstreamcompletedhandler">See the ICoreWebView2PrintToPdfStreamCompletedHandler article.</a>
type TCoreWebView2PrintToPdfStreamCompletedHandler struct {
	TObject
}

func NewCoreWebView2PrintToPdfStreamCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2PrintToPdfStreamCompletedHandler {
	r1 := coreWebView2PrintToPdfStreamCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2PrintToPdfStreamCompletedHandler(r1)
}

var (
	coreWebView2PrintToPdfStreamCompletedHandlerImport       *imports.Imports = nil
	coreWebView2PrintToPdfStreamCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PrintToPdfStreamCompletedHandler_Create", 0),
	}
)

func coreWebView2PrintToPdfStreamCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2PrintToPdfStreamCompletedHandlerImport == nil {
		coreWebView2PrintToPdfStreamCompletedHandlerImport = NewDefaultImports()
		coreWebView2PrintToPdfStreamCompletedHandlerImport.SetImportTable(coreWebView2PrintToPdfStreamCompletedHandlerImportTables)
		coreWebView2PrintToPdfStreamCompletedHandlerImportTables = nil
	}
	return coreWebView2PrintToPdfStreamCompletedHandlerImport
}
