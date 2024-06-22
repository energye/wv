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

// ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Parent: IObject
//
//	The caller implements this interface to handle the result of
//	GetNonDefaultPermissionSettings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler">See the ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler article.</a>
type ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler interface {
	IObject
}

// TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Parent: TObject
//
//	The caller implements this interface to handle the result of
//	GetNonDefaultPermissionSettings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2getnondefaultpermissionsettingscompletedhandler">See the ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler article.</a>
type TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler struct {
	TObject
}

func NewCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(aEvents IWVBrowserEvents) ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	r1 := coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(r1)
}

var (
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport       *imports.Imports = nil
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2GetNonDefaultPermissionSettingsCompletedHandler_Create", 0),
	}
)

func coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport == nil {
		coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport = NewDefaultImports()
		coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport.SetImportTable(coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImportTables)
		coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImportTables = nil
	}
	return coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport
}
