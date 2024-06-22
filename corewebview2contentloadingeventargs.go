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

// ICoreWebView2ContentLoadingEventArgs Parent: IObject
//
//	Receives ContentLoading events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs">See the ICoreWebView2ContentLoadingEventArgs article.</a>
type ICoreWebView2ContentLoadingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContentLoadingEventArgs // property
	// IsErrorPage
	//  `TRUE` if the loaded content is an error page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_iserrorpage">See the ICoreWebView2ContentLoadingEventArgs article.</a>
	IsErrorPage() bool // property
	// NavigationId
	//  The ID of the navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_navigationid">See the ICoreWebView2ContentLoadingEventArgs article.</a>
	NavigationId() (resultUint64 uint64) // property
}

// TCoreWebView2ContentLoadingEventArgs Parent: TObject
//
//	Receives ContentLoading events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs">See the ICoreWebView2ContentLoadingEventArgs article.</a>
type TCoreWebView2ContentLoadingEventArgs struct {
	TObject
}

func NewCoreWebView2ContentLoadingEventArgs(aArgs ICoreWebView2ContentLoadingEventArgs) ICoreWebView2ContentLoadingEventArgs {
	r1 := coreWebView2ContentLoadingEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2ContentLoadingEventArgs(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) Initialized() bool {
	r1 := coreWebView2ContentLoadingEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) BaseIntf() ICoreWebView2ContentLoadingEventArgs {
	var resultCoreWebView2ContentLoadingEventArgs uintptr
	coreWebView2ContentLoadingEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ContentLoadingEventArgs)))
	return AsCoreWebView2ContentLoadingEventArgs(resultCoreWebView2ContentLoadingEventArgs)
}

func (m *TCoreWebView2ContentLoadingEventArgs) IsErrorPage() bool {
	r1 := coreWebView2ContentLoadingEventArgsImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ContentLoadingEventArgs) NavigationId() (resultUint64 uint64) {
	coreWebView2ContentLoadingEventArgsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultUint64)))
	return
}

var (
	coreWebView2ContentLoadingEventArgsImport       *imports.Imports = nil
	coreWebView2ContentLoadingEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ContentLoadingEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ContentLoadingEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2ContentLoadingEventArgs_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2ContentLoadingEventArgs_IsErrorPage", 0),
		/*4*/ imports.NewTable("CoreWebView2ContentLoadingEventArgs_NavigationId", 0),
	}
)

func coreWebView2ContentLoadingEventArgsImportAPI() *imports.Imports {
	if coreWebView2ContentLoadingEventArgsImport == nil {
		coreWebView2ContentLoadingEventArgsImport = NewDefaultImports()
		coreWebView2ContentLoadingEventArgsImport.SetImportTable(coreWebView2ContentLoadingEventArgsImportTables)
		coreWebView2ContentLoadingEventArgsImportTables = nil
	}
	return coreWebView2ContentLoadingEventArgsImport
}
