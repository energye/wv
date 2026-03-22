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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2Settings Parent: IObject
type ICoreWebView2Settings interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Settings // property BaseIntf Getter
	// IsBuiltInErrorPageEnabled
	//  The `IsBuiltInErrorPageEnabled` property is used to disable built in
	//  error page for navigation failure and render process failure. When
	//  disabled, a blank page is displayed when the related error happens.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isbuiltinerrorpageenabled">See the ICoreWebView2Settings article.</see>
	IsBuiltInErrorPageEnabled() bool         // property IsBuiltInErrorPageEnabled Getter
	SetIsBuiltInErrorPageEnabled(value bool) // property IsBuiltInErrorPageEnabled Setter
	// AreDefaultContextMenusEnabled
	//  The `AreDefaultContextMenusEnabled` property is used to prevent default
	//  context menus from being shown to user in WebView.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultcontextmenusenabled">See the ICoreWebView2Settings article.</see>
	AreDefaultContextMenusEnabled() bool         // property AreDefaultContextMenusEnabled Getter
	SetAreDefaultContextMenusEnabled(value bool) // property AreDefaultContextMenusEnabled Setter
	// AreDefaultScriptDialogsEnabled
	//  `AreDefaultScriptDialogsEnabled` is used when loading a new HTML
	//  document. If set to `FALSE`, WebView2 does not render the default JavaScript
	//  dialog box (Specifically those displayed by the JavaScript alert,
	//  confirm, prompt functions and `beforeunload` event). Instead, if an
	//  event handler is set using `add_ScriptDialogOpening`, WebView sends an
	//  event that contains all of the information for the dialog and allow the
	//  host app to show a custom UI.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultscriptdialogsenabled">See the ICoreWebView2Settings article.</see>
	AreDefaultScriptDialogsEnabled() bool         // property AreDefaultScriptDialogsEnabled Getter
	SetAreDefaultScriptDialogsEnabled(value bool) // property AreDefaultScriptDialogsEnabled Setter
	// AreDevToolsEnabled
	//  `AreDevToolsEnabled` controls whether the user is able to use the context
	//  menu or keyboard shortcuts to open the DevTools window.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredevtoolsenabled">See the ICoreWebView2Settings article.</see>
	AreDevToolsEnabled() bool         // property AreDevToolsEnabled Getter
	SetAreDevToolsEnabled(value bool) // property AreDevToolsEnabled Setter
	// IsScriptEnabled
	//  Controls if running JavaScript is enabled in all future navigations in
	//  the WebView. This only affects scripts in the document. Scripts
	//  injected with `ExecuteScript` runs even if script is disabled.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isscriptenabled">See the ICoreWebView2Settings article.</see>
	IsScriptEnabled() bool         // property IsScriptEnabled Getter
	SetIsScriptEnabled(value bool) // property IsScriptEnabled Setter
	// IsStatusBarEnabled
	//  `IsStatusBarEnabled` controls whether the status bar is displayed. The
	//  status bar is usually displayed in the lower left of the WebView and
	//  shows things such as the URI of a link when the user hovers over it and
	//  other information. The default value is `TRUE`. The status bar UI can be
	//  altered by web content and should not be considered secure.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isstatusbarenabled">See the ICoreWebView2Settings article.</see>
	IsStatusBarEnabled() bool         // property IsStatusBarEnabled Getter
	SetIsStatusBarEnabled(value bool) // property IsStatusBarEnabled Setter
	// IsWebMessageEnabled
	//  The `IsWebMessageEnabled` property is used when loading a new HTML
	//  document. If set to `TRUE`, communication from the host to the top-level
	//  HTML document of the WebView is allowed using `PostWebMessageAsJson`,
	//  `PostWebMessageAsString`, and message event of `window.chrome.webview`.
	//  For more information, navigate to PostWebMessageAsJson. Communication
	//  from the top-level HTML document of the WebView to the host is allowed
	//  using the postMessage function of `window.chrome.webview` and
	//  `add_WebMessageReceived` method. For more information, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  If set to false, then communication is disallowed. `PostWebMessageAsJson`
	//  and `PostWebMessageAsString` fails with `E_ACCESSDENIED` and
	//  `window.chrome.webview.postMessage` fails by throwing an instance of an
	//  `Error` object. The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iswebmessageenabled">See the ICoreWebView2Settings article.</see>
	IsWebMessageEnabled() bool         // property IsWebMessageEnabled Getter
	SetIsWebMessageEnabled(value bool) // property IsWebMessageEnabled Setter
	// IsZoomControlEnabled
	//  The `IsZoomControlEnabled` property is used to prevent the user from
	//  impacting the zoom of the WebView. When disabled, the user is not able
	//  to zoom using Ctrl++, Ctrl+-, or Ctrl+mouse wheel, but the zoom
	//  is set using `ZoomFactor` API. The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iszoomcontrolenabled">See the ICoreWebView2Settings article.</see>
	IsZoomControlEnabled() bool         // property IsZoomControlEnabled Getter
	SetIsZoomControlEnabled(value bool) // property IsZoomControlEnabled Setter
	// AreHostObjectsAllowed
	//  The `AreHostObjectsAllowed` property is used to control whether host
	//  objects are accessible from the page in WebView.
	//  The default value is `TRUE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_arehostobjectsallowed">See the ICoreWebView2Settings article.</see>
	AreHostObjectsAllowed() bool         // property AreHostObjectsAllowed Getter
	SetAreHostObjectsAllowed(value bool) // property AreHostObjectsAllowed Setter
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
	//  access features specific to a web browser, including but not limited to:
	//  - Ctrl-F and F3 for Find on Page
	//  - Ctrl-P for Print
	//  - Ctrl-R and F5 for Reload
	//  - Ctrl-Plus and Ctrl-Minus for zooming
	//  - Ctrl-Shift-C and F12 for DevTools
	//  - Special keys for browser functions, such as Back, Forward, and Search
	//
	//  It does not disable accelerator keys related to movement and text editing,
	//  such as:
	//  - Home, End, Page Up, and Page Down
	//  - Ctrl-X, Ctrl-C, Ctrl-V
	//  - Ctrl-A for Select All
	//  - Ctrl-Z for Undo
	//
	//  Those accelerator keys will always be enabled unless they are handled in
	//  the `AcceleratorKeyPressed` event.
	//
	//  This setting has no effect on the `AcceleratorKeyPressed` event. The event
	//  will be fired for all accelerator keys, whether they are enabled or not.
	//
	//  The default value for `AreBrowserAcceleratorKeysEnabled` is TRUE.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings3#get_arebrowseracceleratorkeysenabled">See the ICoreWebView2Settings3 article.</see>
	AreBrowserAcceleratorKeysEnabled() bool         // property AreBrowserAcceleratorKeysEnabled Getter
	SetAreBrowserAcceleratorKeysEnabled(value bool) // property AreBrowserAcceleratorKeysEnabled Setter
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
	//  It will take effect immediately after setting.
	//  The default value is `FALSE`.
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
	// IsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields.
	//  It will take effect immediately after setting.
	//  The default value is `TRUE`.
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
	// IsPinchZoomEnabled
	//  Pinch-zoom, referred to as "Page Scale" zoom, is performed as a post-rendering step,
	//  it changes the page scale factor property and scales the surface the web page is
	//  rendered onto when user performs a pinch zooming action. It does not change the layout
	//  but rather changes the viewport and clips the web content, the content outside of the
	//  viewport isn't visible onscreen and users can't reach this content using mouse.
	//
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
	//
	//  When this property is `TRUE`, then all configured navigation gestures are enabled:
	//  1. Swiping left and right to navigate forward and backward is always configured.
	//  2. Swiping down to refresh is off by default and not exposed via our API currently,
	//  it requires the "--pull-to-refresh" option to be included in the additional browser
	//  arguments to be configured. (See put_AdditionalBrowserArguments.)
	//
	//  When set to `FALSE`, the end user cannot swipe to navigate or pull to refresh.
	//  This API only affects the overscrolling navigation functionality and has no
	//  effect on the scrolling interaction used to explore the web content shown
	//  in WebView2.
	//
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
	//
	//  When this property is `TRUE`, then all the non-client region features
	//  will be enabled:
	//  Draggable Regions will be enabled, they are regions on a webpage that
	//  are marked with the CSS attribute `app-region: drag/no-drag`. When set to
	//  `drag`, these regions will be treated like the window's title bar, supporting
	//  dragging of the entire WebView and its host app window; the system menu shows
	//  upon right click, and a double click will trigger maximizing/restoration of the
	//  window size.
	//
	//  When set to `FALSE`, all non-client region support will be disabled.
	//  The `app-region` CSS style will be ignored on web pages.
	//  \snippet SettingsComponent.cpp ToggleNonClientRegionSupportEnabled
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings9#get_isnonclientregionsupportenabled">See the ICoreWebView2Settings9 article.</see>
	IsNonClientRegionSupportEnabled() bool         // property IsNonClientRegionSupportEnabled Getter
	SetIsNonClientRegionSupportEnabled(value bool) // property IsNonClientRegionSupportEnabled Setter
}

