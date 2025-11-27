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

// ICoreWebView2PermissionSettingCollectionView Parent: lcl.IObject
type ICoreWebView2PermissionSettingCollectionView interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionSettingCollectionView // property BaseIntf Getter
	// ValueAtIndex
	//  Gets the `ICoreWebView2PermissionSetting` at the specified index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview#getvalueatindex">See the ICoreWebView2PermissionSettingCollectionView article.</see>
	ValueAtIndex(idx uint32) ICoreWebView2PermissionSetting // property ValueAtIndex Getter
	// Count
	//  The number of `ICoreWebView2PermissionSetting`s in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview#get_count">See the ICoreWebView2PermissionSettingCollectionView article.</see>
	Count() uint32 // property Count Getter
}

type TCoreWebView2PermissionSettingCollectionView struct {
	lcl.TObject
}

func (m *TCoreWebView2PermissionSettingCollectionView) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionSettingCollectionViewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionSettingCollectionView) BaseIntf() (result ICoreWebView2PermissionSettingCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PermissionSettingCollectionViewAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PermissionSettingCollectionView(resultPtr)
	return
}

func (m *TCoreWebView2PermissionSettingCollectionView) ValueAtIndex(idx uint32) (result ICoreWebView2PermissionSetting) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PermissionSettingCollectionViewAPI().SysCallN(3, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PermissionSetting(resultPtr)
	return
}

func (m *TCoreWebView2PermissionSettingCollectionView) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PermissionSettingCollectionViewAPI().SysCallN(4, m.Instance())
	return uint32(r)
}

// NewCoreWebView2PermissionSettingCollectionView class constructor
func NewCoreWebView2PermissionSettingCollectionView(baseIntf ICoreWebView2PermissionSettingCollectionView) ICoreWebView2PermissionSettingCollectionView {
	r := coreWebView2PermissionSettingCollectionViewAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2PermissionSettingCollectionView(r)
}

var (
	coreWebView2PermissionSettingCollectionViewOnce   base.Once
	coreWebView2PermissionSettingCollectionViewImport *imports.Imports = nil
)

func coreWebView2PermissionSettingCollectionViewAPI() *imports.Imports {
	coreWebView2PermissionSettingCollectionViewOnce.Do(func() {
		coreWebView2PermissionSettingCollectionViewImport = api.NewDefaultImports()
		coreWebView2PermissionSettingCollectionViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PermissionSettingCollectionView_Create", 0), // constructor NewCoreWebView2PermissionSettingCollectionView
			/* 1 */ imports.NewTable("TCoreWebView2PermissionSettingCollectionView_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2PermissionSettingCollectionView_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2PermissionSettingCollectionView_ValueAtIndex", 0), // property ValueAtIndex
			/* 4 */ imports.NewTable("TCoreWebView2PermissionSettingCollectionView_Count", 0), // property Count
		}
	})
	return coreWebView2PermissionSettingCollectionViewImport
}
