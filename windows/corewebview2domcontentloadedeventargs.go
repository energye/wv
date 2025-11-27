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
)

// ICoreWebView2DOMContentLoadedEventArgs Parent: lcl.IObject
type ICoreWebView2DOMContentLoadedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DOMContentLoadedEventArgs // property BaseIntf Getter
	// NavigationId
	//  The ID of the navigation which corresponds to other navigation ID properties on other navigation events.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2domcontentloadedeventargs#get_navigationid">See the ICoreWebView2DOMContentLoadedEventArgs article.</see>
	NavigationId() uint64 // property NavigationId Getter
}

type TCoreWebView2DOMContentLoadedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DOMContentLoadedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) BaseIntf() (result ICoreWebView2DOMContentLoadedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DOMContentLoadedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2DOMContentLoadedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2DOMContentLoadedEventArgs) NavigationId() (result uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2DOMContentLoadedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

// NewCoreWebView2DOMContentLoadedEventArgs class constructor
func NewCoreWebView2DOMContentLoadedEventArgs(args ICoreWebView2DOMContentLoadedEventArgs) ICoreWebView2DOMContentLoadedEventArgs {
	r := coreWebView2DOMContentLoadedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2DOMContentLoadedEventArgs(r)
}

var (
	coreWebView2DOMContentLoadedEventArgsOnce   base.Once
	coreWebView2DOMContentLoadedEventArgsImport *imports.Imports = nil
)

func coreWebView2DOMContentLoadedEventArgsAPI() *imports.Imports {
	coreWebView2DOMContentLoadedEventArgsOnce.Do(func() {
		coreWebView2DOMContentLoadedEventArgsImport = api.NewDefaultImports()
		coreWebView2DOMContentLoadedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2DOMContentLoadedEventArgs_Create", 0), // constructor NewCoreWebView2DOMContentLoadedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2DOMContentLoadedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2DOMContentLoadedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2DOMContentLoadedEventArgs_NavigationId", 0), // property NavigationId
		}
	})
	return coreWebView2DOMContentLoadedEventArgsImport
}
