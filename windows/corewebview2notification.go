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

// ICoreWebView2Notification Parent: IObject
type ICoreWebView2Notification interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// ReportShown
	//  The host may run this to report the notification has been displayed and it
	//  will cause the
	//  [show](https://developer.mozilla.org/docs/Web/API/Notification/show_event)
	//  event to be raised for non-persistent notifications. You must not run this
	//  unless you are handling the `NotificationReceived` event. Returns
	//  `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)` if `Handled` is `FALSE` when
	//  this is called.
	ReportShown() bool // function
	// ReportClicked
	//  The host may run this to report the notification has been clicked, and it
	//  will cause the
	//  [click](https://developer.mozilla.org/docs/Web/API/Notification/click_event)
	//  event to be raised for non-persistent notifications and the
	//  [notificationclick](https://developer.mozilla.org/docs/Web/API/ServiceWorkerGlobalScope/notificationclick_event)
	//  event for persistent notifications. Use `ReportClickedWithActionIndex` to
	//  specify an action to activate a persistent notification. You must not run
	//  this unless you are handling the `NotificationReceived` event. Returns
	//  `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)` if `Handled` is `FALSE` or
	//  `ReportShown` has not been run when this is called.
	ReportClicked() bool // function
	// ReportClosed
	//  The host may run this to report the notification was dismissed, and it
	//  will cause the
	//  [close](https://developer.mozilla.org/docs/Web/API/Notification/close_event)
	//  event to be raised for non-persistent notifications and the
	//  [notificationclose](https://developer.mozilla.org/docs/Web/API/ServiceWorkerGlobalScope/notificationclose_event)
	//  event for persistent notifications. You must not run this unless you are
	//  handling the `NotificationReceived` event. Returns
	//  `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)` if `Handled` is `FALSE` or
	//  `ReportShown` has not been run when this is called.
	ReportClosed() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Notification         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2Notification) // property BaseIntf Setter
	// Body
	//  A string representing the body text of the notification.
	//  The default value is an empty string.
	Body() string // property Body Getter
	// Direction
	//  The text direction in which to display the notification.
	//  This corresponds to
	//  [Notification.dir](https://developer.mozilla.org/docs/Web/API/Notification/dir)
	//  DOM API.
	//  The default value is `COREWEBVIEW2_TEXT_DIRECTION_KIND_DEFAULT`.
	Direction() wvTypes.TWVTextDirectionKind // property Direction Getter
	// Language
	//  The notification's language, as intended to be specified using a string
	//  representing a language tag (such as `en-US`) according to
	//  [BCP47](https://datatracker.ietf.org/doc/html/rfc5646). Note that no
	//  validation is performed on this property and it can be any string the
	//  notification sender specifies.
	//  This corresponds to
	//  [Notification.lang](https://developer.mozilla.org/docs/Web/API/Notification/lang)
	//  DOM API.
	//  The default value is an empty string.
	Language() string // property Language Getter
	// Tag
	//  A string representing an identifying tag for the notification.
	//  This corresponds to
	//  [Notification.tag](https://developer.mozilla.org/docs/Web/API/Notification/tag)
	//  DOM API.
	//  The default value is an empty string.
	Tag() string // property Tag Getter
	// IconUri
	//  A string containing the URI of an icon to be displayed in the
	//  notification.
	//  The default value is an empty string.
	IconUri() string // property IconUri Getter
	// Title
	//  The title of the notification.
	Title() string // property Title Getter
	// BadgeUri
	//  A string containing the URI of the image used to represent the
	//  notification when there isn't enough space to display the notification
	//  itself.
	//  The default value is an empty string.
	BadgeUri() string // property BadgeUri Getter
	// BodyImageUri
	//  A string containing the URI of an image to be displayed in the
	//  notification.
	//  The default value is an empty string.
	BodyImageUri() string // property BodyImageUri Getter
	// ShouldRenotify
	//  Indicates whether the user should be notified after a new notification
	//  replaces an old one.
	//  This corresponds to
	//  [Notification.renotify](https://developer.mozilla.org/docs/Web/API/Notification/renotify)
	//  DOM API.
	//  The default value is `FALSE`.
	ShouldRenotify() bool // property ShouldRenotify Getter
	// RequiresInteraction
	//  A boolean value indicating that a notification should remain active until
	//  the user clicks or dismisses it, rather than closing automatically.
	//  This corresponds to
	//  [Notification.requireInteraction](https://developer.mozilla.org/docs/Web/API/Notification/requireInteraction)
	//  DOM API. Note that you may not be able to necessarily implement this due to native API limitations.
	//  The default value is `FALSE`.
	RequiresInteraction() bool // property RequiresInteraction Getter
	// IsSilent
	//  Indicates whether the notification should be silent -- i.e., no sounds or
	//  vibrations should be issued, regardless of the device settings.
	//  This corresponds to
	//  [Notification.silent](https://developer.mozilla.org/docs/Web/API/Notification/silent)
	//  DOM API.
	//  The default value is `FALSE`.
	IsSilent() bool // property IsSilent Getter
	// Timestamp
	//  Indicates the time at which a notification is created or applicable (past,
	//  present, or future) as the number of milliseconds since the UNIX epoch.
	Timestamp() types.TDateTime // property Timestamp Getter
	// VibrationPattern
	//  Gets the vibration pattern for devices with vibration hardware to emit.
	//  The vibration pattern can be represented by an array of 64-bit unsigned integers
	//  describing a pattern of vibrations and pauses. See [Vibration
	//  API](https://developer.mozilla.org/docs/Web/API/Vibration_API) for more
	//  information.
	//  This corresponds to
	//  [Notification.vibrate](https://developer.mozilla.org/docs/Web/API/Notification/vibrate)
	//  DOM API.
	//  An empty array is returned if no vibration patterns are
	//  specified.
	VibrationPattern() lcl.IUint64ArrayWrap // property VibrationPattern Getter
}

