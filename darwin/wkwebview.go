//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkWebview Parent: lcl.IComponent
type IWkWebview interface {
	lcl.IComponent
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKWebView // function
	// Configuration
	//  The object that contains the configuration details for the web view.
	Configuration() wvTypes.WKWebViewConfiguration // function
	// LoadRequest
	//  Loads the web content that the specified URL request object references and navigates to that content.
	LoadRequest(request NSURLRequest) wvTypes.WKNavigation // function
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
	//  is key-value observing (KVO) compliant for this property.
	EstimatedProgress() float64 // function
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
	GoBack() wvTypes.WKNavigation // function
	// GoForward
	//  Navigates to the forward item in the back-forward list.
	GoForward() wvTypes.WKNavigation // function
	// Reload
	//  Reloads the current webpage.
	Reload() wvTypes.WKNavigation // function
	// ReloadFromOrigin
	//  Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
	ReloadFromOrigin() wvTypes.WKNavigation // function
	// AllowsBackForwardNavigationGestures
	//  Returns A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
	AllowsBackForwardNavigationGestures() bool // function
	// AllowsMagnification
	//  Returns A Boolean value that indicates whether magnify gestures change the web view’s magnification.
	AllowsMagnification() bool // function
	// Magnification
	//  Returns The factor by which the page content is currently scaled.
	Magnification() float64 // function
	// AsReceiveScriptMessageDelegate
	//  Must display conversion to interface type
	AsReceiveScriptMessageDelegate() IWkWebview // function
	// AsWKNavigationDelegate
	//  Must display conversion to interface type
	AsWKNavigationDelegate() IWkWebview // function
	// AsWKURLSchemeHandlerDelegate
	//  Must display conversion to interface type
	AsWKURLSchemeHandlerDelegate() IWkWebview // function
	// AsWKUIDelegate
	//  Must display conversion to interface type
	AsWKUIDelegate() IWkWebview // function
	// AsWKDownloadDelegate
	//  Must display conversion to interface type
	AsWKDownloadDelegate() IWkWebview // function
	// Release
	//  Freeing the class and the objects it implements.
	Release()                          // procedure
	SetAcceptsMouseMovedEvents(V bool) // procedure
	// LoadHTML
	//  Loads the contents of the specified HTML string and navigates to it.
	LoadHTML(html string, baseURL string) // procedure
	// LoadURL
	//  Load the specified url
	LoadURL(uRL string) // procedure
	// SetNavigationDelegate
	//  The object you use to manage navigation behavior for the web view.
	SetNavigationDelegate(navigationDelegate wvTypes.WKNavigationDelegate) // procedure
	// SetUIDelegate
	//  The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
	SetUIDelegate(uiDelegate wvTypes.WKUIDelegate) // procedure
	// InitWithFrameConfiguration
	//  Creates a web view and initializes it with the specified frame and configuration data.
	InitWithFrameConfiguration(frame types.TRect, configuration wvTypes.WKWebViewConfiguration) // procedure
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
	SetMagnificationCenteredAtPoint(magnification float64, point types.TPoint) // procedure
	// RemoveFromSuperview
	//  Called when closing the window properly closes the webview
	RemoveFromSuperview()                                                                                                       // procedure
	SetOnProcessMessage(fn TWkProcessMessageEvent)                                                                              // property event
	SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences)                           // property event
	SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse)                                             // property event
	SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation)                                                           // property event
	SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation)                     // property event
	SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError)                                           // property event
	SetOnCommitNavigation(fn TWkCommitNavigation)                                                                               // property event
	SetOnFinishNavigation(fn TWkFinishNavigation)                                                                               // property event
	SetOnFailNavigationWithError(fn TWkFailNavigationWithError)                                                                 // property event
	SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate)                                                     // property event
	SetOnNavigationActionDidBecomeDownload(fn TWkNavigationActionDidBecomeDownload)                                             // property event
	SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload)                                         // property event
	SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask)                                                                           // property event
	SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask)                                                                             // property event
	SetOnCreateWebView(fn TWKCreateWebView)                                                                                     // property event
	SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert)                                                                           // property event
	SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion)                                                   // property event
	SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion)                                               // property event
	SetOnWebViewDidClose(fn TWKWebViewDidClose)                                                                                 // property event
	SetOnMouseDown(fn TWkMouseEvent)                                                                                            // property event
	SetOnMouseUp(fn TWkMouseEvent)                                                                                              // property event
	SetOnMouseMoved(fn TWkMouseEvent)                                                                                           // property event
	SetOnMouseDragged(fn TWkMouseEvent)                                                                                         // property event
	SetOnRightMouseDown(fn TWkMouseEvent)                                                                                       // property event
	SetOnRightMouseUp(fn TWkMouseEvent)                                                                                         // property event
	SetOnRightMouseDragged(fn TWkMouseEvent)                                                                                    // property event
	SetOnOtherMouseDown(fn TWkMouseEvent)                                                                                       // property event
	SetOnOtherMouseUp(fn TWkMouseEvent)                                                                                         // property event
	SetOnOtherMouseDragged(fn TWkMouseEvent)                                                                                    // property event
	SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler)                                                 // property event
	SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) // property event
	SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest)                       // property event
	SetOnDownloadFinish(fn TWKDownloadFinish)                                                                                   // property event
	SetOnDownloadFailWithError(fn TWKDownloadFailWithError)                                                                     // property event
	SetOnDraggingEntered(fn TWKDraggingEntered)                                                                                 // property event
	SetOnDraggingUpdated(fn TWKDraggingUpdated)                                                                                 // property event
	SetOnDraggingExited(fn TWKDraggingExited)                                                                                   // property event
	SetOnPrepareForDragOperation(fn TWKPrepareForDragOperation)                                                                 // property event
	SetOnPerformDragOperation(fn TWKPerformDragOperation)                                                                       // property event
	SetOnConcludeDragOperation(fn TWKConcludeDragOperation)                                                                     // property event
	SetOnDraggingEnded(fn TWKDraggingEnded)                                                                                     // property event
	SetOnWantsPeriodicDraggingUpdates(fn TWKWantsPeriodicDraggingUpdates)                                                       // property event
	SetOnUpdateDraggingItemsForDrag(fn TWKUpdateDraggingItemsForDrag)                                                           // property event
	AsIntfReceiveScriptMessageDelegateEvent() uintptr
	AsIntfWKNavigationDelegateEvent() uintptr
	AsIntfWKURLSchemeHandlerDelegateEvent() uintptr
	AsIntfWKUIDelegateEvent() uintptr
	AsIntfWKDownloadDelegateEvent() uintptr
}

