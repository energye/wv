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

// ICoreWebView2Cookie Parent: IObject
//
//	Provides a set of properties that are used to manage an
//	ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie">See the ICoreWebView2Cookie article.</a>
type ICoreWebView2Cookie interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Cookie // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2Cookie) // property
	// Name
	//  Cookie name.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_name">See the ICoreWebView2Cookie article.</a>
	Name() string // property
	// Value
	//  Cookie value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_value">See the ICoreWebView2Cookie article.</a>
	Value() string // property
	// SetValue Set Value
	SetValue(AValue string) // property
	// Domain
	//  The domain for which the cookie is valid.
	//  The default is the host that this cookie has been received from.
	//  Note that, for instance, ".bing.com", "bing.com", and "www.bing.com" are
	//  considered different domains.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_domain">See the ICoreWebView2Cookie article.</a>
	Domain() string // property
	// Path
	//  The path for which the cookie is valid. The default is "/", which means
	//  this cookie will be sent to all pages on the Domain.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_path">See the ICoreWebView2Cookie article.</a>
	Path() string // property
	// Expires
	//  The expiration date and time for the cookie as the number of seconds since the UNIX epoch.
	//  The default is -1.0, which means cookies are session cookies by default.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</a>
	Expires() (resultDouble float64) // property
	// SetExpires Set Expires
	SetExpires(AValue float64) // property
	// ExpiresDate
	//  The expiration date and time for the cookie in TDateTime format.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</a>
	ExpiresDate() (resultDateTime float64) // property
	// SetExpiresDate Set ExpiresDate
	SetExpiresDate(AValue float64) // property
	// IsHttpOnly
	//  Whether this cookie is http-only.
	//  True if a page script or other active content cannot access this
	//  cookie. The default is false.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_ishttponly">See the ICoreWebView2Cookie article.</a>
	IsHttpOnly() bool // property
	// SetIsHttpOnly Set IsHttpOnly
	SetIsHttpOnly(AValue bool) // property
	// SameSite
	//  SameSite status of the cookie which represents the enforcement mode of the cookie.
	//  The default is COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_samesite">See the ICoreWebView2Cookie article.</a>
	SameSite() TWVCookieSameSiteKind // property
	// SetSameSite Set SameSite
	SetSameSite(AValue TWVCookieSameSiteKind) // property
	// IsSecure
	//  The security level of this cookie. True if the client is only to return
	//  the cookie in subsequent requests if those requests use HTTPS.
	//  The default is false.
	//  Note that cookie that requests COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE but
	//  is not marked Secure will be rejected.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issecure">See the ICoreWebView2Cookie article.</a>
	IsSecure() bool // property
	// SetIsSecure Set IsSecure
	SetIsSecure(AValue bool) // property
	// IsSession
	//  Whether this is a session cookie. The default is false.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issession">See the ICoreWebView2Cookie article.</a>
	IsSession() bool // property
}

// TCoreWebView2Cookie Parent: TObject
//
//	Provides a set of properties that are used to manage an
//	ICoreWebView2Cookie.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie">See the ICoreWebView2Cookie article.</a>
type TCoreWebView2Cookie struct {
	TObject
}

func NewCoreWebView2Cookie(aBaseIntf ICoreWebView2Cookie) ICoreWebView2Cookie {
	r1 := coreWebView2CookieImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Cookie(r1)
}

func (m *TCoreWebView2Cookie) Initialized() bool {
	r1 := coreWebView2CookieImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) BaseIntf() ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	coreWebView2CookieImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TCoreWebView2Cookie) SetBaseIntf(AValue ICoreWebView2Cookie) {
	coreWebView2CookieImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2Cookie) Name() string {
	r1 := coreWebView2CookieImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Value() string {
	r1 := coreWebView2CookieImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) SetValue(AValue string) {
	coreWebView2CookieImportAPI().SysCallN(12, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2Cookie) Domain() string {
	r1 := coreWebView2CookieImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Path() string {
	r1 := coreWebView2CookieImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Cookie) Expires() (resultDouble float64) {
	coreWebView2CookieImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCoreWebView2Cookie) SetExpires(AValue float64) {
	coreWebView2CookieImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2Cookie) ExpiresDate() (resultDateTime float64) {
	coreWebView2CookieImportAPI().SysCallN(4, 0, m.Instance(), uintptr(unsafePointer(&resultDateTime)), uintptr(unsafePointer(&resultDateTime)))
	return
}

func (m *TCoreWebView2Cookie) SetExpiresDate(AValue float64) {
	coreWebView2CookieImportAPI().SysCallN(4, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCoreWebView2Cookie) IsHttpOnly() bool {
	r1 := coreWebView2CookieImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) SetIsHttpOnly(AValue bool) {
	coreWebView2CookieImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Cookie) SameSite() TWVCookieSameSiteKind {
	r1 := coreWebView2CookieImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TWVCookieSameSiteKind(r1)
}

func (m *TCoreWebView2Cookie) SetSameSite(AValue TWVCookieSameSiteKind) {
	coreWebView2CookieImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Cookie) IsSecure() bool {
	r1 := coreWebView2CookieImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Cookie) SetIsSecure(AValue bool) {
	coreWebView2CookieImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Cookie) IsSession() bool {
	r1 := coreWebView2CookieImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

var (
	coreWebView2CookieImport       *imports.Imports = nil
	coreWebView2CookieImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2Cookie_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2Cookie_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2Cookie_Domain", 0),
		/*3*/ imports.NewTable("CoreWebView2Cookie_Expires", 0),
		/*4*/ imports.NewTable("CoreWebView2Cookie_ExpiresDate", 0),
		/*5*/ imports.NewTable("CoreWebView2Cookie_Initialized", 0),
		/*6*/ imports.NewTable("CoreWebView2Cookie_IsHttpOnly", 0),
		/*7*/ imports.NewTable("CoreWebView2Cookie_IsSecure", 0),
		/*8*/ imports.NewTable("CoreWebView2Cookie_IsSession", 0),
		/*9*/ imports.NewTable("CoreWebView2Cookie_Name", 0),
		/*10*/ imports.NewTable("CoreWebView2Cookie_Path", 0),
		/*11*/ imports.NewTable("CoreWebView2Cookie_SameSite", 0),
		/*12*/ imports.NewTable("CoreWebView2Cookie_Value", 0),
	}
)

func coreWebView2CookieImportAPI() *imports.Imports {
	if coreWebView2CookieImport == nil {
		coreWebView2CookieImport = NewDefaultImports()
		coreWebView2CookieImport.SetImportTable(coreWebView2CookieImportTables)
		coreWebView2CookieImportTables = nil
	}
	return coreWebView2CookieImport
}
