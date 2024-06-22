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

// ICoreWebView2StringCollection Parent: IObject
//
//	A collection of strings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection">See the ICoreWebView2StringCollection article.</a>
type ICoreWebView2StringCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2StringCollection // property
	// Count
	//  The number of strings contained in ICoreWebView2StringCollection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection#get_count">See the ICoreWebView2StringCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the value at a given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection#getvalueatindex">See the ICoreWebView2StringCollection article.</a>
	Items(idx uint32) string // property
}

// TCoreWebView2StringCollection Parent: TObject
//
//	A collection of strings.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2stringcollection">See the ICoreWebView2StringCollection article.</a>
type TCoreWebView2StringCollection struct {
	TObject
}

func NewCoreWebView2StringCollection(aBaseIntf ICoreWebView2StringCollection) ICoreWebView2StringCollection {
	r1 := coreWebView2StringCollectionImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2StringCollection(r1)
}

func (m *TCoreWebView2StringCollection) Initialized() bool {
	r1 := coreWebView2StringCollectionImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2StringCollection) BaseIntf() ICoreWebView2StringCollection {
	var resultCoreWebView2StringCollection uintptr
	coreWebView2StringCollectionImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2StringCollection)))
	return AsCoreWebView2StringCollection(resultCoreWebView2StringCollection)
}

func (m *TCoreWebView2StringCollection) Count() uint32 {
	r1 := coreWebView2StringCollectionImportAPI().SysCallN(1, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2StringCollection) Items(idx uint32) string {
	r1 := coreWebView2StringCollectionImportAPI().SysCallN(4, m.Instance(), uintptr(idx))
	return GoStr(r1)
}

var (
	coreWebView2StringCollectionImport       *imports.Imports = nil
	coreWebView2StringCollectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2StringCollection_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2StringCollection_Count", 0),
		/*2*/ imports.NewTable("CoreWebView2StringCollection_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2StringCollection_Initialized", 0),
		/*4*/ imports.NewTable("CoreWebView2StringCollection_Items", 0),
	}
)

func coreWebView2StringCollectionImportAPI() *imports.Imports {
	if coreWebView2StringCollectionImport == nil {
		coreWebView2StringCollectionImport = NewDefaultImports()
		coreWebView2StringCollectionImport.SetImportTable(coreWebView2StringCollectionImportTables)
		coreWebView2StringCollectionImportTables = nil
	}
	return coreWebView2StringCollectionImport
}
