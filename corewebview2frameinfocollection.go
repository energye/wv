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

// ICoreWebView2FrameInfoCollection Parent: IObject
//
//	Collection of FrameInfos (name and source). Used to list the affected
//	frames' info when a frame-only render process failure occurs in the
//	ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection">See the ICoreWebView2FrameInfoCollection article.</a>
type ICoreWebView2FrameInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfoCollection // property
	// Iterator
	//  Gets an iterator over the collection of `FrameInfo`s.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection#getiterator">See the ICoreWebView2FrameInfoCollection article.</a>
	Iterator() ICoreWebView2FrameInfoCollectionIterator // property
}

// TCoreWebView2FrameInfoCollection Parent: TObject
//
//	Collection of FrameInfos (name and source). Used to list the affected
//	frames' info when a frame-only render process failure occurs in the
//	ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection">See the ICoreWebView2FrameInfoCollection article.</a>
type TCoreWebView2FrameInfoCollection struct {
	TObject
}

func NewCoreWebView2FrameInfoCollection(aBaseIntf ICoreWebView2FrameInfoCollection) ICoreWebView2FrameInfoCollection {
	r1 := coreWebView2FrameInfoCollectionImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2FrameInfoCollection(r1)
}

func (m *TCoreWebView2FrameInfoCollection) Initialized() bool {
	r1 := coreWebView2FrameInfoCollectionImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameInfoCollection) BaseIntf() ICoreWebView2FrameInfoCollection {
	var resultCoreWebView2FrameInfoCollection uintptr
	coreWebView2FrameInfoCollectionImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollection)))
	return AsCoreWebView2FrameInfoCollection(resultCoreWebView2FrameInfoCollection)
}

func (m *TCoreWebView2FrameInfoCollection) Iterator() ICoreWebView2FrameInfoCollectionIterator {
	var resultCoreWebView2FrameInfoCollectionIterator uintptr
	coreWebView2FrameInfoCollectionImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollectionIterator)))
	return AsCoreWebView2FrameInfoCollectionIterator(resultCoreWebView2FrameInfoCollectionIterator)
}

var (
	coreWebView2FrameInfoCollectionImport       *imports.Imports = nil
	coreWebView2FrameInfoCollectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2FrameInfoCollection_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2FrameInfoCollection_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2FrameInfoCollection_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2FrameInfoCollection_Iterator", 0),
	}
)

func coreWebView2FrameInfoCollectionImportAPI() *imports.Imports {
	if coreWebView2FrameInfoCollectionImport == nil {
		coreWebView2FrameInfoCollectionImport = NewDefaultImports()
		coreWebView2FrameInfoCollectionImport.SetImportTable(coreWebView2FrameInfoCollectionImportTables)
		coreWebView2FrameInfoCollectionImportTables = nil
	}
	return coreWebView2FrameInfoCollectionImport
}
