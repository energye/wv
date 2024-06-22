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

// ICoreWebView2BrowserProcessExitedEventArgs Parent: IObject
//
//	Event args for the BrowserProcessExited event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
type ICoreWebView2BrowserProcessExitedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserProcessExitedEventArgs // property
	// BrowserProcessExitKind
	//  The kind of browser process exit that has occurred.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessexitkind">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
	BrowserProcessExitKind() TWVBrowserProcessExitKind // property
	// BrowserProcessId
	//  The process ID of the browser process that has exited.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessid">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
	BrowserProcessId() uint32 // property
}

// TCoreWebView2BrowserProcessExitedEventArgs Parent: TObject
//
//	Event args for the BrowserProcessExited event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
type TCoreWebView2BrowserProcessExitedEventArgs struct {
	TObject
}

func NewCoreWebView2BrowserProcessExitedEventArgs(aArgs ICoreWebView2BrowserProcessExitedEventArgs) ICoreWebView2BrowserProcessExitedEventArgs {
	r1 := coreWebView2BrowserProcessExitedEventArgsImportAPI().SysCallN(3, GetObjectUintptr(aArgs))
	return AsCoreWebView2BrowserProcessExitedEventArgs(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) Initialized() bool {
	r1 := coreWebView2BrowserProcessExitedEventArgsImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BaseIntf() ICoreWebView2BrowserProcessExitedEventArgs {
	var resultCoreWebView2BrowserProcessExitedEventArgs uintptr
	coreWebView2BrowserProcessExitedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2BrowserProcessExitedEventArgs)))
	return AsCoreWebView2BrowserProcessExitedEventArgs(resultCoreWebView2BrowserProcessExitedEventArgs)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessExitKind() TWVBrowserProcessExitKind {
	r1 := coreWebView2BrowserProcessExitedEventArgsImportAPI().SysCallN(1, m.Instance())
	return TWVBrowserProcessExitKind(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessId() uint32 {
	r1 := coreWebView2BrowserProcessExitedEventArgsImportAPI().SysCallN(2, m.Instance())
	return uint32(r1)
}

var (
	coreWebView2BrowserProcessExitedEventArgsImport       *imports.Imports = nil
	coreWebView2BrowserProcessExitedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BrowserProcessExitedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2BrowserProcessExitedEventArgs_BrowserProcessExitKind", 0),
		/*2*/ imports.NewTable("CoreWebView2BrowserProcessExitedEventArgs_BrowserProcessId", 0),
		/*3*/ imports.NewTable("CoreWebView2BrowserProcessExitedEventArgs_Create", 0),
		/*4*/ imports.NewTable("CoreWebView2BrowserProcessExitedEventArgs_Initialized", 0),
	}
)

func coreWebView2BrowserProcessExitedEventArgsImportAPI() *imports.Imports {
	if coreWebView2BrowserProcessExitedEventArgsImport == nil {
		coreWebView2BrowserProcessExitedEventArgsImport = NewDefaultImports()
		coreWebView2BrowserProcessExitedEventArgsImport.SetImportTable(coreWebView2BrowserProcessExitedEventArgsImportTables)
		coreWebView2BrowserProcessExitedEventArgsImportTables = nil
	}
	return coreWebView2BrowserProcessExitedEventArgsImport
}
