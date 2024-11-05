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

// IWkWebview Root Interface
type IWkWebview interface {
	IComponent
	WebView() WebKitWebView                                 // property
	CanGoBack() bool                                        // function
	CanGoForward() bool                                     // function
	IsLoading() bool                                        // function
	GetTitle() string                                       // function
	CookieManager() IWkCookieManager                        // function
	GetSettings() IWkSettings                               // function
	AsCookieManagerDelegate() IWkCookieManagerDelegateEvent // function
	AsSchemeRequestDelegate() IWkSchemeRequestDelegateEvent // function
	CreateBrowser()                                         // procedure
	GoBack()                                                // procedure
	GoForward()                                             // procedure
	Reload()                                                // procedure
	Stop()                                                  // procedure
	StartDrag(aFormWindow IWinControl)                      // procedure
	TerminateWebProcess()                                   // procedure
	TryClose()                                              // procedure
	ExecuteScript(aScript string)                           // procedure
	LoadHTML(aHTML string)                                  // procedure
	LoadURL(aURL string)                                    // procedure
	EnabledDevtools(enabled bool)                           // procedure
	ShowDevtools()                                          // procedure
	RegisterScriptCode(scriptCode string)                   // procedure
	// RegisterScriptMessageHandler
	//  注册JS消息处理器
	//  window.webkit.messageHandlers.<name>.postMessage()
	RegisterScriptMessageHandler(aName string) // procedure
	SetSettings(aSetting IWkSettings)          // procedure
	FreeWebview()                              // procedure
	// SetOnLoadChange
	//  当web_view中的加载操作改变时触发
	//  当发出新的加载请求时，WEBKIT_LOAD_STARTED发出信号，当加载成功完成或由于错误而完成时，WEBKIT_LOAD_FINISHED发出信号
	//  当正在进行的加载操作失败时，在WEBKIT_LOAD_FINISHED发出WebKitWebView::load-changed之前发出WebKitWebView::load-failed信号
	//  如果从服务器接收到重定向，则在WEBKIT_LOAD_STARTED初始发射之后，WEBKIT_LOAD_COMMITTED之前，用webkit_load_redirect发出此信号
	//  当页面内容开始到达时，使用WEBKIT_LOAD_COMMITTED事件发出信号
	//  https://webkitgtk.org/reference/webkit2gtk/stable/signal.WebView.load-changed.html
	SetOnLoadChange(fn TWkLoadChangeEvent) // property event
	// SetOnExecuteScriptFinished
	//  完成由webkit_web_view_run_javascript()开始的异步操作。
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.WebView.run_javascript_finish.html
	SetOnExecuteScriptFinished(fn TWkExecuteScriptFinishedEvent) // property event
	// SetOnLoadFailed
	//  在加载操作期间发生错误时触发。如果在开始加载页面数据时发生错误，load_event将是WEBKIT_LOAD_STARTED。如果它发生在加载提交的数据源时，load_event将是WEBKIT_LOAD_COMMITTED
	//  由于加载错误导致加载操作结束，WebKitWebView::load-changed信号将总是伴随着WEBKIT_LOAD_FINISHED事件紧随其后
	//  默认情况下，如果不处理该信号，将显示一个股票错误页面。如果希望提供自己的错误页面，则需要处理该信号
	//  默认的处理程序: 默认处理程序在通过g_signal_connect()添加处理程序后调用
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.CookieManager.delete_cookie.html?q=load-failed
	SetOnLoadFailed(fn TWkLoadFailedEvent) // property event
	// SetOnURISchemeRequest
	//  Webview 自定义协议: energy
	//  在上下文中注册scheme，这样当WebKitWebContext中发出带有scheme的URI请求时，注册的webkitur缺血erequestcallback将被webkitur缺血erequest调用
	//  可以异步处理URI方案请求，方法是在webkitur缺血性请求中调用g_object_ref()，然后在请求的数据可用时调用webkit_uri_scheme_request_finish()，或者在出现错误时调用webkit_uri_scheme_request_finish_error()
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.WebContext.register_uri_scheme.html
	SetOnURISchemeRequest(fn TWkURISchemeRequestEvent) // property event
	// SetOnProcessMessage
	//  Webview 自定义消息: processMessage, window.webkit.messageHandlers.<processMessage>.postMessage('data')
	//  当web视图中的JavaScript在使用webkit_user_content_manager_register_script_message_handler()注册后调用window.webkit.messageHandlers.<processMessage>.postMessage()时发出该信号。
	//  默认的处理程序: 默认处理程序在通过g_signal_connect()添加处理程序后调用。
	//  信号可以是详细的
	//  https://webkitgtk.org/reference/webkit2gtk/stable/signal.UserContentManager.script-message-received.html
	//  注册一个新的用户脚本消息处理程序
	//  注册后，脚本可以使用window.webkit.messageHandlers.<name>. postmessage(value)发送消息
	//  通过将处理程序连接到WebKitUserContentManager::script-message-received信号来接收这些消息
	//  处理程序名称用作信号的详细信息
	//  为了避免在注册处理程序名称和开始接收信号之间的竞争条件，建议在注册处理程序名称之前连接到信号:
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.UserContentManager.register_script_message_handler.html
	SetOnProcessMessage(fn TWkProcessMessageEvent) // property event
	// SetOnMousePress
	//  g_signal_connect_data button-press-event
	//  网页鼠标按下触发事件连接信号
	SetOnMousePress(fn TWkMousePressEvent) // property event
	// SetOnMouseRelease
	//  g_signal_connect_data button-release-event
	//  网页鼠标释放触发事件连接信号
	SetOnMouseRelease(fn TWkMouseReleaseEvent) // property event
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
	// SetOnDecidePolicy
	//  当WebKit请求客户端决定一个策略决策时，例如是否导航到一个页面，打开一个新窗口，或者是否下载一个资源，就会发出这个信号。
	//  传入的参数是泛型类型，但在做出决策时应将其强制转换为更具体的类型。例如:WebKitNavigationPolicyDecisiondecision
	//  可以通过简单地调用参数并返回阻塞默认信号处理程序来异步地做出策略决策。
	//  如果最后一个引用被删除，并且没有明确做出决策，则将是默认的策略决策。
	//  默认信号处理程序将简单地调用webkit_policy_decision_use()。
	//  只有为给定选择的第一个策略决策才会有任何影响, g_object_ref() decision TRUE WebKitPolicyDecision webkit_policy_decision_use()
	//  https://webkitgtk.org/reference/webkit2gtk/stable/signal.WebView.decide-policy.html
	SetOnDecidePolicy(fn TWkDecidePolicyEvent) // property event
	// SetOnWebProcessTerminated
	//  当web进程由于reason异常终止时，发出此信号
	//  默认处理程序在通过g_signal_connect()添加处理程序后被调用
	//  https://webkitgtk.org/reference/webkit2gtk/stable/method.WebView.terminate_web_process.html
	SetOnWebProcessTerminated(fn TWkWebProcessTerminatedEvent) // property event
	SetOnContextMenu(fn TWkContextMenuEvent)                   // property event
	SetOnContextMenuCommand(fn TWkContextMenuCommandEvent)     // property event
	SetOnContextMenuDismissed(fn TWkContextMenuDismissedEvent) // property event
}

