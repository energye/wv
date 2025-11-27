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

// ICoreWebView2HttpHeadersCollectionIterator Parent: lcl.IObject
type ICoreWebView2HttpHeadersCollectionIterator interface {
	lcl.IObject
	// GetCurrentHeader
	//  Get the name and value of the current HTTP header of the iterator. If
	//  the previous `MoveNext` operation set the `hasNext` parameter to `FALSE`,
	//  this method fails.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#getcurrentheader">See the ICoreWebView2HttpHeadersCollectionIterator article.</see>
	GetCurrentHeader(name *string, value *string) bool // function
	// MoveNext
	//  Move the iterator to the next HTTP header in the collection.
	//  \> [!NOTE]\n \> If no more HTTP headers exist, the `hasNext` parameter is set to
	//  `FALSE`. After this occurs the `GetCurrentHeader` method fails.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#movenext">See the ICoreWebView2HttpHeadersCollectionIterator article.</see>
	MoveNext() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpHeadersCollectionIterator // property BaseIntf Getter
	// HasCurrentHeader
	//  `TRUE` when the iterator has not run out of headers. If the collection
	//  over which the iterator is iterating is empty or if the iterator has gone
	//  past the end of the collection then this is `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#get_hascurrentheader">See the ICoreWebView2HttpHeadersCollectionIterator article.</see>
	HasCurrentHeader() bool // property HasCurrentHeader Getter
}

type TCoreWebView2HttpHeadersCollectionIterator struct {
	lcl.TObject
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) GetCurrentHeader(name *string, value *string) bool {
	if !m.IsValid() {
		return false
	}
	namePtr := api.PasStr(*name)
	valuePtr := api.PasStr(*value)
	r := coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&namePtr)), uintptr(base.UnsafePointer(&valuePtr)))
	*name = api.GoStr(namePtr)
	*value = api.GoStr(valuePtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) BaseIntf() (result ICoreWebView2HttpHeadersCollectionIterator) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpHeadersCollectionIterator(resultPtr)
	return
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) HasCurrentHeader() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2HttpHeadersCollectionIterator class constructor
func NewCoreWebView2HttpHeadersCollectionIterator(baseIntf ICoreWebView2HttpHeadersCollectionIterator) ICoreWebView2HttpHeadersCollectionIterator {
	r := coreWebView2HttpHeadersCollectionIteratorAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2HttpHeadersCollectionIterator(r)
}

var (
	coreWebView2HttpHeadersCollectionIteratorOnce   base.Once
	coreWebView2HttpHeadersCollectionIteratorImport *imports.Imports = nil
)

func coreWebView2HttpHeadersCollectionIteratorAPI() *imports.Imports {
	coreWebView2HttpHeadersCollectionIteratorOnce.Do(func() {
		coreWebView2HttpHeadersCollectionIteratorImport = api.NewDefaultImports()
		coreWebView2HttpHeadersCollectionIteratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_Create", 0), // constructor NewCoreWebView2HttpHeadersCollectionIterator
			/* 1 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_GetCurrentHeader", 0), // function GetCurrentHeader
			/* 2 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_Initialized", 0), // property Initialized
			/* 4 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_BaseIntf", 0), // property BaseIntf
			/* 5 */ imports.NewTable("TCoreWebView2HttpHeadersCollectionIterator_HasCurrentHeader", 0), // property HasCurrentHeader
		}
	})
	return coreWebView2HttpHeadersCollectionIteratorImport
}
