//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

type IWkCookieManagerDelegateEvent interface {
	// SetOnGetAcceptPolicyFinish
	//  完成由webkit_cookie_manager_get_accept_policy()启动的异步操作。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.get_accept_policy_finish.html
	SetOnGetAcceptPolicyFinish(fn TWkGetAcceptPolicyFinishEvent) // property event
	// SetOnAddCookieFinish
	//  完成由webkit_cookie_manager_add_cookie()开始的异步操作。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.add_cookie_finish.html
	SetOnAddCookieFinish(fn TWkAddCookieFinishEvent) // property event
	// SetOnGetCookiesFinish
	//  完成由webkit_cookie_manager_get_cookies()开始的异步操作。
	//  返回值是SoupCookie实例的GSList，应该使用g_list_free_full()和soup_cookie_free()释放。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.get_cookies_finish.html
	SetOnGetCookiesFinish(fn TWkGetCookiesFinishEvent) // property event
	// SetOnDeleteCookieFinish
	//  完成由webkit_cookie_manager_delete_cookie()开始的异步操作。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_cookie_finish.html
	SetOnDeleteCookieFinish(fn TWkDeleteCookieFinishEvent) // property event
}

type IWkSchemeRequestDelegateEvent interface {
	SetOnURISchemeRequest(fn TWkURISchemeRequestEvent)
}

type TWkLoadChangeEvent func(sender IObject, wkLoadEvent WebKitLoadEvent)
type TWkExecuteScriptFinishedEvent func(sender IObject, jsValue IWkJSValue)
type TWkLoadFailedEvent func(sender IObject, wkLoadEvent WebKitLoadEvent, failingUri string, error_ string) bool
type TWkURISchemeRequestEvent func(sender IObject, wkURISchemeRequest WebKitURISchemeRequest)
type TWkProcessMessageEvent func(sender IObject, jsValue IWkJSValue, processId TWkProcessId)
type TWkMousePressEvent func(sender IObject, event TWkButtonEvent) bool
type TWkMouseReleaseEvent func(sender IObject, event TWkButtonEvent) bool
type TWkGetAcceptPolicyFinishEvent func(sender IObject, policy WebKitCookieAcceptPolicy, error_ string)
type TWkAddCookieFinishEvent func(sender IObject, result bool, error_ string)
type TWkGetCookiesFinishEvent func(sender IObject, wkCookieList PList, error_ string)
type TWkDeleteCookieFinishEvent func(sender IObject, result bool, error_ string)
type TWkDecidePolicyEvent func(sender IObject, wkDecision WebKitPolicyDecision, type_ WebKitPolicyDecisionType) bool
type TWkWebProcessTerminatedEvent func(sender IObject, wkReason WebKitWebProcessTerminationReason)
type TWkContextMenuEvent func(sender IObject, wkContextMenu WebKitContextMenu, wkDefaultAction PWkAction) bool
type TWkContextMenuCommandEvent func(sender IObject, menuID int32)
type TWkContextMenuDismissedEvent func(sender IObject)
