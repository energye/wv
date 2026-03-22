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

// ICoreWebView2Frame Parent: IObject
type ICoreWebView2Frame interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// AddHostObjectToScriptWithOrigins
	//  Add the provided host object to script running in the iframe with the
	//  specified name for the list of the specified origins. The host object
	//  will be accessible for this iframe only if the iframe's origin during
	//  access matches one of the origins which are passed. The provided origins
	//  will be normalized before comparing to the origin of the document.
	//  So the scheme name is made lower case, the host will be punycode decoded
	//  as appropriate, default port values will be removed, and so on.
	//  This means the origin's host may be punycode encoded or not and will match
	//  regardless. If list contains malformed origin the call will fail.
	//  The method can be called multiple times in a row without calling
	//  RemoveHostObjectFromScript for the same object name. It will replace
	//  the previous object with the new object and new list of origins.
	//  List of origins will be treated as following:
	//  1. empty list - call will succeed and object will be added for the iframe
	//  but it will not be exposed to any origin;
	//  2. list with origins - during access to host object from iframe the
	//  origin will be checked that it belongs to this list;
	//  3. list with "*" element - host object will be available for iframe for
	//  all origins. We suggest not to use this feature without understanding
	//  security implications of giving access to host object from from iframes
	//  with unknown origins.
	//  4. list with "file://" element - host object will be available for iframes
	//  loaded via file protocol.
	//  Calling this method fails if it is called after the iframe is destroyed.
	//  For more information about host objects navigate to
	//  ICoreWebView2.AddHostObjectToScript.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#addhostobjecttoscriptwithorigins">See the ICoreWebView2Frame article.</see>
	AddHostObjectToScriptWithOrigins(name string, object types.OleVariant, originsCount uint32, origins *string) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the iframe. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the iframe, the JavaScript code continues to have access to that
	//  object. Calling this method for a name that is already removed or was
	//  never added fails. If the iframe is destroyed this method will return fail
	//  also.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#removehostobjectfromscript">See the ICoreWebView2Frame article.</see>
	RemoveHostObjectFromScript(name string) bool // function
	// ExecuteScript
	//  Run JavaScript code from the javascript parameter in the current frame.
	//  The result of evaluating the provided JavaScript is passed to the completion handler.
	//  The result value is a JSON encoded string. If the result is undefined,
	//  contains a reference cycle, or otherwise is not able to be encoded into
	//  JSON, then the result is considered to be null, which is encoded
	//  in JSON as the string "null".
	//  NOTE: A function that has no explicit return value returns undefined. If the
	//  script that was run throws an unhandled exception, then the result is
	//  also "null". This method is applied asynchronously. If the method is
	//  run before `ContentLoading`, the script will not be executed
	//  and the string "null" will be returned.
	//  This operation executes the script even if `ICoreWebView2Settings::IsScriptEnabled` is
	//  set to `FALSE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#executescript">See the ICoreWebView2Frame2 article.</see>
	ExecuteScript(javaScript string, executionID int32, browserComponent lcl.IComponent) bool // function
	// PostWebMessageAsJson
	//  Posts the specified webMessage to the frame.
	//  The frame receives the message by subscribing to the `message` event of
	//  the `window.chrome.webview` of the frame document.
	//
	//  <code>
	//  ```cpp
	//  window.chrome.webview.addEventListener('message', handler)
	//  window.chrome.webview.removeEventListener('message', handler)
	//  ```</code>
	//
	//  The event args is an instances of `MessageEvent`. The
	//  `ICoreWebView2Settings::IsWebMessageEnabled` setting must be `TRUE` or
	//  the message will not be sent. The `data` property of the event
	//  args is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event args is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before the
	//  message is posted to the page, the message is discarded.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#postwebmessageasjson">See the ICoreWebView2Frame2 article.</see>
	PostWebMessageAsJson(webMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  args of the `window.chrome.webview` message is a string with the same
	//  value as `webMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#postwebmessageasstring">See the ICoreWebView2Frame2 article.</see>
	PostWebMessageAsString(webMessageAsString string) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the iframe in the WebView.
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
	//  For example, if we want to send data to script for one time read only consumption.
	//  Sharing a buffer to script has security risk. You should only share buffer with trusted site.
	//  If a buffer is shared to a untrusted site, possible sensitive information could be leaked.
	//  If a buffer is shared as modifiable by the script and the script modifies it in an unexpected way,
	//  it could result in corrupted data that might even crash the application.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame4#postsharedbuffertoscript">See the ICoreWebView2Frame4 article.</see>
	PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Frame // property BaseIntf Getter
	// FrameID
	//  The unique identifier of the current frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2` and via `CoreWebView2FrameInfo`.
	FrameID() uint32 // property FrameID Getter
	// Name
	//  The value of iframe's window.name property. The default value equals to
	//  iframe html tag declaring it. You can access this property even if the
	//  iframe is destroyed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#get_name">See the ICoreWebView2Frame article.</see>
	Name() string // property Name Getter
	// IsDestroyed
	//  Check whether a frame is destroyed. Returns true during
	//  the Destroyed event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#isdestroyed">See the ICoreWebView2Frame article.</see>
	IsDestroyed() bool // property IsDestroyed Getter
}

