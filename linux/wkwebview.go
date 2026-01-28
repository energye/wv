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

// IWkWebview Parent: lcl.IComponent
type IWkWebview interface {
	lcl.IComponent
	CanGoBack() bool                       // function
	CanGoForward() bool                    // function
	IsLoading() bool                       // function
	GetTitle() string                      // function
	GetURI() string                        // function
	CookieManager() IWkCookieManager       // function
	GetSettings() IWkSettings              // function
	AsCookieManagerDelegate() IWkWebview   // function
	AsSchemeRequestDelegate() IWkWebview   // function
	CreateBrowser()                        // procedure
	GoBack()                               // procedure
	GoForward()                            // procedure
	Reload()                               // procedure
	Stop()                                 // procedure
	StartDrag(formWindow lcl.IWinControl)  // procedure
	TerminateWebProcess()                  // procedure
	TryClose()                             // procedure
	ExecuteScript(script string, id int32) // procedure
	LoadHTML(hTML string)                  // procedure
	LoadURL(uRL string)                    // procedure
	EnabledDevtools(enabled bool)          // procedure
	ShowDevtools()                         // procedure
	RegisterScriptCode(scriptCode string)  // procedure
	// RegisterScriptMessageHandler
	//  注册JS消息处理器
	//  window.webkit.messageHandlers.<name>.postMessage()
	RegisterScriptMessageHandler(name string)                    // procedure
	SetSettings(setting IWkSettings)                             // procedure
	FreeWebview()                                                // procedure
	WebView() wvTypes.WebKitWebView                              // property WebView Getter
	SetOnLoadChange(fn TWkLoadChangeEvent)                       // property event
	SetOnExecuteScriptFinished(fn TWkExecuteScriptFinishedEvent) // property event
	SetOnLoadFailed(fn TWkLoadFailedEvent)                       // property event
	SetOnURISchemeRequest(fn TWkURISchemeRequestEvent)           // property event
	SetOnProcessMessage(fn TWkProcessMessageEvent)               // property event
	SetOnMouseMove(fn TWkMouseEvent)                             // property event
	SetOnMousePress(fn TWkMouseEvent)                            // property event
	SetOnMouseRelease(fn TWkMouseEvent)                          // property event
	SetOnGetAcceptPolicyFinish(fn TWkGetAcceptPolicyFinishEvent) // property event
	SetOnAddCookieFinish(fn TWkAddCookieFinishEvent)             // property event
	SetOnGetCookiesFinish(fn TWkGetCookiesFinishEvent)           // property event
	SetOnDeleteCookieFinish(fn TWkDeleteCookieFinishEvent)       // property event
	SetOnDecidePolicy(fn TWkDecidePolicyEvent)                   // property event
	SetOnWebProcessTerminated(fn TWkWebProcessTerminatedEvent)   // property event
	SetOnContextMenu(fn TWkContextMenuEvent)                     // property event
	SetOnContextMenuCommand(fn TWkContextMenuCommandEvent)       // property event
	SetOnContextMenuDismissed(fn TWkContextMenuDismissedEvent)   // property event
	AsIntfWkCookieManagerDelegateEvent() uintptr
	AsIntfWkSchemeRequestDelegateEvent() uintptr
}

type TWkWebview struct {
	lcl.TComponent
}

func (m *TWkWebview) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) GetTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := wkWebviewAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TWkWebview) GetURI() string {
	if !m.IsValid() {
		return ""
	}
	r := wkWebviewAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TWkWebview) CookieManager() IWkCookieManager {
	if !m.IsValid() {
		return nil
	}
	r := wkWebviewAPI().SysCallN(6, m.Instance())
	return AsWkCookieManager(r)
}

func (m *TWkWebview) GetSettings() IWkSettings {
	if !m.IsValid() {
		return nil
	}
	r := wkWebviewAPI().SysCallN(7, m.Instance())
	return AsWkSettings(r)
}

func (m *TWkWebview) AsCookieManagerDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) AsSchemeRequestDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) CreateBrowser() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(10, m.Instance())
}

func (m *TWkWebview) GoBack() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(11, m.Instance())
}

func (m *TWkWebview) GoForward() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(12, m.Instance())
}

func (m *TWkWebview) Reload() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(13, m.Instance())
}

func (m *TWkWebview) Stop() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(14, m.Instance())
}

func (m *TWkWebview) StartDrag(formWindow lcl.IWinControl) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(formWindow))
}

func (m *TWkWebview) TerminateWebProcess() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(16, m.Instance())
}

func (m *TWkWebview) TryClose() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(17, m.Instance())
}

func (m *TWkWebview) ExecuteScript(script string, id int32) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(18, m.Instance(), api.PasStr(script), uintptr(id))
}

