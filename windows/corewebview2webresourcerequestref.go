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

// ICoreWebView2WebResourceRequestRef Parent: IObject
type ICoreWebView2WebResourceRequestRef interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceRequest // property BaseIntf Getter
	// URI
	//  The request URI.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_uri">See the ICoreWebView2WebResourceRequest article.</see>
	URI() string         // property URI Getter
	SetURI(value string) // property URI Setter
	// Method
	//  The HTTP request method.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_method">See the ICoreWebView2WebResourceRequest article.</see>
	Method() string         // property Method Getter
	SetMethod(value string) // property Method Setter
	// Content
	//  The HTTP request message body as stream. POST data should be here. If a
	//  stream is set, which overrides the message body, the stream must have
	//  all the content data available by the time the `WebResourceRequested`
	//  event deferral of this response is completed. Stream should be agile or
	//  be created from a background STA to prevent performance impact to the UI
	//  thread. `Null` means no content data. `IStream` semantics apply
	//  (return `S_OK` to `Read` runs until all data is exhausted).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_content">See the ICoreWebView2WebResourceRequest article.</see>
	Content() lcl.IStreamAdapter         // property Content Getter
	SetContent(value lcl.IStreamAdapter) // property Content Setter
	// Headers
	//  The mutable HTTP request headers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequest#get_headers">See the ICoreWebView2WebResourceRequest article.</see>
	Headers() ICoreWebView2HttpRequestHeaders // property Headers Getter
}

type TCoreWebView2WebResourceRequestRef struct {
	TObject
}

func (m *TCoreWebView2WebResourceRequestRef) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceRequestRefAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceRequestRef) BaseIntf() (result ICoreWebView2WebResourceRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestRefAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceRequestOwn(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceRequestRef) URI() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebResourceRequestRefAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebResourceRequestRef) SetURI(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceRequestRefAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2WebResourceRequestRef) Method() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2WebResourceRequestRefAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2WebResourceRequestRef) SetMethod(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceRequestRefAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2WebResourceRequestRef) Content() lcl.IStreamAdapter {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestRefAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	return lcl.AsStreamAdapter(resultPtr)
}

func (m *TCoreWebView2WebResourceRequestRef) SetContent(value lcl.IStreamAdapter) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceRequestRefAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2WebResourceRequestRef) Headers() (result ICoreWebView2HttpRequestHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestRefAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpRequestHeaders(resultPtr)
	return
}

// NewCoreWebView2WebResourceRequestRef class constructor
func NewCoreWebView2WebResourceRequestRef(baseIntf ICoreWebView2WebResourceRequest) ICoreWebView2WebResourceRequestRef {
	r := coreWebView2WebResourceRequestRefAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2WebResourceRequestRef(r)
}

var (
	coreWebView2WebResourceRequestRefOnce   base.Once
	coreWebView2WebResourceRequestRefImport *imports.Imports = nil
)

func coreWebView2WebResourceRequestRefAPI() *imports.Imports {
	coreWebView2WebResourceRequestRefOnce.Do(func() {
		coreWebView2WebResourceRequestRefImport = api.NewDefaultImports()
		coreWebView2WebResourceRequestRefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_Create", 0), // constructor NewCoreWebView2WebResourceRequestRef
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_URI", 0), // property URI
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_Method", 0), // property Method
			/* 5 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_Content", 0), // property Content
			/* 6 */ imports.NewTable("TCoreWebView2WebResourceRequestRef_Headers", 0), // property Headers
		}
	})
	return coreWebView2WebResourceRequestRefImport
}
