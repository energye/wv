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
	Data() WkWebview                               // function
	Configuration() WKWebViewConfiguration         // function
	LoadRequest(request NSURLRequest) WKNavigation // function
	Title() string                                 // function
	URL() NSURL                                    // function
	IsLoading() bool                               // function
	// EstimatedProgress
	//  根据接收到的总字节数，此值的范围为0.0到1.0，包括主文档及其所有潜在的子资源。
	//  导航加载完成后，估计进度值保持在1.0，直到新的导航开始，此时估计进度值重置为0.0。
	//  WKWebView类符合此属性的键值观察（KVO）标准。
	EstimatedProgress() (resultFloat64 float64)                                      // function
	HasOnlySecureContent() bool                                                      // function
	CanGoBack() bool                                                                 // function
	CanGoForward() bool                                                              // function
	GoBack() WKNavigation                                                            // function
	GoForward() WKNavigation                                                         // function
	Reload() WKNavigation                                                            // function
	ReloadFromOrigin() WKNavigation                                                  // function
	AllowsBackForwardNavigationGestures() bool                                       // function
	AllowsMagnification() bool                                                       // function
	Magnification() (resultFloat64 float64)                                          // function
	AsReceiveScriptMessageDelegate() IReceiveScriptMessageDelegateEvent              // function
	AsWKNavigationDelegate() IWKNavigationDelegateEvent                              // function
	AsWKURLSchemeHandlerDelegate() IWKURLSchemeHandlerDelegateEvent                  // function
	AsWKUIDelegate() IWKUIDelegateEvent                                              // function
	AsWKDownloadDelegate() IWKDownloadDelegateEvent                                  // function
	LoadHTML(aHtml, aBaseURL string)                                                 // procedure
	LoadURL(aURL string)                                                             // procedure
	SetNavigationDelegate(navigationDelegate WKNavigationDelegate)                   // procedure
	SetUIDelegate(uiDelegate WKUIDelegate)                                           // procedure
	InitWithFrameConfiguration(aFrame *TRect, aConfiguration WKWebViewConfiguration) // procedure
	StopLoading()                                                                    // procedure
	EvaluateJavaScript(javaScriptString string)                                      // procedure
	SetAllowsBackForwardNavigationGestures(newValue bool)                            // procedure
	SetAllowsMagnification(newValue bool)                                            // procedure
	SetMagnification(newValue float64)                                               // procedure
	SetMagnificationCenteredAtPoint(magnification float64, point *TPoint)            // procedure
	RemoveFromSuperview()
	RemoveAllSubviews()
	// SetOnProcessMessage
	//  WKScriptMessageHandlerProtocol
	SetOnProcessMessage(fn TWkProcessMessageEvent) // property event
	// SetOnDecidePolicyForNavigationActionPreferences
	//  WKNavigationDelegateProtocol
	SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences)       // property event
	SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse)                         // property event
	SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation)                                       // property event
	SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation) // property event
	SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError)                       // property event
	SetOnCommitNavigation(fn TWkCommitNavigation)                                                           // property event
	SetOnFinishNavigation(fn TWkFinishNavigation)                                                           // property event
	SetOnFailNavigationWithError(fn TWkFailNavigationWithError)                                             // property event
	// SetOnWebContentProcessDidTerminate
	//  property OnReceiveAuthenticationChallenge: TWkReceiveAuthenticationChallenge read FOnReceiveAuthenticationChallenge write FOnReceiveAuthenticationChallenge;
	SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate) // property event
	// SetOnNavigationActionDidBecomeDownload
	//  property OnAuthenticationChallengeShouldAllowDeprecatedTLS: TWkAuthenticationChallengeShouldAllowDeprecatedTLS read FOnAuthenticationChallengeShouldAllowDeprecatedTLS write FOnAuthenticationChallengeShouldAllowDeprecatedTLS;
	SetOnNavigationActionDidBecomeDownload(fn TWkNavigationActionDidBecomeDownload)     // property event
	SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload) // property event
	// SetOnStartURLSchemeTask
	//  WKURLSchemeHandlerProtocol
	SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask) // property event
	SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask)   // property event
	// SetOnCreateWebView
	//  WKUIDelegateProtocol
	SetOnCreateWebView(fn TWKCreateWebView)                                       // property event
	SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert)                             // property event
	SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion)     // property event
	SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion) // property event
	SetOnWebViewDidClose(fn TWKWebViewDidClose)                                   // property event
	// SetOnDownloadCancelCompletionHandler
	//  WKDownload
	SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler)                                                 // property event
	SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) // property event
	SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest)                       // property event
	// SetOnDownloadFinish
	//  property OnDownloadReceiveAuthenticationChallenge: TWKDownloadReceiveAuthenticationChallenge read FOnDownloadReceiveAuthenticationChallenge write FOnDownloadReceiveAuthenticationChallenge;
	SetOnDownloadFinish(fn TWKDownloadFinish)               // property event
	SetOnDownloadFailWithError(fn TWKDownloadFailWithError) // property event
}

