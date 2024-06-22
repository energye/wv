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

// ICoreWebView2BasicAuthenticationResponse Parent: IObject
//
//	Represents a Basic HTTP authentication response that contains a user name
//	and a password as according to RFC7617 (https://tools.ietf.org/html/rfc7617)
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse">See the ICoreWebView2BasicAuthenticationResponse article.</a>
type ICoreWebView2BasicAuthenticationResponse interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BasicAuthenticationResponse // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2BasicAuthenticationResponse) // property
	// UserName
	//  User name provided for authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_username">See the ICoreWebView2BasicAuthenticationResponse article.</a>
	UserName() string // property
	// SetUserName Set UserName
	SetUserName(AValue string) // property
	// Password
	//  Password provided for authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_password">See the ICoreWebView2BasicAuthenticationResponse article.</a>
	Password() string // property
	// SetPassword Set Password
	SetPassword(AValue string) // property
}

// TCoreWebView2BasicAuthenticationResponse Parent: TObject
//
//	Represents a Basic HTTP authentication response that contains a user name
//	and a password as according to RFC7617 (https://tools.ietf.org/html/rfc7617)
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse">See the ICoreWebView2BasicAuthenticationResponse article.</a>
type TCoreWebView2BasicAuthenticationResponse struct {
	TObject
}

func NewCoreWebView2BasicAuthenticationResponse(aBaseIntf ICoreWebView2BasicAuthenticationResponse) ICoreWebView2BasicAuthenticationResponse {
	r1 := coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2BasicAuthenticationResponse(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) Initialized() bool {
	r1 := coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) BaseIntf() ICoreWebView2BasicAuthenticationResponse {
	var resultCoreWebView2BasicAuthenticationResponse uintptr
	coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2BasicAuthenticationResponse)))
	return AsCoreWebView2BasicAuthenticationResponse(resultCoreWebView2BasicAuthenticationResponse)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetBaseIntf(AValue ICoreWebView2BasicAuthenticationResponse) {
	coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2BasicAuthenticationResponse) UserName() string {
	r1 := coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetUserName(AValue string) {
	coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2BasicAuthenticationResponse) Password() string {
	r1 := coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetPassword(AValue string) {
	coreWebView2BasicAuthenticationResponseImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

var (
	coreWebView2BasicAuthenticationResponseImport       *imports.Imports = nil
	coreWebView2BasicAuthenticationResponseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2BasicAuthenticationResponse_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2BasicAuthenticationResponse_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2BasicAuthenticationResponse_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2BasicAuthenticationResponse_Password", 0),
		/*4*/ imports.NewTable("CoreWebView2BasicAuthenticationResponse_UserName", 0),
	}
)

func coreWebView2BasicAuthenticationResponseImportAPI() *imports.Imports {
	if coreWebView2BasicAuthenticationResponseImport == nil {
		coreWebView2BasicAuthenticationResponseImport = NewDefaultImports()
		coreWebView2BasicAuthenticationResponseImport.SetImportTable(coreWebView2BasicAuthenticationResponseImportTables)
		coreWebView2BasicAuthenticationResponseImportTables = nil
	}
	return coreWebView2BasicAuthenticationResponseImport
}