// TWkWebview Root Object
type TWkWebview struct {
	TComponent
	loadChangePtr            uintptr
	executeScriptFinishedPtr uintptr
	loadFailedPtr            uintptr
	uRISchemeRequestPtr      uintptr
	processMessagePtr        uintptr
	mousePressPtr            uintptr
	mouseReleasePtr          uintptr
	getAcceptPolicyFinishPtr uintptr
	addCookieFinishPtr       uintptr
	getCookiesFinishPtr      uintptr
	deleteCookieFinishPtr    uintptr
	decidePolicyPtr          uintptr
	webProcessTerminatedPtr  uintptr
	contextMenuPtr           uintptr
	contextMenuCommandPtr    uintptr
	contextMenuDismissedPtr  uintptr
}

func NewWkWebview(aOwner IComponent) IWkWebview {
	r1 := wkWebviewImportAPI().SysCallN(5, GetObjectUintptr(aOwner))
	return AsWkWebview(r1)
}

func (m *TWkWebview) WebView() WebKitWebView {
	r1 := wkWebviewImportAPI().SysCallN(42, m.Instance())
	return WebKitWebView(r1)
}

func (m *TWkWebview) CanGoBack() bool {
	r1 := wkWebviewImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) CanGoForward() bool {
	r1 := wkWebviewImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) IsLoading() bool {
	r1 := wkWebviewImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) GetTitle() string {
	r1 := wkWebviewImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TWkWebview) CookieManager() IWkCookieManager {
	r1 := wkWebviewImportAPI().SysCallN(4, m.Instance())
	return AsWkCookieManager(r1)
}