// TWkWebview Root Object
type TWkWebview struct {
	TComponent
	processMessagePtr                                          uintptr
	decidePolicyForNavigationActionPreferencesPtr              uintptr
	decidePolicyForNavigationResponsePtr                       uintptr
	startProvisionalNavigationPtr                              uintptr
	receiveServerRedirectForProvisionalNavigationPtr           uintptr
	failProvisionalNavigationWithErrorPtr                      uintptr
	commitNavigationPtr                                        uintptr
	finishNavigationPtr                                        uintptr
	failNavigationWithErrorPtr                                 uintptr
	webContentProcessDidTerminatePtr                           uintptr
	navigationActionDidBecomeDownloadPtr                       uintptr
	navigationResponseDidBecomeDownloadPtr                     uintptr
	startURLSchemeTaskPtr                                      uintptr
	stopURLSchemeTaskPtr                                       uintptr
	createWebViewPtr                                           uintptr
	runJavaScriptAlertPtr                                      uintptr
	runJavaScriptConfirmCompletionPtr                          uintptr
	runJavaScriptTextInputCompletionPtr                        uintptr
	webViewDidClosePtr                                         uintptr
	downloadCancelCompletionHandlerPtr                         uintptr
	downloadDecideDestinationUsingResponseSuggestedFilenamePtr uintptr
	downloadWillPerformHTTPRedirectionNewRequestPtr            uintptr
	downloadFinishPtr                                          uintptr
	downloadFailWithErrorPtr                                   uintptr
}

func NewWkWebview(aOwner IComponent) IWkWebview {
	r1 := wkWebviewImportAPI().SysCallN(10, GetObjectUintptr(aOwner))
	return AsWkWebview(r1)
}

func (m *TWkWebview) Data() WkWebview {
	r1 := wkWebviewImportAPI().SysCallN(11, m.Instance())
	return WkWebview(r1)
}

func (m *TWkWebview) Configuration() WKWebViewConfiguration {
	r1 := wkWebviewImportAPI().SysCallN(9, m.Instance())
	return WKWebViewConfiguration(r1)
}

func (m *TWkWebview) LoadRequest(request NSURLRequest) WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(20, m.Instance(), uintptr(request))
	return WKNavigation(r1)
}

func (m *TWkWebview) Title() string {
	r1 := wkWebviewImportAPI().SysCallN(56, m.Instance())
	return GoStr(r1)
}

func (m *TWkWebview) URL() NSURL {
	r1 := wkWebviewImportAPI().SysCallN(57, m.Instance())
	return NSURL(r1)
}

func (m *TWkWebview) IsLoading() bool {
	r1 := wkWebviewImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) EstimatedProgress() (resultFloat64 float64) {
	wkWebviewImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TWkWebview) HasOnlySecureContent() bool {
	r1 := wkWebviewImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) CanGoBack() bool {
	r1 := wkWebviewImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) CanGoForward() bool {
	r1 := wkWebviewImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) GoBack() WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(14, m.Instance())
	return WKNavigation(r1)
}

func (m *TWkWebview) GoForward() WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(15, m.Instance())
	return WKNavigation(r1)
}

func (m *TWkWebview) Reload() WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(23, m.Instance())
	return WKNavigation(r1)
}

func (m *TWkWebview) ReloadFromOrigin() WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(24, m.Instance())
	return WKNavigation(r1)
}

