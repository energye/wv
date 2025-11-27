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

// ICoreWebView2BasicAuthenticationRequestedEventArgs Parent: lcl.IObject
type ICoreWebView2BasicAuthenticationRequestedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BasicAuthenticationRequestedEventArgs // property BaseIntf Getter
	// Uri
	//  The URI that led to the authentication challenge. For proxy authentication
	//  requests, this will be the URI of the proxy server.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_uri">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</see>
	Uri() string // property Uri Getter
	// Challenge
	//  The authentication challenge string.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_challenge">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</see>
	Challenge() string // property Challenge Getter
	// Response
	//  Response to the authentication request with credentials. This object will be populated by the app
	//  if the host would like to provide authentication credentials.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_response">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</see>
	Response() ICoreWebView2BasicAuthenticationResponse // property Response Getter
	// Cancel
	//  Cancel the authentication request. False by default.
	//  If set to true, Response will be ignored.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#get_cancel">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</see>
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this deferral to
	//  defer the decision to show the Basic Authentication dialog.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationrequestedeventargs#getdeferral">See the ICoreWebView2BasicAuthenticationRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2BasicAuthenticationRequestedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) BaseIntf() (result ICoreWebView2BasicAuthenticationRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BasicAuthenticationRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Uri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Challenge() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Response() (result ICoreWebView2BasicAuthenticationResponse) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BasicAuthenticationResponse(resultPtr)
	return
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2BasicAuthenticationRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2BasicAuthenticationRequestedEventArgs class constructor
func NewCoreWebView2BasicAuthenticationRequestedEventArgs(args ICoreWebView2BasicAuthenticationRequestedEventArgs) ICoreWebView2BasicAuthenticationRequestedEventArgs {
	r := coreWebView2BasicAuthenticationRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2BasicAuthenticationRequestedEventArgs(r)
}

var (
	coreWebView2BasicAuthenticationRequestedEventArgsOnce   base.Once
	coreWebView2BasicAuthenticationRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2BasicAuthenticationRequestedEventArgsAPI() *imports.Imports {
	coreWebView2BasicAuthenticationRequestedEventArgsOnce.Do(func() {
		coreWebView2BasicAuthenticationRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2BasicAuthenticationRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Create", 0), // constructor NewCoreWebView2BasicAuthenticationRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Uri", 0), // property Uri
			/* 4 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Challenge", 0), // property Challenge
			/* 5 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Response", 0), // property Response
			/* 6 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Cancel", 0), // property Cancel
			/* 7 */ imports.NewTable("TCoreWebView2BasicAuthenticationRequestedEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2BasicAuthenticationRequestedEventArgsImport
}
