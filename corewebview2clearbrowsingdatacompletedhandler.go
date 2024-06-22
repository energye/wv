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

// ICoreWebView2ClearBrowsingDataCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the ClearBrowsingData result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clearbrowsingdatacompletedhandler">See the ICoreWebView2ClearBrowsingDataCompletedHandler article.</a>
type ICoreWebView2ClearBrowsingDataCompletedHandler interface {
	IObject
}

// TCoreWebView2ClearBrowsingDataCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the ClearBrowsingData result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clearbrowsingdatacompletedhandler">See the ICoreWebView2ClearBrowsingDataCompletedHandler article.</a>
type TCoreWebView2ClearBrowsingDataCompletedHandler struct {
	TObject
}

func NewCoreWebView2ClearBrowsingDataCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2ClearBrowsingDataCompletedHandler {
	r1 := coreWebView2ClearBrowsingDataCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2ClearBrowsingDataCompletedHandler(r1)
}

var (
	coreWebView2ClearBrowsingDataCompletedHandlerImport       *imports.Imports = nil
	coreWebView2ClearBrowsingDataCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ClearBrowsingDataCompletedHandler_Create", 0),
	}
)

func coreWebView2ClearBrowsingDataCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2ClearBrowsingDataCompletedHandlerImport == nil {
		coreWebView2ClearBrowsingDataCompletedHandlerImport = NewDefaultImports()
		coreWebView2ClearBrowsingDataCompletedHandlerImport.SetImportTable(coreWebView2ClearBrowsingDataCompletedHandlerImportTables)
		coreWebView2ClearBrowsingDataCompletedHandlerImportTables = nil
	}
	return coreWebView2ClearBrowsingDataCompletedHandlerImport
}