func (m *TWkWebview) AllowsBackForwardNavigationGestures() bool {
	r1 := wkWebviewImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) AllowsMagnification() bool {
	r1 := wkWebviewImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TWkWebview) Magnification() (resultFloat64 float64) {
	wkWebviewImportAPI().SysCallN(22, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TWkWebview) AsReceiveScriptMessageDelegate() IReceiveScriptMessageDelegateEvent {
	var resultReceiveScriptMessageDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultReceiveScriptMessageDelegateEvent)))
	return AsReceiveScriptMessageDelegateEvent(resultReceiveScriptMessageDelegateEvent)
}

func (m *TWkWebview) AsWKNavigationDelegate() IWKNavigationDelegateEvent {
	var resultWKNavigationDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultWKNavigationDelegateEvent)))
	return AsWKNavigationDelegateEvent(resultWKNavigationDelegateEvent)
}

func (m *TWkWebview) AsWKURLSchemeHandlerDelegate() IWKURLSchemeHandlerDelegateEvent {
	var resultWKURLSchemeHandlerDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultWKURLSchemeHandlerDelegateEvent)))
	return AsWKURLSchemeHandlerDelegateEvent(resultWKURLSchemeHandlerDelegateEvent)
}

func (m *TWkWebview) AsWKUIDelegate() IWKUIDelegateEvent {
	var resultWKUIDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&resultWKUIDelegateEvent)))
	return AsWKUIDelegateEvent(resultWKUIDelegateEvent)
}

func (m *TWkWebview) AsWKDownloadDelegate() IWKDownloadDelegateEvent {
	var resultWKDownloadDelegateEvent uintptr
	wkWebviewImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultWKDownloadDelegateEvent)))
	return AsWKDownloadDelegateEvent(resultWKDownloadDelegateEvent)
}

func (m *TWkWebview) LoadHTML(aHtml, aBaseURL string) {
	wkWebviewImportAPI().SysCallN(19, m.Instance(), PascalStr(aHtml), PascalStr(aBaseURL))
}

func (m *TWkWebview) LoadURL(aURL string) {
	wkWebviewImportAPI().SysCallN(21, m.Instance(), PascalStr(aURL))
}

func (m *TWkWebview) SetNavigationDelegate(navigationDelegate WKNavigationDelegate) {
	wkWebviewImportAPI().SysCallN(29, m.Instance(), uintptr(navigationDelegate))
}

func (m *TWkWebview) SetUIDelegate(uiDelegate WKUIDelegate) {
	wkWebviewImportAPI().SysCallN(54, m.Instance(), uintptr(uiDelegate))
}

func (m *TWkWebview) InitWithFrameConfiguration(aFrame *TRect, aConfiguration WKWebViewConfiguration) {
	wkWebviewImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(aFrame)), uintptr(aConfiguration))
}

func (m *TWkWebview) StopLoading() {
	wkWebviewImportAPI().SysCallN(55, m.Instance())
}

func (m *TWkWebview) EvaluateJavaScript(javaScriptString string) {
	wkWebviewImportAPI().SysCallN(13, m.Instance(), PascalStr(javaScriptString))
}

func (m *TWkWebview) SetAllowsBackForwardNavigationGestures(newValue bool) {
	wkWebviewImportAPI().SysCallN(25, m.Instance(), PascalBool(newValue))
}

func (m *TWkWebview) SetAllowsMagnification(newValue bool) {
	wkWebviewImportAPI().SysCallN(26, m.Instance(), PascalBool(newValue))
}

func (m *TWkWebview) SetMagnification(newValue float64) {
	wkWebviewImportAPI().SysCallN(27, m.Instance(), uintptr(unsafePointer(&newValue)))
}

func (m *TWkWebview) SetMagnificationCenteredAtPoint(magnification float64, point *TPoint) {
	wkWebviewImportAPI().SysCallN(28, m.Instance(), uintptr(unsafePointer(&magnification)), uintptr(unsafePointer(point)))
}
func (m *TWkWebview) RemoveFromSuperview() {
	wkWebviewImportAPI().SysCallN(58, m.Instance())
}
func (m *TWkWebview) RemoveAllSubviews() {
	wkWebviewImportAPI().SysCallN(59, m.Instance())
}

