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

// ICoreWebView2TrySuspendCompletedHandler Parent: IObject
//
//	The caller implements this interface to receive the TrySuspend result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2trysuspendcompletedhandler">See the ICoreWebView2TrySuspendCompletedHandler article.</a>
type ICoreWebView2TrySuspendCompletedHandler interface {
	IObject
}

// TCoreWebView2TrySuspendCompletedHandler Parent: TObject
//
//	The caller implements this interface to receive the TrySuspend result.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2trysuspendcompletedhandler">See the ICoreWebView2TrySuspendCompletedHandler article.</a>
type TCoreWebView2TrySuspendCompletedHandler struct {
	TObject
}

func NewCoreWebView2TrySuspendCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2TrySuspendCompletedHandler {
	r1 := coreWebView2TrySuspendCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2TrySuspendCompletedHandler(r1)
}

var (
	coreWebView2TrySuspendCompletedHandlerImport       *imports.Imports = nil
	coreWebView2TrySuspendCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2TrySuspendCompletedHandler_Create", 0),
	}
)

func coreWebView2TrySuspendCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2TrySuspendCompletedHandlerImport == nil {
		coreWebView2TrySuspendCompletedHandlerImport = NewDefaultImports()
		coreWebView2TrySuspendCompletedHandlerImport.SetImportTable(coreWebView2TrySuspendCompletedHandlerImportTables)
		coreWebView2TrySuspendCompletedHandlerImportTables = nil
	}
	return coreWebView2TrySuspendCompletedHandlerImport
}
