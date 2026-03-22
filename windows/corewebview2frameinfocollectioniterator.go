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

// ICoreWebView2FrameInfoCollectionIterator Parent: IObject
type ICoreWebView2FrameInfoCollectionIterator interface {
	IObject
	// MoveNext
	//  Move the iterator to the next `FrameInfo` in the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#movenext">See the ICoreWebView2FrameInfoCollectionIterator article.</see>
	MoveNext() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfoCollectionIterator // property BaseIntf Getter
	// HasCurrent
	//  `TRUE` when the iterator has not run out of `FrameInfo`s. If the
	//  collection over which the iterator is iterating is empty or if the
	//  iterator has gone past the end of the collection, then this is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#get_hascurrent">See the ICoreWebView2FrameInfoCollectionIterator article.</see>
	HasCurrent() bool // property HasCurrent Getter
	// Current
	//  Get the current `ICoreWebView2FrameInfo` of the iterator.
	//  Returns `HRESULT_FROM_WIN32(ERROR_INVALID_INDEX)` if HasCurrent is
	//  `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator#getcurrent">See the ICoreWebView2FrameInfoCollectionIterator article.</see>
	Current() ICoreWebView2FrameInfo // property Current Getter
}

type TCoreWebView2FrameInfoCollectionIterator struct {
	TObject
}

func (m *TCoreWebView2FrameInfoCollectionIterator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) BaseIntf() (result ICoreWebView2FrameInfoCollectionIterator) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfoCollectionIterator(resultPtr)
	return
}

func (m *TCoreWebView2FrameInfoCollectionIterator) HasCurrent() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameInfoCollectionIterator) Current() (result ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfo(resultPtr)
	return
}

// NewCoreWebView2FrameInfoCollectionIterator class constructor
func NewCoreWebView2FrameInfoCollectionIterator(baseIntf ICoreWebView2FrameInfoCollectionIterator) ICoreWebView2FrameInfoCollectionIterator {
	r := coreWebView2FrameInfoCollectionIteratorAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2FrameInfoCollectionIterator(r)
}

var (
	coreWebView2FrameInfoCollectionIteratorOnce   base.Once
	coreWebView2FrameInfoCollectionIteratorImport *imports.Imports = nil
)

func coreWebView2FrameInfoCollectionIteratorAPI() *imports.Imports {
	coreWebView2FrameInfoCollectionIteratorOnce.Do(func() {
		coreWebView2FrameInfoCollectionIteratorImport = api.NewDefaultImports()
		coreWebView2FrameInfoCollectionIteratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_Create", 0), // constructor NewCoreWebView2FrameInfoCollectionIterator
			/* 1 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_HasCurrent", 0), // property HasCurrent
			/* 5 */ imports.NewTable("TCoreWebView2FrameInfoCollectionIterator_Current", 0), // property Current
		}
	})
	return coreWebView2FrameInfoCollectionIteratorImport
}