func (m *TWkWebview) SetOnProcessMessage(fn TWkProcessMessageEvent) {
	if m.processMessagePtr != 0 {
		RemoveEventElement(m.processMessagePtr)
	}
	m.processMessagePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(44, m.Instance(), m.processMessagePtr)
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences) {
	if m.decidePolicyForNavigationActionPreferencesPtr != 0 {
		RemoveEventElement(m.decidePolicyForNavigationActionPreferencesPtr)
	}
	m.decidePolicyForNavigationActionPreferencesPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(32, m.Instance(), m.decidePolicyForNavigationActionPreferencesPtr)
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse) {
	if m.decidePolicyForNavigationResponsePtr != 0 {
		RemoveEventElement(m.decidePolicyForNavigationResponsePtr)
	}
	m.decidePolicyForNavigationResponsePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(33, m.Instance(), m.decidePolicyForNavigationResponsePtr)
}

func (m *TWkWebview) SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation) {
	if m.startProvisionalNavigationPtr != 0 {
		RemoveEventElement(m.startProvisionalNavigationPtr)
	}
	m.startProvisionalNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(49, m.Instance(), m.startProvisionalNavigationPtr)
}

func (m *TWkWebview) SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation) {
	if m.receiveServerRedirectForProvisionalNavigationPtr != 0 {
		RemoveEventElement(m.receiveServerRedirectForProvisionalNavigationPtr)
	}
	m.receiveServerRedirectForProvisionalNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(45, m.Instance(), m.receiveServerRedirectForProvisionalNavigationPtr)
}

func (m *TWkWebview) SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError) {
	if m.failProvisionalNavigationWithErrorPtr != 0 {
		RemoveEventElement(m.failProvisionalNavigationWithErrorPtr)
	}
	m.failProvisionalNavigationWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(40, m.Instance(), m.failProvisionalNavigationWithErrorPtr)
}

func (m *TWkWebview) SetOnCommitNavigation(fn TWkCommitNavigation) {
	if m.commitNavigationPtr != 0 {
		RemoveEventElement(m.commitNavigationPtr)
	}
	m.commitNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(30, m.Instance(), m.commitNavigationPtr)
}

func (m *TWkWebview) SetOnFinishNavigation(fn TWkFinishNavigation) {
	if m.finishNavigationPtr != 0 {
		RemoveEventElement(m.finishNavigationPtr)
	}
	m.finishNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(41, m.Instance(), m.finishNavigationPtr)
}

func (m *TWkWebview) SetOnFailNavigationWithError(fn TWkFailNavigationWithError) {
	if m.failNavigationWithErrorPtr != 0 {
		RemoveEventElement(m.failNavigationWithErrorPtr)
	}
	m.failNavigationWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(39, m.Instance(), m.failNavigationWithErrorPtr)
}

func (m *TWkWebview) SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate) {
	if m.webContentProcessDidTerminatePtr != 0 {
		RemoveEventElement(m.webContentProcessDidTerminatePtr)
	}
	m.webContentProcessDidTerminatePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(52, m.Instance(), m.webContentProcessDidTerminatePtr)
}

func (m *TWkWebview) SetOnNavigationActionDidBecomeDownload(fn TWkNavigationActionDidBecomeDownload) {
	if m.navigationActionDidBecomeDownloadPtr != 0 {
		RemoveEventElement(m.navigationActionDidBecomeDownloadPtr)
	}
	m.navigationActionDidBecomeDownloadPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(42, m.Instance(), m.navigationActionDidBecomeDownloadPtr)
}

func (m *TWkWebview) SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload) {
	if m.navigationResponseDidBecomeDownloadPtr != 0 {
		RemoveEventElement(m.navigationResponseDidBecomeDownloadPtr)
	}
	m.navigationResponseDidBecomeDownloadPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(43, m.Instance(), m.navigationResponseDidBecomeDownloadPtr)
}

func (m *TWkWebview) SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask) {
	if m.startURLSchemeTaskPtr != 0 {
		RemoveEventElement(m.startURLSchemeTaskPtr)
	}
	m.startURLSchemeTaskPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(50, m.Instance(), m.startURLSchemeTaskPtr)
}

func (m *TWkWebview) SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask) {
	if m.stopURLSchemeTaskPtr != 0 {
		RemoveEventElement(m.stopURLSchemeTaskPtr)
	}
	m.stopURLSchemeTaskPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(51, m.Instance(), m.stopURLSchemeTaskPtr)
}

