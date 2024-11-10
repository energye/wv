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
	// Data
	//  Returns the object implemented by this class.
	Data() WkWebview // function
	// Configuration
	//  The object that contains the configuration details for the web view.
	Configuration() WKWebViewConfiguration // function
	// LoadRequest
	//  Loads the web content that the specified URL request object references and navigates to that content.
	LoadRequest(request NSURLRequest) WKNavigation // function
	// Title
	//  The page title
	Title() string // function
	// URL
	//  The URL for the current webpage.
	URL() NSURL // function
	// IsLoading
	//  A Boolean value that indicates whether the view is currently loading content.
	IsLoading() bool // function
	// EstimatedProgress
	//  This value ranges from 0.0 to 1.0 based on the total number of bytes received,
	//  including the main document and all of its potential subresources.
	//  After navigation loading completes, the estimatedProgress value remains at 1.0 until a new navigation starts,
	//  at which point the estimatedProgress value resets to 0.0. The WKWebView class
	//  is key-value observing(KVO) compliant for this property.
	EstimatedProgress() (resultFloat64 float64) // function
	// HasOnlySecureContent
	//  A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
	HasOnlySecureContent() bool // function
	// CanGoBack
	//  A Boolean value that indicates whether there is a valid back item in the back-forward list.
	CanGoBack() bool // function
	// CanGoForward
	//  A Boolean value that indicates whether there is a valid forward item in the back-forward list.
	CanGoForward() bool // function
	// GoBack
	//  Navigates to the back item in the back-forward list.
	GoBack() WKNavigation // function
	// GoForward
	//  Navigates to the forward item in the back-forward list.
	GoForward() WKNavigation // function
	// Reload
	//  Reloads the current webpage.
	Reload() WKNavigation // function
	// ReloadFromOrigin
	//  Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
	ReloadFromOrigin() WKNavigation // function
	// AllowsBackForwardNavigationGestures
	//  Returns A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
	AllowsBackForwardNavigationGestures() bool // function
	// AllowsMagnification
	//  Returns A Boolean value that indicates whether magnify gestures change the web view’s magnification.
	AllowsMagnification() bool // function
	// Magnification
	//  Returns The factor by which the page content is currently scaled.
	Magnification() (resultFloat64 float64) // function
	// AsReceiveScriptMessageDelegate
	//  Must display conversion to interface type
	AsReceiveScriptMessageDelegate() IReceiveScriptMessageDelegateEvent // function
	// AsWKNavigationDelegate
	//  Must display conversion to interface type
	AsWKNavigationDelegate() IWKNavigationDelegateEvent // function
	// AsWKURLSchemeHandlerDelegate
	//  Must display conversion to interface type
	AsWKURLSchemeHandlerDelegate() IWKURLSchemeHandlerDelegateEvent // function
	// AsWKUIDelegate
	//  Must display conversion to interface type
	AsWKUIDelegate() IWKUIDelegateEvent // function
	// AsWKDownloadDelegate
	//  Must display conversion to interface type
	AsWKDownloadDelegate() IWKDownloadDelegateEvent // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// LoadHTML
	//  Loads the contents of the specified HTML string and navigates to it.
	LoadHTML(aHtml, aBaseURL string) // procedure
	// LoadURL
	//  Load the specified url
	LoadURL(aURL string) // procedure
	// SetNavigationDelegate
	//  The object you use to manage navigation behavior for the web view.
	SetNavigationDelegate(navigationDelegate WKNavigationDelegate) // procedure
	// SetUIDelegate
	//  The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
	SetUIDelegate(uiDelegate WKUIDelegate) // procedure
	// InitWithFrameConfiguration
	//  Creates a web view and initializes it with the specified frame and configuration data.
	InitWithFrameConfiguration(aFrame *TRect, aConfiguration WKWebViewConfiguration) // procedure
	// StopLoading
	//  Stops loading all resources on the current page.
	StopLoading() // procedure
	// EvaluateJavaScript
	//  Evaluates the specified JavaScript string.
	EvaluateJavaScript(javaScriptString string) // procedure
	// SetAllowsBackForwardNavigationGestures
	//  Sets A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
	SetAllowsBackForwardNavigationGestures(newValue bool) // procedure
	// SetAllowsMagnification
	//  Sets A Boolean value that indicates whether magnify gestures change the web view’s magnification.
	SetAllowsMagnification(newValue bool) // procedure
	// SetMagnification
	//  Sets The factor by which the page content is currently scaled.
	SetMagnification(newValue float64) // procedure
	// SetMagnificationCenteredAtPoint
	//  Scales the page content and centers the result on the specified point.
	SetMagnificationCenteredAtPoint(magnification float64, point *TPoint) // procedure
	// RemoveFromSuperview
	//  Called when closing the window properly closes the webview
	RemoveFromSuperview() // procedure
	// SetOnProcessMessage
	//  WKScriptMessageHandlerProtocol
	//  Tells the handler that a webpage sent a script message.
	SetOnProcessMessage(fn TWkProcessMessageEvent) // property event
	// SetOnDecidePolicyForNavigationActionPreferences
	//  WKNavigationDelegateProtocol
	//  Asks the delegate for permission to navigate to new content based on the specified preferences and action information.
	//  Use this method to allow or deny a navigation request that originated with the specified action.
	//  The web view calls this method after the interaction occurs but before it attempts to load any content.
	SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences) // property event
	// SetOnDecidePolicyForNavigationResponse
	//  Asks the delegate for permission to navigate to new content after the response to the navigation request is known.
	//  Use this method to allow or deny a navigation request after the web view receives the response to its original URL request.
	//  The navigationResponse parameter contains the details of the response, including the type of data that the response contains.
	SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse) // property event
	// SetOnStartProvisionalNavigation
	//  Tells the delegate that navigation from the main frame has started.
	//  The web view calls this method after it receives provisional approval to process a navigation request, but before it receives a response to that request.
	SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation) // property event
	// SetOnReceiveServerRedirectForProvisionalNavigation
	//  Tells the delegate that the web view received a server redirect for a request.
	SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation) // property event
	// SetOnFailProvisionalNavigationWithError
	//  Tells the delegate that an error occurred during the early navigation process.
	SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError) // property event
	// SetOnCommitNavigation
	//  Tells the delegate that the web view has started to receive content for the main frame.
	SetOnCommitNavigation(fn TWkCommitNavigation) // property event
	// SetOnFinishNavigation
	//  Tells the delegate that navigation is complete.
	SetOnFinishNavigation(fn TWkFinishNavigation) // property event
	// SetOnFailNavigationWithError
	//  Tells the delegate that an error occurred during navigation.
	SetOnFailNavigationWithError(fn TWkFailNavigationWithError) // property event
	// SetOnWebContentProcessDidTerminate
	//  Asks the delegate to respond to an authentication challenge.
	//  property OnReceiveAuthenticationChallenge: TWkReceiveAuthenticationChallenge read FOnReceiveAuthenticationChallenge write FOnReceiveAuthenticationChallenge;
	//  Tells the delegate that the web view’s content process was terminated.
	SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate) // property event
	// SetOnNavigationResponseDidBecomeDownload
	//  Asks the delegate whether to continue with a connection that uses a deprecated version of TLS.
	//  property OnAuthenticationChallengeShouldAllowDeprecatedTLS: TWkAuthenticationChallengeShouldAllowDeprecatedTLS read FOnAuthenticationChallengeShouldAllowDeprecatedTLS write FOnAuthenticationChallengeShouldAllowDeprecatedTLS; Tells the delegate that a navigation action became a download. property OnNavigationActionDidBecomeDownload: TWkNavigationActionDidBecomeDownload read FOnNavigationActionDidBecomeDownload write FOnNavigationActionDidBecomeDownload;
	//  Tells the delegate that a navigation response became a download.
	SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload) // property event
	// SetOnStartURLSchemeTask
	//  WKURLSchemeHandlerProtocol
	//  Asks your handler to begin loading the data for the specified resource.
	SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask) // property event
	// SetOnStopURLSchemeTask
	//  Asks your handler to stop loading the data for the specified resource.
	SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask) // property event
	// SetOnCreateWebView
	//  WKUIDelegateProtocol
	//  Creates a new web view.
	SetOnCreateWebView(fn TWKCreateWebView) // property event
	// SetOnRunJavaScriptAlert
	//  Displays a JavaScript alert panel.
	SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert) // property event
	// SetOnRunJavaScriptConfirmCompletion
	//  Displays a JavaScript confirm panel.
	SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion) // property event
	// SetOnRunJavaScriptTextInputCompletion
	//  Displays a JavaScript text input panel.
	SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion) // property event
	// SetOnWebViewDidClose
	//  Notifies your app that the DOM window closed successfully.
	SetOnWebViewDidClose(fn TWKWebViewDidClose) // property event
	// SetOnDownloadCancelCompletionHandler
	//  WKDownload
	SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler) // property event
	// SetOnDownloadDecideDestinationUsingResponseSuggestedFilename
	//  Asks the delegate to provide a file destination where the system should write the download data.
	SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) // property event
	// SetOnDownloadWillPerformHTTPRedirectionNewRequest
	//  Asks the delegate to respond to the download’s redirect response.
	SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest) // property event
	// SetOnDownloadFinish
	//  Asks the delegate to respond to an authentication challenge.
	//  property OnDownloadReceiveAuthenticationChallenge: TWKDownloadReceiveAuthenticationChallenge read FOnDownloadReceiveAuthenticationChallenge write FOnDownloadReceiveAuthenticationChallenge;
	//  Tells the delegate that the download finished.
	SetOnDownloadFinish(fn TWKDownloadFinish) // property event
	// SetOnDownloadFailWithError
	//  Tells the delegate that the download failed, with error information and data you can use to restart the download.
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
	r1 := wkWebviewImportAPI().SysCallN(57, m.Instance())
	return GoStr(r1)
}

