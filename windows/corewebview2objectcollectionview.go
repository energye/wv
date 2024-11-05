//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICoreWebView2ObjectCollectionView Parent: IObject
//
//	Read-only collection of generic objects.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview">See the ICoreWebView2ObjectCollectionView article.</a>
type ICoreWebView2ObjectCollectionView interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ObjectCollectionView // property
	// Count
	//  Gets the number of items in the collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview#get_count">See the ICoreWebView2ObjectCollectionView article.</a>
	Count() uint32 // property
	// Items
	//  Gets the object at the specified index. Cast the object to the native type
	//  to access its specific properties.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview#getvalueatindex">See the ICoreWebView2ObjectCollectionView article.</a>
	Items(idx uint32) IUnknown // property
}

// TCoreWebView2ObjectCollectionView Parent: TObject
//
//	Read-only collection of generic objects.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2objectcollectionview">See the ICoreWebView2ObjectCollectionView article.</a>
type TCoreWebView2ObjectCollectionView struct {
	TObject
}

func NewCoreWebView2ObjectCollectionView(aBaseIntf ICoreWebView2ObjectCollectionView) ICoreWebView2ObjectCollectionView {
	r1 := coreWebView2ObjectCollectionViewImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ObjectCollectionView(r1)
}

func (m *TCoreWebView2ObjectCollectionView) Initialized() bool {
	r1 := coreWebView2ObjectCollectionViewImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ObjectCollectionView) BaseIntf() ICoreWebView2ObjectCollectionView {
	var resultCoreWebView2ObjectCollectionView uintptr
	coreWebView2ObjectCollectionViewImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ObjectCollectionView)))
	return AsCoreWebView2ObjectCollectionView(resultCoreWebView2ObjectCollectionView)
}

func (m *TCoreWebView2ObjectCollectionView) Count() uint32 {
	r1 := coreWebView2ObjectCollectionViewImportAPI().SysCallN(1, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ObjectCollectionView) Items(idx uint32) IUnknown {
	var resultUnknown uintptr
	coreWebView2ObjectCollectionViewImportAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultUnknown)))
	return AsUnknown(resultUnknown)
}

var (
	coreWebView2ObjectCollectionViewImport       *imports.Imports = nil
	coreWebView2ObjectCollectionViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ObjectCollectionView_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ObjectCollectionView_Count", 0),
		/*2*/ imports.NewTable("CoreWebView2ObjectCollectionView_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2ObjectCollectionView_Initialized", 0),
		/*4*/ imports.NewTable("CoreWebView2ObjectCollectionView_Items", 0),
	}
)

func coreWebView2ObjectCollectionViewImportAPI() *imports.Imports {
	if coreWebView2ObjectCollectionViewImport == nil {
		coreWebView2ObjectCollectionViewImport = NewDefaultImports()
		coreWebView2ObjectCollectionViewImport.SetImportTable(coreWebView2ObjectCollectionViewImportTables)
		coreWebView2ObjectCollectionViewImportTables = nil
	}
	return coreWebView2ObjectCollectionViewImport
}