func (m *TWkWebview) SetOnCreateWebView(fn TWKCreateWebView) {
	if m.createWebViewPtr != 0 {
		RemoveEventElement(m.createWebViewPtr)
	}
	m.createWebViewPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(31, m.Instance(), m.createWebViewPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert) {
	if m.runJavaScriptAlertPtr != 0 {
		RemoveEventElement(m.runJavaScriptAlertPtr)
	}
	m.runJavaScriptAlertPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(46, m.Instance(), m.runJavaScriptAlertPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion) {
	if m.runJavaScriptConfirmCompletionPtr != 0 {
		RemoveEventElement(m.runJavaScriptConfirmCompletionPtr)
	}
	m.runJavaScriptConfirmCompletionPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(47, m.Instance(), m.runJavaScriptConfirmCompletionPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion) {
	if m.runJavaScriptTextInputCompletionPtr != 0 {
		RemoveEventElement(m.runJavaScriptTextInputCompletionPtr)
	}
	m.runJavaScriptTextInputCompletionPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(48, m.Instance(), m.runJavaScriptTextInputCompletionPtr)
}

func (m *TWkWebview) SetOnWebViewDidClose(fn TWKWebViewDidClose) {
	if m.webViewDidClosePtr != 0 {
		RemoveEventElement(m.webViewDidClosePtr)
	}
	m.webViewDidClosePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(53, m.Instance(), m.webViewDidClosePtr)
}

func (m *TWkWebview) SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler) {
	if m.downloadCancelCompletionHandlerPtr != 0 {
		RemoveEventElement(m.downloadCancelCompletionHandlerPtr)
	}
	m.downloadCancelCompletionHandlerPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(34, m.Instance(), m.downloadCancelCompletionHandlerPtr)
}

func (m *TWkWebview) SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) {
	if m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr != 0 {
		RemoveEventElement(m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr)
	}
	m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(35, m.Instance(), m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr)
}

func (m *TWkWebview) SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest) {
	if m.downloadWillPerformHTTPRedirectionNewRequestPtr != 0 {
		RemoveEventElement(m.downloadWillPerformHTTPRedirectionNewRequestPtr)
	}
	m.downloadWillPerformHTTPRedirectionNewRequestPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(38, m.Instance(), m.downloadWillPerformHTTPRedirectionNewRequestPtr)
}

func (m *TWkWebview) SetOnDownloadFinish(fn TWKDownloadFinish) {
	if m.downloadFinishPtr != 0 {
		RemoveEventElement(m.downloadFinishPtr)
	}
	m.downloadFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(37, m.Instance(), m.downloadFinishPtr)
}

func (m *TWkWebview) SetOnDownloadFailWithError(fn TWKDownloadFailWithError) {
	if m.downloadFailWithErrorPtr != 0 {
		RemoveEventElement(m.downloadFailWithErrorPtr)
	}
	m.downloadFailWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(36, m.Instance(), m.downloadFailWithErrorPtr)
}

