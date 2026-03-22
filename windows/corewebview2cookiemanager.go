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
)

// ICoreWebView2CookieManager Parent: IObject
type ICoreWebView2CookieManager interface {
	IObject
	// CreateCookie
	//  Create a cookie object with a specified name, value, domain, and path.
	//  One can set other optional properties after cookie creation.
	//  This only creates a cookie object and it is not added to the cookie
	//  manager until you call AddOrUpdateCookie.
	//  Leading or trailing whitespace(s), empty string, and special characters
	//  are not allowed for name.
	//  See ICoreWebView2Cookie for more details.
	CreateCookie(name string, value string, domain string, path string) ICoreWebView2Cookie // function
	// CopyCookie
	//  Creates a cookie whose params matches those of the specified cookie.
	CopyCookie(cookie ICoreWebView2Cookie) ICoreWebView2Cookie // function
	// GetCookies
	//  Gets a list of cookies matching the specific URI.
	//  If uri is empty string or null, all cookies under the same profile are
	//  returned.
	//  You can modify the cookie objects by calling
	//  ICoreWebView2CookieManager::AddOrUpdateCookie, and the changes
	//  will be applied to the webview.
	GetCookies(uRI string, handler ICoreWebView2GetCookiesCompletedHandler) bool // function
	// AddOrUpdateCookie
	//  Adds or updates a cookie with the given cookie data; may overwrite
	//  cookies with matching name, domain, and path if they exist.
	//  This method will fail if the domain of the given cookie is not specified.
	AddOrUpdateCookie(cookie ICoreWebView2Cookie) bool // function
	// DeleteCookie
	//  Deletes a cookie whose name and domain/path pair
	//  match those of the specified cookie.
	DeleteCookie(cookie ICoreWebView2Cookie) bool // function
	// DeleteCookies
	//  Deletes cookies with matching name and uri.
	//  Cookie name is required.
	//  All cookies with the given name where domain
	//  and path match provided URI are deleted.
	DeleteCookies(name string, uRI string) bool // function
	// DeleteCookiesWithDomainAndPath
	//  Deletes cookies with matching name and domain/path pair.
	//  Cookie name is required.
	//  If domain is specified, deletes only cookies with the exact domain.
	//  If path is specified, deletes only cookies with the exact path.
	DeleteCookiesWithDomainAndPath(name string, domain string, path string) bool // function
	// DeleteAllCookies
	//  Deletes all cookies under the same profile.
	//  This could affect other WebViews under the same user profile.
	DeleteAllCookies() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CookieManager // property BaseIntf Getter
}

type TCoreWebView2CookieManager struct {
	TObject
}

func (m *TCoreWebView2CookieManager) CreateCookie(name string, value string, domain string, path string) (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieManagerAPI().SysCallN(1, m.Instance(), api.PasStr(name), api.PasStr(value), api.PasStr(domain), api.PasStr(path), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

func (m *TCoreWebView2CookieManager) CopyCookie(cookie ICoreWebView2Cookie) (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieManagerAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(cookie), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

func (m *TCoreWebView2CookieManager) GetCookies(uRI string, handler ICoreWebView2GetCookiesCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(3, m.Instance(), api.PasStr(uRI), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) AddOrUpdateCookie(cookie ICoreWebView2Cookie) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(cookie))
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) DeleteCookie(cookie ICoreWebView2Cookie) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(cookie))
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) DeleteCookies(name string, uRI string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(6, m.Instance(), api.PasStr(name), api.PasStr(uRI))
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) DeleteCookiesWithDomainAndPath(name string, domain string, path string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(7, m.Instance(), api.PasStr(name), api.PasStr(domain), api.PasStr(path))
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) DeleteAllCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CookieManagerAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2CookieManager) BaseIntf() (result ICoreWebView2CookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CookieManagerAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieManager(resultPtr)
	return
}

// NewCoreWebView2CookieManager class constructor
func NewCoreWebView2CookieManager(baseIntf ICoreWebView2CookieManager) ICoreWebView2CookieManager {
	r := coreWebView2CookieManagerAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2CookieManager(r)
}

var (
	coreWebView2CookieManagerOnce   base.Once
	coreWebView2CookieManagerImport *imports.Imports = nil
)

func coreWebView2CookieManagerAPI() *imports.Imports {
	coreWebView2CookieManagerOnce.Do(func() {
		coreWebView2CookieManagerImport = api.NewDefaultImports()
		coreWebView2CookieManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2CookieManager_Create", 0), // constructor NewCoreWebView2CookieManager
			/* 1 */ imports.NewTable("TCoreWebView2CookieManager_CreateCookie", 0), // function CreateCookie
			/* 2 */ imports.NewTable("TCoreWebView2CookieManager_CopyCookie", 0), // function CopyCookie
			/* 3 */ imports.NewTable("TCoreWebView2CookieManager_GetCookies", 0), // function GetCookies
			/* 4 */ imports.NewTable("TCoreWebView2CookieManager_AddOrUpdateCookie", 0), // function AddOrUpdateCookie
			/* 5 */ imports.NewTable("TCoreWebView2CookieManager_DeleteCookie", 0), // function DeleteCookie
			/* 6 */ imports.NewTable("TCoreWebView2CookieManager_DeleteCookies", 0), // function DeleteCookies
			/* 7 */ imports.NewTable("TCoreWebView2CookieManager_DeleteCookiesWithDomainAndPath", 0), // function DeleteCookiesWithDomainAndPath
			/* 8 */ imports.NewTable("TCoreWebView2CookieManager_DeleteAllCookies", 0), // function DeleteAllCookies
			/* 9 */ imports.NewTable("TCoreWebView2CookieManager_Initialized", 0), // property Initialized
			/* 10 */ imports.NewTable("TCoreWebView2CookieManager_BaseIntf", 0), // property BaseIntf
		}
	})
	return coreWebView2CookieManagerImport
}
