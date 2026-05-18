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

// ICoreWebView2Profile Parent: IObject
type ICoreWebView2Profile interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// ClearBrowsingData
	//  Clear browsing data based on a data type. This method takes two parameters,
	//  the first being a mask of one or more `COREWEBVIEW2_BROWSING_DATA_KINDS`. OR
	//  operation(s) can be applied to multiple `COREWEBVIEW2_BROWSING_DATA_KINDS` to
	//  create a mask representing those data types. The browsing data kinds that are
	//  supported are listed below. These data kinds follow a hierarchical structure in
	//  which nested bullet points are included in their parent bullet point's data kind.
	//  Ex: All DOM storage is encompassed in all site data which is encompassed in
	//  all profile data.<code>
	//  * All Profile
	//  * All Site Data
	//  * All DOM Storage: File Systems, Indexed DB, Local Storage, Web SQL, Cache
	//  Storage
	//  * Cookies
	//  * Disk Cache
	//  * Download History
	//  * General Autofill
	//  * Password Autosave
	//  * Browsing History
	//  * Settings</code>
	//  The completed handler will be invoked when the browsing data has been cleared and
	//  will indicate if the specified data was properly cleared. In the case in which
	//  the operation is interrupted and the corresponding data is not fully cleared
	//  the handler will return `E_ABORT` and otherwise will return `S_OK`.
	//  Because this is an asynchronous operation, code that is dependent on the cleared
	//  data must be placed in the callback of this operation.
	//  If the WebView object is closed before the clear browsing data operation
	//  has completed, the handler will be released, but not invoked. In this case
	//  the clear browsing data operation may or may not be completed.
	//  ClearBrowsingData clears the `dataKinds` regardless of timestamp.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdata">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingData(dataKinds wvTypes.TWVBrowsingDataKinds, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// ClearBrowsingDataInTimeRange
	//  ClearBrowsingDataInTimeRange behaves like ClearBrowsingData except that it
	//  takes in two additional parameters for the start and end time for which it
	//  should clear the data between. The `startTime` and `endTime`
	//  parameters correspond to the number of seconds since the UNIX epoch.
	//  `startTime` is inclusive while `endTime` is exclusive, therefore the data will
	//  be cleared between [startTime, endTime).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataintimerange">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingDataInTimeRange(dataKinds wvTypes.TWVBrowsingDataKinds, startTime types.TDateTime, endTime types.TDateTime, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// ClearBrowsingDataAll
	//  ClearBrowsingDataAll behaves like ClearBrowsingData except that it
	//  clears the entirety of the data associated with the profile it is called on.
	//  It clears the data regardless of timestamp.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataall">See the ICoreWebView2Profile2 article.</see>
	ClearBrowsingDataAll(handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// SetPermissionState
	//  Sets permission state for the given permission kind and origin
	//  asynchronously. The change persists across sessions until it is changed by
	//  another call to `SetPermissionState`, or by setting the `State` property
	//  in `PermissionRequestedEventArgs`. Setting the state to
	//  `COREWEBVIEW2_PERMISSION_STATE_DEFAULT` will erase any state saved in the
	//  profile and restore the default behavior.
	//  The origin should have a valid scheme and host (e.g. "https://www.example.com"),
	//  otherwise the method fails with `E_INVALIDARG`. Additional URI parts like
	//  path and fragment are ignored. For example, "https://wwww.example.com/app1/index.html/"
	//  is treated the same as "https://wwww.example.com". See the
	//  [MDN origin definition](https://developer.mozilla.org/docs/Glossary/Origin)
	//  for more details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#setpermissionstate">See the ICoreWebView2Profile4 article.</see>
	SetPermissionState(permissionKind wvTypes.TWVPermissionKind, origin string, state wvTypes.TWVPermissionState, completedHandler ICoreWebView2SetPermissionStateCompletedHandler) bool // function
	// GetNonDefaultPermissionSettings
	//  Invokes the handler with a collection of all nondefault permission settings.
	//  Use this method to get the permission state set in the current and previous
	//  sessions.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#getnondefaultpermissionsettings">See the ICoreWebView2Profile4 article.</see>
	GetNonDefaultPermissionSettings(completedHandler ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) bool // function
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
	AddBrowserExtension(extensionFolderPath string, completedHandler ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) bool // function
	// GetBrowserExtensions
	//  Gets a snapshot of the set of extensions installed at the time `GetBrowserExtensions` is
	//  called. If an extension is installed or uninstalled after `GetBrowserExtensions` completes,
	//  the list returned by `GetBrowserExtensions` remains the same.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `GetBrowserExtensions` won't return any
	//  extensions on current user profile.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#getbrowserextensions">See the ICoreWebView2Profile7 article.</see>
	GetBrowserExtensions(completedHandler ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) bool // function
	// Delete
	//  After the API is called, the profile will be marked for deletion. The
	//  local profile's directory will be deleted at browser process exit. If it
	//  fails to delete, because something else is holding the files open,
	//  WebView2 will try to delete the profile at all future browser process
	//  starts until successful.
	//  The corresponding CoreWebView2s will be closed and the
	//  ICoreWebView2Profile.Deleted event will be raised. See
	//  `ICoreWebView2Profile.Deleted` for more information.
	//  If you try to create a new profile with the same name as an existing
	//  profile that has been marked as deleted but hasn't yet been deleted,
	//  profile creation will fail with HRESULT_FROM_WIN32(ERROR_DELETE_PENDING).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile8#delete">See the ICoreWebView2Profile8 article.</see>
	Delete() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Profile         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2Profile) // property BaseIntf Setter
	// ProfileName
	//  Name of the profile.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_profilename">See the ICoreWebView2Profile article.</see>
	ProfileName() string // property ProfileName Getter
	// IsInPrivateModeEnabled
	//  InPrivate mode is enabled or not.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_isinprivatemodeenabled">See the ICoreWebView2Profile article.</see>
	IsInPrivateModeEnabled() bool // property IsInPrivateModeEnabled Getter
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
	//  There is `ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention` property to enable/disable tracking prevention feature
	//  for all the WebView2's created in the same environment. If enabled, `PreferredTrackingPreventionLevel` is set to
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED` by default for all the WebView2's and profiles created in the same
	//  environment or is set to the level whatever value was last changed/persisted to the profile. If disabled
	//  `PreferredTrackingPreventionLevel` is not respected by WebView2. If `PreferredTrackingPreventionLevel` is set when the
	//  feature is disabled, the property value get changed and persisted but it will takes effect only if
	//  `ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention` is true.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile3#get_preferredtrackingpreventionlevel">See the ICoreWebView2Profile3 article.</see>
	PreferredTrackingPreventionLevel() wvTypes.TWVTrackingPreventionLevel         // property PreferredTrackingPreventionLevel Getter
	SetPreferredTrackingPreventionLevel(value wvTypes.TWVTrackingPreventionLevel) // property PreferredTrackingPreventionLevel Setter
	// CookieManager
	//  Get the cookie manager for the profile. All CoreWebView2s associated with this
	//  profile share the same cookie values. Changes to cookies in this cookie manager apply to all
	//  CoreWebView2s associated with this profile. See ICoreWebView2CookieManager.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile5#get_cookiemanager">See the ICoreWebView2Profile5 article.</see>
	CookieManager() ICoreWebView2CookieManager // property CookieManager Getter
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
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_ispasswordautosaveenabled">See the ICoreWebView2Profile6 article.</see>
	IsPasswordAutosaveEnabled() bool         // property IsPasswordAutosaveEnabled Getter
	SetIsPasswordAutosaveEnabled(value bool) // property IsPasswordAutosaveEnabled Setter
	// IsGeneralAutofillEnabled
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
	IsGeneralAutofillEnabled() bool         // property IsGeneralAutofillEnabled Getter
	SetIsGeneralAutofillEnabled(value bool) // property IsGeneralAutofillEnabled Setter
}

type TCoreWebView2Profile struct {
	TObject
}

func (m *TCoreWebView2Profile) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) ClearBrowsingData(dataKinds wvTypes.TWVBrowsingDataKinds, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(2, m.Instance(), uintptr(dataKinds), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) ClearBrowsingDataInTimeRange(dataKinds wvTypes.TWVBrowsingDataKinds, startTime types.TDateTime, endTime types.TDateTime, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(3, m.Instance(), uintptr(dataKinds), uintptr(base.UnsafePointer(&startTime)), uintptr(base.UnsafePointer(&endTime)), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) ClearBrowsingDataAll(handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(handler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) SetPermissionState(permissionKind wvTypes.TWVPermissionKind, origin string, state wvTypes.TWVPermissionState, completedHandler ICoreWebView2SetPermissionStateCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(5, m.Instance(), uintptr(permissionKind), api.PasStr(origin), uintptr(state), base.GetObjectUintptr(completedHandler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) GetNonDefaultPermissionSettings(completedHandler ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(completedHandler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) AddBrowserExtension(extensionFolderPath string, completedHandler ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(7, m.Instance(), api.PasStr(extensionFolderPath), base.GetObjectUintptr(completedHandler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) GetBrowserExtensions(completedHandler ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(completedHandler))
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) Delete() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) BaseIntf() (result ICoreWebView2Profile) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProfileAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Profile(resultPtr)
	return
}

func (m *TCoreWebView2Profile) SetBaseIntf(value ICoreWebView2Profile) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2Profile) ProfileName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProfileAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Profile) IsInPrivateModeEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) ProfilePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProfileAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Profile) DefaultDownloadFolderPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProfileAPI().SysCallN(15, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2Profile) SetDefaultDownloadFolderPath(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2Profile) PreferredColorScheme() wvTypes.TWVPreferredColorScheme {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProfileAPI().SysCallN(16, 0, m.Instance())
	return wvTypes.TWVPreferredColorScheme(r)
}

func (m *TCoreWebView2Profile) SetPreferredColorScheme(value wvTypes.TWVPreferredColorScheme) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Profile) PreferredTrackingPreventionLevel() wvTypes.TWVTrackingPreventionLevel {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProfileAPI().SysCallN(17, 0, m.Instance())
	return wvTypes.TWVTrackingPreventionLevel(r)
}

func (m *TCoreWebView2Profile) SetPreferredTrackingPreventionLevel(value wvTypes.TWVTrackingPreventionLevel) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Profile) CookieManager() (result ICoreWebView2CookieManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProfileAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CookieManager(resultPtr)
	return
}

func (m *TCoreWebView2Profile) IsPasswordAutosaveEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) SetIsPasswordAutosaveEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Profile) IsGeneralAutofillEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProfileAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Profile) SetIsGeneralAutofillEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProfileAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2Profile class constructor
func NewCoreWebView2Profile(baseIntf ICoreWebView2Profile) ICoreWebView2Profile {
	r := coreWebView2ProfileAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Profile(r)
}

var (
	coreWebView2ProfileOnce   base.Once
	coreWebView2ProfileImport *imports.Imports = nil
)

func coreWebView2ProfileAPI() *imports.Imports {
	coreWebView2ProfileOnce.Do(func() {
		coreWebView2ProfileImport = api.NewDefaultImports()
		coreWebView2ProfileImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Profile_Create", 0), // constructor NewCoreWebView2Profile
			/* 1 */ imports.NewTable("TCoreWebView2Profile_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2Profile_ClearBrowsingData", 0), // function ClearBrowsingData
			/* 3 */ imports.NewTable("TCoreWebView2Profile_ClearBrowsingDataInTimeRange", 0), // function ClearBrowsingDataInTimeRange
			/* 4 */ imports.NewTable("TCoreWebView2Profile_ClearBrowsingDataAll", 0), // function ClearBrowsingDataAll
			/* 5 */ imports.NewTable("TCoreWebView2Profile_SetPermissionState", 0), // function SetPermissionState
			/* 6 */ imports.NewTable("TCoreWebView2Profile_GetNonDefaultPermissionSettings", 0), // function GetNonDefaultPermissionSettings
			/* 7 */ imports.NewTable("TCoreWebView2Profile_AddBrowserExtension", 0), // function AddBrowserExtension
			/* 8 */ imports.NewTable("TCoreWebView2Profile_GetBrowserExtensions", 0), // function GetBrowserExtensions
			/* 9 */ imports.NewTable("TCoreWebView2Profile_Delete", 0), // function Delete
			/* 10 */ imports.NewTable("TCoreWebView2Profile_Initialized", 0), // property Initialized
			/* 11 */ imports.NewTable("TCoreWebView2Profile_BaseIntf", 0), // property BaseIntf
			/* 12 */ imports.NewTable("TCoreWebView2Profile_ProfileName", 0), // property ProfileName
			/* 13 */ imports.NewTable("TCoreWebView2Profile_IsInPrivateModeEnabled", 0), // property IsInPrivateModeEnabled
			/* 14 */ imports.NewTable("TCoreWebView2Profile_ProfilePath", 0), // property ProfilePath
			/* 15 */ imports.NewTable("TCoreWebView2Profile_DefaultDownloadFolderPath", 0), // property DefaultDownloadFolderPath
			/* 16 */ imports.NewTable("TCoreWebView2Profile_PreferredColorScheme", 0), // property PreferredColorScheme
			/* 17 */ imports.NewTable("TCoreWebView2Profile_PreferredTrackingPreventionLevel", 0), // property PreferredTrackingPreventionLevel
			/* 18 */ imports.NewTable("TCoreWebView2Profile_CookieManager", 0), // property CookieManager
			/* 19 */ imports.NewTable("TCoreWebView2Profile_IsPasswordAutosaveEnabled", 0), // property IsPasswordAutosaveEnabled
			/* 20 */ imports.NewTable("TCoreWebView2Profile_IsGeneralAutofillEnabled", 0), // property IsGeneralAutofillEnabled
		}
	})
	return coreWebView2ProfileImport
}