type TWkWebview struct {
	lcl.TComponent
}

func (m *TWkWebview) Data() wvTypes.WKWebView {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(1, m.Instance())
	return wvTypes.WKWebView(r)
}

func (m *TWkWebview) Configuration() wvTypes.WKWebViewConfiguration {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(2, m.Instance())
	return wvTypes.WKWebViewConfiguration(r)
}

func (m *TWkWebview) LoadRequest(request NSURLRequest) wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(3, m.Instance(), uintptr(request))
	return wvTypes.WKNavigation(r)
}

func (m *TWkWebview) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := wkWebviewAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TWkWebview) URL() NSURL {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(5, m.Instance())
	return NSURL(r)
}

func (m *TWkWebview) IsLoading() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) EstimatedProgress() (result float64) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkWebview) HasOnlySecureContent() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) GoBack() wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(11, m.Instance())
	return wvTypes.WKNavigation(r)
}

func (m *TWkWebview) GoForward() wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(12, m.Instance())
	return wvTypes.WKNavigation(r)
}

func (m *TWkWebview) Reload() wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(13, m.Instance())
	return wvTypes.WKNavigation(r)
}

func (m *TWkWebview) ReloadFromOrigin() wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewAPI().SysCallN(14, m.Instance())
	return wvTypes.WKNavigation(r)
}

