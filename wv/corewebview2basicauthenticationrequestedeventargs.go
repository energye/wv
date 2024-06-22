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

// ICoreWebView2BasicAuthenticationRequestedEventArgs Parent: IObject
//
//	Event args for the BasicAuthenticationRequested event. Will contain the
//	request that led to the HTTP authorization challenge, the challenge
//	and allows the host to provide authentication response or cancel the request.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
type ICoreWebView2BasicAuthenticationRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BasicAuthenticationRequestedEventArgs // property
	// Uri
	//  The URI that led to the authentication challenge. For proxy authentication
	//  requests, this will be the URI of the proxy server.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_uri">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
	Uri() string // property
	// Challenge
	//  The authentication challenge string.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_challenge">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
	Challenge() string // property
	// Response
	//  Response to the authentication request with credentials. This object will be populated by the app
	//  if the host would like to provide authentication credentials.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_response">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
	Response() ICoreWebView2BasicAuthenticationResponse // property
	// Cancel
	//  Cancel the authentication request. False by default.
	//  If set to true, Response will be ignored.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_cancel">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this deferral to
	//  defer the decision to show the Basic Authentication dialog.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#getdeferral">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2BasicAuthenticationRequestedEventArgs Parent: TObject
//
//	Event args for the BasicAuthenticationRequested event. Will contain the
//	request that led to the HTTP authorization challenge, the challenge
//	and allows the host to provide authentication response or cancel the request.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</a>
type TCoreWebView2BasicAuthenticationRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2BasicAuthenticationRequestedEventArgs(aArgs ICoreWebView2BasicAuthenticationRequestedEventArgs) ICoreWebView2BasicAuthenticationRequestedEventArgs {
	r1 := coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(3, GetObjectUintptr(aArgs))
	return AsCoreWebView2BasicAuthenticationRequestedEventArgs(r1)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Initialized() bool {
	r1 := coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) BaseIntf() ICoreWebView2BasicAuthenticationRequestedEventArgs {
	var resultCoreWebView2BasicAuthenticationRequestedEventArgs uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2BasicAuthenticationRequestedEventArgs)))
	return AsCoreWebView2BasicAuthenticationRequestedEventArgs(resultCoreWebView2BasicAuthenticationRequestedEventArgs)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Uri() string {
	r1 := coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Challenge() string {
	r1 := coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Response() ICoreWebView2BasicAuthenticationResponse {
	var resultCoreWebView2BasicAuthenticationResponse uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2BasicAuthenticationResponse)))
	return AsCoreWebView2BasicAuthenticationResponse(resultCoreWebView2BasicAuthenticationResponse)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Cancel() bool {
	r1 := coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) SetCancel(AValue bool) {
	coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

var (
	coreWebView2BasicAuthenticationRequestedEventArgsImport       *imports.Imports = nil
	coreWebView2BasicAuthenticationRequestedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Cancel", 0),
		/*2*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Challenge", 0),
		/*3*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Create", 0),
		/*4*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Deferral", 0),
		/*5*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Initialized", 0),
		/*6*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Response", 0),
		/*7*/ imports.NewTable("CoreWebView2BasicAuthenticationRequestedEventArgs_Uri", 0),
	}
)

func coreWebView2BasicAuthenticationRequestedEventArgsImportAPI() *imports.Imports {
	if coreWebView2BasicAuthenticationRequestedEventArgsImport == nil {
		coreWebView2BasicAuthenticationRequestedEventArgsImport = NewDefaultImports()
		coreWebView2BasicAuthenticationRequestedEventArgsImport.SetImportTable(coreWebView2BasicAuthenticationRequestedEventArgsImportTables)
		coreWebView2BasicAuthenticationRequestedEventArgsImportTables = nil
	}
	return coreWebView2BasicAuthenticationRequestedEventArgsImport
}
