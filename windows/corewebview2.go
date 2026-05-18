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

// ICoreWebView2 Parent: IObject
type ICoreWebView2 interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// SubscribeToDevToolsProtocolEvent
	//  Subscribe to a DevTools protocol event. The TWVBrowserBase.OnDevToolsProtocolEventReceived
	//  event will be triggered on each DevTools event.
	//  <param name="aEventName">The DevTools protocol event name.</param>
	//  <param name="aEventID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	SubscribeToDevToolsProtocolEvent(eventName string, eventID int32, browserComponent lcl.IComponent) bool // function
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
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	CapturePreview(imageFormat wvTypes.TWVCapturePreviewImageFormat, imageStream lcl.IStreamAdapter, browserComponent lcl.IComponent) bool // function
	// ExecuteScript
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptCompleted event is triggered
	//  when it finishes.
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
	//  <param name="JavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	ExecuteScript(javaScript string, executionID int32, browserComponent lcl.IComponent) bool // function
	// GoBack
	//  Navigates the WebView to the previous page in the navigation history.
	GoBack() bool // function
	// GoForward
	//  Navigates the WebView to the next page in the navigation history.
	GoForward() bool // function
	// Navigate
	//  Cause a navigation of the top-level document to run to the specified URI.
	Navigate(uRI string) bool // function
	// NavigateToString
	//  Initiates a navigation to aHTMLContent as source HTML of a new document.
	//  The `aHTMLContent` parameter may not be larger than 2 MB (2 * 1024 * 1024 bytes) in total size.
	//  The origin of the new page is `about:blank`.
	NavigateToString(hTMLContent string) bool // function
	// NavigateWithWebResourceRequest
	//  Navigates using a constructed ICoreWebView2WebResourceRequest object. This lets you
	//  provide post data or additional request headers during navigation.
	//  The headers in aRequest override headers added by WebView2 runtime except for Cookie headers.
	//  Method can only be either "GET" or "POST". Provided post data will only
	//  be sent only if the method is "POST" and the uri scheme is HTTP(S).
	NavigateWithWebResourceRequest(request ICoreWebView2WebResourceRequest) bool // function
	// Reload
	//  Reload the current page. This is similar to navigating to the URI of
	//  current top level document including all navigation events firing and
	//  respecting any entries in the HTTP cache. But, the back or forward
	//  history are not modified.
	Reload() bool // function
	// Stop
	//  Stop all navigations and pending resource fetches. Does not stop scripts.
	Stop() bool // function
	// TrySuspend
	//  An app may call the `TrySuspend` API to have the WebView2 consume less memory.
	//  This is useful when a Win32 app becomes invisible, or when a Universal Windows
	//  Platform app is being suspended, during the suspended event handler before completing
	//  the suspended event.
	//  The CoreWebView2Controller's IsVisible property must be false when the API is called.
	//  Otherwise, the API fails with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Suspending is similar to putting a tab to sleep in the Edge browser. Suspending pauses
	//  WebView script timers and animations, minimizes CPU usage for the associated
	//  browser renderer process and allows the operating system to reuse the memory that was
	//  used by the renderer process for other processes.
	//  Note that Suspend is best effort and considered completed successfully once the request
	//  is sent to browser renderer process. If there is a running script, the script will continue
	//  to run and the renderer process will be suspended after that script is done.
	//  See [Sleeping Tabs FAQ](https://techcommunity.microsoft.com/t5/articles/sleeping-tabs-faq/m-p/1705434)
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
	TrySuspend(handler ICoreWebView2TrySuspendCompletedHandler) bool // function
	// Resume
	//  Resumes the WebView so that it resumes activities on the web page.
	//  This API can be called while the WebView2 controller is invisible.
	//  The app can interact with the WebView immediately after `Resume`.
	//  WebView will be automatically resumed when it becomes visible.
	Resume() bool // function
	// SetVirtualHostNameToFolderMapping
	//  Sets a mapping between a virtual host name and a folder path to make available to web sites
	//  via that host name.
	//
	//  After setting the mapping, documents loaded in the WebView can use HTTP or HTTPS URLs at
	//  the specified host name specified by hostName to access files in the local folder specified
	//  by folderPath.
	//
	//  This mapping applies to both top-level document and iframe navigations as well as subresource
	//  references from a document. This also applies to web workers including dedicated/shared worker
	//  and service worker, for loading either worker scripts or subresources
	//  (importScripts(), fetch(), XHR, etc.) issued from within a worker.
	//  For virtual host mapping to work with service worker, please keep the virtual host name
	//  mappings consistent among all WebViews sharing the same browser instance. As service worker
	//  works independently of WebViews, we merge mappings from all WebViews when resolving virtual
	//  host name, inconsistent mappings between WebViews would lead unexpected behavior.
	//
	//  Due to a current implementation limitation, media files accessed using virtual host name can be
	//  very slow to load.
	//  As the resource loaders for the current page might have already been created and running,
	//  changes to the mapping might not be applied to the current page and a reload of the page is
	//  needed to apply the new mapping.
	//
	//  Both absolute and relative paths are supported for folderPath. Relative paths are interpreted
	//  as relative to the folder where the exe of the app is in.
	//
	//  Note that the folderPath length must not exceed the Windows MAX_PATH limit.
	//
	//  accessKind specifies the level of access to resources under the virtual host from other sites.
	//
	//  For example, after calling
	//  <code>
	//  ```cpp
	//  SetVirtualHostNameToFolderMapping(
	//  L"appassets.example", L"assets",
	//  COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY);
	//  ```
	//  </code>
	//  navigating to `https://appassets.example/my-local-file.html` will
	//  show the content from my-local-file.html in the assets subfolder located on disk under
	//  the same path as the app's executable file.
	//
	//  DOM elements that want to reference local files will have their host reference virtual host in the source.
	//  If there are multiple folders being used, define one unique virtual host per folder.
	//  For example, you can embed a local image like this
	//  <code>
	//  ```cpp
	//  WCHAR c_navString[] = L"<img src=\"http://appassets.example/wv2.png\"/>";
	//  m_webView->NavigateToString(c_navString);
	//  ```
	//  </code>
	//  The example above shows the image wv2.png by resolving the folder mapping above.
	//
	//  You should typically choose virtual host names that are never used by real sites.
	//  If you own a domain such as example.com, another option is to use a subdomain reserved for
	//  the app (like my-app.example.com).
	//
	//  [RFC 6761](https://tools.ietf.org/html/rfc6761) has reserved several special-use domain
	//  names that are guaranteed to not be used by real sites (for example, .example, .test, and
	//  .invalid.)
	//
	//  Note that using `.local` as the top-level domain name will work but can cause a delay
	//  during navigations. You should avoid using `.local` if you can.
	//
	//  Apps should use distinct domain names when mapping folder from different sources that
	//  should be isolated from each other. For instance, the app might use app-file.example for
	//  files that ship as part of the app, and book1.example might be used for files containing
	//  books from a less trusted source that were previously downloaded and saved to the disk by
	//  the app.
	//
	//  The host name used in the APIs is canonicalized using Chromium's host name parsing logic
	//  before being used internally. For more information see [HTML5 2.6 URLs](https://dev.w3.org/html5/spec-LC/urls.html).
	//
	//  All host names that are canonicalized to the same string are considered identical.
	//  For example, `EXAMPLE.COM` and `example.com` are treated as the same host name.
	//  An international host name and its Punycode-encoded host name are considered the same host
	//  name. There is no DNS resolution for host name and the trailing '.' is not normalized as
	//  part of canonicalization.
	//
	//  Therefore `example.com` and `example.com.` are treated as different host names. Similarly,
	//  `virtual-host-name` and `virtual-host-name.example.com` are treated as different host names
	//  even if the machine has a DNS suffix of `example.com`.
	//
	//  Specify the minimal cross-origin access necessary to run the app. If there is not a need to
	//  access local resources from other origins, use COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY.
	SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind wvTypes.TWVHostResourceAcccessKind) bool // function
	// ClearVirtualHostNameToFolderMapping
	//  Clears a host name mapping for local folder that was added by `SetVirtualHostNameToFolderMapping`.
	ClearVirtualHostNameToFolderMapping(hostName string) bool // function
	// OpenTaskManagerWindow
	//  Opens the Browser Task Manager view as a new window in the foreground.
	//  If the Browser Task Manager is already open, this will bring it into
	//  the foreground. WebView2 currently blocks the Shift+Esc shortcut for
	//  opening the task manager. An end user can open the browser task manager
	//  manually via the `Browser task manager` entry of the DevTools window's
	//  title bar's context menu.
	OpenTaskManagerWindow() bool // function
	// PrintToPdf
	//  Print the current page to PDF asynchronously with the provided settings.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//
	//  Use `resultFilePath` to specify the path to the PDF file. The host should
	//  provide an absolute path, including file name. If the path
	//  points to an existing file, the file will be overwritten. If the path is
	//  not valid, the method fails with `E_INVALIDARG`.
	//
	//  The async `PrintToPdf` operation completes when the data has been written
	//  to the PDF file. At this time the
	//  `ICoreWebView2PrintToPdfCompletedHandler` is invoked. If the
	//  application exits before printing is complete, the file is not saved.
	//  Only one `Printing` operation can be in progress at a time. If
	//  `PrintToPdf` is called while a `PrintToPdf` or `PrintToPdfStream` or `Print` or
	//  `ShowPrintUI` job is in progress, the completed handler is immediately invoked
	//  with `isSuccessful` set to FALSE.
	PrintToPdf(resultFilePath string, printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintToPdfCompletedHandler) bool // function
	// OpenDevToolsWindow
	//  Opens the DevTools window for the current document in the WebView. Does
	//  nothing if run when the DevTools window is already open.
	OpenDevToolsWindow() bool // function
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
	//  `ICoreWebView2Settings::IsWebMessageEnabled` setting must be `TRUE` or
	//  the web message will not be sent. The `data` property of the event
	//  arg is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event arg is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before
	//  the message is posted to the page, the message is discarded.
	PostWebMessageAsJson(webMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  arg of the `window.chrome.webview` message is a string with the same
	//  value as `webMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	PostWebMessageAsString(webMessageAsString string) bool // function
	// CallDevToolsProtocolMethod
	//  Runs an asynchronous `DevToolsProtocol` method. For more information
	//  about available methods, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot).
	//  The `methodName` parameter is the full name of the method in the
	//  `{domain}.{method}` format. The `parametersAsJson` parameter is a JSON
	//  formatted string containing the parameters for the corresponding method.
	//  The `Invoke` method of the `handler` is run when the method
	//  asynchronously completes. `Invoke` is run with the return object of the
	//  method as a JSON string. This function returns E_INVALIDARG if the `methodName` is
	//  unknown or the `parametersAsJson` has an error. In the case of such an error, the
	//  `returnObjectAsJson` parameter of the handler will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method's completed handler to run before calling the
	//  next method.
	//  If the method is to run in add_NewWindowRequested handler it should be called
	//  before the new window is set if the cdp message should affect the initial navigation. If
	//  called after setting the NewWindow property, the cdp messages
	//  may or may not apply to the initial navigation and may only apply to the subsequent navigation.
	//  For more details see `ICoreWebView2NewWindowRequestedEventArgs::put_NewWindow`.
	CallDevToolsProtocolMethod(methodName string, parametersAsJson string, executionID int32, browserComponent lcl.IComponent) bool // function
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
	//  For more information about targets and sessions, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot/Target).
	//  For more information about available methods, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot)
	//  The `sessionId` parameter is the sessionId for an attached target.
	//  nullptr or empty string is treated as the session for the default target for the top page.
	//  The `methodName` parameter is the full name of the method in the
	//  `{domain}.{method}` format. The `parametersAsJson` parameter is a JSON
	//  formatted string containing the parameters for the corresponding method.
	//  The `Invoke` method of the `handler` is run when the method
	//  asynchronously completes. `Invoke` is run with the return object of the
	//  method as a JSON string. This function returns E_INVALIDARG if the `methodName` is
	//  unknown or the `parametersAsJson` has an error. In the case of such an error, the
	//  `returnObjectAsJson` parameter of the handler will include information
	//  about the error.
	CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, executionID int32, browserComponent lcl.IComponent) bool // function
	// AddWebResourceRequestedFilter
	//  Warning: This method is deprecated and does not behave as expected for
	//  iframes. It will cause the WebResourceRequested event to fire only for the
	//  main frame and its same-origin iframes. Please use
	//  `AddWebResourceRequestedFilterWithRequestSourceKinds`
	//  instead, which will let the event to fire for all iframes on the
	//  document.
	//
	//  Adds a URI and resource context filter for the `WebResourceRequested`
	//  event. A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string will be raised via the `WebResourceRequested` event.
	//
	//  The `uri` parameter value is a wildcard string matched against the URI
	//  of the web resource request. This is a glob style
	//  wildcard string in which a `*` matches zero or more characters and a `?`
	//  matches exactly one character.
	//  These wildcard characters can be escaped using a backslash just before
	//  the wildcard character in order to represent the literal `*` or `?`.
	//
	//  The matching occurs over the URI as a whole string and not limiting
	//  wildcard matches to particular parts of the URI.
	//  The wildcard filter is compared to the URI after the URI has been
	//  normalized, any URI fragment has been removed, and non-ASCII hostnames
	//  have been converted to punycode.
	//
	//  Specifying a `nullptr` for the uri is equivalent to an empty string which
	//  matches no URIs.
	//
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
	RemoveWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool // function
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
	AddHostObjectToScript(name string, object types.OleVariant) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the WebView. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the WebView, the JavaScript code continues to have access to that
	//  object. Run this method for a name that is already removed or never
	//  added fails.
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
	//  If the method is run in add_NewWindowRequested handler it should be called
	//  before the new window is set. If called after setting the NewWindow property, the initial script
	//  may or may not apply to the initial navigation and may only apply to the subsequent navigation.
	//  For more details see `ICoreWebView2NewWindowRequestedEventArgs::put_NewWindow`.
	//  NOTE: If an HTML document is running in a sandbox of some kind using
	//  [sandbox](https://developer.mozilla.org/docs/Web/HTML/Element/iframe#attr-sandbox)
	//  properties or the
	//  [Content-Security-Policy](https://developer.mozilla.org/docs/Web/HTTP/Headers/Content-Security-Policy)
	//  HTTP header affects the script that runs. For example, if the
	//  `allow-modals` keyword is not set then requests to run the `alert`
	//  function are ignored.
	AddScriptToExecuteOnDocumentCreated(javaScript string, browserComponent lcl.IComponent) bool // function
	// RemoveScriptToExecuteOnDocumentCreated
	//  Remove the corresponding JavaScript added using
	//  `AddScriptToExecuteOnDocumentCreated` with the specified script ID. The
	//  script ID should be the one returned by the `AddScriptToExecuteOnDocumentCreated`.
	//  Both use `AddScriptToExecuteOnDocumentCreated` and this method in `NewWindowRequested`
	//  event handler at the same time sometimes causes trouble. Since invalid scripts will
	//  be ignored, the script IDs you got may not be valid anymore.
	RemoveScriptToExecuteOnDocumentCreated(iD string) bool // function
	// OpenDefaultDownloadDialog
	//  Open the default download dialog. If the dialog is opened before there
	//  are recent downloads, the dialog shows all past downloads for the
	//  current profile. Otherwise, the dialog shows only the recent downloads
	//  with a "See more" button for past downloads. Calling this method raises
	//  the `IsDefaultDownloadDialogOpenChanged` event if the dialog was closed.
	//  No effect if the dialog is already open.
	OpenDefaultDownloadDialog() bool // function
	// CloseDefaultDownloadDialog
	//  Close the default download dialog. Calling this method raises the
	//  `IsDefaultDownloadDialogOpenChanged` event if the dialog was open. No
	//  effect if the dialog is already closed.
	CloseDefaultDownloadDialog() bool // function
	// ClearServerCertificateErrorActions
	//  Clears all cached decisions to proceed with TLS certificate errors from the
	//  ServerCertificateErrorDetected event for all WebView2's sharing the same session.
	ClearServerCertificateErrorActions(browserComponent lcl.IComponent) bool // function
	// GetFavicon
	//  Async function for getting the actual image data of the favicon.
	//  The image is copied to the `imageStream` object in `ICoreWebView2GetFaviconCompletedHandler`.
	//  If there is no image then no data would be copied into the imageStream.
	//  The `format` is the file format to return the image stream.
	//  `completedHandler` is executed at the end of the operation.
	GetFavicon(format wvTypes.TWVFaviconImageFormat, browserComponent lcl.IComponent) bool // function
	// Print
	//  Print the current web page asynchronously to the specified printer with the provided settings.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//  The handler will return `errorCode` as `S_OK` and `printStatus` as COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE
	//  if `printerName` doesn't match with the name of any installed printers on the user OS. The handler
	//  will return `errorCode` as `E_INVALIDARG` and `printStatus` as COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR
	//  if the caller provides invalid settings for a given printer.
	//  The async `Print` operation completes when it finishes printing to the printer.
	//  At this time the `ICoreWebView2PrintCompletedHandler` is invoked.
	//  Only one `Printing` operation can be in progress at a time. If `Print` is called while a `Print` or `PrintToPdf`
	//  or `PrintToPdfStream` or `ShowPrintUI` job is in progress, the completed handler is immediately invoked
	//  with `E_ABORT` and `printStatus` is COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR.
	//  This is only for printing operation on one webview.
	//  <code>
	//  | errorCode | printStatus | Notes |
	//  | --- | --- | --- |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_SUCCEEDED | Print operation succeeded. |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE | If specified printer is not found or printer status is not available, offline or error state. |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | Print operation is failed. |
	//  | E_INVALIDARG | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | If the caller provides invalid settings for the specified printer. |
	//  | E_ABORT | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | Print operation is failed as printing job already in progress. |
	//  </code>
	Print(printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintCompletedHandler) bool // function
	// ShowPrintUI
	//  Opens the print dialog to print the current web page. See `COREWEBVIEW2_PRINT_DIALOG_KIND`
	//  for descriptions of print dialog kinds.
	//  Invoking browser or system print dialog doesn't open new print dialog if
	//  it is already open.
	ShowPrintUI(printDialogKind wvTypes.TWVPrintDialogKind) bool // function
	// PrintToPdfStream
	//  Provides the Pdf data of current web page asynchronously for the provided settings.
	//  Stream will be rewound to the start of the pdf data.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//  The async `PrintToPdfStream` operation completes when it finishes
	//  writing to the stream. At this time the `ICoreWebView2PrintToPdfStreamCompletedHandler`
	//  is invoked. Only one `Printing` operation can be in progress at a time. If
	//  `PrintToPdfStream` is called while a `PrintToPdfStream` or `PrintToPdf` or `Print`
	//  or `ShowPrintUI` job is in progress, the completed handler is immediately invoked with `E_ABORT`.
	//  This is only for printing operation on one webview.
	PrintToPdfStream(printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintToPdfStreamCompletedHandler) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the main frame in the WebView.
	//  The script will receive a `sharedbufferreceived` event from chrome.webview.
	//  The event arg for that event will have the following methods and properties:
	//   `getBuffer()`: return an ArrayBuffer object with the backing content from the shared buffer.
	//   `additionalData`: an object as the result of parsing `additionalDataAsJson` as JSON string.
	//  This property will be `undefined` if `additionalDataAsJson` is nullptr or empty string.
	//   `source`: with a value set as `chrome.webview` object.
	//  If a string is provided as `additionalDataAsJson` but it is not a valid JSON string,
	//  the API will fail with `E_INVALIDARG`.
	//  If `access` is COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY, the script will only have read access to the buffer.
	//  If the script tries to modify the content in a read only buffer, it will cause an access
	//  violation in WebView renderer process and crash the renderer process.
	//  If the shared buffer is already closed, the API will fail with `RO_E_CLOSED`.
	//  The script code should call `chrome.webview.releaseBuffer` with
	//  the shared buffer as the parameter to release underlying resources as soon
	//  as it does not need access to the shared buffer any more.
	//  The application can post the same shared buffer object to multiple web pages or iframes, or
	//  post to the same web page or iframe multiple times. Each `PostSharedBufferToScript` will
	//  create a separate ArrayBuffer object with its own view of the memory and is separately
	//  released. The underlying shared memory will be released when all the views are released.
	PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool // function
	// ExecuteScriptWithResult
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The result of the execution is returned asynchronously in the CoreWebView2ExecuteScriptResult object
	//  which has methods and properties to obtain the successful result of script execution as well as any
	//  unhandled JavaScript exceptions.
	//  If this method is
	//  run after the NavigationStarting event during a navigation, the script
	//  runs in the new document when loading it, around the time
	//  ContentLoading is run. This operation executes the script even if
	//  ICoreWebView2Settings::IsScriptEnabled is set to FALSE.
	//
	//  \snippet ScriptComponent.cpp ExecuteScriptWithResult
	ExecuteScriptWithResult(javaScript string, executionID int32, browserComponent lcl.IComponent) bool // function
	// AddWebResourceRequestedFilterWithRequestSourceKinds
	//  A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string for corresponding request sources will be raised via
	//  the `WebResourceRequested` event. To receive all raised events filters
	//  have to be added before main page navigation.
	//
	//  The `uri` parameter value is a wildcard string matched against the URI
	//  of the web resource request. This is a glob style
	//  wildcard string in which a `*` matches zero or more characters and a `?`
	//  matches exactly one character.
	//  These wildcard characters can be escaped using a backslash just before
	//  the wildcard character in order to represent the literal `*` or `?`.
	//
	//  The matching occurs over the URI as a whole string and not limiting
	//  wildcard matches to particular parts of the URI.
	//  The wildcard filter is compared to the URI after the URI has been
	//  normalized, any URI fragment has been removed, and non-ASCII hostnames
	//  have been converted to punycode.
	//
	//  Specifying a `nullptr` for the uri is equivalent to an empty string which
	//  matches no URIs.
	//
	//  For more information about resource context filters, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_CONTEXT](/microsoft-edge/webview2/reference/win32/icorewebview2#corewebview2_web_resource_context).
	//
	//  The `requestSourceKinds` is a mask of one or more
	//  `COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS`. OR operation(s) can be
	//  applied to multiple `COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS` to
	//  create a mask representing those data types. API returns `E_INVALIDARG` if
	//  `requestSourceKinds` equals to zero. For more information about request
	//  source kinds, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS](/microsoft-edge/webview2/reference/win32/icorewebview2#corewebview2_web_resource_request_source_kinds).
	//
	//  Because service workers and shared workers run separately from any one
	//  HTML document their WebResourceRequested will be raised for all
	//  CoreWebView2s that have appropriate filters added in the corresponding
	//  CoreWebView2Environment. You should only add a WebResourceRequested filter
	//  for COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SERVICE_WORKER or
	//  COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SHARED_WORKER on
	//  one CoreWebView2 to avoid handling the same WebResourceRequested
	//  event multiple times.
	//
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
	//
	//  \snippet ScenarioSharedWorkerWRR.cpp WebResourceRequested2
	AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool // function
	// RemoveWebResourceRequestedFilterWithRequestSourceKinds
	//  Removes a matching WebResource filter that was previously added for the
	//  `WebResourceRequested` event. If the same filter was added multiple
	//  times, then it must be removed as many times as it was added for the
	//  removal to be effective. Returns `E_INVALIDARG` for a filter that was
	//  not added or is already removed.
	//  If the filter was added for multiple requestSourceKinds and removed just for one of them
	//  the filter remains for the non-removed requestSourceKinds.
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
	PostWebMessageAsJsonWithAdditionalObjects(webMessageAsJson string, additionalObjects ICoreWebView2ObjectCollectionView) bool // function
	// ShowSaveAsUI
	//  Programmatically trigger a Save As action for the currently loaded document.
	//  The `SaveAsUIShowing` event is raised.
	//
	//  Opens a system modal dialog by default. If the `SuppressDefaultDialog` property
	//  is `TRUE`, the system dialog is not opened.
	//
	//  This method returns `COREWEBVIEW2_SAVE_AS_UI_RESULT`. See
	//  `COREWEBVIEW2_SAVE_AS_UI_RESULT` for details.
	ShowSaveAsUI(browserComponent lcl.IComponent) bool // function
	AfterConstruction()                                // procedure
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2 // property BaseIntf Getter
	// Settings
	//  The `ICoreWebView2Settings` object contains various modifiable settings
	//  for the running WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_settings">See the ICoreWebView2 article.</see>
	Settings() ICoreWebView2Settings // property Settings Getter
	// BrowserProcessID
	//  The process ID of the browser process that hosts the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_browserprocessid">See the ICoreWebView2 article.</see>
	BrowserProcessID() types.DWORD // property BrowserProcessID Getter
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
	// Environment
	//  Exposes the CoreWebView2Environment used to create this CoreWebView2.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#get_environment">See the ICoreWebView2_2 article.</see>
	Environment() ICoreWebView2Environment // property Environment Getter
	// IsSuspended
	//  Whether WebView is suspended.
	//  `TRUE` when WebView is suspended, from the time when TrySuspend has completed
	//  successfully until WebView is resumed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#get_issuspended">See the ICoreWebView2_3 article.</see>
	IsSuspended() bool // property IsSuspended Getter
	// IsMuted
	//  Indicates whether all audio output from this CoreWebView2 is muted or not.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</see>
	IsMuted() bool         // property IsMuted Getter
	SetIsMuted(value bool) // property IsMuted Setter
	// IsDocumentPlayingAudio
	//  Indicates whether any audio output from this CoreWebView2 is playing.
	//  This property will be true if audio is playing even if IsMuted is true.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_isdocumentplayingaudio">See the ICoreWebView2_8 article.</see>
	IsDocumentPlayingAudio() bool // property IsDocumentPlayingAudio Getter
	// IsDefaultDownloadDialogOpen
	//  `TRUE` if the default download dialog is currently open. The value of this
	//  property changes only when the default download dialog is explicitly
	//  opened or closed. Hiding the WebView implicitly hides the dialog, but does
	//  not change the value of this property.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_isdefaultdownloaddialogopen">See the ICoreWebView2_9 article.</see>
	IsDefaultDownloadDialogOpen() bool // property IsDefaultDownloadDialogOpen Getter
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
	// StatusBarText
	//  The status message text.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_12#get_statusbartext">See the ICoreWebView2_12 article.</see>
	StatusBarText() string // property StatusBarText Getter
	// Profile
	//  The associated `ICoreWebView2Profile` object. If this CoreWebView2 was created with a
	//  CoreWebView2ControllerOptions, the CoreWebView2Profile will match those specified options.
	//  Otherwise if this CoreWebView2 was created without a CoreWebView2ControllerOptions, then
	//  this will be the default CoreWebView2Profile for the corresponding CoreWebView2Environment.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_13#get_profile">See the ICoreWebView2_13 article.</see>
	Profile() ICoreWebView2Profile // property Profile Getter
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
	// FrameId
	//  The unique identifier of the main frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2Frame` and via `ICoreWebView2FrameInfo`.
	//  Note that `FrameId` may not be valid if `ICoreWebView2` has not done
	//  any navigation. It's safe to get this value during or after the first
	//  `ContentLoading` event. Otherwise, it could return the invalid frame Id 0.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_20#get_frameid">See the ICoreWebView2_20 article.</see>
	FrameId() uint32 // property FrameId Getter
}

