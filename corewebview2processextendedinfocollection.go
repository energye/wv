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

// ICoreWebView2ProcessExtendedInfoCollection Parent: IObject
//
//	A list containing processInfo and associated extended information.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
type ICoreWebView2ProcessExtendedInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessExtendedInfoCollection // property
	// Count
	//  The number of process contained in the `ICoreWebView2ProcessExtendedInfoCollection`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#get_count">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the `ICoreWebView2ProcessExtendedInfo` located in the
	//  `ICoreWebView2ProcessExtendedInfoCollection` at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection#getvalueatindex">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
	Items(idx uint32) ICoreWebView2ProcessExtendedInfo // property
}

// TCoreWebView2ProcessExtendedInfoCollection Parent: TObject
//
//	A list containing processInfo and associated extended information.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfocollection">See the ICoreWebView2ProcessExtendedInfoCollection article.</a>
type TCoreWebView2ProcessExtendedInfoCollection struct {
	TObject
}

func NewCoreWebView2ProcessExtendedInfoCollection(aBaseIntf ICoreWebView2ProcessExtendedInfoCollection) ICoreWebView2ProcessExtendedInfoCollection {
	r1 := coreWebView2ProcessExtendedInfoCollectionImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessExtendedInfoCollection(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Initialized() bool {
	r1 := coreWebView2ProcessExtendedInfoCollectionImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) BaseIntf() ICoreWebView2ProcessExtendedInfoCollection {
	var resultCoreWebView2ProcessExtendedInfoCollection uintptr
	coreWebView2ProcessExtendedInfoCollectionImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessExtendedInfoCollection)))
	return AsCoreWebView2ProcessExtendedInfoCollection(resultCoreWebView2ProcessExtendedInfoCollection)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Count() uint32 {
	r1 := coreWebView2ProcessExtendedInfoCollectionImportAPI().SysCallN(1, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ProcessExtendedInfoCollection) Items(idx uint32) ICoreWebView2ProcessExtendedInfo {
	var resultCoreWebView2ProcessExtendedInfo uintptr
	coreWebView2ProcessExtendedInfoCollectionImportAPI().SysCallN(4, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2ProcessExtendedInfo)))
	return AsCoreWebView2ProcessExtendedInfo(resultCoreWebView2ProcessExtendedInfo)
}

var (
	coreWebView2ProcessExtendedInfoCollectionImport       *imports.Imports = nil
	coreWebView2ProcessExtendedInfoCollectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ProcessExtendedInfoCollection_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ProcessExtendedInfoCollection_Count", 0),
		/*2*/ imports.NewTable("CoreWebView2ProcessExtendedInfoCollection_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2ProcessExtendedInfoCollection_Initialized", 0),
		/*4*/ imports.NewTable("CoreWebView2ProcessExtendedInfoCollection_Items", 0),
	}
)

func coreWebView2ProcessExtendedInfoCollectionImportAPI() *imports.Imports {
	if coreWebView2ProcessExtendedInfoCollectionImport == nil {
		coreWebView2ProcessExtendedInfoCollectionImport = NewDefaultImports()
		coreWebView2ProcessExtendedInfoCollectionImport.SetImportTable(coreWebView2ProcessExtendedInfoCollectionImportTables)
		coreWebView2ProcessExtendedInfoCollectionImportTables = nil
	}
	return coreWebView2ProcessExtendedInfoCollectionImport
}
