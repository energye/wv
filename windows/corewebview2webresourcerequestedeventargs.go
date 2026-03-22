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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2WebResourceRequestedEventArgs Parent: IObject
type ICoreWebView2WebResourceRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceRequestedEventArgs // property BaseIntf Getter
	// Request
	//  The Web resource request. The request object may be missing some headers
	//  that are added by network stack at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_request">See the ICoreWebView2WebResourceRequestedEventArgs article.</see>
	Request() ICoreWebView2WebResourceRequest // property Request Getter
	// Response
	//  A placeholder for the web resource response object. If this object is
	//  set, the web resource request is completed with the specified response.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_response">See the ICoreWebView2WebResourceRequestedEventArgs article.</see>
	Response() ICoreWebView2WebResourceResponse         // property Response Getter
	SetResponse(value ICoreWebView2WebResourceResponse) // property Response Setter
	// Deferral
	//  Obtain an `ICoreWebView2Deferral` object and put the event into a
	//  deferred state. Use the `ICoreWebView2Deferral` object to complete the
	//  request at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#getdeferral">See the ICoreWebView2WebResourceRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
	// ResourceContext
	//  The web resource request context.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs#get_resourcecontext">See the ICoreWebView2WebResourceRequestedEventArgs article.</see>
	ResourceContext() wvTypes.TWVWebResourceContext // property ResourceContext Getter
	// RequestedSourceKind
	//  The web resource requested source.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourcerequestedeventargs2#get_requestedsourcekind">See the ICoreWebView2WebResourceRequestedEventArgs2 article.</see>
	RequestedSourceKind() wvTypes.TWVWebResourceRequestSourceKind // property RequestedSourceKind Getter
}

type TCoreWebView2WebResourceRequestedEventArgs struct {
	TObject
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) BaseIntf() (result ICoreWebView2WebResourceRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Request() (result ICoreWebView2WebResourceRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceRequestOwn(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Response() (result ICoreWebView2WebResourceResponse) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceResponse(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) SetResponse(value ICoreWebView2WebResourceResponse) {
	if !m.IsValid() {
		return
	}
	coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) ResourceContext() wvTypes.TWVWebResourceContext {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(6, m.Instance())
	return wvTypes.TWVWebResourceContext(r)
}

func (m *TCoreWebView2WebResourceRequestedEventArgs) RequestedSourceKind() wvTypes.TWVWebResourceRequestSourceKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(7, m.Instance())
	return wvTypes.TWVWebResourceRequestSourceKind(r)
}

// NewCoreWebView2WebResourceRequestedEventArgs class constructor
func NewCoreWebView2WebResourceRequestedEventArgs(args ICoreWebView2WebResourceRequestedEventArgs) ICoreWebView2WebResourceRequestedEventArgs {
	r := coreWebView2WebResourceRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2WebResourceRequestedEventArgs(r)
}

var (
	coreWebView2WebResourceRequestedEventArgsOnce   base.Once
	coreWebView2WebResourceRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2WebResourceRequestedEventArgsAPI() *imports.Imports {
	coreWebView2WebResourceRequestedEventArgsOnce.Do(func() {
		coreWebView2WebResourceRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2WebResourceRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_Create", 0), // constructor NewCoreWebView2WebResourceRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_Request", 0), // property Request
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_Response", 0), // property Response
			/* 5 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_Deferral", 0), // property Deferral
			/* 6 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_ResourceContext", 0), // property ResourceContext
			/* 7 */ imports.NewTable("TCoreWebView2WebResourceRequestedEventArgs_RequestedSourceKind", 0), // property RequestedSourceKind
		}
	})
	return coreWebView2WebResourceRequestedEventArgsImport
}
