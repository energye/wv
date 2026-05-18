//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkSettings Parent: IObject
type IWkSettings interface {
	IObject
	GetEnableJavascript() bool                                                            // function
	GetAutoLoadImages() bool                                                              // function
	GetEnableHtml5LocalStorage() bool                                                     // function
	GetEnableHtml5Database() bool                                                         // function
	GetJavascriptCanOpenWindowsAutomatically() bool                                       // function
	GetEnableHyperlinkAuditing() bool                                                     // function
	GetDefaultFontFamily() string                                                         // function
	GetMonospaceFontFamily() string                                                       // function
	GetSerifFontFamily() string                                                           // function
	GetSansSerifFontFamily() string                                                       // function
	GetCursiveFontFamily() string                                                         // function
	GetFantasyFontFamily() string                                                         // function
	GetPictographFontFamily() string                                                      // function
	GetDefaultFontSize() int32                                                            // function
	GetDefaultMonospaceFontSize() int32                                                   // function
	GetMinimumFontSize() int32                                                            // function
	GetDefaultCharset() string                                                            // function
	GetEnableDeveloperExtras() bool                                                       // function
	GetEnableResizableTextAreas() bool                                                    // function
	GetEnableTabsToLinks() bool                                                           // function
	GetEnableCaretBrowsing() bool                                                         // function
	GetEnableFullscreen() bool                                                            // function
	GetPrintBackgrounds() bool                                                            // function
	GetEnableWebaudio() bool                                                              // function
	GetEnableWebgl() bool                                                                 // function
	GetAllowModalDialogs() bool                                                           // function
	GetZoomTextOnly() bool                                                                // function
	GetJavascriptCanAccessClipboard() bool                                                // function
	GetMediaPlaybackRequiresUserGesture() bool                                            // function
	GetMediaPlaybackAllowsInline() bool                                                   // function
	GetDrawCompositingIndicators() bool                                                   // function
	GetEnableSiteSpecificQuirks() bool                                                    // function
	GetEnablePageCache() bool                                                             // function
	GetUserAgent() string                                                                 // function
	GetEnableSmoothScrolling() bool                                                       // function
	GetEnable2dCanvasAcceleration() bool                                                  // function
	GetEnableWriteConsoleMessagesToStdout() bool                                          // function
	GetEnableMediaStream() bool                                                           // function
	GetEnableMockCaptureDevices() bool                                                    // function
	GetEnableSpatialNavigation() bool                                                     // function
	GetEnableMediasource() bool                                                           // function
	GetEnableEncryptedMedia() bool                                                        // function
	GetEnableMediaCapabilities() bool                                                     // function
	GetAllowFileAccessFromFileUrls() bool                                                 // function
	GetAllowUniversalAccessFromFileUrls() bool                                            // function
	GetAllowTopNavigationToDataUrls() bool                                                // function
	GetHardwareAccelerationPolicy() wvTypes.WebKitHardwareAccelerationPolicy              // function
	GetEnableBackForwardNavigationGestures() bool                                         // function
	GetEnableJavascriptMarkup() bool                                                      // function
	GetEnableMedia() bool                                                                 // function
	GetMediaContentTypesRequiringHardwareSupport() string                                 // function
	GetEnableWebrtc() bool                                                                // function
	GetDisableWebSecurity() bool                                                          // function
	Data() wvTypes.WebKitSettings                                                         // function
	SetEnableJavascript(enabled bool)                                                     // procedure
	SetAutoLoadImages(enabled bool)                                                       // procedure
	SetEnableHtml5LocalStorage(enabled bool)                                              // procedure
	SetEnableHtml5Database(enabled bool)                                                  // procedure
	SetJavascriptCanOpenWindowsAutomatically(enabled bool)                                // procedure
	SetEnableHyperlinkAuditing(enabled bool)                                              // procedure
	SetDefaultFontFamily(defaultFontFamily string)                                        // procedure
	SetMonospaceFontFamily(monospaceFontFamily string)                                    // procedure
	SetSerifFontFamily(serifFontFamily string)                                            // procedure
	SetSansSerifFontFamily(sansSerifFontFamily string)                                    // procedure
	SetCursiveFontFamily(cursiveFontFamily string)                                        // procedure
	SetFantasyFontFamily(fantasyFontFamily string)                                        // procedure
	SetPictographFontFamily(pictographFontFamily string)                                  // procedure
	SetDefaultFontSize(fontSize int32)                                                    // procedure
	SetDefaultMonospaceFontSize(fontSize int32)                                           // procedure
	SetMinimumFontSize(fontSize int32)                                                    // procedure
	SetDefaultCharset(defaultCharset string)                                              // procedure
	SetEnableDeveloperExtras(enabled bool)                                                // procedure
	SetEnableResizableTextAreas(enabled bool)                                             // procedure
	SetEnableTabsToLinks(enabled bool)                                                    // procedure
	SetEnableCaretBrowsing(enabled bool)                                                  // procedure
	SetEnableFullscreen(enabled bool)                                                     // procedure
	SetPrintBackgrounds(printBackgrounds bool)                                            // procedure
	SetEnableWebaudio(enabled bool)                                                       // procedure
	SetEnableWebgl(enabled bool)                                                          // procedure
	SetAllowModalDialogs(allowed bool)                                                    // procedure
	SetZoomTextOnly(zoomTextOnly bool)                                                    // procedure
	SetJavascriptCanAccessClipboard(enabled bool)                                         // procedure
	SetMediaPlaybackRequiresUserGesture(enabled bool)                                     // procedure
	SetMediaPlaybackAllowsInline(enabled bool)                                            // procedure
	SetDrawCompositingIndicators(enabled bool)                                            // procedure
	SetEnableSiteSpecificQuirks(enabled bool)                                             // procedure
	SetEnablePageCache(enabled bool)                                                      // procedure
	SetUserAgent(userAgent string)                                                        // procedure
	SetUserAgentWithApplicationDetails(applicationName string, applicationVersion string) // procedure
	SetEnableSmoothScrolling(enabled bool)                                                // procedure
	SetEnable2dCanvasAcceleration(enabled bool)                                           // procedure
	SetEnableWriteConsoleMessagesToStdout(enabled bool)                                   // procedure
	SetEnableMediaStream(enabled bool)                                                    // procedure
	SetEnableMockCaptureDevices(enabled bool)                                             // procedure
	SetEnableSpatialNavigation(enabled bool)                                              // procedure
	SetEnableMediasource(enabled bool)                                                    // procedure
	SetEnableEncryptedMedia(enabled bool)                                                 // procedure
	SetEnableMediaCapabilities(enabled bool)                                              // procedure
	SetAllowFileAccessFromFileUrls(allowed bool)                                          // procedure
	SetAllowUniversalAccessFromFileUrls(allowed bool)                                     // procedure
	SetAllowTopNavigationToDataUrls(allowed bool)                                         // procedure
	SetHardwareAccelerationPolicy(policy wvTypes.WebKitHardwareAccelerationPolicy)        // procedure
	SetEnableBackForwardNavigationGestures(enabled bool)                                  // procedure
	SetEnableJavascriptMarkup(enabled bool)                                               // procedure
	SetEnableMedia(enabled bool)                                                          // procedure
	SetMediaContentTypesRequiringHardwareSupport(contentTypes string)                     // procedure
	SetEnableWebrtc(enabled bool)                                                         // procedure
	SetDisableWebSecurity(disabled bool)                                                  // procedure
}

