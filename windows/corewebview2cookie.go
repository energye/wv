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
	"github.com/energye/lcl/types"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2Cookie Parent: IObject
type ICoreWebView2Cookie interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Cookie         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2Cookie) // property BaseIntf Setter
	// Name
	//  Cookie name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_name">See the ICoreWebView2Cookie article.</see>
	Name() string // property Name Getter
	// Value
	//  Cookie value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_value">See the ICoreWebView2Cookie article.</see>
	Value() string         // property Value Getter
	SetValue(value string) // property Value Setter
	// Domain
	//  The domain for which the cookie is valid.
	//  The default is the host that this cookie has been received from.
	//  Note that, for instance, ".bing.com", "bing.com", and "www.bing.com" are
	//  considered different domains.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_domain">See the ICoreWebView2Cookie article.</see>
	Domain() string // property Domain Getter
	// Path
	//  The path for which the cookie is valid. The default is "/", which means
	//  this cookie will be sent to all pages on the Domain.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_path">See the ICoreWebView2Cookie article.</see>
	Path() string // property Path Getter
	// Expires
	//  The expiration date and time for the cookie as the number of seconds since the UNIX epoch.
	//  The default is -1.0, which means cookies are session cookies by default.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</see>
	Expires() float64         // property Expires Getter
	SetExpires(value float64) // property Expires Setter
	// ExpiresDate
	//  The expiration date and time for the cookie in TDateTime format.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_expires">See the ICoreWebView2Cookie article.</see>
	ExpiresDate() types.TDateTime         // property ExpiresDate Getter
	SetExpiresDate(value types.TDateTime) // property ExpiresDate Setter
	// IsHttpOnly
	//  Whether this cookie is http-only.
	//  True if a page script or other active content cannot access this
	//  cookie. The default is false.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_ishttponly">See the ICoreWebView2Cookie article.</see>
	IsHttpOnly() bool         // property IsHttpOnly Getter
	SetIsHttpOnly(value bool) // property IsHttpOnly Setter
	// SameSite
	//  SameSite status of the cookie which represents the enforcement mode of the cookie.
	//  The default is COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_samesite">See the ICoreWebView2Cookie article.</see>
	SameSite() wvTypes.TWVCookieSameSiteKind         // property SameSite Getter
	SetSameSite(value wvTypes.TWVCookieSameSiteKind) // property SameSite Setter
	// IsSecure
	//  The security level of this cookie. True if the client is only to return
	//  the cookie in subsequent requests if those requests use HTTPS.
	//  The default is false.
	//  Note that cookie that requests COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE but
	//  is not marked Secure will be rejected.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issecure">See the ICoreWebView2Cookie article.</see>
	IsSecure() bool         // property IsSecure Getter
	SetIsSecure(value bool) // property IsSecure Setter
	// IsSession
	//  Whether this is a session cookie. The default is false.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookie#get_issession">See the ICoreWebView2Cookie article.</see>
	IsSession() bool // property IsSession Getter
}

type TCoreWebView2Cookie struct {
	TObject
}

func (m *TCoreWebView2Cookie) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Cookie) BaseIntf() (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

func (m *TCoreWebView2Cookie) SetBaseIntf(value ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2Cookie) Name() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2CookieAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Cookie) Value() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2CookieAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Cookie) SetValue(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2Cookie) Domain() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2CookieAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Cookie) Path() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2CookieAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Cookie) Expires() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Cookie) SetExpires(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2Cookie) ExpiresDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Cookie) SetExpiresDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2Cookie) IsHttpOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Cookie) SetIsHttpOnly(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Cookie) SameSite() wvTypes.TWVCookieSameSiteKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CookieAPI().SysCallN(10, 0, m.Instance())
	return wvTypes.TWVCookieSameSiteKind(r)
}

func (m *TCoreWebView2Cookie) SetSameSite(value wvTypes.TWVCookieSameSiteKind) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Cookie) IsSecure() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Cookie) SetIsSecure(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2CookieAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Cookie) IsSession() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2Cookie class constructor
func NewCoreWebView2Cookie(baseIntf ICoreWebView2Cookie) ICoreWebView2Cookie {
	r := coreWebView2CookieAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Cookie(r)
}

var (
	coreWebView2CookieOnce   base.Once
	coreWebView2CookieImport *imports.Imports = nil
)

func coreWebView2CookieAPI() *imports.Imports {
	coreWebView2CookieOnce.Do(func() {
		coreWebView2CookieImport = api.NewDefaultImports()
		coreWebView2CookieImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Cookie_Create", 0), // constructor NewCoreWebView2Cookie
			/* 1 */ imports.NewTable("TCoreWebView2Cookie_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2Cookie_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2Cookie_Name", 0), // property Name
			/* 4 */ imports.NewTable("TCoreWebView2Cookie_Value", 0), // property Value
			/* 5 */ imports.NewTable("TCoreWebView2Cookie_Domain", 0), // property Domain
			/* 6 */ imports.NewTable("TCoreWebView2Cookie_Path", 0), // property Path
			/* 7 */ imports.NewTable("TCoreWebView2Cookie_Expires", 0), // property Expires
			/* 8 */ imports.NewTable("TCoreWebView2Cookie_ExpiresDate", 0), // property ExpiresDate
			/* 9 */ imports.NewTable("TCoreWebView2Cookie_IsHttpOnly", 0), // property IsHttpOnly
			/* 10 */ imports.NewTable("TCoreWebView2Cookie_SameSite", 0), // property SameSite
			/* 11 */ imports.NewTable("TCoreWebView2Cookie_IsSecure", 0), // property IsSecure
			/* 12 */ imports.NewTable("TCoreWebView2Cookie_IsSession", 0), // property IsSession
		}
	})
	return coreWebView2CookieImport
}
