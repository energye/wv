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

// ICoreWebView2ProcessExtendedInfoCollection Parent: IObject
type ICoreWebView2ProcessExtendedInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessExtendedInfoCollection // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#get_count">See the ICoreWebView2ProcessExtendedInfoCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#getvalueatindex">See the ICoreWebView2ProcessExtendedInfoCollection article.</see>
	Items(idx uint32) ICoreWebView2ProcessExtendedInfo // property Items Getter
}

type TCoreWebView2ProcessExtendedInfoCollection struct {
	TObject
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProcessExtendedInfoCollectionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) BaseIntf() (result ICoreWebView2ProcessExtendedInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessExtendedInfoCollectionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessExtendedInfoCollection(resultPtr)
	return
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessExtendedInfoCollectionAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Items(idx uint32) (result ICoreWebView2ProcessExtendedInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessExtendedInfoCollectionAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessExtendedInfo(resultPtr)
	return
}

// NewCoreWebView2ProcessExtendedInfoCollection class constructor
func NewCoreWebView2ProcessExtendedInfoCollection(baseIntf ICoreWebView2ProcessExtendedInfoCollection) ICoreWebView2ProcessExtendedInfoCollection {
	r := coreWebView2ProcessExtendedInfoCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ProcessExtendedInfoCollection(r)
}

var (
	coreWebView2ProcessExtendedInfoCollectionOnce   base.Once
	coreWebView2ProcessExtendedInfoCollectionImport *imports.Imports = nil
)

func coreWebView2ProcessExtendedInfoCollectionAPI() *imports.Imports {
	coreWebView2ProcessExtendedInfoCollectionOnce.Do(func() {
		coreWebView2ProcessExtendedInfoCollectionImport = api.NewDefaultImports()
		coreWebView2ProcessExtendedInfoCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProcessExtendedInfoCollection_Create", 0), // constructor NewCoreWebView2ProcessExtendedInfoCollection
			/* 1 */ imports.NewTable("TCoreWebView2ProcessExtendedInfoCollection_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ProcessExtendedInfoCollection_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ProcessExtendedInfoCollection_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2ProcessExtendedInfoCollection_Items", 0), // property Items
		}
	})
	return coreWebView2ProcessExtendedInfoCollectionImport
}
