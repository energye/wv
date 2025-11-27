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

// ICoreWebView2Controller Parent: lcl.IObject
type ICoreWebView2Controller interface {
	lcl.IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// MoveFocus
	//  Moves focus into WebView. WebView gets focus and focus is set to
	//  correspondent element in the page hosted in the WebView. For
	//  Programmatic reason, focus is set to previously focused element or the
	//  default element if no previously focused element exists. For `Next`
	//  reason, focus is set to the first element. For `Previous` reason, focus
	//  is set to the last element. WebView changes focus through user
	//  interaction including selecting into a WebView or Tab into it. For
	//  tabbing, the app runs MoveFocus with Next or Previous to align with Tab
	//  and Shift+Tab respectively when it decides the WebView is the next
	//  element that may exist in a tab. Or, the app runs `IsDialogMessage`
	//  as part of the associated message loop to allow the platform to auto
	//  handle tabbing. The platform rotates through all windows with
	//  `WS_TABSTOP`. When the WebView gets focus from `IsDialogMessage`, it is
	//  internally put the focus on the first or last element for tab and
	//  Shift+Tab respectively.
	MoveFocus(reason wvTypes.TWVMoveFocusReason) bool // function
	// Close
	//  Closes the WebView and cleans up the underlying browser instance.
	//  Cleaning up the browser instance releases the resources powering the
	//  WebView. The browser instance is shut down if no other WebViews are
	//  using it.
	//
	//  After running `Close`, most methods will fail and event handlers stop
	//  running. Specifically, the WebView releases the associated references to
	//  any associated event handlers when `Close` is run.
	//
	//  `Close` is implicitly run when the `CoreWebView2Controller` loses the
	//  final reference and is destructed. But it is best practice to
	//  explicitly run `Close` to avoid any accidental cycle of references
	//  between the WebView and the app code. Specifically, if you capture a
	//  reference to the WebView in an event handler you create a reference cycle
	//  between the WebView and the event handler. Run `Close` to break the
	//  cycle by releasing all event handlers. But to avoid the situation, it is
	//  best to both explicitly run `Close` on the WebView and to not capture a
	//  reference to the WebView to ensure the WebView is cleaned up correctly.
	//  `Close` is synchronous and won't trigger the `beforeunload` event.
	Close() bool // function
	// SetBoundsAndZoomFactor
	//  Updates `aBounds` and `aZoomFactor` properties at the same time. This
	//  operation is atomic from the perspective of the host. After returning
	//  from this function, the `aBounds` and `aZoomFactor` properties are both
	//  updated if the function is successful, or neither is updated if the
	//  function fails. If `aBounds` and `aZoomFactor` are both updated by the
	//  same scale (for example, `aBounds` and `aZoomFactor` are both doubled),
	//  then the page does not display a change in `window.innerWidth` or
	//  `window.innerHeight` and the WebView renders the content at the new size
	//  and zoom without intermediate renderings. This function also updates
	//  just one of `aZoomFactor` or `aBounds` by passing in the new value for one
	//  and the current value for the other.
	SetBoundsAndZoomFactor(bounds types.TRect, zoomFactor float64) bool // function
	// NotifyParentWindowPositionChanged
	//  This is a notification separate from `Bounds` that tells WebView that the
	//  main WebView parent (or any ancestor) `HWND` moved. This is needed
	//  for accessibility and certain dialogs in WebView to work correctly.
	NotifyParentWindowPositionChanged() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Controller // property BaseIntf Getter
	// ZoomFactor
	//  The zoom factor for the WebView.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</see>
	ZoomFactor() float64         // property ZoomFactor Getter
	SetZoomFactor(value float64) // property ZoomFactor Setter
	// IsVisible
	//  The `IsVisible` property determines whether to show or hide the WebView2.
	//  If `IsVisible` is set to `FALSE`, the WebView2 is transparent and is
	//  not rendered. However, this does not affect the window containing the
	//  WebView2 (the `HWND` parameter that was passed to
	//  `CreateCoreWebView2Controller`). If you want that window to disappear
	//  too, run `ShowWindow` on it directly in addition to modifying the
	//  `IsVisible` property. WebView2 as a child window does not get window
	//  messages when the top window is minimized or restored. For performance
	//  reasons, developers should set the `IsVisible` property of the WebView to
	//  `FALSE` when the app window is minimized and back to `TRUE` when the app
	//  window is restored. The app window does this by handling
	//  `SIZE_MINIMIZED and SIZE_RESTORED` command upon receiving `WM_SIZE`
	//  message.
	//  There are CPU and memory benefits when the page is hidden. For instance,
	//  Chromium has code that throttles activities on the page like animations
	//  and some tasks are run less frequently. Similarly, WebView2 will
	//  purge some caches to reduce memory usage.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_isvisible">See the ICoreWebView2Controller article.</see>
	IsVisible() bool         // property IsVisible Getter
	SetIsVisible(value bool) // property IsVisible Setter
	// Bounds
	//  The WebView bounds. Bounds are relative to the parent `HWND`. The app
	//  has two ways to position a WebView.
	//  * Create a child `HWND` that is the WebView parent `HWND`. Position
	//  the window where the WebView should be. Use `(0, 0)` for the
	//  top-left corner (the offset) of the `Bounds` of the WebView.
	//  * Use the top-most window of the app as the WebView parent HWND. For
	//  example, to position WebView correctly in the app, set the top-left
	//  corner of the Bound of the WebView.
	//  The values of `Bounds` are limited by the coordinate space of the host.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_bounds">See the ICoreWebView2Controller article.</see>
	Bounds() types.TRect         // property Bounds Getter
	SetBounds(value types.TRect) // property Bounds Setter
	// ParentWindow
	//  The parent window provided by the app that this WebView is using to
	//  render content. This API initially returns the window passed into
	//  `CreateCoreWebView2Controller`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_parentwindow">See the ICoreWebView2Controller article.</see>
	ParentWindow() types.THandle         // property ParentWindow Getter
	SetParentWindow(value types.THandle) // property ParentWindow Setter
	// DefaultBackgroundColor
	//  The `DefaultBackgroundColor` property is the color WebView renders
	//  underneath all web content. This means WebView renders this color when
	//  there is no web content loaded such as before the initial navigation or
	//  between navigations. This also means web pages with undefined css
	//  background properties or background properties containing transparent
	//  pixels will render their contents over this color. Web pages with defined
	//  and opaque background properties that span the page will obscure the
	//  `DefaultBackgroundColor` and display normally. The default value for this
	//  property is white to resemble the native browser experience.
	//  The Color is specified by the COREWEBVIEW2_COLOR that represents an RGBA
	//  value. The `A` represents an Alpha value, meaning
	//  `DefaultBackgroundColor` can be transparent. In the case of a transparent
	//  `DefaultBackgroundColor` WebView will render hosting app content as the
	//  background. This Alpha value is not supported on Windows 7. Any `A` value
	//  other than 255 will result in E_INVALIDARG on Windows 7.
	//  It is supported on all other WebView compatible platforms.
	//  Semi-transparent colors are not currently supported by this API and
	//  setting `DefaultBackgroundColor` to a semi-transparent color will fail
	//  with E_INVALIDARG. The only supported alpha values are 0 and 255, all
	//  other values will result in E_INVALIDARG.
	//  `DefaultBackgroundColor` can only be an opaque color or transparent.
	//  This value may also be set by using the
	//  `WEBVIEW2_DEFAULT_BACKGROUND_COLOR` environment variable. There is a
	//  known issue with background color where setting the color by API can
	//  still leave the app with a white flicker before the
	//  `DefaultBackgroundColor` takes effect. Setting the color via environment
	//  variable solves this issue. The value must be a hex value that can
	//  optionally prepend a 0x. The value must account for the alpha value
	//  which is represented by the first 2 digits. So any hex value fewer than 8
	//  digits will assume a prepended 00 to the hex value and result in a
	//  transparent color.
	//  `get_DefaultBackgroundColor` will return the result of this environment
	//  variable if used. This environment variable can only set the
	//  `DefaultBackgroundColor` once. Subsequent updates to background color
	//  must be done through API call.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller2#get_defaultbackgroundcolor">See the ICoreWebView2Controller2 article.</see>
	DefaultBackgroundColor() types.TColor         // property DefaultBackgroundColor Getter
	SetDefaultBackgroundColor(value types.TColor) // property DefaultBackgroundColor Setter
	// RasterizationScale
	//  The rasterization scale for the WebView. The rasterization scale is the
	//  combination of the monitor DPI scale and text scaling set by the user.
	//  This value should be updated when the DPI scale of the app's top level
	//  window changes (i.e. monitor DPI scale changes or window changes monitor)
	//  or when the text scale factor of the system changes.
	//  Rasterization scale applies to the WebView content, as well as
	//  popups, context menus, scroll bars, and so on. Normal app scaling
	//  scenarios should use the ZoomFactor property or SetBoundsAndZoomFactor
	//  API which only scale the rendered HTML content and not popups, context
	//  menus, scroll bars, and so on.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_rasterizationscale">See the ICoreWebView2Controller3 article.</see>
	RasterizationScale() float64         // property RasterizationScale Getter
	SetRasterizationScale(value float64) // property RasterizationScale Setter
	// ShouldDetectMonitorScaleChanges
	//  ShouldDetectMonitorScaleChanges property determines whether the WebView
	//  attempts to track monitor DPI scale changes. When true, the WebView will
	//  track monitor DPI scale changes, update the RasterizationScale property,
	//  and raises RasterizationScaleChanged event. When false, the WebView will
	//  not track monitor DPI scale changes, and the app must update the
	//  RasterizationScale property itself. RasterizationScaleChanged event will
	//  never raise when ShouldDetectMonitorScaleChanges is false. Apps that want
	//  to set their own rasterization scale should set this property to false to
	//  avoid the WebView2 updating the RasterizationScale property to match the
	//  monitor DPI scale.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_shoulddetectmonitorscalechanges">See the ICoreWebView2Controller3 article.</see>
	ShouldDetectMonitorScaleChanges() bool         // property ShouldDetectMonitorScaleChanges Getter
	SetShouldDetectMonitorScaleChanges(value bool) // property ShouldDetectMonitorScaleChanges Setter
	// BoundsMode
	//  BoundsMode affects how setting the Bounds and RasterizationScale
	//  properties work. Bounds mode can either be in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS
	//  mode or COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE mode.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS, setting the bounds
	//  property will set the size of the WebView in raw screen pixels. Changing
	//  the rasterization scale in this mode won't change the raw pixel size of
	//  the WebView and will only change the rasterization scale.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE, setting the
	//  bounds property will change the logical size of the WebView which can be
	//  described by the following equation: Logical size * rasterization scale = Raw Pixel size
	//  In this case, changing the rasterization scale will keep the logical size
	//  the same and change the raw pixel size.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_boundsmode">See the ICoreWebView2Controller3 article.</see>
	BoundsMode() wvTypes.TWVBoundsMode         // property BoundsMode Getter
	SetBoundsMode(value wvTypes.TWVBoundsMode) // property BoundsMode Setter
	// CoreWebView2
	//  Gets the `CoreWebView2` associated with this `CoreWebView2Controller`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_corewebview2">See the ICoreWebView2Controller article.</see>
	CoreWebView2() ICoreWebView2 // property CoreWebView2 Getter
	// AllowExternalDrop
	//  Gets the `AllowExternalDrop` property which is used to configure the
	//  capability that dragging objects from outside the bounds of webview2 and
	//  dropping into webview2 is allowed or disallowed. The default value is
	//  TRUE.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller4#get_allowexternaldrop">See the ICoreWebView2Controller4 article.</see>
	AllowExternalDrop() bool         // property AllowExternalDrop Getter
	SetAllowExternalDrop(value bool) // property AllowExternalDrop Setter
}