type TCoreWebView2Notification struct {
	TObject
}

func (m *TCoreWebView2Notification) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) ReportShown() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) ReportClicked() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) ReportClosed() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) BaseIntf() (result ICoreWebView2Notification) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NotificationAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Notification(resultPtr)
	return
}

func (m *TCoreWebView2Notification) SetBaseIntf(value ICoreWebView2Notification) {
	if !m.IsValid() {
		return
	}
	coreWebView2NotificationAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2Notification) Body() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) Direction() wvTypes.TWVTextDirectionKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2NotificationAPI().SysCallN(8, m.Instance())
	return wvTypes.TWVTextDirectionKind(r)
}

func (m *TCoreWebView2Notification) Language() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) Tag() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(10, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) IconUri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(11, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) BadgeUri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(13, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) BodyImageUri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NotificationAPI().SysCallN(14, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2Notification) ShouldRenotify() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) RequiresInteraction() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(16, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) IsSilent() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NotificationAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Notification) Timestamp() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2NotificationAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Notification) VibrationPattern() lcl.IUint64ArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := coreWebView2NotificationAPI().SysCallN(19, m.Instance())
	return lcl.AsUint64ArrayWrap(r)
}

// NewCoreWebView2Notification class constructor
func NewCoreWebView2Notification(baseIntf ICoreWebView2Notification) ICoreWebView2Notification {
	r := coreWebView2NotificationAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Notification(r)
}

var (
	coreWebView2NotificationOnce   base.Once
	coreWebView2NotificationImport *imports.Imports = nil
)

func coreWebView2NotificationAPI() *imports.Imports {
	coreWebView2NotificationOnce.Do(func() {
		coreWebView2NotificationImport = api.NewDefaultImports()
		coreWebView2NotificationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Notification_Create", 0), // constructor NewCoreWebView2Notification
			/* 1 */ imports.NewTable("TCoreWebView2Notification_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2Notification_ReportShown", 0), // function ReportShown
			/* 3 */ imports.NewTable("TCoreWebView2Notification_ReportClicked", 0), // function ReportClicked
			/* 4 */ imports.NewTable("TCoreWebView2Notification_ReportClosed", 0), // function ReportClosed
			/* 5 */ imports.NewTable("TCoreWebView2Notification_Initialized", 0), // property Initialized
			/* 6 */ imports.NewTable("TCoreWebView2Notification_BaseIntf", 0), // property BaseIntf
			/* 7 */ imports.NewTable("TCoreWebView2Notification_Body", 0), // property Body
			/* 8 */ imports.NewTable("TCoreWebView2Notification_Direction", 0), // property Direction
			/* 9 */ imports.NewTable("TCoreWebView2Notification_Language", 0), // property Language
			/* 10 */ imports.NewTable("TCoreWebView2Notification_Tag", 0), // property Tag
			/* 11 */ imports.NewTable("TCoreWebView2Notification_IconUri", 0), // property IconUri
			/* 12 */ imports.NewTable("TCoreWebView2Notification_Title", 0), // property Title
			/* 13 */ imports.NewTable("TCoreWebView2Notification_BadgeUri", 0), // property BadgeUri
			/* 14 */ imports.NewTable("TCoreWebView2Notification_BodyImageUri", 0), // property BodyImageUri
			/* 15 */ imports.NewTable("TCoreWebView2Notification_ShouldRenotify", 0), // property ShouldRenotify
			/* 16 */ imports.NewTable("TCoreWebView2Notification_RequiresInteraction", 0), // property RequiresInteraction
			/* 17 */ imports.NewTable("TCoreWebView2Notification_IsSilent", 0), // property IsSilent
			/* 18 */ imports.NewTable("TCoreWebView2Notification_Timestamp", 0), // property Timestamp
			/* 19 */ imports.NewTable("TCoreWebView2Notification_VibrationPattern", 0), // property VibrationPattern
		}
	})
	return coreWebView2NotificationImport
}