func (m *TWkWebview) URL() NSURL {
	r1 := wkWebviewImportAPI().SysCallN(58, m.Instance())
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
	r1 := wkWebviewImportAPI().SysCallN(24, m.Instance())
	return WKNavigation(r1)
}

func (m *TWkWebview) ReloadFromOrigin() WKNavigation {
	r1 := wkWebviewImportAPI().SysCallN(25, m.Instance())
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

func (m *TWkWebview) Release() {
	wkWebviewImportAPI().SysCallN(23, m.Instance())
}

func (m *TWkWebview) LoadHTML(aHtml, aBaseURL string) {
	wkWebviewImportAPI().SysCallN(19, m.Instance(), PascalStr(aHtml), PascalStr(aBaseURL))
}

func (m *TWkWebview) LoadURL(aURL string) {
	wkWebviewImportAPI().SysCallN(21, m.Instance(), PascalStr(aURL))
}

func (m *TWkWebview) SetNavigationDelegate(navigationDelegate WKNavigationDelegate) {
	wkWebviewImportAPI().SysCallN(31, m.Instance(), uintptr(navigationDelegate))
}

func (m *TWkWebview) SetUIDelegate(uiDelegate WKUIDelegate) {
	wkWebviewImportAPI().SysCallN(55, m.Instance(), uintptr(uiDelegate))
}

func (m *TWkWebview) InitWithFrameConfiguration(aFrame *TRect, aConfiguration WKWebViewConfiguration) {
	wkWebviewImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(aFrame)), uintptr(aConfiguration))
}

