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

// ICoreWebView2CookieList Parent: lcl.IObject
type ICoreWebView2CookieList interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CookieList // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist#get_count">See the ICoreWebView2CookieList article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookielist#getvalueatindex">See the ICoreWebView2CookieList article.</see>
	Items(idx uint32) ICoreWebView2Cookie // property Items Getter
}

type TCoreWebView2CookieList struct {
	lcl.TObject
}

func (m *TCoreWebView2CookieList) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieListAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieList) BaseIntf() (result ICoreWebView2CookieList) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieListAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieList(resultPtr)
	return
}

func (m *TCoreWebView2CookieList) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CookieListAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2CookieList) Items(idx uint32) (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieListAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

// NewCoreWebView2CookieList class constructor
func NewCoreWebView2CookieList(baseIntf ICoreWebView2CookieList) ICoreWebView2CookieList {
	r := coreWebView2CookieListAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2CookieList(r)
}

var (
	coreWebView2CookieListOnce   base.Once
	coreWebView2CookieListImport *imports.Imports = nil
)

func coreWebView2CookieListAPI() *imports.Imports {
	coreWebView2CookieListOnce.Do(func() {
		coreWebView2CookieListImport = api.NewDefaultImports()
		coreWebView2CookieListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2CookieList_Create", 0), // constructor NewCoreWebView2CookieList
			/* 1 */ imports.NewTable("TCoreWebView2CookieList_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2CookieList_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2CookieList_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2CookieList_Items", 0), // property Items
		}
	})
	return coreWebView2CookieListImport
}
