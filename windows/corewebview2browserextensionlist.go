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

// ICoreWebView2BrowserExtensionList Parent: IObject
type ICoreWebView2BrowserExtensionList interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserExtensionList // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist#get_count">See the ICoreWebView2ProcessExtendedInfoCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextensionlist#getvalueatindex">See the ICoreWebView2ProcessExtendedInfoCollection article.</see>
	Items(idx uint32) ICoreWebView2BrowserExtension // property Items Getter
}

type TCoreWebView2BrowserExtensionList struct {
	TObject
}

func (m *TCoreWebView2BrowserExtensionList) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BrowserExtensionListAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BrowserExtensionList) BaseIntf() (result ICoreWebView2BrowserExtensionList) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BrowserExtensionListAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BrowserExtensionList(resultPtr)
	return
}

func (m *TCoreWebView2BrowserExtensionList) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2BrowserExtensionListAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2BrowserExtensionList) Items(idx uint32) (result ICoreWebView2BrowserExtension) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BrowserExtensionListAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BrowserExtension(resultPtr)
	return
}

// NewCoreWebView2BrowserExtensionList class constructor
func NewCoreWebView2BrowserExtensionList(baseIntf ICoreWebView2BrowserExtensionList) ICoreWebView2BrowserExtensionList {
	r := coreWebView2BrowserExtensionListAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2BrowserExtensionList(r)
}

var (
	coreWebView2BrowserExtensionListOnce   base.Once
	coreWebView2BrowserExtensionListImport *imports.Imports = nil
)

func coreWebView2BrowserExtensionListAPI() *imports.Imports {
	coreWebView2BrowserExtensionListOnce.Do(func() {
		coreWebView2BrowserExtensionListImport = api.NewDefaultImports()
		coreWebView2BrowserExtensionListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BrowserExtensionList_Create", 0), // constructor NewCoreWebView2BrowserExtensionList
			/* 1 */ imports.NewTable("TCoreWebView2BrowserExtensionList_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2BrowserExtensionList_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2BrowserExtensionList_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2BrowserExtensionList_Items", 0), // property Items
		}
	})
	return coreWebView2BrowserExtensionListImport
}