func (m *TWkWebview) StopLoading() {
	wkWebviewImportAPI().SysCallN(56, m.Instance())
}

func (m *TWkWebview) EvaluateJavaScript(javaScriptString string) {
	wkWebviewImportAPI().SysCallN(13, m.Instance(), PascalStr(javaScriptString))
}

func (m *TWkWebview) SetAllowsBackForwardNavigationGestures(newValue bool) {
	wkWebviewImportAPI().SysCallN(27, m.Instance(), PascalBool(newValue))
}

func (m *TWkWebview) SetAllowsMagnification(newValue bool) {
	wkWebviewImportAPI().SysCallN(28, m.Instance(), PascalBool(newValue))
}

func (m *TWkWebview) SetMagnification(newValue float64) {
	wkWebviewImportAPI().SysCallN(29, m.Instance(), uintptr(unsafePointer(&newValue)))
}

func (m *TWkWebview) SetMagnificationCenteredAtPoint(magnification float64, point *TPoint) {
	wkWebviewImportAPI().SysCallN(30, m.Instance(), uintptr(unsafePointer(&magnification)), uintptr(unsafePointer(point)))
}

func (m *TWkWebview) RemoveFromSuperview() {
	wkWebviewImportAPI().SysCallN(26, m.Instance())
}

func (m *TWkWebview) SetOnProcessMessage(fn TWkProcessMessageEvent) {
	if m.processMessagePtr != 0 {
		RemoveEventElement(m.processMessagePtr)
	}
	m.processMessagePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(45, m.Instance(), m.processMessagePtr)
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences) {
	if m.decidePolicyForNavigationActionPreferencesPtr != 0 {
		RemoveEventElement(m.decidePolicyForNavigationActionPreferencesPtr)
	}
	m.decidePolicyForNavigationActionPreferencesPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(34, m.Instance(), m.decidePolicyForNavigationActionPreferencesPtr)
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse) {
	if m.decidePolicyForNavigationResponsePtr != 0 {
		RemoveEventElement(m.decidePolicyForNavigationResponsePtr)
	}
	m.decidePolicyForNavigationResponsePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(35, m.Instance(), m.decidePolicyForNavigationResponsePtr)
}

