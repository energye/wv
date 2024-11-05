//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

var globalWebView2Loader IWVLoader

// GlobalWebView2Loader
// Global instance of TWVLoader used to simplify the WebView2 initialization and destruction.
func GlobalWebView2Loader() IWVLoader {
	if globalWebView2Loader != nil && globalWebView2Loader.IsValid() {
		return globalWebView2Loader
	}
	// Get Global TWVLoader
	r1 := predefWVLoaderImportAPI().SysCallN(0)
	if r1 != 0 {
		globalWebView2Loader = AsWVLoader(r1)
	} else {
		// Create New Global TWVLoader
		SetGlobalWebView2Loader(NewWVLoader(nil))
	}
	return globalWebView2Loader
}

// SetGlobalWebView2Loader
// Global instance of TWVLoader used to simplify the WebView2 initialization and destruction.
func SetGlobalWebView2Loader(wvLoader IWVLoader) {
	if globalWebView2Loader == nil || !globalWebView2Loader.IsValid() {
		predefWVLoaderImportAPI().SysCallN(1, GetObjectUintptr(wvLoader))
		globalWebView2Loader = wvLoader
	}
}

func DestroyGlobalWebView2Loader() {
	//predefWVLoaderImportAPI().SysCallN(2)
	if globalWebView2Loader != nil {
		var data = globalWebView2Loader.Instance()
		api.DFreeAndNil(uintptr(unsafePointer(&data)))
		globalWebView2Loader = nil
	}
}

func (m *TWVLoader) SetOnEnvironmentCreated(fn TOnLoaderNotifyEvent) {
	predefWVLoaderImportAPI().SysCallN(3, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnInitializationError(fn TOnLoaderNotifyEvent) {
	predefWVLoaderImportAPI().SysCallN(4, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnGetCustomSchemes(fn TOnLoaderGetCustomSchemesEvent) {
	predefWVLoaderImportAPI().SysCallN(5, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnNewBrowserVersionAvailable(fn TOnLoaderNewBrowserVersionAvailableEvent) {
	predefWVLoaderImportAPI().SysCallN(6, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnBrowserProcessExited(fn TOnLoaderBrowserProcessExitedEvent) {
	predefWVLoaderImportAPI().SysCallN(7, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnProcessInfosChanged(fn TOnLoaderProcessInfosChangedEvent) {
	predefWVLoaderImportAPI().SysCallN(8, api.MakeEventDataPtr(fn))
}

var (
	predefWVLoaderImport       *imports.Imports = nil
	predefWVLoaderImportTables                  = []*imports.Table{
		imports.NewTable("WVLoader_GetGlobalWebView2Loader", 0),
		imports.NewTable("WVLoader_SetGlobalWebView2Loader", 0),
		imports.NewTable("WVLoader_DestroyGlobalWebView2Loader", 0),
		imports.NewTable("WVLoader_SetOnEnvironmentCreated", 0),
		imports.NewTable("WVLoader_SetOnInitializationError", 0),
		imports.NewTable("WVLoader_SetOnGetCustomSchemes", 0),
		imports.NewTable("WVLoader_SetOnNewBrowserVersionAvailable", 0),
		imports.NewTable("WVLoader_SetOnBrowserProcessExited", 0),
		imports.NewTable("WVLoader_SetOnProcessInfosChanged", 0),
	}
)

func predefWVLoaderImportAPI() *imports.Imports {
	if predefWVLoaderImport == nil {
		predefWVLoaderImport = api.NewDefaultImports()
		predefWVLoaderImport.SetImportTable(predefWVLoaderImportTables)
		predefWVLoaderImportTables = nil
	}
	return predefWVLoaderImport
}
