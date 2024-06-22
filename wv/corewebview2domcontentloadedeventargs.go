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

// ICoreWebView2DOMContentLoadedEventArgs Parent: IObject
//
//	Event args for the DOMContentLoaded event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
type ICoreWebView2DOMContentLoadedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DOMContentLoadedEventArgs // property
	// NavigationId
	//  The ID of the navigation which corresponds to other navigation ID properties on other navigation events.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs#get_navigationid">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
	NavigationId() (resultUint64 uint64) // property
}

// TCoreWebView2DOMContentLoadedEventArgs Parent: TObject
//
//	Event args for the DOMContentLoaded event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs">See the ICoreWebView2DOMContentLoadedEventArgs article.</a>
type TCoreWebView2DOMContentLoadedEventArgs struct {
	TObject
}

func NewCoreWebView2DOMContentLoadedEventArgs(aArgs ICoreWebView2DOMContentLoadedEventArgs) ICoreWebView2DOMContentLoadedEventArgs {
	r1 := coreWebView2DOMContentLoadedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2DOMContentLoadedEventArgs(r1)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) Initialized() bool {
	r1 := coreWebView2DOMContentLoadedEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) BaseIntf() ICoreWebView2DOMContentLoadedEventArgs {
	var resultCoreWebView2DOMContentLoadedEventArgs uintptr
	coreWebView2DOMContentLoadedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DOMContentLoadedEventArgs)))
	return AsCoreWebView2DOMContentLoadedEventArgs(resultCoreWebView2DOMContentLoadedEventArgs)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) NavigationId() (resultUint64 uint64) {
	coreWebView2DOMContentLoadedEventArgsImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

var (
	coreWebView2DOMContentLoadedEventArgsImport       *imports.Imports = nil
	coreWebView2DOMContentLoadedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2DOMContentLoadedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2DOMContentLoadedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2DOMContentLoadedEventArgs_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2DOMContentLoadedEventArgs_NavigationId", 0),
	}
)

func coreWebView2DOMContentLoadedEventArgsImportAPI() *imports.Imports {
	if coreWebView2DOMContentLoadedEventArgsImport == nil {
		coreWebView2DOMContentLoadedEventArgsImport = NewDefaultImports()
		coreWebView2DOMContentLoadedEventArgsImport.SetImportTable(coreWebView2DOMContentLoadedEventArgsImportTables)
		coreWebView2DOMContentLoadedEventArgsImportTables = nil
	}
	return coreWebView2DOMContentLoadedEventArgsImport
}
