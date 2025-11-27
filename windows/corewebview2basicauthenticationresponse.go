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

// ICoreWebView2BasicAuthenticationResponse Parent: lcl.IObject
type ICoreWebView2BasicAuthenticationResponse interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BasicAuthenticationResponse         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2BasicAuthenticationResponse) // property BaseIntf Setter
	// UserName
	//  User name provided for authentication.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_username">See the ICoreWebView2BasicAuthenticationResponse article.</see>
	UserName() string         // property UserName Getter
	SetUserName(value string) // property UserName Setter
	// Password
	//  Password provided for authentication.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2basicauthenticationresponse#get_password">See the ICoreWebView2BasicAuthenticationResponse article.</see>
	Password() string         // property Password Getter
	SetPassword(value string) // property Password Setter
}

type TCoreWebView2BasicAuthenticationResponse struct {
	lcl.TObject
}

func (m *TCoreWebView2BasicAuthenticationResponse) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2BasicAuthenticationResponseAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2BasicAuthenticationResponse) BaseIntf() (result ICoreWebView2BasicAuthenticationResponse) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2BasicAuthenticationResponseAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2BasicAuthenticationResponse(resultPtr)
	return
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetBaseIntf(value ICoreWebView2BasicAuthenticationResponse) {
	if !m.IsValid() {
		return
	}
	coreWebView2BasicAuthenticationResponseAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2BasicAuthenticationResponse) UserName() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2BasicAuthenticationResponseAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetUserName(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2BasicAuthenticationResponseAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2BasicAuthenticationResponse) Password() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2BasicAuthenticationResponseAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2BasicAuthenticationResponse) SetPassword(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2BasicAuthenticationResponseAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

// NewCoreWebView2BasicAuthenticationResponse class constructor
func NewCoreWebView2BasicAuthenticationResponse(baseIntf ICoreWebView2BasicAuthenticationResponse) ICoreWebView2BasicAuthenticationResponse {
	r := coreWebView2BasicAuthenticationResponseAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2BasicAuthenticationResponse(r)
}

var (
	coreWebView2BasicAuthenticationResponseOnce   base.Once
	coreWebView2BasicAuthenticationResponseImport *imports.Imports = nil
)

func coreWebView2BasicAuthenticationResponseAPI() *imports.Imports {
	coreWebView2BasicAuthenticationResponseOnce.Do(func() {
		coreWebView2BasicAuthenticationResponseImport = api.NewDefaultImports()
		coreWebView2BasicAuthenticationResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BasicAuthenticationResponse_Create", 0), // constructor NewCoreWebView2BasicAuthenticationResponse
			/* 1 */ imports.NewTable("TCoreWebView2BasicAuthenticationResponse_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2BasicAuthenticationResponse_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2BasicAuthenticationResponse_UserName", 0), // property UserName
			/* 4 */ imports.NewTable("TCoreWebView2BasicAuthenticationResponse_Password", 0), // property Password
		}
	})
	return coreWebView2BasicAuthenticationResponseImport
}
