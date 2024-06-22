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

// ICoreWebView2SourceChangedEventArgs Parent: IObject
//
//	Event args for the SourceChanged event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs">See the ICoreWebView2SourceChangedEventArgs article.</a>
type ICoreWebView2SourceChangedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SourceChangedEventArgs // property
	// IsNewDocument
	//  `TRUE` if the page being navigated to is a new document.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs#get_isnewdocument">See the ICoreWebView2SourceChangedEventArgs article.</a>
	IsNewDocument() bool // property
}

// TCoreWebView2SourceChangedEventArgs Parent: TObject
//
//	Event args for the SourceChanged event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sourcechangedeventargs">See the ICoreWebView2SourceChangedEventArgs article.</a>
type TCoreWebView2SourceChangedEventArgs struct {
	TObject
}

func NewCoreWebView2SourceChangedEventArgs(aArgs ICoreWebView2SourceChangedEventArgs) ICoreWebView2SourceChangedEventArgs {
	r1 := coreWebView2SourceChangedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2SourceChangedEventArgs(r1)
}

func (m *TCoreWebView2SourceChangedEventArgs) Initialized() bool {
	r1 := coreWebView2SourceChangedEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2SourceChangedEventArgs) BaseIntf() ICoreWebView2SourceChangedEventArgs {
	var resultCoreWebView2SourceChangedEventArgs uintptr
	coreWebView2SourceChangedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2SourceChangedEventArgs)))
	return AsCoreWebView2SourceChangedEventArgs(resultCoreWebView2SourceChangedEventArgs)
}

func (m *TCoreWebView2SourceChangedEventArgs) IsNewDocument() bool {
	r1 := coreWebView2SourceChangedEventArgsImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

var (
	coreWebView2SourceChangedEventArgsImport       *imports.Imports = nil
	coreWebView2SourceChangedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2SourceChangedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2SourceChangedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2SourceChangedEventArgs_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2SourceChangedEventArgs_IsNewDocument", 0),
	}
)

func coreWebView2SourceChangedEventArgsImportAPI() *imports.Imports {
	if coreWebView2SourceChangedEventArgsImport == nil {
		coreWebView2SourceChangedEventArgsImport = NewDefaultImports()
		coreWebView2SourceChangedEventArgsImport.SetImportTable(coreWebView2SourceChangedEventArgsImportTables)
		coreWebView2SourceChangedEventArgsImportTables = nil
	}
	return coreWebView2SourceChangedEventArgsImport
}
