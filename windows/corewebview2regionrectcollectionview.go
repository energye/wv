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
	"github.com/energye/lcl/types"
)

// ICoreWebView2RegionRectCollectionView Parent: lcl.IObject
type ICoreWebView2RegionRectCollectionView interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2RegionRectCollectionView // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2regionrectcollectionview#get_count">See the ICoreWebView2RegionRectCollectionView article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the element at the given index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2regionrectcollectionview#getvalueatindex">See the ICoreWebView2RegionRectCollectionView article.</see>
	Items(idx uint32) types.TRect // property Items Getter
}

type TCoreWebView2RegionRectCollectionView struct {
	lcl.TObject
}

func (m *TCoreWebView2RegionRectCollectionView) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2RegionRectCollectionViewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2RegionRectCollectionView) BaseIntf() (result ICoreWebView2RegionRectCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2RegionRectCollectionViewAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2RegionRectCollectionView(resultPtr)
	return
}

func (m *TCoreWebView2RegionRectCollectionView) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2RegionRectCollectionViewAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2RegionRectCollectionView) Items(idx uint32) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2RegionRectCollectionViewAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&result)))
	return
}

// NewCoreWebView2RegionRectCollectionView class constructor
func NewCoreWebView2RegionRectCollectionView(baseIntf ICoreWebView2RegionRectCollectionView) ICoreWebView2RegionRectCollectionView {
	r := coreWebView2RegionRectCollectionViewAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2RegionRectCollectionView(r)
}

var (
	coreWebView2RegionRectCollectionViewOnce   base.Once
	coreWebView2RegionRectCollectionViewImport *imports.Imports = nil
)

func coreWebView2RegionRectCollectionViewAPI() *imports.Imports {
	coreWebView2RegionRectCollectionViewOnce.Do(func() {
		coreWebView2RegionRectCollectionViewImport = api.NewDefaultImports()
		coreWebView2RegionRectCollectionViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2RegionRectCollectionView_Create", 0), // constructor NewCoreWebView2RegionRectCollectionView
			/* 1 */ imports.NewTable("TCoreWebView2RegionRectCollectionView_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2RegionRectCollectionView_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2RegionRectCollectionView_Count", 0), // property Count
			/* 4 */ imports.NewTable("TCoreWebView2RegionRectCollectionView_Items", 0), // property Items
		}
	})
	return coreWebView2RegionRectCollectionViewImport
}
