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

	wvTypes "github.com/energye/wv/types/windows"
)

// IWVLoaderEvents Parent: IObject
type IWVLoaderEvents interface {
	IObject
}

// IWVLoader Parent: IWVLoaderEvents IComponent
type IWVLoader interface {
	IWVLoaderEvents
	IComponent
	// StartWebView2
	//  This function is used to initialize WebView2.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</see>
	StartWebView2() bool // function
	// CompareVersions
	//  This method is for anyone want to compare version correctly to determine
	//  which version is newer, older or same.Use it to determine whether
	//  to use webview2 or certain feature based upon version. Sets the value of
	//  aCompRslt to -1, 0 or 1 if aVersion1 is less than, equal or greater
	//  than aVersion2 respectively. Returns false if it fails to parse
	//  any of the version strings.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#comparebrowserversions">See the Globals article.</see>
	CompareVersions(version1 string, version2 string, compRslt *int32) bool // function
	// UpdateDeviceScaleFactor
	//  Update the DeviceScaleFactor property value with the current scale or the ForcedDeviceScaleFactor value.
	UpdateDeviceScaleFactor() // procedure
	// AppendErrorLog
	//  Append aText to the ErrorMessage property.
	AppendErrorLog(text string) // procedure
	// Environment
	//  Represents the global WebView2 Environment.
	Environment() ICoreWebView2Environment // property Environment Getter
	// Status
	//  Returns the TWVLoader initialization status.
	Status() wvTypes.TWV2LoaderStatus // property Status Getter
	// AvailableBrowserVersion
	//  Get the browser version info including channel name if it is not the
	//  WebView2 Runtime. Channel names are Beta, Dev, and Canary.
	AvailableBrowserVersion() string // property AvailableBrowserVersion Getter
	// AvailableBrowserVersionWithOptions
	//  Get the browser version info of the release channel used when creating
	//  an environment with the same options. Channel names are Beta, Dev, and
	//  Canary.
	AvailableBrowserVersionWithOptions() string // property AvailableBrowserVersionWithOptions Getter
	// ErrorMessage
	//  Returns all the text appended to the error log with AppendErrorLog.
	ErrorMessage() string // property ErrorMessage Getter
	// ErrorCode
	//  Returns the last initialization error code.
	ErrorCode() int64 // property ErrorCode Getter
	// SetCurrentDir
	//  Used to set the current directory when the WebView2 library is loaded. This is required if the application is launched from a different application.
	SetCurrentDir() bool         // property SetCurrentDir Getter
	SetSetCurrentDir(value bool) // property SetCurrentDir Setter
	// Initialized
	//  Returns true if the Status is wvlsInitialized.
	Initialized() bool // property Initialized Getter
	// InitializationError
	//  Returns true if the Status is wvlsError.
	InitializationError() bool // property InitializationError Getter
	// CheckFiles
	//  Checks if the WebView2 library is present and the DLL version.
	CheckFiles() bool         // property CheckFiles Getter
	SetCheckFiles(value bool) // property CheckFiles Setter
	// ShowMessageDlg
	//  Set to true when you need to use a showmessage dialog to show the error messages.
	ShowMessageDlg() bool         // property ShowMessageDlg Getter
	SetShowMessageDlg(value bool) // property ShowMessageDlg Setter
	// InitCOMLibrary
	//  Set to true to call CoInitializeEx and CoUnInitialize in TWVLoader.Create and TWVLoader.Destroy.
	InitCOMLibrary() bool         // property InitCOMLibrary Getter
	SetInitCOMLibrary(value bool) // property InitCOMLibrary Setter
	// CustomCommandLineSwitches
	//  Custom command line switches used by TCoreWebView2EnvironmentOptions.Create to initialize WebView2.
	CustomCommandLineSwitches() string // property CustomCommandLineSwitches Getter
	// DeviceScaleFactor
	//  Returns the device scale factor.
	DeviceScaleFactor() float32 // property DeviceScaleFactor Getter
	// ReRaiseExceptions
	//  Set to true to raise all exceptions.
	ReRaiseExceptions() bool         // property ReRaiseExceptions Getter
	SetReRaiseExceptions(value bool) // property ReRaiseExceptions Setter
	// InstalledRuntimeVersion
	//  Returns the installed WebView2 runtime version.
	InstalledRuntimeVersion() string // property InstalledRuntimeVersion Getter
	// LoaderDllPath
	//  Full path to WebView2Loader.dll. Leave empty to load WebView2Loader.dll from the current directory.
	LoaderDllPath() string         // property LoaderDllPath Getter
	SetLoaderDllPath(value string) // property LoaderDllPath Setter
	// UseInternalLoader
	//  Use a WebView2Loader.dll replacement based on the OpenWebView2Loader project.
	//  <see href="https://github.com/jchv/OpenWebView2Loader">See the OpenWebView2Loader project repository at GitHub.</see>
	UseInternalLoader() bool         // property UseInternalLoader Getter
	SetUseInternalLoader(value bool) // property UseInternalLoader Setter
	// AllowOldRuntime
	//  Allow using old WebView2 Runtime versions.
	AllowOldRuntime() bool         // property AllowOldRuntime Getter
	SetAllowOldRuntime(value bool) // property AllowOldRuntime Setter
	// BrowserExecPath
	//  Use BrowserExecPath to specify whether WebView2 controls use a fixed or
	//  installed version of the WebView2 Runtime that exists on a user machine.
	//  To use a fixed version of the WebView2 Runtime, pass the folder path that
	//  contains the fixed version of the WebView2 Runtime to BrowserExecPath.
	//  BrowserExecPath supports both relative (to the application's executable)
	//  and absolute files paths. To create WebView2 controls that use the installed
	//  version of the WebView2 Runtime that exists on user machines,
	//  pass an empty string to BrowserExecPath.
	//  Property used to create the environment. Used as the browserExecutableFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</see>
	BrowserExecPath() string         // property BrowserExecPath Getter
	SetBrowserExecPath(value string) // property BrowserExecPath Setter
	// UserDataFolder
	//  You may specify the userDataFolder to change the default user data folder
	//  location for WebView2. The path is either an absolute file path or a relative
	//  file path that is interpreted as relative to the compiled code for the
	//  current process.
	//  Property used to create the environment. Used as the userDataFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</see>
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
	// EnableGPU
	//  Enable GPU hardware acceleration.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu-compositing</see>
	EnableGPU() bool         // property EnableGPU Getter
	SetEnableGPU(value bool) // property EnableGPU Setter
	// EnableFeatures
	//  List of feature names to enable.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-features</see>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/concepts/webview-features-flags">See the WebView2 browser flags article.</see>
	//  The list of features you can enable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	EnableFeatures() string         // property EnableFeatures Getter
	SetEnableFeatures(value string) // property EnableFeatures Setter
	// DisableFeatures
	//  List of feature names to disable.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-features</see>
	//  The list of features you can disable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	DisableFeatures() string         // property DisableFeatures Getter
	SetDisableFeatures(value string) // property DisableFeatures Setter
	// EnableBlinkFeatures
	//  Enable one or more Blink runtime-enabled features.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-blink-features</see>
	//  The list of Blink features you can enable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	EnableBlinkFeatures() string         // property EnableBlinkFeatures Getter
	SetEnableBlinkFeatures(value string) // property EnableBlinkFeatures Setter
	// DisableBlinkFeatures
	//  Disable one or more Blink runtime-enabled features.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-blink-features</see>
	//  The list of Blink features you can disable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	DisableBlinkFeatures() string         // property DisableBlinkFeatures Getter
	SetDisableBlinkFeatures(value string) // property DisableBlinkFeatures Setter
	// BlinkSettings
	//  Set blink settings. Format is <name>[=<value],<name>[=<value>],...
	//  The names are declared in Settings.json5. For boolean type, use "true", "false",
	//  or omit '=<value>' part to set to true. For enum type, use the int value of the
	//  enum value. Applied after other command line flags and prefs.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --blink-settings</see>
	//  The list of Blink settings you can disable is here:
	//  https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/settings.json5
	BlinkSettings() string         // property BlinkSettings Getter
	SetBlinkSettings(value string) // property BlinkSettings Setter
	// ForceFieldTrials
	//  This option can be used to force field trials when testing changes locally.
	//  The argument is a list of name and value pairs, separated by slashes.
	//  If a trial name is prefixed with an asterisk, that trial will start activated.
	//  For example, the following argument defines two trials, with the second one
	//  activated: "GoogleNow/Enable/*MaterialDesignNTP/Default/" This option can also
	//  be used by the browser process to send the list of trials to a non-browser
	//  process, using the same format. See FieldTrialList::CreateTrialsFromString()
	//  in field_trial.h for details.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrials</see>
	//  https://source.chromium.org/chromium/chromium/src/+/master:base/base_switches.cc
	ForceFieldTrials() string         // property ForceFieldTrials Getter
	SetForceFieldTrials(value string) // property ForceFieldTrials Setter
	// ForceFieldTrialParams
	//  This option can be used to force parameters of field trials when testing
	//  changes locally. The argument is a param list of (key, value) pairs prefixed
	//  by an associated (trial, group) pair. You specify the param list for multiple
	//  (trial, group) pairs with a comma separator.
	//  Example: "Trial1.Group1:k1/v1/k2/v2,Trial2.Group2:k3/v3/k4/v4"
	//  Trial names, groups names, parameter names, and value should all be URL
	//  escaped for all non-alphanumeric characters.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrial-params</see>
	//  https://source.chromium.org/chromium/chromium/src/+/master:components/variations/variations_switches.cc
	ForceFieldTrialParams() string         // property ForceFieldTrialParams Getter
	SetForceFieldTrialParams(value string) // property ForceFieldTrialParams Setter
	// SmartScreenProtectionEnabled
	//  Workaround given my Microsoft to disable the SmartScreen protection.
	SmartScreenProtectionEnabled() bool         // property SmartScreenProtectionEnabled Getter
	SetSmartScreenProtectionEnabled(value bool) // property SmartScreenProtectionEnabled Setter
	// AllowInsecureLocalhost
	//  Enables TLS/SSL errors on localhost to be ignored (no interstitial, no blocking of requests).
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-insecure-localhost</see>
	AllowInsecureLocalhost() bool         // property AllowInsecureLocalhost Getter
	SetAllowInsecureLocalhost(value bool) // property AllowInsecureLocalhost Setter
	// DisableWebSecurity
	//  Don't enforce the same-origin policy.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-web-security</see>
	DisableWebSecurity() bool         // property DisableWebSecurity Getter
	SetDisableWebSecurity(value bool) // property DisableWebSecurity Setter
	// TouchEvents
	//  Enable support for touch event feature detection.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --touch-events</see>
	TouchEvents() wvTypes.TWVState         // property TouchEvents Getter
	SetTouchEvents(value wvTypes.TWVState) // property TouchEvents Setter
	// HyperlinkAuditing
	//  Don't send hyperlink auditing pings.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-pings</see>
	HyperlinkAuditing() bool         // property HyperlinkAuditing Getter
	SetHyperlinkAuditing(value bool) // property HyperlinkAuditing Setter
	// AutoplayPolicy
	//  Autoplay policy.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --autoplay-policy</see>
	AutoplayPolicy() wvTypes.TWVAutoplayPolicy         // property AutoplayPolicy Getter
	SetAutoplayPolicy(value wvTypes.TWVAutoplayPolicy) // property AutoplayPolicy Setter
	// MuteAudio
	//  Mutes audio sent to the audio device so it is not audible during automated testing.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --mute-audio</see>
	MuteAudio() bool         // property MuteAudio Getter
	SetMuteAudio(value bool) // property MuteAudio Setter
	// DefaultEncoding
	//  Default encoding.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --default-encoding</see>
	DefaultEncoding() string         // property DefaultEncoding Getter
	SetDefaultEncoding(value string) // property DefaultEncoding Setter
	// KioskPrinting
	//  Enable automatically pressing the print button in print preview.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --kiosk-printing</see>
	KioskPrinting() bool         // property KioskPrinting Getter
	SetKioskPrinting(value bool) // property KioskPrinting Setter
	// ProxySettings
	//  Configure the browser to use a proxy server.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</see>
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</see>
	ProxySettings() IWVProxySettings // property ProxySettings Getter
	// AllowFileAccessFromFiles
	//  By default, file:// URIs cannot read other file:// URIs. This is an override for developers who need the old behavior for testing.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-file-access-from-files</see>
	AllowFileAccessFromFiles() bool         // property AllowFileAccessFromFiles Getter
	SetAllowFileAccessFromFiles(value bool) // property AllowFileAccessFromFiles Setter
	// AllowRunningInsecureContent
	//  By default, an https page cannot run JavaScript, CSS or plugins from http URLs. This provides an override to get the old insecure behavior.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-running-insecure-content</see>
	AllowRunningInsecureContent() bool         // property AllowRunningInsecureContent Getter
	SetAllowRunningInsecureContent(value bool) // property AllowRunningInsecureContent Setter
	// DisableBackgroundNetworking
	//  Disable several subsystems which run network requests in the background.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-background-networking</see>
	DisableBackgroundNetworking() bool         // property DisableBackgroundNetworking Getter
	SetDisableBackgroundNetworking(value bool) // property DisableBackgroundNetworking Setter
	// ForcedDeviceScaleFactor
	//  Overrides the device scale factor for the browser UI and the contents.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-device-scale-factor</see>
	ForcedDeviceScaleFactor() float32         // property ForcedDeviceScaleFactor Getter
	SetForcedDeviceScaleFactor(value float32) // property ForcedDeviceScaleFactor Setter
	// RemoteDebuggingPort
	//  Set to a value between 1024 and 65535 to enable remote debugging on the
	//  specified port. Also configurable using the "remote-debugging-port"
	//  command-line switch. Remote debugging can be accessed by loading the
	//  chrome://inspect page in Google Chrome. Port numbers 9222 and 9229 are
	//  discoverable by default. Other port numbers may need to be configured via
	//  "Discover network targets" on the Devices tab.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-debugging-port</see>
	RemoteDebuggingPort() int32         // property RemoteDebuggingPort Getter
	SetRemoteDebuggingPort(value int32) // property RemoteDebuggingPort Setter
	// RemoteAllowOrigins
	//  Enables web socket connections from the specified origins only. '*' allows any origin.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-allow-origins</see>
	RemoteAllowOrigins() string         // property RemoteAllowOrigins Getter
	SetRemoteAllowOrigins(value string) // property RemoteAllowOrigins Setter
	// DebugLog
	//  Force logging to be enabled. Logging is disabled by default in release builds.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-logging</see>
	DebugLog() wvTypes.TWV2DebugLog         // property DebugLog Getter
	SetDebugLog(value wvTypes.TWV2DebugLog) // property DebugLog Setter
	// DebugLogLevel
	//  Sets the minimum log level.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-level</see>
	DebugLogLevel() wvTypes.TWV2DebugLogLevel         // property DebugLogLevel Getter
	SetDebugLogLevel(value wvTypes.TWV2DebugLogLevel) // property DebugLogLevel Setter
	// JavaScriptFlags
	//  Specifies the flags passed to JS engine.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --js-flags</see>
	JavaScriptFlags() string         // property JavaScriptFlags Getter
	SetJavaScriptFlags(value string) // property JavaScriptFlags Setter
	// DisableEdgePitchNotification
	//  Workaround given my Microsoft to disable the "Download Edge" notifications.
	DisableEdgePitchNotification() bool         // property DisableEdgePitchNotification Getter
	SetDisableEdgePitchNotification(value bool) // property DisableEdgePitchNotification Setter
	// TreatInsecureOriginAsSecure
	//  Treat given (insecure) origins as secure origins.
	//  Multiple origins can be supplied as a comma-separated list.
	//  For the definition of secure contexts, see https://w3c.github.io/webappsec-secure-contexts/
	//  and https://www.w3.org/TR/powerful-features/#is-origin-trustworthy
	//  Example: --unsafely-treat-insecure-origin-as-secure=http://a.test,http://b.test
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --unsafely-treat-insecure-origin-as-secure</see>
	TreatInsecureOriginAsSecure() string         // property TreatInsecureOriginAsSecure Getter
	SetTreatInsecureOriginAsSecure(value string) // property TreatInsecureOriginAsSecure Setter
	// OpenOfficeDocumentsInWebViewer
	//  Enable the MS Office file viewer in the browser.
	//  This is a workaround given by Microsoft to open MS Office documents in the web browser instead of downloading the files.
	OpenOfficeDocumentsInWebViewer() bool         // property OpenOfficeDocumentsInWebViewer Getter
	SetOpenOfficeDocumentsInWebViewer(value bool) // property OpenOfficeDocumentsInWebViewer Setter
	// MicrosoftSignIn
	//  If enabled, allows implicit sign-in to Microsoft webpages using any account, by using the information from the primary OS account.
	//  This property uses the msSingleSignOnOSForPrimaryAccountIsShared flag.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/concepts/webview-features-flags">See the WebView2 browser flags article.</see>
	MicrosoftSignIn() bool         // property MicrosoftSignIn Getter
	SetMicrosoftSignIn(value bool) // property MicrosoftSignIn Setter
	// TLS13HybridizedKyberSupport
	//  This option enables a combination of X25519 and Kyber in TLS 1.3.
	TLS13HybridizedKyberSupport() wvTypes.TWVState         // property TLS13HybridizedKyberSupport Getter
	SetTLS13HybridizedKyberSupport(value wvTypes.TWVState) // property TLS13HybridizedKyberSupport Setter
	// UserAgent
	//  A string used to override the default user agent with a custom one.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --user-agent</see>
	UserAgent() string         // property UserAgent Getter
	SetUserAgent(value string) // property UserAgent Setter
	// AutoAcceptCamAndMicCapture
	//  Bypasses the dialog prompting the user for permission to capture cameras and microphones.
	//  Useful in automatic tests of video-conferencing Web applications. This is nearly
	//  identical to kUseFakeUIForMediaStream, with the exception being that this flag does NOT
	//  affect screen-capture.
	//  <see href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --auto-accept-camera-and-microphone-capture</see>
	AutoAcceptCamAndMicCapture() bool         // property AutoAcceptCamAndMicCapture Getter
	SetAutoAcceptCamAndMicCapture(value bool) // property AutoAcceptCamAndMicCapture Setter
	// SupportsCompositionController
	//  Returns true if the current WebView2 runtime version supports Composition Controllers.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3">See the ICoreWebView2Environment3 article.</see>
	SupportsCompositionController() bool // property SupportsCompositionController Getter
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`
	//  Provide a list of all process using same user data folder except for crashpad process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8">See the ICoreWebView2Environment8 article.</see>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property ProcessInfos Getter
	// SupportsControllerOptions
	//  Returns true if the current WebView2 runtime version supports Controller Options.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10">See the ICoreWebView2Environment10 article.</see>
	SupportsControllerOptions() bool // property SupportsControllerOptions Getter
	// FailureReportFolderPath
	//  `FailureReportFolderPath` returns the path of the folder where minidump files are written.
	//  Whenever a WebView2 process crashes, a crash dump file will be created in the crash dump folder.
	//  The crash dump format is minidump files.
	//  Please see [Minidump Files documentation](/windows/win32/debug/minidump-files) for detailed information.
	//  Normally when a single child process fails, a minidump will be generated and written to disk,
	//  then the `ProcessFailed` event is raised. But for unexpected crashes, a minidump file might not be generated
	//  at all, despite whether `ProcessFailed` event is raised. If there are multiple
	//  process failures at once, multiple minidump files could be generated. Thus `FailureReportFolderPath`
	//  could contain old minidump files that are not associated with a specific `ProcessFailed` event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment11#get_failurereportfolderpath">See the ICoreWebView2Environment11 article.</see>
	FailureReportFolderPath() string                                           // property FailureReportFolderPath Getter
	SetOnEnvironmentCreated(fn TLoaderNotifyEvent)                             // property event
	SetOnInitializationError(fn TLoaderNotifyEvent)                            // property event
	SetOnGetCustomSchemes(fn TLoaderGetCustomSchemesEvent)                     // property event
	SetOnNewBrowserVersionAvailable(fn TLoaderNewBrowserVersionAvailableEvent) // property event
	SetOnBrowserProcessExited(fn TLoaderBrowserProcessExitedEvent)             // property event
	SetOnProcessInfosChanged(fn TLoaderProcessInfosChangedEvent)               // property event
	AsIntfWVLoaderEvents() uintptr
}

type TWVLoader struct {
	TComponent
}

func (m *TWVLoader) StartWebView2() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) CompareVersions(version1 string, version2 string, compRslt *int32) bool {
	if !m.IsValid() {
		return false
	}
	compRsltPtr := uintptr(*compRslt)
	r := wVLoaderAPI().SysCallN(2, m.Instance(), api.PasStr(version1), api.PasStr(version2), uintptr(base.UnsafePointer(&compRsltPtr)))
	*compRslt = int32(compRsltPtr)
	return api.GoBool(r)
}

func (m *TWVLoader) UpdateDeviceScaleFactor() {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(3, m.Instance())
}

func (m *TWVLoader) AppendErrorLog(text string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(4, m.Instance(), api.PasStr(text))
}

func (m *TWVLoader) Environment() (result ICoreWebView2Environment) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVLoaderAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Environment(resultPtr)
	return
}

func (m *TWVLoader) Status() wvTypes.TWV2LoaderStatus {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(6, m.Instance())
	return wvTypes.TWV2LoaderStatus(r)
}

func (m *TWVLoader) AvailableBrowserVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) AvailableBrowserVersionWithOptions() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) ErrorMessage() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) ErrorCode() (result int64) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVLoader) SetCurrentDir() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetSetCurrentDir(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) InitializationError() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) CheckFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetCheckFiles(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) ShowMessageDlg() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetShowMessageDlg(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) InitCOMLibrary() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetInitCOMLibrary(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) CustomCommandLineSwitches() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(17, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) DeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVLoader) ReRaiseExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetReRaiseExceptions(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) InstalledRuntimeVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(20, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) LoaderDllPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(21, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetLoaderDllPath(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(21, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) UseInternalLoader() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetUseInternalLoader(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) AllowOldRuntime() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAllowOldRuntime(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) BrowserExecPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(24, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetBrowserExecPath(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(24, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) UserDataFolder() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetUserDataFolder(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) AdditionalBrowserArguments() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetAdditionalBrowserArguments(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) Language() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(27, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetLanguage(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(27, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) TargetCompatibleBrowserVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetTargetCompatibleBrowserVersion(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) AllowSingleSignOnUsingOSPrimaryAccount() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAllowSingleSignOnUsingOSPrimaryAccount(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) ExclusiveUserDataFolderAccess() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(30, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetExclusiveUserDataFolderAccess(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(30, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) CustomCrashReportingEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(31, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetCustomCrashReportingEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(31, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) EnableTrackingPrevention() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetEnableTrackingPrevention(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) AreBrowserExtensionsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAreBrowserExtensionsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) ChannelSearchKind() wvTypes.TWVChannelSearchKind {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(34, 0, m.Instance())
	return wvTypes.TWVChannelSearchKind(r)
}

func (m *TWVLoader) SetChannelSearchKind(value wvTypes.TWVChannelSearchKind) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) ReleaseChannels() wvTypes.TWVReleaseChannels {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(35, 0, m.Instance())
	return wvTypes.TWVReleaseChannels(r)
}

func (m *TWVLoader) SetReleaseChannels(value wvTypes.TWVReleaseChannels) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) ScrollBarStyle() wvTypes.TWVScrollBarStyle {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(36, 0, m.Instance())
	return wvTypes.TWVScrollBarStyle(r)
}

func (m *TWVLoader) SetScrollBarStyle(value wvTypes.TWVScrollBarStyle) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) EnableGPU() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetEnableGPU(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) EnableFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(38, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetEnableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(38, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) DisableFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(39, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetDisableFeatures(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(39, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) EnableBlinkFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(40, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetEnableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(40, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) DisableBlinkFeatures() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(41, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetDisableBlinkFeatures(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(41, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) BlinkSettings() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(42, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetBlinkSettings(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(42, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) ForceFieldTrials() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(43, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetForceFieldTrials(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(43, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) ForceFieldTrialParams() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(44, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetForceFieldTrialParams(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(44, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) SmartScreenProtectionEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(45, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetSmartScreenProtectionEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(45, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) AllowInsecureLocalhost() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(46, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAllowInsecureLocalhost(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(46, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) DisableWebSecurity() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(47, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetDisableWebSecurity(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(47, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) TouchEvents() wvTypes.TWVState {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(48, 0, m.Instance())
	return wvTypes.TWVState(r)
}

func (m *TWVLoader) SetTouchEvents(value wvTypes.TWVState) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) HyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(49, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetHyperlinkAuditing(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(49, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) AutoplayPolicy() wvTypes.TWVAutoplayPolicy {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(50, 0, m.Instance())
	return wvTypes.TWVAutoplayPolicy(r)
}

func (m *TWVLoader) SetAutoplayPolicy(value wvTypes.TWVAutoplayPolicy) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) MuteAudio() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(51, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetMuteAudio(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(51, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) DefaultEncoding() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(52, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetDefaultEncoding(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(52, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) KioskPrinting() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(53, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetKioskPrinting(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(53, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) ProxySettings() IWVProxySettings {
	if !m.IsValid() {
		return nil
	}
	r := wVLoaderAPI().SysCallN(54, m.Instance())
	return AsWVProxySettings(r)
}

func (m *TWVLoader) AllowFileAccessFromFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAllowFileAccessFromFiles(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) AllowRunningInsecureContent() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(56, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAllowRunningInsecureContent(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(56, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) DisableBackgroundNetworking() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(57, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetDisableBackgroundNetworking(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(57, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) ForcedDeviceScaleFactor() (result float32) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(58, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWVLoader) SetForcedDeviceScaleFactor(value float32) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(58, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TWVLoader) RemoteDebuggingPort() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(59, 0, m.Instance())
	return int32(r)
}

func (m *TWVLoader) SetRemoteDebuggingPort(value int32) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) RemoteAllowOrigins() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(60, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetRemoteAllowOrigins(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(60, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) DebugLog() wvTypes.TWV2DebugLog {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(61, 0, m.Instance())
	return wvTypes.TWV2DebugLog(r)
}

func (m *TWVLoader) SetDebugLog(value wvTypes.TWV2DebugLog) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(61, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) DebugLogLevel() wvTypes.TWV2DebugLogLevel {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(62, 0, m.Instance())
	return wvTypes.TWV2DebugLogLevel(r)
}

func (m *TWVLoader) SetDebugLogLevel(value wvTypes.TWV2DebugLogLevel) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(62, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) JavaScriptFlags() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(63, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetJavaScriptFlags(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(63, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) DisableEdgePitchNotification() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetDisableEdgePitchNotification(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) TreatInsecureOriginAsSecure() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(65, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetTreatInsecureOriginAsSecure(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(65, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) OpenOfficeDocumentsInWebViewer() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetOpenOfficeDocumentsInWebViewer(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) MicrosoftSignIn() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(67, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetMicrosoftSignIn(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(67, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) TLS13HybridizedKyberSupport() wvTypes.TWVState {
	if !m.IsValid() {
		return 0
	}
	r := wVLoaderAPI().SysCallN(68, 0, m.Instance())
	return wvTypes.TWVState(r)
}

func (m *TWVLoader) SetTLS13HybridizedKyberSupport(value wvTypes.TWVState) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(68, 1, m.Instance(), uintptr(value))
}

func (m *TWVLoader) UserAgent() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(69, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetUserAgent(value string) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(69, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVLoader) AutoAcceptCamAndMicCapture() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(70, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) SetAutoAcceptCamAndMicCapture(value bool) {
	if !m.IsValid() {
		return
	}
	wVLoaderAPI().SysCallN(70, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVLoader) SupportsCompositionController() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(71, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) ProcessInfos() (result ICoreWebView2ProcessInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	wVLoaderAPI().SysCallN(72, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfoCollection(resultPtr)
	return
}

func (m *TWVLoader) SupportsControllerOptions() bool {
	if !m.IsValid() {
		return false
	}
	r := wVLoaderAPI().SysCallN(73, m.Instance())
	return api.GoBool(r)
}

func (m *TWVLoader) FailureReportFolderPath() string {
	if !m.IsValid() {
		return ""
	}
	r := wVLoaderAPI().SysCallN(74, m.Instance())
	return api.GoStr(r)
}

func (m *TWVLoader) SetOnEnvironmentCreated(fn TLoaderNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderNotifyEvent(fn)
	base.SetEvent(m, 75, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) SetOnInitializationError(fn TLoaderNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderNotifyEvent(fn)
	base.SetEvent(m, 76, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) SetOnGetCustomSchemes(fn TLoaderGetCustomSchemesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderGetCustomSchemesEvent(fn)
	base.SetEvent(m, 77, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) SetOnNewBrowserVersionAvailable(fn TLoaderNewBrowserVersionAvailableEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderNewBrowserVersionAvailableEvent(fn)
	base.SetEvent(m, 78, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) SetOnBrowserProcessExited(fn TLoaderBrowserProcessExitedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderBrowserProcessExitedEvent(fn)
	base.SetEvent(m, 79, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) SetOnProcessInfosChanged(fn TLoaderProcessInfosChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLoaderProcessInfosChangedEvent(fn)
	base.SetEvent(m, 80, wVLoaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVLoader) AsIntfWVLoaderEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewLoader class constructor
func NewLoader(owner lcl.IComponent) IWVLoader {
	var wVLoaderEventsPtr uintptr // IWVLoaderEvents
	r := wVLoaderAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&wVLoaderEventsPtr)))
	ret := AsWVLoader(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, wVLoaderEventsPtr)
	}
	return ret
}

var (
	wVLoaderOnce   base.Once
	wVLoaderImport *imports.Imports = nil
)

func wVLoaderAPI() *imports.Imports {
	wVLoaderOnce.Do(func() {
		wVLoaderImport = api.NewDefaultImports()
		wVLoaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVLoader_Create", 0), // constructor NewLoader
			/* 1 */ imports.NewTable("TWVLoader_StartWebView2", 0), // function StartWebView2
			/* 2 */ imports.NewTable("TWVLoader_CompareVersions", 0), // function CompareVersions
			/* 3 */ imports.NewTable("TWVLoader_UpdateDeviceScaleFactor", 0), // procedure UpdateDeviceScaleFactor
			/* 4 */ imports.NewTable("TWVLoader_AppendErrorLog", 0), // procedure AppendErrorLog
			/* 5 */ imports.NewTable("TWVLoader_Environment", 0), // property Environment
			/* 6 */ imports.NewTable("TWVLoader_Status", 0), // property Status
			/* 7 */ imports.NewTable("TWVLoader_AvailableBrowserVersion", 0), // property AvailableBrowserVersion
			/* 8 */ imports.NewTable("TWVLoader_AvailableBrowserVersionWithOptions", 0), // property AvailableBrowserVersionWithOptions
			/* 9 */ imports.NewTable("TWVLoader_ErrorMessage", 0), // property ErrorMessage
			/* 10 */ imports.NewTable("TWVLoader_ErrorCode", 0), // property ErrorCode
			/* 11 */ imports.NewTable("TWVLoader_SetCurrentDir", 0), // property SetCurrentDir
			/* 12 */ imports.NewTable("TWVLoader_Initialized", 0), // property Initialized
			/* 13 */ imports.NewTable("TWVLoader_InitializationError", 0), // property InitializationError
			/* 14 */ imports.NewTable("TWVLoader_CheckFiles", 0), // property CheckFiles
			/* 15 */ imports.NewTable("TWVLoader_ShowMessageDlg", 0), // property ShowMessageDlg
			/* 16 */ imports.NewTable("TWVLoader_InitCOMLibrary", 0), // property InitCOMLibrary
			/* 17 */ imports.NewTable("TWVLoader_CustomCommandLineSwitches", 0), // property CustomCommandLineSwitches
			/* 18 */ imports.NewTable("TWVLoader_DeviceScaleFactor", 0), // property DeviceScaleFactor
			/* 19 */ imports.NewTable("TWVLoader_ReRaiseExceptions", 0), // property ReRaiseExceptions
			/* 20 */ imports.NewTable("TWVLoader_InstalledRuntimeVersion", 0), // property InstalledRuntimeVersion
			/* 21 */ imports.NewTable("TWVLoader_LoaderDllPath", 0), // property LoaderDllPath
			/* 22 */ imports.NewTable("TWVLoader_UseInternalLoader", 0), // property UseInternalLoader
			/* 23 */ imports.NewTable("TWVLoader_AllowOldRuntime", 0), // property AllowOldRuntime
			/* 24 */ imports.NewTable("TWVLoader_BrowserExecPath", 0), // property BrowserExecPath
			/* 25 */ imports.NewTable("TWVLoader_UserDataFolder", 0), // property UserDataFolder
			/* 26 */ imports.NewTable("TWVLoader_AdditionalBrowserArguments", 0), // property AdditionalBrowserArguments
			/* 27 */ imports.NewTable("TWVLoader_Language", 0), // property Language
			/* 28 */ imports.NewTable("TWVLoader_TargetCompatibleBrowserVersion", 0), // property TargetCompatibleBrowserVersion
			/* 29 */ imports.NewTable("TWVLoader_AllowSingleSignOnUsingOSPrimaryAccount", 0), // property AllowSingleSignOnUsingOSPrimaryAccount
			/* 30 */ imports.NewTable("TWVLoader_ExclusiveUserDataFolderAccess", 0), // property ExclusiveUserDataFolderAccess
			/* 31 */ imports.NewTable("TWVLoader_CustomCrashReportingEnabled", 0), // property CustomCrashReportingEnabled
			/* 32 */ imports.NewTable("TWVLoader_EnableTrackingPrevention", 0), // property EnableTrackingPrevention
			/* 33 */ imports.NewTable("TWVLoader_AreBrowserExtensionsEnabled", 0), // property AreBrowserExtensionsEnabled
			/* 34 */ imports.NewTable("TWVLoader_ChannelSearchKind", 0), // property ChannelSearchKind
			/* 35 */ imports.NewTable("TWVLoader_ReleaseChannels", 0), // property ReleaseChannels
			/* 36 */ imports.NewTable("TWVLoader_ScrollBarStyle", 0), // property ScrollBarStyle
			/* 37 */ imports.NewTable("TWVLoader_EnableGPU", 0), // property EnableGPU
			/* 38 */ imports.NewTable("TWVLoader_EnableFeatures", 0), // property EnableFeatures
			/* 39 */ imports.NewTable("TWVLoader_DisableFeatures", 0), // property DisableFeatures
			/* 40 */ imports.NewTable("TWVLoader_EnableBlinkFeatures", 0), // property EnableBlinkFeatures
			/* 41 */ imports.NewTable("TWVLoader_DisableBlinkFeatures", 0), // property DisableBlinkFeatures
			/* 42 */ imports.NewTable("TWVLoader_BlinkSettings", 0), // property BlinkSettings
			/* 43 */ imports.NewTable("TWVLoader_ForceFieldTrials", 0), // property ForceFieldTrials
			/* 44 */ imports.NewTable("TWVLoader_ForceFieldTrialParams", 0), // property ForceFieldTrialParams
			/* 45 */ imports.NewTable("TWVLoader_SmartScreenProtectionEnabled", 0), // property SmartScreenProtectionEnabled
			/* 46 */ imports.NewTable("TWVLoader_AllowInsecureLocalhost", 0), // property AllowInsecureLocalhost
			/* 47 */ imports.NewTable("TWVLoader_DisableWebSecurity", 0), // property DisableWebSecurity
			/* 48 */ imports.NewTable("TWVLoader_TouchEvents", 0), // property TouchEvents
			/* 49 */ imports.NewTable("TWVLoader_HyperlinkAuditing", 0), // property HyperlinkAuditing
			/* 50 */ imports.NewTable("TWVLoader_AutoplayPolicy", 0), // property AutoplayPolicy
			/* 51 */ imports.NewTable("TWVLoader_MuteAudio", 0), // property MuteAudio
			/* 52 */ imports.NewTable("TWVLoader_DefaultEncoding", 0), // property DefaultEncoding
			/* 53 */ imports.NewTable("TWVLoader_KioskPrinting", 0), // property KioskPrinting
			/* 54 */ imports.NewTable("TWVLoader_ProxySettings", 0), // property ProxySettings
			/* 55 */ imports.NewTable("TWVLoader_AllowFileAccessFromFiles", 0), // property AllowFileAccessFromFiles
			/* 56 */ imports.NewTable("TWVLoader_AllowRunningInsecureContent", 0), // property AllowRunningInsecureContent
			/* 57 */ imports.NewTable("TWVLoader_DisableBackgroundNetworking", 0), // property DisableBackgroundNetworking
			/* 58 */ imports.NewTable("TWVLoader_ForcedDeviceScaleFactor", 0), // property ForcedDeviceScaleFactor
			/* 59 */ imports.NewTable("TWVLoader_RemoteDebuggingPort", 0), // property RemoteDebuggingPort
			/* 60 */ imports.NewTable("TWVLoader_RemoteAllowOrigins", 0), // property RemoteAllowOrigins
			/* 61 */ imports.NewTable("TWVLoader_DebugLog", 0), // property DebugLog
			/* 62 */ imports.NewTable("TWVLoader_DebugLogLevel", 0), // property DebugLogLevel
			/* 63 */ imports.NewTable("TWVLoader_JavaScriptFlags", 0), // property JavaScriptFlags
			/* 64 */ imports.NewTable("TWVLoader_DisableEdgePitchNotification", 0), // property DisableEdgePitchNotification
			/* 65 */ imports.NewTable("TWVLoader_TreatInsecureOriginAsSecure", 0), // property TreatInsecureOriginAsSecure
			/* 66 */ imports.NewTable("TWVLoader_OpenOfficeDocumentsInWebViewer", 0), // property OpenOfficeDocumentsInWebViewer
			/* 67 */ imports.NewTable("TWVLoader_MicrosoftSignIn", 0), // property MicrosoftSignIn
			/* 68 */ imports.NewTable("TWVLoader_TLS13HybridizedKyberSupport", 0), // property TLS13HybridizedKyberSupport
			/* 69 */ imports.NewTable("TWVLoader_UserAgent", 0), // property UserAgent
			/* 70 */ imports.NewTable("TWVLoader_AutoAcceptCamAndMicCapture", 0), // property AutoAcceptCamAndMicCapture
			/* 71 */ imports.NewTable("TWVLoader_SupportsCompositionController", 0), // property SupportsCompositionController
			/* 72 */ imports.NewTable("TWVLoader_ProcessInfos", 0), // property ProcessInfos
			/* 73 */ imports.NewTable("TWVLoader_SupportsControllerOptions", 0), // property SupportsControllerOptions
			/* 74 */ imports.NewTable("TWVLoader_FailureReportFolderPath", 0), // property FailureReportFolderPath
			/* 75 */ imports.NewTable("TWVLoader_OnEnvironmentCreated", 0), // event OnEnvironmentCreated
			/* 76 */ imports.NewTable("TWVLoader_OnInitializationError", 0), // event OnInitializationError
			/* 77 */ imports.NewTable("TWVLoader_OnGetCustomSchemes", 0), // event OnGetCustomSchemes
			/* 78 */ imports.NewTable("TWVLoader_OnNewBrowserVersionAvailable", 0), // event OnNewBrowserVersionAvailable
			/* 79 */ imports.NewTable("TWVLoader_OnBrowserProcessExited", 0), // event OnBrowserProcessExited
			/* 80 */ imports.NewTable("TWVLoader_OnProcessInfosChanged", 0), // event OnProcessInfosChanged
		}
	})
	return wVLoaderImport
}