func (m *TWkWebview) SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation) {
	if m.startProvisionalNavigationPtr != 0 {
		RemoveEventElement(m.startProvisionalNavigationPtr)
	}
	m.startProvisionalNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(50, m.Instance(), m.startProvisionalNavigationPtr)
}

func (m *TWkWebview) SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation) {
	if m.receiveServerRedirectForProvisionalNavigationPtr != 0 {
		RemoveEventElement(m.receiveServerRedirectForProvisionalNavigationPtr)
	}
	m.receiveServerRedirectForProvisionalNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(46, m.Instance(), m.receiveServerRedirectForProvisionalNavigationPtr)
}

func (m *TWkWebview) SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError) {
	if m.failProvisionalNavigationWithErrorPtr != 0 {
		RemoveEventElement(m.failProvisionalNavigationWithErrorPtr)
	}
	m.failProvisionalNavigationWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(42, m.Instance(), m.failProvisionalNavigationWithErrorPtr)
}

func (m *TWkWebview) SetOnCommitNavigation(fn TWkCommitNavigation) {
	if m.commitNavigationPtr != 0 {
		RemoveEventElement(m.commitNavigationPtr)
	}
	m.commitNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(32, m.Instance(), m.commitNavigationPtr)
}

func (m *TWkWebview) SetOnFinishNavigation(fn TWkFinishNavigation) {
	if m.finishNavigationPtr != 0 {
		RemoveEventElement(m.finishNavigationPtr)
	}
	m.finishNavigationPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(43, m.Instance(), m.finishNavigationPtr)
}

func (m *TWkWebview) SetOnFailNavigationWithError(fn TWkFailNavigationWithError) {
	if m.failNavigationWithErrorPtr != 0 {
		RemoveEventElement(m.failNavigationWithErrorPtr)
	}
	m.failNavigationWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(41, m.Instance(), m.failNavigationWithErrorPtr)
}