func (m *TWkWebview) GetSettings() IWkSettings {
	r1 := wkWebviewImportAPI().SysCallN(10, m.Instance())
	return AsWkSettings(r1)
}

func (m *TWkWebview) AsCookieManagerDelegate() IWkCookieManagerDelegateEvent {
	var resultWkCookieManagerDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultWkCookieManagerDelegateEvent)))
	return AsWkCookieManagerDelegateEvent(resultWkCookieManagerDelegateEvent)
}

func (m *TWkWebview) AsSchemeRequestDelegate() IWkSchemeRequestDelegateEvent {
	var resultWkSchemeRequestDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultWkSchemeRequestDelegateEvent)))
	return AsWkSchemeRequestDelegateEvent(resultWkSchemeRequestDelegateEvent)
}

func (m *TWkWebview) CreateBrowser() {
	wkWebviewImportAPI().SysCallN(6, m.Instance())
}

func (m *TWkWebview) GoBack() {
	wkWebviewImportAPI().SysCallN(12, m.Instance())
}

func (m *TWkWebview) GoForward() {
	wkWebviewImportAPI().SysCallN(13, m.Instance())
}

func (m *TWkWebview) Reload() {
	wkWebviewImportAPI().SysCallN(19, m.Instance())
}

func (m *TWkWebview) Stop() {
	wkWebviewImportAPI().SysCallN(39, m.Instance())
}

func (m *TWkWebview) StartDrag(aFormWindow IWinControl) {
	wkWebviewImportAPI().SysCallN(38, m.Instance(), GetObjectUintptr(aFormWindow))
}

func (m *TWkWebview) TerminateWebProcess() {
	wkWebviewImportAPI().SysCallN(40, m.Instance())
}

func (m *TWkWebview) TryClose() {
	wkWebviewImportAPI().SysCallN(41, m.Instance())
}

func (m *TWkWebview) ExecuteScript(aScript string) {
	wkWebviewImportAPI().SysCallN(8, m.Instance(), PascalStr(aScript))
}

func (m *TWkWebview) LoadHTML(aHTML string) {
	wkWebviewImportAPI().SysCallN(15, m.Instance(), PascalStr(aHTML))
}

func (m *TWkWebview) LoadURL(aURL string) {
	wkWebviewImportAPI().SysCallN(16, m.Instance(), PascalStr(aURL))
}

func (m *TWkWebview) EnabledDevtools(enabled bool) {
	wkWebviewImportAPI().SysCallN(7, m.Instance(), PascalBool(enabled))
}

func (m *TWkWebview) ShowDevtools() {
	wkWebviewImportAPI().SysCallN(37, m.Instance())
}

func (m *TWkWebview) RegisterScriptCode(scriptCode string) {
	wkWebviewImportAPI().SysCallN(17, m.Instance(), PascalStr(scriptCode))
}

func (m *TWkWebview) RegisterScriptMessageHandler(aName string) {
	wkWebviewImportAPI().SysCallN(18, m.Instance(), PascalStr(aName))
}

func (m *TWkWebview) SetSettings(aSetting IWkSettings) {
	wkWebviewImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(aSetting))
}

func (m *TWkWebview) FreeWebview() {
	wkWebviewImportAPI().SysCallN(9, m.Instance())
}

func (m *TWkWebview) SetOnLoadChange(fn TWkLoadChangeEvent) {
	if m.loadChangePtr != 0 {
		RemoveEventElement(m.loadChangePtr)
	}
	m.loadChangePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(29, m.Instance(), m.loadChangePtr)
}

func (m *TWkWebview) SetOnExecuteScriptFinished(fn TWkExecuteScriptFinishedEvent) {
	if m.executeScriptFinishedPtr != 0 {
		RemoveEventElement(m.executeScriptFinishedPtr)
	}
	m.executeScriptFinishedPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(26, m.Instance(), m.executeScriptFinishedPtr)
}