func (m *TWkWebview) AllowsBackForwardNavigationGestures() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) AllowsMagnification() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebviewAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebview) Magnification() (result float64) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkWebview) AsReceiveScriptMessageDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) AsWKNavigationDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) AsWKURLSchemeHandlerDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) AsWKUIDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) AsWKDownloadDelegate() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	wkWebviewAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return AsWkWebview(resultPtr)
}

func (m *TWkWebview) Release() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(23, m.Instance())
}

func (m *TWkWebview) SetAcceptsMouseMovedEvents(V bool) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(24, m.Instance(), api.PasBool(V))
}

func (m *TWkWebview) LoadHTML(html string, baseURL string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(25, m.Instance(), api.PasStr(html), api.PasStr(baseURL))
}

func (m *TWkWebview) LoadURL(uRL string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(26, m.Instance(), api.PasStr(uRL))
}

func (m *TWkWebview) SetNavigationDelegate(navigationDelegate wvTypes.WKNavigationDelegate) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(27, m.Instance(), uintptr(navigationDelegate))
}

func (m *TWkWebview) SetUIDelegate(uiDelegate wvTypes.WKUIDelegate) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(28, m.Instance(), uintptr(uiDelegate))
}

func (m *TWkWebview) InitWithFrameConfiguration(frame types.TRect, configuration wvTypes.WKWebViewConfiguration) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&frame)), uintptr(configuration))
}

func (m *TWkWebview) StopLoading() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(30, m.Instance())
}

func (m *TWkWebview) EvaluateJavaScript(javaScriptString string) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(31, m.Instance(), api.PasStr(javaScriptString))
}

func (m *TWkWebview) SetAllowsBackForwardNavigationGestures(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(32, m.Instance(), api.PasBool(newValue))
}

func (m *TWkWebview) SetAllowsMagnification(newValue bool) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(33, m.Instance(), api.PasBool(newValue))
}

func (m *TWkWebview) SetMagnification(newValue float64) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(34, m.Instance(), uintptr(base.UnsafePointer(&newValue)))
}

func (m *TWkWebview) SetMagnificationCenteredAtPoint(magnification float64, point types.TPoint) {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(35, m.Instance(), uintptr(base.UnsafePointer(&magnification)), uintptr(base.UnsafePointer(&point)))
}

func (m *TWkWebview) RemoveFromSuperview() {
	if !m.IsValid() {
		return
	}
	wkWebviewAPI().SysCallN(36, m.Instance())
}