var (
	wkWebviewImport       *imports.Imports = nil
	wkWebviewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebview_AllowsBackForwardNavigationGestures", 0),
		/*1*/ imports.NewTable("WkWebview_AllowsMagnification", 0),
		/*2*/ imports.NewTable("WkWebview_AsReceiveScriptMessageDelegate", 0),
		/*3*/ imports.NewTable("WkWebview_AsWKDownloadDelegate", 0),
		/*4*/ imports.NewTable("WkWebview_AsWKNavigationDelegate", 0),
		/*5*/ imports.NewTable("WkWebview_AsWKUIDelegate", 0),
		/*6*/ imports.NewTable("WkWebview_AsWKURLSchemeHandlerDelegate", 0),
		/*7*/ imports.NewTable("WkWebview_CanGoBack", 0),
		/*8*/ imports.NewTable("WkWebview_CanGoForward", 0),
		/*9*/ imports.NewTable("WkWebview_Configuration", 0),
		/*10*/ imports.NewTable("WkWebview_Create", 0),
		/*11*/ imports.NewTable("WkWebview_Data", 0),
		/*12*/ imports.NewTable("WkWebview_EstimatedProgress", 0),
		/*13*/ imports.NewTable("WkWebview_EvaluateJavaScript", 0),
		/*14*/ imports.NewTable("WkWebview_GoBack", 0),
		/*15*/ imports.NewTable("WkWebview_GoForward", 0),
		/*16*/ imports.NewTable("WkWebview_HasOnlySecureContent", 0),
		/*17*/ imports.NewTable("WkWebview_InitWithFrameConfiguration", 0),
		/*18*/ imports.NewTable("WkWebview_IsLoading", 0),
		/*19*/ imports.NewTable("WkWebview_LoadHTML", 0),
		/*20*/ imports.NewTable("WkWebview_LoadRequest", 0),
		/*21*/ imports.NewTable("WkWebview_LoadURL", 0),
		/*22*/ imports.NewTable("WkWebview_Magnification", 0),
		/*23*/ imports.NewTable("WkWebview_Reload", 0),
		/*24*/ imports.NewTable("WkWebview_ReloadFromOrigin", 0),
		/*25*/ imports.NewTable("WkWebview_SetAllowsBackForwardNavigationGestures", 0),
		/*26*/ imports.NewTable("WkWebview_SetAllowsMagnification", 0),
		/*27*/ imports.NewTable("WkWebview_SetMagnification", 0),
		/*28*/ imports.NewTable("WkWebview_SetMagnificationCenteredAtPoint", 0),
		/*29*/ imports.NewTable("WkWebview_SetNavigationDelegate", 0),
		/*30*/ imports.NewTable("WkWebview_SetOnCommitNavigation", 0),
		/*31*/ imports.NewTable("WkWebview_SetOnCreateWebView", 0),
		/*32*/ imports.NewTable("WkWebview_SetOnDecidePolicyForNavigationActionPreferences", 0),
		/*33*/ imports.NewTable("WkWebview_SetOnDecidePolicyForNavigationResponse", 0),
		/*34*/ imports.NewTable("WkWebview_SetOnDownloadCancelCompletionHandler", 0),
		/*35*/ imports.NewTable("WkWebview_SetOnDownloadDecideDestinationUsingResponseSuggestedFilename", 0),
		/*36*/ imports.NewTable("WkWebview_SetOnDownloadFailWithError", 0),
		/*37*/ imports.NewTable("WkWebview_SetOnDownloadFinish", 0),
		/*38*/ imports.NewTable("WkWebview_SetOnDownloadWillPerformHTTPRedirectionNewRequest", 0),
		/*39*/ imports.NewTable("WkWebview_SetOnFailNavigationWithError", 0),
		/*40*/ imports.NewTable("WkWebview_SetOnFailProvisionalNavigationWithError", 0),
		/*41*/ imports.NewTable("WkWebview_SetOnFinishNavigation", 0),
		/*42*/ imports.NewTable("WkWebview_SetOnNavigationActionDidBecomeDownload", 0),
		/*43*/ imports.NewTable("WkWebview_SetOnNavigationResponseDidBecomeDownload", 0),
		/*44*/ imports.NewTable("WkWebview_SetOnProcessMessage", 0),
		/*45*/ imports.NewTable("WkWebview_SetOnReceiveServerRedirectForProvisionalNavigation", 0),
		/*46*/ imports.NewTable("WkWebview_SetOnRunJavaScriptAlert", 0),
		/*47*/ imports.NewTable("WkWebview_SetOnRunJavaScriptConfirmCompletion", 0),
		/*48*/ imports.NewTable("WkWebview_SetOnRunJavaScriptTextInputCompletion", 0),
		/*49*/ imports.NewTable("WkWebview_SetOnStartProvisionalNavigation", 0),
		/*50*/ imports.NewTable("WkWebview_SetOnStartURLSchemeTask", 0),
		/*51*/ imports.NewTable("WkWebview_SetOnStopURLSchemeTask", 0),
		/*52*/ imports.NewTable("WkWebview_SetOnWebContentProcessDidTerminate", 0),
		/*53*/ imports.NewTable("WkWebview_SetOnWebViewDidClose", 0),
		/*54*/ imports.NewTable("WkWebview_SetUIDelegate", 0),
		/*55*/ imports.NewTable("WkWebview_StopLoading", 0),
		/*56*/ imports.NewTable("WkWebview_Title", 0),
		/*57*/ imports.NewTable("WkWebview_URL", 0),
		imports.NewTable("WkWebview_RemoveFromSuperview", 0),
		imports.NewTable("WkWebview_RemoveAllSubviews", 0),
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
