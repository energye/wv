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

// ICoreWebView2ProcessInfoCollection Parent: IObject
type ICoreWebView2ProcessInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessInfoCollection // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection#get_count">See the ICoreWebView2ProcessInfoCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfocollection#getvalueatindex">See the ICoreWebView2ProcessInfoCollection article.</see>
	Items(idx uint32) ICoreWebView2ProcessInfo // property Items Getter
}

type TCoreWebView2ProcessInfoCollection struct {
	TObject
}

func (m *TCoreWebView2ProcessInfoCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProcessInfoCollectionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ProcessInfoCollection) BaseIntf() (result ICoreWebView2ProcessInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessInfoCollectionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfoCollection(resultPtr)
	return
}

func (m *TCoreWebView2ProcessInfoCollection) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessInfoCollectionAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ProcessInfoCollection) Items(idx uint32) (result ICoreWebView2ProcessInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessInfoCollectionAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfo(resultPtr)
	return
}

// NewCoreWebView2ProcessInfoCollection class constructor
func NewCoreWebView2ProcessInfoCollection(baseIntf ICoreWebView2ProcessInfoCollection) ICoreWebView2ProcessInfoCollection {
	r := coreWebView2ProcessInfoCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ProcessInfoCollection(r)
}

var (
	coreWebView2ProcessInfoCollectionOnce   base.Once
	coreWebView2ProcessInfoCollectionImport *imports.Imports = nil
)

func coreWebView2ProcessInfoCollectionAPI() *imports.Imports {
	coreWebView2ProcessInfoCollectionOnce.Do(func() {
		coreWebView2ProcessInfoCollectionImport = api.NewDefaultImports()
		coreWebView2ProcessInfoCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProcessInfoCollection_Create", 0), // constructor NewCoreWebView2ProcessInfoCollection
			/* 1 */ imports.NewTable("TCoreWebView2ProcessInfoCollection_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ProcessInfoCollection_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ProcessInfoCollection_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2ProcessInfoCollection_Items", 0), // property Items
		}
	})
	return coreWebView2ProcessInfoCollectionImport
}
