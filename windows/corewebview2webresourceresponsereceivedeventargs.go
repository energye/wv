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

// ICoreWebView2WebResourceResponseReceivedEventArgs Parent: lcl.IObject
type ICoreWebView2WebResourceResponseReceivedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponseReceivedEventArgs // property BaseIntf Getter
	// Request
	//  The request object for the web resource, as committed. This includes
	//  headers added by the network stack that were not be included during the
	//  associated WebResourceRequested event, such as Authentication headers.
	//  Modifications to this object have no effect on how the request is
	//  processed as it has already been sent.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs#get_request">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</see>
	Request() ICoreWebView2WebResourceRequest // property Request Getter
	// Response
	//  View of the response object received for the web resource.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs#get_response">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</see>
	Response() ICoreWebView2WebResourceResponseView // property Response Getter
}

type TCoreWebView2WebResourceResponseReceivedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WebResourceResponseReceivedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) BaseIntf() (result ICoreWebView2WebResourceResponseReceivedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseReceivedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceResponseReceivedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Request() (result ICoreWebView2WebResourceRequest) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseReceivedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceRequestOwn(resultPtr)
	return
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Response() (result ICoreWebView2WebResourceResponseView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WebResourceResponseReceivedEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WebResourceResponseView(resultPtr)
	return
}

// NewCoreWebView2WebResourceResponseReceivedEventArgs class constructor
func NewCoreWebView2WebResourceResponseReceivedEventArgs(args ICoreWebView2WebResourceResponseReceivedEventArgs) ICoreWebView2WebResourceResponseReceivedEventArgs {
	r := coreWebView2WebResourceResponseReceivedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2WebResourceResponseReceivedEventArgs(r)
}

var (
	coreWebView2WebResourceResponseReceivedEventArgsOnce   base.Once
	coreWebView2WebResourceResponseReceivedEventArgsImport *imports.Imports = nil
)

func coreWebView2WebResourceResponseReceivedEventArgsAPI() *imports.Imports {
	coreWebView2WebResourceResponseReceivedEventArgsOnce.Do(func() {
		coreWebView2WebResourceResponseReceivedEventArgsImport = api.NewDefaultImports()
		coreWebView2WebResourceResponseReceivedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceResponseReceivedEventArgs_Create", 0), // constructor NewCoreWebView2WebResourceResponseReceivedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceResponseReceivedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceResponseReceivedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2WebResourceResponseReceivedEventArgs_Request", 0), // property Request
			/* 4 */ imports.NewTable("TCoreWebView2WebResourceResponseReceivedEventArgs_Response", 0), // property Response
		}
	})
	return coreWebView2WebResourceResponseReceivedEventArgsImport
}
