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
)

// ICoreWebView2ContentLoadingEventArgs Parent: IObject
type ICoreWebView2ContentLoadingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContentLoadingEventArgs // property BaseIntf Getter
	// IsErrorPage
	//  `TRUE` if the loaded content is an error page.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_iserrorpage">See the ICoreWebView2ContentLoadingEventArgs article.</see>
	IsErrorPage() bool // property IsErrorPage Getter
	// NavigationId
	//  The ID of the navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contentloadingeventargs#get_navigationid">See the ICoreWebView2ContentLoadingEventArgs article.</see>
	NavigationId() uint64 // property NavigationId Getter
}

type TCoreWebView2ContentLoadingEventArgs struct {
	TObject
}

func (m *TCoreWebView2ContentLoadingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContentLoadingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContentLoadingEventArgs) BaseIntf() (result ICoreWebView2ContentLoadingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContentLoadingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContentLoadingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ContentLoadingEventArgs) IsErrorPage() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContentLoadingEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContentLoadingEventArgs) NavigationId() (result uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContentLoadingEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

// NewCoreWebView2ContentLoadingEventArgs class constructor
func NewCoreWebView2ContentLoadingEventArgs(args ICoreWebView2ContentLoadingEventArgs) ICoreWebView2ContentLoadingEventArgs {
	r := coreWebView2ContentLoadingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ContentLoadingEventArgs(r)
}

var (
	coreWebView2ContentLoadingEventArgsOnce   base.Once
	coreWebView2ContentLoadingEventArgsImport *imports.Imports = nil
)

func coreWebView2ContentLoadingEventArgsAPI() *imports.Imports {
	coreWebView2ContentLoadingEventArgsOnce.Do(func() {
		coreWebView2ContentLoadingEventArgsImport = api.NewDefaultImports()
		coreWebView2ContentLoadingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ContentLoadingEventArgs_Create", 0), // constructor NewCoreWebView2ContentLoadingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ContentLoadingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ContentLoadingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ContentLoadingEventArgs_IsErrorPage", 0), // property IsErrorPage
			/* 4 */ imports.NewTable("TCoreWebView2ContentLoadingEventArgs_NavigationId", 0), // property NavigationId
		}
	})
	return coreWebView2ContentLoadingEventArgsImport
}
