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

// ICoreWebView2ObjectCollectionView Parent: lcl.IObject
type ICoreWebView2ObjectCollectionView interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ObjectCollectionView // property BaseIntf Getter
	// Count
	//  The number of elements contained in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview#get_count">See the ICoreWebView2ObjectCollectionView article.</see>
	Count() uint32 // property Count Getter
}

type TCoreWebView2ObjectCollectionView struct {
	lcl.TObject
}

func (m *TCoreWebView2ObjectCollectionView) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ObjectCollectionViewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ObjectCollectionView) BaseIntf() (result ICoreWebView2ObjectCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ObjectCollectionViewAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ObjectCollectionView(resultPtr)
	return
}

func (m *TCoreWebView2ObjectCollectionView) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ObjectCollectionViewAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

// NewCoreWebView2ObjectCollectionView class constructor
func NewCoreWebView2ObjectCollectionView(baseIntf ICoreWebView2ObjectCollectionView) ICoreWebView2ObjectCollectionView {
	r := coreWebView2ObjectCollectionViewAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ObjectCollectionView(r)
}

var (
	coreWebView2ObjectCollectionViewOnce   base.Once
	coreWebView2ObjectCollectionViewImport *imports.Imports = nil
)

func coreWebView2ObjectCollectionViewAPI() *imports.Imports {
	coreWebView2ObjectCollectionViewOnce.Do(func() {
		coreWebView2ObjectCollectionViewImport = api.NewDefaultImports()
		coreWebView2ObjectCollectionViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ObjectCollectionView_Create", 0), // constructor NewCoreWebView2ObjectCollectionView
			/* 1 */ imports.NewTable("TCoreWebView2ObjectCollectionView_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ObjectCollectionView_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ObjectCollectionView_Count", 0), // property Count
		}
	})
	return coreWebView2ObjectCollectionViewImport
}
