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

// ICoreWebView2StringCollection Parent: lcl.IObject
type ICoreWebView2StringCollection interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2StringCollection // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection#get_count">See the ICoreWebView2StringCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection#getvalueatindex">See the ICoreWebView2StringCollection article.</see>
	Items(idx uint32) string // property Items Getter
}

type TCoreWebView2StringCollection struct {
	lcl.TObject
}

func (m *TCoreWebView2StringCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2StringCollectionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2StringCollection) BaseIntf() (result ICoreWebView2StringCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2StringCollectionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2StringCollection(resultPtr)
	return
}

func (m *TCoreWebView2StringCollection) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2StringCollectionAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2StringCollection) Items(idx uint32) string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2StringCollectionAPI().SysCallN(4, m.Instance(), uintptr(idx))
	return api.GoStr(r)
}

// NewCoreWebView2StringCollection class constructor
func NewCoreWebView2StringCollection(baseIntf ICoreWebView2StringCollection) ICoreWebView2StringCollection {
	r := coreWebView2StringCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2StringCollection(r)
}

var (
	coreWebView2StringCollectionOnce   base.Once
	coreWebView2StringCollectionImport *imports.Imports = nil
)

func coreWebView2StringCollectionAPI() *imports.Imports {
	coreWebView2StringCollectionOnce.Do(func() {
		coreWebView2StringCollectionImport = api.NewDefaultImports()
		coreWebView2StringCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2StringCollection_Create", 0), // constructor NewCoreWebView2StringCollection
			/* 1 */ imports.NewTable("TCoreWebView2StringCollection_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2StringCollection_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2StringCollection_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2StringCollection_Items", 0), // property Items
		}
	})
	return coreWebView2StringCollectionImport
}
