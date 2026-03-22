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
	"github.com/energye/lcl/types"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2EnvironmentOptions2 Parent: IObject
type ICoreWebView2EnvironmentOptions2 interface {
	IObject

	// GetExclusiveUserDataFolderAccess
	//  Whether other processes can create WebView2 from WebView2Environment created with the
	//  same user data folder and therefore sharing the same WebView browser process instance.
	//  Default is FALSE.
	GetExclusiveUserDataFolderAccess(outValue *int32) types.HRESULT // function
	// SetExclusiveUserDataFolderAccess
	//  Sets the `ExclusiveUserDataFolderAccess` property.
	//  The `ExclusiveUserDataFolderAccess` property specifies that the WebView environment
	//  obtains exclusive access to the user data folder.
	//  If the user data folder is already being used by another WebView environment with a
	//  different value for `ExclusiveUserDataFolderAccess` property, the creation of a WebView2Controller
	//  using the environment object will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  When set as TRUE, no other WebView can be created from other processes using WebView2Environment
	//  objects with the same UserDataFolder. This prevents other processes from creating WebViews
	//  which share the same browser process instance, since sharing is performed among
	//  WebViews that have the same UserDataFolder. When another process tries to create a
	//  WebView2Controller from an WebView2Environment object created with the same user data folder,
	//  it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	SetExclusiveUserDataFolderAccess(value int32) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions3 Parent: IObject
type ICoreWebView2EnvironmentOptions3 interface {
	IObject

	// GetIsCustomCrashReportingEnabled
	//  When `IsCustomCrashReportingEnabled` is set to `TRUE`, Windows won't send crash data to Microsoft endpoint.
	//  `IsCustomCrashReportingEnabled` is default to be `FALSE`, in this case, WebView will respect OS consent.
	GetIsCustomCrashReportingEnabled(outValue *int32) types.HRESULT // function
	// SetIsCustomCrashReportingEnabled
	//  Sets the `IsCustomCrashReportingEnabled` property.
	SetIsCustomCrashReportingEnabled(value int32) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions4 Parent: IObject
type ICoreWebView2EnvironmentOptions4 interface {
	IObject

	// GetCustomSchemeRegistrations
	//  Array of custom scheme registrations. The returned
	//  ICoreWebView2CustomSchemeRegistration pointers must be released, and the
	//  array itself must be deallocated with CoTaskMemFree.
	//  * out schemeRegistrations: PPUserType3 --> out schemeRegistrations: PPCoreWebView2CustomSchemeRegistration ************** WEBVIEW4DELPHI ************** *
	GetCustomSchemeRegistrations(outCount *uint32, outSchemeRegistrations *PPCoreWebView2CustomSchemeRegistration) types.HRESULT // function
	// SetCustomSchemeRegistrations
	//  Set the array of custom scheme registrations to be used.
	//  \snippet AppWindow.cpp CoreWebView2CustomSchemeRegistration
	//  * var schemeRegistrations: ICoreWebView2CustomSchemeRegistration --> schemeRegistrations: PPCoreWebView2CustomSchemeRegistration ************** WEBVIEW4DELPHI ************** *
	SetCustomSchemeRegistrations(count uint32, schemeRegistrations PPCoreWebView2CustomSchemeRegistration) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions5 Parent: IObject
type ICoreWebView2EnvironmentOptions5 interface {
	IObject

	// GetEnableTrackingPrevention
	//  The `EnableTrackingPrevention` property is used to enable/disable tracking prevention
	//  feature in WebView2. This property enable/disable tracking prevention for all the
	//  WebView2's created in the same environment. By default this feature is enabled to block
	//  potentially harmful trackers and trackers from sites that aren't visited before and set to
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED` or whatever value was last changed/persisted
	//  on the profile.
	//
	//  You can set this property to false to disable the tracking prevention feature if the app only
	//  renders content in the WebView2 that is known to be safe. Disabling this feature when creating
	//  environment also improves runtime performance by skipping related code.
	//
	//  You shouldn't disable this property if WebView2 is being used as a "full browser" with arbitrary
	//  navigation and should protect end user privacy.
	//
	//  There is `ICoreWebView2Profile3::PreferredTrackingPreventionLevel` property to control levels of
	//  tracking prevention of the WebView2's associated with a same profile. However, you can also disable
	//  tracking prevention later using `ICoreWebView2Profile3::PreferredTrackingPreventionLevel` property and
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE` value but that doesn't improves runtime performance.
	//
	//  See `ICoreWebView2Profile3::PreferredTrackingPreventionLevel` for more details.
	//
	//  Tracking prevention protects users from online tracking by restricting the ability of trackers to
	//  access browser-based storage as well as the network. See [Tracking prevention](/microsoft-edge/web-platform/tracking-prevention).
	GetEnableTrackingPrevention(outValue *int32) types.HRESULT // function
	// SetEnableTrackingPrevention
	//  Sets the `EnableTrackingPrevention` property.
	SetEnableTrackingPrevention(value int32) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions6 Parent: IObject
type ICoreWebView2EnvironmentOptions6 interface {
	IObject

	// GetAreBrowserExtensionsEnabled
	//  When `AreBrowserExtensionsEnabled` is set to `TRUE`, new extensions can be added to user
	//  profile and used. `AreBrowserExtensionsEnabled` is default to be `FALSE`, in this case,
	//  new extensions can't be installed, and already installed extension won't be
	//  available to use in user profile.
	//  If connecting to an already running environment with a different value for `AreBrowserExtensionsEnabled`
	//  property, it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  See `ICoreWebView2BrowserExtension` for Extensions API details.
	GetAreBrowserExtensionsEnabled(outValue *int32) types.HRESULT // function
	// SetAreBrowserExtensionsEnabled
	//  Sets the `AreBrowserExtensionsEnabled` property.
	SetAreBrowserExtensionsEnabled(value int32) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions7 Parent: IObject
type ICoreWebView2EnvironmentOptions7 interface {
	IObject

	// GetChannelSearchKind
	//  Gets the `ChannelSearchKind` property.
	GetChannelSearchKind(outValue *wvTypes.COREWEBVIEW2_CHANNEL_SEARCH_KIND) types.HRESULT // function
	// SetChannelSearchKind
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
	SetChannelSearchKind(value wvTypes.COREWEBVIEW2_CHANNEL_SEARCH_KIND) types.HRESULT // function
	// GetReleaseChannels
	//  Gets the `ReleaseChannels` property.
	GetReleaseChannels(outValue *wvTypes.COREWEBVIEW2_RELEASE_CHANNELS) types.HRESULT // function
	// SetReleaseChannels
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
	SetReleaseChannels(value wvTypes.COREWEBVIEW2_RELEASE_CHANNELS) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions8 Parent: IObject
type ICoreWebView2EnvironmentOptions8 interface {
	IObject

	// GetScrollBarStyle
	//  Gets the `ScrollBarStyle` property.
	GetScrollBarStyle(outValue *wvTypes.COREWEBVIEW2_SCROLLBAR_STYLE) types.HRESULT // function
	// SetScrollBarStyle
	//  The ScrollBar style being set on the WebView2 Environment.
	//  The default value is `COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT`
	//  which specifies the default browser ScrollBar style.
	//  The `color-scheme` CSS property needs to be set on the corresponding page
	//  to allow ScrollBar to follow light or dark theme. Please see
	//  [color-scheme](https://developer.mozilla.org/docs/Web/CSS/color-scheme#declaring_color_scheme_preferences)
	//  for how `color-scheme` can be set.
	//  CSS styles that modify the ScrollBar applied on top of native ScrollBar styling
	//  that is selected with `ScrollBarStyle`.
	SetScrollBarStyle(value wvTypes.COREWEBVIEW2_SCROLLBAR_STYLE) types.HRESULT // function
}

// ICoreWebView2EnvironmentOptions Parent: ICoreWebView2EnvironmentOptions2 ICoreWebView2EnvironmentOptions3 ICoreWebView2EnvironmentOptions4 ICoreWebView2EnvironmentOptions5 ICoreWebView2EnvironmentOptions6 ICoreWebView2EnvironmentOptions7 ICoreWebView2EnvironmentOptions8 IInterfacedObject
type ICoreWebView2EnvironmentOptions interface {
	ICoreWebView2EnvironmentOptions2
	ICoreWebView2EnvironmentOptions3
	ICoreWebView2EnvironmentOptions4
	ICoreWebView2EnvironmentOptions5
	ICoreWebView2EnvironmentOptions6
	ICoreWebView2EnvironmentOptions7
	ICoreWebView2EnvironmentOptions8
	IInterfacedObject
	// GetAdditionalBrowserArguments
	//  Changes the behavior of the WebView. The arguments are passed to the
	//  browser process as part of the command. For more information about
	//  using command-line switches with Chromium browser processes, navigate to
	//  [Run Chromium with Flags](https://www.chromium.org/developers/how-tos/run-chromium-with-flags).
	//  The value appended to a switch is appended to the browser process, for
	//  example, in `--edge-webview-switches=xxx` the value is `xxx`. If you
	//  specify a switch that is important to WebView functionality, it is
	//  ignored, for example, `--user-data-dir`. Specific features are disabled
	//  internally and blocked from being enabled. If a switch is specified
	//  multiple times, only the last instance is used.
	//
	//  \> [!NOTE]\n\> A merge of the different values of the same switch is not attempted,
	//  except for disabled and enabled features. The features specified by
	//  `--enable-features` and `--disable-features` are merged with simple
	//  logic.\n\> * The features is the union of the specified features
	//  and built-in features. If a feature is disabled, it is removed from the
	//  enabled features list.
	//
	//  If you specify command-line switches and use the
	//  `additionalBrowserArguments` parameter, the `--edge-webview-switches`
	//  value takes precedence and is processed last. If a switch fails to
	//  parse, the switch is ignored. The default state for the operation is
	//  to run the browser process with no extra flags.
	//
	//  The caller must free the returned string with `CoTaskMemFree`. See
	//  [API Conventions](/microsoft-edge/webview2/concepts/win32-api-conventions#strings).
	GetAdditionalBrowserArguments(outValue *string) types.HRESULT // function
	// SetAdditionalBrowserArguments
	//  Sets the `AdditionalBrowserArguments` property.
	//
	//  Please note that calling this API twice will replace the previous value
	//  rather than appending to it. If there are multiple switches, there
	//  should be a space in between them. The one exception is if multiple
	//  features are being enabled/disabled for a single switch, in which
	//  case the features should be comma-seperated.
	//  Ex. "--disable-features=feature1,feature2 --some-other-switch --do-something"
	SetAdditionalBrowserArguments(value string) types.HRESULT // function
	// GetLanguage
	//  The default display language for WebView. It applies to browser UI such as
	//  context menu and dialogs. It also applies to the `accept-languages` HTTP
	//  header that WebView sends to websites. The intended locale value is in the
	//  format of BCP 47 Language Tags. More information can be found from
	//  [IETF BCP47](https://www.ietf.org/rfc/bcp/bcp47.html).
	//
	//  The caller must free the returned string with `CoTaskMemFree`. See
	//  [API Conventions](/microsoft-edge/webview2/concepts/win32-api-conventions#strings).
	GetLanguage(outValue *string) types.HRESULT // function
	// SetLanguage
	//  Sets the `Language` property.
	SetLanguage(value string) types.HRESULT // function
	// GetTargetCompatibleBrowserVersion
	//  Specifies the version of the WebView2 Runtime binaries required to be
	//  compatible with your app. This defaults to the WebView2 Runtime version
	//  that corresponds with the version of the SDK the app is using. The
	//  format of this value is the same as the format of the
	//  `BrowserVersionString` property and other `BrowserVersion` values. Only
	//  the version part of the `BrowserVersion` value is respected. The channel
	//  suffix, if it exists, is ignored. The version of the WebView2 Runtime
	//  binaries actually used may be different from the specified
	//  `TargetCompatibleBrowserVersion`. The binaries are only guaranteed to be
	//  compatible. Verify the actual version on the `BrowserVersionString`
	//  property on the `ICoreWebView2Environment`.
	//
	//  The caller must free the returned string with `CoTaskMemFree`. See
	//  [API Conventions](/microsoft-edge/webview2/concepts/win32-api-conventions#strings).
	GetTargetCompatibleBrowserVersion(outValue *string) types.HRESULT // function
	// SetTargetCompatibleBrowserVersion
	//  Sets the `TargetCompatibleBrowserVersion` property.
	SetTargetCompatibleBrowserVersion(value string) types.HRESULT // function
	// GetAllowSingleSignOnUsingOSPrimaryAccount
	//  The `AllowSingleSignOnUsingOSPrimaryAccount` property is used to enable
	//  single sign on with Azure Active Directory (AAD) and personal Microsoft
	//  Account (MSA) resources inside WebView. All AAD accounts, connected to
	//  Windows and shared for all apps, are supported. For MSA, SSO is only enabled
	//  for the account associated for Windows account login, if any.
	//  Default is disabled. Universal Windows Platform apps must also declare
	//  `enterpriseCloudSSO`
	//  [Restricted capabilities](/windows/uwp/packaging/app-capability-declarations\#restricted-capabilities)
	//  for the single sign on (SSO) to work.
	GetAllowSingleSignOnUsingOSPrimaryAccount(outAllow *int32) types.HRESULT // function
	// SetAllowSingleSignOnUsingOSPrimaryAccount
	//  Sets the `AllowSingleSignOnUsingOSPrimaryAccount` property.
	SetAllowSingleSignOnUsingOSPrimaryAccount(allow int32) types.HRESULT // function
	AsIntfEnvironmentOptions() uintptr
	AsIntfEnvironmentOptions2() uintptr
	AsIntfEnvironmentOptions3() uintptr
	AsIntfEnvironmentOptions4() uintptr
	AsIntfEnvironmentOptions5() uintptr
	AsIntfEnvironmentOptions6() uintptr
	AsIntfEnvironmentOptions7() uintptr
	AsIntfEnvironmentOptions8() uintptr
}

type TCoreWebView2EnvironmentOptions struct {
	TInterfacedObject
}

func (m *TCoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments(outValue *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = api.GoStr(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetAdditionalBrowserArguments(value string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(2, m.Instance(), api.PasStr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetLanguage(outValue *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = api.GoStr(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetLanguage(value string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(4, m.Instance(), api.PasStr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetTargetCompatibleBrowserVersion(outValue *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = api.GoStr(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetTargetCompatibleBrowserVersion(value string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(6, m.Instance(), api.PasStr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetAllowSingleSignOnUsingOSPrimaryAccount(outAllow *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var allowPtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&allowPtr)))
	*outAllow = int32(allowPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetAllowSingleSignOnUsingOSPrimaryAccount(allow int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(8, m.Instance(), uintptr(allow))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetExclusiveUserDataFolderAccess(outValue *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = int32(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetExclusiveUserDataFolderAccess(value int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(10, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetIsCustomCrashReportingEnabled(outValue *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = int32(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetIsCustomCrashReportingEnabled(value int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(12, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetCustomSchemeRegistrations(outCount *uint32, outSchemeRegistrations *PPCoreWebView2CustomSchemeRegistration) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var countPtr uintptr
	var schemeRegistrationsPtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&countPtr)), uintptr(base.UnsafePointer(&schemeRegistrationsPtr)))
	*outCount = uint32(countPtr)
	*outSchemeRegistrations = PPCoreWebView2CustomSchemeRegistration(schemeRegistrationsPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetCustomSchemeRegistrations(count uint32, schemeRegistrations PPCoreWebView2CustomSchemeRegistration) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(14, m.Instance(), uintptr(count), uintptr(schemeRegistrations))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetEnableTrackingPrevention(outValue *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = int32(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetEnableTrackingPrevention(value int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(16, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetAreBrowserExtensionsEnabled(outValue *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = int32(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetAreBrowserExtensionsEnabled(value int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(18, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetChannelSearchKind(outValue *wvTypes.COREWEBVIEW2_CHANNEL_SEARCH_KIND) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = wvTypes.COREWEBVIEW2_CHANNEL_SEARCH_KIND(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetChannelSearchKind(value wvTypes.COREWEBVIEW2_CHANNEL_SEARCH_KIND) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(20, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetReleaseChannels(outValue *wvTypes.COREWEBVIEW2_RELEASE_CHANNELS) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = wvTypes.COREWEBVIEW2_RELEASE_CHANNELS(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetReleaseChannels(value wvTypes.COREWEBVIEW2_RELEASE_CHANNELS) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(22, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) GetScrollBarStyle(outValue *wvTypes.COREWEBVIEW2_SCROLLBAR_STYLE) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var valuePtr uintptr
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&valuePtr)))
	*outValue = wvTypes.COREWEBVIEW2_SCROLLBAR_STYLE(valuePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) SetScrollBarStyle(value wvTypes.COREWEBVIEW2_SCROLLBAR_STYLE) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(24, m.Instance(), uintptr(value))
	return types.HRESULT(r)
}

func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions() uintptr {
	return m.GetIntfPointer(0)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions2() uintptr {
	return m.GetIntfPointer(1)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions3() uintptr {
	return m.GetIntfPointer(2)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions4() uintptr {
	return m.GetIntfPointer(3)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions5() uintptr {
	return m.GetIntfPointer(4)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions6() uintptr {
	return m.GetIntfPointer(5)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions7() uintptr {
	return m.GetIntfPointer(6)
}
func (m *TCoreWebView2EnvironmentOptions) AsIntfEnvironmentOptions8() uintptr {
	return m.GetIntfPointer(7)
}

// NewCoreWebView2EnvironmentOptions class constructor
func NewCoreWebView2EnvironmentOptions(additionalBrowserArguments string, language string, targetCompatibleBrowserVersion string, allowSingleSignOnUsingOSPrimaryAccount bool, exclusiveUserDataFolderAccess bool, customCrashReportingEnabled bool, schemeRegistrations IWVCustomSchemeRegistrationArray, enableTrackingPrevention bool, areBrowserExtensionsEnabled bool, channelSearchKind wvTypes.TWVChannelSearchKind, releaseChannels wvTypes.TWVReleaseChannels, scrollBarStyle wvTypes.TWVScrollBarStyle) ICoreWebView2EnvironmentOptions {
	var environmentOptionsPtr uintptr  // ICoreWebView2EnvironmentOptions
	var environmentOptions2Ptr uintptr // ICoreWebView2EnvironmentOptions2
	var environmentOptions3Ptr uintptr // ICoreWebView2EnvironmentOptions3
	var environmentOptions4Ptr uintptr // ICoreWebView2EnvironmentOptions4
	var environmentOptions5Ptr uintptr // ICoreWebView2EnvironmentOptions5
	var environmentOptions6Ptr uintptr // ICoreWebView2EnvironmentOptions6
	var environmentOptions7Ptr uintptr // ICoreWebView2EnvironmentOptions7
	var environmentOptions8Ptr uintptr // ICoreWebView2EnvironmentOptions8
	r := coreWebView2EnvironmentOptionsAPI().SysCallN(0, api.PasStr(additionalBrowserArguments), api.PasStr(language), api.PasStr(targetCompatibleBrowserVersion), api.PasBool(allowSingleSignOnUsingOSPrimaryAccount), api.PasBool(exclusiveUserDataFolderAccess), api.PasBool(customCrashReportingEnabled), schemeRegistrations.Instance(), uintptr(int32(schemeRegistrations.Count())), api.PasBool(enableTrackingPrevention), api.PasBool(areBrowserExtensionsEnabled), uintptr(channelSearchKind), uintptr(releaseChannels), uintptr(scrollBarStyle), uintptr(base.UnsafePointer(&environmentOptionsPtr)), uintptr(base.UnsafePointer(&environmentOptions2Ptr)), uintptr(base.UnsafePointer(&environmentOptions3Ptr)), uintptr(base.UnsafePointer(&environmentOptions4Ptr)), uintptr(base.UnsafePointer(&environmentOptions5Ptr)), uintptr(base.UnsafePointer(&environmentOptions6Ptr)), uintptr(base.UnsafePointer(&environmentOptions7Ptr)), uintptr(base.UnsafePointer(&environmentOptions8Ptr)))
	ret := AsCoreWebView2EnvironmentOptions(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(8)
		intf.SetIntfPointer(0, environmentOptionsPtr)
		intf.SetIntfPointer(1, environmentOptions2Ptr)
		intf.SetIntfPointer(2, environmentOptions3Ptr)
		intf.SetIntfPointer(3, environmentOptions4Ptr)
		intf.SetIntfPointer(4, environmentOptions5Ptr)
		intf.SetIntfPointer(5, environmentOptions6Ptr)
		intf.SetIntfPointer(6, environmentOptions7Ptr)
		intf.SetIntfPointer(7, environmentOptions8Ptr)
	}
	return ret
}

var (
	coreWebView2EnvironmentOptionsOnce   base.Once
	coreWebView2EnvironmentOptionsImport *imports.Imports = nil
)

func coreWebView2EnvironmentOptionsAPI() *imports.Imports {
	coreWebView2EnvironmentOptionsOnce.Do(func() {
		coreWebView2EnvironmentOptionsImport = api.NewDefaultImports()
		coreWebView2EnvironmentOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Create", 0), // constructor NewCoreWebView2EnvironmentOptions
			/* 1 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_AdditionalBrowserArguments", 0), // function GetAdditionalBrowserArguments
			/* 2 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_AdditionalBrowserArguments", 0), // function SetAdditionalBrowserArguments
			/* 3 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_Language", 0), // function GetLanguage
			/* 4 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_Language", 0), // function SetLanguage
			/* 5 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_TargetCompatibleBrowserVersion", 0), // function GetTargetCompatibleBrowserVersion
			/* 6 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_TargetCompatibleBrowserVersion", 0), // function SetTargetCompatibleBrowserVersion
			/* 7 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_AllowSingleSignOnUsingOSPrimaryAccount", 0), // function GetAllowSingleSignOnUsingOSPrimaryAccount
			/* 8 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_AllowSingleSignOnUsingOSPrimaryAccount", 0), // function SetAllowSingleSignOnUsingOSPrimaryAccount
			/* 9 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_ExclusiveUserDataFolderAccess", 0), // function GetExclusiveUserDataFolderAccess
			/* 10 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_ExclusiveUserDataFolderAccess", 0), // function SetExclusiveUserDataFolderAccess
			/* 11 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_IsCustomCrashReportingEnabled", 0), // function GetIsCustomCrashReportingEnabled
			/* 12 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_IsCustomCrashReportingEnabled", 0), // function SetIsCustomCrashReportingEnabled
			/* 13 */ imports.NewTable("TCoreWebView2EnvironmentOptions_GetCustomSchemeRegistrations", 0), // function GetCustomSchemeRegistrations
			/* 14 */ imports.NewTable("TCoreWebView2EnvironmentOptions_SetCustomSchemeRegistrations", 0), // function SetCustomSchemeRegistrations
			/* 15 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_EnableTrackingPrevention", 0), // function GetEnableTrackingPrevention
			/* 16 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_EnableTrackingPrevention", 0), // function SetEnableTrackingPrevention
			/* 17 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_AreBrowserExtensionsEnabled", 0), // function GetAreBrowserExtensionsEnabled
			/* 18 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_AreBrowserExtensionsEnabled", 0), // function SetAreBrowserExtensionsEnabled
			/* 19 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_ChannelSearchKind", 0), // function GetChannelSearchKind
			/* 20 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_ChannelSearchKind", 0), // function SetChannelSearchKind
			/* 21 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_ReleaseChannels", 0), // function GetReleaseChannels
			/* 22 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_ReleaseChannels", 0), // function SetReleaseChannels
			/* 23 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Get_ScrollBarStyle", 0), // function GetScrollBarStyle
			/* 24 */ imports.NewTable("TCoreWebView2EnvironmentOptions_Set_ScrollBarStyle", 0), // function SetScrollBarStyle
		}
	})
	return coreWebView2EnvironmentOptionsImport
}