type TCoreWebView2 struct {
	TObject
}

func (m *TCoreWebView2) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) SubscribeToDevToolsProtocolEvent(eventName string, eventID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(2, m.Instance(), api.PasStr(eventName), uintptr(eventID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) CapturePreview(imageFormat wvTypes.TWVCapturePreviewImageFormat, imageStream lcl.IStreamAdapter, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(3, m.Instance(), uintptr(imageFormat), base.GetObjectUintptr(imageStream), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) ExecuteScript(javaScript string, executionID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(4, m.Instance(), api.PasStr(javaScript), uintptr(executionID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) GoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) GoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) Navigate(uRI string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(7, m.Instance(), api.PasStr(uRI))
	return api.GoBool(r)
}

func (m *TCoreWebView2) NavigateToString(hTMLContent string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(8, m.Instance(), api.PasStr(hTMLContent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) NavigateWithWebResourceRequest(request ICoreWebView2WebResourceRequest) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(9, m.Instance(), base.GetObjectUintptr(request))
	return api.GoBool(r)
}

func (m *TCoreWebView2) Reload() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) Stop() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) TrySuspend(handler ICoreWebView2TrySuspendCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(12, m.Instance(), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2) Resume() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind wvTypes.TWVHostResourceAcccessKind) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(14, m.Instance(), api.PasStr(hostName), api.PasStr(folderPath), uintptr(accessKind))
	return api.GoBool(r)
}

func (m *TCoreWebView2) ClearVirtualHostNameToFolderMapping(hostName string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(15, m.Instance(), api.PasStr(hostName))
	return api.GoBool(r)
}

func (m *TCoreWebView2) OpenTaskManagerWindow() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) PrintToPdf(resultFilePath string, printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintToPdfCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(17, m.Instance(), api.PasStr(resultFilePath), base.GetObjectUintptr(printSettings), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2) OpenDevToolsWindow() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) PostWebMessageAsJson(webMessageAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(19, m.Instance(), api.PasStr(webMessageAsJson))
	return api.GoBool(r)
}

