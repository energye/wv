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
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2ScriptDialogOpeningEventArgs Parent: lcl.IObject
type ICoreWebView2ScriptDialogOpeningEventArgs interface {
	lcl.IObject
	// Accept
	//  The host may run this to respond with **OK** to `confirm`, `prompt`, and
	//  `beforeunload` dialogs. Do not run this method to indicate cancel.
	//  From JavaScript, this means that the `confirm` and `beforeunload` function
	//  returns `TRUE` if `Accept` is run. And for the prompt function it returns
	//  the value of `ResultText` if `Accept` is run and otherwise returns
	//  `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#accept">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	Accept() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScriptDialogOpeningEventArgs // property BaseIntf Getter
	// URI
	//  The URI of the page that requested the dialog box.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_uri">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	URI() string // property URI Getter
	// Kind
	//  The kind of JavaScript dialog box. `alert`, `confirm`, `prompt`, or
	//  `beforeunload`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_kind">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	Kind() wvTypes.TWVScriptDialogKind // property Kind Getter
	// Message
	//  The message of the dialog box. From JavaScript this is the first
	//  parameter passed to `alert`, `confirm`, and `prompt` and is empty for
	//  `beforeunload`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_message">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	Message() string // property Message_ Getter
	// DefaultText
	//  The second parameter passed to the JavaScript prompt dialog.
	//  The result of the prompt JavaScript function uses this value as the
	//  default value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_defaulttext">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	DefaultText() string // property DefaultText Getter
	// ResultText
	//  The return value from the JavaScript prompt function if `Accept` is run.
	//  This value is ignored for dialog kinds other than prompt. If `Accept`
	//  is not run, this value is ignored and `FALSE` is returned from prompt.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_resulttext">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	ResultText() string         // property ResultText Getter
	SetResultText(value string) // property ResultText Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#getdeferral">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2ScriptDialogOpeningEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Accept() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) BaseIntf() (result ICoreWebView2ScriptDialogOpeningEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ScriptDialogOpeningEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) URI() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Kind() wvTypes.TWVScriptDialogKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(5, m.Instance())
	return wvTypes.TWVScriptDialogKind(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Message() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) DefaultText() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) ResultText() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(8, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) SetResultText(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2ScriptDialogOpeningEventArgs class constructor
func NewCoreWebView2ScriptDialogOpeningEventArgs(args ICoreWebView2ScriptDialogOpeningEventArgs) ICoreWebView2ScriptDialogOpeningEventArgs {
	r := coreWebView2ScriptDialogOpeningEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ScriptDialogOpeningEventArgs(r)
}

var (
	coreWebView2ScriptDialogOpeningEventArgsOnce   base.Once
	coreWebView2ScriptDialogOpeningEventArgsImport *imports.Imports = nil
)

func coreWebView2ScriptDialogOpeningEventArgsAPI() *imports.Imports {
	coreWebView2ScriptDialogOpeningEventArgsOnce.Do(func() {
		coreWebView2ScriptDialogOpeningEventArgsImport = api.NewDefaultImports()
		coreWebView2ScriptDialogOpeningEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Create", 0), // constructor NewCoreWebView2ScriptDialogOpeningEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Accept", 0), // function Accept
			/* 2 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_URI", 0), // property URI
			/* 5 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Kind", 0), // property Kind
			/* 6 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Message_", 0), // property Message
			/* 7 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_DefaultText", 0), // property DefaultText
			/* 8 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_ResultText", 0), // property ResultText
			/* 9 */ imports.NewTable("TCoreWebView2ScriptDialogOpeningEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2ScriptDialogOpeningEventArgsImport
}
