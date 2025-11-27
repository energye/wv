//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkCookieManager Parent: lcl.IObject
type IWkCookieManager interface {
	lcl.IObject
	Data() wvTypes.WebKitCookieManager                                                   // function
	SetPersistentStorage(filename string, storage wvTypes.WebKitCookiePersistentStorage) // procedure
	SetAcceptPolicy(policy wvTypes.WebKitCookieAcceptPolicy)                             // procedure
	// GetAcceptPolicy
	//  异步获取cookie_manager的cookie接受策略。
	//  注意，当策略被设置为WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY并且ITP被启用时，这将返回WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS。参见webkit_website_data_manager_set_itp_enabled()。
	//  当操作完成时，回调将被调用。然后可以调用webkit_cookie_manager_get_accept_policy_finish()来获取操作的结果。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.get_accept_policy.html
	GetAcceptPolicy() // procedure
	// AddCookie
	//  异步地将SoupCookie添加到底层存储。
	//  当操作完成时，回调将被调用。然后可以调用webkit_cookie_manager_add_cookie_finish()来获取操作的结果。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.add_cookie.html
	AddCookie(cookie wvTypes.PSoupCookie) // procedure
	// GetCookies
	//  从cookie_manager异步获取SoupCookie的列表。
	//  从与uri关联的cookie_manager异步获取SoupCookie的列表，该列表必须是HTTP或HTTPS URL。
	//  当操作完成时，回调将被调用。然后可以调用webkit_cookie_manager_get_cookies_finish()来获取操作的结果。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.get_cookies.html
	GetCookies(uri string) // procedure
	// DeleteCookie
	//  从当前会话中异步删除SoupCookie。
	//  当操作完成时，回调将被调用。然后可以调用webkit_cookie_manager_delete_cookie_finish()来获取操作的结果。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_cookie.html
	DeleteCookie(cookie wvTypes.PSoupCookie) // procedure
	// DeleteCookiesForDomain
	//  删除给定域的cookie_manager的所有cookie。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_cookies_for_domain.html
	DeleteCookiesForDomain(domain string) // procedure
	// DeleteAllCookies
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_all_cookies.html
	DeleteAllCookies() // procedure
}

type TWkCookieManager struct {
	lcl.TObject
}

func (m *TWkCookieManager) Data() wvTypes.WebKitCookieManager {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieManagerAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitCookieManager(r)
}

func (m *TWkCookieManager) SetPersistentStorage(filename string, storage wvTypes.WebKitCookiePersistentStorage) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(3, m.Instance(), api.PasStr(filename), uintptr(storage))
}

func (m *TWkCookieManager) SetAcceptPolicy(policy wvTypes.WebKitCookieAcceptPolicy) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(4, m.Instance(), uintptr(policy))
}

func (m *TWkCookieManager) GetAcceptPolicy() {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(5, m.Instance())
}

func (m *TWkCookieManager) AddCookie(cookie wvTypes.PSoupCookie) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(6, m.Instance(), uintptr(cookie))
}

func (m *TWkCookieManager) GetCookies(uri string) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(7, m.Instance(), api.PasStr(uri))
}

func (m *TWkCookieManager) DeleteCookie(cookie wvTypes.PSoupCookie) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(8, m.Instance(), uintptr(cookie))
}

func (m *TWkCookieManager) DeleteCookiesForDomain(domain string) {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(9, m.Instance(), api.PasStr(domain))
}

func (m *TWkCookieManager) DeleteAllCookies() {
	if !m.IsValid() {
		return
	}
	wkCookieManagerAPI().SysCallN(10, m.Instance())
}

// CookieManager  is static instance
var CookieManager _CookieManagerClass

// _CookieManagerClass is class type defined by TWkCookieManager
type _CookieManagerClass uintptr

func (_CookieManagerClass) NewDelegate(cookieManager wvTypes.WebKitCookieManager, delegateEvent IWkWebview) IWkCookieManager {
	r := wkCookieManagerAPI().SysCallN(2, uintptr(cookieManager), base.GetObjectUintptr(delegateEvent))
	return AsWkCookieManager(r)
}

// NewCookieManager class constructor
func NewCookieManager(cookieManager wvTypes.WebKitCookieManager) IWkCookieManager {
	r := wkCookieManagerAPI().SysCallN(0, uintptr(cookieManager))
	return AsWkCookieManager(r)
}

var (
	wkCookieManagerOnce   base.Once
	wkCookieManagerImport *imports.Imports = nil
)

func wkCookieManagerAPI() *imports.Imports {
	wkCookieManagerOnce.Do(func() {
		wkCookieManagerImport = api.NewDefaultImports()
		wkCookieManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkCookieManager_Create", 0), // constructor NewCookieManager
			/* 1 */ imports.NewTable("TWkCookieManager_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkCookieManager_NewDelegate", 0), // static function NewDelegate
			/* 3 */ imports.NewTable("TWkCookieManager_SetPersistentStorage", 0), // procedure SetPersistentStorage
			/* 4 */ imports.NewTable("TWkCookieManager_SetAcceptPolicy", 0), // procedure SetAcceptPolicy
			/* 5 */ imports.NewTable("TWkCookieManager_GetAcceptPolicy", 0), // procedure GetAcceptPolicy
			/* 6 */ imports.NewTable("TWkCookieManager_AddCookie", 0), // procedure AddCookie
			/* 7 */ imports.NewTable("TWkCookieManager_GetCookies", 0), // procedure GetCookies
			/* 8 */ imports.NewTable("TWkCookieManager_DeleteCookie", 0), // procedure DeleteCookie
			/* 9 */ imports.NewTable("TWkCookieManager_DeleteCookiesForDomain", 0), // procedure DeleteCookiesForDomain
			/* 10 */ imports.NewTable("TWkCookieManager_DeleteAllCookies", 0), // procedure DeleteAllCookies
		}
	})
	return wkCookieManagerImport
}