func (m *TWkWebview) LoadHTML(hTML string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(19, m.Instance(), api.PasStr(hTML))
}

func (m *TWkWebview) LoadURL(uRL string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(20, m.Instance(), api.PasStr(uRL))
}

func (m *TWkWebview) EnabledDevtools(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(21, m.Instance(), api.PasBool(enabled))
}

func (m *TWkWebview) ShowDevtools() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(22, m.Instance())
}

func (m *TWkWebview) RegisterScriptCode(scriptCode string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(23, m.Instance(), api.PasStr(scriptCode))
}

func (m *TWkWebview) RegisterScriptMessageHandler(name string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(24, m.Instance(), api.PasStr(name))
}

func (m *TWkWebview) SetSettings(setting IWkSettings) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(setting))
}

func (m *TWkWebview) FreeWebview() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(26, m.Instance())
}

func (m *TWkWebview) WebView() wvTypes.WebKitWebView {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(27, m.Instance())
	return wvTypes.WebKitWebView(r)
}

func (m *TWkWebview) SetOnLoadChange(fn TWkLoadChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkLoadChangeEvent(fn)
	base.SetEvent(m, 28, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnExecuteScriptFinished(fn TWkExecuteScriptFinishedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkExecuteScriptFinishedEvent(fn)
	base.SetEvent(m, 29, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnLoadFailed(fn TWkLoadFailedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkLoadFailedEvent(fn)
	base.SetEvent(m, 30, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnURISchemeRequest(fn TWkURISchemeRequestEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkURISchemeRequestEvent(fn)
	base.SetEvent(m, 31, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnProcessMessage(fn TWkProcessMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkProcessMessageEvent(fn)
	base.SetEvent(m, 32, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseMove(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 33, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMousePress(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 34, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseRelease(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 35, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnGetAcceptPolicyFinish(fn TWkGetAcceptPolicyFinishEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkGetAcceptPolicyFinishEvent(fn)
	base.SetEvent(m, 36, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnAddCookieFinish(fn TWkAddCookieFinishEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkAddCookieFinishEvent(fn)
	base.SetEvent(m, 37, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnGetCookiesFinish(fn TWkGetCookiesFinishEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkGetCookiesFinishEvent(fn)
	base.SetEvent(m, 38, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDeleteCookieFinish(fn TWkDeleteCookieFinishEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkDeleteCookieFinishEvent(fn)
	base.SetEvent(m, 39, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDecidePolicy(fn TWkDecidePolicyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkDecidePolicyEvent(fn)
	base.SetEvent(m, 40, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnWebProcessTerminated(fn TWkWebProcessTerminatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkWebProcessTerminatedEvent(fn)
	base.SetEvent(m, 41, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnContextMenu(fn TWkContextMenuEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkContextMenuEvent(fn)
	base.SetEvent(m, 42, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnContextMenuCommand(fn TWkContextMenuCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkContextMenuCommandEvent(fn)
	base.SetEvent(m, 43, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnContextMenuDismissed(fn TWkContextMenuDismissedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkContextMenuDismissedEvent(fn)
	base.SetEvent(m, 44, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) AsIntfWkCookieManagerDelegateEvent() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TWkWebview) AsIntfWkSchemeRequestDelegateEvent() uintptr {
	return m.GetIntfPointer(1)
}

// NewWebview class constructor
func NewWebview(owner lcl.IComponent) IWkWebview {
	var wkCookieManagerDelegateEventPtr uintptr // IWkCookieManagerDelegateEvent
	var wkSchemeRequestDelegateEventPtr uintptr // IWkSchemeRequestDelegateEvent
	r := wkWebviewAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&wkCookieManagerDelegateEventPtr)), uintptr(base.UnsafePointer(&wkSchemeRequestDelegateEventPtr)))
	ret := AsWkWebview(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(2)
		intf.SetIntfPointer(0, wkCookieManagerDelegateEventPtr)
		intf.SetIntfPointer(1, wkSchemeRequestDelegateEventPtr)
	}
	return ret
}

var (
	wkWebviewOnce   base.Once
	wkWebviewImport *imports.Imports = nil
)

func wkWebviewAPI() *imports.Imports {
	wkWebviewOnce.Do(func() {
		wkWebviewImport = api.NewDefaultImports()
		wkWebviewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebview_Create", 0), // constructor NewWebview
			/* 1 */ imports.NewTable("TWkWebview_CanGoBack", 0), // function CanGoBack
			/* 2 */ imports.NewTable("TWkWebview_CanGoForward", 0), // function CanGoForward
			/* 3 */ imports.NewTable("TWkWebview_IsLoading", 0), // function IsLoading
			/* 4 */ imports.NewTable("TWkWebview_GetTitle", 0), // function GetTitle
			/* 5 */ imports.NewTable("TWkWebview_GetURI", 0), // function GetURI
			/* 6 */ imports.NewTable("TWkWebview_CookieManager", 0), // function CookieManager
			/* 7 */ imports.NewTable("TWkWebview_GetSettings", 0), // function GetSettings
			/* 8 */ imports.NewTable("TWkWebview_AsCookieManagerDelegate", 0), // function AsCookieManagerDelegate
			/* 9 */ imports.NewTable("TWkWebview_AsSchemeRequestDelegate", 0), // function AsSchemeRequestDelegate
			/* 10 */ imports.NewTable("TWkWebview_CreateBrowser", 0), // procedure CreateBrowser
			/* 11 */ imports.NewTable("TWkWebview_GoBack", 0), // procedure GoBack
			/* 12 */ imports.NewTable("TWkWebview_GoForward", 0), // procedure GoForward
			/* 13 */ imports.NewTable("TWkWebview_Reload", 0), // procedure Reload
			/* 14 */ imports.NewTable("TWkWebview_Stop", 0), // procedure Stop
			/* 15 */ imports.NewTable("TWkWebview_StartDrag", 0), // procedure StartDrag
			/* 16 */ imports.NewTable("TWkWebview_TerminateWebProcess", 0), // procedure TerminateWebProcess
			/* 17 */ imports.NewTable("TWkWebview_TryClose", 0), // procedure TryClose
			/* 18 */ imports.NewTable("TWkWebview_ExecuteScript", 0), // procedure ExecuteScript
			/* 19 */ imports.NewTable("TWkWebview_LoadHTML", 0), // procedure LoadHTML
			/* 20 */ imports.NewTable("TWkWebview_LoadURL", 0), // procedure LoadURL
			/* 21 */ imports.NewTable("TWkWebview_EnabledDevtools", 0), // procedure EnabledDevtools
			/* 22 */ imports.NewTable("TWkWebview_ShowDevtools", 0), // procedure ShowDevtools
			/* 23 */ imports.NewTable("TWkWebview_RegisterScriptCode", 0), // procedure RegisterScriptCode
			/* 24 */ imports.NewTable("TWkWebview_RegisterScriptMessageHandler", 0), // procedure RegisterScriptMessageHandler
			/* 25 */ imports.NewTable("TWkWebview_SetSettings", 0), // procedure SetSettings
			/* 26 */ imports.NewTable("TWkWebview_FreeWebview", 0), // procedure FreeWebview
			/* 27 */ imports.NewTable("TWkWebview_WebView", 0), // property WebView
			/* 28 */ imports.NewTable("TWkWebview_OnLoadChange", 0), // event OnLoadChange
			/* 29 */ imports.NewTable("TWkWebview_OnExecuteScriptFinished", 0), // event OnExecuteScriptFinished
			/* 30 */ imports.NewTable("TWkWebview_OnLoadFailed", 0), // event OnLoadFailed
			/* 31 */ imports.NewTable("TWkWebview_OnURISchemeRequest", 0), // event OnURISchemeRequest
			/* 32 */ imports.NewTable("TWkWebview_OnProcessMessage", 0), // event OnProcessMessage
			/* 33 */ imports.NewTable("TWkWebview_OnMouseMove", 0), // event OnMouseMove
			/* 34 */ imports.NewTable("TWkWebview_OnMousePress", 0), // event OnMousePress
			/* 35 */ imports.NewTable("TWkWebview_OnMouseRelease", 0), // event OnMouseRelease
			/* 36 */ imports.NewTable("TWkWebview_OnGetAcceptPolicyFinish", 0), // event OnGetAcceptPolicyFinish
			/* 37 */ imports.NewTable("TWkWebview_OnAddCookieFinish", 0), // event OnAddCookieFinish
			/* 38 */ imports.NewTable("TWkWebview_OnGetCookiesFinish", 0), // event OnGetCookiesFinish
			/* 39 */ imports.NewTable("TWkWebview_OnDeleteCookieFinish", 0), // event OnDeleteCookieFinish
			/* 40 */ imports.NewTable("TWkWebview_OnDecidePolicy", 0), // event OnDecidePolicy
			/* 41 */ imports.NewTable("TWkWebview_OnWebProcessTerminated", 0), // event OnWebProcessTerminated
			/* 42 */ imports.NewTable("TWkWebview_OnContextMenu", 0), // event OnContextMenu
			/* 43 */ imports.NewTable("TWkWebview_OnContextMenuCommand", 0), // event OnContextMenuCommand
			/* 44 */ imports.NewTable("TWkWebview_OnContextMenuDismissed", 0), // event OnContextMenuDismissed
		}
	})
	return wkWebviewImport
}
