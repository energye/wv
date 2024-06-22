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

// ICoreWebView2MoveFocusRequestedEventArgs Parent: IObject
//
//	Event args for the MoveFocusRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs">See the ICoreWebView2MoveFocusRequestedEventArgs article.</a>
type ICoreWebView2MoveFocusRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2MoveFocusRequestedEventArgs // property
	// Reason
	//  The reason for WebView to run the `MoveFocusRequested` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs#get_reason">See the ICoreWebView2MoveFocusRequestedEventArgs article.</a>
	Reason() TWVMoveFocusReason // property
	// Handled
	//  Indicates whether the event has been handled by the app. If the app has
	//  moved the focus to another desired location, it should set the `Handled`
	//  property to `TRUE`. When the `Handled` property is `FALSE` after the
	//  event handler returns, default action is taken. The default action is to
	//  try to find the next tab stop child window in the app and try to move
	//  focus to that window. If no other window exists to move focus, focus is
	//  cycled within the web content of the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs#get_handled">See the ICoreWebView2MoveFocusRequestedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
}

// TCoreWebView2MoveFocusRequestedEventArgs Parent: TObject
//
//	Event args for the MoveFocusRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs">See the ICoreWebView2MoveFocusRequestedEventArgs article.</a>
type TCoreWebView2MoveFocusRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2MoveFocusRequestedEventArgs(aArgs ICoreWebView2MoveFocusRequestedEventArgs) ICoreWebView2MoveFocusRequestedEventArgs {
	r1 := coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2MoveFocusRequestedEventArgs(r1)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Initialized() bool {
	r1 := coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) BaseIntf() ICoreWebView2MoveFocusRequestedEventArgs {
	var resultCoreWebView2MoveFocusRequestedEventArgs uintptr
	coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2MoveFocusRequestedEventArgs)))
	return AsCoreWebView2MoveFocusRequestedEventArgs(resultCoreWebView2MoveFocusRequestedEventArgs)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Reason() TWVMoveFocusReason {
	r1 := coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(4, m.Instance())
	return TWVMoveFocusReason(r1)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Handled() bool {
	r1 := coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) SetHandled(AValue bool) {
	coreWebView2MoveFocusRequestedEventArgsImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

var (
	coreWebView2MoveFocusRequestedEventArgsImport       *imports.Imports = nil
	coreWebView2MoveFocusRequestedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2MoveFocusRequestedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2MoveFocusRequestedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2MoveFocusRequestedEventArgs_Handled", 0),
		/*3*/ imports.NewTable("CoreWebView2MoveFocusRequestedEventArgs_Initialized", 0),
		/*4*/ imports.NewTable("CoreWebView2MoveFocusRequestedEventArgs_Reason", 0),
	}
)

func coreWebView2MoveFocusRequestedEventArgsImportAPI() *imports.Imports {
	if coreWebView2MoveFocusRequestedEventArgsImport == nil {
		coreWebView2MoveFocusRequestedEventArgsImport = NewDefaultImports()
		coreWebView2MoveFocusRequestedEventArgsImport.SetImportTable(coreWebView2MoveFocusRequestedEventArgsImportTables)
		coreWebView2MoveFocusRequestedEventArgsImportTables = nil
	}
	return coreWebView2MoveFocusRequestedEventArgsImport
}