func (m *TWkWebview) SetOnProcessMessage(fn TWkProcessMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkProcessMessageEvent(fn)
	base.SetEvent(m, 37, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationActionPreferences(fn TWKDecidePolicyForNavigationActionPreferences) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDecidePolicyForNavigationActionPreferences(fn)
	base.SetEvent(m, 38, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDecidePolicyForNavigationResponse(fn TWKDecidePolicyForNavigationResponse) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDecidePolicyForNavigationResponse(fn)
	base.SetEvent(m, 39, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnStartProvisionalNavigation(fn TWkStartProvisionalNavigation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkStartProvisionalNavigation(fn)
	base.SetEvent(m, 40, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnReceiveServerRedirectForProvisionalNavigation(fn TWkReceiveServerRedirectForProvisionalNavigation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkReceiveServerRedirectForProvisionalNavigation(fn)
	base.SetEvent(m, 41, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnFailProvisionalNavigationWithError(fn TWkFailProvisionalNavigationWithError) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkFailProvisionalNavigationWithError(fn)
	base.SetEvent(m, 42, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnCommitNavigation(fn TWkCommitNavigation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkCommitNavigation(fn)
	base.SetEvent(m, 43, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnFinishNavigation(fn TWkFinishNavigation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkFinishNavigation(fn)
	base.SetEvent(m, 44, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnFailNavigationWithError(fn TWkFailNavigationWithError) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkFailNavigationWithError(fn)
	base.SetEvent(m, 45, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnWebContentProcessDidTerminate(fn TWkWebContentProcessDidTerminate) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkWebContentProcessDidTerminate(fn)
	base.SetEvent(m, 46, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnNavigationActionDidBecomeDownload(fn TWkNavigationActionDidBecomeDownload) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkNavigationActionDidBecomeDownload(fn)
	base.SetEvent(m, 47, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnNavigationResponseDidBecomeDownload(fn TWkNavigationResponseDidBecomeDownload) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkNavigationResponseDidBecomeDownload(fn)
	base.SetEvent(m, 48, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnStartURLSchemeTask(fn TWKStartURLSchemeTask) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKStartURLSchemeTask(fn)
	base.SetEvent(m, 49, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnStopURLSchemeTask(fn TWKStopURLSchemeTask) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKStopURLSchemeTask(fn)
	base.SetEvent(m, 50, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnCreateWebView(fn TWKCreateWebView) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKCreateWebView(fn)
	base.SetEvent(m, 51, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRunJavaScriptAlert(fn TWKRunJavaScriptAlert) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKRunJavaScriptAlert(fn)
	base.SetEvent(m, 52, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRunJavaScriptConfirmCompletion(fn TWKRunJavaScriptConfirmCompletion) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKRunJavaScriptConfirmCompletion(fn)
	base.SetEvent(m, 53, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRunJavaScriptTextInputCompletion(fn TWKRunJavaScriptTextInputCompletion) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKRunJavaScriptTextInputCompletion(fn)
	base.SetEvent(m, 54, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnWebViewDidClose(fn TWKWebViewDidClose) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKWebViewDidClose(fn)
	base.SetEvent(m, 55, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseDown(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 56, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseUp(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 57, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseMoved(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 58, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnMouseDragged(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 59, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRightMouseDown(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 60, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRightMouseUp(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 61, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnRightMouseDragged(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 62, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnOtherMouseDown(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 63, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnOtherMouseUp(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 64, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnOtherMouseDragged(fn TWkMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTWkMouseEvent(fn)
	base.SetEvent(m, 65, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDownloadCancelCompletionHandler(fn TWKDownloadCancelCompletionHandler) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDownloadCancelCompletionHandler(fn)
	base.SetEvent(m, 66, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDownloadDecideDestinationUsingResponseSuggestedFilename(fn TWKDownloadDecideDestinationUsingResponseSuggestedFilename) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDownloadDecideDestinationUsingResponseSuggestedFilename(fn)
	base.SetEvent(m, 67, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDownloadWillPerformHTTPRedirectionNewRequest(fn TWKDownloadWillPerformHTTPRedirectionNewRequest) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDownloadWillPerformHTTPRedirectionNewRequest(fn)
	base.SetEvent(m, 68, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDownloadFinish(fn TWKDownloadFinish) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDownloadFinish(fn)
	base.SetEvent(m, 69, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDownloadFailWithError(fn TWKDownloadFailWithError) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDownloadFailWithError(fn)
	base.SetEvent(m, 70, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDraggingEntered(fn TWKDraggingEntered) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDraggingEntered(fn)
	base.SetEvent(m, 71, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDraggingUpdated(fn TWKDraggingUpdated) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDraggingUpdated(fn)
	base.SetEvent(m, 72, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDraggingExited(fn TWKDraggingExited) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDraggingExited(fn)
	base.SetEvent(m, 73, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnPrepareForDragOperation(fn TWKPrepareForDragOperation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKPrepareForDragOperation(fn)
	base.SetEvent(m, 74, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnPerformDragOperation(fn TWKPerformDragOperation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKPerformDragOperation(fn)
	base.SetEvent(m, 75, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnConcludeDragOperation(fn TWKConcludeDragOperation) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKConcludeDragOperation(fn)
	base.SetEvent(m, 76, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnDraggingEnded(fn TWKDraggingEnded) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKDraggingEnded(fn)
	base.SetEvent(m, 77, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnWantsPeriodicDraggingUpdates(fn TWKWantsPeriodicDraggingUpdates) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKWantsPeriodicDraggingUpdates(fn)
	base.SetEvent(m, 78, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) SetOnUpdateDraggingItemsForDrag(fn TWKUpdateDraggingItemsForDrag) {
	if !m.IsValid() {
		return
	}
	cb := makeTWKUpdateDraggingItemsForDrag(fn)
	base.SetEvent(m, 79, wkWebviewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWkWebview) AsIntfReceiveScriptMessageDelegateEvent() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TWkWebview) AsIntfWKNavigationDelegateEvent() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TWkWebview) AsIntfWKURLSchemeHandlerDelegateEvent() uintptr {
	return m.GetIntfPointer(2)
}
func (m *TWkWebview) AsIntfWKUIDelegateEvent() uintptr {
	return m.GetIntfPointer(3)
}
func (m *TWkWebview) AsIntfWKDownloadDelegateEvent() uintptr {
	return m.GetIntfPointer(4)
}

// NewWebview class constructor
func NewWebview(owner lcl.IComponent) IWkWebview {
	var receiveScriptMessageDelegateEventPtr uintptr // IReceiveScriptMessageDelegateEvent
	var wKNavigationDelegateEventPtr uintptr         // IWKNavigationDelegateEvent
	var wKURLSchemeHandlerDelegateEventPtr uintptr   // IWKURLSchemeHandlerDelegateEvent
	var wKUIDelegateEventPtr uintptr                 // IWKUIDelegateEvent
	var wKDownloadDelegateEventPtr uintptr           // IWKDownloadDelegateEvent
	r := wkWebviewAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&receiveScriptMessageDelegateEventPtr)), uintptr(base.UnsafePointer(&wKNavigationDelegateEventPtr)), uintptr(base.UnsafePointer(&wKURLSchemeHandlerDelegateEventPtr)), uintptr(base.UnsafePointer(&wKUIDelegateEventPtr)), uintptr(base.UnsafePointer(&wKDownloadDelegateEventPtr)))
	ret := AsWkWebview(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(5)
		intf.SetIntfPointer(0, receiveScriptMessageDelegateEventPtr)
		intf.SetIntfPointer(1, wKNavigationDelegateEventPtr)
		intf.SetIntfPointer(2, wKURLSchemeHandlerDelegateEventPtr)
		intf.SetIntfPointer(3, wKUIDelegateEventPtr)
		intf.SetIntfPointer(4, wKDownloadDelegateEventPtr)
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
			/* 1 */ imports.NewTable("TWkWebview_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkWebview_Configuration", 0), // function Configuration
			/* 3 */ imports.NewTable("TWkWebview_LoadRequest", 0), // function LoadRequest
			/* 4 */ imports.NewTable("TWkWebview_Title", 0), // function Title
			/* 5 */ imports.NewTable("TWkWebview_URL", 0), // function URL
			/* 6 */ imports.NewTable("TWkWebview_IsLoading", 0), // function IsLoading
			/* 7 */ imports.NewTable("TWkWebview_EstimatedProgress", 0), // function EstimatedProgress
			/* 8 */ imports.NewTable("TWkWebview_HasOnlySecureContent", 0), // function HasOnlySecureContent
			/* 9 */ imports.NewTable("TWkWebview_CanGoBack", 0), // function CanGoBack
			/* 10 */ imports.NewTable("TWkWebview_CanGoForward", 0), // function CanGoForward
			/* 11 */ imports.NewTable("TWkWebview_GoBack", 0), // function GoBack
			/* 12 */ imports.NewTable("TWkWebview_GoForward", 0), // function GoForward
			/* 13 */ imports.NewTable("TWkWebview_Reload", 0), // function Reload
			/* 14 */ imports.NewTable("TWkWebview_ReloadFromOrigin", 0), // function ReloadFromOrigin
			/* 15 */ imports.NewTable("TWkWebview_AllowsBackForwardNavigationGestures", 0), // function AllowsBackForwardNavigationGestures
			/* 16 */ imports.NewTable("TWkWebview_AllowsMagnification", 0), // function AllowsMagnification
			/* 17 */ imports.NewTable("TWkWebview_Magnification", 0), // function Magnification
			/* 18 */ imports.NewTable("TWkWebview_AsReceiveScriptMessageDelegate", 0), // function AsReceiveScriptMessageDelegate
			/* 19 */ imports.NewTable("TWkWebview_AsWKNavigationDelegate", 0), // function AsWKNavigationDelegate
			/* 20 */ imports.NewTable("TWkWebview_AsWKURLSchemeHandlerDelegate", 0), // function AsWKURLSchemeHandlerDelegate
			/* 21 */ imports.NewTable("TWkWebview_AsWKUIDelegate", 0), // function AsWKUIDelegate
			/* 22 */ imports.NewTable("TWkWebview_AsWKDownloadDelegate", 0), // function AsWKDownloadDelegate
			/* 23 */ imports.NewTable("TWkWebview_Release", 0), // procedure Release
			/* 24 */ imports.NewTable("TWkWebview_SetAcceptsMouseMovedEvents", 0), // procedure SetAcceptsMouseMovedEvents
			/* 25 */ imports.NewTable("TWkWebview_LoadHTML", 0), // procedure LoadHTML
			/* 26 */ imports.NewTable("TWkWebview_LoadURL", 0), // procedure LoadURL
			/* 27 */ imports.NewTable("TWkWebview_SetNavigationDelegate", 0), // procedure SetNavigationDelegate
			/* 28 */ imports.NewTable("TWkWebview_SetUIDelegate", 0), // procedure SetUIDelegate
			/* 29 */ imports.NewTable("TWkWebview_InitWithFrameConfiguration", 0), // procedure InitWithFrameConfiguration
			/* 30 */ imports.NewTable("TWkWebview_StopLoading", 0), // procedure StopLoading
			/* 31 */ imports.NewTable("TWkWebview_EvaluateJavaScript", 0), // procedure EvaluateJavaScript
			/* 32 */ imports.NewTable("TWkWebview_SetAllowsBackForwardNavigationGestures", 0), // procedure SetAllowsBackForwardNavigationGestures
			/* 33 */ imports.NewTable("TWkWebview_SetAllowsMagnification", 0), // procedure SetAllowsMagnification
			/* 34 */ imports.NewTable("TWkWebview_SetMagnification", 0), // procedure SetMagnification
			/* 35 */ imports.NewTable("TWkWebview_SetMagnificationCenteredAtPoint", 0), // procedure SetMagnificationCenteredAtPoint
			/* 36 */ imports.NewTable("TWkWebview_RemoveFromSuperview", 0), // procedure RemoveFromSuperview
			/* 37 */ imports.NewTable("TWkWebview_OnProcessMessage", 0), // event OnProcessMessage
			/* 38 */ imports.NewTable("TWkWebview_OnDecidePolicyForNavigationActionPreferences", 0), // event OnDecidePolicyForNavigationActionPreferences
			/* 39 */ imports.NewTable("TWkWebview_OnDecidePolicyForNavigationResponse", 0), // event OnDecidePolicyForNavigationResponse
			/* 40 */ imports.NewTable("TWkWebview_OnStartProvisionalNavigation", 0), // event OnStartProvisionalNavigation
			/* 41 */ imports.NewTable("TWkWebview_OnReceiveServerRedirectForProvisionalNavigation", 0), // event OnReceiveServerRedirectForProvisionalNavigation
			/* 42 */ imports.NewTable("TWkWebview_OnFailProvisionalNavigationWithError", 0), // event OnFailProvisionalNavigationWithError
			/* 43 */ imports.NewTable("TWkWebview_OnCommitNavigation", 0), // event OnCommitNavigation
			/* 44 */ imports.NewTable("TWkWebview_OnFinishNavigation", 0), // event OnFinishNavigation
			/* 45 */ imports.NewTable("TWkWebview_OnFailNavigationWithError", 0), // event OnFailNavigationWithError
			/* 46 */ imports.NewTable("TWkWebview_OnWebContentProcessDidTerminate", 0), // event OnWebContentProcessDidTerminate
			/* 47 */ imports.NewTable("TWkWebview_OnNavigationActionDidBecomeDownload", 0), // event OnNavigationActionDidBecomeDownload
			/* 48 */ imports.NewTable("TWkWebview_OnNavigationResponseDidBecomeDownload", 0), // event OnNavigationResponseDidBecomeDownload
			/* 49 */ imports.NewTable("TWkWebview_OnStartURLSchemeTask", 0), // event OnStartURLSchemeTask
			/* 50 */ imports.NewTable("TWkWebview_OnStopURLSchemeTask", 0), // event OnStopURLSchemeTask
			/* 51 */ imports.NewTable("TWkWebview_OnCreateWebView", 0), // event OnCreateWebView
			/* 52 */ imports.NewTable("TWkWebview_OnRunJavaScriptAlert", 0), // event OnRunJavaScriptAlert
			/* 53 */ imports.NewTable("TWkWebview_OnRunJavaScriptConfirmCompletion", 0), // event OnRunJavaScriptConfirmCompletion
			/* 54 */ imports.NewTable("TWkWebview_OnRunJavaScriptTextInputCompletion", 0), // event OnRunJavaScriptTextInputCompletion
			/* 55 */ imports.NewTable("TWkWebview_OnWebViewDidClose", 0), // event OnWebViewDidClose
			/* 56 */ imports.NewTable("TWkWebview_OnMouseDown", 0), // event OnMouseDown
			/* 57 */ imports.NewTable("TWkWebview_OnMouseUp", 0), // event OnMouseUp
			/* 58 */ imports.NewTable("TWkWebview_OnMouseMoved", 0), // event OnMouseMoved
			/* 59 */ imports.NewTable("TWkWebview_OnMouseDragged", 0), // event OnMouseDragged
			/* 60 */ imports.NewTable("TWkWebview_OnRightMouseDown", 0), // event OnRightMouseDown
			/* 61 */ imports.NewTable("TWkWebview_OnRightMouseUp", 0), // event OnRightMouseUp
			/* 62 */ imports.NewTable("TWkWebview_OnRightMouseDragged", 0), // event OnRightMouseDragged
			/* 63 */ imports.NewTable("TWkWebview_OnOtherMouseDown", 0), // event OnOtherMouseDown
			/* 64 */ imports.NewTable("TWkWebview_OnOtherMouseUp", 0), // event OnOtherMouseUp
			/* 65 */ imports.NewTable("TWkWebview_OnOtherMouseDragged", 0), // event OnOtherMouseDragged
			/* 66 */ imports.NewTable("TWkWebview_OnDownloadCancelCompletionHandler", 0), // event OnDownloadCancelCompletionHandler
			/* 67 */ imports.NewTable("TWkWebview_OnDownloadDecideDestinationUsingResponseSuggestedFilename", 0), // event OnDownloadDecideDestinationUsingResponseSuggestedFilename
			/* 68 */ imports.NewTable("TWkWebview_OnDownloadWillPerformHTTPRedirectionNewRequest", 0), // event OnDownloadWillPerformHTTPRedirectionNewRequest
			/* 69 */ imports.NewTable("TWkWebview_OnDownloadFinish", 0), // event OnDownloadFinish
			/* 70 */ imports.NewTable("TWkWebview_OnDownloadFailWithError", 0), // event OnDownloadFailWithError
			/* 71 */ imports.NewTable("TWkWebview_OnDraggingEntered", 0), // event OnDraggingEntered
			/* 72 */ imports.NewTable("TWkWebview_OnDraggingUpdated", 0), // event OnDraggingUpdated
			/* 73 */ imports.NewTable("TWkWebview_OnDraggingExited", 0), // event OnDraggingExited
			/* 74 */ imports.NewTable("TWkWebview_OnPrepareForDragOperation", 0), // event OnPrepareForDragOperation
			/* 75 */ imports.NewTable("TWkWebview_OnPerformDragOperation", 0), // event OnPerformDragOperation
			/* 76 */ imports.NewTable("TWkWebview_OnConcludeDragOperation", 0), // event OnConcludeDragOperation
			/* 77 */ imports.NewTable("TWkWebview_OnDraggingEnded", 0), // event OnDraggingEnded
			/* 78 */ imports.NewTable("TWkWebview_OnWantsPeriodicDraggingUpdates", 0), // event OnWantsPeriodicDraggingUpdates
			/* 79 */ imports.NewTable("TWkWebview_OnUpdateDraggingItemsForDrag", 0), // event OnUpdateDraggingItemsForDrag
		}
	})
	return wkWebviewImport
}
