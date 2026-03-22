//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2MoveFocusRequestedEventArgs Parent: IObject
type ICoreWebView2MoveFocusRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2MoveFocusRequestedEventArgs // property BaseIntf Getter
	// Reason
	//  The reason for WebView to run the `MoveFocusRequested` event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs#get_reason">See the ICoreWebView2MoveFocusRequestedEventArgs article.</see>
	Reason() wvTypes.TWVMoveFocusReason // property Reason Getter
	// Handled
	//  Indicates whether the event has been handled by the app. If the app has
	//  moved the focus to another desired location, it should set the `Handled`
	//  property to `TRUE`. When the `Handled` property is `FALSE` after the
	//  event handler returns, default action is taken. The default action is to
	//  try to find the next tab stop child window in the app and try to move
	//  focus to that window. If no other window exists to move focus, focus is
	//  cycled within the web content of the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2movefocusrequestedeventargs#get_handled">See the ICoreWebView2MoveFocusRequestedEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
}

type TCoreWebView2MoveFocusRequestedEventArgs struct {
	TObject
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) BaseIntf() (result ICoreWebView2MoveFocusRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2MoveFocusRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Reason() wvTypes.TWVMoveFocusReason {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVMoveFocusReason(r)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2MoveFocusRequestedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2MoveFocusRequestedEventArgs class constructor
func NewCoreWebView2MoveFocusRequestedEventArgs(args ICoreWebView2MoveFocusRequestedEventArgs) ICoreWebView2MoveFocusRequestedEventArgs {
	r := coreWebView2MoveFocusRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2MoveFocusRequestedEventArgs(r)
}

var (
	coreWebView2MoveFocusRequestedEventArgsOnce   base.Once
	coreWebView2MoveFocusRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2MoveFocusRequestedEventArgsAPI() *imports.Imports {
	coreWebView2MoveFocusRequestedEventArgsOnce.Do(func() {
		coreWebView2MoveFocusRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2MoveFocusRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2MoveFocusRequestedEventArgs_Create", 0), // constructor NewCoreWebView2MoveFocusRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2MoveFocusRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2MoveFocusRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2MoveFocusRequestedEventArgs_Reason", 0), // property Reason
			/* 4 */ imports.NewTable("TCoreWebView2MoveFocusRequestedEventArgs_Handled", 0), // property Handled
		}
	})
	return coreWebView2MoveFocusRequestedEventArgsImport
}
