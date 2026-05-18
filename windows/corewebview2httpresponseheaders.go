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

// ICoreWebView2HttpResponseHeaders Parent: IObject
type ICoreWebView2HttpResponseHeaders interface {
	IObject
	// GetHeader
	//  Gets the first header value in the collection matching the name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheader">See the ICoreWebView2HttpResponseHeaders article.</see>
	GetHeader(name string) string // function
	// GetHeaders
	//  Gets the header values matching the name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getheaders">See the ICoreWebView2HttpResponseHeaders article.</see>
	GetHeaders(name string, iterator *ICoreWebView2HttpHeadersCollectionIterator) bool // function
	// Contains
	//  Verifies that the headers contain entries that match the header name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#contains">See the ICoreWebView2HttpResponseHeaders article.</see>
	Contains(name string) bool // function
	// AppendHeader
	//  Appends header line with name and value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#appendheader">See the ICoreWebView2HttpResponseHeaders article.</see>
	AppendHeader(name string, value string) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpResponseHeaders // property BaseIntf Getter
	// Iterator
	//  Gets an iterator over the collection of entire response headers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httpresponseheaders#getiterator">See the ICoreWebView2HttpResponseHeaders article.</see>
	Iterator() ICoreWebView2HttpHeadersCollectionIterator // property Iterator Getter
}

type TCoreWebView2HttpResponseHeaders struct {
	TObject
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeader(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2HttpResponseHeadersAPI().SysCallN(1, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2HttpResponseHeaders) GetHeaders(name string, iterator *ICoreWebView2HttpHeadersCollectionIterator) bool {
	if !m.IsValid() {
		return false
	}
	iteratorPtr := base.GetObjectUintptr(*iterator)
	r := coreWebView2HttpResponseHeadersAPI().SysCallN(2, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&iteratorPtr)))
	*iterator = AsCoreWebView2HttpHeadersCollectionIterator(iteratorPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpResponseHeaders) Contains(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpResponseHeadersAPI().SysCallN(3, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpResponseHeaders) AppendHeader(name string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpResponseHeadersAPI().SysCallN(4, m.Instance(), api.PasStr(name), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpResponseHeaders) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpResponseHeadersAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpResponseHeaders) BaseIntf() (result ICoreWebView2HttpResponseHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2HttpResponseHeadersAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpResponseHeaders(resultPtr)
	return
}

func (m *TCoreWebView2HttpResponseHeaders) Iterator() (result ICoreWebView2HttpHeadersCollectionIterator) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2HttpResponseHeadersAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpHeadersCollectionIterator(resultPtr)
	return
}

// NewCoreWebView2HttpResponseHeaders class constructor
func NewCoreWebView2HttpResponseHeaders(baseIntf ICoreWebView2HttpResponseHeaders) ICoreWebView2HttpResponseHeaders {
	r := coreWebView2HttpResponseHeadersAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2HttpResponseHeaders(r)
}

var (
	coreWebView2HttpResponseHeadersOnce   base.Once
	coreWebView2HttpResponseHeadersImport *imports.Imports = nil
)

func coreWebView2HttpResponseHeadersAPI() *imports.Imports {
	coreWebView2HttpResponseHeadersOnce.Do(func() {
		coreWebView2HttpResponseHeadersImport = api.NewDefaultImports()
		coreWebView2HttpResponseHeadersImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_Create", 0), // constructor NewCoreWebView2HttpResponseHeaders
			/* 1 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_GetHeader", 0), // function GetHeader
			/* 2 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_GetHeaders", 0), // function GetHeaders
			/* 3 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_Contains", 0), // function Contains
			/* 4 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_AppendHeader", 0), // function AppendHeader
			/* 5 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_Initialized", 0), // property Initialized
			/* 6 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_BaseIntf", 0), // property BaseIntf
			/* 7 */ imports.NewTable("TCoreWebView2HttpResponseHeaders_Iterator", 0), // property Iterator
		}
	})
	return coreWebView2HttpResponseHeadersImport
}
