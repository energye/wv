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

// ICoreWebView2HttpResponseHeaders Parent: IObject
//
//	HTTP response headers.  Used to construct a WebResourceResponse for the
//	WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
type ICoreWebView2HttpResponseHeaders interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpResponseHeaders // property
	// Iterator
	//  Gets an iterator over the collection of entire response headers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getiterator">See the ICoreWebView2HttpResponseHeaders article.</a>
	Iterator() ICoreWebView2HttpHeadersCollectionIterator // property
	// GetHeader
	//  Gets the first header value in the collection matching the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheader">See the ICoreWebView2HttpResponseHeaders article.</a>
	GetHeader(aName string) string // function
	// GetHeaders
	//  Gets the header values matching the name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
	GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool // function
	// Contains
	//  Verifies that the headers contain entries that match the header name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#contains">See the ICoreWebView2HttpResponseHeaders article.</a>
	Contains(aName string) bool // function
	// AppendHeader
	//  Appends header line with name and value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#appendheader">See the ICoreWebView2HttpResponseHeaders article.</a>
	AppendHeader(aName, aValue string) bool // function
}

// TCoreWebView2HttpResponseHeaders Parent: TObject
//
//	HTTP response headers.  Used to construct a WebResourceResponse for the
//	WebResourceRequested event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders">See the ICoreWebView2HttpResponseHeaders article.</a>
type TCoreWebView2HttpResponseHeaders struct {
	TObject
}

func NewCoreWebView2HttpResponseHeaders(aBaseIntf ICoreWebView2HttpResponseHeaders) ICoreWebView2HttpResponseHeaders {
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(3, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2HttpResponseHeaders(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) Initialized() bool {
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) BaseIntf() ICoreWebView2HttpResponseHeaders {
	var resultCoreWebView2HttpResponseHeaders uintptr
	coreWebView2HttpResponseHeadersImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpResponseHeaders)))
	return AsCoreWebView2HttpResponseHeaders(resultCoreWebView2HttpResponseHeaders)
}

func (m *TCoreWebView2HttpResponseHeaders) Iterator() ICoreWebView2HttpHeadersCollectionIterator {
	var resultCoreWebView2HttpHeadersCollectionIterator uintptr
	coreWebView2HttpResponseHeadersImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2HttpHeadersCollectionIterator)))
	return AsCoreWebView2HttpHeadersCollectionIterator(resultCoreWebView2HttpHeadersCollectionIterator)
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeader(aName string) string {
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(4, m.Instance(), PascalStr(aName))
	return GoStr(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeaders(aName string, aIterator *ICoreWebView2HttpHeadersCollectionIterator) bool {
	var result1 uintptr
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(5, m.Instance(), PascalStr(aName), uintptr(unsafePointer(&result1)))
	*aIterator = AsCoreWebView2HttpHeadersCollectionIterator(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) Contains(aName string) bool {
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(2, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TCoreWebView2HttpResponseHeaders) AppendHeader(aName, aValue string) bool {
	r1 := coreWebView2HttpResponseHeadersImportAPI().SysCallN(0, m.Instance(), PascalStr(aName), PascalStr(aValue))
	return GoBool(r1)
}

var (
	coreWebView2HttpResponseHeadersImport       *imports.Imports = nil
	coreWebView2HttpResponseHeadersImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2HttpResponseHeaders_AppendHeader", 0),
		/*1*/ imports.NewTable("CoreWebView2HttpResponseHeaders_BaseIntf", 0),
		/*2*/ imports.NewTable("CoreWebView2HttpResponseHeaders_Contains", 0),
		/*3*/ imports.NewTable("CoreWebView2HttpResponseHeaders_Create", 0),
		/*4*/ imports.NewTable("CoreWebView2HttpResponseHeaders_GetHeader", 0),
		/*5*/ imports.NewTable("CoreWebView2HttpResponseHeaders_GetHeaders", 0),
		/*6*/ imports.NewTable("CoreWebView2HttpResponseHeaders_Initialized", 0),
		/*7*/ imports.NewTable("CoreWebView2HttpResponseHeaders_Iterator", 0),
	}
)

func coreWebView2HttpResponseHeadersImportAPI() *imports.Imports {
	if coreWebView2HttpResponseHeadersImport == nil {
		coreWebView2HttpResponseHeadersImport = NewDefaultImports()
		coreWebView2HttpResponseHeadersImport.SetImportTable(coreWebView2HttpResponseHeadersImportTables)
		coreWebView2HttpResponseHeadersImportTables = nil
	}
	return coreWebView2HttpResponseHeadersImport
}