func (m *TWkWebview) SetOnLoadFailed(fn TWkLoadFailedEvent) {
	if m.loadFailedPtr != 0 {
		RemoveEventElement(m.loadFailedPtr)
	}
	m.loadFailedPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(30, m.Instance(), m.loadFailedPtr)
}

func (m *TWkWebview) SetOnURISchemeRequest(fn TWkURISchemeRequestEvent) {
	if m.uRISchemeRequestPtr != 0 {
		RemoveEventElement(m.uRISchemeRequestPtr)
	}
	m.uRISchemeRequestPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(34, m.Instance(), m.uRISchemeRequestPtr)
}

func (m *TWkWebview) SetOnProcessMessage(fn TWkProcessMessageEvent) {
	if m.processMessagePtr != 0 {
		RemoveEventElement(m.processMessagePtr)
	}
	m.processMessagePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(33, m.Instance(), m.processMessagePtr)
}

func (m *TWkWebview) SetOnMousePress(fn TWkMousePressEvent) {
	if m.mousePressPtr != 0 {
		RemoveEventElement(m.mousePressPtr)
	}
	m.mousePressPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(31, m.Instance(), m.mousePressPtr)
}

func (m *TWkWebview) SetOnMouseRelease(fn TWkMouseReleaseEvent) {
	if m.mouseReleasePtr != 0 {
		RemoveEventElement(m.mouseReleasePtr)
	}
	m.mouseReleasePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(32, m.Instance(), m.mouseReleasePtr)
}

func (m *TWkWebview) SetOnGetAcceptPolicyFinish(fn TWkGetAcceptPolicyFinishEvent) {
	if m.getAcceptPolicyFinishPtr != 0 {
		RemoveEventElement(m.getAcceptPolicyFinishPtr)
	}
	m.getAcceptPolicyFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(27, m.Instance(), m.getAcceptPolicyFinishPtr)
}

func (m *TWkWebview) SetOnAddCookieFinish(fn TWkAddCookieFinishEvent) {
	if m.addCookieFinishPtr != 0 {
		RemoveEventElement(m.addCookieFinishPtr)
	}
	m.addCookieFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(20, m.Instance(), m.addCookieFinishPtr)
}

func (m *TWkWebview) SetOnGetCookiesFinish(fn TWkGetCookiesFinishEvent) {
	if m.getCookiesFinishPtr != 0 {
		RemoveEventElement(m.getCookiesFinishPtr)
	}
	m.getCookiesFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(28, m.Instance(), m.getCookiesFinishPtr)
}

func (m *TWkWebview) SetOnDeleteCookieFinish(fn TWkDeleteCookieFinishEvent) {
	if m.deleteCookieFinishPtr != 0 {
		RemoveEventElement(m.deleteCookieFinishPtr)
	}
	m.deleteCookieFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(25, m.Instance(), m.deleteCookieFinishPtr)
}

func (m *TWkWebview) SetOnDecidePolicy(fn TWkDecidePolicyEvent) {
	if m.decidePolicyPtr != 0 {
		RemoveEventElement(m.decidePolicyPtr)
	}
	m.decidePolicyPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(24, m.Instance(), m.decidePolicyPtr)
}

func (m *TWkWebview) SetOnWebProcessTerminated(fn TWkWebProcessTerminatedEvent) {
	if m.webProcessTerminatedPtr != 0 {
		RemoveEventElement(m.webProcessTerminatedPtr)
	}
	m.webProcessTerminatedPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(35, m.Instance(), m.webProcessTerminatedPtr)
}

func (m *TWkWebview) SetOnContextMenu(fn TWkContextMenuEvent) {
	if m.contextMenuPtr != 0 {
		RemoveEventElement(m.contextMenuPtr)
	}
	m.contextMenuPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(21, m.Instance(), m.contextMenuPtr)
}

func (m *TWkWebview) SetOnContextMenuCommand(fn TWkContextMenuCommandEvent) {
	if m.contextMenuCommandPtr != 0 {
		RemoveEventElement(m.contextMenuCommandPtr)
	}
	m.contextMenuCommandPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(22, m.Instance(), m.contextMenuCommandPtr)
}

