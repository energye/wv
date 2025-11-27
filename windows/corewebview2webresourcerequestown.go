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
	"github.com/energye/lcl/types"
)

// ICoreWebView2WebResourceRequest Parent: lcl.IInterfacedObject
type ICoreWebView2WebResourceRequest interface {
	lcl.IInterfacedObject
	// Get_uri
	//  The request URI.
	//
	//  The caller must free the returned string with `CoTaskMemFree`. See
	//  [API Conventions](/microsoft-edge/webview2/concepts/win32-api-conventions#strings).
	Get_uri(outUri *string) types.HRESULT // function
	// Set_uri
	//  Sets the `Uri` property.
	Set_uri(uri string) types.HRESULT // function
	// Get_Method
	//  The HTTP request method.
	//
	//  The caller must free the returned string with `CoTaskMemFree`. See
	//  [API Conventions](/microsoft-edge/webview2/concepts/win32-api-conventions#strings).
	Get_Method(outMethod *string) types.HRESULT // function
	// Set_Method
	//  Sets the `Method` property.
	Set_Method(method string) types.HRESULT // function
	// Get_Content
	//  The HTTP request message body as stream. POST data should be here. If a
	//  stream is set, which overrides the message body, the stream must have
	//  all the content data available by the time the `WebResourceRequested`
	//  event deferral of this response is completed. Stream should be agile or
	//  be created from a background STA to prevent performance impact to the UI
	//  thread. `Null` means no content data. `IStream` semantics apply
	//  (return `S_OK` to `Read` runs until all data is exhausted).
	Get_Content(outContent *lcl.IStreamAdapter) types.HRESULT // function
	// Set_Content
	//  Sets the `Content` property.
	Set_Content(content lcl.IStreamAdapter) types.HRESULT // function
	// Get_Headers
	//  The mutable HTTP request headers
	Get_Headers(outHeaders *ICoreWebView2HttpRequestHeaders) types.HRESULT // function
}

// ICoreWebView2WebResourceRequestOwn Parent: ICoreWebView2WebResourceRequest
type ICoreWebView2WebResourceRequestOwn interface {
	ICoreWebView2WebResourceRequest
	AsIntfWebResourceRequest() uintptr
}

type TCoreWebView2WebResourceRequestOwn struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2WebResourceRequestOwn) Get_uri(outUri *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var uriPtr uintptr
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&uriPtr)))
	*outUri = api.GoStr(uriPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Set_uri(uri string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(2, m.Instance(), api.PasStr(uri))
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Get_Method(outMethod *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var methodPtr uintptr
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&methodPtr)))
	*outMethod = api.GoStr(methodPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Set_Method(method string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(4, m.Instance(), api.PasStr(method))
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Get_Content(outContent *lcl.IStreamAdapter) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var contentPtr uintptr
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&contentPtr)))
	*outContent = lcl.AsStreamAdapter(contentPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Set_Content(content lcl.IStreamAdapter) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(content))
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) Get_Headers(outHeaders *ICoreWebView2HttpRequestHeaders) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var headersPtr uintptr
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&headersPtr)))
	*outHeaders = AsCoreWebView2HttpRequestHeaders(headersPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceRequestOwn) AsIntfWebResourceRequest() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2WebResourceRequestOwn class constructor
func NewCoreWebView2WebResourceRequestOwn(uRI string, method string, content lcl.IStreamAdapter, headers ICoreWebView2HttpRequestHeaders) ICoreWebView2WebResourceRequestOwn {
	var webResourceRequestPtr uintptr // ICoreWebView2WebResourceRequest
	r := coreWebView2WebResourceRequestOwnAPI().SysCallN(0, api.PasStr(uRI), api.PasStr(method), base.GetObjectUintptr(content), base.GetObjectUintptr(headers), uintptr(base.UnsafePointer(&webResourceRequestPtr)))
	ret := AsCoreWebView2WebResourceRequestOwn(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, webResourceRequestPtr)
	}
	return ret
}

var (
	coreWebView2WebResourceRequestOwnOnce   base.Once
	coreWebView2WebResourceRequestOwnImport *imports.Imports = nil
)

func coreWebView2WebResourceRequestOwnAPI() *imports.Imports {
	coreWebView2WebResourceRequestOwnOnce.Do(func() {
		coreWebView2WebResourceRequestOwnImport = api.NewDefaultImports()
		coreWebView2WebResourceRequestOwnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Create", 0), // constructor NewCoreWebView2WebResourceRequestOwn
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Get_uri", 0), // function Get_uri
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Set_uri", 0), // function Set_uri
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Get_Method", 0), // function Get_Method
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Set_Method", 0), // function Set_Method
			/* 5 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Get_Content", 0), // function Get_Content
			/* 6 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Set_Content", 0), // function Set_Content
			/* 7 */ imports.NewTable("TCoreWebView2WebResourceRequestOwn_Get_Headers", 0), // function Get_Headers
		}
	})
	return coreWebView2WebResourceRequestOwnImport
}