type TCoreWebView2Controller struct {
	lcl.TObject
}

func (m *TCoreWebView2Controller) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) MoveFocus(reason wvTypes.TWVMoveFocusReason) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(2, m.Instance(), uintptr(reason))
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) Close() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) SetBoundsAndZoomFactor(bounds types.TRect, zoomFactor float64) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&bounds)), uintptr(base.UnsafePointer(&zoomFactor)))
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) NotifyParentWindowPositionChanged() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) BaseIntf() (result ICoreWebView2Controller) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ControllerAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Controller(resultPtr)
	return
}

func (m *TCoreWebView2Controller) ZoomFactor() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Controller) SetZoomFactor(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2Controller) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) SetIsVisible(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Controller) Bounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Controller) SetBounds(value types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2Controller) ParentWindow() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ControllerAPI().SysCallN(11, 0, m.Instance())
	return types.THandle(r)
}

func (m *TCoreWebView2Controller) SetParentWindow(value types.THandle) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Controller) DefaultBackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ControllerAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCoreWebView2Controller) SetDefaultBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Controller) RasterizationScale() (result float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2Controller) SetRasterizationScale(value float64) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(13, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2Controller) ShouldDetectMonitorScaleChanges() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) SetShouldDetectMonitorScaleChanges(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2Controller) BoundsMode() wvTypes.TWVBoundsMode {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ControllerAPI().SysCallN(15, 0, m.Instance())
	return wvTypes.TWVBoundsMode(r)
}