type TCoreWebView2Frame struct {
	TObject
}

func (m *TCoreWebView2Frame) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) AddHostObjectToScriptWithOrigins(name string, object types.OleVariant, originsCount uint32, origins *string) bool {
	if !m.IsValid() {
		return false
	}
	originsPtr := api.PasStr(*origins)
	r := coreWebView2FrameAPI().SysCallN(2, m.Instance(), api.PasStr(name), uintptr(object), uintptr(originsCount), uintptr(base.UnsafePointer(&originsPtr)))
	*origins = api.GoStr(originsPtr)
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) RemoveHostObjectFromScript(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(3, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) ExecuteScript(javaScript string, executionID int32, browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(4, m.Instance(), api.PasStr(javaScript), uintptr(executionID), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) PostWebMessageAsJson(webMessageAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(5, m.Instance(), api.PasStr(webMessageAsJson))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) PostWebMessageAsString(webMessageAsString string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(6, m.Instance(), api.PasStr(webMessageAsString))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) PostSharedBufferToScript(sharedBuffer ICoreWebView2SharedBuffer, access wvTypes.TWVSharedBufferAccess, additionalDataAsJson string) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(sharedBuffer), uintptr(access), api.PasStr(additionalDataAsJson))
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Frame) BaseIntf() (result ICoreWebView2Frame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Frame(resultPtr)
	return
}

func (m *TCoreWebView2Frame) FrameID() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2FrameAPI().SysCallN(10, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2Frame) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2FrameAPI().SysCallN(11, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Frame) IsDestroyed() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2Frame class constructor
func NewCoreWebView2Frame(baseIntf ICoreWebView2Frame) ICoreWebView2Frame {
	r := coreWebView2FrameAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Frame(r)
}

var (
	coreWebView2FrameOnce   base.Once
	coreWebView2FrameImport *imports.Imports = nil
)

func coreWebView2FrameAPI() *imports.Imports {
	coreWebView2FrameOnce.Do(func() {
		coreWebView2FrameImport = api.NewDefaultImports()
		coreWebView2FrameImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Frame_Create", 0), // constructor NewCoreWebView2Frame
			/* 1 */ imports.NewTable("TCoreWebView2Frame_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2Frame_AddHostObjectToScriptWithOrigins", 0), // function AddHostObjectToScriptWithOrigins
			/* 3 */ imports.NewTable("TCoreWebView2Frame_RemoveHostObjectFromScript", 0), // function RemoveHostObjectFromScript
			/* 4 */ imports.NewTable("TCoreWebView2Frame_ExecuteScript", 0), // function ExecuteScript
			/* 5 */ imports.NewTable("TCoreWebView2Frame_PostWebMessageAsJson", 0), // function PostWebMessageAsJson
			/* 6 */ imports.NewTable("TCoreWebView2Frame_PostWebMessageAsString", 0), // function PostWebMessageAsString
			/* 7 */ imports.NewTable("TCoreWebView2Frame_PostSharedBufferToScript", 0), // function PostSharedBufferToScript
			/* 8 */ imports.NewTable("TCoreWebView2Frame_Initialized", 0), // property Initialized
			/* 9 */ imports.NewTable("TCoreWebView2Frame_BaseIntf", 0), // property BaseIntf
			/* 10 */ imports.NewTable("TCoreWebView2Frame_FrameID", 0), // property FrameID
			/* 11 */ imports.NewTable("TCoreWebView2Frame_Name", 0), // property Name
			/* 12 */ imports.NewTable("TCoreWebView2Frame_IsDestroyed", 0), // property IsDestroyed
		}
	})
	return coreWebView2FrameImport
}
