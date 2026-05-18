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
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"

	wvTypes "github.com/energye/wv/types/windows"
)

// IWVBrowserEvents Parent: IObject
type IWVBrowserEvents interface {
	IObject
}

// IWVBrowserBase Parent: IWVBrowserEvents IComponent
type IWVBrowserBase interface {
	IWVBrowserEvents
	IComponent
	// CreateBrowserWithHandleBool
	//  Used to create the browser using the global environment by default.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aHandle">The TWVWindowParent handle.</param>
	//  <param name="aUseDefaultEnvironment">Use the global environment or create a new one for this browser.</param>
	CreateBrowserWithHandleBool(handle types.THandle, useDefaultEnvironment bool) bool // function
	// CreateBrowserWithHandleEnvironment
	//  Used to create the browser using a custom environment. The browser will be
	//  fully initialized when the TWVBrowserBase.OnAfterCreated event is triggered.
	//  <param name="aHandle">The TWVWindowParent handle.</param>
	//  <param name="aEnvironment">Custom environment to be used by this browser.</param>
	CreateBrowserWithHandleEnvironment(handle types.THandle, environment ICoreWebView2Environment) bool // function
	// CreateWindowlessBrowserWithHandleBool
	//  Used to create a windowless browser using the global environment by default.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aHandle">The TWVDirectCompositionHost handle.</param>
	//  <param name="aUseDefaultEnvironment">Use the global environment or create a new one for this browser.</param>
	CreateWindowlessBrowserWithHandleBool(handle types.THandle, useDefaultEnvironment bool) bool // function
	// CreateWindowlessBrowserWithHandleEnvironment
	//  Used to create a windowless browser using a custom environment. The browser will be
	//  fully initialized when the TWVBrowserBase.OnAfterCreated event is triggered.
	//  <param name="aHandle">The TWVDirectCompositionHost handle.</param>
	//  <param name="aEnvironment">Custom environment to be used by this browser.</param>
	CreateWindowlessBrowserWithHandleEnvironment(handle types.THandle, environment ICoreWebView2Environment) bool // function
	// CreateInvisibleBrowserWithBool
	//  Used to create an invisible browser using the global environment by default.
	//  You are not able to reparent the window after you have created the browser.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aUseDefaultEnvironment">Use the global environment or create a new one for this browser.</param>
	CreateInvisibleBrowserWithBool(useDefaultEnvironment bool) bool // function
	// CreateInvisibleBrowserWithEnvironment
	//  Used to create an invisible browser using a custom environment.
	//  You are not able to reparent the window after you have created the browser.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aEnvironment">Custom environment to be used by this browser.</param>
	CreateInvisibleBrowserWithEnvironment(environment ICoreWebView2Environment) bool // function
	// GoBack
	//  Navigates the WebView to the previous page in the navigation history.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#goback">See the ICoreWebView2 article.</see>
	GoBack() bool // function
	// GoForward
	//  Navigates the WebView to the next page in the navigation history.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#goforward">See the ICoreWebView2 article.</see>
	GoForward() bool // function
	// Refresh
	//  Reload the current page. This is similar to navigating to the URI of
	//  current top level document including all navigation events firing and
	//  respecting any entries in the HTTP cache. But, the back or forward
	//  history are not modified.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#reload">See the ICoreWebView2 article.</see>
	Refresh() bool // function
	// RefreshIgnoreCache
	//  Reload the current page. Browser cache is ignored as if the user pressed Shift+refresh.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnRefreshIgnoreCacheCompleted event when it finishes executing.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload">See the Page Domain article.</see>
	RefreshIgnoreCache() bool // function
	// Stop
	//  Stop all navigations and pending resource fetches. Does not stop scripts.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#stop">See the ICoreWebView2 article.</see>
	Stop() bool // function
	// Navigate
	//  Cause a navigation of the top-level document to run to the specified URI.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#navigate">See the ICoreWebView2 article.</see>
	Navigate(uRI string) bool // function
	// NavigateToString
	//  Initiates a navigation to aHTMLContent as source HTML of a new document.
	//  The origin of the new page is `about:blank`.
	//  <param name="aHTMLContent">Source HTML. It may not be larger than 2 MB (2 * 1024 * 1024 bytes) in total size.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#navigatetostring">See the ICoreWebView2 article.</see>
	NavigateToString(hTMLContent string) bool // function
	// NavigateWithWebResourceRequest
	//  Navigates using a constructed ICoreWebView2WebResourceRequest object. This lets you
	//  provide post data or additional request headers during navigation.
	//  The headers in aRequest override headers added by WebView2 runtime except for Cookie headers.
	//  Method can only be either "GET" or "POST". Provided post data will only
	//  be sent only if the method is "POST" and the uri scheme is HTTP(S).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#navigatewithwebresourcerequest">See the ICoreWebView2_2 article.</see>
	NavigateWithWebResourceRequest(request ICoreWebView2WebResourceRequest) bool // function
	// SubscribeToDevToolsProtocolEvent
	//  Subscribe to a DevTools protocol event. The TWVBrowserBase.OnDevToolsProtocolEventReceived
	//  event will be triggered on each DevTools event.
	//  <param name="aEventName">The DevTools protocol event name.</param>
	//  <param name="aEventID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceiver#add_devtoolsprotocoleventreceived">See the ICoreWebView2DevToolsProtocolEventReceiver article.</see>
	SubscribeToDevToolsProtocolEvent(eventName string, eventID int32) bool // function
	// CallDevToolsProtocolMethod
	//  Runs an asynchronous `DevToolsProtocol` method.
	//  The TWVBrowserBase.OnCallDevToolsProtocolMethodCompleted event is triggered
	//  when it finishes executing. This function returns E_INVALIDARG if the `aMethodName` is
	//  unknown or the `aParametersAsJson` has an error. In the case of such an error, the
	//  `aReturnObjectAsJson` parameter of the event will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method is finished before calling the next method.
	//  If the method is to run in add_NewWindowRequested handler it should be called
	//  before the new window is set if the cdp message should affect the initial navigation. If
	//  called after setting the NewWindow property, the cdp messages
	//  may or may not apply to the initial navigation and may only apply to the subsequent navigation.
	//  For more details see `ICoreWebView2NewWindowRequestedEventArgs::put_NewWindow`.
	//  <param name="aMethodName">The DevTools protocol full method name.</param>
	//  <param name="aParametersAsJson">JSON formatted string containing the parameters for the corresponding method.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot">See the Chrome DevTools Protocol web page.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#calldevtoolsprotocolmethod">See the ICoreWebView2 article.</see>
	CallDevToolsProtocolMethod(methodName string, parametersAsJson string, executionID int32) bool // function
	// CallDevToolsProtocolMethodForSession
	//  Runs an asynchronous `DevToolsProtocol` method for a specific session of
	//  an attached target.
	//  There could be multiple `DevToolsProtocol` targets in a WebView.
	//  Besides the top level page, iframes from different origin and web workers
	//  are also separate targets. Attaching to these targets allows interaction with them.
	//  When the DevToolsProtocol is attached to a target, the connection is identified
	//  by a sessionId.
	//  To use this API, you must set the `flatten` parameter to true when calling
	//  `Target.attachToTarget` or `Target.setAutoAttach` `DevToolsProtocol` method.
	//  Using `Target.setAutoAttach` is recommended as that would allow you to attach
	//  to dedicated worker targets, which are not discoverable via other APIs like
	//  `Target.getTargets`.
	//  The TWVBrowserBase.OnCallDevToolsProtocolMethodCompleted event is triggered
	//  when it finishes executing. This function returns E_INVALIDARG if the `aMethodName` is
	//  unknown or the `aParametersAsJson` has an error. In the case of such an error, the
	//  `aReturnObjectAsJson` parameter of the event will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method is finished before calling the next method.
	//  <param name="aSessionId">The sessionId for an attached target. An empty string is treated as the session for the default target for the top page.</param>
	//  <param name="aMethodName">The DevTools protocol full method name.</param>
	//  <param name="aParametersAsJson">JSON formatted string containing the parameters for the corresponding method.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot">See the Chrome DevTools Protocol web page.</see>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Target">Information about targets and sessions.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_11#calldevtoolsprotocolmethodforsession">See the ICoreWebView2_11 article.</see>
	CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, executionID int32) bool // function
	// SetFocus
	//  Moves focus into WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</see>
	SetFocus() bool // function
	// FocusNext
	//  Moves the focus to the next element.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</see>
	FocusNext() bool // function
	// FocusPrevious
	//  Moves the focus to the previous element.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</see>
	FocusPrevious() bool // function
	// ExecuteScriptWithResult
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptWithResultCompleted event is triggered
	//  when it finishes executing.
	//  The result of the execution is returned asynchronously in the ICoreWebView2ExecuteScriptResult object
	//  which has methods and properties to obtain the successful result of script execution as well as any
	//  unhandled JavaScript exceptions.
	//  If this method is run after the NavigationStarting event during a navigation, the script
	//  runs in the new document when loading it, around the time
	//  ContentLoading is run. This operation executes the script even if
	//  ICoreWebView2Settings.IsScriptEnabled is set to FALSE.
	//  <param name="aJavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_21#executescriptwithresult">See the icorewebview2_21 article.</see>
	ExecuteScriptWithResult(javaScript string, executionID int32) bool // function
	// ExecuteScript
	//  Run JavaScript code from the aJavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptCompleted event is triggered
	//  when it finishes executing.
	//  The result of evaluating the provided JavaScript is available in the
	//  aResultObjectAsJson parameter of the TWVBrowserBase.OnExecuteScriptCompleted
	//  event as a JSON encoded string. If the result is undefined, contains a reference
	//  cycle, or otherwise is not able to be encoded into JSON, then the result
	//  is considered to be null, which is encoded in JSON as the string "null".
	//  If the script that was run throws an unhandled exception, then the result is
	//  also "null".
	//  If the method is run after the `NavigationStarting` event during a navigation,
	//  the script runs in the new document when loading it, around the time
	//  `ContentLoading` is run. This operation executes the script even if
	//  `ICoreWebView2Settings.IsScriptEnabled` is set to `FALSE`.
	//  <param name="aJavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#executescript">See the ICoreWebView2 article.</see>
	ExecuteScript(javaScript string, executionID int32) bool // function
	// CapturePreview
	//  Capture an image of what WebView is displaying. Specify the format of
	//  the image with the aImageFormat parameter. The resulting image binary
	//  data is written to the provided aImageStream parameter. This method fails if called
	//  before the first ContentLoading event. For example if this is called in
	//  the NavigationStarting event for the first navigation it will fail.
	//  For subsequent navigations, the method may not fail, but will not capture
	//  an image of a given webpage until the ContentLoading event has been fired
	//  for it. Any call to this method prior to that will result in a capture of
	//  the page being navigated away from. When this function finishes writing to the stream,
	//  the TWVBrowserBase.OnCapturePreviewCompleted event is triggered.
	//  <param name="aImageFormat">The format of the image.</param>
	//  <param name="aImageStream">The resulting image binary data is written to this stream.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#capturepreview">See the ICoreWebView2 article.</see>
	CapturePreview(imageFormat wvTypes.TWVCapturePreviewImageFormat, imageStream lcl.IStreamAdapter) bool // function
	// NotifyParentWindowPositionChanged
	//  This is a notification that tells WebView that the main WebView parent
	//  (or any ancestor) `HWND` moved. This is needed for accessibility and
	//  certain dialogs in WebView to work correctly.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#notifyparentwindowpositionchanged">See the ICoreWebView2Controller article.</see>
	NotifyParentWindowPositionChanged() bool // function
	// SetPermissionState
	//  Sets permission state for the given permission kind and origin
	//  asynchronously. The change persists across sessions until it is changed by
	//  another call to `SetPermissionState`, or by setting the `State` property
	//  in `PermissionRequestedEventArgs`.
	//  Setting the state to `COREWEBVIEW2_PERMISSION_STATE_DEFAULT` will
	//  erase any state saved in the profile and restore the default behavior.
	//  The origin should have a valid scheme and host (e.g. "https://www.example.com"),
	//  otherwise the method fails with `E_INVALIDARG`. Additional URI parts like
	//  path and fragment are ignored. For example, "https://wwww.example.com/app1/index.html/"
	//  is treated the same as "https://wwww.example.com".
	//  This function triggers the TWVBrowserBase.OnSetPermissionStateCompleted event.
	//  <see href="https://developer.mozilla.org/en-US/docs/Glossary/Origin">See the MDN origin definition.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#setpermissionstate">See the ICoreWebView2Profile4 article.</see>
	SetPermissionState(permissionKind wvTypes.TWVPermissionKind, origin string, state wvTypes.TWVPermissionState) bool // function
	// GetNonDefaultPermissionSettings
	//  Invokes the handler with a collection of all nondefault permission settings.
	//  Use this method to get the permission state set in the current and previous
	//  sessions.
	//  This function triggers the TWVBrowserBase.OnGetNonDefaultPermissionSettingsCompleted event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#getnondefaultpermissionsettings">See the ICoreWebView2Profile4 article.</see>
	GetNonDefaultPermissionSettings() bool // function
	// AddBrowserExtension
	//  Adds the [browser extension](https://developer.mozilla.org/docs/Mozilla/Add-ons/WebExtensions)
	//  using the extension path for unpacked extensions from the local device. Extension is
	//  running right after installation.
	//  The extension folder path is the topmost folder of an unpacked browser extension and
	//  contains the browser extension manifest file.
	//  If the `extensionFolderPath` is an invalid path or doesn't contain the extension manifest.json
	//  file, this function will return `ERROR_FILE_NOT_FOUND` to callers.
	//  Installed extension will default `IsEnabled` to true.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `AddBrowserExtension` will fail and return
	//  HRESULT `ERROR_NOT_SUPPORTED`.
	//  During installation, the content of the extension is not copied to the user data folder.
	//  Once the extension is installed, changing the content of the extension will cause the
	//  extension to be removed from the installed profile.
	//  When an extension is added the extension is persisted in the corresponding profile. The
	//  extension will still be installed the next time you use this profile.
	//  When an extension is installed from a folder path, adding the same extension from the same
	//  folder path means reinstalling this extension. When two extensions with the same Id are
	//  installed, only the later installed extension will be kept.
	//  Extensions that are designed to include any UI interactions (e.g. icon, badge, pop up, etc.)
	//  can be loaded and used but will have missing UI entry points due to the lack of browser
	//  UI elements to host these entry points in WebView2.
	//  The following summarizes the possible error values that can be returned from
	//  `AddBrowserExtension` and a description of why these errors occur.
	//  <code>
	//  Error value | Description
	//  ----------------------------------------------- | --------------------------
	//  `HRESULT_FROM_WIN32(ERROR_NOT_SUPPORTED)` | Extensions are disabled.
	//  `HRESULT_FROM_WIN32(ERROR_FILE_NOT_FOUND)` | Cannot find `manifest.json` file or it is not a valid extension manifest.
	//  `E_ACCESSDENIED` | Cannot load extension with file or directory name starting with \"_\", reserved for use by the system.
	//  `E_FAIL` | Extension failed to install with other unknown reasons.
	//  </code>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#addbrowserextension">See the ICoreWebView2Profile7 article.</see>
	//  The TWVBrowserBase.OnProfileAddBrowserExtensionCompleted event is triggered when TWVBrowserBase.AddBrowserExtension finishes executing.
	AddBrowserExtension(extensionFolderPath string) bool // function
	// GetBrowserExtensions
	//  Gets a snapshot of the set of extensions installed at the time `GetBrowserExtensions` is
	//  called. If an extension is installed or uninstalled after `GetBrowserExtensions` completes,
	//  the list returned by `GetBrowserExtensions` remains the same.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `GetBrowserExtensions` won't return any
	//  extensions on current user profile.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#getbrowserextensions">See the ICoreWebView2Profile7 article.</see>
	//  The TWVBrowserBase.OnProfileGetBrowserExtensionsCompleted event is triggered when TWVBrowserBase.GetBrowserExtensions finishes executing.
	GetBrowserExtensions() bool // function
	// DeleteProfile
	//  After the API is called, the profile will be marked for deletion. The
	//  local profile's directory will be deleted at browser process exit. If it
	//  fails to delete, because something else is holding the files open,
	//  WebView2 will try to delete the profile at all future browser process
	//  starts until successful.
	//  The corresponding CoreWebView2s will be closed and the
	//  CoreWebView2Profile.Deleted event will be raised. See
	//  `CoreWebView2Profile.Deleted` for more information.
	//  If you try to create a new profile with the same name as an existing
	//  profile that has been marked as deleted but hasn't yet been deleted,
	//  profile creation will fail with HRESULT_FROM_WIN32(ERROR_DELETE_PENDING).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile8#delete">See the ICoreWebView2Profile8 article.</see>
	DeleteProfile() bool // function
	// TrySuspend
	//  An app may call the `TrySuspend` API to have the WebView2 consume less memory.
	//  This is useful when a Win32 app becomes invisible, or when a Universal Windows
	//  Platform app is being suspended, during the suspended event handler before completing
	//  the suspended event.
	//  The IsVisible property must be false when the API is called.
	//  Otherwise, the API fails with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Suspending is similar to putting a tab to sleep in the Edge browser. Suspending pauses
	//  WebView script timers and animations, minimizes CPU usage for the associated
	//  browser renderer process and allows the operating system to reuse the memory that was
	//  used by the renderer process for other processes.
	//  Note that Suspend is best effort and considered completed successfully once the request
	//  is sent to browser renderer process. If there is a running script, the script will continue
	//  to run and the renderer process will be suspended after that script is done.
	//  for conditions that might prevent WebView from being suspended. In those situations,
	//  the completed handler will be invoked with isSuccessful as false and errorCode as S_OK.
	//  The WebView will be automatically resumed when it becomes visible. Therefore, the
	//  app normally does not have to call `Resume` explicitly.
	//  The app can call `Resume` and then `TrySuspend` periodically for an invisible WebView so that
	//  the invisible WebView can sync up with latest data and the page ready to show fresh content
	//  when it becomes visible.
	//  All WebView APIs can still be accessed when a WebView is suspended. Some APIs like Navigate
	//  will auto resume the WebView. To avoid unexpected auto resume, check `IsSuspended` property
	//  before calling APIs that might change WebView state.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnTrySuspendCompleted event when it finishes executing.
	//  <see href="https://techcommunity.microsoft.com/t5/articles/sleeping-tabs-faq/m-p/1705434">See the sleeping Tabs FAQ.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#trysuspend">See the ICoreWebView2_3 article.</see>
	TrySuspend() bool // function
	// Resume
	//  Resumes the WebView so that it resumes activities on the web page.
	//  This API can be called while the WebView2 controller is invisible.
	//  The app can interact with the WebView immediately after `Resume`.
	//  WebView will be automatically resumed when it becomes visible.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#resume">See the ICoreWebView2_3 article.</see>
	Resume() bool // function
	// SetVirtualHostNameToFolderMapping
	//  Sets a mapping between a virtual host name and a folder path to make available to web sites
	//  via that host name.
	//  After setting the mapping, documents loaded in the WebView can use HTTP or HTTPS URLs at
	//  the specified host name specified by hostName to access files in the local folder specified
	//  by folderPath.
	//  <param name="aHostName">Host name to access files in the local folder specified by aFolderPath.</param>
	//  <param name="aFolderPath">The path to the local files. Both absolute and relative paths are supported. Relative paths are interpreted as relative to the folder where the exe of the app is in. Note that the aFolderPath length must not exceed the Windows MAX_PATH limit.</param>
	//  <param name="aAccessKind">aAccessKind specifies the level of access to resources under the virtual host from other sites.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#setvirtualhostnametofoldermapping">See the ICoreWebView2_3 article.</see>
	SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind wvTypes.TWVHostResourceAcccessKind) bool // function
	// ClearVirtualHostNameToFolderMapping
	//  Clears a host name mapping for local folder that was added by `SetVirtualHostNameToFolderMapping`.
	//  <param name="aHostName">Host name used previously with SetVirtualHostNameToFolderMapping to access files in the local folder.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#clearvirtualhostnametofoldermapping">See the ICoreWebView2_3 article.</see>
	ClearVirtualHostNameToFolderMapping(hostName string) bool // function
	// RetrieveHTML
	//  Retrieve the HTML contents. The TWVBrowserBase.OnRetrieveHTMLCompleted event is triggered asynchronously with the HTML contents.
	RetrieveHTML() bool // function
	// RetrieveText
	//  Retrieve the text contents. The TWVBrowserBase.OnRetrieveTextCompleted event is triggered asynchronously with the text contents.
	//  <param name="aVisibleTextOnly">Exclude text that is hidden with CSS or rendered as invisible due to its parent's visibility settings.</param>
	RetrieveText(visibleTextOnly bool) bool // function
	// RetrieveMHTML
	//  Retrieve the web page contents in MHTML format. The TWVBrowserBase.OnRetrieveMHTMLCompleted event is triggered asynchronously with the MHTML contents.
	RetrieveMHTML() bool // function
	// Print
	//  Print the current web page asynchronously to the specified printer with the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#print">See the ICoreWebView2_16 article.</see>
	Print() bool // function
	// ShowPrintUI
	//  Opens the print dialog to print the current web page using the system print dialog or the browser print preview dialog.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#showprintui">See the ICoreWebView2_16 article.</see>
	ShowPrintUI(useSystemPrintDialog bool) bool // function
	// PrintToPdf
	//  Print the current page to PDF asynchronously with the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintToPdfCompleted event when it finishes executing.
	//  <param name="aResultFilePath">The path to the PDF file.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_7#printtopdf">See the ICoreWebView2_7 article.</see>
	PrintToPdf(resultFilePath string) bool // function
	// PrintToPdfStream
	//  Provides the Pdf data of current web page asynchronously for the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  Stream will be rewound to the start of the pdf data.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintToPdfStreamCompleted event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#printtopdfstream">See the ICoreWebView2_16 article.</see>
	PrintToPdfStream() bool // function
	// OpenDevToolsWindow
	//  Opens the DevTools window for the current document in the WebView. Does
	//  nothing if run when the DevTools window is already open.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#opendevtoolswindow">See the ICoreWebView2 article.</see>
	OpenDevToolsWindow() bool // function
	// OpenTaskManagerWindow
	//  Opens the Browser Task Manager view as a new window in the foreground.
	//  If the Browser Task Manager is already open, this will bring it into
	//  the foreground. WebView2 currently blocks the Shift+Esc shortcut for
	//  opening the task manager. An end user can open the browser task manager
	//  manually via the `Browser task manager` entry of the DevTools window's
	//  title bar's context menu.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_6#opentaskmanagerwindow">See the ICoreWebView2_6 article.</see>
	OpenTaskManagerWindow() bool // function
	// PostWebMessageAsJson
	//  Post the specified webMessage to the top level document in this WebView.
	//  The main page receives the message by subscribing to the `message` event of the
	//  `window.chrome.webview` of the page document.
	//  <code>
	//  ```cpp
	//  window.chrome.webview.addEventListener('message', handler)
	//  window.chrome.webview.removeEventListener('message', handler)
	//  ```
	//  </code>
	//  The event args is an instance of `MessageEvent`. The
	//  `ICoreWebView2Settings.IsWebMessageEnabled` setting must be `TRUE` or
	//  the web message will not be sent. The `data` property of the event
	//  arg is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event arg is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before
	//  the message is posted to the page, the message is discarded.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#postwebmessageasjson">See the ICoreWebView2 article.</see>
	PostWebMessageAsJson(webMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  arg of the `window.chrome.webview` message is a string with the same
	//  value as `aWebMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#postwebmessageasstring">See the ICoreWebView2 article.</see>
	PostWebMessageAsString(webMessageAsString string) bool // function
	// AddWebResourceRequestedFilter
	//  Warning: This method is deprecated and does not behave as expected for
	//  iframes. It will cause the WebResourceRequested event to fire only for the
	//  main frame and its same-origin iframes. Please use
	//  `AddWebResourceRequestedFilterWithRequestSourceKinds`
	//  instead, which will let the event to fire for all iframes on the
	//  document.
	//  Adds a URI and resource context filter for the `WebResourceRequested`
	//  event. A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string will be raised via the `WebResourceRequested` event.
	//  The `aURI` parameter value is a wildcard string matched against the URI
	//  of the web resource request. This is a glob style
	//  wildcard string in which a `*` matches zero or more characters and a `?`
	//  matches exactly one character.
	//  These wildcard characters can be escaped using a backslash just before
	//  the wildcard character in order to represent the literal `*` or `?`.
	//  The matching occurs over the URI as a whole string and not limiting
	//  wildcard matches to particular parts of the URI.
	//  The wildcard filter is compared to the URI after the URI has been
	//  normalized, any URI fragment has been removed, and non-ASCII hostnames
	//  have been converted to punycode.
	//  Specifying an empty string for aURI matches no URIs.
	//  For more information about resource context filters, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_CONTEXT](/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_web_resource_context).
	//  <code>
	//  | URI Filter String | Request URI | Match | Notes |
	//  | ---- | ---- | ---- | ---- |
	//  | `*` | `https://contoso.com/a/b/c` | Yes | A single * will match all URIs |
	//  | `*://contoso.com/*` | `https://contoso.com/a/b/c` | Yes | Matches everything in contoso.com across all schemes |
	//  | `*://contoso.com/*` | `https://example.com/?https://contoso.com/` | Yes | But also matches a URI with just the same text anywhere in the URI |
	//  | `example` | `https://contoso.com/example` | No | The filter does not perform partial matches |
	//  | `*example` | `https://contoso.com/example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?query#example` | No | The filter is matched against the URI with no fragment |
	//  | `*example` | `https://example` | No | The URI is normalized before filter matching so the actual URI used for comparison is `https://example/` |
	//  | `*example/` | `https://example` | Yes | Just like above, but this time the filter ends with a / just like the normalized URI |
	//  | `https://xn--qei.example/` | `https://&#x2764;.example/` | Yes | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  | `https://&#x2764;.example/` | `https://xn--qei.example/` | No | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  </code>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter">See the ICoreWebView2 article.</see>
	AddWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool // function
	// RemoveWebResourceRequestedFilter
	//  Warning: This method and `AddWebResourceRequestedFilter` are deprecated.
	//  Please use `AddWebResourceRequestedFilterWithRequestSourceKinds` and
	//  `RemoveWebResourceRequestedFilterWithRequestSourceKinds` instead.
	//  Removes a matching WebResource filter that was previously added for the
	//  `WebResourceRequested` event. If the same filter was added multiple
	//  times, then it must be removed as many times as it was added for the
	//  removal to be effective. Returns `E_INVALIDARG` for a filter that was
	//  never added.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removewebresourcerequestedfilter">See the ICoreWebView2 article.</see>
	RemoveWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool // function
	// AddWebResourceRequestedFilterWithRequestSourceKinds
	//  A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string for corresponding request sources will be raised via
	//  the `WebResourceRequested` event. To receive all raised events filters
	//  have to be added before main page navigation.
	//  The `uri` parameter value is a wildcard string matched against the URI
	//  of the web resource request. This is a glob style
	//  wildcard string in which a `*` matches zero or more characters and a `?`
	//  matches exactly one character.
	//  These wildcard characters can be escaped using a backslash just before
	//  the wildcard character in order to represent the literal `*` or `?`.
	//  The matching occurs over the URI as a whole string and not limiting
	//  wildcard matches to particular parts of the URI.
	//  The wildcard filter is compared to the URI after the URI has been
	//  normalized, any URI fragment has been removed, and non-ASCII hostnames
	//  have been converted to punycode.
	//  Specifying a `nullptr` for the uri is equivalent to an empty string which
	//  matches no URIs.
	//  For more information about resource context filters, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_CONTEXT](/microsoft-edge/webview2/reference/win32/icorewebview2#corewebview2_web_resource_context).
	//  The `requestSourceKinds` is a mask of one or more
	//  `COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS`. OR operation(s) can be
	//  applied to multiple `COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS` to
	//  create a mask representing those data types. API returns `E_INVALIDARG` if
	//  `requestSourceKinds` equals to zero. For more information about request
	//  source kinds, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS](/microsoft-edge/webview2/reference/win32/icorewebview2#corewebview2_web_resource_request_source_kinds).
	//  Because service workers and shared workers run separately from any one
	//  HTML document their WebResourceRequested will be raised for all
	//  CoreWebView2s that have appropriate filters added in the corresponding
	//  CoreWebView2Environment. You should only add a WebResourceRequested filter
	//  for COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SERVICE_WORKER or
	//  COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SHARED_WORKER on
	//  one CoreWebView2 to avoid handling the same WebResourceRequested
	//  event multiple times.
	//  <code>
	//  | URI Filter String | Request URI | Match | Notes |
	//  | ---- | ---- | ---- | ---- |
	//  | `*` | `https://contoso.com/a/b/c` | Yes | A single * will match all URIs |
	//  | `*://contoso.com/*` | `https://contoso.com/a/b/c` | Yes | Matches everything in contoso.com across all schemes |
	//  | `*://contoso.com/*` | `https://example.com/?https://contoso.com/` | Yes | But also matches a URI with just the same text anywhere in the URI |
	//  | `example` | `https://contoso.com/example` | No | The filter does not perform partial matches |
	//  | `*example` | `https://contoso.com/example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?query#example` | No | The filter is matched against the URI with no fragment |
	//  | `*example` | `https://example` | No | The URI is normalized before filter matching so the actual URI used for comparison is `https://example/` |
	//  | `*example/` | `https://example` | Yes | Just like above, but this time the filter ends with a / just like the normalized URI |
	//  | `https://xn--qei.example/` | `https://&#x2764;.example/` | Yes | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  | `https://&#x2764;.example/` | `https://xn--qei.example/` | No | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  </code>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_22#addwebresourcerequestedfilterwithrequestsourcekinds">See the ICoreWebView2_22 article.</see>
	AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool // function
	// RemoveWebResourceRequestedFilterWithRequestSourceKinds
	//  Removes a matching WebResource filter that was previously added for the
	//  `WebResourceRequested` event. If the same filter was added multiple
	//  times, then it must be removed as many times as it was added for the
	//  removal to be effective. Returns `E_INVALIDARG` for a filter that was
	//  not added or is already removed.
	//  If the filter was added for multiple requestSourceKinds and removed just for one of them
	//  the filter remains for the non-removed requestSourceKinds.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_22#removewebresourcerequestedfilterwithrequestsourcekinds">See the ICoreWebView2_22 article.</see>
	RemoveWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool // function
	// PostWebMessageAsJsonWithAdditionalObjects
	//  Same as `PostWebMessageAsJson`, but also has support for posting DOM objects
	//  to page content. The `additionalObjects` property in the DOM MessageEvent
	//  fired on the page content is an array-like list of DOM objects that can
	//  be posted to the web content. Currently these can be the following
	//  types of objects:
	//  <code>
	//  | Win32 | DOM type |
	//  |-------------------|-------------|
	//  | ICoreWebView2FileSystemHandle | [FileSystemHandle](https://developer.mozilla.org/docs/Web/API/FileSystemHandle) |
	//  | nullptr | null |
	//  </code>
	//  The objects are posted to the web content, following the
	//  [structured-clone](https://developer.mozilla.org/docs/Web/API/Web_Workers_API/Structured_clone_algorithm)
	//  semantics, meaning only objects that can be cloned can be posted.
	//  They will also behave as if they had been created by the web content they are
	//  posted to. For example, if a FileSystemFileHandle is posted to a web content
	//  it can only be re-transferred via postMessage to other web content
	//  [with the same origin](https://fs.spec.whatwg.org/#filesystemhandle).
	//  Warning: An app needs to be mindful when using this API to post DOM objects
	//  as this API provides the web content with unusual access to sensitive Web
	//  Platform features such as filesystem access! Similar to PostWebMessageAsJson
	//  the app should check the `Source` property of WebView2 right before posting the message
	//  to ensure the message and objects will only be sent to the target web content
	//  that it expects to receive the DOM objects. Additionally, the order
	//  of messages that are posted between `PostWebMessageAsJson` and `PostWebMessageAsJsonWithAdditionalObjects`
	//  may not be preserved.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_23#postwebmessageasjsonwithadditionalobjects">See the ICoreWebView2_23 article.</see>
	PostWebMessageAsJsonWithAdditionalObjects(webMessageAsJson string, additionalObjects ICoreWebView2ObjectCollectionView) bool // function
	// AddHostObjectToScript
	//  Add the provided host object to script running in the WebView with the
	//  specified name. Host objects are exposed as host object proxies using
	//  `window.chrome.webview.hostObjects.{name}`. Host object proxies are
	//  promises and resolves to an object representing the host object. The
	//  promise is rejected if the app has not added an object with the name.
	//  When JavaScript code access a property or method of the object, a promise
	//  is return, which resolves to the value returned from the host for the
	//  property or method, or rejected in case of error, for example, no
	//  property or method on the object or parameters are not valid.
	//
	//  NOTE: While simple types, `IDispatch` and array are supported, and
	//  `IUnknown` objects that also implement `IDispatch` are treated as `IDispatch`,
	//  generic `IUnknown`, `VT_DECIMAL`, or `VT_RECORD` variant is not supported.
	//  Remote JavaScript objects like callback functions are represented as an
	//  `VT_DISPATCH` `VARIANT` with the object implementing `IDispatch`. The
	//  JavaScript callback method may be invoked using `DISPID_VALUE` for the
	//  `DISPID`. Such callback method invocations will return immediately and will
	//  not wait for the JavaScript function to run and so will not provide the return
	//  value of the JavaScript function.
	//  Nested arrays are supported up to a depth of 3. Arrays of by
	//  reference types are not supported. `VT_EMPTY` and `VT_NULL` are mapped
	//  into JavaScript as `null`. In JavaScript, `null` and undefined are
	//  mapped to `VT_EMPTY`.
	//
	//  Additionally, all host objects are exposed as
	//  `window.chrome.webview.hostObjects.sync.{name}`. Here the host objects
	//  are exposed as synchronous host object proxies. These are not promises
	//  and function runtimes or property access synchronously block running
	//  script waiting to communicate cross process for the host code to run.
	//  Accordingly the result may have reliability issues and it is recommended
	//  that you use the promise-based asynchronous
	//  `window.chrome.webview.hostObjects.{name}` API.
	//
	//  Synchronous host object proxies and asynchronous host object proxies may
	//  both use a proxy to the same host object. Remote changes made by one
	//  proxy propagates to any other proxy of that same host object whether
	//  the other proxies and synchronous or asynchronous.
	//
	//  While JavaScript is blocked on a synchronous run to native code, that
	//  native code is unable to run back to JavaScript. Attempts to do so fail
	//  with `HRESULT_FROM_WIN32(ERROR_POSSIBLE_DEADLOCK)`.
	//
	//  Host object proxies are JavaScript Proxy objects that intercept all
	//  property get, property set, and method invocations. Properties or methods
	//  that are a part of the Function or Object prototype are run locally.
	//  Additionally any property or method in the
	//  `chrome.webview.hostObjects.options.forceLocalProperties`
	//  array are also run locally. This defaults to including optional methods
	//  that have meaning in JavaScript like `toJSON` and `Symbol.toPrimitive`.
	//  Add more to the array as required.
	//
	//  The `chrome.webview.hostObjects.cleanupSome` method performs a best
	//  effort garbage collection on host object proxies.
	//
	//  The `chrome.webview.hostObjects.options` object provides the ability to
	//  change some functionality of host objects.
	//  <code>
	//  Options property | Details
	//  ---|---
	//  `forceLocalProperties` | This is an array of host object property names that will be run locally, instead of being called on the native host object. This defaults to `then`, `toJSON`, `Symbol.toString`, and `Symbol.toPrimitive`. You can add other properties to specify that they should be run locally on the javascript host object proxy.
	//  `log` | This is a callback that will be called with debug information. For example, you can set this to `console.log.bind(console)` to have it print debug information to the console to help when troubleshooting host object usage. By default this is null.
	//  `shouldSerializeDates` | By default this is false, and javascript Date objects will be sent to host objects as a string using `JSON.stringify`. You can set this property to true to have Date objects properly serialize as a `VT_DATE` when sending to the native host object, and have `VT_DATE` properties and return values create a javascript Date object.
	//  `defaultSyncProxy` | When calling a method on a synchronous proxy, the result should also be a synchronous proxy. But in some cases, the sync/async context is lost (for example, when providing to native code a reference to a function, and then calling that function in native code). In these cases, the proxy will be asynchronous, unless this property is set.
	//  `forceAsyncMethodMatches ` | This is an array of regular expressions. When calling a method on a synchronous proxy, the method call will be performed asynchronously if the method name matches a string or regular expression in this array. Setting this value to `Async` will make any method that ends with Async be an asynchronous method call. If an async method doesn't match here and isn't forced to be asynchronous, the method will be invoked synchronously, blocking execution of the calling JavaScript and then returning the resolution of the promise, rather than returning a promise.
	//  `ignoreMemberNotFoundError` | By default, an exception is thrown when attempting to get the value of a proxy property that doesn't exist on the corresponding native class. Setting this property to `true` switches the behavior to match Chakra WinRT projection (and general JavaScript) behavior of returning `undefined` with no error.
	//  `shouldPassTypedArraysAsArrays` | By default, typed arrays are passed to the host as `IDispatch`. To instead pass typed arrays to the host as `array`, set this to `true`.
	//  </code>
	//  Host object proxies additionally have the following methods which run
	//  locally.
	//  <code>
	//  Method name | Details
	//  ---|---
	//  `applyHostFunction`, `getHostProperty`, `setHostProperty` | Perform a method invocation, property get, or property set on the host object. Use the methods to explicitly force a method or property to run remotely if a conflicting local method or property exists. For instance, `proxy.toString()` runs the local `toString` method on the proxy object. But proxy.applyHostFunction('toString') runs `toString` on the host proxied object instead.
	//  `getLocalProperty`, `setLocalProperty` | Perform property get, or property set locally. Use the methods to force getting or setting a property on the host object proxy rather than on the host object it represents. For instance, `proxy.unknownProperty` gets the property named `unknownProperty` from the host proxied object. But proxy.getLocalProperty('unknownProperty') gets the value of the property `unknownProperty` on the proxy object.
	//  `sync` | Asynchronous host object proxies expose a sync method which returns a promise for a synchronous host object proxy for the same host object. For example, `chrome.webview.hostObjects.sample.methodCall()` returns an asynchronous host object proxy. Use the `sync` method to obtain a synchronous host object proxy instead: `const syncProxy = await chrome.webview.hostObjects.sample.methodCall().sync()`.
	//  `async` | Synchronous host object proxies expose an async method which blocks and returns an asynchronous host object proxy for the same host object. For example, `chrome.webview.hostObjects.sync.sample.methodCall()` returns a synchronous host object proxy. Running the `async` method on this blocks and then returns an asynchronous host object proxy for the same host object: `const asyncProxy = chrome.webview.hostObjects.sync.sample.methodCall().async()`.
	//  `then` | Asynchronous host object proxies have a `then` method. Allows proxies to be awaitable. `then` returns a promise that resolves with a representation of the host object. If the proxy represents a JavaScript literal, a copy of that is returned locally. If the proxy represents a function, a non-awaitable proxy is returned. If the proxy represents a JavaScript object with a mix of literal properties and function properties, the a copy of the object is returned with some properties as host object proxies.
	//  </code>
	//  All other property and method invocations (other than the above Remote
	//  object proxy methods, `forceLocalProperties` list, and properties on
	//  Function and Object prototypes) are run remotely. Asynchronous host
	//  object proxies return a promise representing asynchronous completion of
	//  remotely invoking the method, or getting the property. The promise
	//  resolves after the remote operations complete and the promises resolve to
	//  the resulting value of the operation. Synchronous host object proxies
	//  work similarly, but block running JavaScript and wait for the remote
	//  operation to complete.
	//
	//  Setting a property on an asynchronous host object proxy works slightly
	//  differently. The set returns immediately and the return value is the
	//  value that is set. This is a requirement of the JavaScript Proxy object.
	//  If you need to asynchronously wait for the property set to complete, use
	//  the `setHostProperty` method which returns a promise as described above.
	//  Synchronous object property set property synchronously blocks until the
	//  property is set.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addhostobjecttoscript">See the ICoreWebView2 article.</see>
	AddHostObjectToScript(name string, object types.OleVariant) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the WebView. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the WebView, the JavaScript code continues to have access to that
	//  object. Run this method for a name that is already removed or never
	//  added fails.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removehostobjectfromscript">See the ICoreWebView2 article.</see>
	RemoveHostObjectFromScript(name string) bool // function
	// AddScriptToExecuteOnDocumentCreated
	//  Add the provided JavaScript to a list of scripts that should be run after
	//  the global object has been created, but before the HTML document has
	//  been parsed and before any other script included by the HTML document is
	//  run. This method injects a script that runs on all top-level document
	//  and child frame page navigations. This method runs asynchronously, and
	//  you must wait for the completion handler to finish before the injected
	//  script is ready to run. When this method completes, the `Invoke` method
	//  of the handler is run with the `id` of the injected script. `id` is a
	//  string. To remove the injected script, use
	//  `RemoveScriptToExecuteOnDocumentCreated`.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnAddScriptToExecuteOnDocumentCreatedCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addscripttoexecuteondocumentcreated">See the ICoreWebView2 article.</see>
	AddScriptToExecuteOnDocumentCreated(javaScript string) bool // function
	// RemoveScriptToExecuteOnDocumentCreated
	//  Remove the corresponding JavaScript added using
	//  `AddScriptToExecuteOnDocumentCreated` with the specified script ID. The
	//  script ID should be the one returned by the `AddScriptToExecuteOnDocumentCreated`.
	//  Both use `AddScriptToExecuteOnDocumentCreated` and this method in `NewWindowRequested`
	//  event handler at the same time sometimes causes trouble. Since invalid scripts will
	//  be ignored, the script IDs you got may not be valid anymore.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removescripttoexecuteondocumentcreated">See the ICoreWebView2 article.</see>
	RemoveScriptToExecuteOnDocumentCreated(iD string) bool // function
	// CreateCookie
	//  Create a cookie object with a specified name, value, domain, and path.
	//  One can set other optional properties after cookie creation.
	//  This only creates a cookie object and it is not added to the cookie
	//  manager until you call AddOrUpdateCookie.
	//  Leading or trailing whitespace(s), empty string, and special characters
	//  are not allowed for name.
	//  See ICoreWebView2Cookie for more details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#createcookie">See the ICoreWebView2CookieManager article.</see>
	CreateCookie(name string, value string, domain string, path string) ICoreWebView2Cookie // function
	// CopyCookie
	//  Creates a cookie whose params matches those of the specified cookie.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#copycookie">See the ICoreWebView2CookieManager article.</see>
	CopyCookie(cookie ICoreWebView2Cookie) ICoreWebView2Cookie // function
	// GetCookies
	//  Gets a list of cookies matching the specific URI.
	//  If uri is empty string or null, all cookies under the same profile are
	//  returned.
	//  You can modify the cookie objects by calling
	//  ICoreWebView2CookieManager.AddOrUpdateCookie, and the changes
	//  will be applied to the webview.
	//  The TWVBrowserBase.OnGetCookiesCompleted event is triggered asynchronously with the cookies.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#getcookies">See the ICoreWebView2CookieManager article.</see>
	GetCookies(uRI string) bool // function
	// AddOrUpdateCookie
	//  Adds or updates a cookie with the given cookie data; may overwrite
	//  cookies with matching name, domain, and path if they exist.
	//  This method will fail if the domain of the given cookie is not specified.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#addorupdatecookie">See the ICoreWebView2CookieManager article.</see>
	AddOrUpdateCookie(cookie ICoreWebView2Cookie) bool // function
	// DeleteCookie
	//  Deletes a cookie whose name and domain/path pair
	//  match those of the specified cookie.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookie">See the ICoreWebView2CookieManager article.</see>
	DeleteCookie(cookie ICoreWebView2Cookie) bool // function
	// DeleteCookies
	//  Deletes cookies with matching name and uri.
	//  Cookie name is required.
	//  All cookies with the given name where domain
	//  and path match provided URI are deleted.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookies">See the ICoreWebView2CookieManager article.</see>
	DeleteCookies(name string, uRI string) bool // function
	// DeleteCookiesWithDomainAndPath
	//  Deletes cookies with matching name and domain/path pair.
	//  Cookie name is required.
	//  If domain is specified, deletes only cookies with the exact domain.
	//  If path is specified, deletes only cookies with the exact path.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookieswithdomainandpath">See the ICoreWebView2CookieManager article.</see>
	DeleteCookiesWithDomainAndPath(name string, domain string, path string) bool // function
	// DeleteAllCookies
	//  Deletes all cookies under the same profile.
	//  This could affect other WebViews under the same user profile.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deleteallcookies">See the ICoreWebView2CookieManager article.</see>
	DeleteAllCookies() bool // function
	// SetBoundsAndZoomFactor
	//  Updates `Bounds` and `ZoomFactor` properties at the same time. This
	//  operation is atomic from the perspective of the host. After returning
	//  from this function, the `Bounds` and `ZoomFactor` properties are both
	//  updated if the function is successful, or neither is updated if the
	//  function fails. If `Bounds` and `ZoomFactor` are both updated by the
	//  same scale (for example, `Bounds` and `ZoomFactor` are both doubled),
	//  then the page does not display a change in `window.innerWidth` or
	//  `window.innerHeight` and the WebView renders the content at the new size
	//  and zoom without intermediate renderings. This function also updates
	//  just one of `ZoomFactor` or `Bounds` by passing in the new value for one
	//  and the current value for the other.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#setboundsandzoomfactor">See the ICoreWebView2Controller article.</see>
	SetBoundsAndZoomFactor(bounds types.TRect, zoomFactor float64) bool // function
	// OpenDefaultDownloadDialog
	//  Open the default download dialog. If the dialog is opened before there
	//  are recent downloads, the dialog shows all past downloads for the
	//  current profile. Otherwise, the dialog shows only the recent downloads
	//  with a "See more" button for past downloads. Calling this method raises
	//  the `IsDefaultDownloadDialogOpenChanged` event if the dialog was closed.
	//  No effect if the dialog is already open.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#opendefaultdownloaddialog">See the ICoreWebView2_9 article.</see>
	OpenDefaultDownloadDialog() bool // function
	// CloseDefaultDownloadDialog
	//  Close the default download dialog. Calling this method raises the
	//  IsDefaultDownloadDialogOpenChanged event if the dialog was open.
	//  No effect if the dialog is already closed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#opendefaultdownloaddialog">See the ICoreWebView2_9 article.</see>
	CloseDefaultDownloadDialog() bool // function
	// SimulateEditingCommand
	//  Simulate editing commands using the "Input.dispatchKeyEvent" DevTools method.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</see>
	SimulateEditingCommand(editingCommand wvTypes.TWV2EditingCommand) bool // function
	// SimulateKeyEvent
	//  Dispatches a key event to the page using the "Input.dispatchKeyEvent"
	//  DevTools method. The browser has to be focused before simulating any
	//  key event. This function is asynchronous and it triggers the
	//  TWVBrowserBase.OnSimulateKeyEventCompleted event when it finishes executing.
	//  <param name="type_">Type of the key event.</param>
	//  <param name="modifiers">Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8.</param>
	//  <param name="windowsVirtualKeyCode">Windows virtual key code.</param>
	//  <param name="nativeVirtualKeyCode">Native virtual key code.</param>
	//  <param name="timestamp">Time at which the event occurred.</param>
	//  <param name="location">Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right.</param>
	//  <param name="autoRepeat">Whether the event was generated from auto repeat.</param>
	//  <param name="isKeypad">Whether the event was generated from the keypad.</param>
	//  <param name="isSystemKey">Whether the event was a system key event.</param>
	//  <param name="text">Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events.</param>
	//  <param name="unmodifiedtext">Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling.</param>
	//  <param name="keyIdentifier">Unique key identifier (e.g., 'U+0041').</param>
	//  <param name="code">Unique DOM defined string value for each physical key (e.g., 'KeyA').</param>
	//  <param name="key">Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr').</param>
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	SimulateKeyEvent(type_ wvTypes.TWV2KeyEventType, modifiers int32, windowsVirtualKeyCode int32, nativeVirtualKeyCode int32, timestamp int32, location int32, autoRepeat bool, isKeypad bool, isSystemKey bool, text string, unmodifiedtext string, keyIdentifier string, code string, key string) bool // function
	// KeyboardShortcutSearch
	//  Simulate that the F3 key was pressed and released.
	//  The browser has to be focused before simulating any key event.
	//  This key information was logged using a Spanish keyboard. It might not work with different keyboard layouts.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnSimulateKeyEventCompleted event several times.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	KeyboardShortcutSearch() bool // function
	// KeyboardShortcutRefreshIgnoreCache
	//  Simulate that SHIFT + F5 keys were pressed and released.
	//  The browser has to be focused before simulating any key event.
	//  This key information was logged using a Spanish keyboard. It might not work with different keyboard layouts.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnSimulateKeyEventCompleted event several times.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</see>
	KeyboardShortcutRefreshIgnoreCache() bool // function
	// SendMouseInput
	//  This function is only available in "Windowless mode" and provides mouse input meant for the WebView.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL, then mouseData specifies the amount of
	//  wheel movement. A positive value indicates that the wheel was rotated
	//  forward, away from the user; a negative value indicates that the wheel was
	//  rotated backward, toward the user. One wheel click is defined as
	//  WHEEL_DELTA, which is 120.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN, or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP, then mouseData specifies which X
	//  buttons were pressed or released. This value should be 1 if the first X
	//  button is pressed/released and 2 if the second X button is
	//  pressed/released.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE, then virtualKeys,
	//  mouseData, and point should all be zero.
	//  If eventKind is any other value, then mouseData should be zero.
	//  Point is expected to be in the client coordinate space of the WebView.
	//  To track mouse events that start in the WebView and can potentially move
	//  outside of the WebView and host application, calling SetCapture and
	//  ReleaseCapture is recommended.
	//  To dismiss hover popups, it is also recommended to send
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE messages.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#sendmouseinput">See the ICoreWebView2CompositionController article.</see>
	SendMouseInput(eventKind wvTypes.TWVMouseEventKind, virtualKeys wvTypes.TWVMouseEventVirtualKeys, mouseData uint32, point types.TPoint) bool // function
	// SendPointerInput
	//  This function is only available in "Windowless mode" and provides pointer input meant for the WebView.
	//  SendPointerInput accepts touch or pen pointer input of types defined in
	//  COREWEBVIEW2_POINTER_EVENT_KIND. Any pointer input from the system must be
	//  converted into an ICoreWebView2PointerInfo first.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#sendpointerinput">See the ICoreWebView2CompositionController article.</see>
	SendPointerInput(eventKind wvTypes.TWVPointerEventKind, pointerInfo ICoreWebView2PointerInfo) bool // function
	// DragLeave
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::DragLeave](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragLeave calls to this function.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#dragleave">See the ICoreWebView2CompositionController3 article.</see>
	DragLeave() types.HRESULT // function
	// DragOver
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::DragOver](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragOver calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates (Similar to how SendMouseInput works).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#dragover">See the ICoreWebView2CompositionController3 article.</see>
	DragOver(keyState uint32, point types.TPoint, outEffect *uint32) types.HRESULT // function
	// GetNonClientRegionAtPoint
	//  If you are hosting a WebView2 using CoreWebView2CompositionController, you can call
	//  this method in your Win32 WndProc to determine if the mouse is moving over or
	//  clicking on WebView2 web content that should be considered part of a non-client region.
	//  The point parameter is expected to be in the client coordinate space of WebView2.
	//  The method sets the out parameter value as follows:
	//  <code>
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_CAPTION when point corresponds to
	//  a region (HTML element) within the WebView2 with
	//  `-webkit-app-region: drag` CSS style set.
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_CLIENT when point corresponds to
	//  a region (HTML element) within the WebView2 without
	//  `-webkit-app-region: drag` CSS style set.
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_NOWHERE when point is not within the WebView2.
	//  </code>
	//  NOTE: in order for WebView2 to properly handle the title bar system menu,
	//  the app needs to send WM_NCRBUTTONDOWN and WM_NCRBUTTONUP to SendMouseInput.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4#getnonclientregionatpoint">See the ICoreWebView2CompositionController4 article.</see>
	GetNonClientRegionAtPoint(point types.TPoint) wvTypes.TWVNonClientRegionKind // function
	// QueryNonClientRegion
	//  This method is used to get the collection of rects that correspond
	//  to a particular TWVNonClientRegionKind. This is to be used in
	//  the callback of add_NonClientRegionChanged whose event args object contains
	//  a region property of type TWVNonClientRegionKind.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4#querynonclientregion">See the ICoreWebView2CompositionController4 article.</see>
	QueryNonClientRegion(kind wvTypes.TWVNonClientRegionKind) ICoreWebView2RegionRectCollectionView // function
	// ClearCache
	//  Clears the browser cache. This function is asynchronous and it triggers the TWVBrowserBase.OnClearCacheCompleted event when it finishes executing.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache">See the Chrome DevTools Protocol page about the Network.clearBrowserCache method.</see>
	ClearCache() bool // function
	// ClearDataForOrigin
	//  Clears the storage for origin. This function is asynchronous and it triggers the TWVBrowserBase.OnClearDataForOriginCompleted event when it finishes executing.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin">See the Chrome DevTools Protocol page about the Storage.clearDataForOrigin method.</see>
	ClearDataForOrigin(origin string, storageTypes wvTypes.TWVClearDataStorageTypes) bool // function
	// ClearBrowsingData
	//  Clear browsing data based on a data type. This method takes two parameters,
	//  the first being a mask of one or more `COREWEBVIEW2_BROWSING_DATA_KINDS`. OR
	//  operation(s) can be applied to multiple `COREWEBVIEW2_BROWSING_DATA_KINDS` to
	//  create a mask representing those data types.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdata">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingData(dataKinds wvTypes.TWVBrowsingDataKinds) bool // function
	// ClearBrowsingDataInTimeRange
	//  ClearBrowsingDataInTimeRange behaves like ClearBrowsingData except that it
	//  takes in two additional parameters for the start and end time for which it
	//  should clear the data between. The `startTime` and `endTime`
	//  parameters correspond to the number of seconds since the UNIX epoch.
	//  `startTime` is inclusive while `endTime` is exclusive, therefore the data will
	//  be cleared between startTime and endTime.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataintimerange">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingDataInTimeRange(dataKinds wvTypes.TWVBrowsingDataKinds, startTime types.TDateTime, endTime types.TDateTime) bool // function
	// ClearBrowsingDataAll
	//  ClearBrowsingDataAll behaves like ClearBrowsingData except that it
	//  clears the entirety of the data associated with the profile it is called on.
	//  It clears the data regardless of timestamp.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataall">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingDataAll() bool // function
	// ClearServerCertificateErrorActions
	//  Clears all cached decisions to proceed with TLS certificate errors from the
	//  OnServerCertificateErrorDetected event for all WebView2's sharing the same session.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnServerCertificateErrorActionsCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_14#clearservercertificateerroractions">See the ICoreWebView2_14 article.</see>
	ClearServerCertificateErrorActions() bool // function
	// GetFavicon
	//  Async function for getting the actual image data of the favicon.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnGetFaviconCompleted event when it finishes executing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#getfavicon">See the ICoreWebView2_15 article.</see>
	GetFavicon(format wvTypes.TWVFaviconImageFormat) bool // function
	// CreateSharedBuffer
	//  Create a shared memory based buffer with the specified size in bytes.
	//  The buffer can be shared with web contents in WebView by calling
	//  `PostSharedBufferToScript` on `CoreWebView2` or `CoreWebView2Frame` object.
	//  Once shared, the same content of the buffer will be accessible from both
	//  the app process and script in WebView. Modification to the content will be visible
	//  to all parties that have access to the buffer.
	//  The shared buffer is presented to the script as ArrayBuffer. All JavaScript APIs
	//  that work for ArrayBuffer including Atomics APIs can be used on it.
	//  There is currently a limitation that only size less than 2GB is supported.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the ICoreWebView2Environment12 article.</see>
	CreateSharedBuffer(size int64, sharedBuffer *ICoreWebView2SharedBuffer) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the main frame in the WebView.
	//  The script will receive a `sharedbufferreceived` event from chrome.webview.
	//  Read the linked article for all the details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_17#postsharedbuffertoscript">See the ICoreWebView2_17 article.</see>
	PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool // function
	// GetProcessExtendedInfos
	//  Gets a snapshot collection of `ProcessExtendedInfo`s corresponding to all
	//  currently running processes associated with this `ICoreWebView2Environment`
	//  excludes crashpad process.
	//  This provides the same list of `ProcessInfo`s as what's provided in
	//  `GetProcessInfos`, but additionally provides a list of associated `FrameInfo`s
	//  which are actively running (showing or hiding UI elements) in the renderer
	//  process. See `AssociatedFrameInfos` for more information.
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment13#getprocessextendedinfos">See the ICoreWebView2Environment13 article.</see>
	//  This function triggers the TWVBrowserBase.OnGetProcessExtendedInfosCompleted event.
	GetProcessExtendedInfos() bool // function
	// ShowSaveAsUI
	//  Programmatically trigger a Save As action for the currently loaded document.
	//  Opens a system modal dialog by default. If the `SuppressDefaultDialog` property
	//  is `TRUE`, the system dialog is not opened.
	//  This method returns `COREWEBVIEW2_SAVE_AS_UI_RESULT`. See
	//  `COREWEBVIEW2_SAVE_AS_UI_RESULT` for details.
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_25#showsaveasui">See the ICoreWebView2_25 article.</see>
	//  This function triggers the TWVBrowserBase.OnSaveAsUIShowing and TWVBrowserBase.OnShowSaveAsUICompleted events.
	ShowSaveAsUI() bool // function
	// MoveFormTo
	//  Move the parent form to the x and y coordinates.
	MoveFormTo(X int32, Y int32) // procedure
	// MoveFormBy
	//  Move the parent form adding x and y to the coordinates.
	MoveFormBy(X int32, Y int32) // procedure
	// ResizeFormWidthTo
	//  Add x to the parent form width.
	ResizeFormWidthTo(X int32) // procedure
	// ResizeFormHeightTo
	//  Add y to the parent form height.
	ResizeFormHeightTo(Y int32) // procedure
	// SetFormLeftTo
	//  Set the parent form left property to x.
	SetFormLeftTo(X int32) // procedure
	// SetFormTopTo
	//  Set the parent form top property to y.
	SetFormTopTo(Y int32) // procedure
	// IncZoomStep
	//  Increments the browser zoom.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	IncZoomStep() // procedure
	// DecZoomStep
	//  Decrements the browser zoom.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	DecZoomStep() // procedure
	// ResetZoom
	//  Sets the browser zoom to 100%.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	ResetZoom() // procedure
	// ToggleMuteState
	//  Enables or disables all audio output from this browser.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</see>
	ToggleMuteState() // procedure
	// Initialized
	//  Custom properties
	Initialized() bool // property Initialized Getter
	// CoreWebView2PrintSettings
	//  Settings used for printing.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings">See the ICoreWebView2PrintSettings article.</see>
	CoreWebView2PrintSettings() ICoreWebView2PrintSettings // property CoreWebView2PrintSettings Getter
	// CoreWebView2Settings
	//  CoreWebView2Settings contains various modifiable settings for the running WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings">See the ICoreWebView2Settings article.</see>
	CoreWebView2Settings() ICoreWebView2Settings // property CoreWebView2Settings Getter
	// CoreWebView2Environment
	//  Represents the WebView2 Environment.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment">See the ICoreWebView2Environment article.</see>
	CoreWebView2Environment() ICoreWebView2Environment // property CoreWebView2Environment Getter
	// CoreWebView2Controller
	//  The owner of the `CoreWebView2` object that provides support for resizing,
	//  showing and hiding, focusing, and other functionality related to
	//  windowing and composition. The `CoreWebView2Controller` owns the
	//  `CoreWebView2`, and if all references to the `CoreWebView2Controller` go
	//  away, the WebView is closed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller">See the ICoreWebView2Controller article.</see>
	CoreWebView2Controller() ICoreWebView2Controller // property CoreWebView2Controller Getter
	// CoreWebView2CompositionController
	//  ICoreWebView2CompositionController wrapper used by this browser.
	//  This interface is an extension of the ICoreWebView2Controller interface to
	//  support visual hosting. An object implementing the
	//  ICoreWebView2CompositionController interface will also implement
	//  ICoreWebView2Controller. Callers are expected to use
	//  ICoreWebView2Controller for resizing, visibility, focus, and so on, and
	//  then use ICoreWebView2CompositionController to connect to a composition
	//  tree and provide input meant for the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller">See the ICoreWebView2CompositionController article.</see>
	CoreWebView2CompositionController() ICoreWebView2CompositionController // property CoreWebView2CompositionController Getter
	// CoreWebView2
	//  ICoreWebView2 wrapper used by this browser.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2">See the ICoreWebView2 article.</see>
	CoreWebView2() ICoreWebView2 // property CoreWebView2 Getter
	// Profile
	//  The associated `ICoreWebView2Profile` object. If this CoreWebView2 was created with a
	//  CoreWebView2ControllerOptions, the CoreWebView2Profile will match those specified options.
	//  Otherwise if this CoreWebView2 was created without a CoreWebView2ControllerOptions, then
	//  this will be the default CoreWebView2Profile for the corresponding CoreWebView2Environment.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_13#get_profile">See the ICoreWebView2_13 article.</see>
	Profile() ICoreWebView2Profile // property Profile Getter
	// DefaultURL
	//  First URL loaded by the browser after its creation.
	DefaultURL() string         // property DefaultURL Getter
	SetDefaultURL(value string) // property DefaultURL Setter
	// IsNavigating
	//  Returns true after OnNavigationStarting and before OnNavigationCompleted.
	IsNavigating() bool // property IsNavigating Getter
	// ZoomPct
	//  Returns the current zoom value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	ZoomPct() float64         // property ZoomPct Getter
	SetZoomPct(value float64) // property ZoomPct Setter
	// ZoomStep
	//  Returns the current zoom value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	ZoomStep() byte         // property ZoomStep Getter
	SetZoomStep(value byte) // property ZoomStep Setter
	// Widget0CompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	Widget0CompHWND() types.THandle // property Widget0CompHWND Getter
	// Widget1CompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	Widget1CompHWND() types.THandle // property Widget1CompHWND Getter
	// RenderCompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	RenderCompHWND() types.THandle // property RenderCompHWND Getter
	// D3DWindowCompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	D3DWindowCompHWND() types.THandle // property D3DWindowCompHWND Getter
	// ScreenScale
	//  Returns the GlobalWebView2Loader.DeviceScaleFactor value.
	ScreenScale() float32 // property ScreenScale Getter
	// Offline
	//  Uses the Network.emulateNetworkConditions DevTool method to set the browser in offline mode.
	//  The TWVBrowserBase.OnOfflineCompleted event is triggered asynchronously after setting this property.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions">See the Network Domain article.</see>
	Offline() bool         // property Offline Getter
	SetOffline(value bool) // property Offline Setter
	// IgnoreCertificateErrors
	//  Uses the Security.setIgnoreCertificateErrors DevTool method to enable/disable whether all certificate errors should be ignored.
	//  The TWVBrowserBase.OnIgnoreCertificateErrorsCompleted event is triggered asynchronously after setting this property.
	//  <see href="https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors">See the Security Domain article.</see>
	IgnoreCertificateErrors() bool         // property IgnoreCertificateErrors Getter
	SetIgnoreCertificateErrors(value bool) // property IgnoreCertificateErrors Setter
	// BrowserExecPath
	//  Properties used in the ICoreWebView2Environment creation
	BrowserExecPath() string         // property BrowserExecPath Getter
	SetBrowserExecPath(value string) // property BrowserExecPath Setter
	// UserDataFolder
	//  Returns the user data folder that all CoreWebView2's created from this
	//  environment are using.
	//  This could be either the value passed in by the developer when creating
	//  the environment object or the calculated one for default handling. It
	//  will always be an absolute path.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment7#get_userdatafolder">See the ICoreWebView2Environment7 article.</see>
	UserDataFolder() string         // property UserDataFolder Getter
	SetUserDataFolder(value string) // property UserDataFolder Setter
	// AdditionalBrowserArguments
	//  Additional command line switches.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AdditionalBrowserArguments.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</see>
	AdditionalBrowserArguments() string         // property AdditionalBrowserArguments Getter
	SetAdditionalBrowserArguments(value string) // property AdditionalBrowserArguments Setter
	// Language
	//  The default display language for WebView. It applies to browser UI such as
	//  context menu and dialogs. It also applies to the `accept-languages` HTTP
	//  header that WebView sends to websites.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_Language.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</see>
	Language() string         // property Language Getter
	SetLanguage(value string) // property Language Setter
	// TargetCompatibleBrowserVersion
	//  Specifies the version of the WebView2 Runtime binaries required to be
	//  compatible with your app.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_TargetCompatibleBrowserVersion.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</see>
	TargetCompatibleBrowserVersion() string         // property TargetCompatibleBrowserVersion Getter
	SetTargetCompatibleBrowserVersion(value string) // property TargetCompatibleBrowserVersion Setter
	// AllowSingleSignOnUsingOSPrimaryAccount
	//  Used to enable single sign on with Azure Active Directory (AAD) and personal Microsoft
	//  Account (MSA) resources inside WebView.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AllowSingleSignOnUsingOSPrimaryAccount.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</see>
	AllowSingleSignOnUsingOSPrimaryAccount() bool         // property AllowSingleSignOnUsingOSPrimaryAccount Getter
	SetAllowSingleSignOnUsingOSPrimaryAccount(value bool) // property AllowSingleSignOnUsingOSPrimaryAccount Setter
	// ExclusiveUserDataFolderAccess
	//  Whether other processes can create WebView2 from WebView2Environment created with the
	//  same user data folder and therefore sharing the same WebView browser process instance.
	//  Default is FALSE.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions2.Get_ExclusiveUserDataFolderAccess.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2">See the ICoreWebView2EnvironmentOptions2 article.</see>
	ExclusiveUserDataFolderAccess() bool         // property ExclusiveUserDataFolderAccess Getter
	SetExclusiveUserDataFolderAccess(value bool) // property ExclusiveUserDataFolderAccess Setter
	// CustomCrashReportingEnabled
	//  When `CustomCrashReportingEnabled` is set to `TRUE`, Windows won't send crash data to Microsoft endpoint.
	//  `CustomCrashReportingEnabled` is default to be `FALSE`, in this case, WebView will respect OS consent.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions3.Get_IsCustomCrashReportingEnabled.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3">See the ICoreWebView2EnvironmentOptions3 article.</see>
	CustomCrashReportingEnabled() bool         // property CustomCrashReportingEnabled Getter
	SetCustomCrashReportingEnabled(value bool) // property CustomCrashReportingEnabled Setter
	// EnableTrackingPrevention
	//  The `EnableTrackingPrevention` property is used to enable/disable tracking prevention
	//  feature in WebView2. This property enable/disable tracking prevention for all the
	//  WebView2's created in the same environment.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions5.Get_EnableTrackingPrevention.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5">See the ICoreWebView2EnvironmentOptions5 article.</see>
	EnableTrackingPrevention() bool         // property EnableTrackingPrevention Getter
	SetEnableTrackingPrevention(value bool) // property EnableTrackingPrevention Setter
	// AreBrowserExtensionsEnabled
	//  When `AreBrowserExtensionsEnabled` is set to `TRUE`, new extensions can be added to user
	//  profile and used. `AreBrowserExtensionsEnabled` is default to be `FALSE`, in this case,
	//  new extensions can't be installed, and already installed extension won't be
	//  available to use in user profile.
	//  If connecting to an already running environment with a different value for `AreBrowserExtensionsEnabled`
	//  property, it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions6.Get_AreBrowserExtensionsEnabled.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions6">See the ICoreWebView2EnvironmentOptions6 article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article for Extensions API details.</see>
	AreBrowserExtensionsEnabled() bool         // property AreBrowserExtensionsEnabled Getter
	SetAreBrowserExtensionsEnabled(value bool) // property AreBrowserExtensionsEnabled Setter
	// ChannelSearchKind
	//  The `ChannelSearchKind` property is `COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE`
	//  by default; environment creation searches for a release channel on the machine
	//  from most to least stable using the first channel found. The default search order is:
	//  WebView2 Runtime -&gt; Beta -&gt; Dev -&gt; Canary. Set `ChannelSearchKind` to
	//  `COREWEBVIEW2_CHANNEL_SEARCH_KIND_LEAST_STABLE` to reverse the search order
	//  so that environment creation searches for a channel from least to most stable. If
	//  `ReleaseChannels` has been provided, the loader will only search
	//  for channels in the set. See `COREWEBVIEW2_RELEASE_CHANNELS` for more details
	//  on channels.
	//  This property can be overridden by the corresponding
	//  registry key `ChannelSearchKind` or the environment variable
	//  `WEBVIEW2_CHANNEL_SEARCH_KIND`. Set the value to `1` to set the search kind to
	//  `COREWEBVIEW2_CHANNEL_SEARCH_KIND_LEAST_STABLE`. See
	//  `CreateCoreWebView2EnvironmentWithOptions` for more details on overrides.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions7.Get_ChannelSearchKind.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7">See the ICoreWebView2EnvironmentOptions7 article.</see>
	ChannelSearchKind() wvTypes.TWVChannelSearchKind         // property ChannelSearchKind Getter
	SetChannelSearchKind(value wvTypes.TWVChannelSearchKind) // property ChannelSearchKind Setter
	// ReleaseChannels
	//  Sets the `ReleaseChannels`, which is a mask of one or more
	//  `COREWEBVIEW2_RELEASE_CHANNELS` indicating which channels environment
	//  creation should search for. OR operation(s) can be applied to multiple
	//  `COREWEBVIEW2_RELEASE_CHANNELS` to create a mask. The default value is a
	//  a mask of all the channels. By default, environment creation searches for
	//  channels from most to least stable, using the first channel found on the
	//  device. When `ReleaseChannels` is provided, environment creation will only
	//  search for the channels specified in the set. Set `ChannelSearchKind` to
	//  `COREWEBVIEW2_CHANNEL_SEARCH_KIND_LEAST_STABLE` to reverse the search order
	//  so environment creation searches for least stable build first. See
	//  `COREWEBVIEW2_RELEASE_CHANNELS` for descriptions of each channel.
	//  `CreateCoreWebView2EnvironmentWithOptions` fails with
	//  `HRESULT_FROM_WIN32(ERROR_FILE_NOT_FOUND)` if environment creation is unable
	//  to find any channel from the `ReleaseChannels` installed on the device.
	//  Use `GetAvailableCoreWebView2BrowserVersionStringWithOptions` on
	//  `ICoreWebView2Environment` to verify which channel is used when this option
	//  is set.
	//  Examples:
	//  <code>
	//  | ReleaseChannels | Channel Search Kind: Most Stable (default) | Channel Search Kind: Least Stable |
	//  | --- | --- | --- |
	//  |COREWEBVIEW2_RELEASE_CHANNELS_BETA \| COREWEBVIEW2_RELEASE_CHANNELS_STABLE| WebView2 Runtime -&gt; Beta | Beta -&gt; WebView2 Runtime|
	//  |COREWEBVIEW2_RELEASE_CHANNELS_CANARY \| COREWEBVIEW2_RELEASE_CHANNELS_DEV \| COREWEBVIEW2_RELEASE_CHANNELS_BETA \| COREWEBVIEW2_RELEASE_CHANNELS_STABLE| WebView2 Runtime -&gt; Beta -&gt; Dev -&gt; Canary | Canary -&gt; Dev -&gt; Beta -&gt; WebView2 Runtime |
	//  |COREWEBVIEW2_RELEASE_CHANNELS_CANARY| Canary | Canary |
	//  |COREWEBVIEW2_RELEASE_CHANNELS_BETA \| COREWEBVIEW2_RELEASE_CHANNELS_CANARY \| COREWEBVIEW2_RELEASE_CHANNELS_STABLE | WebView2 Runtime -&gt; Beta -&gt; Canary | Canary -&gt; Beta -&gt; WebView2 Runtime |
	//  </code>
	//  If both `BrowserExecutableFolder` and `ReleaseChannels` are provided, the
	//  `BrowserExecutableFolder` takes precedence, regardless of whether or not the
	//  channel of `BrowserExecutableFolder` is included in the `ReleaseChannels`.
	//  `ReleaseChannels` can be overridden by the corresponding registry override
	//  `ReleaseChannels` or the environment variable `WEBVIEW2_RELEASE_CHANNELS`.
	//  Set the value to a comma-separated string of integers, which map to the
	//  following release channel values: Stable (0), Beta (1), Dev (2), and
	//  Canary (3). For example, the values "0,2" and "2,0" indicate that environment
	//  creation should only search for Dev channel and the WebView2 Runtime, using the
	//  order indicated by `ChannelSearchKind`. Environment creation attempts to
	//  interpret each integer and treats any invalid entry as Stable channel. See
	//  `CreateCoreWebView2EnvironmentWithOptions` for more details on overrides.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions7.Get_ReleaseChannels.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7">See the ICoreWebView2EnvironmentOptions7 article.</see>
	ReleaseChannels() wvTypes.TWVReleaseChannels         // property ReleaseChannels Getter
	SetReleaseChannels(value wvTypes.TWVReleaseChannels) // property ReleaseChannels Setter
	// ScrollBarStyle
	//  The ScrollBar style being set on the WebView2 Environment.
	//  The default value is `COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT`
	//  which specifies the default browser ScrollBar style.
	//  The `color-scheme` CSS property needs to be set on the corresponding page
	//  to allow ScrollBar to follow light or dark theme. Please see
	//  [color-scheme](https://developer.mozilla.org/docs/Web/CSS/color-scheme#declaring_color_scheme_preferences)
	//  for how `color-scheme` can be set.
	//  CSS styles that modify the ScrollBar applied on top of native ScrollBar styling
	//  that is selected with `ScrollBarStyle`.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions8.Get_ScrollBarStyle.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions8">See the ICoreWebView2EnvironmentOptions8 article.</see>
	ScrollBarStyle() wvTypes.TWVScrollBarStyle         // property ScrollBarStyle Getter
	SetScrollBarStyle(value wvTypes.TWVScrollBarStyle) // property ScrollBarStyle Setter
	// BrowserVersionInfo
	//  The browser version info of the current `ICoreWebView2Environment`,
	//  including channel name if it is not the WebView2 Runtime. It matches the
	//  format of the `GetAvailableCoreWebView2BrowserVersionString` API.
	//  Channel names are `beta`, `dev`, and `canary`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#get_browserversionstring">See the ICoreWebView2Environment article.</see>
	BrowserVersionInfo() string // property BrowserVersionInfo Getter
	// BrowserProcessID
	//  The process ID of the browser process that hosts the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_browserprocessid">See the ICoreWebView2 article.</see>
	BrowserProcessID() uint32 // property BrowserProcessID Getter
	// CanGoBack
	//  `TRUE` if the WebView is able to navigate to a previous page in the
	//  navigation history. If `CanGoBack` changes value, the `HistoryChanged`
	//  event runs.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2##get_cangoback">See the ICoreWebView2 article.</see>
	CanGoBack() bool // property CanGoBack Getter
	// CanGoForward
	//  `TRUE` if the WebView is able to navigate to a next page in the
	//  navigation history. If `CanGoForward` changes value, the
	//  `HistoryChanged` event runs.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_cangoforward">See the ICoreWebView2 article.</see>
	CanGoForward() bool // property CanGoForward Getter
	// ContainsFullScreenElement
	//  Indicates if the WebView contains a fullscreen HTML element.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_containsfullscreenelement">See the ICoreWebView2 article.</see>
	ContainsFullScreenElement() bool // property ContainsFullScreenElement Getter
	// DocumentTitle
	//  The title for the current top-level document. If the document has no
	//  explicit title or is otherwise empty, a default that may or may not match
	//  the URI of the document is used.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_documenttitle">See the ICoreWebView2 article.</see>
	DocumentTitle() string // property DocumentTitle Getter
	// Source
	//  The URI of the current top level document. This value potentially
	//  changes as a part of the `SourceChanged` event that runs for some cases
	//  such as navigating to a different site or fragment navigations. It
	//  remains the same for other types of navigations such as page refreshes
	//  or `history.pushState` with the same URL as the current page.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_source">See the ICoreWebView2 article.</see>
	Source() string // property Source Getter
	// CookieManager
	//  Gets the cookie manager object associated with this ICoreWebView2.
	//  See ICoreWebView2CookieManager.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#get_cookiemanager">See the ICoreWebView2_2 article.</see>
	CookieManager() ICoreWebView2CookieManager // property CookieManager Getter
	// IsSuspended
	//  Whether WebView is suspended.
	//  `TRUE` when WebView is suspended, from the time when TrySuspend has completed
	//  successfully until WebView is resumed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#get_issuspended">See the ICoreWebView2_3 article.</see>
	IsSuspended() bool // property IsSuspended Getter
	// IsDocumentPlayingAudio
	//  Indicates whether any audio output from this CoreWebView2 is playing.
	//  This property will be true if audio is playing even if IsMuted is true.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_isdocumentplayingaudio">See the ICoreWebView2_8 article.</see>
	IsDocumentPlayingAudio() bool // property IsDocumentPlayingAudio Getter
	// IsMuted
	//  Indicates whether all audio output from this CoreWebView2 is muted or not.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</see>
	IsMuted() bool         // property IsMuted Getter
	SetIsMuted(value bool) // property IsMuted Setter
	// DefaultDownloadDialogCornerAlignment
	//  Get the default download dialog corner alignment.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_defaultdownloaddialogcorneralignment">See the ICoreWebView2_9 article.</see>
	DefaultDownloadDialogCornerAlignment() wvTypes.TWVDefaultDownloadDialogCornerAlignment         // property DefaultDownloadDialogCornerAlignment Getter
	SetDefaultDownloadDialogCornerAlignment(value wvTypes.TWVDefaultDownloadDialogCornerAlignment) // property DefaultDownloadDialogCornerAlignment Setter
	// DefaultDownloadDialogMargin
	//  Get the default download dialog margin.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_defaultdownloaddialogmargin">See the ICoreWebView2_9 article.</see>
	DefaultDownloadDialogMargin() types.TPoint         // property DefaultDownloadDialogMargin Getter
	SetDefaultDownloadDialogMargin(value types.TPoint) // property DefaultDownloadDialogMargin Setter
	// IsDefaultDownloadDialogOpen
	//  `TRUE` if the default download dialog is currently open. The value of this
	//  property changes only when the default download dialog is explicitly
	//  opened or closed. Hiding the WebView implicitly hides the dialog, but does
	//  not change the value of this property.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_isdefaultdownloaddialogopen">See the ICoreWebView2_9 article.</see>
	IsDefaultDownloadDialogOpen() bool // property IsDefaultDownloadDialogOpen Getter
	// StatusBarText
	//  The status message text.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_12#get_statusbartext">See the ICoreWebView2_12 article.</see>
	StatusBarText() string // property StatusBarText Getter
	// FaviconURI
	//  Get the current Uri of the favicon as a string.
	//  If the value is null, then the return value is `E_POINTER`, otherwise it is `S_OK`.
	//  If a page has no favicon then the value is an empty string.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#get_faviconuri">See the ICoreWebView2_15 article.</see>
	FaviconURI() string // property FaviconURI Getter
	// MemoryUsageTargetLevel
	//  `MemoryUsageTargetLevel` indicates desired memory consumption level of
	//  WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_19#get_memoryusagetargetlevel">See the ICoreWebView2_19 article.</see>
	MemoryUsageTargetLevel() wvTypes.TWVMemoryUsageTargetLevel         // property MemoryUsageTargetLevel Getter
	SetMemoryUsageTargetLevel(value wvTypes.TWVMemoryUsageTargetLevel) // property MemoryUsageTargetLevel Setter
	// Bounds
	//  The WebView bounds. Bounds are relative to the parent `HWND`. The app
	//  has two ways to position a WebView.
	//  * Create a child `HWND` that is the WebView parent `HWND`. Position
	//  the window where the WebView should be. Use `(0, 0)` for the
	//  top-left corner (the offset) of the `Bounds` of the WebView.
	//  * Use the top-most window of the app as the WebView parent HWND. For
	//  example, to position WebView correctly in the app, set the top-left
	//  corner of the Bound of the WebView.
	//  The values of `Bounds` are limited by the coordinate space of the host.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_bounds">See the ICoreWebView2Controller article.</see>
	Bounds() types.TRect         // property Bounds Getter
	SetBounds(value types.TRect) // property Bounds Setter
	// IsVisible
	//  The `IsVisible` property determines whether to show or hide the WebView2.
	//  If `IsVisible` is set to `FALSE`, the WebView2 is transparent and is
	//  not rendered. However, this does not affect the window containing the
	//  WebView2 (the `HWND` parameter that was passed to
	//  `CreateCoreWebView2Controller`). If you want that window to disappear
	//  too, run `ShowWindow` on it directly in addition to modifying the
	//  `IsVisible` property. WebView2 as a child window does not get window
	//  messages when the top window is minimized or restored. For performance
	//  reasons, developers should set the `IsVisible` property of the WebView to
	//  `FALSE` when the app window is minimized and back to `TRUE` when the app
	//  window is restored. The app window does this by handling
	//  `SIZE_MINIMIZED and SIZE_RESTORED` command upon receiving `WM_SIZE`
	//  message.
	//  There are CPU and memory benefits when the page is hidden. For instance,
	//  Chromium has code that throttles activities on the page like animations
	//  and some tasks are run less frequently. Similarly, WebView2 will
	//  purge some caches to reduce memory usage.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_isvisible">See the ICoreWebView2Controller article.</see>
	IsVisible() bool         // property IsVisible Getter
	SetIsVisible(value bool) // property IsVisible Setter
	// ParentWindow
	//  The parent window provided by the app that this WebView is using to
	//  render content. This API initially returns the window passed into
	//  `CreateCoreWebView2Controller`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_parentwindow">See the ICoreWebView2Controller article.</see>
	//  <see href="https://github.com/salvadordf/WebView4Delphi/issues/13">See the WebView4Delphi issue #13 to know how to reparent a browser.</see>
	ParentWindow() types.THandle         // property ParentWindow Getter
	SetParentWindow(value types.THandle) // property ParentWindow Setter
	// ZoomFactor
	//  The zoom factor for the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	ZoomFactor() float64         // property ZoomFactor Getter
	SetZoomFactor(value float64) // property ZoomFactor Setter
	// DefaultBackgroundColor
	//  The `DefaultBackgroundColor` property is the color WebView renders
	//  underneath all web content. This means WebView renders this color when
	//  there is no web content loaded such as before the initial navigation or
	//  between navigations. This also means web pages with undefined css
	//  background properties or background properties containing transparent
	//  pixels will render their contents over this color. Web pages with defined
	//  and opaque background properties that span the page will obscure the
	//  `DefaultBackgroundColor` and display normally. The default value for this
	//  property is white to resemble the native browser experience.
	//  The Color is specified by the COREWEBVIEW2_COLOR that represents an RGBA
	//  value. The `A` represents an Alpha value, meaning
	//  `DefaultBackgroundColor` can be transparent. In the case of a transparent
	//  `DefaultBackgroundColor` WebView will render hosting app content as the
	//  background. This Alpha value is not supported on Windows 7. Any `A` value
	//  other than 255 will result in E_INVALIDARG on Windows 7.
	//  It is supported on all other WebView compatible platforms.
	//  Semi-transparent colors are not currently supported by this API and
	//  setting `DefaultBackgroundColor` to a semi-transparent color will fail
	//  with E_INVALIDARG. The only supported alpha values are 0 and 255, all
	//  other values will result in E_INVALIDARG.
	//  `DefaultBackgroundColor` can only be an opaque color or transparent.
	//  This value may also be set by using the
	//  `WEBVIEW2_DEFAULT_BACKGROUND_COLOR` environment variable. There is a
	//  known issue with background color where setting the color by API can
	//  still leave the app with a white flicker before the
	//  `DefaultBackgroundColor` takes effect. Setting the color via environment
	//  variable solves this issue. The value must be a hex value that can
	//  optionally prepend a 0x. The value must account for the alpha value
	//  which is represented by the first 2 digits. So any hex value fewer than 8
	//  digits will assume a prepended 00 to the hex value and result in a
	//  transparent color.
	//  `get_DefaultBackgroundColor` will return the result of this environment
	//  variable if used. This environment variable can only set the
	//  `DefaultBackgroundColor` once. Subsequent updates to background color
	//  must be done through API call.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller2#get_defaultbackgroundcolor">See the ICoreWebView2Controller2 article.</see>
	DefaultBackgroundColor() types.TColor         // property DefaultBackgroundColor Getter
	SetDefaultBackgroundColor(value types.TColor) // property DefaultBackgroundColor Setter
	// BoundsMode
	//  BoundsMode affects how setting the Bounds and RasterizationScale
	//  properties work. Bounds mode can either be in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS
	//  mode or COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE mode.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS, setting the bounds
	//  property will set the size of the WebView in raw screen pixels. Changing
	//  the rasterization scale in this mode won't change the raw pixel size of
	//  the WebView and will only change the rasterization scale.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE, setting the
	//  bounds property will change the logical size of the WebView which can be
	//  described by the following equation: Logical size * rasterization scale = Raw Pixel size
	//  In this case, changing the rasterization scale will keep the logical size
	//  the same and change the raw pixel size.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_boundsmode">See the ICoreWebView2Controller3 article.</see>
	BoundsMode() wvTypes.TWVBoundsMode         // property BoundsMode Getter
	SetBoundsMode(value wvTypes.TWVBoundsMode) // property BoundsMode Setter
	// RasterizationScale
	//  The rasterization scale for the WebView. The rasterization scale is the
	//  combination of the monitor DPI scale and text scaling set by the user.
	//  This value should be updated when the DPI scale of the app's top level
	//  window changes (i.e. monitor DPI scale changes or window changes monitor)
	//  or when the text scale factor of the system changes.
	//  Rasterization scale applies to the WebView content, as well as
	//  popups, context menus, scroll bars, and so on. Normal app scaling
	//  scenarios should use the ZoomFactor property or SetBoundsAndZoomFactor
	//  API which only scale the rendered HTML content and not popups, context
	//  menus, scroll bars, and so on.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_rasterizationscale">See the ICoreWebView2Controller3 article.</see>
	RasterizationScale() float64         // property RasterizationScale Getter
	SetRasterizationScale(value float64) // property RasterizationScale Setter
	// ShouldDetectMonitorScaleChanges
	//  ShouldDetectMonitorScaleChanges property determines whether the WebView
	//  attempts to track monitor DPI scale changes. When true, the WebView will
	//  track monitor DPI scale changes, update the RasterizationScale property,
	//  and raises RasterizationScaleChanged event. When false, the WebView will
	//  not track monitor DPI scale changes, and the app must update the
	//  RasterizationScale property itself. RasterizationScaleChanged event will
	//  never raise when ShouldDetectMonitorScaleChanges is false. Apps that want
	//  to set their own rasterization scale should set this property to false to
	//  avoid the WebView2 updating the RasterizationScale property to match the
	//  monitor DPI scale.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_shoulddetectmonitorscalechanges">See the ICoreWebView2Controller3 article.</see>
	ShouldDetectMonitorScaleChanges() bool         // property ShouldDetectMonitorScaleChanges Getter
	SetShouldDetectMonitorScaleChanges(value bool) // property ShouldDetectMonitorScaleChanges Setter
	// AllowExternalDrop
	//  Gets the `AllowExternalDrop` property which is used to configure the
	//  capability that dragging objects from outside the bounds of webview2 and
	//  dropping into webview2 is allowed or disallowed. The default value is
	//  TRUE.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller4#get_allowexternaldrop">See the ICoreWebView2Controller4 article.</see>
	AllowExternalDrop() bool         // property AllowExternalDrop Getter
	SetAllowExternalDrop(value bool) // property AllowExternalDrop Setter
	// DefaultContextMenusEnabled
	//  The `DefaultContextMenusEnabled` property is used to prevent default
	//  context menus from being shown to user in WebView.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultcontextmenusenabled">See the ICoreWebView2Settings article.</see>
	DefaultContextMenusEnabled() bool         // property DefaultContextMenusEnabled Getter
	SetDefaultContextMenusEnabled(value bool) // property DefaultContextMenusEnabled Setter
	// DefaultScriptDialogsEnabled
	//  `DefaultScriptDialogsEnabled` is used when loading a new HTML
	//  document. If set to `FALSE`, WebView2 does not render the default JavaScript
	//  dialog box (Specifically those displayed by the JavaScript alert,
	//  confirm, prompt functions and `beforeunload` event). Instead, if an
	//  event handler is set using `add_ScriptDialogOpening`, WebView sends an
	//  event that contains all of the information for the dialog and allow the
	//  host app to show a custom UI. The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultscriptdialogsenabled">See the ICoreWebView2Settings article.</see>
	DefaultScriptDialogsEnabled() bool         // property DefaultScriptDialogsEnabled Getter
	SetDefaultScriptDialogsEnabled(value bool) // property DefaultScriptDialogsEnabled Setter
	// DevToolsEnabled
	//  `DevToolsEnabled` controls whether the user is able to use the context
	//  menu or keyboard shortcuts to open the DevTools window.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredevtoolsenabled">See the ICoreWebView2Settings article.</see>
	DevToolsEnabled() bool         // property DevToolsEnabled Getter
	SetDevToolsEnabled(value bool) // property DevToolsEnabled Setter
	// AreHostObjectsAllowed
	//  The `AreHostObjectsAllowed` property is used to control whether host
	//  objects are accessible from the page in WebView.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_arehostobjectsallowed">See the ICoreWebView2Settings article.</see>
	AreHostObjectsAllowed() bool         // property AreHostObjectsAllowed Getter
	SetAreHostObjectsAllowed(value bool) // property AreHostObjectsAllowed Setter
	// BuiltInErrorPageEnabled
	//  The `BuiltInErrorPageEnabled` property is used to disable built in
	//  error page for navigation failure and render process failure. When
	//  disabled, a blank page is displayed when the related error happens.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isbuiltinerrorpageenabled">See the ICoreWebView2Settings article.</see>
	BuiltInErrorPageEnabled() bool         // property BuiltInErrorPageEnabled Getter
	SetBuiltInErrorPageEnabled(value bool) // property BuiltInErrorPageEnabled Setter
	// ScriptEnabled
	//  Controls if running JavaScript is enabled in all future navigations in
	//  the WebView. This only affects scripts in the document. Scripts
	//  injected with `ExecuteScript` runs even if script is disabled.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isscriptenabled">See the ICoreWebView2Settings article.</see>
	ScriptEnabled() bool         // property ScriptEnabled Getter
	SetScriptEnabled(value bool) // property ScriptEnabled Setter
	// StatusBarEnabled
	//  `StatusBarEnabled` controls whether the status bar is displayed. The
	//  status bar is usually displayed in the lower left of the WebView and
	//  shows things such as the URI of a link when the user hovers over it and
	//  other information. The default value is `TRUE`. The status bar UI can be
	//  altered by web content and should not be considered secure.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isstatusbarenabled">See the ICoreWebView2Settings article.</see>
	StatusBarEnabled() bool         // property StatusBarEnabled Getter
	SetStatusBarEnabled(value bool) // property StatusBarEnabled Setter
	// WebMessageEnabled
	//  The `WebMessageEnabled` property is used when loading a new HTML
	//  document. If set to `TRUE`, communication from the host to the top-level
	//  HTML document of the WebView is allowed using `PostWebMessageAsJson`,
	//  `PostWebMessageAsString`, and message event of `window.chrome.webview`.
	//  For more information, navigate to PostWebMessageAsJson. Communication
	//  from the top-level HTML document of the WebView to the host is allowed
	//  using the postMessage function of `window.chrome.webview` and
	//  `add_WebMessageReceived` method.
	//  If set to false, then communication is disallowed. `PostWebMessageAsJson`
	//  and `PostWebMessageAsString` fails with `E_ACCESSDENIED` and
	//  `window.chrome.webview.postMessage` fails by throwing an instance of an
	//  `Error` object. The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived">See the add_WebMessageReceived method article.</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iswebmessageenabled">See the ICoreWebView2Settings article.</see>
	WebMessageEnabled() bool         // property WebMessageEnabled Getter
	SetWebMessageEnabled(value bool) // property WebMessageEnabled Setter
	// ZoomControlEnabled
	//  The `ZoomControlEnabled` property is used to prevent the user from
	//  impacting the zoom of the WebView. When disabled, the user is not able
	//  to zoom using Ctrl++, Ctrl+-, or Ctrl+mouse wheel, but the zoom
	//  is set using `ZoomFactor` API. The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iszoomcontrolenabled">See the ICoreWebView2Settings article.</see>
	ZoomControlEnabled() bool         // property ZoomControlEnabled Getter
	SetZoomControlEnabled(value bool) // property ZoomControlEnabled Setter
	// UserAgent
	//  Returns the User Agent. The default value is the default User Agent of the
	//  Microsoft Edge browser.
	//  This property may be overridden if
	//  the User-Agent header is set in a request. If the parameter is empty
	//  the User Agent will not be updated and the current User Agent will remain.
	//  Setting this property may clear User Agent Client Hints headers
	//  Sec-CH-UA-* and script values from navigator.userAgentData. Current
	//  implementation behavior is subject to change.
	//  The User Agent set will also be effective on service workers
	//  and shared workers associated with the WebView. If there are
	//  multiple WebViews associated with the same service worker or
	//  shared worker, the last User Agent set will be used.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings2#get_useragent">See the ICoreWebView2Settings2 article.</see>
	UserAgent() string         // property UserAgent Getter
	SetUserAgent(value string) // property UserAgent Setter
	// AreBrowserAcceleratorKeysEnabled
	//  When this setting is set to FALSE, it disables all accelerator keys that
	//  access features specific to a web browser.
	//  The default value for `AreBrowserAcceleratorKeysEnabled` is TRUE.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings3#get_arebrowseracceleratorkeysenabled">See the ICoreWebView2Settings3 article.</see>
	AreBrowserAcceleratorKeysEnabled() bool         // property AreBrowserAcceleratorKeysEnabled Getter
	SetAreBrowserAcceleratorKeysEnabled(value bool) // property AreBrowserAcceleratorKeysEnabled Setter
	// IsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields. It will take effect
	//  immediately after setting. The default value is `TRUE`.
	//  This property has the same value as
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled` and
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled` will always have the same
	//  value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings4#get_isgeneralautofillenabled">See the ICoreWebView2Settings4 article.</see>
	IsGeneralAutofillEnabled() bool         // property IsGeneralAutofillEnabled Getter
	SetIsGeneralAutofillEnabled(value bool) // property IsGeneralAutofillEnabled Setter
	// IsPasswordAutosaveEnabled
	//  IsPasswordAutosaveEnabled controls whether autosave for password
	//  information is enabled. The IsPasswordAutosaveEnabled property behaves
	//  independently of the IsGeneralAutofillEnabled property. When IsPasswordAutosaveEnabled is
	//  false, no new password data is saved and no Save/Update Password prompts are displayed.
	//  However, if there was password data already saved before disabling this setting,
	//  then that password information is auto-populated, suggestions are shown and clicking on
	//  one will populate the fields.
	//  When IsPasswordAutosaveEnabled is true, password information is auto-populated,
	//  suggestions are shown and clicking on one will populate the fields, new data
	//  is saved, and a Save/Update Password prompt is displayed.
	//  It will take effect immediately after setting. The default value is `FALSE`.
	//  This property has the same value as
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings4#get_ispasswordautosaveenabled">See the ICoreWebView2Settings4 article.</see>
	IsPasswordAutosaveEnabled() bool         // property IsPasswordAutosaveEnabled Getter
	SetIsPasswordAutosaveEnabled(value bool) // property IsPasswordAutosaveEnabled Setter
	// IsPinchZoomEnabled
	//  Pinch-zoom, referred to as "Page Scale" zoom, is performed as a post-rendering step,
	//  it changes the page scale factor property and scales the surface the web page is
	//  rendered onto when user performs a pinch zooming action. It does not change the layout
	//  but rather changes the viewport and clips the web content, the content outside of the
	//  viewport isn't visible onscreen and users can't reach this content using mouse.
	//  The `IsPinchZoomEnabled` property enables or disables the ability of
	//  the end user to use a pinching motion on touch input enabled devices
	//  to scale the web content in the WebView2. It defaults to `TRUE`.
	//  When set to `FALSE`, the end user cannot pinch zoom after the next navigation.
	//  Disabling/Enabling `IsPinchZoomEnabled` only affects the end user's ability to use
	//  pinch motions and does not change the page scale factor.
	//  This API only affects the Page Scale zoom and has no effect on the
	//  existing browser zoom properties (`IsZoomControlEnabled` and `ZoomFactor`)
	//  or other end user mechanisms for zooming.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings5#get_ispinchzoomenabled">See the ICoreWebView2Settings5 article.</see>
	IsPinchZoomEnabled() bool         // property IsPinchZoomEnabled Getter
	SetIsPinchZoomEnabled(value bool) // property IsPinchZoomEnabled Setter
	// IsSwipeNavigationEnabled
	//  The `IsSwipeNavigationEnabled` property enables or disables the ability of the
	//  end user to use swiping gesture on touch input enabled devices to
	//  navigate in WebView2. It defaults to `TRUE`.
	//  When this property is `TRUE`, then all configured navigation gestures are enabled:
	//  1. Swiping left and right to navigate forward and backward is always configured.
	//  2. Swiping down to refresh is off by default and not exposed via our API currently,
	//  it requires the "--pull-to-refresh" option to be included in the additional browser
	//  arguments to be configured. (See put_AdditionalBrowserArguments.)
	//  When set to `FALSE`, the end user cannot swipe to navigate or pull to refresh.
	//  This API only affects the overscrolling navigation functionality and has no
	//  effect on the scrolling interaction used to explore the web content shown
	//  in WebView2.
	//  Disabling/Enabling IsSwipeNavigationEnabled takes effect after the
	//  next navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings6#get_isswipenavigationenabled">See the ICoreWebView2Settings6 article.</see>
	IsSwipeNavigationEnabled() bool         // property IsSwipeNavigationEnabled Getter
	SetIsSwipeNavigationEnabled(value bool) // property IsSwipeNavigationEnabled Setter
	// HiddenPdfToolbarItems
	//  `HiddenPdfToolbarItems` is used to customize the PDF toolbar items. By default, it is COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE and so it displays all of the items.
	//  Changes to this property apply to all CoreWebView2s in the same environment and using the same profile.
	//  Changes to this setting apply only after the next navigation.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings7#get_hiddenpdftoolbaritems">See the ICoreWebView2Settings7 article.</see>
	HiddenPdfToolbarItems() wvTypes.TWVPDFToolbarItems         // property HiddenPdfToolbarItems Getter
	SetHiddenPdfToolbarItems(value wvTypes.TWVPDFToolbarItems) // property HiddenPdfToolbarItems Setter
	// IsReputationCheckingRequired
	//  SmartScreen helps webviews identify reported phishing and malware websites
	//  and also helps users make informed decisions about downloads.
	//  `IsReputationCheckingRequired` is used to control whether SmartScreen
	//  enabled or not. SmartScreen is enabled or disabled for all CoreWebView2s
	//  using the same user data folder. If
	//  CoreWebView2Setting.IsReputationCheckingRequired is true for any
	//  CoreWebView2 using the same user data folder, then SmartScreen is enabled.
	//  If CoreWebView2Setting.IsReputationCheckingRequired is false for all
	//  CoreWebView2 using the same user data folder, then SmartScreen is
	//  disabled. When it is changed, the change will be applied to all WebViews
	//  using the same user data folder on the next navigation or download. The
	//  default value for `IsReputationCheckingRequired` is true. If the newly
	//  created CoreWebview2 does not set SmartScreen to false, when
	//  navigating(Such as Navigate(), LoadDataUrl(), ExecuteScript(), etc.), the
	//  default value will be applied to all CoreWebview2 using the same user data
	//  folder.
	//  SmartScreen of WebView2 apps can be controlled by Windows system setting
	//  "SmartScreen for Microsoft Edge", specially, for WebView2 in Windows
	//  Store apps, SmartScreen is controlled by another Windows system setting
	//  "SmartScreen for Microsoft Store apps". When the Windows setting is enabled, the
	//  SmartScreen operates under the control of the `IsReputationCheckingRequired`.
	//  When the Windows setting is disabled, the SmartScreen will be disabled
	//  regardless of the `IsReputationCheckingRequired` value set in WebView2 apps.
	//  In other words, under this circumstance the value of
	//  `IsReputationCheckingRequired` will be saved but overridden by system setting.
	//  Upon re-enabling the Windows setting, the CoreWebview2 will reference the
	//  `IsReputationCheckingRequired` to determine the SmartScreen status.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings8#get_isreputationcheckingrequired">See the ICoreWebView2Settings8 article.</see>
	IsReputationCheckingRequired() bool         // property IsReputationCheckingRequired Getter
	SetIsReputationCheckingRequired(value bool) // property IsReputationCheckingRequired Setter
	// IsNonClientRegionSupportEnabled
	//  The `IsNonClientRegionSupportEnabled` property enables web pages to use the
	//  `app-region` CSS style. Disabling/Enabling the `IsNonClientRegionSupportEnabled`
	//  takes effect after the next navigation. Defaults to `FALSE`.
	//  When this property is `TRUE`, then all the non-client region features
	//  will be enabled:
	//  Draggable Regions will be enabled, they are regions on a webpage that
	//  are marked with the CSS attribute `app-region: drag/no-drag`. When set to
	//  `drag`, these regions will be treated like the window's title bar, supporting
	//  dragging of the entire WebView and its host app window; the system menu shows
	//  upon right click, and a double click will trigger maximizing/restoration of the
	//  window size.
	//  When set to `FALSE`, all non-client region support will be disabled.
	//  The `app-region` CSS style will be ignored on web pages.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings9#get_isnonclientregionsupportenabled">See the ICoreWebView2Settings9 article.</see>
	IsNonClientRegionSupportEnabled() bool         // property IsNonClientRegionSupportEnabled Getter
	SetIsNonClientRegionSupportEnabled(value bool) // property IsNonClientRegionSupportEnabled Setter
	// Cursor
	//  The current cursor that WebView thinks it should be. The cursor should be
	//  set in WM_SETCURSOR through \::SetCursor or set on the corresponding
	//  parent/ancestor HWND of the WebView through \::SetClassLongPtr. The HCURSOR
	//  can be freed so CopyCursor/DestroyCursor is recommended to keep your own
	//  copy if you are doing more than immediately setting the cursor.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_cursor">See the ICoreWebView2CompositionController article.</see>
	Cursor() types.HCURSOR // property Cursor Getter
	// RootVisualTarget
	//  The RootVisualTarget is a visual in the hosting app's visual tree. This
	//  visual is where the WebView will connect its visual tree. The app uses
	//  this visual to position the WebView within the app. The app still needs
	//  to use the Bounds property to size the WebView. The RootVisualTarget
	//  property can be an IDCompositionVisual or a
	//  Windows::UI::Composition::ContainerVisual. WebView will connect its visual
	//  tree to the provided visual before returning from the property setter. The
	//  app needs to commit on its device setting the RootVisualTarget property.
	//  The RootVisualTarget property supports being set to nullptr to disconnect
	//  the WebView from the app's visual tree.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_rootvisualtarget">See the ICoreWebView2CompositionController article.</see>
	RootVisualTarget() lcl.IUnknown         // property RootVisualTarget Getter
	SetRootVisualTarget(value lcl.IUnknown) // property RootVisualTarget Setter
	// SystemCursorID
	//  The current system cursor ID reported by the underlying rendering engine
	//  for WebView. For example, most of the time, when the cursor is over text,
	//  this will return the int value for IDC_IBEAM. The systemCursorId is only
	//  valid if the rendering engine reports a default Windows cursor resource
	//  value. Navigate to
	//  [LoadCursorW](/windows/win32/api/winuser/nf-winuser-loadcursorw) for more
	//  details. Otherwise, if custom CSS cursors are being used, this will return
	//  0. To actually use systemCursorId in LoadCursor or LoadImage,
	//  MAKEINTRESOURCE must be called on it first.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_systemcursorid">See the ICoreWebView2CompositionController article.</see>
	SystemCursorID() uint32 // property SystemCursorID Getter
	// AutomationProvider
	//  Returns the Automation Provider for the WebView. This object implements
	//  IRawElementProviderSimple.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller2#get_automationprovider">See the ICoreWebView2CompositionController2 article.</see>
	AutomationProvider() lcl.IUnknown // property AutomationProvider Getter
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`. Provide a list of all
	//  process using same user data folder except for crashpad process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8#getprocessinfos">See the ICoreWebView2Environment8 article.</see>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property ProcessInfos Getter
	// ProfileName
	//  `ProfileName` property is to specify a profile name, which is only allowed to contain
	//  the following ASCII characters. It has a maximum length of 64 characters excluding the null-terminator.
	//  It is ASCII case insensitive.
	//  * alphabet characters: a-z and A-Z
	//  * digit characters: 0-9
	//  * and '#', '@', '$', '(', ')', '+', '-', '_', '~', '.', ' ' (space).
	//  Note: the text must not end with a period '.' or ' ' (space). And, although upper-case letters are
	//  allowed, they're treated just as lower-case counterparts because the profile name will be mapped to
	//  the real profile directory path on disk and Windows file system handles path names in a case-insensitive way.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_profilename">See the ICoreWebView2ControllerOptions article.</see>
	ProfileName() string         // property ProfileName Getter
	SetProfileName(value string) // property ProfileName Setter
	// IsInPrivateModeEnabled
	//  `IsInPrivateModeEnabled` property is to enable/disable InPrivate mode.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_isinprivatemodeenabled">See the ICoreWebView2ControllerOptions article.</see>
	IsInPrivateModeEnabled() bool         // property IsInPrivateModeEnabled Getter
	SetIsInPrivateModeEnabled(value bool) // property IsInPrivateModeEnabled Setter
	// ScriptLocale
	//  The default locale for the WebView2. It sets the default locale for all
	//  Intl JavaScript APIs and other JavaScript APIs that depend on it, namely
	//  `Intl.DateTimeFormat()` which affects string formatting like
	//  in the time/date formats. Example: `Intl.DateTimeFormat().format(new Date())`
	//  The intended locale value is in the format of
	//  BCP 47 Language Tags. More information can be found from
	//  [IETF BCP47](https://www.ietf.org/rfc/bcp/bcp47.html).
	//  This property sets the locale for a CoreWebView2Environment used to create the
	//  WebView2ControllerOptions object, which is passed as a parameter in
	//  `CreateCoreWebView2ControllerWithOptions`.
	//  Changes to the ScriptLocale property apply to renderer processes created after
	//  the change. Any existing renderer processes will continue to use the previous
	//  ScriptLocale value. To ensure changes are applied to all renderer process,
	//  close and restart the CoreWebView2Environment and all associated WebView2 objects.
	//  The default value for ScriptLocale will depend on the WebView2 language
	//  and OS region. If the language portions of the WebView2 language and OS region
	//  match, then it will use the OS region. Otherwise, it will use the WebView2
	//  language.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2#get_scriptlocale">See the ICoreWebView2ControllerOptions2 article.</see>
	ScriptLocale() string         // property ScriptLocale Getter
	SetScriptLocale(value string) // property ScriptLocale Setter
	// ProfilePath
	//  Full path of the profile directory.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_profilepath">See the ICoreWebView2Profile article.</see>
	ProfilePath() string // property ProfilePath Getter
	// DefaultDownloadFolderPath
	//  Gets the `DefaultDownloadFolderPath` property. The default value is the
	//  system default download folder path for the user.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_defaultdownloadfolderpath">See the ICoreWebView2Profile article.</see>
	DefaultDownloadFolderPath() string         // property DefaultDownloadFolderPath Getter
	SetDefaultDownloadFolderPath(value string) // property DefaultDownloadFolderPath Setter
	// PreferredColorScheme
	//  The PreferredColorScheme property sets the overall color scheme of the
	//  WebView2s associated with this profile. This sets the color scheme for
	//  WebView2 UI like dialogs, prompts, and context menus by setting the
	//  media feature `prefers-color-scheme` for websites to respond to.
	//  The default value for this is COREWEBVIEW2_PREFERRED_COLOR_AUTO,
	//  which will follow whatever theme the OS is currently set to.
	//  Returns the value of the `PreferredColorScheme` property.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_preferredcolorscheme">See the ICoreWebView2Profile article.</see>
	PreferredColorScheme() wvTypes.TWVPreferredColorScheme         // property PreferredColorScheme Getter
	SetPreferredColorScheme(value wvTypes.TWVPreferredColorScheme) // property PreferredColorScheme Setter
	// PreferredTrackingPreventionLevel
	//  The `PreferredTrackingPreventionLevel` property allows you to control levels of tracking prevention for WebView2
	//  which are associated with a profile. This level would apply to the context of the profile. That is, all WebView2s
	//  sharing the same profile will be affected and also the value is persisted in the user data folder.
	//  See `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL` for descriptions of levels.
	//  If tracking prevention feature is enabled when creating the WebView2 environment, you can also disable tracking
	//  prevention later using this property and `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE` value but that doesn't
	//  improves runtime performance.
	//  There is `ICoreWebView2EnvironmentOptions5.EnableTrackingPrevention` property to enable/disable tracking prevention feature
	//  for all the WebView2's created in the same environment. If enabled, `PreferredTrackingPreventionLevel` is set to
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED` by default for all the WebView2's and profiles created in the same
	//  environment or is set to the level whatever value was last changed/persisted to the profile. If disabled
	//  `PreferredTrackingPreventionLevel` is not respected by WebView2. If `PreferredTrackingPreventionLevel` is set when the
	//  feature is disabled, the property value get changed and persisted but it will takes effect only if
	//  `ICoreWebView2EnvironmentOptions5.EnableTrackingPrevention` is true.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile3#get_preferredtrackingpreventionlevel">See the ICoreWebView2Profile3 article.</see>
	PreferredTrackingPreventionLevel() wvTypes.TWVTrackingPreventionLevel         // property PreferredTrackingPreventionLevel Getter
	SetPreferredTrackingPreventionLevel(value wvTypes.TWVTrackingPreventionLevel) // property PreferredTrackingPreventionLevel Setter
	// ProfileCookieManager
	//  Get the cookie manager for the profile. All CoreWebView2s associated with this
	//  profile share the same cookie values. Changes to cookies in this cookie manager apply to all
	//  CoreWebView2s associated with this profile. See ICoreWebView2CookieManager.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile5#get_cookiemanager">See the ICoreWebView2Profile5 article.</see>
	ProfileCookieManager() ICoreWebView2CookieManager // property ProfileCookieManager Getter
	// ProfileIsPasswordAutosaveEnabled
	//  IsPasswordAutosaveEnabled controls whether autosave for password
	//  information is enabled. The IsPasswordAutosaveEnabled property behaves
	//  independently of the IsGeneralAutofillEnabled property. When IsPasswordAutosaveEnabled is
	//  false, no new password data is saved and no Save/Update Password prompts are displayed.
	//  However, if there was password data already saved before disabling this setting,
	//  then that password information is auto-populated, suggestions are shown and clicking on
	//  one will populate the fields.
	//  When IsPasswordAutosaveEnabled is true, password information is auto-populated,
	//  suggestions are shown and clicking on one will populate the fields, new data
	//  is saved, and a Save/Update Password prompt is displayed.
	//  It will take effect immediately after setting. The default value is `FALSE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_ispasswordautosaveenabled">See the ICoreWebView2Profile6 article.</see>
	ProfileIsPasswordAutosaveEnabled() bool         // property ProfileIsPasswordAutosaveEnabled Getter
	SetProfileIsPasswordAutosaveEnabled(value bool) // property ProfileIsPasswordAutosaveEnabled Setter
	// ProfileIsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields.
	//  It will take effect immediately after setting. The default value is `TRUE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled` and
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled` will always have the same
	//  value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_isgeneralautofillenabled">See the ICoreWebView2Profile6 article.</see>
	ProfileIsGeneralAutofillEnabled() bool         // property ProfileIsGeneralAutofillEnabled Getter
	SetProfileIsGeneralAutofillEnabled(value bool) // property ProfileIsGeneralAutofillEnabled Setter
	// FrameId
	//  The unique identifier of the main frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2Frame` and via `ICoreWebView2FrameInfo`.
	//  Note that `FrameId` may not be valid if `ICoreWebView2` has not done
	//  any navigation. It's safe to get this value during or after the first
	//  `ContentLoading` event. Otherwise, it could return the invalid frame Id 0.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_20#get_frameid">See the ICoreWebView2_20 article.</see>
	FrameId() uint32                                                                                           // property FrameId Getter
	SetOnBrowserProcessExited(fn TOnBrowserProcessExitedEvent)                                                 // property event
	SetOnProcessInfosChanged(fn TOnProcessInfosChangedEvent)                                                   // property event
	SetOnContainsFullScreenElementChanged(fn TNotifyEvent)                                                     // property event
	SetOnContentLoading(fn TOnContentLoadingEvent)                                                             // property event
	SetOnDocumentTitleChanged(fn TNotifyEvent)                                                                 // property event
	SetOnFrameNavigationCompleted(fn TOnNavigationCompletedEvent)                                              // property event
	SetOnFrameNavigationStarting(fn TOnNavigationStartingEvent)                                                // property event
	SetOnHistoryChanged(fn TNotifyEvent)                                                                       // property event
	SetOnNavigationCompleted(fn TOnNavigationCompletedEvent)                                                   // property event
	SetOnNavigationStarting(fn TOnNavigationStartingEvent)                                                     // property event
	SetOnNewWindowRequested(fn TOnNewWindowRequestedEvent)                                                     // property event
	SetOnPermissionRequested(fn TOnPermissionRequestedEvent)                                                   // property event
	SetOnProcessFailed(fn TOnProcessFailedEvent)                                                               // property event
	SetOnScriptDialogOpening(fn TOnScriptDialogOpeningEvent)                                                   // property event
	SetOnSourceChanged(fn TOnSourceChangedEvent)                                                               // property event
	SetOnWebMessageReceived(fn TOnWebMessageReceivedEvent)                                                     // property event
	SetOnWebResourceRequested(fn TOnWebResourceRequestedEvent)                                                 // property event
	SetOnWindowCloseRequested(fn TNotifyEvent)                                                                 // property event
	SetOnDOMContentLoaded(fn TOnDOMContentLoadedEvent)                                                         // property event
	SetOnWebResourceResponseReceived(fn TOnWebResourceResponseReceivedEvent)                                   // property event
	SetOnDownloadStarting(fn TOnDownloadStartingEvent)                                                         // property event
	SetOnFrameCreated(fn TOnFrameCreatedEvent)                                                                 // property event
	SetOnClientCertificateRequested(fn TOnClientCertificateRequestedEvent)                                     // property event
	SetOnIsDocumentPlayingAudioChanged(fn TOnIsDocumentPlayingAudioChangedEvent)                               // property event
	SetOnIsMutedChanged(fn TOnIsMutedChangedEvent)                                                             // property event
	SetOnIsDefaultDownloadDialogOpenChanged(fn TOnIsDefaultDownloadDialogOpenChangedEvent)                     // property event
	SetOnBasicAuthenticationRequested(fn TOnBasicAuthenticationRequestedEvent)                                 // property event
	SetOnContextMenuRequested(fn TOnContextMenuRequestedEvent)                                                 // property event
	SetOnStatusBarTextChanged(fn TOnStatusBarTextChangedEvent)                                                 // property event
	SetOnServerCertificateErrorActionsCompleted(fn TOnServerCertificateErrorActionsCompletedEvent)             // property event
	SetOnServerCertificateErrorDetected(fn TOnServerCertificateErrorDetectedEvent)                             // property event
	SetOnFaviconChanged(fn TOnFaviconChangedEvent)                                                             // property event
	SetOnGetFaviconCompleted(fn TOnGetFaviconCompletedEvent)                                                   // property event
	SetOnPrintCompleted(fn TOnPrintCompletedEvent)                                                             // property event
	SetOnPrintToPdfStreamCompleted(fn TOnPrintToPdfStreamCompletedEvent)                                       // property event
	SetOnAcceleratorKeyPressed(fn TOnAcceleratorKeyPressedEvent)                                               // property event
	SetOnGotFocus(fn TNotifyEvent)                                                                             // property event
	SetOnLostFocus(fn TNotifyEvent)                                                                            // property event
	SetOnMoveFocusRequested(fn TOnMoveFocusRequestedEvent)                                                     // property event
	SetOnZoomFactorChanged(fn TNotifyEvent)                                                                    // property event
	SetOnRasterizationScaleChanged(fn TNotifyEvent)                                                            // property event
	SetOnCursorChanged(fn TNotifyEvent)                                                                        // property event
	SetOnBytesReceivedChanged(fn TOnBytesReceivedChangedEvent)                                                 // property event
	SetOnEstimatedEndTimeChanged(fn TOnEstimatedEndTimeChangedEvent)                                           // property event
	SetOnDownloadStateChanged(fn TOnDownloadStateChangedEvent)                                                 // property event
	SetOnFrameDestroyed(fn TOnFrameDestroyedEvent)                                                             // property event
	SetOnFrameNameChanged(fn TOnFrameNameChangedEvent)                                                         // property event
	SetOnFrameNavigationStarting2(fn TOnFrameNavigationStartingEvent)                                          // property event
	SetOnFrameNavigationCompleted2(fn TOnFrameNavigationCompletedEvent)                                        // property event
	SetOnFrameContentLoading(fn TOnFrameContentLoadingEvent)                                                   // property event
	SetOnFrameDOMContentLoaded(fn TOnFrameDOMContentLoadedEvent)                                               // property event
	SetOnFrameWebMessageReceived(fn TOnFrameWebMessageReceivedEvent)                                           // property event
	SetOnFramePermissionRequested(fn TOnFramePermissionRequestedEvent)                                         // property event
	SetOnDevToolsProtocolEventReceived(fn TOnDevToolsProtocolEventReceivedEvent)                               // property event
	SetOnCustomItemSelected(fn TOnCustomItemSelectedEvent)                                                     // property event
	SetOnClearBrowsingDataCompleted(fn TOnClearBrowsingDataCompletedEvent)                                     // property event
	SetOnInitializationError(fn TOnInitializationErrorEvent)                                                   // property event
	SetOnEnvironmentCompleted(fn TNotifyEvent)                                                                 // property event
	SetOnControllerCompleted(fn TNotifyEvent)                                                                  // property event
	SetOnAfterCreated(fn TNotifyEvent)                                                                         // property event
	SetOnExecuteScriptCompleted(fn TOnExecuteScriptCompletedEvent)                                             // property event
	SetOnCapturePreviewCompleted(fn TOnCapturePreviewCompletedEvent)                                           // property event
	SetOnGetCookiesCompleted(fn TOnGetCookiesCompletedEvent)                                                   // property event
	SetOnTrySuspendCompleted(fn TOnTrySuspendCompletedEvent)                                                   // property event
	SetOnPrintToPdfCompleted(fn TOnPrintToPdfCompletedEvent)                                                   // property event
	SetOnCompositionControllerCompleted(fn TNotifyEvent)                                                       // property event
	SetOnCallDevToolsProtocolMethodCompleted(fn TOnCallDevToolsProtocolMethodCompletedEvent)                   // property event
	SetOnAddScriptToExecuteOnDocumentCreatedCompleted(fn TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent) // property event
	SetOnWebResourceResponseViewGetContentCompleted(fn TOnWebResourceResponseViewGetContentCompletedEvent)     // property event
	SetOnWidget0CompMsg(fn TOnCompMsgEvent)                                                                    // property event
	SetOnWidget1CompMsg(fn TOnCompMsgEvent)                                                                    // property event
	SetOnRenderCompMsg(fn TOnCompMsgEvent)                                                                     // property event
	SetOnD3DWindowCompMsg(fn TOnCompMsgEvent)                                                                  // property event
	SetOnRetrieveHTMLCompleted(fn TOnRetrieveHTMLCompletedEvent)                                               // property event
	SetOnRetrieveTextCompleted(fn TOnRetrieveTextCompletedEvent)                                               // property event
	SetOnRetrieveMHTMLCompleted(fn TOnRetrieveMHTMLCompletedEvent)                                             // property event
	SetOnClearCacheCompleted(fn TOnClearCacheCompletedEvent)                                                   // property event
	SetOnClearDataForOriginCompleted(fn TOnClearDataForOriginCompletedEvent)                                   // property event
	SetOnOfflineCompleted(fn TOnOfflineCompletedEvent)                                                         // property event
	SetOnIgnoreCertificateErrorsCompleted(fn TOnIgnoreCertificateErrorsCompletedEvent)                         // property event
	SetOnRefreshIgnoreCacheCompleted(fn TOnRefreshIgnoreCacheCompletedEvent)                                   // property event
	SetOnSimulateKeyEventCompleted(fn TOnSimulateKeyEventCompletedEvent)                                       // property event
	SetOnGetCustomSchemes(fn TOnGetCustomSchemesEvent)                                                         // property event
	SetOnGetNonDefaultPermissionSettingsCompleted(fn TOnGetNonDefaultPermissionSettingsCompletedEvent)         // property event
	SetOnSetPermissionStateCompleted(fn TOnSetPermissionStateCompletedEvent)                                   // property event
	SetOnLaunchingExternalUriScheme(fn TOnLaunchingExternalUriSchemeEvent)                                     // property event
	SetOnGetProcessExtendedInfosCompleted(fn TOnGetProcessExtendedInfosCompletedEvent)                         // property event
	SetOnBrowserExtensionRemoveCompleted(fn TOnBrowserExtensionRemoveCompletedEvent)                           // property event
	SetOnBrowserExtensionEnableCompleted(fn TOnBrowserExtensionEnableCompletedEvent)                           // property event
	SetOnProfileAddBrowserExtensionCompleted(fn TOnProfileAddBrowserExtensionCompletedEvent)                   // property event
	SetOnProfileGetBrowserExtensionsCompleted(fn TOnProfileGetBrowserExtensionsCompletedEvent)                 // property event
	SetOnProfileDeleted(fn TOnProfileDeletedEvent)                                                             // property event
	SetOnExecuteScriptWithResultCompleted(fn TOnExecuteScriptWithResultCompletedEvent)                         // property event
	SetOnNonClientRegionChanged(fn TOnNonClientRegionChangedEvent)                                             // property event
	SetOnNotificationReceived(fn TOnNotificationReceivedEvent)                                                 // property event
	SetOnNotificationCloseRequested(fn TOnNotificationCloseRequestedEvent)                                     // property event
	SetOnSaveAsUIShowing(fn TOnSaveAsUIShowingEvent)                                                           // property event
	SetOnShowSaveAsUICompleted(fn TOnShowSaveAsUICompletedEvent)                                               // property event
	SetOnSaveFileSecurityCheckStarting(fn TOnSaveFileSecurityCheckStartingEvent)                               // property event
	SetOnScreenCaptureStarting(fn TOnScreenCaptureStartingEvent)                                               // property event
	SetOnFrameScreenCaptureStarting(fn TOnFrameScreenCaptureStartingEvent)                                     // property event
}

