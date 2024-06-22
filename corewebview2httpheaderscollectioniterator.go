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

// ICoreWebView2HttpHeadersCollectionIterator Parent: IObject
//
//	Iterator for a collection of HTTP headers.  For more information, navigate
//	to ICoreWebView2HttpRequestHeaders and ICoreWebView2HttpResponseHeaders.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator">See the ICoreWebView2HttpHeadersCollectionIterator article.</a>
type ICoreWebView2HttpHeadersCollectionIterator interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpHeadersCollectionIterator // property
	// HasCurrentHeader
	//  `TRUE` when the iterator has not run out of headers. If the collection
	//  over which the iterator is iterating is empty or if the iterator has gone
	//  past the end of the collection then this is `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#get_hascurrentheader">See the ICoreWebView2HttpHeadersCollectionIterator article.</a>
	HasCurrentHeader() bool // property
	// GetCurrentHeader
	//  Get the name and value of the current HTTP header of the iterator. If
	//  the previous `MoveNext` operation set the `hasNext` parameter to `FALSE`,
	//  this method fails.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#getcurrentheader">See the ICoreWebView2HttpHeadersCollectionIterator article.</a>
	GetCurrentHeader(aName, aValue *string) bool // function
	// MoveNext
	//  Move the iterator to the next HTTP header in the collection.
	//   [!NOTE]\n \> If no more HTTP headers exist, the `hasNext` parameter is set to
	//  `FALSE`. After this occurs the `GetCurrentHeader` method fails.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator#movenext">See the ICoreWebView2HttpHeadersCollectionIterator article.</a>
	MoveNext() bool // function
}

// TCoreWebView2HttpHeadersCollectionIterator Parent: TObject
//
//	Iterator for a collection of HTTP headers.  For more information, navigate
//	to ICoreWebView2HttpRequestHeaders and ICoreWebView2HttpResponseHeaders.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpheaderscollectioniterator">See the ICoreWebView2HttpHeadersCollectionIterator article.</a>
type TCoreWebView2HttpHeadersCollectionIterator struct {
	TObject
}

func NewCoreWebView2HttpHeadersCollectionIterator(aBaseIntf ICoreWebView2HttpHeadersCollectionIterator) ICoreWebView2HttpHeadersCollectionIterator {
	r1 := coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2HttpHeadersCollectionIterator(r1)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) Initialized() bool {
	r1 := coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) BaseIntf() ICoreWebView2HttpHeadersCollectionIterator {
	var resultCoreWebView2HttpHeadersCollectionIterator uintptr
	coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpHeadersCollectionIterator)))
	return AsCoreWebView2HttpHeadersCollectionIterator(resultCoreWebView2HttpHeadersCollectionIterator)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) HasCurrentHeader() bool {
	r1 := coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) GetCurrentHeader(aName, aValue *string) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*aName = GoStr(result0)
	*aValue = GoStr(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2HttpHeadersCollectionIterator) MoveNext() bool {
	r1 := coreWebView2HttpHeadersCollectionIteratorImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

var (
	coreWebView2HttpHeadersCollectionIteratorImport       *imports.Imports = nil
	coreWebView2HttpHeadersCollectionIteratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_GetCurrentHeader", 0),
		/*3*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_HasCurrentHeader", 0),
		/*4*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2HttpHeadersCollectionIterator_MoveNext", 0),
	}
)

func coreWebView2HttpHeadersCollectionIteratorImportAPI() *imports.Imports {
	if coreWebView2HttpHeadersCollectionIteratorImport == nil {
		coreWebView2HttpHeadersCollectionIteratorImport = NewDefaultImports()
		coreWebView2HttpHeadersCollectionIteratorImport.SetImportTable(coreWebView2HttpHeadersCollectionIteratorImportTables)
		coreWebView2HttpHeadersCollectionIteratorImportTables = nil
	}
	return coreWebView2HttpHeadersCollectionIteratorImport
}