type TWkSettings struct {
	TObject
}

func (m *TWkSettings) GetEnableJavascript() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetAutoLoadImages() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableHtml5LocalStorage() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableHtml5Database() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetJavascriptCanOpenWindowsAutomatically() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableHyperlinkAuditing() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetDefaultFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetMonospaceFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetSerifFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetSansSerifFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetCursiveFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetFantasyFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetPictographFontFamily() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetDefaultFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkSettingsAPI().SysCallN(15, m.Instance())
	return int32(r)
}

func (m *TWkSettings) GetDefaultMonospaceFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkSettingsAPI().SysCallN(16, m.Instance())
	return int32(r)
}

func (m *TWkSettings) GetMinimumFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkSettingsAPI().SysCallN(17, m.Instance())
	return int32(r)
}

func (m *TWkSettings) GetDefaultCharset() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetEnableDeveloperExtras() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(19, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableResizableTextAreas() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableTabsToLinks() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableCaretBrowsing() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(22, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetPrintBackgrounds() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(24, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableWebaudio() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(25, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableWebgl() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(26, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetAllowModalDialogs() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(27, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetZoomTextOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetJavascriptCanAccessClipboard() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetMediaPlaybackRequiresUserGesture() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetMediaPlaybackAllowsInline() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetDrawCompositingIndicators() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(32, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableSiteSpecificQuirks() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(33, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnablePageCache() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(34, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetUserAgent() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(35, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetEnableSmoothScrolling() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(36, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnable2dCanvasAcceleration() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(37, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableWriteConsoleMessagesToStdout() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(38, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableMediaStream() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(39, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableMockCaptureDevices() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(40, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableSpatialNavigation() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(41, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableMediasource() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(42, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableEncryptedMedia() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(43, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableMediaCapabilities() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(44, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetAllowFileAccessFromFileUrls() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(45, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetAllowUniversalAccessFromFileUrls() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(46, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetAllowTopNavigationToDataUrls() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(47, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetHardwareAccelerationPolicy() wvTypes.WebKitHardwareAccelerationPolicy {
	if !m.IsValid() {
		return 0
	}
	r := wkSettingsAPI().SysCallN(48, m.Instance())
	return wvTypes.WebKitHardwareAccelerationPolicy(r)
}

func (m *TWkSettings) GetEnableBackForwardNavigationGestures() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(49, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableJavascriptMarkup() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(50, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetEnableMedia() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(51, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetMediaContentTypesRequiringHardwareSupport() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkSettingsAPI().SysCallN(52, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkSettings) GetEnableWebrtc() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(53, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) GetDisableWebSecurity() bool {
	if !m.IsValid() {
		return false
	}
	r := wkSettingsAPI().SysCallN(54, m.Instance())
	return api.GoBool(r)
}

func (m *TWkSettings) Data() wvTypes.WebKitSettings {
	if !m.IsValid() {
		return 0
	}
	r := wkSettingsAPI().SysCallN(55, m.Instance())
	return wvTypes.WebKitSettings(r)
}

func (m *TWkSettings) SetEnableJavascript(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(56, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetAutoLoadImages(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(57, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableHtml5LocalStorage(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(58, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableHtml5Database(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(59, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetJavascriptCanOpenWindowsAutomatically(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(60, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableHyperlinkAuditing(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(61, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetDefaultFontFamily(defaultFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(62, m.Instance(), api.PasStr(defaultFontFamily))
}

func (m *TWkSettings) SetMonospaceFontFamily(monospaceFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(63, m.Instance(), api.PasStr(monospaceFontFamily))
}

func (m *TWkSettings) SetSerifFontFamily(serifFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(64, m.Instance(), api.PasStr(serifFontFamily))
}

func (m *TWkSettings) SetSansSerifFontFamily(sansSerifFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(65, m.Instance(), api.PasStr(sansSerifFontFamily))
}

func (m *TWkSettings) SetCursiveFontFamily(cursiveFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(66, m.Instance(), api.PasStr(cursiveFontFamily))
}

func (m *TWkSettings) SetFantasyFontFamily(fantasyFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(67, m.Instance(), api.PasStr(fantasyFontFamily))
}

func (m *TWkSettings) SetPictographFontFamily(pictographFontFamily string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(68, m.Instance(), api.PasStr(pictographFontFamily))
}

func (m *TWkSettings) SetDefaultFontSize(fontSize int32) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(69, m.Instance(), uintptr(fontSize))
}

func (m *TWkSettings) SetDefaultMonospaceFontSize(fontSize int32) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(70, m.Instance(), uintptr(fontSize))
}

func (m *TWkSettings) SetMinimumFontSize(fontSize int32) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(71, m.Instance(), uintptr(fontSize))
}

func (m *TWkSettings) SetDefaultCharset(defaultCharset string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(72, m.Instance(), api.PasStr(defaultCharset))
}

func (m *TWkSettings) SetEnableDeveloperExtras(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(73, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableResizableTextAreas(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(74, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableTabsToLinks(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(75, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableCaretBrowsing(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(76, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableFullscreen(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(77, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetPrintBackgrounds(printBackgrounds bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(78, m.Instance(), api.PasBool(printBackgrounds))
}

func (m *TWkSettings) SetEnableWebaudio(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(79, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableWebgl(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(80, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetAllowModalDialogs(allowed bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(81, m.Instance(), api.PasBool(allowed))
}

func (m *TWkSettings) SetZoomTextOnly(zoomTextOnly bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(82, m.Instance(), api.PasBool(zoomTextOnly))
}

func (m *TWkSettings) SetJavascriptCanAccessClipboard(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(83, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetMediaPlaybackRequiresUserGesture(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(84, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetMediaPlaybackAllowsInline(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(85, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetDrawCompositingIndicators(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(86, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableSiteSpecificQuirks(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(87, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnablePageCache(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(88, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetUserAgent(userAgent string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(89, m.Instance(), api.PasStr(userAgent))
}

func (m *TWkSettings) SetUserAgentWithApplicationDetails(applicationName string, applicationVersion string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(90, m.Instance(), api.PasStr(applicationName), api.PasStr(applicationVersion))
}

func (m *TWkSettings) SetEnableSmoothScrolling(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(91, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnable2dCanvasAcceleration(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(92, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableWriteConsoleMessagesToStdout(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(93, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableMediaStream(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(94, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableMockCaptureDevices(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(95, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableSpatialNavigation(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(96, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableMediasource(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(97, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableEncryptedMedia(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(98, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableMediaCapabilities(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(99, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetAllowFileAccessFromFileUrls(allowed bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(100, m.Instance(), api.PasBool(allowed))
}

func (m *TWkSettings) SetAllowUniversalAccessFromFileUrls(allowed bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(101, m.Instance(), api.PasBool(allowed))
}

func (m *TWkSettings) SetAllowTopNavigationToDataUrls(allowed bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(102, m.Instance(), api.PasBool(allowed))
}

func (m *TWkSettings) SetHardwareAccelerationPolicy(policy wvTypes.WebKitHardwareAccelerationPolicy) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(103, m.Instance(), uintptr(policy))
}

func (m *TWkSettings) SetEnableBackForwardNavigationGestures(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(104, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableJavascriptMarkup(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(105, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetEnableMedia(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(106, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetMediaContentTypesRequiringHardwareSupport(contentTypes string) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(107, m.Instance(), api.PasStr(contentTypes))
}

func (m *TWkSettings) SetEnableWebrtc(enabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(108, m.Instance(), api.PasBool(enabled))
}

func (m *TWkSettings) SetDisableWebSecurity(disabled bool) {
	if !m.IsValid() {
		return
	}
	wkSettingsAPI().SysCallN(109, m.Instance(), api.PasBool(disabled))
}

// NewSettings class constructor
func NewSettings() IWkSettings {
	r := wkSettingsAPI().SysCallN(0)
	return AsWkSettings(r)
}

// NewSettingsWithWebKitSettings class constructor
func NewSettingsWithWebKitSettings(setting wvTypes.WebKitSettings) IWkSettings {
	r := wkSettingsAPI().SysCallN(1, uintptr(setting))
	return AsWkSettings(r)
}

var (
	wkSettingsOnce   base.Once
	wkSettingsImport *imports.Imports = nil
)

func wkSettingsAPI() *imports.Imports {
	wkSettingsOnce.Do(func() {
		wkSettingsImport = api.NewDefaultImports()
		wkSettingsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkSettings_Create", 0), // constructor NewSettings
			/* 1 */ imports.NewTable("TWkSettings_CreateWithWebKitSettings", 0), // constructor NewSettingsWithWebKitSettings
			/* 2 */ imports.NewTable("TWkSettings_get_enable_javascript", 0), // function GetEnableJavascript
			/* 3 */ imports.NewTable("TWkSettings_get_auto_load_images", 0), // function GetAutoLoadImages
			/* 4 */ imports.NewTable("TWkSettings_get_enable_html5_local_storage", 0), // function GetEnableHtml5LocalStorage
			/* 5 */ imports.NewTable("TWkSettings_get_enable_html5_database", 0), // function GetEnableHtml5Database
			/* 6 */ imports.NewTable("TWkSettings_get_javascript_can_open_windows_automatically", 0), // function GetJavascriptCanOpenWindowsAutomatically
			/* 7 */ imports.NewTable("TWkSettings_get_enable_hyperlink_auditing", 0), // function GetEnableHyperlinkAuditing
			/* 8 */ imports.NewTable("TWkSettings_get_default_font_family", 0), // function GetDefaultFontFamily
			/* 9 */ imports.NewTable("TWkSettings_get_monospace_font_family", 0), // function GetMonospaceFontFamily
			/* 10 */ imports.NewTable("TWkSettings_get_serif_font_family", 0), // function GetSerifFontFamily
			/* 11 */ imports.NewTable("TWkSettings_get_sans_serif_font_family", 0), // function GetSansSerifFontFamily
			/* 12 */ imports.NewTable("TWkSettings_get_cursive_font_family", 0), // function GetCursiveFontFamily
			/* 13 */ imports.NewTable("TWkSettings_get_fantasy_font_family", 0), // function GetFantasyFontFamily
			/* 14 */ imports.NewTable("TWkSettings_get_pictograph_font_family", 0), // function GetPictographFontFamily
			/* 15 */ imports.NewTable("TWkSettings_get_default_font_size", 0), // function GetDefaultFontSize
			/* 16 */ imports.NewTable("TWkSettings_get_default_monospace_font_size", 0), // function GetDefaultMonospaceFontSize
			/* 17 */ imports.NewTable("TWkSettings_get_minimum_font_size", 0), // function GetMinimumFontSize
			/* 18 */ imports.NewTable("TWkSettings_get_default_charset", 0), // function GetDefaultCharset
			/* 19 */ imports.NewTable("TWkSettings_get_enable_developer_extras", 0), // function GetEnableDeveloperExtras
			/* 20 */ imports.NewTable("TWkSettings_get_enable_resizable_text_areas", 0), // function GetEnableResizableTextAreas
			/* 21 */ imports.NewTable("TWkSettings_get_enable_tabs_to_links", 0), // function GetEnableTabsToLinks
			/* 22 */ imports.NewTable("TWkSettings_get_enable_caret_browsing", 0), // function GetEnableCaretBrowsing
			/* 23 */ imports.NewTable("TWkSettings_get_enable_fullscreen", 0), // function GetEnableFullscreen
			/* 24 */ imports.NewTable("TWkSettings_get_print_backgrounds", 0), // function GetPrintBackgrounds
			/* 25 */ imports.NewTable("TWkSettings_get_enable_webaudio", 0), // function GetEnableWebaudio
			/* 26 */ imports.NewTable("TWkSettings_get_enable_webgl", 0), // function GetEnableWebgl
			/* 27 */ imports.NewTable("TWkSettings_get_allow_modal_dialogs", 0), // function GetAllowModalDialogs
			/* 28 */ imports.NewTable("TWkSettings_get_zoom_text_only", 0), // function GetZoomTextOnly
			/* 29 */ imports.NewTable("TWkSettings_get_javascript_can_access_clipboard", 0), // function GetJavascriptCanAccessClipboard
			/* 30 */ imports.NewTable("TWkSettings_get_media_playback_requires_user_gesture", 0), // function GetMediaPlaybackRequiresUserGesture
			/* 31 */ imports.NewTable("TWkSettings_get_media_playback_allows_inline", 0), // function GetMediaPlaybackAllowsInline
			/* 32 */ imports.NewTable("TWkSettings_get_draw_compositing_indicators", 0), // function GetDrawCompositingIndicators
			/* 33 */ imports.NewTable("TWkSettings_get_enable_site_specific_quirks", 0), // function GetEnableSiteSpecificQuirks
			/* 34 */ imports.NewTable("TWkSettings_get_enable_page_cache", 0), // function GetEnablePageCache
			/* 35 */ imports.NewTable("TWkSettings_get_user_agent", 0), // function GetUserAgent
			/* 36 */ imports.NewTable("TWkSettings_get_enable_smooth_scrolling", 0), // function GetEnableSmoothScrolling
			/* 37 */ imports.NewTable("TWkSettings_get_enable_2d_canvas_acceleration", 0), // function GetEnable2dCanvasAcceleration
			/* 38 */ imports.NewTable("TWkSettings_get_enable_write_console_messages_to_stdout", 0), // function GetEnableWriteConsoleMessagesToStdout
			/* 39 */ imports.NewTable("TWkSettings_get_enable_media_stream", 0), // function GetEnableMediaStream
			/* 40 */ imports.NewTable("TWkSettings_get_enable_mock_capture_devices", 0), // function GetEnableMockCaptureDevices
			/* 41 */ imports.NewTable("TWkSettings_get_enable_spatial_navigation", 0), // function GetEnableSpatialNavigation
			/* 42 */ imports.NewTable("TWkSettings_get_enable_mediasource", 0), // function GetEnableMediasource
			/* 43 */ imports.NewTable("TWkSettings_get_enable_encrypted_media", 0), // function GetEnableEncryptedMedia
			/* 44 */ imports.NewTable("TWkSettings_get_enable_media_capabilities", 0), // function GetEnableMediaCapabilities
			/* 45 */ imports.NewTable("TWkSettings_get_allow_file_access_from_file_urls", 0), // function GetAllowFileAccessFromFileUrls
			/* 46 */ imports.NewTable("TWkSettings_get_allow_universal_access_from_file_urls", 0), // function GetAllowUniversalAccessFromFileUrls
			/* 47 */ imports.NewTable("TWkSettings_get_allow_top_navigation_to_data_urls", 0), // function GetAllowTopNavigationToDataUrls
			/* 48 */ imports.NewTable("TWkSettings_get_hardware_acceleration_policy", 0), // function GetHardwareAccelerationPolicy
			/* 49 */ imports.NewTable("TWkSettings_get_enable_back_forward_navigation_gestures", 0), // function GetEnableBackForwardNavigationGestures
			/* 50 */ imports.NewTable("TWkSettings_get_enable_javascript_markup", 0), // function GetEnableJavascriptMarkup
			/* 51 */ imports.NewTable("TWkSettings_get_enable_media", 0), // function GetEnableMedia
			/* 52 */ imports.NewTable("TWkSettings_get_media_content_types_requiring_hardware_support", 0), // function GetMediaContentTypesRequiringHardwareSupport
			/* 53 */ imports.NewTable("TWkSettings_get_enable_webrtc", 0), // function GetEnableWebrtc
			/* 54 */ imports.NewTable("TWkSettings_get_disable_web_security", 0), // function GetDisableWebSecurity
			/* 55 */ imports.NewTable("TWkSettings_Data", 0), // function Data
			/* 56 */ imports.NewTable("TWkSettings_set_enable_javascript", 0), // procedure SetEnableJavascript
			/* 57 */ imports.NewTable("TWkSettings_set_auto_load_images", 0), // procedure SetAutoLoadImages
			/* 58 */ imports.NewTable("TWkSettings_set_enable_html5_local_storage", 0), // procedure SetEnableHtml5LocalStorage
			/* 59 */ imports.NewTable("TWkSettings_set_enable_html5_database", 0), // procedure SetEnableHtml5Database
			/* 60 */ imports.NewTable("TWkSettings_set_javascript_can_open_windows_automatically", 0), // procedure SetJavascriptCanOpenWindowsAutomatically
			/* 61 */ imports.NewTable("TWkSettings_set_enable_hyperlink_auditing", 0), // procedure SetEnableHyperlinkAuditing
			/* 62 */ imports.NewTable("TWkSettings_set_default_font_family", 0), // procedure SetDefaultFontFamily
			/* 63 */ imports.NewTable("TWkSettings_set_monospace_font_family", 0), // procedure SetMonospaceFontFamily
			/* 64 */ imports.NewTable("TWkSettings_set_serif_font_family", 0), // procedure SetSerifFontFamily
			/* 65 */ imports.NewTable("TWkSettings_set_sans_serif_font_family", 0), // procedure SetSansSerifFontFamily
			/* 66 */ imports.NewTable("TWkSettings_set_cursive_font_family", 0), // procedure SetCursiveFontFamily
			/* 67 */ imports.NewTable("TWkSettings_set_fantasy_font_family", 0), // procedure SetFantasyFontFamily
			/* 68 */ imports.NewTable("TWkSettings_set_pictograph_font_family", 0), // procedure SetPictographFontFamily
			/* 69 */ imports.NewTable("TWkSettings_set_default_font_size", 0), // procedure SetDefaultFontSize
			/* 70 */ imports.NewTable("TWkSettings_set_default_monospace_font_size", 0), // procedure SetDefaultMonospaceFontSize
			/* 71 */ imports.NewTable("TWkSettings_set_minimum_font_size", 0), // procedure SetMinimumFontSize
			/* 72 */ imports.NewTable("TWkSettings_set_default_charset", 0), // procedure SetDefaultCharset
			/* 73 */ imports.NewTable("TWkSettings_set_enable_developer_extras", 0), // procedure SetEnableDeveloperExtras
			/* 74 */ imports.NewTable("TWkSettings_set_enable_resizable_text_areas", 0), // procedure SetEnableResizableTextAreas
			/* 75 */ imports.NewTable("TWkSettings_set_enable_tabs_to_links", 0), // procedure SetEnableTabsToLinks
			/* 76 */ imports.NewTable("TWkSettings_set_enable_caret_browsing", 0), // procedure SetEnableCaretBrowsing
			/* 77 */ imports.NewTable("TWkSettings_set_enable_fullscreen", 0), // procedure SetEnableFullscreen
			/* 78 */ imports.NewTable("TWkSettings_set_print_backgrounds", 0), // procedure SetPrintBackgrounds
			/* 79 */ imports.NewTable("TWkSettings_set_enable_webaudio", 0), // procedure SetEnableWebaudio
			/* 80 */ imports.NewTable("TWkSettings_set_enable_webgl", 0), // procedure SetEnableWebgl
			/* 81 */ imports.NewTable("TWkSettings_set_allow_modal_dialogs", 0), // procedure SetAllowModalDialogs
			/* 82 */ imports.NewTable("TWkSettings_set_zoom_text_only", 0), // procedure SetZoomTextOnly
			/* 83 */ imports.NewTable("TWkSettings_set_javascript_can_access_clipboard", 0), // procedure SetJavascriptCanAccessClipboard
			/* 84 */ imports.NewTable("TWkSettings_set_media_playback_requires_user_gesture", 0), // procedure SetMediaPlaybackRequiresUserGesture
			/* 85 */ imports.NewTable("TWkSettings_set_media_playback_allows_inline", 0), // procedure SetMediaPlaybackAllowsInline
			/* 86 */ imports.NewTable("TWkSettings_set_draw_compositing_indicators", 0), // procedure SetDrawCompositingIndicators
			/* 87 */ imports.NewTable("TWkSettings_set_enable_site_specific_quirks", 0), // procedure SetEnableSiteSpecificQuirks
			/* 88 */ imports.NewTable("TWkSettings_set_enable_page_cache", 0), // procedure SetEnablePageCache
			/* 89 */ imports.NewTable("TWkSettings_set_user_agent", 0), // procedure SetUserAgent
			/* 90 */ imports.NewTable("TWkSettings_set_user_agent_with_application_details", 0), // procedure SetUserAgentWithApplicationDetails
			/* 91 */ imports.NewTable("TWkSettings_set_enable_smooth_scrolling", 0), // procedure SetEnableSmoothScrolling
			/* 92 */ imports.NewTable("TWkSettings_set_enable_2d_canvas_acceleration", 0), // procedure SetEnable2dCanvasAcceleration
			/* 93 */ imports.NewTable("TWkSettings_set_enable_write_console_messages_to_stdout", 0), // procedure SetEnableWriteConsoleMessagesToStdout
			/* 94 */ imports.NewTable("TWkSettings_set_enable_media_stream", 0), // procedure SetEnableMediaStream
			/* 95 */ imports.NewTable("TWkSettings_set_enable_mock_capture_devices", 0), // procedure SetEnableMockCaptureDevices
			/* 96 */ imports.NewTable("TWkSettings_set_enable_spatial_navigation", 0), // procedure SetEnableSpatialNavigation
			/* 97 */ imports.NewTable("TWkSettings_set_enable_mediasource", 0), // procedure SetEnableMediasource
			/* 98 */ imports.NewTable("TWkSettings_set_enable_encrypted_media", 0), // procedure SetEnableEncryptedMedia
			/* 99 */ imports.NewTable("TWkSettings_set_enable_media_capabilities", 0), // procedure SetEnableMediaCapabilities
			/* 100 */ imports.NewTable("TWkSettings_set_allow_file_access_from_file_urls", 0), // procedure SetAllowFileAccessFromFileUrls
			/* 101 */ imports.NewTable("TWkSettings_set_allow_universal_access_from_file_urls", 0), // procedure SetAllowUniversalAccessFromFileUrls
			/* 102 */ imports.NewTable("TWkSettings_set_allow_top_navigation_to_data_urls", 0), // procedure SetAllowTopNavigationToDataUrls
			/* 103 */ imports.NewTable("TWkSettings_set_hardware_acceleration_policy", 0), // procedure SetHardwareAccelerationPolicy
			/* 104 */ imports.NewTable("TWkSettings_set_enable_back_forward_navigation_gestures", 0), // procedure SetEnableBackForwardNavigationGestures
			/* 105 */ imports.NewTable("TWkSettings_set_enable_javascript_markup", 0), // procedure SetEnableJavascriptMarkup
			/* 106 */ imports.NewTable("TWkSettings_set_enable_media", 0), // procedure SetEnableMedia
			/* 107 */ imports.NewTable("TWkSettings_set_media_content_types_requiring_hardware_support", 0), // procedure SetMediaContentTypesRequiringHardwareSupport
			/* 108 */ imports.NewTable("TWkSettings_set_enable_webrtc", 0), // procedure SetEnableWebrtc
			/* 109 */ imports.NewTable("TWkSettings_set_disable_web_security", 0), // procedure SetDisableWebSecurity
		}
	})
	return wkSettingsImport
}
