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

// ICoreWebView2CustomItemSelectedEventHandler Parent: IObject
//
//	Raised to notify the host that the end user selected a custom
//	ContextMenuItem. CustomItemSelected event is raised on the specific
//	ContextMenuItem that the end user selected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2customitemselectedeventhandler">See the ICoreWebView2CustomItemSelectedEventHandler article.</a>
type ICoreWebView2CustomItemSelectedEventHandler interface {
	IObject
}

// TCoreWebView2CustomItemSelectedEventHandler Parent: TObject
//
//	Raised to notify the host that the end user selected a custom
//	ContextMenuItem. CustomItemSelected event is raised on the specific
//	ContextMenuItem that the end user selected.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2customitemselectedeventhandler">See the ICoreWebView2CustomItemSelectedEventHandler article.</a>
type TCoreWebView2CustomItemSelectedEventHandler struct {
	TObject
}

func NewCoreWebView2CustomItemSelectedEventHandler(aEvents IWVBrowserEvents) ICoreWebView2CustomItemSelectedEventHandler {
	r1 := coreWebView2CustomItemSelectedEventHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents))
	return AsCoreWebView2CustomItemSelectedEventHandler(r1)
}

var (
	coreWebView2CustomItemSelectedEventHandlerImport       *imports.Imports = nil
	coreWebView2CustomItemSelectedEventHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2CustomItemSelectedEventHandler_Create", 0),
	}
)

func coreWebView2CustomItemSelectedEventHandlerImportAPI() *imports.Imports {
	if coreWebView2CustomItemSelectedEventHandlerImport == nil {
		coreWebView2CustomItemSelectedEventHandlerImport = NewDefaultImports()
		coreWebView2CustomItemSelectedEventHandlerImport.SetImportTable(coreWebView2CustomItemSelectedEventHandlerImportTables)
		coreWebView2CustomItemSelectedEventHandlerImportTables = nil
	}
	return coreWebView2CustomItemSelectedEventHandlerImport
}
