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

// ICoreWebView2Environment Parent: lcl.IObject
type ICoreWebView2Environment interface {
	lcl.IObject
	// AddAllLoaderEvents
	//  Adds all the events of this class to an existing TWVLoader instance.
	//  <param name="aLoaderComponent">The TWVLoader instance.</param>
	AddAllLoaderEvents(loaderComponent lcl.IComponent) bool // function
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// CreateCoreWebView2Controller
	//  Asynchronously create a new WebView.
	//  <param name="aParentWindow">Handle of the control in which the WebView should be displayed.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#createcorewebview2controller">See the ICoreWebView2Environment article.</see>
	CreateCoreWebView2Controller(parentWindow types.THandle, browserEvents IWVBrowserBase, result *types.HRESULT) bool // function
	// CreateWebResourceResponse
	//  Create a new web resource response object.
	//  <param name="aContent">HTTP response content as stream.</param>
	//  <param name="aStatusCode">The HTTP response status code.</param>
	//  <param name="aReasonPhrase">The HTTP response reason phrase.</param>
	//  <param name="aHeaders">Overridden HTTP response headers.</param>
	//  <param name="aResponse">The new ICoreWebView2WebResourceResponse instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#createcorewebview2controller">See the ICoreWebView2Environment article.</see>
	CreateWebResourceResponse(content lcl.IStreamAdapter, statusCode int32, reasonPhrase string, headers string, response *ICoreWebView2WebResourceResponse) bool // function
	// CreateWebResourceRequest
	//  Create a new web resource request object.
	//  URI parameter must be absolute URI.
	//  The headers string is the raw request header string delimited by CRLF
	//  (optional in last header).
	//  It's also possible to create this object with null headers string
	//  and then use the ICoreWebView2HttpRequestHeaders to construct the headers
	//  line by line.
	//  <param name="aURI">The request URI.</param>
	//  <param name="aMethod">The HTTP request method.</param>
	//  <param name="aPostData">The HTTP request message body as stream.</param>
	//  <param name="aHeaders">The mutable HTTP request headers.</param>
	//  <param name="aRequest">The new ICoreWebView2WebResourceRequest instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment2#createwebresourcerequest">See the ICoreWebView2Environment2 article.</see>
	CreateWebResourceRequest(uRI string, method string, postData lcl.IStreamAdapter, headers string, request *ICoreWebView2WebResourceRequest) bool // function
	// CreateCoreWebView2CompositionController
	//  Asynchronously create a new WebView for use with visual hosting.
	//  <param name="aParentWindow">Handle of the control in which the app will connect the visual tree of the WebView.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3#createcorewebview2compositioncontroller">See the ICoreWebView2Environment3 article.</see>
	CreateCoreWebView2CompositionController(parentWindow types.THandle, browserEvents IWVBrowserBase, result *types.HRESULT) bool // function
	// CreateCoreWebView2PointerInfo
	//  Create an empty ICoreWebView2PointerInfo. The returned
	//  ICoreWebView2PointerInfo needs to be populated with all of the relevant
	//  info before calling SendPointerInput.
	//  <param name="aPointerInfo">The new ICoreWebView2PointerInfo instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3#createcorewebview2pointerinfo">See the ICoreWebView2Environment3 article.</see>
	CreateCoreWebView2PointerInfo(pointerInfo *ICoreWebView2PointerInfo) bool // function
	// CreatePrintSettings
	//  Creates the `ICoreWebView2PrintSettings` used by the `PrintToPdf` method.
	//  <param name="aPrintSettings">The new ICoreWebView2PrintSettings instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment6#createprintsettings">See the ICoreWebView2Environment6 article.</see>
	CreatePrintSettings(printSettings *ICoreWebView2PrintSettings) bool // function
	// CreateContextMenuItem
	//  Create a custom `ContextMenuItem` object to insert into the WebView context menu.
	//  CoreWebView2 will rewind the icon stream before decoding.
	//  There is a limit of 1000 active custom context menu items at a given time.
	//  Attempting to create more before deleting existing ones will fail with
	//  ERROR_NOT_ENOUGH_QUOTA.
	//  It is recommended to reuse ContextMenuItems across ContextMenuRequested events
	//  for performance.
	//  The returned ContextMenuItem object's `IsEnabled` property will default to `TRUE`
	//  and `IsChecked` property will default to `FALSE`. A `CommandId` will be assigned
	//  to the ContextMenuItem object that's unique across active custom context menu items,
	//  but command ID values of deleted ContextMenuItems can be reassigned.
	//  <param name="aLabel">Context menu item label.</param>
	//  <param name="aIconStream">Context menu item icon as stream.</param>
	//  <param name="aKind">Context menu item kind.</param>
	//  <param name="aMenuItem">The new ICoreWebView2ContextMenuItem instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment9#createcontextmenuitem">See the ICoreWebView2Environment9 article.</see>
	CreateContextMenuItem(label string, iconStream lcl.IStreamAdapter, kind wvTypes.TWVMenuItemKind, menuItem *ICoreWebView2ContextMenuItem) bool // function
	// CreateCoreWebView2ControllerOptions
	//  Create a new ICoreWebView2ControllerOptions to be passed as a parameter of
	//  CreateCoreWebView2ControllerWithOptions and CreateCoreWebView2CompositionControllerWithOptions.
	//  The 'options' is settable and in it the default value for profile name is the empty string,
	//  and the default value for IsInPrivateModeEnabled is false.
	//  Also the profile name can be reused.
	//  <param name="aOptions">The new ICoreWebView2ControllerOptions instance.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2controlleroptions">See the ICoreWebView2Environment10 article.</see>
	CreateCoreWebView2ControllerOptions(options *ICoreWebView2ControllerOptions, result *types.HRESULT) bool // function
	// CreateCoreWebView2ControllerWithOptions
	//  Create a new WebView with options.
	//  <param name="aParentWindow">Handle of the control in which the WebView should be displayed.</param>
	//  <param name="aOptions">The ICoreWebView2ControllerOptions instance created with CreateCoreWebView2ControllerOptions.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2controllerwithoptions">See the ICoreWebView2Environment10 article.</see>
	CreateCoreWebView2ControllerWithOptions(parentWindow types.HWND, options ICoreWebView2ControllerOptions, browserEvents IWVBrowserBase, result *types.HRESULT) bool // function
	// CreateCoreWebView2CompositionControllerWithOptions
	//  Create a new WebView in visual hosting mode with options.
	//  <param name="aParentWindow">Handle of the control in which the app will connect the visual tree of the WebView.</param>
	//  <param name="aOptions">The ICoreWebView2ControllerOptions instance created with CreateCoreWebView2ControllerOptions.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2compositioncontrollerwithoptions">See the ICoreWebView2Environment10 article.</see>
	CreateCoreWebView2CompositionControllerWithOptions(parentWindow types.HWND, options ICoreWebView2ControllerOptions, browserEvents IWVBrowserBase, result *types.HRESULT) bool // function
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
	//  <param name="aSize">Buffer size in bytes.</param>
	//  <param name="aSharedBuffer">The new ICoreWebView2SharedBuffer instance.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the ICoreWebView2Environment12 article.</see>
	CreateSharedBuffer(size int64, sharedBuffer *ICoreWebView2SharedBuffer) bool // function
	// GetProcessExtendedInfos
	//  Gets a snapshot collection of `ProcessExtendedInfo`s corresponding to all
	//  currently running processes associated with this `ICoreWebView2Environment`
	//  excludes crashpad process.
	//  This provides the same list of `ProcessInfo`s as what's provided in
	//  `GetProcessInfos`, but additionally provides a list of associated `FrameInfo`s
	//  which are actively running (showing or hiding UI elements) in the renderer
	//  process. See `AssociatedFrameInfos` for more information.
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment13#getprocessextendedinfos">See the ICoreWebView2Environment13 article.</see>
	GetProcessExtendedInfos(browserEvents IWVBrowserBase) bool // function
	// CreateWebFileSystemFileHandle
	//  Create a `ICoreWebView2FileSystemHandle` object from a path that represents a Web
	//  [FileSystemFileHandle](https://developer.mozilla.org/docs/Web/API/FileSystemFileHandle).
	//
	//  The `path` is the path pointed by the file and must be a syntactically correct fully qualified
	//  path, but it is not checked here whether it currently points to a file. If an invalid path is
	//  passed, an E_INVALIDARG will be returned and the object will fail to create. Any other state
	//  validation will be done when this handle is accessed from web content
	//  and will cause the DOM exceptions described in
	//  [FileSystemFileHandle methods](https://developer.mozilla.org/docs/Web/API/FileSystemDirectoryHandle#instance_methods)
	//  if access operations fail.
	//
	//  `Permission` property is used to specify whether the Handle should be created with a Read-only or
	//  Read-and-write web permission. For the `permission` value specified here, the DOM
	//  [PermissionStatus](https://developer.mozilla.org/docs/Web/API/PermissionStatus) property
	//  will be [granted](https://developer.mozilla.org/docs/Web/API/PermissionStatus/state)
	//  and the unspecified permission will be
	//  [prompt](https://developer.mozilla.org/docs/Web/API/PermissionStatus/state). Therefore,
	//  the web content then does not need to call
	//  [requestPermission](https://developer.mozilla.org/docs/Web/API/FileSystemHandle/requestPermission)
	//  for the permission that was specified before attempting the permitted operation,
	//  but if it does, the promise will immediately be resolved
	//  with 'granted' PermissionStatus without firing the WebView2
	//  [PermissionRequested](/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs)
	//  event or prompting the user for permission. Otherwise, `requestPermission` will behave as the
	//  status of permission is currently `prompt`, which means the `PermissionRequested` event will fire
	//  or the user will be prompted.
	//  Additionally, the app must have the same OS permissions that have propagated to the
	//  [WebView2 browser process](/microsoft-edge/webview2/concepts/process-model)
	//  for the path it wishes to give the web content to read/write the file.
	//  Specifically, the WebView2 browser process will run in same user, package identity, and app
	//  container of the app, but other means such as security context impersonations do not get
	//  propagated, so such permissions that the app has, will not be effective in WebView2.
	//  Note: An exception to this is, if there is a parent directory handle that
	//  has broader permissions in the same page context than specified in here, the handle will automatically
	//  inherit the most permissive permission of the parent handle when posted to that page context.
	//  i.e. If there is already a `FileSystemDirectoryHandle` to `C:\example` that has write permission on
	//  a page, even though a WebFileSystemHandle to file `C:\example\file.txt` is created with
	//  `COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_ONLY` permission, when posted to that page, write permission
	//  will be automatically granted to the created handle.
	//
	//  An app needs to be mindful that this object, when posted to the web content, provides it with unusual
	//  access to OS file system via the Web FileSystem API! The app should therefore only post objects
	//  for paths that it wants to allow access to the web content and it is not recommended that the web content
	//  "asks" for this path. The app should also check the source property of the target to ensure
	//  that it is sending to the web content of intended origin.
	//
	//  Once the object is passed to web content, if the content is attempting a read,
	//  the file must be existing and available to read similar to a file chosen by
	//  [open file picker](https://developer.mozilla.org/docs/Web/API/Window/showOpenFilePicker),
	//  otherwise the read operation will
	//  [throw a DOM exception](https://developer.mozilla.org/docs/Web/API/FileSystemFileHandle/getFile#exceptions).
	//  For write operations, the file does not need to exist as `FileSystemFileHandle` will behave
	//  as a file path chosen by
	//  [save file picker](https://developer.mozilla.org/docs/Web/API/Window/showSaveFilePicker)
	//  and will create or overwrite the file, but the parent directory structure pointed
	//  by the file must exist and an existing file must be available to write and delete
	//  or the write operation will
	//  [throw a DOM exception](https://developer.mozilla.org/docs/Web/API/FileSystemFileHandle/createWritable#exceptions).
	//  <param name="aPath">The path pointed by the file.</param>
	//  <param name="aPermission">Used to specify whether the Handle should be created with a Read-only or Read-and-write web permission.</param>
	//  <param name="aValue">The ICoreWebView2FileSystemHandle object created from a path that represents a Web FileSystemFileHandle.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment14#createwebfilesystemfilehandle">See the ICoreWebView2Environment14 article.</see>
	CreateWebFileSystemFileHandle(path string, permission wvTypes.TWVFileSystemHandlePermission, value *ICoreWebView2FileSystemHandle) bool // function
	// CreateWebFileSystemDirectoryHandle
	//  Create a `ICoreWebView2FileSystemHandle` object from a path that represents a Web
	//  [FileSystemDirectoryHandle](https://developer.mozilla.org/docs/Web/API/FileSystemDirectoryHandle).
	//
	//  The `path` is the path pointed by the directory and must be a syntactically correct fully qualified
	//  path, but it is not checked here whether it currently points to a directory. Any other state
	//  validation will be done when this handle is accessed from web content
	//  and will cause DOM exceptions if access operations fail. If an invalid path is
	//  passed, an E_INVALIDARG will be returned and the object will fail to create.
	//
	//  `Permission` property is used to specify whether the Handle should be created with a Read-only or
	//  Read-and-write web permission. For the `permission` value specified here, the Web
	//  [PermissionStatus](https://developer.mozilla.org/docs/Web/API/PermissionStatus)
	//  will be [granted](https://developer.mozilla.org/docs/Web/API/PermissionStatus/state)
	//  and the unspecified permission will be
	//  [prompt](https://developer.mozilla.org/docs/Web/API/PermissionStatus/state). Therefore,
	//  the web content then does not need to call
	//  [requestPermission](https://developer.mozilla.org/docs/Web/API/FileSystemHandle/requestPermission)
	//  for the permission that was specified before attempting the permitted operation,
	//  but if it does, the promise will immediately be resolved
	//  with 'granted' PermissionStatus without firing the WebView2
	//  [PermissionRequested](/microsoft-edge/webview2/reference/win32/icorewebview2permissionrequestedeventargs)
	//  event or prompting the user for permission. Otherwise, `requestPermission` will behave as the
	//  status of permission is currently `Prompt`, which means the `PermissionRequested` event will fire
	//  or the user will be prompted.
	//  Additionally, the app must have the same OS permissions that have propagated to the
	//  [WebView2 browser process](/microsoft-edge/webview2/concepts/process-model)
	//  for the path it wishes to give the web content to make any operations on the directory.
	//  Specifically, the WebView2 browser process will run in same user, package identity, and app
	//  container of the app, but other means such as security context impersonations do not get
	//  propagated, so such permissions that the app has, will not be effective in WebView2.
	//  Note: An exception to this is, if there is a parent directory handle that
	//  has broader permissions in the same page context than specified in here, the handle will automatically
	//  inherit the most permissive permission of the parent handle when posted to that page context.
	//  i.e. If there is already a `FileSystemDirectoryHandle` to `C:\example` that has write permission on
	//  a page, even though a WebFileSystemHandle to directory `C:\example\directory` is created with
	//  `COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_ONLY` permission, when posted to that page, write permission
	//  will be automatically granted to the created handle.
	//
	//  An app needs to be mindful that this object, when posted to the web content, provides it with unusual
	//  access to OS file system via the Web FileSystem API! The app should therefore only post objects
	//  for paths that it wants to allow access to the web content and it is not recommended that the web content
	//  "asks" for this path. The app should also check the source property of the target to ensure
	//  that it is sending to the web content of intended origin.
	//
	//  Once the object is passed to web content, the path must point to a directory as if it was chosen via
	//  [directory picker](https://developer.mozilla.org/docs/Web/API/Window/showDirectoryPicker)
	//  otherwise any IO operation done on the `FileSystemDirectoryHandle` will throw a DOM exception.
	//  <param name="aPath">The path pointed by the directory.</param>
	//  <param name="aPermission">Used to specify whether the Handle should be created with a Read-only or Read-and-write web permission.</param>
	//  <param name="aValue">The ICoreWebView2FileSystemHandle object created from a path that represents a Web FileSystemDirectoryHandle.</param>
	//  <returns>True if successfull.</return>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment14#createwebfilesystemdirectoryhandle">See the ICoreWebView2Environment14 article.</see>
	CreateWebFileSystemDirectoryHandle(path string, permission wvTypes.TWVFileSystemHandlePermission, value *ICoreWebView2FileSystemHandle) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Environment // property BaseIntf Getter
	// BrowserVersionInfo
	//  The browser version info of the current `ICoreWebView2Environment`,
	//  including channel name if it is not the WebView2 Runtime. It matches the
	//  format of the `GetAvailableCoreWebView2BrowserVersionString` API.
	//  Channel names are `beta`, `dev`, and `canary`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#get_browserversionstring">See the ICoreWebView2Environment article.</see>
	BrowserVersionInfo() string // property BrowserVersionInfo Getter
	// SupportsCompositionController
	//  Returns true if the current WebView2 runtime version supports Composition Controllers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3">See the ICoreWebView2Environment3 article.</see>
	SupportsCompositionController() bool // property SupportsCompositionController Getter
	// SupportsControllerOptions
	//  Returns true if the current WebView2 runtime version supports Controller Options.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10">See the ICoreWebView2Environment10 article.</see>
	SupportsControllerOptions() bool // property SupportsControllerOptions Getter
	// UserDataFolder
	//  Returns the user data folder that all CoreWebView2's created from this
	//  environment are using.
	//  This could be either the value passed in by the developer when creating
	//  the environment object or the calculated one for default handling. It
	//  will always be an absolute path.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment7#get_userdatafolder">See the ICoreWebView2Environment7 article.</see>
	UserDataFolder() string // property UserDataFolder Getter
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`. Provide a list of all
	//  process using same user data folder except for crashpad process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8#getprocessinfos">See the ICoreWebView2Environment8 article.</see>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property ProcessInfos Getter
	// FailureReportFolderPath
	//  `FailureReportFolderPath` returns the path of the folder where minidump files are written.
	//  Whenever a WebView2 process crashes, a crash dump file will be created in the crash dump folder.
	//  The crash dump format is minidump files. Please see
	//  [Minidump Files documentation](/windows/win32/debug/minidump-files) for detailed information.
	//  Normally when a single child process fails, a minidump will be generated and written to disk,
	//  then the `ProcessFailed` event is raised. But for unexpected crashes, a minidump file might not be generated
	//  at all, despite whether `ProcessFailed` event is raised. If there are multiple
	//  process failures at once, multiple minidump files could be generated. Thus `FailureReportFolderPath`
	//  could contain old minidump files that are not associated with a specific `ProcessFailed` event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment11#get_failurereportfolderpath">See the ICoreWebView2Environment11 article.</see>
	FailureReportFolderPath() string // property FailureReportFolderPath Getter
}

type TCoreWebView2Environment struct {
	lcl.TObject
}

func (m *TCoreWebView2Environment) AddAllLoaderEvents(loaderComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(loaderComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2Controller(parentWindow types.THandle, browserEvents IWVBrowserBase, result *types.HRESULT) bool {
	if !m.IsValid() {
		return false
	}
	resultPtr := uintptr(*result)
	r := coreWebView2EnvironmentAPI().SysCallN(3, m.Instance(), uintptr(parentWindow), base.GetObjectUintptr(browserEvents), uintptr(base.UnsafePointer(&resultPtr)))
	*result = types.HRESULT(resultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateWebResourceResponse(content lcl.IStreamAdapter, statusCode int32, reasonPhrase string, headers string, response *ICoreWebView2WebResourceResponse) bool {
	if !m.IsValid() {
		return false
	}
	responsePtr := base.GetObjectUintptr(*response)
	r := coreWebView2EnvironmentAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(content), uintptr(statusCode), api.PasStr(reasonPhrase), api.PasStr(headers), uintptr(base.UnsafePointer(&responsePtr)))
	*response = AsCoreWebView2WebResourceResponse(responsePtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateWebResourceRequest(uRI string, method string, postData lcl.IStreamAdapter, headers string, request *ICoreWebView2WebResourceRequest) bool {
	if !m.IsValid() {
		return false
	}
	requestPtr := base.GetObjectUintptr(*request)
	r := coreWebView2EnvironmentAPI().SysCallN(5, m.Instance(), api.PasStr(uRI), api.PasStr(method), base.GetObjectUintptr(postData), api.PasStr(headers), uintptr(base.UnsafePointer(&requestPtr)))
	*request = AsCoreWebView2WebResourceRequestOwn(requestPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2CompositionController(parentWindow types.THandle, browserEvents IWVBrowserBase, result *types.HRESULT) bool {
	if !m.IsValid() {
		return false
	}
	resultPtr := uintptr(*result)
	r := coreWebView2EnvironmentAPI().SysCallN(6, m.Instance(), uintptr(parentWindow), base.GetObjectUintptr(browserEvents), uintptr(base.UnsafePointer(&resultPtr)))
	*result = types.HRESULT(resultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2PointerInfo(pointerInfo *ICoreWebView2PointerInfo) bool {
	if !m.IsValid() {
		return false
	}
	pointerInfoPtr := base.GetObjectUintptr(*pointerInfo)
	r := coreWebView2EnvironmentAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&pointerInfoPtr)))
	*pointerInfo = AsCoreWebView2PointerInfo(pointerInfoPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreatePrintSettings(printSettings *ICoreWebView2PrintSettings) bool {
	if !m.IsValid() {
		return false
	}
	printSettingsPtr := base.GetObjectUintptr(*printSettings)
	r := coreWebView2EnvironmentAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&printSettingsPtr)))
	*printSettings = AsCoreWebView2PrintSettings(printSettingsPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateContextMenuItem(label string, iconStream lcl.IStreamAdapter, kind wvTypes.TWVMenuItemKind, menuItem *ICoreWebView2ContextMenuItem) bool {
	if !m.IsValid() {
		return false
	}
	menuItemPtr := base.GetObjectUintptr(*menuItem)
	r := coreWebView2EnvironmentAPI().SysCallN(9, m.Instance(), api.PasStr(label), base.GetObjectUintptr(iconStream), uintptr(kind), uintptr(base.UnsafePointer(&menuItemPtr)))
	*menuItem = AsCoreWebView2ContextMenuItem(menuItemPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2ControllerOptions(options *ICoreWebView2ControllerOptions, result *types.HRESULT) bool {
	if !m.IsValid() {
		return false
	}
	optionsPtr := base.GetObjectUintptr(*options)
	resultPtr := uintptr(*result)
	r := coreWebView2EnvironmentAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&optionsPtr)), uintptr(base.UnsafePointer(&resultPtr)))
	*options = AsCoreWebView2ControllerOptions(optionsPtr)
	*result = types.HRESULT(resultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2ControllerWithOptions(parentWindow types.HWND, options ICoreWebView2ControllerOptions, browserEvents IWVBrowserBase, result *types.HRESULT) bool {
	if !m.IsValid() {
		return false
	}
	resultPtr := uintptr(*result)
	r := coreWebView2EnvironmentAPI().SysCallN(11, m.Instance(), uintptr(parentWindow), base.GetObjectUintptr(options), base.GetObjectUintptr(browserEvents), uintptr(base.UnsafePointer(&resultPtr)))
	*result = types.HRESULT(resultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2CompositionControllerWithOptions(parentWindow types.HWND, options ICoreWebView2ControllerOptions, browserEvents IWVBrowserBase, result *types.HRESULT) bool {
	if !m.IsValid() {
		return false
	}
	resultPtr := uintptr(*result)
	r := coreWebView2EnvironmentAPI().SysCallN(12, m.Instance(), uintptr(parentWindow), base.GetObjectUintptr(options), base.GetObjectUintptr(browserEvents), uintptr(base.UnsafePointer(&resultPtr)))
	*result = types.HRESULT(resultPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateSharedBuffer(size int64, sharedBuffer *ICoreWebView2SharedBuffer) bool {
	if !m.IsValid() {
		return false
	}
	sharedBufferPtr := base.GetObjectUintptr(*sharedBuffer)
	r := coreWebView2EnvironmentAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&size)), uintptr(base.UnsafePointer(&sharedBufferPtr)))
	*sharedBuffer = AsCoreWebView2SharedBuffer(sharedBufferPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) GetProcessExtendedInfos(browserEvents IWVBrowserBase) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(browserEvents))
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateWebFileSystemFileHandle(path string, permission wvTypes.TWVFileSystemHandlePermission, value *ICoreWebView2FileSystemHandle) bool {
	if !m.IsValid() {
		return false
	}
	valuePtr := base.GetObjectUintptr(*value)
	r := coreWebView2EnvironmentAPI().SysCallN(15, m.Instance(), api.PasStr(path), uintptr(permission), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCoreWebView2FileSystemHandle(valuePtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) CreateWebFileSystemDirectoryHandle(path string, permission wvTypes.TWVFileSystemHandlePermission, value *ICoreWebView2FileSystemHandle) bool {
	if !m.IsValid() {
		return false
	}
	valuePtr := base.GetObjectUintptr(*value)
	r := coreWebView2EnvironmentAPI().SysCallN(16, m.Instance(), api.PasStr(path), uintptr(permission), uintptr(base.UnsafePointer(&valuePtr)))
	*value = AsCoreWebView2FileSystemHandle(valuePtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) BaseIntf() (result ICoreWebView2Environment) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2EnvironmentAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Environment(resultPtr)
	return
}

func (m *TCoreWebView2Environment) BrowserVersionInfo() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2EnvironmentAPI().SysCallN(19, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Environment) SupportsCompositionController() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) SupportsControllerOptions() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2EnvironmentAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Environment) UserDataFolder() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2EnvironmentAPI().SysCallN(22, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Environment) ProcessInfos() (result ICoreWebView2ProcessInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2EnvironmentAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfoCollection(resultPtr)
	return
}

func (m *TCoreWebView2Environment) FailureReportFolderPath() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2EnvironmentAPI().SysCallN(24, m.Instance())
	return api.GoStr(r)
}

// NewCoreWebView2Environment class constructor
func NewCoreWebView2Environment(baseIntf ICoreWebView2Environment) ICoreWebView2Environment {
	r := coreWebView2EnvironmentAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Environment(r)
}

var (
	coreWebView2EnvironmentOnce   base.Once
	coreWebView2EnvironmentImport *imports.Imports = nil
)

func coreWebView2EnvironmentAPI() *imports.Imports {
	coreWebView2EnvironmentOnce.Do(func() {
		coreWebView2EnvironmentImport = api.NewDefaultImports()
		coreWebView2EnvironmentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Environment_Create", 0), // constructor NewCoreWebView2Environment
			/* 1 */ imports.NewTable("TCoreWebView2Environment_AddAllLoaderEvents", 0), // function AddAllLoaderEvents
			/* 2 */ imports.NewTable("TCoreWebView2Environment_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 3 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2Controller", 0), // function CreateCoreWebView2Controller
			/* 4 */ imports.NewTable("TCoreWebView2Environment_CreateWebResourceResponse", 0), // function CreateWebResourceResponse
			/* 5 */ imports.NewTable("TCoreWebView2Environment_CreateWebResourceRequest", 0), // function CreateWebResourceRequest
			/* 6 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2CompositionController", 0), // function CreateCoreWebView2CompositionController
			/* 7 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2PointerInfo", 0), // function CreateCoreWebView2PointerInfo
			/* 8 */ imports.NewTable("TCoreWebView2Environment_CreatePrintSettings", 0), // function CreatePrintSettings
			/* 9 */ imports.NewTable("TCoreWebView2Environment_CreateContextMenuItem", 0), // function CreateContextMenuItem
			/* 10 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2ControllerOptions", 0), // function CreateCoreWebView2ControllerOptions
			/* 11 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2ControllerWithOptions", 0), // function CreateCoreWebView2ControllerWithOptions
			/* 12 */ imports.NewTable("TCoreWebView2Environment_CreateCoreWebView2CompositionControllerWithOptions", 0), // function CreateCoreWebView2CompositionControllerWithOptions
			/* 13 */ imports.NewTable("TCoreWebView2Environment_CreateSharedBuffer", 0), // function CreateSharedBuffer
			/* 14 */ imports.NewTable("TCoreWebView2Environment_GetProcessExtendedInfos", 0), // function GetProcessExtendedInfos
			/* 15 */ imports.NewTable("TCoreWebView2Environment_CreateWebFileSystemFileHandle", 0), // function CreateWebFileSystemFileHandle
			/* 16 */ imports.NewTable("TCoreWebView2Environment_CreateWebFileSystemDirectoryHandle", 0), // function CreateWebFileSystemDirectoryHandle
			/* 17 */ imports.NewTable("TCoreWebView2Environment_Initialized", 0), // property Initialized
			/* 18 */ imports.NewTable("TCoreWebView2Environment_BaseIntf", 0), // property BaseIntf
			/* 19 */ imports.NewTable("TCoreWebView2Environment_BrowserVersionInfo", 0), // property BrowserVersionInfo
			/* 20 */ imports.NewTable("TCoreWebView2Environment_SupportsCompositionController", 0), // property SupportsCompositionController
			/* 21 */ imports.NewTable("TCoreWebView2Environment_SupportsControllerOptions", 0), // property SupportsControllerOptions
			/* 22 */ imports.NewTable("TCoreWebView2Environment_UserDataFolder", 0), // property UserDataFolder
			/* 23 */ imports.NewTable("TCoreWebView2Environment_ProcessInfos", 0), // property ProcessInfos
			/* 24 */ imports.NewTable("TCoreWebView2Environment_FailureReportFolderPath", 0), // property FailureReportFolderPath
		}
	})
	return coreWebView2EnvironmentImport
}
