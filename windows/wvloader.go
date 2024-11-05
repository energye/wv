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

// IWVLoader Parent: IComponent
//
//	Class used to simplify the WebView2 initialization and destruction.
type IWVLoader interface {
	IComponent
	SetOnEnvironmentCreated(fn TOnLoaderNotifyEvent)
	SetOnInitializationError(fn TOnLoaderNotifyEvent)
	SetOnGetCustomSchemes(fn TOnLoaderGetCustomSchemesEvent)
	SetOnNewBrowserVersionAvailable(fn TOnLoaderNewBrowserVersionAvailableEvent)
	SetOnBrowserProcessExited(fn TOnLoaderBrowserProcessExitedEvent)
	SetOnProcessInfosChanged(fn TOnLoaderProcessInfosChangedEvent)
	// Environment
	//  Represents the global WebView2 Environment.
	Environment() ICoreWebView2Environment // property
	// Status
	//  Returns the TWVLoader initialization status.
	Status() TWV2LoaderStatus // property
	// AvailableBrowserVersion
	//  Get the browser version info including channel name if it is not the
	//  WebView2 Runtime. Channel names are Beta, Dev, and Canary.
	AvailableBrowserVersion() string // property
	// ErrorMessage
	//  Returns all the text appended to the error log with AppendErrorLog.
	ErrorMessage() string // property
	// ErrorCode
	//  Returns the last initialization error code.
	ErrorCode() (resultInt64 int64) // property
	// SetCurrentDir
	//  Used to set the current directory when the WebView2 library is loaded. This is required if the application is launched from a different application.
	SetCurrentDir() bool // property
	// SetSetCurrentDir Set SetCurrentDir
	SetSetCurrentDir(AValue bool) // property
	// Initialized
	//  Returns true if the Status is wvlsInitialized.
	Initialized() bool // property
	// InitializationError
	//  Returns true if the Status is wvlsError.
	InitializationError() bool // property
	// CheckFiles
	//  Checks if the WebView2 library is present and the DLL version.
	CheckFiles() bool // property
	// SetCheckFiles Set CheckFiles
	SetCheckFiles(AValue bool) // property
	// ShowMessageDlg
	//  Set to true when you need to use a showmessage dialog to show the error messages.
	ShowMessageDlg() bool // property
	// SetShowMessageDlg Set ShowMessageDlg
	SetShowMessageDlg(AValue bool) // property
	// InitCOMLibrary
	//  Set to true to call CoInitializeEx and CoUnInitialize in TWVLoader.Create and TWVLoader.Destroy.
	InitCOMLibrary() bool // property
	// SetInitCOMLibrary Set InitCOMLibrary
	SetInitCOMLibrary(AValue bool) // property
	// CustomCommandLineSwitches
	//  Custom command line switches used by TCoreWebView2EnvironmentOptions.Create to initialize WebView2.
	CustomCommandLineSwitches() string // property
	// DeviceScaleFactor
	//  Returns the device scale factor.
	DeviceScaleFactor() (resultSingle float32) // property
	// ReRaiseExceptions
	//  Set to true to raise all exceptions.
	ReRaiseExceptions() bool // property
	// SetReRaiseExceptions Set ReRaiseExceptions
	SetReRaiseExceptions(AValue bool) // property
	// InstalledRuntimeVersion
	//  Returns the installed WebView2 runtime version.
	InstalledRuntimeVersion() string // property
	// LoaderDllPath
	//  Full path to WebView2Loader.dll. Leave empty to load WebView2Loader.dll from the current directory.
	LoaderDllPath() string // property
	// SetLoaderDllPath Set LoaderDllPath
	SetLoaderDllPath(AValue string) // property
	// UseInternalLoader
	//  Use a WebView2Loader.dll replacement based on the OpenWebView2Loader project.
	//  <a href="https://github.com/jchv/OpenWebView2Loader">See the OpenWebView2Loader project repository at GitHub.</a>
	UseInternalLoader() bool // property
	// SetUseInternalLoader Set UseInternalLoader
	SetUseInternalLoader(AValue bool) // property
	// BrowserExecPath
	//  Use BrowserExecPath to specify whether WebView2 controls use a fixed or
	//  installed version of the WebView2 Runtime that exists on a user machine.
	//  To use a fixed version of the WebView2 Runtime, pass the folder path that
	//  contains the fixed version of the WebView2 Runtime to BrowserExecPath.
	//  BrowserExecPath supports both relative(to the application's executable)
	//  and absolute files paths. To create WebView2 controls that use the installed
	//  version of the WebView2 Runtime that exists on user machines,
	//  pass an empty string to BrowserExecPath.
	//  Property used to create the environment. Used as the browserExecutableFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	BrowserExecPath() string // property
	// SetBrowserExecPath Set BrowserExecPath
	SetBrowserExecPath(AValue string) // property
	// UserDataFolder
	//  You may specify the userDataFolder to change the default user data folder
	//  location for WebView2. The path is either an absolute file path or a relative
	//  file path that is interpreted as relative to the compiled code for the
	//  current process.
	//  Property used to create the environment. Used as the userDataFolder parameter of CreateCoreWebView2EnvironmentWithOptions.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	UserDataFolder() string // property
	// SetUserDataFolder Set UserDataFolder
	SetUserDataFolder(AValue string) // property
	// AdditionalBrowserArguments
	//  Additional command line switches.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AdditionalBrowserArguments.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AdditionalBrowserArguments() string // property
	// SetAdditionalBrowserArguments Set AdditionalBrowserArguments
	SetAdditionalBrowserArguments(AValue string) // property
	// Language
	//  The default display language for WebView. It applies to browser UI such as
	//  context menu and dialogs. It also applies to the `accept-languages` HTTP
	//  header that WebView sends to websites.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_Language.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	Language() string // property
	// SetLanguage Set Language
	SetLanguage(AValue string) // property
	// TargetCompatibleBrowserVersion
	//  Specifies the version of the WebView2 Runtime binaries required to be
	//  compatible with your app.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_TargetCompatibleBrowserVersion.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	TargetCompatibleBrowserVersion() string // property
	// SetTargetCompatibleBrowserVersion Set TargetCompatibleBrowserVersion
	SetTargetCompatibleBrowserVersion(AValue string) // property
	// AllowSingleSignOnUsingOSPrimaryAccount
	//  Used to enable single sign on with Azure Active Directory(AAD) and personal Microsoft
	//  Account(MSA) resources inside WebView.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AllowSingleSignOnUsingOSPrimaryAccount.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AllowSingleSignOnUsingOSPrimaryAccount() bool // property
	// SetAllowSingleSignOnUsingOSPrimaryAccount Set AllowSingleSignOnUsingOSPrimaryAccount
	SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) // property
	// ExclusiveUserDataFolderAccess
	//  Whether other processes can create WebView2 from WebView2Environment created with the
	//  same user data folder and therefore sharing the same WebView browser process instance.
	//  Default is FALSE.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions2.Get_ExclusiveUserDataFolderAccess.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2">See the ICoreWebView2EnvironmentOptions2 article.</a>
	ExclusiveUserDataFolderAccess() bool // property
	// SetExclusiveUserDataFolderAccess Set ExclusiveUserDataFolderAccess
	SetExclusiveUserDataFolderAccess(AValue bool) // property
	// CustomCrashReportingEnabled
	//  When `CustomCrashReportingEnabled` is set to `TRUE`, Windows won't send crash data to Microsoft endpoint.
	//  `CustomCrashReportingEnabled` is default to be `FALSE`, in this case, WebView will respect OS consent.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions3.Get_IsCustomCrashReportingEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3">See the ICoreWebView2EnvironmentOptions3 article.</a>
	CustomCrashReportingEnabled() bool // property
	// SetCustomCrashReportingEnabled Set CustomCrashReportingEnabled
	SetCustomCrashReportingEnabled(AValue bool) // property
	// EnableTrackingPrevention
	//  The `EnableTrackingPrevention` property is used to enable/disable tracking prevention
	//  feature in WebView2. This property enable/disable tracking prevention for all the
	//  WebView2's created in the same environment.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions5.Get_EnableTrackingPrevention.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5">See the ICoreWebView2EnvironmentOptions5 article.</a>
	EnableTrackingPrevention() bool // property
	// SetEnableTrackingPrevention Set EnableTrackingPrevention
	SetEnableTrackingPrevention(AValue bool) // property
	// AreBrowserExtensionsEnabled
	//  When `AreBrowserExtensionsEnabled` is set to `TRUE`, new extensions can be added to user
	//  profile and used. `AreBrowserExtensionsEnabled` is default to be `FALSE`, in this case,
	//  new extensions can't be installed, and already installed extension won't be
	//  available to use in user profile.
	//  If connecting to an already running environment with a different value for `AreBrowserExtensionsEnabled`
	//  property, it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions6.Get_AreBrowserExtensionsEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions6">See the ICoreWebView2EnvironmentOptions6 article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article for Extensions API details.</a>
	AreBrowserExtensionsEnabled() bool // property
	// SetAreBrowserExtensionsEnabled Set AreBrowserExtensionsEnabled
	SetAreBrowserExtensionsEnabled(AValue bool) // property
	// EnableGPU
	//  Enable GPU hardware acceleration.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-gpu-compositing</a>
	EnableGPU() bool // property
	// SetEnableGPU Set EnableGPU
	SetEnableGPU(AValue bool) // property
	// EnableFeatures
	//  List of feature names to enable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-features</a>
	//  The list of features you can enable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	EnableFeatures() string // property
	// SetEnableFeatures Set EnableFeatures
	SetEnableFeatures(AValue string) // property
	// DisableFeatures
	//  List of feature names to disable.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-features</a>
	//  The list of features you can disable is here:
	//  https://chromium.googlesource.com/chromium/src/+/master/chrome/common/chrome_features.cc
	//  https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/content_features.cc
	//  https://source.chromium.org/search?q=base::Feature
	DisableFeatures() string // property
	// SetDisableFeatures Set DisableFeatures
	SetDisableFeatures(AValue string) // property
	// EnableBlinkFeatures
	//  Enable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-blink-features</a>
	//  The list of Blink features you can enable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	EnableBlinkFeatures() string // property
	// SetEnableBlinkFeatures Set EnableBlinkFeatures
	SetEnableBlinkFeatures(AValue string) // property
	// DisableBlinkFeatures
	//  Disable one or more Blink runtime-enabled features.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-blink-features</a>
	//  The list of Blink features you can disable is here:
	//  https://cs.chromium.org/chromium/src/third_party/blink/renderer/platform/runtime_enabled_features.json5
	DisableBlinkFeatures() string // property
	// SetDisableBlinkFeatures Set DisableBlinkFeatures
	SetDisableBlinkFeatures(AValue string) // property
	// BlinkSettings
	//  Set blink settings. Format is <name>[=<value],<name>[=<value>],...
	//  The names are declared in Settings.json5. For boolean type, use "true", "false",
	//  or omit '=<value>' part to set to true. For enum type, use the int value of the
	//  enum value. Applied after other command line flags and prefs.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --blink-settings</a>
	//  The list of Blink settings you can disable is here:
	//  https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/settings.json5
	BlinkSettings() string // property
	// SetBlinkSettings Set BlinkSettings
	SetBlinkSettings(AValue string) // property
	// ForceFieldTrials
	//  This option can be used to force field trials when testing changes locally.
	//  The argument is a list of name and value pairs, separated by slashes.
	//  If a trial name is prefixed with an asterisk, that trial will start activated.
	//  For example, the following argument defines two trials, with the second one
	//  activated: "GoogleNow/Enable/*MaterialDesignNTP/Default/" This option can also
	//  be used by the browser process to send the list of trials to a non-browser
	//  process, using the same format. See FieldTrialList::CreateTrialsFromString()
	//  in field_trial.h for details.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrials</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:base/base_switches.cc
	ForceFieldTrials() string // property
	// SetForceFieldTrials Set ForceFieldTrials
	SetForceFieldTrials(AValue string) // property
	// ForceFieldTrialParams
	//  This option can be used to force parameters of field trials when testing
	//  changes locally. The argument is a param list of(key, value) pairs prefixed
	//  by an associated(trial, group) pair. You specify the param list for multiple
	// (trial, group) pairs with a comma separator.
	//  Example: "Trial1.Group1:k1/v1/k2/v2,Trial2.Group2:k3/v3/k4/v4"
	//  Trial names, groups names, parameter names, and value should all be URL
	//  escaped for all non-alphanumeric characters.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-fieldtrial-params</a>
	//  https://source.chromium.org/chromium/chromium/src/+/master:components/variations/variations_switches.cc
	ForceFieldTrialParams() string // property
	// SetForceFieldTrialParams Set ForceFieldTrialParams
	SetForceFieldTrialParams(AValue string) // property
	// SmartScreenProtectionEnabled
	//  Workaround given my Microsoft to disable the SmartScreen protection.
	SmartScreenProtectionEnabled() bool // property
	// SetSmartScreenProtectionEnabled Set SmartScreenProtectionEnabled
	SetSmartScreenProtectionEnabled(AValue bool) // property
	// AllowInsecureLocalhost
	//  Enables TLS/SSL errors on localhost to be ignored(no interstitial, no blocking of requests).
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-insecure-localhost</a>
	AllowInsecureLocalhost() bool // property
	// SetAllowInsecureLocalhost Set AllowInsecureLocalhost
	SetAllowInsecureLocalhost(AValue bool) // property
	// DisableWebSecurity
	//  Don't enforce the same-origin policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-web-security</a>
	DisableWebSecurity() bool // property
	// SetDisableWebSecurity Set DisableWebSecurity
	SetDisableWebSecurity(AValue bool) // property
	// TouchEvents
	//  Enable support for touch event feature detection.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --touch-events</a>
	TouchEvents() TWVState // property
	// SetTouchEvents Set TouchEvents
	SetTouchEvents(AValue TWVState) // property
	// HyperlinkAuditing
	//  Don't send hyperlink auditing pings.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-pings</a>
	HyperlinkAuditing() bool // property
	// SetHyperlinkAuditing Set HyperlinkAuditing
	SetHyperlinkAuditing(AValue bool) // property
	// AutoplayPolicy
	//  Autoplay policy.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --autoplay-policy</a>
	AutoplayPolicy() TWVAutoplayPolicy // property
	// SetAutoplayPolicy Set AutoplayPolicy
	SetAutoplayPolicy(AValue TWVAutoplayPolicy) // property
	// MuteAudio
	//  Mutes audio sent to the audio device so it is not audible during automated testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --mute-audio</a>
	MuteAudio() bool // property
	// SetMuteAudio Set MuteAudio
	SetMuteAudio(AValue bool) // property
	// KioskPrinting
	//  Default encoding.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/libcef/common/cef_switches.cc">Uses the following command line switch: --default-encoding</a>
	//  Enable automatically pressing the print button in print preview.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --kiosk-printing</a>
	KioskPrinting() bool // property
	// SetKioskPrinting Set KioskPrinting
	SetKioskPrinting(AValue bool) // property
	// ProxySettings
	//  Configure the browser to use a proxy server.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
	ProxySettings() IWVProxySettings // property
	// AllowFileAccessFromFiles
	//  By default, file:// URIs cannot read other file:// URIs. This is an override for developers who need the old behavior for testing.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-file-access-from-files</a>
	AllowFileAccessFromFiles() bool // property
	// SetAllowFileAccessFromFiles Set AllowFileAccessFromFiles
	SetAllowFileAccessFromFiles(AValue bool) // property
	// AllowRunningInsecureContent
	//  By default, an https page cannot run JavaScript, CSS or plugins from http URLs. This provides an override to get the old insecure behavior.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --allow-running-insecure-content</a>
	AllowRunningInsecureContent() bool // property
	// SetAllowRunningInsecureContent Set AllowRunningInsecureContent
	SetAllowRunningInsecureContent(AValue bool) // property
	// DisableBackgroundNetworking
	//  Disable several subsystems which run network requests in the background.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --disable-background-networking</a>
	DisableBackgroundNetworking() bool // property
	// SetDisableBackgroundNetworking Set DisableBackgroundNetworking
	SetDisableBackgroundNetworking(AValue bool) // property
	// ForcedDeviceScaleFactor
	//  Overrides the device scale factor for the browser UI and the contents.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --force-device-scale-factor</a>
	ForcedDeviceScaleFactor() (resultSingle float32) // property
	// SetForcedDeviceScaleFactor Set ForcedDeviceScaleFactor
	SetForcedDeviceScaleFactor(AValue float32) // property
	// RemoteDebuggingPort
	//  Set to a value between 1024 and 65535 to enable remote debugging on the
	//  specified port. Also configurable using the "remote-debugging-port"
	//  command-line switch. Remote debugging can be accessed by loading the
	//  chrome://inspect page in Google Chrome. Port numbers 9222 and 9229 are
	//  discoverable by default. Other port numbers may need to be configured via
	//  "Discover network targets" on the Devices tab.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-debugging-port</a>
	RemoteDebuggingPort() int32 // property
	// SetRemoteDebuggingPort Set RemoteDebuggingPort
	SetRemoteDebuggingPort(AValue int32) // property
	// RemoteAllowOrigins
	//  Enables web socket connections from the specified origins only. '*' allows any origin.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --remote-allow-origins</a>
	RemoteAllowOrigins() string // property
	// SetRemoteAllowOrigins Set RemoteAllowOrigins
	SetRemoteAllowOrigins(AValue string) // property
	// DebugLog
	//  Force logging to be enabled. Logging is disabled by default in release builds.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --enable-logging</a>
	DebugLog() TWV2DebugLog // property
	// SetDebugLog Set DebugLog
	SetDebugLog(AValue TWV2DebugLog) // property
	// DebugLogLevel
	//  Sets the minimum log level.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --log-level</a>
	DebugLogLevel() TWV2DebugLogLevel // property
	// SetDebugLogLevel Set DebugLogLevel
	SetDebugLogLevel(AValue TWV2DebugLogLevel) // property
	// JavaScriptFlags
	//  Specifies the flags passed to JS engine.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --js-flags</a>
	JavaScriptFlags() string // property
	// SetJavaScriptFlags Set JavaScriptFlags
	SetJavaScriptFlags(AValue string) // property
	// DisableEdgePitchNotification
	//  Workaround given my Microsoft to disable the "Download Edge" notifications.
	DisableEdgePitchNotification() bool // property
	// SetDisableEdgePitchNotification Set DisableEdgePitchNotification
	SetDisableEdgePitchNotification(AValue bool) // property
	// TreatInsecureOriginAsSecure
	//  Treat given(insecure) origins as secure origins.
	//  Multiple origins can be supplied as a comma-separated list.
	//  For the definition of secure contexts, see https://w3c.github.io/webappsec-secure-contexts/
	//  and https://www.w3.org/TR/powerful-features/#is-origin-trustworthy
	//  Example: --unsafely-treat-insecure-origin-as-secure=http://a.test,http://b.test
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --unsafely-treat-insecure-origin-as-secure</a>
	TreatInsecureOriginAsSecure() string // property
	// SetTreatInsecureOriginAsSecure Set TreatInsecureOriginAsSecure
	SetTreatInsecureOriginAsSecure(AValue string) // property
	// AutoAcceptCamAndMicCapture
	//  Bypasses the dialog prompting the user for permission to capture cameras and microphones.
	//  Useful in automatic tests of video-conferencing Web applications. This is nearly
	//  identical to kUseFakeUIForMediaStream, with the exception being that this flag does NOT
	//  affect screen-capture.
	//  <a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --auto-accept-camera-and-microphone-capture</a>
	AutoAcceptCamAndMicCapture() bool // property
	// SetAutoAcceptCamAndMicCapture Set AutoAcceptCamAndMicCapture
	SetAutoAcceptCamAndMicCapture(AValue bool) // property
	// SupportsCompositionController
	//  Returns true if the current WebView2 runtime version supports Composition Controllers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3">See the ICoreWebView2Environment3 article.</a>
	SupportsCompositionController() bool // property
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`
	//  Provide a list of all process using same user data folder except for crashpad process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8">See the ICoreWebView2Environment8 article.</a>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property
	// SupportsControllerOptions
	//  Returns true if the current WebView2 runtime version supports Controller Options.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10">See the ICoreWebView2Environment10 article.</a>
	SupportsControllerOptions() bool // property
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
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment11#get_failurereportfolderpath">See the ICoreWebView2Environment11 article.</a>
	FailureReportFolderPath() string // property
	// StartWebView2
	//  This function is used to initialize WebView2.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#createcorewebview2environmentwithoptions">See the Globals article.</a>
	StartWebView2() bool // function
	// CompareVersions
	//  This method is for anyone want to compare version correctly to determine
	//  which version is newer, older or same.Use it to determine whether
	//  to use webview2 or certain feature based upon version. Sets the value of
	//  aCompRslt to -1, 0 or 1 if aVersion1 is less than, equal or greater
	//  than aVersion2 respectively. Returns false if it fails to parse
	//  any of the version strings.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#comparebrowserversions">See the Globals article.</a>
	CompareVersions(aVersion1, aVersion2 string, aCompRslt *int32) bool // function
	// UpdateDeviceScaleFactor
	//  Update the DeviceScaleFactor property value with the current scale or the ForcedDeviceScaleFactor value.
	UpdateDeviceScaleFactor() // procedure
	// AppendErrorLog
	//  Append aText to the ErrorMessage property.
	AppendErrorLog(aText string) // procedure
}

// TWVLoader Parent: TComponent
//
//	Class used to simplify the WebView2 initialization and destruction.
type TWVLoader struct {
	TComponent
}

func NewWVLoader(AOwner IComponent) IWVLoader {
	r1 := wVLoaderImportAPI().SysCallN(14, GetObjectUintptr(AOwner))
	return AsWVLoader(r1)
}

func (m *TWVLoader) Environment() ICoreWebView2Environment {
	var resultCoreWebView2Environment uintptr
	wVLoaderImportAPI().SysCallN(29, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Environment)))
	return AsCoreWebView2Environment(resultCoreWebView2Environment)
}

func (m *TWVLoader) Status() TWV2LoaderStatus {
	r1 := wVLoaderImportAPI().SysCallN(56, m.Instance())
	return TWV2LoaderStatus(r1)
}

func (m *TWVLoader) AvailableBrowserVersion() string {
	r1 := wVLoaderImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) ErrorMessage() string {
	r1 := wVLoaderImportAPI().SysCallN(31, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) ErrorCode() (resultInt64 int64) {
	wVLoaderImportAPI().SysCallN(30, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TWVLoader) SetCurrentDir() bool {
	r1 := wVLoaderImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetSetCurrentDir(AValue bool) {
	wVLoaderImportAPI().SysCallN(52, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) Initialized() bool {
	r1 := wVLoaderImportAPI().SysCallN(40, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) InitializationError() bool {
	r1 := wVLoaderImportAPI().SysCallN(39, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) CheckFiles() bool {
	r1 := wVLoaderImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetCheckFiles(AValue bool) {
	wVLoaderImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ShowMessageDlg() bool {
	r1 := wVLoaderImportAPI().SysCallN(53, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetShowMessageDlg(AValue bool) {
	wVLoaderImportAPI().SysCallN(53, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) InitCOMLibrary() bool {
	r1 := wVLoaderImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetInitCOMLibrary(AValue bool) {
	wVLoaderImportAPI().SysCallN(38, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) CustomCommandLineSwitches() string {
	r1 := wVLoaderImportAPI().SysCallN(15, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) DeviceScaleFactor() (resultSingle float32) {
	r1 := wVLoaderImportAPI().SysCallN(19, m.Instance())
	return float32(r1)
}

func (m *TWVLoader) ReRaiseExceptions() bool {
	r1 := wVLoaderImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetReRaiseExceptions(AValue bool) {
	wVLoaderImportAPI().SysCallN(49, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) InstalledRuntimeVersion() string {
	r1 := wVLoaderImportAPI().SysCallN(41, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) LoaderDllPath() string {
	r1 := wVLoaderImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetLoaderDllPath(AValue string) {
	wVLoaderImportAPI().SysCallN(45, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) UseInternalLoader() bool {
	r1 := wVLoaderImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetUseInternalLoader(AValue bool) {
	wVLoaderImportAPI().SysCallN(63, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) BrowserExecPath() string {
	r1 := wVLoaderImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetBrowserExecPath(AValue string) {
	wVLoaderImportAPI().SysCallN(11, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) UserDataFolder() string {
	r1 := wVLoaderImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetUserDataFolder(AValue string) {
	wVLoaderImportAPI().SysCallN(64, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AdditionalBrowserArguments() string {
	r1 := wVLoaderImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetAdditionalBrowserArguments(AValue string) {
	wVLoaderImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) Language() string {
	r1 := wVLoaderImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetLanguage(AValue string) {
	wVLoaderImportAPI().SysCallN(44, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) TargetCompatibleBrowserVersion() string {
	r1 := wVLoaderImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetTargetCompatibleBrowserVersion(AValue string) {
	wVLoaderImportAPI().SysCallN(59, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AllowSingleSignOnUsingOSPrimaryAccount() bool {
	r1 := wVLoaderImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) {
	wVLoaderImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ExclusiveUserDataFolderAccess() bool {
	r1 := wVLoaderImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetExclusiveUserDataFolderAccess(AValue bool) {
	wVLoaderImportAPI().SysCallN(32, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) CustomCrashReportingEnabled() bool {
	r1 := wVLoaderImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetCustomCrashReportingEnabled(AValue bool) {
	wVLoaderImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableTrackingPrevention() bool {
	r1 := wVLoaderImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetEnableTrackingPrevention(AValue bool) {
	wVLoaderImportAPI().SysCallN(28, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AreBrowserExtensionsEnabled() bool {
	r1 := wVLoaderImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAreBrowserExtensionsEnabled(AValue bool) {
	wVLoaderImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableGPU() bool {
	r1 := wVLoaderImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetEnableGPU(AValue bool) {
	wVLoaderImportAPI().SysCallN(27, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) EnableFeatures() string {
	r1 := wVLoaderImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetEnableFeatures(AValue string) {
	wVLoaderImportAPI().SysCallN(26, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableFeatures() string {
	r1 := wVLoaderImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetDisableFeatures(AValue string) {
	wVLoaderImportAPI().SysCallN(23, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) EnableBlinkFeatures() string {
	r1 := wVLoaderImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetEnableBlinkFeatures(AValue string) {
	wVLoaderImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableBlinkFeatures() string {
	r1 := wVLoaderImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetDisableBlinkFeatures(AValue string) {
	wVLoaderImportAPI().SysCallN(21, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) BlinkSettings() string {
	r1 := wVLoaderImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetBlinkSettings(AValue string) {
	wVLoaderImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) ForceFieldTrials() string {
	r1 := wVLoaderImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetForceFieldTrials(AValue string) {
	wVLoaderImportAPI().SysCallN(35, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) ForceFieldTrialParams() string {
	r1 := wVLoaderImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetForceFieldTrialParams(AValue string) {
	wVLoaderImportAPI().SysCallN(34, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) SmartScreenProtectionEnabled() bool {
	r1 := wVLoaderImportAPI().SysCallN(54, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetSmartScreenProtectionEnabled(AValue bool) {
	wVLoaderImportAPI().SysCallN(54, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AllowInsecureLocalhost() bool {
	r1 := wVLoaderImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowInsecureLocalhost(AValue bool) {
	wVLoaderImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) DisableWebSecurity() bool {
	r1 := wVLoaderImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableWebSecurity(AValue bool) {
	wVLoaderImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) TouchEvents() TWVState {
	r1 := wVLoaderImportAPI().SysCallN(60, 0, m.Instance(), 0)
	return TWVState(r1)
}

func (m *TWVLoader) SetTouchEvents(AValue TWVState) {
	wVLoaderImportAPI().SysCallN(60, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) HyperlinkAuditing() bool {
	r1 := wVLoaderImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetHyperlinkAuditing(AValue bool) {
	wVLoaderImportAPI().SysCallN(37, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AutoplayPolicy() TWVAutoplayPolicy {
	r1 := wVLoaderImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TWVAutoplayPolicy(r1)
}

func (m *TWVLoader) SetAutoplayPolicy(AValue TWVAutoplayPolicy) {
	wVLoaderImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) MuteAudio() bool {
	r1 := wVLoaderImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetMuteAudio(AValue bool) {
	wVLoaderImportAPI().SysCallN(46, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) KioskPrinting() bool {
	r1 := wVLoaderImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetKioskPrinting(AValue bool) {
	wVLoaderImportAPI().SysCallN(43, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ProxySettings() IWVProxySettings {
	r1 := wVLoaderImportAPI().SysCallN(48, m.Instance())
	return AsWVProxySettings(r1)
}

func (m *TWVLoader) AllowFileAccessFromFiles() bool {
	r1 := wVLoaderImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowFileAccessFromFiles(AValue bool) {
	wVLoaderImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) AllowRunningInsecureContent() bool {
	r1 := wVLoaderImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAllowRunningInsecureContent(AValue bool) {
	wVLoaderImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) DisableBackgroundNetworking() bool {
	r1 := wVLoaderImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableBackgroundNetworking(AValue bool) {
	wVLoaderImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) ForcedDeviceScaleFactor() (resultSingle float32) {
	r1 := wVLoaderImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return float32(r1)
}

func (m *TWVLoader) SetForcedDeviceScaleFactor(AValue float32) {
	wVLoaderImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) RemoteDebuggingPort() int32 {
	r1 := wVLoaderImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TWVLoader) SetRemoteDebuggingPort(AValue int32) {
	wVLoaderImportAPI().SysCallN(51, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) RemoteAllowOrigins() string {
	r1 := wVLoaderImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetRemoteAllowOrigins(AValue string) {
	wVLoaderImportAPI().SysCallN(50, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DebugLog() TWV2DebugLog {
	r1 := wVLoaderImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TWV2DebugLog(r1)
}

func (m *TWVLoader) SetDebugLog(AValue TWV2DebugLog) {
	wVLoaderImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) DebugLogLevel() TWV2DebugLogLevel {
	r1 := wVLoaderImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TWV2DebugLogLevel(r1)
}

func (m *TWVLoader) SetDebugLogLevel(AValue TWV2DebugLogLevel) {
	wVLoaderImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVLoader) JavaScriptFlags() string {
	r1 := wVLoaderImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetJavaScriptFlags(AValue string) {
	wVLoaderImportAPI().SysCallN(42, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) DisableEdgePitchNotification() bool {
	r1 := wVLoaderImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetDisableEdgePitchNotification(AValue bool) {
	wVLoaderImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) TreatInsecureOriginAsSecure() string {
	r1 := wVLoaderImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVLoader) SetTreatInsecureOriginAsSecure(AValue string) {
	wVLoaderImportAPI().SysCallN(61, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVLoader) AutoAcceptCamAndMicCapture() bool {
	r1 := wVLoaderImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVLoader) SetAutoAcceptCamAndMicCapture(AValue bool) {
	wVLoaderImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVLoader) SupportsCompositionController() bool {
	r1 := wVLoaderImportAPI().SysCallN(57, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) ProcessInfos() ICoreWebView2ProcessInfoCollection {
	var resultCoreWebView2ProcessInfoCollection uintptr
	wVLoaderImportAPI().SysCallN(47, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfoCollection)))
	return AsCoreWebView2ProcessInfoCollection(resultCoreWebView2ProcessInfoCollection)
}

func (m *TWVLoader) SupportsControllerOptions() bool {
	r1 := wVLoaderImportAPI().SysCallN(58, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) FailureReportFolderPath() string {
	r1 := wVLoaderImportAPI().SysCallN(33, m.Instance())
	return GoStr(r1)
}

func (m *TWVLoader) StartWebView2() bool {
	r1 := wVLoaderImportAPI().SysCallN(55, m.Instance())
	return GoBool(r1)
}

func (m *TWVLoader) CompareVersions(aVersion1, aVersion2 string, aCompRslt *int32) bool {
	var result1 uintptr
	r1 := wVLoaderImportAPI().SysCallN(13, m.Instance(), PascalStr(aVersion1), PascalStr(aVersion2), uintptr(unsafePointer(&result1)))
	*aCompRslt = int32(result1)
	return GoBool(r1)
}

func (m *TWVLoader) UpdateDeviceScaleFactor() {
	wVLoaderImportAPI().SysCallN(62, m.Instance())
}

func (m *TWVLoader) AppendErrorLog(aText string) {
	wVLoaderImportAPI().SysCallN(5, m.Instance(), PascalStr(aText))
}

var (
	wVLoaderImport       *imports.Imports = nil
	wVLoaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WVLoader_AdditionalBrowserArguments", 0),
		/*1*/ imports.NewTable("WVLoader_AllowFileAccessFromFiles", 0),
		/*2*/ imports.NewTable("WVLoader_AllowInsecureLocalhost", 0),
		/*3*/ imports.NewTable("WVLoader_AllowRunningInsecureContent", 0),
		/*4*/ imports.NewTable("WVLoader_AllowSingleSignOnUsingOSPrimaryAccount", 0),
		/*5*/ imports.NewTable("WVLoader_AppendErrorLog", 0),
		/*6*/ imports.NewTable("WVLoader_AreBrowserExtensionsEnabled", 0),
		/*7*/ imports.NewTable("WVLoader_AutoAcceptCamAndMicCapture", 0),
		/*8*/ imports.NewTable("WVLoader_AutoplayPolicy", 0),
		/*9*/ imports.NewTable("WVLoader_AvailableBrowserVersion", 0),
		/*10*/ imports.NewTable("WVLoader_BlinkSettings", 0),
		/*11*/ imports.NewTable("WVLoader_BrowserExecPath", 0),
		/*12*/ imports.NewTable("WVLoader_CheckFiles", 0),
		/*13*/ imports.NewTable("WVLoader_CompareVersions", 0),
		/*14*/ imports.NewTable("WVLoader_Create", 0),
		/*15*/ imports.NewTable("WVLoader_CustomCommandLineSwitches", 0),
		/*16*/ imports.NewTable("WVLoader_CustomCrashReportingEnabled", 0),
		/*17*/ imports.NewTable("WVLoader_DebugLog", 0),
		/*18*/ imports.NewTable("WVLoader_DebugLogLevel", 0),
		/*19*/ imports.NewTable("WVLoader_DeviceScaleFactor", 0),
		/*20*/ imports.NewTable("WVLoader_DisableBackgroundNetworking", 0),
		/*21*/ imports.NewTable("WVLoader_DisableBlinkFeatures", 0),
		/*22*/ imports.NewTable("WVLoader_DisableEdgePitchNotification", 0),
		/*23*/ imports.NewTable("WVLoader_DisableFeatures", 0),
		/*24*/ imports.NewTable("WVLoader_DisableWebSecurity", 0),
		/*25*/ imports.NewTable("WVLoader_EnableBlinkFeatures", 0),
		/*26*/ imports.NewTable("WVLoader_EnableFeatures", 0),
		/*27*/ imports.NewTable("WVLoader_EnableGPU", 0),
		/*28*/ imports.NewTable("WVLoader_EnableTrackingPrevention", 0),
		/*29*/ imports.NewTable("WVLoader_Environment", 0),
		/*30*/ imports.NewTable("WVLoader_ErrorCode", 0),
		/*31*/ imports.NewTable("WVLoader_ErrorMessage", 0),
		/*32*/ imports.NewTable("WVLoader_ExclusiveUserDataFolderAccess", 0),
		/*33*/ imports.NewTable("WVLoader_FailureReportFolderPath", 0),
		/*34*/ imports.NewTable("WVLoader_ForceFieldTrialParams", 0),
		/*35*/ imports.NewTable("WVLoader_ForceFieldTrials", 0),
		/*36*/ imports.NewTable("WVLoader_ForcedDeviceScaleFactor", 0),
		/*37*/ imports.NewTable("WVLoader_HyperlinkAuditing", 0),
		/*38*/ imports.NewTable("WVLoader_InitCOMLibrary", 0),
		/*39*/ imports.NewTable("WVLoader_InitializationError", 0),
		/*40*/ imports.NewTable("WVLoader_Initialized", 0),
		/*41*/ imports.NewTable("WVLoader_InstalledRuntimeVersion", 0),
		/*42*/ imports.NewTable("WVLoader_JavaScriptFlags", 0),
		/*43*/ imports.NewTable("WVLoader_KioskPrinting", 0),
		/*44*/ imports.NewTable("WVLoader_Language", 0),
		/*45*/ imports.NewTable("WVLoader_LoaderDllPath", 0),
		/*46*/ imports.NewTable("WVLoader_MuteAudio", 0),
		/*47*/ imports.NewTable("WVLoader_ProcessInfos", 0),
		/*48*/ imports.NewTable("WVLoader_ProxySettings", 0),
		/*49*/ imports.NewTable("WVLoader_ReRaiseExceptions", 0),
		/*50*/ imports.NewTable("WVLoader_RemoteAllowOrigins", 0),
		/*51*/ imports.NewTable("WVLoader_RemoteDebuggingPort", 0),
		/*52*/ imports.NewTable("WVLoader_SetCurrentDir", 0),
		/*53*/ imports.NewTable("WVLoader_ShowMessageDlg", 0),
		/*54*/ imports.NewTable("WVLoader_SmartScreenProtectionEnabled", 0),
		/*55*/ imports.NewTable("WVLoader_StartWebView2", 0),
		/*56*/ imports.NewTable("WVLoader_Status", 0),
		/*57*/ imports.NewTable("WVLoader_SupportsCompositionController", 0),
		/*58*/ imports.NewTable("WVLoader_SupportsControllerOptions", 0),
		/*59*/ imports.NewTable("WVLoader_TargetCompatibleBrowserVersion", 0),
		/*60*/ imports.NewTable("WVLoader_TouchEvents", 0),
		/*61*/ imports.NewTable("WVLoader_TreatInsecureOriginAsSecure", 0),
		/*62*/ imports.NewTable("WVLoader_UpdateDeviceScaleFactor", 0),
		/*63*/ imports.NewTable("WVLoader_UseInternalLoader", 0),
		/*64*/ imports.NewTable("WVLoader_UserDataFolder", 0),
	}
)

func wVLoaderImportAPI() *imports.Imports {
	if wVLoaderImport == nil {
		wVLoaderImport = NewDefaultImports()
		wVLoaderImport.SetImportTable(wVLoaderImportTables)
		wVLoaderImportTables = nil
	}
	return wVLoaderImport
}
