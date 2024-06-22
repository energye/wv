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

// ICoreWebView2SetPermissionStateCompletedHandler Parent: IObject
//
//	The caller implements this interface to handle the result of
//	SetPermissionState.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler">See the ICoreWebView2SetPermissionStateCompletedHandler article.</a>
type ICoreWebView2SetPermissionStateCompletedHandler interface {
	IObject
}

// TCoreWebView2SetPermissionStateCompletedHandler Parent: TObject
//
//	The caller implements this interface to handle the result of
//	SetPermissionState.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2setpermissionstatecompletedhandler">See the ICoreWebView2SetPermissionStateCompletedHandler article.</a>
type TCoreWebView2SetPermissionStateCompletedHandler struct {
	TObject
}

func NewCoreWebView2SetPermissionStateCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2SetPermissionStateCompletedHandler {
	r1 := coreWebView2SetPermissionStateCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2SetPermissionStateCompletedHandler(r1)
}

var (
	coreWebView2SetPermissionStateCompletedHandlerImport       *imports.Imports = nil
	coreWebView2SetPermissionStateCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2SetPermissionStateCompletedHandler_Create", 0),
	}
)

func coreWebView2SetPermissionStateCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2SetPermissionStateCompletedHandlerImport == nil {
		coreWebView2SetPermissionStateCompletedHandlerImport = NewDefaultImports()
		coreWebView2SetPermissionStateCompletedHandlerImport.SetImportTable(coreWebView2SetPermissionStateCompletedHandlerImportTables)
		coreWebView2SetPermissionStateCompletedHandlerImportTables = nil
	}
	return coreWebView2SetPermissionStateCompletedHandlerImport
}