func (m *TWkWebview) SetOnContextMenuDismissed(fn TWkContextMenuDismissedEvent) {
	if m.contextMenuDismissedPtr != 0 {
		RemoveEventElement(m.contextMenuDismissedPtr)
	}
	m.contextMenuDismissedPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(23, m.Instance(), m.contextMenuDismissedPtr)
}

var (
	wkWebviewImport       *imports.Imports = nil
	wkWebviewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebview_AsCookieManagerDelegate", 0),
		/*1*/ imports.NewTable("WkWebview_AsSchemeRequestDelegate", 0),
		/*2*/ imports.NewTable("WkWebview_CanGoBack", 0),
		/*3*/ imports.NewTable("WkWebview_CanGoForward", 0),
		/*4*/ imports.NewTable("WkWebview_CookieManager", 0),
		/*5*/ imports.NewTable("WkWebview_Create", 0),
		/*6*/ imports.NewTable("WkWebview_CreateBrowser", 0),
		/*7*/ imports.NewTable("WkWebview_EnabledDevtools", 0),
		/*8*/ imports.NewTable("WkWebview_ExecuteScript", 0),
		/*9*/ imports.NewTable("WkWebview_FreeWebview", 0),
		/*10*/ imports.NewTable("WkWebview_GetSettings", 0),
		/*11*/ imports.NewTable("WkWebview_GetTitle", 0),
		/*12*/ imports.NewTable("WkWebview_GoBack", 0),
		/*13*/ imports.NewTable("WkWebview_GoForward", 0),
		/*14*/ imports.NewTable("WkWebview_IsLoading", 0),
		/*15*/ imports.NewTable("WkWebview_LoadHTML", 0),
		/*16*/ imports.NewTable("WkWebview_LoadURL", 0),
		/*17*/ imports.NewTable("WkWebview_RegisterScriptCode", 0),
		/*18*/ imports.NewTable("WkWebview_RegisterScriptMessageHandler", 0),
		/*19*/ imports.NewTable("WkWebview_Reload", 0),
		/*20*/ imports.NewTable("WkWebview_SetOnAddCookieFinish", 0),
		/*21*/ imports.NewTable("WkWebview_SetOnContextMenu", 0),
		/*22*/ imports.NewTable("WkWebview_SetOnContextMenuCommand", 0),
		/*23*/ imports.NewTable("WkWebview_SetOnContextMenuDismissed", 0),
		/*24*/ imports.NewTable("WkWebview_SetOnDecidePolicy", 0),
		/*25*/ imports.NewTable("WkWebview_SetOnDeleteCookieFinish", 0),
		/*26*/ imports.NewTable("WkWebview_SetOnExecuteScriptFinished", 0),
		/*27*/ imports.NewTable("WkWebview_SetOnGetAcceptPolicyFinish", 0),
		/*28*/ imports.NewTable("WkWebview_SetOnGetCookiesFinish", 0),
		/*29*/ imports.NewTable("WkWebview_SetOnLoadChange", 0),
		/*30*/ imports.NewTable("WkWebview_SetOnLoadFailed", 0),
		/*31*/ imports.NewTable("WkWebview_SetOnMousePress", 0),
		/*32*/ imports.NewTable("WkWebview_SetOnMouseRelease", 0),
		/*33*/ imports.NewTable("WkWebview_SetOnProcessMessage", 0),
		/*34*/ imports.NewTable("WkWebview_SetOnURISchemeRequest", 0),
		/*35*/ imports.NewTable("WkWebview_SetOnWebProcessTerminated", 0),
		/*36*/ imports.NewTable("WkWebview_SetSettings", 0),
		/*37*/ imports.NewTable("WkWebview_ShowDevtools", 0),
		/*38*/ imports.NewTable("WkWebview_StartDrag", 0),
		/*39*/ imports.NewTable("WkWebview_Stop", 0),
		/*40*/ imports.NewTable("WkWebview_TerminateWebProcess", 0),
		/*41*/ imports.NewTable("WkWebview_TryClose", 0),
		/*42*/ imports.NewTable("WkWebview_WebView", 0),
	}
)

func wkWebviewImportAPI() *imports.Imports {
	if wkWebviewImport == nil {
		wkWebviewImport = NewDefaultImports()
		wkWebviewImport.SetImportTable(wkWebviewImportTables)
		wkWebviewImportTables = nil
	}
	return wkWebviewImport
}