func (m *TCoreWebView2Controller) SetBoundsMode(value wvTypes.TWVBoundsMode) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2Controller) CoreWebView2() (result ICoreWebView2) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ControllerAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2(resultPtr)
	return
}

func (m *TCoreWebView2Controller) AllowExternalDrop() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ControllerAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2Controller) SetAllowExternalDrop(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ControllerAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

// NewCoreWebView2Controller class constructor
func NewCoreWebView2Controller(baseIntf ICoreWebView2Controller) ICoreWebView2Controller {
	r := coreWebView2ControllerAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2Controller(r)
}

var (
	coreWebView2ControllerOnce   base.Once
	coreWebView2ControllerImport *imports.Imports = nil
)

func coreWebView2ControllerAPI() *imports.Imports {
	coreWebView2ControllerOnce.Do(func() {
		coreWebView2ControllerImport = api.NewDefaultImports()
		coreWebView2ControllerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2Controller_Create", 0), // constructor NewCoreWebView2Controller
			/* 1 */ imports.NewTable("TCoreWebView2Controller_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2Controller_MoveFocus", 0), // function MoveFocus
			/* 3 */ imports.NewTable("TCoreWebView2Controller_Close", 0), // function Close
			/* 4 */ imports.NewTable("TCoreWebView2Controller_SetBoundsAndZoomFactor", 0), // function SetBoundsAndZoomFactor
			/* 5 */ imports.NewTable("TCoreWebView2Controller_NotifyParentWindowPositionChanged", 0), // function NotifyParentWindowPositionChanged
			/* 6 */ imports.NewTable("TCoreWebView2Controller_Initialized", 0), // property Initialized
			/* 7 */ imports.NewTable("TCoreWebView2Controller_BaseIntf", 0), // property BaseIntf
			/* 8 */ imports.NewTable("TCoreWebView2Controller_ZoomFactor", 0), // property ZoomFactor
			/* 9 */ imports.NewTable("TCoreWebView2Controller_IsVisible", 0), // property IsVisible
			/* 10 */ imports.NewTable("TCoreWebView2Controller_Bounds", 0), // property Bounds
			/* 11 */ imports.NewTable("TCoreWebView2Controller_ParentWindow", 0), // property ParentWindow
			/* 12 */ imports.NewTable("TCoreWebView2Controller_DefaultBackgroundColor", 0), // property DefaultBackgroundColor
			/* 13 */ imports.NewTable("TCoreWebView2Controller_RasterizationScale", 0), // property RasterizationScale
			/* 14 */ imports.NewTable("TCoreWebView2Controller_ShouldDetectMonitorScaleChanges", 0), // property ShouldDetectMonitorScaleChanges
			/* 15 */ imports.NewTable("TCoreWebView2Controller_BoundsMode", 0), // property BoundsMode
			/* 16 */ imports.NewTable("TCoreWebView2Controller_CoreWebView2", 0), // property CoreWebView2
			/* 17 */ imports.NewTable("TCoreWebView2Controller_AllowExternalDrop", 0), // property AllowExternalDrop
		}
	})
	return coreWebView2ControllerImport
}