type TCoreWebView2Settings struct {
	TObject
}

func (m *TCoreWebView2Settings) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) BaseIntf() (result ICoreWebView2Settings) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SettingsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Settings(resultPtr)
	return
}

func (m *TCoreWebView2Settings) IsBuiltInErrorPageEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsBuiltInErrorPageEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) AreDefaultContextMenusEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetAreDefaultContextMenusEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) AreDefaultScriptDialogsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetAreDefaultScriptDialogsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) AreDevToolsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetAreDevToolsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsScriptEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsScriptEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsStatusBarEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsStatusBarEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsWebMessageEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsWebMessageEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsZoomControlEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsZoomControlEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) AreHostObjectsAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetAreHostObjectsAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) UserAgent() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SettingsAPI().SysCallN(12, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Settings) SetUserAgent(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2Settings) AreBrowserAcceleratorKeysEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetAreBrowserAcceleratorKeysEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsPasswordAutosaveEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsPasswordAutosaveEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsGeneralAutofillEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsGeneralAutofillEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsPinchZoomEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsPinchZoomEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsSwipeNavigationEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsSwipeNavigationEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) HiddenPdfToolbarItems() wvTypes.TWVPDFToolbarItems {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2SettingsAPI().SysCallN(18, 0, m.Instance())
	return wvTypes.TWVPDFToolbarItems(r)
}

