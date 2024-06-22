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

// ICoreWebView2WebResourceResponseReceivedEventArgs Parent: IObject
//
//	Event args for the WebResourceResponseReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</a>
type ICoreWebView2WebResourceResponseReceivedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WebResourceResponseReceivedEventArgs // property
	// Request
	//  The request object for the web resource, as committed. This includes
	//  headers added by the network stack that were not be included during the
	//  associated WebResourceRequested event, such as Authentication headers.
	//  Modifications to this object have no effect on how the request is
	//  processed as it has already been sent.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs#get_request">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</a>
	Request() ICoreWebView2WebResourceRequestRef // property
	// Response
	//  View of the response object received for the web resource.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs#get_response">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</a>
	Response() ICoreWebView2WebResourceResponseView // property
}

// TCoreWebView2WebResourceResponseReceivedEventArgs Parent: TObject
//
//	Event args for the WebResourceResponseReceived event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponsereceivedeventargs">See the ICoreWebView2WebResourceResponseReceivedEventArgs article.</a>
type TCoreWebView2WebResourceResponseReceivedEventArgs struct {
	TObject
}

func NewCoreWebView2WebResourceResponseReceivedEventArgs(aArgs ICoreWebView2WebResourceResponseReceivedEventArgs) ICoreWebView2WebResourceResponseReceivedEventArgs {
	r1 := coreWebView2WebResourceResponseReceivedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2WebResourceResponseReceivedEventArgs(r1)
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Initialized() bool {
	r1 := coreWebView2WebResourceResponseReceivedEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) BaseIntf() ICoreWebView2WebResourceResponseReceivedEventArgs {
	var resultCoreWebView2WebResourceResponseReceivedEventArgs uintptr
	coreWebView2WebResourceResponseReceivedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceResponseReceivedEventArgs)))
	return AsCoreWebView2WebResourceResponseReceivedEventArgs(resultCoreWebView2WebResourceResponseReceivedEventArgs)
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Request() ICoreWebView2WebResourceRequestRef {
	var resultCoreWebView2WebResourceRequest uintptr
	coreWebView2WebResourceResponseReceivedEventArgsImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceRequest)))
	return AsCoreWebView2WebResourceRequest(resultCoreWebView2WebResourceRequest)
}

func (m *TCoreWebView2WebResourceResponseReceivedEventArgs) Response() ICoreWebView2WebResourceResponseView {
	var resultCoreWebView2WebResourceResponseView uintptr
	coreWebView2WebResourceResponseReceivedEventArgsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WebResourceResponseView)))
	return AsCoreWebView2WebResourceResponseView(resultCoreWebView2WebResourceResponseView)
}

var (
	coreWebView2WebResourceResponseReceivedEventArgsImport       *imports.Imports = nil
	coreWebView2WebResourceResponseReceivedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2WebResourceResponseReceivedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2WebResourceResponseReceivedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2WebResourceResponseReceivedEventArgs_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2WebResourceResponseReceivedEventArgs_Request", 0),
		/*4*/ imports.NewTable("CoreWebView2WebResourceResponseReceivedEventArgs_Response", 0),
	}
)

func coreWebView2WebResourceResponseReceivedEventArgsImportAPI() *imports.Imports {
	if coreWebView2WebResourceResponseReceivedEventArgsImport == nil {
		coreWebView2WebResourceResponseReceivedEventArgsImport = NewDefaultImports()
		coreWebView2WebResourceResponseReceivedEventArgsImport.SetImportTable(coreWebView2WebResourceResponseReceivedEventArgsImportTables)
		coreWebView2WebResourceResponseReceivedEventArgsImportTables = nil
	}
	return coreWebView2WebResourceResponseReceivedEventArgsImport
}
