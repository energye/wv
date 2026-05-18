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

// ICoreWebView2HttpRequestHeaders Parent: IObject
type ICoreWebView2HttpRequestHeaders interface {
	IObject
	// SetHeader
	//  Adds or updates header that matches the name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#setheader">See the ICoreWebView2HttpRequestHeaders article.</see>
	SetHeader(name string, value string) bool // function
	// GetHeader
	//  Gets the header value matching the name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getheader">See the ICoreWebView2HttpRequestHeaders article.</see>
	GetHeader(name string) string // function
	// GetHeaders
	//  Gets the header value matching the name using an iterator.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getheaders">See the ICoreWebView2HttpRequestHeaders article.</see>
	GetHeaders(name string, value *ICoreWebView2HttpHeadersCollectionIterator) bool // function
	// Contains
	//  Checks whether the headers contain an entry that matches the header name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#contains">See the ICoreWebView2HttpRequestHeaders article.</see>
	Contains(value string) bool // function
	// RemoveHeader
	//  Removes header that matches the name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#removeheader">See the ICoreWebView2HttpRequestHeaders article.</see>
	RemoveHeader(name string) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2HttpRequestHeaders // property BaseIntf Getter
	// Iterator
	//  Gets an iterator over the collection of request headers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2httprequestheaders#getiterator">See the ICoreWebView2HttpRequestHeaders article.</see>
	Iterator() ICoreWebView2HttpHeadersCollectionIterator // property Iterator Getter
}

type TCoreWebView2HttpRequestHeaders struct {
	TObject
}

func (m *TCoreWebView2HttpRequestHeaders) SetHeader(name string, value string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(1, m.Instance(), api.PasStr(name), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpRequestHeaders) GetHeader(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2HttpRequestHeadersAPI().SysCallN(2, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2HttpRequestHeaders) GetHeaders(name string, value *ICoreWebView2HttpHeadersCollectionIterator) bool {
	if !m.IsValid() {
		return false
	}
	valuePtr := base.GetObjectUintptr(*value)
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(3, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCoreWebView2HttpHeadersCollectionIterator(valuePtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpRequestHeaders) Contains(value string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(4, m.Instance(), api.PasStr(value))
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpRequestHeaders) RemoveHeader(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(5, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpRequestHeaders) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2HttpRequestHeaders) BaseIntf() (result ICoreWebView2HttpRequestHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2HttpRequestHeadersAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpRequestHeaders(resultPtr)
	return
}

func (m *TCoreWebView2HttpRequestHeaders) Iterator() (result ICoreWebView2HttpHeadersCollectionIterator) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2HttpRequestHeadersAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpHeadersCollectionIterator(resultPtr)
	return
}

// NewCoreWebView2HttpRequestHeaders class constructor
func NewCoreWebView2HttpRequestHeaders(baseIntf ICoreWebView2HttpRequestHeaders) ICoreWebView2HttpRequestHeaders {
	r := coreWebView2HttpRequestHeadersAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2HttpRequestHeaders(r)
}

var (
	coreWebView2HttpRequestHeadersOnce   base.Once
	coreWebView2HttpRequestHeadersImport *imports.Imports = nil
)

func coreWebView2HttpRequestHeadersAPI() *imports.Imports {
	coreWebView2HttpRequestHeadersOnce.Do(func() {
		coreWebView2HttpRequestHeadersImport = api.NewDefaultImports()
		coreWebView2HttpRequestHeadersImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_Create", 0), // constructor NewCoreWebView2HttpRequestHeaders
			/* 1 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_SetHeader", 0), // function SetHeader
			/* 2 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_GetHeader", 0), // function GetHeader
			/* 3 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_GetHeaders", 0), // function GetHeaders
			/* 4 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_Contains", 0), // function Contains
			/* 5 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_RemoveHeader", 0), // function RemoveHeader
			/* 6 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_Initialized", 0), // property Initialized
			/* 7 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_BaseIntf", 0), // property BaseIntf
			/* 8 */ imports.NewTable("TCoreWebView2HttpRequestHeaders_Iterator", 0), // property Iterator
		}
	})
	return coreWebView2HttpRequestHeadersImport
}