func (m *TWkWebview) SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate) {
	if m.webContentProcessDidTerminatePtr != 0 {
		RemoveEventElement(m.webContentProcessDidTerminatePtr)
	}
	m.webContentProcessDidTerminatePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(53, m.Instance(), m.webContentProcessDidTerminatePtr)
}

func (m *TWkWebview) SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload) {
	if m.navigationResponseDidBecomeDownloadPtr != 0 {
		RemoveEventElement(m.navigationResponseDidBecomeDownloadPtr)
	}
	m.navigationResponseDidBecomeDownloadPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(44, m.Instance(), m.navigationResponseDidBecomeDownloadPtr)
}

func (m *TWkWebview) SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask) {
	if m.startURLSchemeTaskPtr != 0 {
		RemoveEventElement(m.startURLSchemeTaskPtr)
	}
	m.startURLSchemeTaskPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(51, m.Instance(), m.startURLSchemeTaskPtr)
}

func (m *TWkWebview) SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask) {
	if m.stopURLSchemeTaskPtr != 0 {
		RemoveEventElement(m.stopURLSchemeTaskPtr)
	}
	m.stopURLSchemeTaskPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(52, m.Instance(), m.stopURLSchemeTaskPtr)
}

func (m *TWkWebview) SetOnCreateWebView(fn TWKCreateWebView) {
	if m.createWebViewPtr != 0 {
		RemoveEventElement(m.createWebViewPtr)
	}
	m.createWebViewPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(33, m.Instance(), m.createWebViewPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert) {
	if m.runJavaScriptAlertPtr != 0 {
		RemoveEventElement(m.runJavaScriptAlertPtr)
	}
	m.runJavaScriptAlertPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(47, m.Instance(), m.runJavaScriptAlertPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion) {
	if m.runJavaScriptConfirmCompletionPtr != 0 {
		RemoveEventElement(m.runJavaScriptConfirmCompletionPtr)
	}
	m.runJavaScriptConfirmCompletionPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(48, m.Instance(), m.runJavaScriptConfirmCompletionPtr)
}

func (m *TWkWebview) SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion) {
	if m.runJavaScriptTextInputCompletionPtr != 0 {
		RemoveEventElement(m.runJavaScriptTextInputCompletionPtr)
	}
	m.runJavaScriptTextInputCompletionPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(49, m.Instance(), m.runJavaScriptTextInputCompletionPtr)
}

func (m *TWkWebview) SetOnWebViewDidClose(fn TWKWebViewDidClose) {
	if m.webViewDidClosePtr != 0 {
		RemoveEventElement(m.webViewDidClosePtr)
	}
	m.webViewDidClosePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(54, m.Instance(), m.webViewDidClosePtr)
}

func (m *TWkWebview) SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler) {
	if m.downloadCancelCompletionHandlerPtr != 0 {
		RemoveEventElement(m.downloadCancelCompletionHandlerPtr)
	}
	m.downloadCancelCompletionHandlerPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(36, m.Instance(), m.downloadCancelCompletionHandlerPtr)
}

func (m *TWkWebview) SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) {
	if m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr != 0 {
		RemoveEventElement(m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr)
	}
	m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(37, m.Instance(), m.downloadDecideDestinationUsingResponseSuggestedFilenamePtr)
}

func (m *TWkWebview) SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest) {
	if m.downloadWillPerformHTTPRedirectionNewRequestPtr != 0 {
		RemoveEventElement(m.downloadWillPerformHTTPRedirectionNewRequestPtr)
	}
	m.downloadWillPerformHTTPRedirectionNewRequestPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(40, m.Instance(), m.downloadWillPerformHTTPRedirectionNewRequestPtr)
}

func (m *TWkWebview) SetOnDownloadFinish(fn TWKDownloadFinish) {
	if m.downloadFinishPtr != 0 {
		RemoveEventElement(m.downloadFinishPtr)
	}
	m.downloadFinishPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(39, m.Instance(), m.downloadFinishPtr)
}