type TWVBrowserBase struct {
	TComponent
}

func (m *TWVBrowserBase) CreateBrowserWithHandleBool(handle types.THandle, useDefaultEnvironment bool) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(0, m.Instance(), uintptr(handle), api.PasBool(useDefaultEnvironment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateBrowserWithHandleEnvironment(handle types.THandle, environment ICoreWebView2Environment) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(1, m.Instance(), uintptr(handle), base.GetObjectUintptr(environment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateWindowlessBrowserWithHandleBool(handle types.THandle, useDefaultEnvironment bool) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(2, m.Instance(), uintptr(handle), api.PasBool(useDefaultEnvironment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateWindowlessBrowserWithHandleEnvironment(handle types.THandle, environment ICoreWebView2Environment) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(3, m.Instance(), uintptr(handle), base.GetObjectUintptr(environment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateInvisibleBrowserWithBool(useDefaultEnvironment bool) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(4, m.Instance(), api.PasBool(useDefaultEnvironment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateInvisibleBrowserWithEnvironment(environment ICoreWebView2Environment) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(environment))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) Refresh() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RefreshIgnoreCache() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) Stop() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) Navigate(uRI string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(11, m.Instance(), api.PasStr(uRI))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) NavigateToString(hTMLContent string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(12, m.Instance(), api.PasStr(hTMLContent))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) NavigateWithWebResourceRequest(request ICoreWebView2WebResourceRequest) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(request))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SubscribeToDevToolsProtocolEvent(eventName string, eventID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(14, m.Instance(), api.PasStr(eventName), uintptr(eventID))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CallDevToolsProtocolMethod(methodName string, parametersAsJson string, executionID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(15, m.Instance(), api.PasStr(methodName), api.PasStr(parametersAsJson), uintptr(executionID))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, executionID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(16, m.Instance(), api.PasStr(sessionId), api.PasStr(methodName), api.PasStr(parametersAsJson), uintptr(executionID))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) FocusNext() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) FocusPrevious() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(19, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ExecuteScriptWithResult(javaScript string, executionID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(20, m.Instance(), api.PasStr(javaScript), uintptr(executionID))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ExecuteScript(javaScript string, executionID int32) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(21, m.Instance(), api.PasStr(javaScript), uintptr(executionID))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CapturePreview(imageFormat wvTypes.TWVCapturePreviewImageFormat, imageStream lcl.IStreamAdapter) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(22, m.Instance(), uintptr(imageFormat), base.GetObjectUintptr(imageStream))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) NotifyParentWindowPositionChanged() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetPermissionState(permissionKind wvTypes.TWVPermissionKind, origin string, state wvTypes.TWVPermissionState) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(24, m.Instance(), uintptr(permissionKind), api.PasStr(origin), uintptr(state))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GetNonDefaultPermissionSettings() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(25, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddBrowserExtension(extensionFolderPath string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(26, m.Instance(), api.PasStr(extensionFolderPath))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GetBrowserExtensions() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(27, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DeleteProfile() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) TrySuspend() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) Resume() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind wvTypes.TWVHostResourceAcccessKind) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(31, m.Instance(), api.PasStr(hostName), api.PasStr(folderPath), uintptr(accessKind))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearVirtualHostNameToFolderMapping(hostName string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(32, m.Instance(), api.PasStr(hostName))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RetrieveHTML() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(33, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RetrieveText(visibleTextOnly bool) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(34, m.Instance(), api.PasBool(visibleTextOnly))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RetrieveMHTML() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(35, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) Print() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(36, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ShowPrintUI(useSystemPrintDialog bool) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(37, m.Instance(), api.PasBool(useSystemPrintDialog))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PrintToPdf(resultFilePath string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(38, m.Instance(), api.PasStr(resultFilePath))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PrintToPdfStream() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(39, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) OpenDevToolsWindow() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(40, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) OpenTaskManagerWindow() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(41, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PostWebMessageAsJson(webMessageAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(42, m.Instance(), api.PasStr(webMessageAsJson))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PostWebMessageAsString(webMessageAsString string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(43, m.Instance(), api.PasStr(webMessageAsString))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(44, m.Instance(), api.PasStr(uRI), uintptr(resourceContext))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RemoveWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(45, m.Instance(), api.PasStr(uRI), uintptr(resourceContext))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(46, m.Instance(), api.PasStr(uri), uintptr(resourceContext), uintptr(requestSourceKinds))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RemoveWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(47, m.Instance(), api.PasStr(uri), uintptr(resourceContext), uintptr(requestSourceKinds))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PostWebMessageAsJsonWithAdditionalObjects(webMessageAsJson string, additionalObjects ICoreWebView2ObjectCollectionView) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(48, m.Instance(), api.PasStr(webMessageAsJson), base.GetObjectUintptr(additionalObjects))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddHostObjectToScript(name string, object types.OleVariant) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(49, m.Instance(), api.PasStr(name), uintptr(object))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RemoveHostObjectFromScript(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(50, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddScriptToExecuteOnDocumentCreated(javaScript string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(51, m.Instance(), api.PasStr(javaScript))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) RemoveScriptToExecuteOnDocumentCreated(iD string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(52, m.Instance(), api.PasStr(iD))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateCookie(name string, value string, domain string, path string) (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(53, m.Instance(), api.PasStr(name), api.PasStr(value), api.PasStr(domain), api.PasStr(path), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

func (m *TWVBrowserBase) CopyCookie(cookie ICoreWebView2Cookie) (result ICoreWebView2Cookie) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(54, m.Instance(), base.GetObjectUintptr(cookie), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Cookie(resultPtr)
	return
}

func (m *TWVBrowserBase) GetCookies(uRI string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(55, m.Instance(), api.PasStr(uRI))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) AddOrUpdateCookie(cookie ICoreWebView2Cookie) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(56, m.Instance(), base.GetObjectUintptr(cookie))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DeleteCookie(cookie ICoreWebView2Cookie) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(57, m.Instance(), base.GetObjectUintptr(cookie))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DeleteCookies(name string, uRI string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(58, m.Instance(), api.PasStr(name), api.PasStr(uRI))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DeleteCookiesWithDomainAndPath(name string, domain string, path string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(59, m.Instance(), api.PasStr(name), api.PasStr(domain), api.PasStr(path))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DeleteAllCookies() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(60, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetBoundsAndZoomFactor(bounds types.TRect, zoomFactor float64) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(61, m.Instance(), uintptr(base.UnsafePointer(&bounds)), uintptr(base.UnsafePointer(&zoomFactor)))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) OpenDefaultDownloadDialog() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(62, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CloseDefaultDownloadDialog() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(63, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SimulateEditingCommand(editingCommand wvTypes.TWV2EditingCommand) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(64, m.Instance(), uintptr(editingCommand))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SimulateKeyEvent(type_ wvTypes.TWV2KeyEventType, modifiers int32, windowsVirtualKeyCode int32, nativeVirtualKeyCode int32, timestamp int32, location int32, autoRepeat bool, isKeypad bool, isSystemKey bool, text string, unmodifiedtext string, keyIdentifier string, code string, key string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(65, m.Instance(), uintptr(type_), uintptr(modifiers), uintptr(windowsVirtualKeyCode), uintptr(nativeVirtualKeyCode), uintptr(timestamp), uintptr(location), api.PasBool(autoRepeat), api.PasBool(isKeypad), api.PasBool(isSystemKey), api.PasStr(text), api.PasStr(unmodifiedtext), api.PasStr(keyIdentifier), api.PasStr(code), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) KeyboardShortcutSearch() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(66, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) KeyboardShortcutRefreshIgnoreCache() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(67, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SendMouseInput(eventKind wvTypes.TWVMouseEventKind, virtualKeys wvTypes.TWVMouseEventVirtualKeys, mouseData uint32, point types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(68, m.Instance(), uintptr(eventKind), uintptr(virtualKeys), uintptr(mouseData), uintptr(base.UnsafePointer(&point)))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SendPointerInput(eventKind wvTypes.TWVPointerEventKind, pointerInfo ICoreWebView2PointerInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(69, m.Instance(), uintptr(eventKind), base.GetObjectUintptr(pointerInfo))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DragLeave() types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(70, m.Instance())
	return types.HRESULT(r)
}

func (m *TWVBrowserBase) DragOver(keyState uint32, point types.TPoint, outEffect *uint32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var effectPtr uintptr
	r := wVBrowserBaseAPI().SysCallN(71, m.Instance(), uintptr(keyState), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&effectPtr)))
	*outEffect = uint32(effectPtr)
	return types.HRESULT(r)
}

func (m *TWVBrowserBase) GetNonClientRegionAtPoint(point types.TPoint) wvTypes.TWVNonClientRegionKind {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(72, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return wvTypes.TWVNonClientRegionKind(r)
}

func (m *TWVBrowserBase) QueryNonClientRegion(kind wvTypes.TWVNonClientRegionKind) (result ICoreWebView2RegionRectCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(73, m.Instance(), uintptr(kind), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2RegionRectCollectionView(resultPtr)
	return
}

func (m *TWVBrowserBase) ClearCache() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(74, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearDataForOrigin(origin string, storageTypes wvTypes.TWVClearDataStorageTypes) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(75, m.Instance(), api.PasStr(origin), uintptr(storageTypes))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearBrowsingData(dataKinds wvTypes.TWVBrowsingDataKinds) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(76, m.Instance(), uintptr(dataKinds))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearBrowsingDataInTimeRange(dataKinds wvTypes.TWVBrowsingDataKinds, startTime types.TDateTime, endTime types.TDateTime) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(77, m.Instance(), uintptr(dataKinds), uintptr(base.UnsafePointer(&startTime)), uintptr(base.UnsafePointer(&endTime)))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearBrowsingDataAll() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(78, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ClearServerCertificateErrorActions() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(79, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GetFavicon(format wvTypes.TWVFaviconImageFormat) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(80, m.Instance(), uintptr(format))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CreateSharedBuffer(size int64, sharedBuffer *ICoreWebView2SharedBuffer) bool {
	if !m.IsValid() {
		return false
	}
	sharedBufferPtr := base.GetObjectUintptr(*sharedBuffer)
	r := wVBrowserBaseAPI().SysCallN(81, m.Instance(), uintptr(base.UnsafePointer(&size)), uintptr(base.UnsafePointer(&sharedBufferPtr)))
	*sharedBuffer = AsCoreWebView2SharedBuffer(sharedBufferPtr)
	return api.GoBool(r)
}

func (m *TWVBrowserBase) PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(82, m.Instance(), base.GetObjectUintptr(sharedBuffer), uintptr(access), api.PasStr(additionalDataAsJson))
	return api.GoBool(r)
}

func (m *TWVBrowserBase) GetProcessExtendedInfos() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(83, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ShowSaveAsUI() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(84, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) MoveFormTo(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(85, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TWVBrowserBase) MoveFormBy(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(86, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TWVBrowserBase) ResizeFormWidthTo(X int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(87, m.Instance(), uintptr(X))
}

func (m *TWVBrowserBase) ResizeFormHeightTo(Y int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(88, m.Instance(), uintptr(Y))
}

func (m *TWVBrowserBase) SetFormLeftTo(X int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(89, m.Instance(), uintptr(X))
}

func (m *TWVBrowserBase) SetFormTopTo(Y int32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(90, m.Instance(), uintptr(Y))
}

func (m *TWVBrowserBase) IncZoomStep() {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(91, m.Instance())
}

func (m *TWVBrowserBase) DecZoomStep() {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(92, m.Instance())
}

func (m *TWVBrowserBase) ResetZoom() {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(93, m.Instance())
}

func (m *TWVBrowserBase) ToggleMuteState() {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(94, m.Instance())
}

func (m *TWVBrowserBase) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(95, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CoreWebView2PrintSettings() ICoreWebView2PrintSettings {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(96, m.Instance())
	return AsCoreWebView2PrintSettings(r)
}

func (m *TWVBrowserBase) CoreWebView2Settings() ICoreWebView2Settings {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(97, m.Instance())
	return AsCoreWebView2Settings(r)
}

func (m *TWVBrowserBase) CoreWebView2Environment() ICoreWebView2Environment {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(98, m.Instance())
	return AsCoreWebView2Environment(r)
}

func (m *TWVBrowserBase) CoreWebView2Controller() ICoreWebView2Controller {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(99, m.Instance())
	return AsCoreWebView2Controller(r)
}

func (m *TWVBrowserBase) CoreWebView2CompositionController() ICoreWebView2CompositionController {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(100, m.Instance())
	return AsCoreWebView2CompositionController(r)
}

func (m *TWVBrowserBase) CoreWebView2() ICoreWebView2 {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(101, m.Instance())
	return AsCoreWebView2(r)
}

func (m *TWVBrowserBase) Profile() ICoreWebView2Profile {
	if !m.IsValid() {
		return nil
	}
	r := wVBrowserBaseAPI().SysCallN(102, m.Instance())
	return AsCoreWebView2Profile(r)
}

func (m *TWVBrowserBase) DefaultURL() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(103, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetDefaultURL(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(103, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) IsNavigating() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(104, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ZoomPct() (result float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(105, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) SetZoomPct(value float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(105, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVBrowserBase) ZoomStep() byte {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(106, 0, m.Instance())
	return byte(r)
}

func (m *TWVBrowserBase) SetZoomStep(value byte) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(106, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) Widget0CompHWND() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(107, m.Instance())
	return types.THandle(r)
}

func (m *TWVBrowserBase) Widget1CompHWND() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(108, m.Instance())
	return types.THandle(r)
}

func (m *TWVBrowserBase) RenderCompHWND() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(109, m.Instance())
	return types.THandle(r)
}

func (m *TWVBrowserBase) D3DWindowCompHWND() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(110, m.Instance())
	return types.THandle(r)
}

func (m *TWVBrowserBase) ScreenScale() (result float32) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(111, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) Offline() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(112, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetOffline(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(112, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IgnoreCertificateErrors() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(113, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIgnoreCertificateErrors(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(113, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) BrowserExecPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(114, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetBrowserExecPath(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(114, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) UserDataFolder() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(115, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetUserDataFolder(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(115, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) AdditionalBrowserArguments() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(116, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetAdditionalBrowserArguments(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(116, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) Language() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(117, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetLanguage(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(117, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) TargetCompatibleBrowserVersion() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(118, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetTargetCompatibleBrowserVersion(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(118, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) AllowSingleSignOnUsingOSPrimaryAccount() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(119, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetAllowSingleSignOnUsingOSPrimaryAccount(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(119, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ExclusiveUserDataFolderAccess() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(120, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetExclusiveUserDataFolderAccess(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(120, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) CustomCrashReportingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(121, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetCustomCrashReportingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(121, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) EnableTrackingPrevention() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(122, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetEnableTrackingPrevention(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(122, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) AreBrowserExtensionsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(123, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetAreBrowserExtensionsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(123, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ChannelSearchKind() wvTypes.TWVChannelSearchKind {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(124, 0, m.Instance())
	return wvTypes.TWVChannelSearchKind(r)
}

func (m *TWVBrowserBase) SetChannelSearchKind(value wvTypes.TWVChannelSearchKind) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(124, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) ReleaseChannels() wvTypes.TWVReleaseChannels {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(125, 0, m.Instance())
	return wvTypes.TWVReleaseChannels(r)
}

func (m *TWVBrowserBase) SetReleaseChannels(value wvTypes.TWVReleaseChannels) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(125, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) ScrollBarStyle() wvTypes.TWVScrollBarStyle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(126, 0, m.Instance())
	return wvTypes.TWVScrollBarStyle(r)
}

func (m *TWVBrowserBase) SetScrollBarStyle(value wvTypes.TWVScrollBarStyle) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(126, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) BrowserVersionInfo() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(127, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) BrowserProcessID() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(128, m.Instance())
	return uint32(r)
}

func (m *TWVBrowserBase) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(129, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(130, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) ContainsFullScreenElement() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(131, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) DocumentTitle() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(132, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) Source() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(133, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) CookieManager() (result ICoreWebView2CookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(134, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieManager(resultPtr)
	return
}

func (m *TWVBrowserBase) IsSuspended() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(135, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) IsDocumentPlayingAudio() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(136, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) IsMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(137, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsMuted(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(137, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) DefaultDownloadDialogCornerAlignment() wvTypes.TWVDefaultDownloadDialogCornerAlignment {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(138, 0, m.Instance())
	return wvTypes.TWVDefaultDownloadDialogCornerAlignment(r)
}

func (m *TWVBrowserBase) SetDefaultDownloadDialogCornerAlignment(value wvTypes.TWVDefaultDownloadDialogCornerAlignment) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(138, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) DefaultDownloadDialogMargin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(139, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) SetDefaultDownloadDialogMargin(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(139, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVBrowserBase) IsDefaultDownloadDialogOpen() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(140, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) StatusBarText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(141, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) FaviconURI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(142, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) MemoryUsageTargetLevel() wvTypes.TWVMemoryUsageTargetLevel {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(143, 0, m.Instance())
	return wvTypes.TWVMemoryUsageTargetLevel(r)
}

func (m *TWVBrowserBase) SetMemoryUsageTargetLevel(value wvTypes.TWVMemoryUsageTargetLevel) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(143, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) Bounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(144, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) SetBounds(value types.TRect) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(144, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVBrowserBase) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(145, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsVisible(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(145, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ParentWindow() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(146, 0, m.Instance())
	return types.THandle(r)
}

func (m *TWVBrowserBase) SetParentWindow(value types.THandle) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(146, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) ZoomFactor() (result float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(147, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) SetZoomFactor(value float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(147, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVBrowserBase) DefaultBackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(148, 0, m.Instance())
	return types.TColor(r)
}

func (m *TWVBrowserBase) SetDefaultBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(148, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) BoundsMode() wvTypes.TWVBoundsMode {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(149, 0, m.Instance())
	return wvTypes.TWVBoundsMode(r)
}

func (m *TWVBrowserBase) SetBoundsMode(value wvTypes.TWVBoundsMode) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(149, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) RasterizationScale() (result float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(150, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVBrowserBase) SetRasterizationScale(value float64) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(150, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVBrowserBase) ShouldDetectMonitorScaleChanges() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(151, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetShouldDetectMonitorScaleChanges(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(151, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) AllowExternalDrop() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(152, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetAllowExternalDrop(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(152, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) DefaultContextMenusEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(153, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetDefaultContextMenusEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(153, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) DefaultScriptDialogsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(154, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetDefaultScriptDialogsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(154, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) DevToolsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(155, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetDevToolsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(155, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) AreHostObjectsAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(156, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetAreHostObjectsAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(156, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) BuiltInErrorPageEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(157, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetBuiltInErrorPageEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(157, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ScriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(158, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetScriptEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(158, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) StatusBarEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(159, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetStatusBarEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(159, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) WebMessageEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(160, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetWebMessageEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(160, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ZoomControlEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(161, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetZoomControlEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(161, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) UserAgent() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(162, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetUserAgent(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(162, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) AreBrowserAcceleratorKeysEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(163, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetAreBrowserAcceleratorKeysEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(163, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IsGeneralAutofillEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(164, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsGeneralAutofillEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(164, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IsPasswordAutosaveEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(165, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsPasswordAutosaveEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(165, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IsPinchZoomEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(166, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsPinchZoomEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(166, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IsSwipeNavigationEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(167, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsSwipeNavigationEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(167, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) HiddenPdfToolbarItems() wvTypes.TWVPDFToolbarItems {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(168, 0, m.Instance())
	return wvTypes.TWVPDFToolbarItems(r)
}

func (m *TWVBrowserBase) SetHiddenPdfToolbarItems(value wvTypes.TWVPDFToolbarItems) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(168, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) IsReputationCheckingRequired() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(169, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsReputationCheckingRequired(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(169, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) IsNonClientRegionSupportEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(170, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsNonClientRegionSupportEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(170, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) Cursor() types.HCURSOR {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(171, m.Instance())
	return types.HCURSOR(r)
}

func (m *TWVBrowserBase) RootVisualTarget() (result lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(172, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = lcl.AsUnknown(resultPtr)
	return
}

func (m *TWVBrowserBase) SetRootVisualTarget(value lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(172, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TWVBrowserBase) SystemCursorID() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(173, m.Instance())
	return uint32(r)
}

func (m *TWVBrowserBase) AutomationProvider() (result lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(174, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = lcl.AsUnknown(resultPtr)
	return
}

func (m *TWVBrowserBase) ProcessInfos() (result ICoreWebView2ProcessInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(175, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfoCollection(resultPtr)
	return
}

func (m *TWVBrowserBase) ProfileName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(176, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetProfileName(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(176, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) IsInPrivateModeEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(177, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetIsInPrivateModeEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(177, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ScriptLocale() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(178, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetScriptLocale(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(178, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) ProfilePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(179, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) DefaultDownloadFolderPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVBrowserBaseAPI().SysCallN(180, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVBrowserBase) SetDefaultDownloadFolderPath(value string) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(180, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVBrowserBase) PreferredColorScheme() wvTypes.TWVPreferredColorScheme {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(181, 0, m.Instance())
	return wvTypes.TWVPreferredColorScheme(r)
}

func (m *TWVBrowserBase) SetPreferredColorScheme(value wvTypes.TWVPreferredColorScheme) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(181, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) PreferredTrackingPreventionLevel() wvTypes.TWVTrackingPreventionLevel {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(182, 0, m.Instance())
	return wvTypes.TWVTrackingPreventionLevel(r)
}

func (m *TWVBrowserBase) SetPreferredTrackingPreventionLevel(value wvTypes.TWVTrackingPreventionLevel) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(182, 1, m.Instance(), uintptr(value))
}

func (m *TWVBrowserBase) ProfileCookieManager() (result ICoreWebView2CookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVBrowserBaseAPI().SysCallN(183, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieManager(resultPtr)
	return
}

func (m *TWVBrowserBase) ProfileIsPasswordAutosaveEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(184, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetProfileIsPasswordAutosaveEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(184, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) ProfileIsGeneralAutofillEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVBrowserBaseAPI().SysCallN(185, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVBrowserBase) SetProfileIsGeneralAutofillEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVBrowserBaseAPI().SysCallN(185, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVBrowserBase) FrameId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := wVBrowserBaseAPI().SysCallN(186, m.Instance())
	return uint32(r)
}

func (m *TWVBrowserBase) SetOnBrowserProcessExited(fn TOnBrowserProcessExitedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserProcessExitedEvent(fn)
	base.SetEvent(m, 187, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnProcessInfosChanged(fn TOnProcessInfosChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessInfosChangedEvent(fn)
	base.SetEvent(m, 188, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnContainsFullScreenElementChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 189, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnContentLoading(fn TOnContentLoadingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContentLoadingEvent(fn)
	base.SetEvent(m, 190, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnDocumentTitleChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 191, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameNavigationCompleted(fn TOnNavigationCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationCompletedEvent(fn)
	base.SetEvent(m, 192, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameNavigationStarting(fn TOnNavigationStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationStartingEvent(fn)
	base.SetEvent(m, 193, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnHistoryChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 194, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNavigationCompleted(fn TOnNavigationCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationCompletedEvent(fn)
	base.SetEvent(m, 195, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNavigationStarting(fn TOnNavigationStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNavigationStartingEvent(fn)
	base.SetEvent(m, 196, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNewWindowRequested(fn TOnNewWindowRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNewWindowRequestedEvent(fn)
	base.SetEvent(m, 197, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnPermissionRequested(fn TOnPermissionRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPermissionRequestedEvent(fn)
	base.SetEvent(m, 198, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnProcessFailed(fn TOnProcessFailedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProcessFailedEvent(fn)
	base.SetEvent(m, 199, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnScriptDialogOpening(fn TOnScriptDialogOpeningEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScriptDialogOpeningEvent(fn)
	base.SetEvent(m, 200, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnSourceChanged(fn TOnSourceChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSourceChangedEvent(fn)
	base.SetEvent(m, 201, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWebMessageReceived(fn TOnWebMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebMessageReceivedEvent(fn)
	base.SetEvent(m, 202, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWebResourceRequested(fn TOnWebResourceRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebResourceRequestedEvent(fn)
	base.SetEvent(m, 203, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWindowCloseRequested(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 204, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnDOMContentLoaded(fn TOnDOMContentLoadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDOMContentLoadedEvent(fn)
	base.SetEvent(m, 205, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWebResourceResponseReceived(fn TOnWebResourceResponseReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebResourceResponseReceivedEvent(fn)
	base.SetEvent(m, 206, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnDownloadStarting(fn TOnDownloadStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadStartingEvent(fn)
	base.SetEvent(m, 207, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameCreated(fn TOnFrameCreatedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameCreatedEvent(fn)
	base.SetEvent(m, 208, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnClientCertificateRequested(fn TOnClientCertificateRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClientCertificateRequestedEvent(fn)
	base.SetEvent(m, 209, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnIsDocumentPlayingAudioChanged(fn TOnIsDocumentPlayingAudioChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsDocumentPlayingAudioChangedEvent(fn)
	base.SetEvent(m, 210, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnIsMutedChanged(fn TOnIsMutedChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsMutedChangedEvent(fn)
	base.SetEvent(m, 211, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnIsDefaultDownloadDialogOpenChanged(fn TOnIsDefaultDownloadDialogOpenChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIsDefaultDownloadDialogOpenChangedEvent(fn)
	base.SetEvent(m, 212, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnBasicAuthenticationRequested(fn TOnBasicAuthenticationRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBasicAuthenticationRequestedEvent(fn)
	base.SetEvent(m, 213, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnContextMenuRequested(fn TOnContextMenuRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnContextMenuRequestedEvent(fn)
	base.SetEvent(m, 214, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnStatusBarTextChanged(fn TOnStatusBarTextChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnStatusBarTextChangedEvent(fn)
	base.SetEvent(m, 215, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnServerCertificateErrorActionsCompleted(fn TOnServerCertificateErrorActionsCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerCertificateErrorActionsCompletedEvent(fn)
	base.SetEvent(m, 216, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnServerCertificateErrorDetected(fn TOnServerCertificateErrorDetectedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnServerCertificateErrorDetectedEvent(fn)
	base.SetEvent(m, 217, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFaviconChanged(fn TOnFaviconChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFaviconChangedEvent(fn)
	base.SetEvent(m, 218, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGetFaviconCompleted(fn TOnGetFaviconCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetFaviconCompletedEvent(fn)
	base.SetEvent(m, 219, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnPrintCompleted(fn TOnPrintCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintCompletedEvent(fn)
	base.SetEvent(m, 220, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnPrintToPdfStreamCompleted(fn TOnPrintToPdfStreamCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintToPdfStreamCompletedEvent(fn)
	base.SetEvent(m, 221, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnAcceleratorKeyPressed(fn TOnAcceleratorKeyPressedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAcceleratorKeyPressedEvent(fn)
	base.SetEvent(m, 222, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGotFocus(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 223, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnLostFocus(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 224, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnMoveFocusRequested(fn TOnMoveFocusRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnMoveFocusRequestedEvent(fn)
	base.SetEvent(m, 225, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnZoomFactorChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 226, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRasterizationScaleChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 227, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnCursorChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 228, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnBytesReceivedChanged(fn TOnBytesReceivedChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBytesReceivedChangedEvent(fn)
	base.SetEvent(m, 229, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnEstimatedEndTimeChanged(fn TOnEstimatedEndTimeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnEstimatedEndTimeChangedEvent(fn)
	base.SetEvent(m, 230, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnDownloadStateChanged(fn TOnDownloadStateChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDownloadStateChangedEvent(fn)
	base.SetEvent(m, 231, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameDestroyed(fn TOnFrameDestroyedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameDestroyedEvent(fn)
	base.SetEvent(m, 232, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameNameChanged(fn TOnFrameNameChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameNameChangedEvent(fn)
	base.SetEvent(m, 233, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameNavigationStarting2(fn TOnFrameNavigationStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameNavigationStartingEvent(fn)
	base.SetEvent(m, 234, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameNavigationCompleted2(fn TOnFrameNavigationCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameNavigationCompletedEvent(fn)
	base.SetEvent(m, 235, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameContentLoading(fn TOnFrameContentLoadingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameContentLoadingEvent(fn)
	base.SetEvent(m, 236, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameDOMContentLoaded(fn TOnFrameDOMContentLoadedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameDOMContentLoadedEvent(fn)
	base.SetEvent(m, 237, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameWebMessageReceived(fn TOnFrameWebMessageReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameWebMessageReceivedEvent(fn)
	base.SetEvent(m, 238, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFramePermissionRequested(fn TOnFramePermissionRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFramePermissionRequestedEvent(fn)
	base.SetEvent(m, 239, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnDevToolsProtocolEventReceived(fn TOnDevToolsProtocolEventReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDevToolsProtocolEventReceivedEvent(fn)
	base.SetEvent(m, 240, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnCustomItemSelected(fn TOnCustomItemSelectedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCustomItemSelectedEvent(fn)
	base.SetEvent(m, 241, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnClearBrowsingDataCompleted(fn TOnClearBrowsingDataCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClearBrowsingDataCompletedEvent(fn)
	base.SetEvent(m, 242, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnInitializationError(fn TOnInitializationErrorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnInitializationErrorEvent(fn)
	base.SetEvent(m, 243, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnEnvironmentCompleted(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 244, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnControllerCompleted(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 245, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnAfterCreated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 246, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnExecuteScriptCompleted(fn TOnExecuteScriptCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExecuteScriptCompletedEvent(fn)
	base.SetEvent(m, 247, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnCapturePreviewCompleted(fn TOnCapturePreviewCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCapturePreviewCompletedEvent(fn)
	base.SetEvent(m, 248, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGetCookiesCompleted(fn TOnGetCookiesCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetCookiesCompletedEvent(fn)
	base.SetEvent(m, 249, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnTrySuspendCompleted(fn TOnTrySuspendCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnTrySuspendCompletedEvent(fn)
	base.SetEvent(m, 250, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnPrintToPdfCompleted(fn TOnPrintToPdfCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrintToPdfCompletedEvent(fn)
	base.SetEvent(m, 251, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnCompositionControllerCompleted(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 252, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnCallDevToolsProtocolMethodCompleted(fn TOnCallDevToolsProtocolMethodCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCallDevToolsProtocolMethodCompletedEvent(fn)
	base.SetEvent(m, 253, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnAddScriptToExecuteOnDocumentCreatedCompleted(fn TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnAddScriptToExecuteOnDocumentCreatedCompletedEvent(fn)
	base.SetEvent(m, 254, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWebResourceResponseViewGetContentCompleted(fn TOnWebResourceResponseViewGetContentCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnWebResourceResponseViewGetContentCompletedEvent(fn)
	base.SetEvent(m, 255, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWidget0CompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 256, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnWidget1CompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 257, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRenderCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 258, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnD3DWindowCompMsg(fn TOnCompMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompMsgEvent(fn)
	base.SetEvent(m, 259, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRetrieveHTMLCompleted(fn TOnRetrieveHTMLCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRetrieveHTMLCompletedEvent(fn)
	base.SetEvent(m, 260, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRetrieveTextCompleted(fn TOnRetrieveTextCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRetrieveTextCompletedEvent(fn)
	base.SetEvent(m, 261, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRetrieveMHTMLCompleted(fn TOnRetrieveMHTMLCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRetrieveMHTMLCompletedEvent(fn)
	base.SetEvent(m, 262, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnClearCacheCompleted(fn TOnClearCacheCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClearCacheCompletedEvent(fn)
	base.SetEvent(m, 263, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnClearDataForOriginCompleted(fn TOnClearDataForOriginCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnClearDataForOriginCompletedEvent(fn)
	base.SetEvent(m, 264, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnOfflineCompleted(fn TOnOfflineCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnOfflineCompletedEvent(fn)
	base.SetEvent(m, 265, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnIgnoreCertificateErrorsCompleted(fn TOnIgnoreCertificateErrorsCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnIgnoreCertificateErrorsCompletedEvent(fn)
	base.SetEvent(m, 266, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnRefreshIgnoreCacheCompleted(fn TOnRefreshIgnoreCacheCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnRefreshIgnoreCacheCompletedEvent(fn)
	base.SetEvent(m, 267, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnSimulateKeyEventCompleted(fn TOnSimulateKeyEventCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSimulateKeyEventCompletedEvent(fn)
	base.SetEvent(m, 268, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGetCustomSchemes(fn TOnGetCustomSchemesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetCustomSchemesEvent(fn)
	base.SetEvent(m, 269, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGetNonDefaultPermissionSettingsCompleted(fn TOnGetNonDefaultPermissionSettingsCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetNonDefaultPermissionSettingsCompletedEvent(fn)
	base.SetEvent(m, 270, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnSetPermissionStateCompleted(fn TOnSetPermissionStateCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSetPermissionStateCompletedEvent(fn)
	base.SetEvent(m, 271, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnLaunchingExternalUriScheme(fn TOnLaunchingExternalUriSchemeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnLaunchingExternalUriSchemeEvent(fn)
	base.SetEvent(m, 272, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnGetProcessExtendedInfosCompleted(fn TOnGetProcessExtendedInfosCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnGetProcessExtendedInfosCompletedEvent(fn)
	base.SetEvent(m, 273, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnBrowserExtensionRemoveCompleted(fn TOnBrowserExtensionRemoveCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserExtensionRemoveCompletedEvent(fn)
	base.SetEvent(m, 274, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnBrowserExtensionEnableCompleted(fn TOnBrowserExtensionEnableCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBrowserExtensionEnableCompletedEvent(fn)
	base.SetEvent(m, 275, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnProfileAddBrowserExtensionCompleted(fn TOnProfileAddBrowserExtensionCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProfileAddBrowserExtensionCompletedEvent(fn)
	base.SetEvent(m, 276, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnProfileGetBrowserExtensionsCompleted(fn TOnProfileGetBrowserExtensionsCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProfileGetBrowserExtensionsCompletedEvent(fn)
	base.SetEvent(m, 277, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnProfileDeleted(fn TOnProfileDeletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnProfileDeletedEvent(fn)
	base.SetEvent(m, 278, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnExecuteScriptWithResultCompleted(fn TOnExecuteScriptWithResultCompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnExecuteScriptWithResultCompletedEvent(fn)
	base.SetEvent(m, 279, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNonClientRegionChanged(fn TOnNonClientRegionChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNonClientRegionChangedEvent(fn)
	base.SetEvent(m, 280, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNotificationReceived(fn TOnNotificationReceivedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNotificationReceivedEvent(fn)
	base.SetEvent(m, 281, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnNotificationCloseRequested(fn TOnNotificationCloseRequestedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnNotificationCloseRequestedEvent(fn)
	base.SetEvent(m, 282, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnSaveAsUIShowing(fn TOnSaveAsUIShowingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSaveAsUIShowingEvent(fn)
	base.SetEvent(m, 283, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnShowSaveAsUICompleted(fn TOnShowSaveAsUICompletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnShowSaveAsUICompletedEvent(fn)
	base.SetEvent(m, 284, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnSaveFileSecurityCheckStarting(fn TOnSaveFileSecurityCheckStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSaveFileSecurityCheckStartingEvent(fn)
	base.SetEvent(m, 285, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnScreenCaptureStarting(fn TOnScreenCaptureStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnScreenCaptureStartingEvent(fn)
	base.SetEvent(m, 286, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVBrowserBase) SetOnFrameScreenCaptureStarting(fn TOnFrameScreenCaptureStartingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnFrameScreenCaptureStartingEvent(fn)
	base.SetEvent(m, 287, wVBrowserBaseAPI(), api.MakeEventDataPtr(cb))
}

var (
	wVBrowserBaseOnce   base.Once
	wVBrowserBaseImport *imports.Imports = nil
)

func wVBrowserBaseAPI() *imports.Imports {
	wVBrowserBaseOnce.Do(func() {
		wVBrowserBaseImport = api.NewDefaultImports()
		wVBrowserBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVBrowserBase_CreateBrowserWithHandleBool", 0), // function CreateBrowserWithHandleBool
			/* 1 */ imports.NewTable("TWVBrowserBase_CreateBrowserWithHandleEnvironment", 0), // function CreateBrowserWithHandleEnvironment
			/* 2 */ imports.NewTable("TWVBrowserBase_CreateWindowlessBrowserWithHandleBool", 0), // function CreateWindowlessBrowserWithHandleBool
			/* 3 */ imports.NewTable("TWVBrowserBase_CreateWindowlessBrowserWithHandleEnvironment", 0), // function CreateWindowlessBrowserWithHandleEnvironment
			/* 4 */ imports.NewTable("TWVBrowserBase_CreateInvisibleBrowserWithBool", 0), // function CreateInvisibleBrowserWithBool
			/* 5 */ imports.NewTable("TWVBrowserBase_CreateInvisibleBrowserWithEnvironment", 0), // function CreateInvisibleBrowserWithEnvironment
			/* 6 */ imports.NewTable("TWVBrowserBase_GoBack", 0), // function GoBack
			/* 7 */ imports.NewTable("TWVBrowserBase_GoForward", 0), // function GoForward
			/* 8 */ imports.NewTable("TWVBrowserBase_Refresh", 0), // function Refresh
			/* 9 */ imports.NewTable("TWVBrowserBase_RefreshIgnoreCache", 0), // function RefreshIgnoreCache
			/* 10 */ imports.NewTable("TWVBrowserBase_Stop", 0), // function Stop
			/* 11 */ imports.NewTable("TWVBrowserBase_Navigate", 0), // function Navigate
			/* 12 */ imports.NewTable("TWVBrowserBase_NavigateToString", 0), // function NavigateToString
			/* 13 */ imports.NewTable("TWVBrowserBase_NavigateWithWebResourceRequest", 0), // function NavigateWithWebResourceRequest
			/* 14 */ imports.NewTable("TWVBrowserBase_SubscribeToDevToolsProtocolEvent", 0), // function SubscribeToDevToolsProtocolEvent
			/* 15 */ imports.NewTable("TWVBrowserBase_CallDevToolsProtocolMethod", 0), // function CallDevToolsProtocolMethod
			/* 16 */ imports.NewTable("TWVBrowserBase_CallDevToolsProtocolMethodForSession", 0), // function CallDevToolsProtocolMethodForSession
			/* 17 */ imports.NewTable("TWVBrowserBase_SetFocus", 0), // function SetFocus
			/* 18 */ imports.NewTable("TWVBrowserBase_FocusNext", 0), // function FocusNext
			/* 19 */ imports.NewTable("TWVBrowserBase_FocusPrevious", 0), // function FocusPrevious
			/* 20 */ imports.NewTable("TWVBrowserBase_ExecuteScriptWithResult", 0), // function ExecuteScriptWithResult
			/* 21 */ imports.NewTable("TWVBrowserBase_ExecuteScript", 0), // function ExecuteScript
			/* 22 */ imports.NewTable("TWVBrowserBase_CapturePreview", 0), // function CapturePreview
			/* 23 */ imports.NewTable("TWVBrowserBase_NotifyParentWindowPositionChanged", 0), // function NotifyParentWindowPositionChanged
			/* 24 */ imports.NewTable("TWVBrowserBase_SetPermissionState", 0), // function SetPermissionState
			/* 25 */ imports.NewTable("TWVBrowserBase_GetNonDefaultPermissionSettings", 0), // function GetNonDefaultPermissionSettings
			/* 26 */ imports.NewTable("TWVBrowserBase_AddBrowserExtension", 0), // function AddBrowserExtension
			/* 27 */ imports.NewTable("TWVBrowserBase_GetBrowserExtensions", 0), // function GetBrowserExtensions
			/* 28 */ imports.NewTable("TWVBrowserBase_DeleteProfile", 0), // function DeleteProfile
			/* 29 */ imports.NewTable("TWVBrowserBase_TrySuspend", 0), // function TrySuspend
			/* 30 */ imports.NewTable("TWVBrowserBase_Resume", 0), // function Resume
			/* 31 */ imports.NewTable("TWVBrowserBase_SetVirtualHostNameToFolderMapping", 0), // function SetVirtualHostNameToFolderMapping
			/* 32 */ imports.NewTable("TWVBrowserBase_ClearVirtualHostNameToFolderMapping", 0), // function ClearVirtualHostNameToFolderMapping
			/* 33 */ imports.NewTable("TWVBrowserBase_RetrieveHTML", 0), // function RetrieveHTML
			/* 34 */ imports.NewTable("TWVBrowserBase_RetrieveText", 0), // function RetrieveText
			/* 35 */ imports.NewTable("TWVBrowserBase_RetrieveMHTML", 0), // function RetrieveMHTML
			/* 36 */ imports.NewTable("TWVBrowserBase_Print", 0), // function Print
			/* 37 */ imports.NewTable("TWVBrowserBase_ShowPrintUI", 0), // function ShowPrintUI
			/* 38 */ imports.NewTable("TWVBrowserBase_PrintToPdf", 0), // function PrintToPdf
			/* 39 */ imports.NewTable("TWVBrowserBase_PrintToPdfStream", 0), // function PrintToPdfStream
			/* 40 */ imports.NewTable("TWVBrowserBase_OpenDevToolsWindow", 0), // function OpenDevToolsWindow
			/* 41 */ imports.NewTable("TWVBrowserBase_OpenTaskManagerWindow", 0), // function OpenTaskManagerWindow
			/* 42 */ imports.NewTable("TWVBrowserBase_PostWebMessageAsJson", 0), // function PostWebMessageAsJson
			/* 43 */ imports.NewTable("TWVBrowserBase_PostWebMessageAsString", 0), // function PostWebMessageAsString
			/* 44 */ imports.NewTable("TWVBrowserBase_AddWebResourceRequestedFilter", 0), // function AddWebResourceRequestedFilter
			/* 45 */ imports.NewTable("TWVBrowserBase_RemoveWebResourceRequestedFilter", 0), // function RemoveWebResourceRequestedFilter
			/* 46 */ imports.NewTable("TWVBrowserBase_AddWebResourceRequestedFilterWithRequestSourceKinds", 0), // function AddWebResourceRequestedFilterWithRequestSourceKinds
			/* 47 */ imports.NewTable("TWVBrowserBase_RemoveWebResourceRequestedFilterWithRequestSourceKinds", 0), // function RemoveWebResourceRequestedFilterWithRequestSourceKinds
			/* 48 */ imports.NewTable("TWVBrowserBase_PostWebMessageAsJsonWithAdditionalObjects", 0), // function PostWebMessageAsJsonWithAdditionalObjects
			/* 49 */ imports.NewTable("TWVBrowserBase_AddHostObjectToScript", 0), // function AddHostObjectToScript
			/* 50 */ imports.NewTable("TWVBrowserBase_RemoveHostObjectFromScript", 0), // function RemoveHostObjectFromScript
			/* 51 */ imports.NewTable("TWVBrowserBase_AddScriptToExecuteOnDocumentCreated", 0), // function AddScriptToExecuteOnDocumentCreated
			/* 52 */ imports.NewTable("TWVBrowserBase_RemoveScriptToExecuteOnDocumentCreated", 0), // function RemoveScriptToExecuteOnDocumentCreated
			/* 53 */ imports.NewTable("TWVBrowserBase_CreateCookie", 0), // function CreateCookie
			/* 54 */ imports.NewTable("TWVBrowserBase_CopyCookie", 0), // function CopyCookie
			/* 55 */ imports.NewTable("TWVBrowserBase_GetCookies", 0), // function GetCookies
			/* 56 */ imports.NewTable("TWVBrowserBase_AddOrUpdateCookie", 0), // function AddOrUpdateCookie
			/* 57 */ imports.NewTable("TWVBrowserBase_DeleteCookie", 0), // function DeleteCookie
			/* 58 */ imports.NewTable("TWVBrowserBase_DeleteCookies", 0), // function DeleteCookies
			/* 59 */ imports.NewTable("TWVBrowserBase_DeleteCookiesWithDomainAndPath", 0), // function DeleteCookiesWithDomainAndPath
			/* 60 */ imports.NewTable("TWVBrowserBase_DeleteAllCookies", 0), // function DeleteAllCookies
			/* 61 */ imports.NewTable("TWVBrowserBase_SetBoundsAndZoomFactor", 0), // function SetBoundsAndZoomFactor
			/* 62 */ imports.NewTable("TWVBrowserBase_OpenDefaultDownloadDialog", 0), // function OpenDefaultDownloadDialog
			/* 63 */ imports.NewTable("TWVBrowserBase_CloseDefaultDownloadDialog", 0), // function CloseDefaultDownloadDialog
			/* 64 */ imports.NewTable("TWVBrowserBase_SimulateEditingCommand", 0), // function SimulateEditingCommand
			/* 65 */ imports.NewTable("TWVBrowserBase_SimulateKeyEvent", 0), // function SimulateKeyEvent
			/* 66 */ imports.NewTable("TWVBrowserBase_KeyboardShortcutSearch", 0), // function KeyboardShortcutSearch
			/* 67 */ imports.NewTable("TWVBrowserBase_KeyboardShortcutRefreshIgnoreCache", 0), // function KeyboardShortcutRefreshIgnoreCache
			/* 68 */ imports.NewTable("TWVBrowserBase_SendMouseInput", 0), // function SendMouseInput
			/* 69 */ imports.NewTable("TWVBrowserBase_SendPointerInput", 0), // function SendPointerInput
			/* 70 */ imports.NewTable("TWVBrowserBase_DragLeave", 0), // function DragLeave
			/* 71 */ imports.NewTable("TWVBrowserBase_DragOver", 0), // function DragOver
			/* 72 */ imports.NewTable("TWVBrowserBase_GetNonClientRegionAtPoint", 0), // function GetNonClientRegionAtPoint
			/* 73 */ imports.NewTable("TWVBrowserBase_QueryNonClientRegion", 0), // function QueryNonClientRegion
			/* 74 */ imports.NewTable("TWVBrowserBase_ClearCache", 0), // function ClearCache
			/* 75 */ imports.NewTable("TWVBrowserBase_ClearDataForOrigin", 0), // function ClearDataForOrigin
			/* 76 */ imports.NewTable("TWVBrowserBase_ClearBrowsingData", 0), // function ClearBrowsingData
			/* 77 */ imports.NewTable("TWVBrowserBase_ClearBrowsingDataInTimeRange", 0), // function ClearBrowsingDataInTimeRange
			/* 78 */ imports.NewTable("TWVBrowserBase_ClearBrowsingDataAll", 0), // function ClearBrowsingDataAll
			/* 79 */ imports.NewTable("TWVBrowserBase_ClearServerCertificateErrorActions", 0), // function ClearServerCertificateErrorActions
			/* 80 */ imports.NewTable("TWVBrowserBase_GetFavicon", 0), // function GetFavicon
			/* 81 */ imports.NewTable("TWVBrowserBase_CreateSharedBuffer", 0), // function CreateSharedBuffer
			/* 82 */ imports.NewTable("TWVBrowserBase_PostSharedBufferToScript", 0), // function PostSharedBufferToScript
			/* 83 */ imports.NewTable("TWVBrowserBase_GetProcessExtendedInfos", 0), // function GetProcessExtendedInfos
			/* 84 */ imports.NewTable("TWVBrowserBase_ShowSaveAsUI", 0), // function ShowSaveAsUI
			/* 85 */ imports.NewTable("TWVBrowserBase_MoveFormTo", 0), // procedure MoveFormTo
			/* 86 */ imports.NewTable("TWVBrowserBase_MoveFormBy", 0), // procedure MoveFormBy
			/* 87 */ imports.NewTable("TWVBrowserBase_ResizeFormWidthTo", 0), // procedure ResizeFormWidthTo
			/* 88 */ imports.NewTable("TWVBrowserBase_ResizeFormHeightTo", 0), // procedure ResizeFormHeightTo
			/* 89 */ imports.NewTable("TWVBrowserBase_SetFormLeftTo", 0), // procedure SetFormLeftTo
			/* 90 */ imports.NewTable("TWVBrowserBase_SetFormTopTo", 0), // procedure SetFormTopTo
			/* 91 */ imports.NewTable("TWVBrowserBase_IncZoomStep", 0), // procedure IncZoomStep
			/* 92 */ imports.NewTable("TWVBrowserBase_DecZoomStep", 0), // procedure DecZoomStep
			/* 93 */ imports.NewTable("TWVBrowserBase_ResetZoom", 0), // procedure ResetZoom
			/* 94 */ imports.NewTable("TWVBrowserBase_ToggleMuteState", 0), // procedure ToggleMuteState
			/* 95 */ imports.NewTable("TWVBrowserBase_Initialized", 0), // property Initialized
			/* 96 */ imports.NewTable("TWVBrowserBase_CoreWebView2PrintSettings", 0), // property CoreWebView2PrintSettings
			/* 97 */ imports.NewTable("TWVBrowserBase_CoreWebView2Settings", 0), // property CoreWebView2Settings
			/* 98 */ imports.NewTable("TWVBrowserBase_CoreWebView2Environment", 0), // property CoreWebView2Environment
			/* 99 */ imports.NewTable("TWVBrowserBase_CoreWebView2Controller", 0), // property CoreWebView2Controller
			/* 100 */ imports.NewTable("TWVBrowserBase_CoreWebView2CompositionController", 0), // property CoreWebView2CompositionController
			/* 101 */ imports.NewTable("TWVBrowserBase_CoreWebView2", 0), // property CoreWebView2
			/* 102 */ imports.NewTable("TWVBrowserBase_Profile", 0), // property Profile
			/* 103 */ imports.NewTable("TWVBrowserBase_DefaultURL", 0), // property DefaultURL
			/* 104 */ imports.NewTable("TWVBrowserBase_IsNavigating", 0), // property IsNavigating
			/* 105 */ imports.NewTable("TWVBrowserBase_ZoomPct", 0), // property ZoomPct
			/* 106 */ imports.NewTable("TWVBrowserBase_ZoomStep", 0), // property ZoomStep
			/* 107 */ imports.NewTable("TWVBrowserBase_Widget0CompHWND", 0), // property Widget0CompHWND
			/* 108 */ imports.NewTable("TWVBrowserBase_Widget1CompHWND", 0), // property Widget1CompHWND
			/* 109 */ imports.NewTable("TWVBrowserBase_RenderCompHWND", 0), // property RenderCompHWND
			/* 110 */ imports.NewTable("TWVBrowserBase_D3DWindowCompHWND", 0), // property D3DWindowCompHWND
			/* 111 */ imports.NewTable("TWVBrowserBase_ScreenScale", 0), // property ScreenScale
			/* 112 */ imports.NewTable("TWVBrowserBase_Offline", 0), // property Offline
			/* 113 */ imports.NewTable("TWVBrowserBase_IgnoreCertificateErrors", 0), // property IgnoreCertificateErrors
			/* 114 */ imports.NewTable("TWVBrowserBase_BrowserExecPath", 0), // property BrowserExecPath
			/* 115 */ imports.NewTable("TWVBrowserBase_UserDataFolder", 0), // property UserDataFolder
			/* 116 */ imports.NewTable("TWVBrowserBase_AdditionalBrowserArguments", 0), // property AdditionalBrowserArguments
			/* 117 */ imports.NewTable("TWVBrowserBase_Language", 0), // property Language
			/* 118 */ imports.NewTable("TWVBrowserBase_TargetCompatibleBrowserVersion", 0), // property TargetCompatibleBrowserVersion
			/* 119 */ imports.NewTable("TWVBrowserBase_AllowSingleSignOnUsingOSPrimaryAccount", 0), // property AllowSingleSignOnUsingOSPrimaryAccount
			/* 120 */ imports.NewTable("TWVBrowserBase_ExclusiveUserDataFolderAccess", 0), // property ExclusiveUserDataFolderAccess
			/* 121 */ imports.NewTable("TWVBrowserBase_CustomCrashReportingEnabled", 0), // property CustomCrashReportingEnabled
			/* 122 */ imports.NewTable("TWVBrowserBase_EnableTrackingPrevention", 0), // property EnableTrackingPrevention
			/* 123 */ imports.NewTable("TWVBrowserBase_AreBrowserExtensionsEnabled", 0), // property AreBrowserExtensionsEnabled
			/* 124 */ imports.NewTable("TWVBrowserBase_ChannelSearchKind", 0), // property ChannelSearchKind
			/* 125 */ imports.NewTable("TWVBrowserBase_ReleaseChannels", 0), // property ReleaseChannels
			/* 126 */ imports.NewTable("TWVBrowserBase_ScrollBarStyle", 0), // property ScrollBarStyle
			/* 127 */ imports.NewTable("TWVBrowserBase_BrowserVersionInfo", 0), // property BrowserVersionInfo
			/* 128 */ imports.NewTable("TWVBrowserBase_BrowserProcessID", 0), // property BrowserProcessID
			/* 129 */ imports.NewTable("TWVBrowserBase_CanGoBack", 0), // property CanGoBack
			/* 130 */ imports.NewTable("TWVBrowserBase_CanGoForward", 0), // property CanGoForward
			/* 131 */ imports.NewTable("TWVBrowserBase_ContainsFullScreenElement", 0), // property ContainsFullScreenElement
			/* 132 */ imports.NewTable("TWVBrowserBase_DocumentTitle", 0), // property DocumentTitle
			/* 133 */ imports.NewTable("TWVBrowserBase_Source", 0), // property Source
			/* 134 */ imports.NewTable("TWVBrowserBase_CookieManager", 0), // property CookieManager
			/* 135 */ imports.NewTable("TWVBrowserBase_IsSuspended", 0), // property IsSuspended
			/* 136 */ imports.NewTable("TWVBrowserBase_IsDocumentPlayingAudio", 0), // property IsDocumentPlayingAudio
			/* 137 */ imports.NewTable("TWVBrowserBase_IsMuted", 0), // property IsMuted
			/* 138 */ imports.NewTable("TWVBrowserBase_DefaultDownloadDialogCornerAlignment", 0), // property DefaultDownloadDialogCornerAlignment
			/* 139 */ imports.NewTable("TWVBrowserBase_DefaultDownloadDialogMargin", 0), // property DefaultDownloadDialogMargin
			/* 140 */ imports.NewTable("TWVBrowserBase_IsDefaultDownloadDialogOpen", 0), // property IsDefaultDownloadDialogOpen
			/* 141 */ imports.NewTable("TWVBrowserBase_StatusBarText", 0), // property StatusBarText
			/* 142 */ imports.NewTable("TWVBrowserBase_FaviconURI", 0), // property FaviconURI
			/* 143 */ imports.NewTable("TWVBrowserBase_MemoryUsageTargetLevel", 0), // property MemoryUsageTargetLevel
			/* 144 */ imports.NewTable("TWVBrowserBase_Bounds", 0), // property Bounds
			/* 145 */ imports.NewTable("TWVBrowserBase_IsVisible", 0), // property IsVisible
			/* 146 */ imports.NewTable("TWVBrowserBase_ParentWindow", 0), // property ParentWindow
			/* 147 */ imports.NewTable("TWVBrowserBase_ZoomFactor", 0), // property ZoomFactor
			/* 148 */ imports.NewTable("TWVBrowserBase_DefaultBackgroundColor", 0), // property DefaultBackgroundColor
			/* 149 */ imports.NewTable("TWVBrowserBase_BoundsMode", 0), // property BoundsMode
			/* 150 */ imports.NewTable("TWVBrowserBase_RasterizationScale", 0), // property RasterizationScale
			/* 151 */ imports.NewTable("TWVBrowserBase_ShouldDetectMonitorScaleChanges", 0), // property ShouldDetectMonitorScaleChanges
			/* 152 */ imports.NewTable("TWVBrowserBase_AllowExternalDrop", 0), // property AllowExternalDrop
			/* 153 */ imports.NewTable("TWVBrowserBase_DefaultContextMenusEnabled", 0), // property DefaultContextMenusEnabled
			/* 154 */ imports.NewTable("TWVBrowserBase_DefaultScriptDialogsEnabled", 0), // property DefaultScriptDialogsEnabled
			/* 155 */ imports.NewTable("TWVBrowserBase_DevToolsEnabled", 0), // property DevToolsEnabled
			/* 156 */ imports.NewTable("TWVBrowserBase_AreHostObjectsAllowed", 0), // property AreHostObjectsAllowed
			/* 157 */ imports.NewTable("TWVBrowserBase_BuiltInErrorPageEnabled", 0), // property BuiltInErrorPageEnabled
			/* 158 */ imports.NewTable("TWVBrowserBase_ScriptEnabled", 0), // property ScriptEnabled
			/* 159 */ imports.NewTable("TWVBrowserBase_StatusBarEnabled", 0), // property StatusBarEnabled
			/* 160 */ imports.NewTable("TWVBrowserBase_WebMessageEnabled", 0), // property WebMessageEnabled
			/* 161 */ imports.NewTable("TWVBrowserBase_ZoomControlEnabled", 0), // property ZoomControlEnabled
			/* 162 */ imports.NewTable("TWVBrowserBase_UserAgent", 0), // property UserAgent
			/* 163 */ imports.NewTable("TWVBrowserBase_AreBrowserAcceleratorKeysEnabled", 0), // property AreBrowserAcceleratorKeysEnabled
			/* 164 */ imports.NewTable("TWVBrowserBase_IsGeneralAutofillEnabled", 0), // property IsGeneralAutofillEnabled
			/* 165 */ imports.NewTable("TWVBrowserBase_IsPasswordAutosaveEnabled", 0), // property IsPasswordAutosaveEnabled
			/* 166 */ imports.NewTable("TWVBrowserBase_IsPinchZoomEnabled", 0), // property IsPinchZoomEnabled
			/* 167 */ imports.NewTable("TWVBrowserBase_IsSwipeNavigationEnabled", 0), // property IsSwipeNavigationEnabled
			/* 168 */ imports.NewTable("TWVBrowserBase_HiddenPdfToolbarItems", 0), // property HiddenPdfToolbarItems
			/* 169 */ imports.NewTable("TWVBrowserBase_IsReputationCheckingRequired", 0), // property IsReputationCheckingRequired
			/* 170 */ imports.NewTable("TWVBrowserBase_IsNonClientRegionSupportEnabled", 0), // property IsNonClientRegionSupportEnabled
			/* 171 */ imports.NewTable("TWVBrowserBase_Cursor", 0), // property Cursor
			/* 172 */ imports.NewTable("TWVBrowserBase_RootVisualTarget", 0), // property RootVisualTarget
			/* 173 */ imports.NewTable("TWVBrowserBase_SystemCursorID", 0), // property SystemCursorID
			/* 174 */ imports.NewTable("TWVBrowserBase_AutomationProvider", 0), // property AutomationProvider
			/* 175 */ imports.NewTable("TWVBrowserBase_ProcessInfos", 0), // property ProcessInfos
			/* 176 */ imports.NewTable("TWVBrowserBase_ProfileName", 0), // property ProfileName
			/* 177 */ imports.NewTable("TWVBrowserBase_IsInPrivateModeEnabled", 0), // property IsInPrivateModeEnabled
			/* 178 */ imports.NewTable("TWVBrowserBase_ScriptLocale", 0), // property ScriptLocale
			/* 179 */ imports.NewTable("TWVBrowserBase_ProfilePath", 0), // property ProfilePath
			/* 180 */ imports.NewTable("TWVBrowserBase_DefaultDownloadFolderPath", 0), // property DefaultDownloadFolderPath
			/* 181 */ imports.NewTable("TWVBrowserBase_PreferredColorScheme", 0), // property PreferredColorScheme
			/* 182 */ imports.NewTable("TWVBrowserBase_PreferredTrackingPreventionLevel", 0), // property PreferredTrackingPreventionLevel
			/* 183 */ imports.NewTable("TWVBrowserBase_ProfileCookieManager", 0), // property ProfileCookieManager
			/* 184 */ imports.NewTable("TWVBrowserBase_ProfileIsPasswordAutosaveEnabled", 0), // property ProfileIsPasswordAutosaveEnabled
			/* 185 */ imports.NewTable("TWVBrowserBase_ProfileIsGeneralAutofillEnabled", 0), // property ProfileIsGeneralAutofillEnabled
			/* 186 */ imports.NewTable("TWVBrowserBase_FrameId", 0), // property FrameId
			/* 187 */ imports.NewTable("TWVBrowserBase_OnBrowserProcessExited", 0), // event OnBrowserProcessExited
			/* 188 */ imports.NewTable("TWVBrowserBase_OnProcessInfosChanged", 0), // event OnProcessInfosChanged
			/* 189 */ imports.NewTable("TWVBrowserBase_OnContainsFullScreenElementChanged", 0), // event OnContainsFullScreenElementChanged
			/* 190 */ imports.NewTable("TWVBrowserBase_OnContentLoading", 0), // event OnContentLoading
			/* 191 */ imports.NewTable("TWVBrowserBase_OnDocumentTitleChanged", 0), // event OnDocumentTitleChanged
			/* 192 */ imports.NewTable("TWVBrowserBase_OnFrameNavigationCompleted", 0), // event OnFrameNavigationCompleted
			/* 193 */ imports.NewTable("TWVBrowserBase_OnFrameNavigationStarting", 0), // event OnFrameNavigationStarting
			/* 194 */ imports.NewTable("TWVBrowserBase_OnHistoryChanged", 0), // event OnHistoryChanged
			/* 195 */ imports.NewTable("TWVBrowserBase_OnNavigationCompleted", 0), // event OnNavigationCompleted
			/* 196 */ imports.NewTable("TWVBrowserBase_OnNavigationStarting", 0), // event OnNavigationStarting
			/* 197 */ imports.NewTable("TWVBrowserBase_OnNewWindowRequested", 0), // event OnNewWindowRequested
			/* 198 */ imports.NewTable("TWVBrowserBase_OnPermissionRequested", 0), // event OnPermissionRequested
			/* 199 */ imports.NewTable("TWVBrowserBase_OnProcessFailed", 0), // event OnProcessFailed
			/* 200 */ imports.NewTable("TWVBrowserBase_OnScriptDialogOpening", 0), // event OnScriptDialogOpening
			/* 201 */ imports.NewTable("TWVBrowserBase_OnSourceChanged", 0), // event OnSourceChanged
			/* 202 */ imports.NewTable("TWVBrowserBase_OnWebMessageReceived", 0), // event OnWebMessageReceived
			/* 203 */ imports.NewTable("TWVBrowserBase_OnWebResourceRequested", 0), // event OnWebResourceRequested
			/* 204 */ imports.NewTable("TWVBrowserBase_OnWindowCloseRequested", 0), // event OnWindowCloseRequested
			/* 205 */ imports.NewTable("TWVBrowserBase_OnDOMContentLoaded", 0), // event OnDOMContentLoaded
			/* 206 */ imports.NewTable("TWVBrowserBase_OnWebResourceResponseReceived", 0), // event OnWebResourceResponseReceived
			/* 207 */ imports.NewTable("TWVBrowserBase_OnDownloadStarting", 0), // event OnDownloadStarting
			/* 208 */ imports.NewTable("TWVBrowserBase_OnFrameCreated", 0), // event OnFrameCreated
			/* 209 */ imports.NewTable("TWVBrowserBase_OnClientCertificateRequested", 0), // event OnClientCertificateRequested
			/* 210 */ imports.NewTable("TWVBrowserBase_OnIsDocumentPlayingAudioChanged", 0), // event OnIsDocumentPlayingAudioChanged
			/* 211 */ imports.NewTable("TWVBrowserBase_OnIsMutedChanged", 0), // event OnIsMutedChanged
			/* 212 */ imports.NewTable("TWVBrowserBase_OnIsDefaultDownloadDialogOpenChanged", 0), // event OnIsDefaultDownloadDialogOpenChanged
			/* 213 */ imports.NewTable("TWVBrowserBase_OnBasicAuthenticationRequested", 0), // event OnBasicAuthenticationRequested
			/* 214 */ imports.NewTable("TWVBrowserBase_OnContextMenuRequested", 0), // event OnContextMenuRequested
			/* 215 */ imports.NewTable("TWVBrowserBase_OnStatusBarTextChanged", 0), // event OnStatusBarTextChanged
			/* 216 */ imports.NewTable("TWVBrowserBase_OnServerCertificateErrorActionsCompleted", 0), // event OnServerCertificateErrorActionsCompleted
			/* 217 */ imports.NewTable("TWVBrowserBase_OnServerCertificateErrorDetected", 0), // event OnServerCertificateErrorDetected
			/* 218 */ imports.NewTable("TWVBrowserBase_OnFaviconChanged", 0), // event OnFaviconChanged
			/* 219 */ imports.NewTable("TWVBrowserBase_OnGetFaviconCompleted", 0), // event OnGetFaviconCompleted
			/* 220 */ imports.NewTable("TWVBrowserBase_OnPrintCompleted", 0), // event OnPrintCompleted
			/* 221 */ imports.NewTable("TWVBrowserBase_OnPrintToPdfStreamCompleted", 0), // event OnPrintToPdfStreamCompleted
			/* 222 */ imports.NewTable("TWVBrowserBase_OnAcceleratorKeyPressed", 0), // event OnAcceleratorKeyPressed
			/* 223 */ imports.NewTable("TWVBrowserBase_OnGotFocus", 0), // event OnGotFocus
			/* 224 */ imports.NewTable("TWVBrowserBase_OnLostFocus", 0), // event OnLostFocus
			/* 225 */ imports.NewTable("TWVBrowserBase_OnMoveFocusRequested", 0), // event OnMoveFocusRequested
			/* 226 */ imports.NewTable("TWVBrowserBase_OnZoomFactorChanged", 0), // event OnZoomFactorChanged
			/* 227 */ imports.NewTable("TWVBrowserBase_OnRasterizationScaleChanged", 0), // event OnRasterizationScaleChanged
			/* 228 */ imports.NewTable("TWVBrowserBase_OnCursorChanged", 0), // event OnCursorChanged
			/* 229 */ imports.NewTable("TWVBrowserBase_OnBytesReceivedChanged", 0), // event OnBytesReceivedChanged
			/* 230 */ imports.NewTable("TWVBrowserBase_OnEstimatedEndTimeChanged", 0), // event OnEstimatedEndTimeChanged
			/* 231 */ imports.NewTable("TWVBrowserBase_OnDownloadStateChanged", 0), // event OnDownloadStateChanged
			/* 232 */ imports.NewTable("TWVBrowserBase_OnFrameDestroyed", 0), // event OnFrameDestroyed
			/* 233 */ imports.NewTable("TWVBrowserBase_OnFrameNameChanged", 0), // event OnFrameNameChanged
			/* 234 */ imports.NewTable("TWVBrowserBase_OnFrameNavigationStarting2", 0), // event OnFrameNavigationStarting2
			/* 235 */ imports.NewTable("TWVBrowserBase_OnFrameNavigationCompleted2", 0), // event OnFrameNavigationCompleted2
			/* 236 */ imports.NewTable("TWVBrowserBase_OnFrameContentLoading", 0), // event OnFrameContentLoading
			/* 237 */ imports.NewTable("TWVBrowserBase_OnFrameDOMContentLoaded", 0), // event OnFrameDOMContentLoaded
			/* 238 */ imports.NewTable("TWVBrowserBase_OnFrameWebMessageReceived", 0), // event OnFrameWebMessageReceived
			/* 239 */ imports.NewTable("TWVBrowserBase_OnFramePermissionRequested", 0), // event OnFramePermissionRequested
			/* 240 */ imports.NewTable("TWVBrowserBase_OnDevToolsProtocolEventReceived", 0), // event OnDevToolsProtocolEventReceived
			/* 241 */ imports.NewTable("TWVBrowserBase_OnCustomItemSelected", 0), // event OnCustomItemSelected
			/* 242 */ imports.NewTable("TWVBrowserBase_OnClearBrowsingDataCompleted", 0), // event OnClearBrowsingDataCompleted
			/* 243 */ imports.NewTable("TWVBrowserBase_OnInitializationError", 0), // event OnInitializationError
			/* 244 */ imports.NewTable("TWVBrowserBase_OnEnvironmentCompleted", 0), // event OnEnvironmentCompleted
			/* 245 */ imports.NewTable("TWVBrowserBase_OnControllerCompleted", 0), // event OnControllerCompleted
			/* 246 */ imports.NewTable("TWVBrowserBase_OnAfterCreated", 0), // event OnAfterCreated
			/* 247 */ imports.NewTable("TWVBrowserBase_OnExecuteScriptCompleted", 0), // event OnExecuteScriptCompleted
			/* 248 */ imports.NewTable("TWVBrowserBase_OnCapturePreviewCompleted", 0), // event OnCapturePreviewCompleted
			/* 249 */ imports.NewTable("TWVBrowserBase_OnGetCookiesCompleted", 0), // event OnGetCookiesCompleted
			/* 250 */ imports.NewTable("TWVBrowserBase_OnTrySuspendCompleted", 0), // event OnTrySuspendCompleted
			/* 251 */ imports.NewTable("TWVBrowserBase_OnPrintToPdfCompleted", 0), // event OnPrintToPdfCompleted
			/* 252 */ imports.NewTable("TWVBrowserBase_OnCompositionControllerCompleted", 0), // event OnCompositionControllerCompleted
			/* 253 */ imports.NewTable("TWVBrowserBase_OnCallDevToolsProtocolMethodCompleted", 0), // event OnCallDevToolsProtocolMethodCompleted
			/* 254 */ imports.NewTable("TWVBrowserBase_OnAddScriptToExecuteOnDocumentCreatedCompleted", 0), // event OnAddScriptToExecuteOnDocumentCreatedCompleted
			/* 255 */ imports.NewTable("TWVBrowserBase_OnWebResourceResponseViewGetContentCompleted", 0), // event OnWebResourceResponseViewGetContentCompleted
			/* 256 */ imports.NewTable("TWVBrowserBase_OnWidget0CompMsg", 0), // event OnWidget0CompMsg
			/* 257 */ imports.NewTable("TWVBrowserBase_OnWidget1CompMsg", 0), // event OnWidget1CompMsg
			/* 258 */ imports.NewTable("TWVBrowserBase_OnRenderCompMsg", 0), // event OnRenderCompMsg
			/* 259 */ imports.NewTable("TWVBrowserBase_OnD3DWindowCompMsg", 0), // event OnD3DWindowCompMsg
			/* 260 */ imports.NewTable("TWVBrowserBase_OnRetrieveHTMLCompleted", 0), // event OnRetrieveHTMLCompleted
			/* 261 */ imports.NewTable("TWVBrowserBase_OnRetrieveTextCompleted", 0), // event OnRetrieveTextCompleted
			/* 262 */ imports.NewTable("TWVBrowserBase_OnRetrieveMHTMLCompleted", 0), // event OnRetrieveMHTMLCompleted
			/* 263 */ imports.NewTable("TWVBrowserBase_OnClearCacheCompleted", 0), // event OnClearCacheCompleted
			/* 264 */ imports.NewTable("TWVBrowserBase_OnClearDataForOriginCompleted", 0), // event OnClearDataForOriginCompleted
			/* 265 */ imports.NewTable("TWVBrowserBase_OnOfflineCompleted", 0), // event OnOfflineCompleted
			/* 266 */ imports.NewTable("TWVBrowserBase_OnIgnoreCertificateErrorsCompleted", 0), // event OnIgnoreCertificateErrorsCompleted
			/* 267 */ imports.NewTable("TWVBrowserBase_OnRefreshIgnoreCacheCompleted", 0), // event OnRefreshIgnoreCacheCompleted
			/* 268 */ imports.NewTable("TWVBrowserBase_OnSimulateKeyEventCompleted", 0), // event OnSimulateKeyEventCompleted
			/* 269 */ imports.NewTable("TWVBrowserBase_OnGetCustomSchemes", 0), // event OnGetCustomSchemes
			/* 270 */ imports.NewTable("TWVBrowserBase_OnGetNonDefaultPermissionSettingsCompleted", 0), // event OnGetNonDefaultPermissionSettingsCompleted
			/* 271 */ imports.NewTable("TWVBrowserBase_OnSetPermissionStateCompleted", 0), // event OnSetPermissionStateCompleted
			/* 272 */ imports.NewTable("TWVBrowserBase_OnLaunchingExternalUriScheme", 0), // event OnLaunchingExternalUriScheme
			/* 273 */ imports.NewTable("TWVBrowserBase_OnGetProcessExtendedInfosCompleted", 0), // event OnGetProcessExtendedInfosCompleted
			/* 274 */ imports.NewTable("TWVBrowserBase_OnBrowserExtensionRemoveCompleted", 0), // event OnBrowserExtensionRemoveCompleted
			/* 275 */ imports.NewTable("TWVBrowserBase_OnBrowserExtensionEnableCompleted", 0), // event OnBrowserExtensionEnableCompleted
			/* 276 */ imports.NewTable("TWVBrowserBase_OnProfileAddBrowserExtensionCompleted", 0), // event OnProfileAddBrowserExtensionCompleted
			/* 277 */ imports.NewTable("TWVBrowserBase_OnProfileGetBrowserExtensionsCompleted", 0), // event OnProfileGetBrowserExtensionsCompleted
			/* 278 */ imports.NewTable("TWVBrowserBase_OnProfileDeleted", 0), // event OnProfileDeleted
			/* 279 */ imports.NewTable("TWVBrowserBase_OnExecuteScriptWithResultCompleted", 0), // event OnExecuteScriptWithResultCompleted
			/* 280 */ imports.NewTable("TWVBrowserBase_OnNonClientRegionChanged", 0), // event OnNonClientRegionChanged
			/* 281 */ imports.NewTable("TWVBrowserBase_OnNotificationReceived", 0), // event OnNotificationReceived
			/* 282 */ imports.NewTable("TWVBrowserBase_OnNotificationCloseRequested", 0), // event OnNotificationCloseRequested
			/* 283 */ imports.NewTable("TWVBrowserBase_OnSaveAsUIShowing", 0), // event OnSaveAsUIShowing
			/* 284 */ imports.NewTable("TWVBrowserBase_OnShowSaveAsUICompleted", 0), // event OnShowSaveAsUICompleted
			/* 285 */ imports.NewTable("TWVBrowserBase_OnSaveFileSecurityCheckStarting", 0), // event OnSaveFileSecurityCheckStarting
			/* 286 */ imports.NewTable("TWVBrowserBase_OnScreenCaptureStarting", 0), // event OnScreenCaptureStarting
			/* 287 */ imports.NewTable("TWVBrowserBase_OnFrameScreenCaptureStarting", 0), // event OnFrameScreenCaptureStarting
		}
	})
	return wVBrowserBaseImport
}
