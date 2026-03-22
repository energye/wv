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

// ICoreWebView2BrowserProcessExitedEventArgs Parent: IObject
type ICoreWebView2BrowserProcessExitedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserProcessExitedEventArgs // property BaseIntf Getter
	// BrowserProcessExitKind
	//  The kind of browser process exit that has occurred.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessexitkind">See the ICoreWebView2BrowserProcessExitedEventArgs article.</see>
	BrowserProcessExitKind() wvTypes.TWVBrowserProcessExitKind // property BrowserProcessExitKind Getter
	// BrowserProcessId
	//  The process ID of the browser process that has exited.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessid">See the ICoreWebView2BrowserProcessExitedEventArgs article.</see>
	BrowserProcessId() uint32 // property BrowserProcessId Getter
}

type TCoreWebView2BrowserProcessExitedEventArgs struct {
	TObject
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserProcessExitedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BaseIntf() (result ICoreWebView2BrowserProcessExitedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BrowserProcessExitedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BrowserProcessExitedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessExitKind() wvTypes.TWVBrowserProcessExitKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2BrowserProcessExitedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVBrowserProcessExitKind(r)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2BrowserProcessExitedEventArgsAPI().SysCallN(4, m.Instance())
	return uint32(r)
}

// NewCoreWebView2BrowserProcessExitedEventArgs class constructor
func NewCoreWebView2BrowserProcessExitedEventArgs(args ICoreWebView2BrowserProcessExitedEventArgs) ICoreWebView2BrowserProcessExitedEventArgs {
	r := coreWebView2BrowserProcessExitedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2BrowserProcessExitedEventArgs(r)
}

var (
	coreWebView2BrowserProcessExitedEventArgsOnce   base.Once
	coreWebView2BrowserProcessExitedEventArgsImport *imports.Imports = nil
)

func coreWebView2BrowserProcessExitedEventArgsAPI() *imports.Imports {
	coreWebView2BrowserProcessExitedEventArgsOnce.Do(func() {
		coreWebView2BrowserProcessExitedEventArgsImport = api.NewDefaultImports()
		coreWebView2BrowserProcessExitedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BrowserProcessExitedEventArgs_Create", 0), // constructor NewCoreWebView2BrowserProcessExitedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2BrowserProcessExitedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2BrowserProcessExitedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2BrowserProcessExitedEventArgs_BrowserProcessExitKind", 0), // property BrowserProcessExitKind
			/* 4 */ imports.NewTable("TCoreWebView2BrowserProcessExitedEventArgs_BrowserProcessId", 0), // property BrowserProcessId
		}
	})
	return coreWebView2BrowserProcessExitedEventArgsImport
}