func (m *TWkWebview) SetOnDownloadFailWithError(fn TWKDownloadFailWithError) {
	if m.downloadFailWithErrorPtr != 0 {
		RemoveEventElement(m.downloadFailWithErrorPtr)
	}
	m.downloadFailWithErrorPtr = MakeEventDataPtr(fn)
	wkWebviewImportAPI().SysCallN(38, m.Instance(), m.downloadFailWithErrorPtr)
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
		/*23*/ imports.NewTable("WkWebview_Release", 0),
		/*24*/ imports.NewTable("WkWebview_Reload", 0),
		/*25*/ imports.NewTable("WkWebview_ReloadFromOrigin", 0),
		/*26*/ imports.NewTable("WkWebview_RemoveFromSuperview", 0),
		/*27*/ imports.NewTable("WkWebview_SetAllowsBackForwardNavigationGestures", 0),
		/*28*/ imports.NewTable("WkWebview_SetAllowsMagnification", 0),
		/*29*/ imports.NewTable("WkWebview_SetMagnification", 0),
		/*30*/ imports.NewTable("WkWebview_SetMagnificationCenteredAtPoint", 0),
		/*31*/ imports.NewTable("WkWebview_SetNavigationDelegate", 0),
		/*32*/ imports.NewTable("WkWebview_SetOnCommitNavigation", 0),
		/*33*/ imports.NewTable("WkWebview_SetOnCreateWebView", 0),
		/*34*/ imports.NewTable("WkWebview_SetOnDecidePolicyForNavigationActionPreferences", 0),
		/*35*/ imports.NewTable("WkWebview_SetOnDecidePolicyForNavigationResponse", 0),
		/*36*/ imports.NewTable("WkWebview_SetOnDownloadCancelCompletionHandler", 0),
		/*37*/ imports.NewTable("WkWebview_SetOnDownloadDecideDestinationUsingResponseSuggestedFilename", 0),
		/*38*/ imports.NewTable("WkWebview_SetOnDownloadFailWithError", 0),
		/*39*/ imports.NewTable("WkWebview_SetOnDownloadFinish", 0),
		/*40*/ imports.NewTable("WkWebview_SetOnDownloadWillPerformHTTPRedirectionNewRequest", 0),
		/*41*/ imports.NewTable("WkWebview_SetOnFailNavigationWithError", 0),
		/*42*/ imports.NewTable("WkWebview_SetOnFailProvisionalNavigationWithError", 0),
		/*43*/ imports.NewTable("WkWebview_SetOnFinishNavigation", 0),
		/*44*/ imports.NewTable("WkWebview_SetOnNavigationResponseDidBecomeDownload", 0),
		/*45*/ imports.NewTable("WkWebview_SetOnProcessMessage", 0),
		/*46*/ imports.NewTable("WkWebview_SetOnReceiveServerRedirectForProvisionalNavigation", 0),
		/*47*/ imports.NewTable("WkWebview_SetOnRunJavaScriptAlert", 0),
		/*48*/ imports.NewTable("WkWebview_SetOnRunJavaScriptConfirmCompletion", 0),
		/*49*/ imports.NewTable("WkWebview_SetOnRunJavaScriptTextInputCompletion", 0),
		/*50*/ imports.NewTable("WkWebview_SetOnStartProvisionalNavigation", 0),
		/*51*/ imports.NewTable("WkWebview_SetOnStartURLSchemeTask", 0),
		/*52*/ imports.NewTable("WkWebview_SetOnStopURLSchemeTask", 0),
		/*53*/ imports.NewTable("WkWebview_SetOnWebContentProcessDidTerminate", 0),
		/*54*/ imports.NewTable("WkWebview_SetOnWebViewDidClose", 0),
		/*55*/ imports.NewTable("WkWebview_SetUIDelegate", 0),
		/*56*/ imports.NewTable("WkWebview_StopLoading", 0),
		/*57*/ imports.NewTable("WkWebview_Title", 0),
		/*58*/ imports.NewTable("WkWebview_URL", 0),
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