func (m *TCoreWebView2Settings) SetHiddenPdfToolbarItems(value wvTypes.TWVPDFToolbarItems) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Settings) IsReputationCheckingRequired() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsReputationCheckingRequired(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Settings) IsNonClientRegionSupportEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SettingsAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Settings) SetIsNonClientRegionSupportEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SettingsAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2Settings class constructor
func NewCoreWebView2Settings(baseIntf ICoreWebView2Settings) ICoreWebView2Settings {
	r := coreWebView2SettingsAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Settings(r)
}

var (
	coreWebView2SettingsOnce   base.Once
	coreWebView2SettingsImport *imports.Imports = nil
)

func coreWebView2SettingsAPI() *imports.Imports {
	coreWebView2SettingsOnce.Do(func() {
		coreWebView2SettingsImport = api.NewDefaultImports()
		coreWebView2SettingsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Settings_Create", 0), // constructor NewCoreWebView2Settings
			/* 1 */ imports.NewTable("TCoreWebView2Settings_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2Settings_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2Settings_IsBuiltInErrorPageEnabled", 0), // property IsBuiltInErrorPageEnabled
			/* 4 */ imports.NewTable("TCoreWebView2Settings_AreDefaultContextMenusEnabled", 0), // property AreDefaultContextMenusEnabled
			/* 5 */ imports.NewTable("TCoreWebView2Settings_AreDefaultScriptDialogsEnabled", 0), // property AreDefaultScriptDialogsEnabled
			/* 6 */ imports.NewTable("TCoreWebView2Settings_AreDevToolsEnabled", 0), // property AreDevToolsEnabled
			/* 7 */ imports.NewTable("TCoreWebView2Settings_IsScriptEnabled", 0), // property IsScriptEnabled
			/* 8 */ imports.NewTable("TCoreWebView2Settings_IsStatusBarEnabled", 0), // property IsStatusBarEnabled
			/* 9 */ imports.NewTable("TCoreWebView2Settings_IsWebMessageEnabled", 0), // property IsWebMessageEnabled
			/* 10 */ imports.NewTable("TCoreWebView2Settings_IsZoomControlEnabled", 0), // property IsZoomControlEnabled
			/* 11 */ imports.NewTable("TCoreWebView2Settings_AreHostObjectsAllowed", 0), // property AreHostObjectsAllowed
			/* 12 */ imports.NewTable("TCoreWebView2Settings_UserAgent", 0), // property UserAgent
			/* 13 */ imports.NewTable("TCoreWebView2Settings_AreBrowserAcceleratorKeysEnabled", 0), // property AreBrowserAcceleratorKeysEnabled
			/* 14 */ imports.NewTable("TCoreWebView2Settings_IsPasswordAutosaveEnabled", 0), // property IsPasswordAutosaveEnabled
			/* 15 */ imports.NewTable("TCoreWebView2Settings_IsGeneralAutofillEnabled", 0), // property IsGeneralAutofillEnabled
			/* 16 */ imports.NewTable("TCoreWebView2Settings_IsPinchZoomEnabled", 0), // property IsPinchZoomEnabled
			/* 17 */ imports.NewTable("TCoreWebView2Settings_IsSwipeNavigationEnabled", 0), // property IsSwipeNavigationEnabled
			/* 18 */ imports.NewTable("TCoreWebView2Settings_HiddenPdfToolbarItems", 0), // property HiddenPdfToolbarItems
			/* 19 */ imports.NewTable("TCoreWebView2Settings_IsReputationCheckingRequired", 0), // property IsReputationCheckingRequired
			/* 20 */ imports.NewTable("TCoreWebView2Settings_IsNonClientRegionSupportEnabled", 0), // property IsNonClientRegionSupportEnabled
		}
	})
	return coreWebView2SettingsImport
}