func (m *TCoreWebView2) PostWebMessageAsString(webMessageAsString string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(20, m.Instance(), api.PasStr(webMessageAsString))
	return api.GoBool(r)
}

func (m *TCoreWebView2) CallDevToolsProtocolMethod(methodName string, parametersAsJson string, executionID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(21, m.Instance(), api.PasStr(methodName), api.PasStr(parametersAsJson), uintptr(executionID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, executionID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(22, m.Instance(), api.PasStr(sessionId), api.PasStr(methodName), api.PasStr(parametersAsJson), uintptr(executionID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) AddWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(23, m.Instance(), api.PasStr(uRI), uintptr(resourceContext))
	return api.GoBool(r)
}

func (m *TCoreWebView2) RemoveWebResourceRequestedFilter(uRI string, resourceContext wvTypes.TWVWebResourceContext) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(24, m.Instance(), api.PasStr(uRI), uintptr(resourceContext))
	return api.GoBool(r)
}

func (m *TCoreWebView2) AddHostObjectToScript(name string, object types.OleVariant) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(25, m.Instance(), api.PasStr(name), uintptr(object))
	return api.GoBool(r)
}

func (m *TCoreWebView2) RemoveHostObjectFromScript(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(26, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCoreWebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(27, m.Instance(), api.PasStr(javaScript), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) RemoveScriptToExecuteOnDocumentCreated(iD string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(28, m.Instance(), api.PasStr(iD))
	return api.GoBool(r)
}

func (m *TCoreWebView2) OpenDefaultDownloadDialog() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) CloseDefaultDownloadDialog() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) ClearServerCertificateErrorActions(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(31, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) GetFavicon(format wvTypes.TWVFaviconImageFormat, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(32, m.Instance(), uintptr(format), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) Print(printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(33, m.Instance(), base.GetObjectUintptr(printSettings), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2) ShowPrintUI(printDialogKind wvTypes.TWVPrintDialogKind) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(34, m.Instance(), uintptr(printDialogKind))
	return api.GoBool(r)
}

func (m *TCoreWebView2) PrintToPdfStream(printSettings ICoreWebView2PrintSettings, handler ICoreWebView2PrintToPdfStreamCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(35, m.Instance(), base.GetObjectUintptr(printSettings), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2) PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(36, m.Instance(), base.GetObjectUintptr(sharedBuffer), uintptr(access), api.PasStr(additionalDataAsJson))
	return api.GoBool(r)
}

func (m *TCoreWebView2) ExecuteScriptWithResult(javaScript string, executionID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(37, m.Instance(), api.PasStr(javaScript), uintptr(executionID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(38, m.Instance(), api.PasStr(uri), uintptr(resourceContext), uintptr(requestSourceKinds))
	return api.GoBool(r)
}

func (m *TCoreWebView2) RemoveWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext wvTypes.TWVWebResourceContext, requestSourceKinds wvTypes.TWVWebResourceRequestSourceKind) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(39, m.Instance(), api.PasStr(uri), uintptr(resourceContext), uintptr(requestSourceKinds))
	return api.GoBool(r)
}

func (m *TCoreWebView2) PostWebMessageAsJsonWithAdditionalObjects(webMessageAsJson string, additionalObjects ICoreWebView2ObjectCollectionView) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(40, m.Instance(), api.PasStr(webMessageAsJson), base.GetObjectUintptr(additionalObjects))
	return api.GoBool(r)
}

func (m *TCoreWebView2) ShowSaveAsUI(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(41, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2) AfterConstruction() {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(42, m.Instance())
}

func (m *TCoreWebView2) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(43, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) BaseIntf() (result ICoreWebView2) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2API().SysCallN(44, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2(resultPtr)
	return
}

func (m *TCoreWebView2) Settings() (result ICoreWebView2Settings) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2API().SysCallN(45, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Settings(resultPtr)
	return
}

func (m *TCoreWebView2) BrowserProcessID() types.DWORD {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2API().SysCallN(46, m.Instance())
	return types.DWORD(r)
}

func (m *TCoreWebView2) CanGoBack() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(47, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) CanGoForward() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(48, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) ContainsFullScreenElement() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(49, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) DocumentTitle() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2API().SysCallN(50, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2) Source() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2API().SysCallN(51, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2) CookieManager() (result ICoreWebView2CookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2API().SysCallN(52, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieManager(resultPtr)
	return
}

func (m *TCoreWebView2) Environment() (result ICoreWebView2Environment) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2API().SysCallN(53, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Environment(resultPtr)
	return
}

func (m *TCoreWebView2) IsSuspended() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(54, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) IsMuted() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) SetIsMuted(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2) IsDocumentPlayingAudio() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(56, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) IsDefaultDownloadDialogOpen() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2API().SysCallN(57, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2) DefaultDownloadDialogCornerAlignment() wvTypes.TWVDefaultDownloadDialogCornerAlignment {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2API().SysCallN(58, 0, m.Instance())
	return wvTypes.TWVDefaultDownloadDialogCornerAlignment(r)
}

func (m *TCoreWebView2) SetDefaultDownloadDialogCornerAlignment(value wvTypes.TWVDefaultDownloadDialogCornerAlignment) {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2) DefaultDownloadDialogMargin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(59, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2) SetDefaultDownloadDialogMargin(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(59, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2) StatusBarText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2API().SysCallN(60, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2) Profile() (result ICoreWebView2Profile) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2API().SysCallN(61, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Profile(resultPtr)
	return
}

func (m *TCoreWebView2) FaviconURI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2API().SysCallN(62, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2) MemoryUsageTargetLevel() wvTypes.TWVMemoryUsageTargetLevel {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2API().SysCallN(63, 0, m.Instance())
	return wvTypes.TWVMemoryUsageTargetLevel(r)
}

func (m *TCoreWebView2) SetMemoryUsageTargetLevel(value wvTypes.TWVMemoryUsageTargetLevel) {
	if !m.IsValid() {
		return
	}
	coreWebView2API().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2) FrameId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2API().SysCallN(64, m.Instance())
	return uint32(r)
}

// NewCoreWebView2 class constructor
func NewCoreWebView2(baseIntf ICoreWebView2) ICoreWebView2 {
	r := coreWebView2API().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2(r)
}

var (
	coreWebView2Once   base.Once
	coreWebView2Import *imports.Imports = nil
)

func coreWebView2API() *imports.Imports {
	coreWebView2Once.Do(func() {
		coreWebView2Import = api.NewDefaultImports()
		coreWebView2Import.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2_Create", 0), // constructor NewCoreWebView2
			/* 1 */ imports.NewTable("TCoreWebView2_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2_SubscribeToDevToolsProtocolEvent", 0), // function SubscribeToDevToolsProtocolEvent
			/* 3 */ imports.NewTable("TCoreWebView2_CapturePreview", 0), // function CapturePreview
			/* 4 */ imports.NewTable("TCoreWebView2_ExecuteScript", 0), // function ExecuteScript
			/* 5 */ imports.NewTable("TCoreWebView2_GoBack", 0), // function GoBack
			/* 6 */ imports.NewTable("TCoreWebView2_GoForward", 0), // function GoForward
			/* 7 */ imports.NewTable("TCoreWebView2_Navigate", 0), // function Navigate
			/* 8 */ imports.NewTable("TCoreWebView2_NavigateToString", 0), // function NavigateToString
			/* 9 */ imports.NewTable("TCoreWebView2_NavigateWithWebResourceRequest", 0), // function NavigateWithWebResourceRequest
			/* 10 */ imports.NewTable("TCoreWebView2_Reload", 0), // function Reload
			/* 11 */ imports.NewTable("TCoreWebView2_Stop", 0), // function Stop
			/* 12 */ imports.NewTable("TCoreWebView2_TrySuspend", 0), // function TrySuspend
			/* 13 */ imports.NewTable("TCoreWebView2_Resume", 0), // function Resume
			/* 14 */ imports.NewTable("TCoreWebView2_SetVirtualHostNameToFolderMapping", 0), // function SetVirtualHostNameToFolderMapping
			/* 15 */ imports.NewTable("TCoreWebView2_ClearVirtualHostNameToFolderMapping", 0), // function ClearVirtualHostNameToFolderMapping
			/* 16 */ imports.NewTable("TCoreWebView2_OpenTaskManagerWindow", 0), // function OpenTaskManagerWindow
			/* 17 */ imports.NewTable("TCoreWebView2_PrintToPdf", 0), // function PrintToPdf
			/* 18 */ imports.NewTable("TCoreWebView2_OpenDevToolsWindow", 0), // function OpenDevToolsWindow
			/* 19 */ imports.NewTable("TCoreWebView2_PostWebMessageAsJson", 0), // function PostWebMessageAsJson
			/* 20 */ imports.NewTable("TCoreWebView2_PostWebMessageAsString", 0), // function PostWebMessageAsString
			/* 21 */ imports.NewTable("TCoreWebView2_CallDevToolsProtocolMethod", 0), // function CallDevToolsProtocolMethod
			/* 22 */ imports.NewTable("TCoreWebView2_CallDevToolsProtocolMethodForSession", 0), // function CallDevToolsProtocolMethodForSession
			/* 23 */ imports.NewTable("TCoreWebView2_AddWebResourceRequestedFilter", 0), // function AddWebResourceRequestedFilter
			/* 24 */ imports.NewTable("TCoreWebView2_RemoveWebResourceRequestedFilter", 0), // function RemoveWebResourceRequestedFilter
			/* 25 */ imports.NewTable("TCoreWebView2_AddHostObjectToScript", 0), // function AddHostObjectToScript
			/* 26 */ imports.NewTable("TCoreWebView2_RemoveHostObjectFromScript", 0), // function RemoveHostObjectFromScript
			/* 27 */ imports.NewTable("TCoreWebView2_AddScriptToExecuteOnDocumentCreated", 0), // function AddScriptToExecuteOnDocumentCreated
			/* 28 */ imports.NewTable("TCoreWebView2_RemoveScriptToExecuteOnDocumentCreated", 0), // function RemoveScriptToExecuteOnDocumentCreated
			/* 29 */ imports.NewTable("TCoreWebView2_OpenDefaultDownloadDialog", 0), // function OpenDefaultDownloadDialog
			/* 30 */ imports.NewTable("TCoreWebView2_CloseDefaultDownloadDialog", 0), // function CloseDefaultDownloadDialog
			/* 31 */ imports.NewTable("TCoreWebView2_ClearServerCertificateErrorActions", 0), // function ClearServerCertificateErrorActions
			/* 32 */ imports.NewTable("TCoreWebView2_GetFavicon", 0), // function GetFavicon
			/* 33 */ imports.NewTable("TCoreWebView2_Print", 0), // function Print
			/* 34 */ imports.NewTable("TCoreWebView2_ShowPrintUI", 0), // function ShowPrintUI
			/* 35 */ imports.NewTable("TCoreWebView2_PrintToPdfStream", 0), // function PrintToPdfStream
			/* 36 */ imports.NewTable("TCoreWebView2_PostSharedBufferToScript", 0), // function PostSharedBufferToScript
			/* 37 */ imports.NewTable("TCoreWebView2_ExecuteScriptWithResult", 0), // function ExecuteScriptWithResult
			/* 38 */ imports.NewTable("TCoreWebView2_AddWebResourceRequestedFilterWithRequestSourceKinds", 0), // function AddWebResourceRequestedFilterWithRequestSourceKinds
			/* 39 */ imports.NewTable("TCoreWebView2_RemoveWebResourceRequestedFilterWithRequestSourceKinds", 0), // function RemoveWebResourceRequestedFilterWithRequestSourceKinds
			/* 40 */ imports.NewTable("TCoreWebView2_PostWebMessageAsJsonWithAdditionalObjects", 0), // function PostWebMessageAsJsonWithAdditionalObjects
			/* 41 */ imports.NewTable("TCoreWebView2_ShowSaveAsUI", 0), // function ShowSaveAsUI
			/* 42 */ imports.NewTable("TCoreWebView2_AfterConstruction", 0), // procedure AfterConstruction
			/* 43 */ imports.NewTable("TCoreWebView2_Initialized", 0), // property Initialized
			/* 44 */ imports.NewTable("TCoreWebView2_BaseIntf", 0), // property BaseIntf
			/* 45 */ imports.NewTable("TCoreWebView2_Settings", 0), // property Settings
			/* 46 */ imports.NewTable("TCoreWebView2_BrowserProcessID", 0), // property BrowserProcessID
			/* 47 */ imports.NewTable("TCoreWebView2_CanGoBack", 0), // property CanGoBack
			/* 48 */ imports.NewTable("TCoreWebView2_CanGoForward", 0), // property CanGoForward
			/* 49 */ imports.NewTable("TCoreWebView2_ContainsFullScreenElement", 0), // property ContainsFullScreenElement
			/* 50 */ imports.NewTable("TCoreWebView2_DocumentTitle", 0), // property DocumentTitle
			/* 51 */ imports.NewTable("TCoreWebView2_Source", 0), // property Source
			/* 52 */ imports.NewTable("TCoreWebView2_CookieManager", 0), // property CookieManager
			/* 53 */ imports.NewTable("TCoreWebView2_Environment", 0), // property Environment
			/* 54 */ imports.NewTable("TCoreWebView2_IsSuspended", 0), // property IsSuspended
			/* 55 */ imports.NewTable("TCoreWebView2_IsMuted", 0), // property IsMuted
			/* 56 */ imports.NewTable("TCoreWebView2_IsDocumentPlayingAudio", 0), // property IsDocumentPlayingAudio
			/* 57 */ imports.NewTable("TCoreWebView2_IsDefaultDownloadDialogOpen", 0), // property IsDefaultDownloadDialogOpen
			/* 58 */ imports.NewTable("TCoreWebView2_DefaultDownloadDialogCornerAlignment", 0), // property DefaultDownloadDialogCornerAlignment
			/* 59 */ imports.NewTable("TCoreWebView2_DefaultDownloadDialogMargin", 0), // property DefaultDownloadDialogMargin
			/* 60 */ imports.NewTable("TCoreWebView2_StatusBarText", 0), // property StatusBarText
			/* 61 */ imports.NewTable("TCoreWebView2_Profile", 0), // property Profile
			/* 62 */ imports.NewTable("TCoreWebView2_FaviconURI", 0), // property FaviconURI
			/* 63 */ imports.NewTable("TCoreWebView2_MemoryUsageTargetLevel", 0), // property MemoryUsageTargetLevel
			/* 64 */ imports.NewTable("TCoreWebView2_FrameId", 0), // property FrameId
		}
	})
	return coreWebView2Import
}
