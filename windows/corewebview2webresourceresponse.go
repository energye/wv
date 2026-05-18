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

// ICoreWebView2WebResourceResponse Parent: IObject
type ICoreWebView2WebResourceResponse interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponse // property BaseIntf Getter
	// StatusCode
	//  The HTTP response status code.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_statuscode">See the ICoreWebView2WebResourceResponse article.</see>
	StatusCode() int32         // property StatusCode Getter
	SetStatusCode(value int32) // property StatusCode Setter
	// ReasonPhrase
	//  The HTTP response reason phrase.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_reasonphrase">See the ICoreWebView2WebResourceResponse article.</see>
	ReasonPhrase() string         // property ReasonPhrase Getter
	SetReasonPhrase(value string) // property ReasonPhrase Setter
	// Content
	//  HTTP response content as stream. Stream must have all the content data
	//  available by the time the `WebResourceRequested` event deferral of this
	//  response is completed. Stream should be agile or be created from a
	//  background thread to prevent performance impact to the UI thread. `Null`
	//  means no content data. `IStream` semantics apply (return `S_OK` to
	//  `Read` runs until all data is exhausted).
	//  When providing the response data, you should consider relevant HTTP
	//  request headers just like an HTTP server would do. For example, if the
	//  request was for a video resource in a HTML video element, the request may
	//  contain the [Range](https://developer.mozilla.org/docs/Web/HTTP/Headers/Range)
	//  header to request only a part of the video that is streaming. In this
	//  case, your response stream should be only the portion of the video
	//  specified by the range HTTP request headers and you should set the
	//  appropriate
	//  [Content-Range](https://developer.mozilla.org/docs/Web/HTTP/Headers/Content-Range)
	//  header in the response.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_content">See the ICoreWebView2WebResourceResponse article.</see>
	Content() lcl.IStreamAdapter         // property Content Getter
	SetContent(value lcl.IStreamAdapter) // property Content Setter
	// Headers
	//  Overridden HTTP response headers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponse#get_headers">See the ICoreWebView2WebResourceResponse article.</see>
	Headers() ICoreWebView2HttpResponseHeaders // property Headers Getter
}

type TCoreWebView2WebResourceResponse struct {
	TObject
}

func (m *TCoreWebView2WebResourceResponse) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceResponseAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceResponse) BaseIntf() (result ICoreWebView2WebResourceResponse) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceResponse(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceResponse) StatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceResponseAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2WebResourceResponse) SetStatusCode(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceResponseAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2WebResourceResponse) ReasonPhrase() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2WebResourceResponseAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2WebResourceResponse) SetReasonPhrase(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceResponseAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2WebResourceResponse) Content() lcl.IStreamAdapter {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	return lcl.AsStreamAdapter(resultPtr)
}

func (m *TCoreWebView2WebResourceResponse) SetContent(value lcl.IStreamAdapter) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceResponseAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2WebResourceResponse) Headers() (result ICoreWebView2HttpResponseHeaders) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2HttpResponseHeaders(resultPtr)
	return
}

// NewCoreWebView2WebResourceResponse class constructor
func NewCoreWebView2WebResourceResponse(baseIntf ICoreWebView2WebResourceResponse) ICoreWebView2WebResourceResponse {
	r := coreWebView2WebResourceResponseAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2WebResourceResponse(r)
}

var (
	coreWebView2WebResourceResponseOnce   base.Once
	coreWebView2WebResourceResponseImport *imports.Imports = nil
)

func coreWebView2WebResourceResponseAPI() *imports.Imports {
	coreWebView2WebResourceResponseOnce.Do(func() {
		coreWebView2WebResourceResponseImport = api.NewDefaultImports()
		coreWebView2WebResourceResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceResponse_Create", 0), // constructor NewCoreWebView2WebResourceResponse
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceResponse_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceResponse_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceResponse_StatusCode", 0), // property StatusCode
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceResponse_ReasonPhrase", 0), // property ReasonPhrase
			/* 5 */ imports.NewTable("TCoreWebView2WebResourceResponse_Content", 0), // property Content
			/* 6 */ imports.NewTable("TCoreWebView2WebResourceResponse_Headers", 0), // property Headers
		}
	})
	return coreWebView2WebResourceResponseImport
}
