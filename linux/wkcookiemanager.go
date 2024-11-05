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

// IWkCookieManager Root Interface
type IWkCookieManager interface {
	IObject
	Data() WebKitCookieManager                                                   // function
	SetPersistentStorage(filename string, storage WebKitCookiePersistentStorage) // procedure
	SetAcceptPolicy(policy WebKitCookieAcceptPolicy)                             // procedure
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
	AddCookie(cookie PSoupCookie) // procedure
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
	DeleteCookie(cookie PSoupCookie) // procedure
	// DeleteCookiesForDomain
	//  删除给定域的cookie_manager的所有cookie。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_cookies_for_domain.html
	DeleteCookiesForDomain(domain string) // procedure
	// DeleteAllCookies
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_all_cookies.html
	DeleteAllCookies() // procedure
}

// TWkCookieManager Root Object
type TWkCookieManager struct {
	TObject
}

func NewWkCookieManager(aCookieManager WebKitCookieManager) IWkCookieManager {
	r1 := wkCookieManagerImportAPI().SysCallN(1, uintptr(aCookieManager))
	return AsWkCookieManager(r1)
}

// WkCookieManagerRef -> IWkCookieManager
var WkCookieManagerRef wkCookieManager

// wkCookieManager TWkCookieManager Ref
type wkCookieManager uintptr

func (m *wkCookieManager) NewDelegate(aCookieManager WebKitCookieManager, aDelegateEvent IWkCookieManagerDelegateEvent) IWkCookieManager {
	r1 := wkCookieManagerImportAPI().SysCallN(8, uintptr(aCookieManager), GetObjectUintptr(aDelegateEvent))
	return AsWkCookieManager(r1)
}

func (m *TWkCookieManager) Data() WebKitCookieManager {
	r1 := wkCookieManagerImportAPI().SysCallN(2, m.Instance())
	return WebKitCookieManager(r1)
}

func (m *TWkCookieManager) SetPersistentStorage(filename string, storage WebKitCookiePersistentStorage) {
	wkCookieManagerImportAPI().SysCallN(10, m.Instance(), PascalStr(filename), uintptr(storage))
}

func (m *TWkCookieManager) SetAcceptPolicy(policy WebKitCookieAcceptPolicy) {
	wkCookieManagerImportAPI().SysCallN(9, m.Instance(), uintptr(policy))
}

func (m *TWkCookieManager) GetAcceptPolicy() {
	wkCookieManagerImportAPI().SysCallN(6, m.Instance())
}

func (m *TWkCookieManager) AddCookie(cookie PSoupCookie) {
	wkCookieManagerImportAPI().SysCallN(0, m.Instance(), uintptr(cookie))
}

func (m *TWkCookieManager) GetCookies(uri string) {
	wkCookieManagerImportAPI().SysCallN(7, m.Instance(), PascalStr(uri))
}

func (m *TWkCookieManager) DeleteCookie(cookie PSoupCookie) {
	wkCookieManagerImportAPI().SysCallN(4, m.Instance(), uintptr(cookie))
}

func (m *TWkCookieManager) DeleteCookiesForDomain(domain string) {
	wkCookieManagerImportAPI().SysCallN(5, m.Instance(), PascalStr(domain))
}

func (m *TWkCookieManager) DeleteAllCookies() {
	wkCookieManagerImportAPI().SysCallN(3, m.Instance())
}

var (
	wkCookieManagerImport       *imports.Imports = nil
	wkCookieManagerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkCookieManager_AddCookie", 0),
		/*1*/ imports.NewTable("WkCookieManager_Create", 0),
		/*2*/ imports.NewTable("WkCookieManager_Data", 0),
		/*3*/ imports.NewTable("WkCookieManager_DeleteAllCookies", 0),
		/*4*/ imports.NewTable("WkCookieManager_DeleteCookie", 0),
		/*5*/ imports.NewTable("WkCookieManager_DeleteCookiesForDomain", 0),
		/*6*/ imports.NewTable("WkCookieManager_GetAcceptPolicy", 0),
		/*7*/ imports.NewTable("WkCookieManager_GetCookies", 0),
		/*8*/ imports.NewTable("WkCookieManager_NewDelegate", 0),
		/*9*/ imports.NewTable("WkCookieManager_SetAcceptPolicy", 0),
		/*10*/ imports.NewTable("WkCookieManager_SetPersistentStorage", 0),
	}
)

func wkCookieManagerImportAPI() *imports.Imports {
	if wkCookieManagerImport == nil {
		wkCookieManagerImport = NewDefaultImports()
		wkCookieManagerImport.SetImportTable(wkCookieManagerImportTables)
		wkCookieManagerImportTables = nil
	}
	return wkCookieManagerImport
}
