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

// ICoreWebView2FrameInfoCollectionIterator Parent: IObject
//
//	Iterator for a collection of FrameInfos. For more info, see
//	ICoreWebView2ProcessFailedEventArgs2 and
//	ICoreWebView2FrameInfoCollection.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator">See the ICoreWebView2FrameInfoCollectionIterator article.</a>
type ICoreWebView2FrameInfoCollectionIterator interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfoCollectionIterator // property
	// HasCurrent
	//  `TRUE` when the iterator has not run out of `FrameInfo`s. If the
	//  collection over which the iterator is iterating is empty or if the
	//  iterator has gone past the end of the collection, then this is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#get_hascurrent">See the ICoreWebView2FrameInfoCollectionIterator article.</a>
	HasCurrent() bool // property
	// Current
	//  Get the current `ICoreWebView2FrameInfo` of the iterator.
	//  Returns `HRESULT_FROM_WIN32(ERROR_INVALID_INDEX)` if HasCurrent is
	//  `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#getcurrent">See the ICoreWebView2FrameInfoCollectionIterator article.</a>
	Current() ICoreWebView2FrameInfo // property
	// MoveNext
	//  Move the iterator to the next `FrameInfo` in the collection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#movenext">See the ICoreWebView2FrameInfoCollectionIterator article.</a>
	MoveNext() bool // function
}

// TCoreWebView2FrameInfoCollectionIterator Parent: TObject
//
//	Iterator for a collection of FrameInfos. For more info, see
//	ICoreWebView2ProcessFailedEventArgs2 and
//	ICoreWebView2FrameInfoCollection.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator">See the ICoreWebView2FrameInfoCollectionIterator article.</a>
type TCoreWebView2FrameInfoCollectionIterator struct {
	TObject
}

func NewCoreWebView2FrameInfoCollectionIterator(aBaseIntf ICoreWebView2FrameInfoCollectionIterator) ICoreWebView2FrameInfoCollectionIterator {
	r1 := coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2FrameInfoCollectionIterator(r1)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) Initialized() bool {
	r1 := coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) BaseIntf() ICoreWebView2FrameInfoCollectionIterator {
	var resultCoreWebView2FrameInfoCollectionIterator uintptr
	coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollectionIterator)))
	return AsCoreWebView2FrameInfoCollectionIterator(resultCoreWebView2FrameInfoCollectionIterator)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) HasCurrent() bool {
	r1 := coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) Current() ICoreWebView2FrameInfo {
	var resultCoreWebView2FrameInfo uintptr
	coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfo)))
	return AsCoreWebView2FrameInfo(resultCoreWebView2FrameInfo)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) MoveNext() bool {
	r1 := coreWebView2FrameInfoCollectionIteratorImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

var (
	coreWebView2FrameInfoCollectionIteratorImport       *imports.Imports = nil
	coreWebView2FrameInfoCollectionIteratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_Current", 0),
		/*3*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_HasCurrent", 0),
		/*4*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2FrameInfoCollectionIterator_MoveNext", 0),
	}
)

func coreWebView2FrameInfoCollectionIteratorImportAPI() *imports.Imports {
	if coreWebView2FrameInfoCollectionIteratorImport == nil {
		coreWebView2FrameInfoCollectionIteratorImport = NewDefaultImports()
		coreWebView2FrameInfoCollectionIteratorImport.SetImportTable(coreWebView2FrameInfoCollectionIteratorImportTables)
		coreWebView2FrameInfoCollectionIteratorImportTables = nil
	}
	return coreWebView2FrameInfoCollectionIteratorImport
}
