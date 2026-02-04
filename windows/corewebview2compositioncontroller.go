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

// ICoreWebView2CompositionController Parent: lcl.IObject
type ICoreWebView2CompositionController interface {
	lcl.IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// SendMouseInput
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL, then mouseData specifies the amount of
	//  wheel movement. A positive value indicates that the wheel was rotated
	//  forward, away from the user; a negative value indicates that the wheel was
	//  rotated backward, toward the user. One wheel click is defined as
	//  WHEEL_DELTA, which is 120.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN, or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP, then mouseData specifies which X
	//  buttons were pressed or released. This value should be 1 if the first X
	//  button is pressed/released and 2 if the second X button is
	//  pressed/released.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE, then virtualKeys,
	//  mouseData, and point should all be zero.
	//  If eventKind is any other value, then mouseData should be zero.
	//  Point is expected to be in the client coordinate space of the WebView.
	//  To track mouse events that start in the WebView and can potentially move
	//  outside of the WebView and host application, calling SetCapture and
	//  ReleaseCapture is recommended.
	//  To dismiss hover popups, it is also recommended to send
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE messages.
	SendMouseInput(eventKind wvTypes.TWVMouseEventKind, virtualKeys wvTypes.TWVMouseEventVirtualKeys, mouseData uint32, point types.TPoint) bool // function
	// SendPointerInput
	//  SendPointerInput accepts touch or pen pointer input of types defined in
	//  COREWEBVIEW2_POINTER_EVENT_KIND. Any pointer input from the system must be
	//  converted into an ICoreWebView2PointerInfo first.
	SendPointerInput(eventKind wvTypes.TWVPointerEventKind, pointerInfo ICoreWebView2PointerInfo) bool // function
	// DragLeave
	//  This function corresponds to [IDropTarget::DragLeave](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave).
	//
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragLeave calls to this function.
	DragLeave() types.HRESULT // function
	// DragOver
	//  This function corresponds to [IDropTarget::DragOver](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover).
	//
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragOver calls to this function.
	//
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates (Similar to how SendMouseInput works).
	DragOver(keyState uint32, point types.TPoint, outEffect *uint32) types.HRESULT // function
	// GetNonClientRegionAtPoint
	//  If you are hosting a WebView2 using CoreWebView2CompositionController, you can call
	//  this method in your Win32 WndProc to determine if the mouse is moving over or
	//  clicking on WebView2 web content that should be considered part of a non-client region.
	//  The point parameter is expected to be in the client coordinate space of WebView2.
	//  The method sets the out parameter value as follows:
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_CAPTION when point corresponds to
	//  a region (HTML element) within the WebView2 with
	//  `-webkit-app-region: drag` CSS style set.
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_CLIENT when point corresponds to
	//  a region (HTML element) within the WebView2 without
	//  `-webkit-app-region: drag` CSS style set.
	//  - COREWEBVIEW2_NON_CLIENT_REGION_KIND_NOWHERE when point is not within the WebView2.
	//
	//  NOTE: in order for WebView2 to properly handle the title bar system menu,
	//  the app needs to send WM_NCRBUTTONDOWN and WM_NCRBUTTONUP to SendMouseInput.
	//  See sample code below.
	//  \snippet ViewComponent.cpp DraggableRegions2
	//
	//  \snippet ViewComponent.cpp DraggableRegions1
	GetNonClientRegionAtPoint(point types.TPoint) wvTypes.TWVNonClientRegionKind // function
	// QueryNonClientRegion
	//  This method is used to get the collection of rects that correspond
	//  to a particular COREWEBVIEW2_NON_CLIENT_REGION_KIND. This is to be used in
	//  the callback of add_NonClientRegionChanged whose event args object contains
	//  a region property of type COREWEBVIEW2_NON_CLIENT_REGION_KIND.
	//
	//  \snippet ScenarioNonClientRegionSupport.cpp AddChangeListener
	QueryNonClientRegion(kind wvTypes.TWVNonClientRegionKind) ICoreWebView2RegionRectCollectionView // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CompositionController // property BaseIntf Getter
	// RootVisualTarget
	//  The RootVisualTarget is a visual in the hosting app's visual tree. This
	//  visual is where the WebView will connect its visual tree. The app uses
	//  this visual to position the WebView within the app. The app still needs
	//  to use the Bounds property to size the WebView. The RootVisualTarget
	//  property can be an IDCompositionVisual or a
	//  Windows::UI::Composition::ContainerVisual. WebView will connect its visual
	//  tree to the provided visual before returning from the property setter. The
	//  app needs to commit on its device setting the RootVisualTarget property.
	//  The RootVisualTarget property supports being set to nullptr to disconnect
	//  the WebView from the app's visual tree.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_rootvisualtarget">See the ICoreWebView2CompositionController article.</see>
	RootVisualTarget() lcl.IUnknown         // property RootVisualTarget Getter
	SetRootVisualTarget(value lcl.IUnknown) // property RootVisualTarget Setter
	// Cursor
	//  The current cursor that WebView thinks it should be. The cursor should be
	//  set in WM_SETCURSOR through \::SetCursor or set on the corresponding
	//  parent/ancestor HWND of the WebView through \::SetClassLongPtr. The HCURSOR
	//  can be freed so CopyCursor/DestroyCursor is recommended to keep your own
	//  copy if you are doing more than immediately setting the cursor.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_cursor">See the ICoreWebView2CompositionController article.</see>
	Cursor() types.HCURSOR // property Cursor Getter
	// SystemCursorID
	//  The current system cursor ID reported by the underlying rendering engine
	//  for WebView. For example, most of the time, when the cursor is over text,
	//  this will return the int value for IDC_IBEAM. The systemCursorId is only
	//  valid if the rendering engine reports a default Windows cursor resource
	//  value. Navigate to
	//  [LoadCursorW](/windows/win32/api/winuser/nf-winuser-loadcursorw) for more
	//  details. Otherwise, if custom CSS cursors are being used, this will return
	//  0. To actually use systemCursorId in LoadCursor or LoadImage,
	//  MAKEINTRESOURCE must be called on it first.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_systemcursorid">See the ICoreWebView2CompositionController article.</see>
	SystemCursorID() uint32 // property SystemCursorID Getter
	// AutomationProvider
	//  Returns the Automation Provider for the WebView. This object implements
	//  IRawElementProviderSimple.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller2#get_automationprovider">See the ICoreWebView2CompositionController2 article.</see>
	AutomationProvider() lcl.IUnknown // property AutomationProvider Getter
}

type TCoreWebView2CompositionController struct {
	lcl.TObject
}

func (m *TCoreWebView2CompositionController) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2CompositionController) SendMouseInput(eventKind wvTypes.TWVMouseEventKind, virtualKeys wvTypes.TWVMouseEventVirtualKeys, mouseData uint32, point types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(2, m.Instance(), uintptr(eventKind), uintptr(virtualKeys), uintptr(mouseData), uintptr(base.UnsafePointer(&point)))
	return api.GoBool(r)
}

func (m *TCoreWebView2CompositionController) SendPointerInput(eventKind wvTypes.TWVPointerEventKind, pointerInfo ICoreWebView2PointerInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(3, m.Instance(), uintptr(eventKind), base.GetObjectUintptr(pointerInfo))
	return api.GoBool(r)
}

func (m *TCoreWebView2CompositionController) DragLeave() types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(4, m.Instance())
	return types.HRESULT(r)
}

func (m *TCoreWebView2CompositionController) DragOver(keyState uint32, point types.TPoint, outEffect *uint32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var effectPtr uintptr
	r := coreWebView2CompositionControllerAPI().SysCallN(5, m.Instance(), uintptr(keyState), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&effectPtr)))
	*outEffect = uint32(effectPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2CompositionController) GetNonClientRegionAtPoint(point types.TPoint) wvTypes.TWVNonClientRegionKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return wvTypes.TWVNonClientRegionKind(r)
}

func (m *TCoreWebView2CompositionController) QueryNonClientRegion(kind wvTypes.TWVNonClientRegionKind) (result ICoreWebView2RegionRectCollectionView) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CompositionControllerAPI().SysCallN(7, m.Instance(), uintptr(kind), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2RegionRectCollectionView(resultPtr)
	return
}

func (m *TCoreWebView2CompositionController) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2CompositionController) BaseIntf() (result ICoreWebView2CompositionController) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CompositionControllerAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2CompositionController(resultPtr)
	return
}

func (m *TCoreWebView2CompositionController) RootVisualTarget() (result lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CompositionControllerAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = lcl.AsUnknown(resultPtr)
	return
}

func (m *TCoreWebView2CompositionController) SetRootVisualTarget(value lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	coreWebView2CompositionControllerAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2CompositionController) Cursor() types.HCURSOR {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(11, m.Instance())
	return types.HCURSOR(r)
}

func (m *TCoreWebView2CompositionController) SystemCursorID() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CompositionControllerAPI().SysCallN(12, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2CompositionController) AutomationProvider() (result lcl.IUnknown) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2CompositionControllerAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = lcl.AsUnknown(resultPtr)
	return
}

// NewCoreWebView2CompositionController class constructor
func NewCoreWebView2CompositionController(baseIntf ICoreWebView2CompositionController) ICoreWebView2CompositionController {
	r := coreWebView2CompositionControllerAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2CompositionController(r)
}

var (
	coreWebView2CompositionControllerOnce   base.Once
	coreWebView2CompositionControllerImport *imports.Imports = nil
)

func coreWebView2CompositionControllerAPI() *imports.Imports {
	coreWebView2CompositionControllerOnce.Do(func() {
		coreWebView2CompositionControllerImport = api.NewDefaultImports()
		coreWebView2CompositionControllerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2CompositionController_Create", 0), // constructor NewCoreWebView2CompositionController
			/* 1 */ imports.NewTable("TCoreWebView2CompositionController_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2CompositionController_SendMouseInput", 0), // function SendMouseInput
			/* 3 */ imports.NewTable("TCoreWebView2CompositionController_SendPointerInput", 0), // function SendPointerInput
			/* 4 */ imports.NewTable("TCoreWebView2CompositionController_DragLeave", 0), // function DragLeave
			/* 5 */ imports.NewTable("TCoreWebView2CompositionController_DragOver", 0), // function DragOver
			/* 6 */ imports.NewTable("TCoreWebView2CompositionController_GetNonClientRegionAtPoint", 0), // function GetNonClientRegionAtPoint
			/* 7 */ imports.NewTable("TCoreWebView2CompositionController_QueryNonClientRegion", 0), // function QueryNonClientRegion
			/* 8 */ imports.NewTable("TCoreWebView2CompositionController_Initialized", 0), // property Initialized
			/* 9 */ imports.NewTable("TCoreWebView2CompositionController_BaseIntf", 0), // property BaseIntf
			/* 10 */ imports.NewTable("TCoreWebView2CompositionController_RootVisualTarget", 0), // property RootVisualTarget
			/* 11 */ imports.NewTable("TCoreWebView2CompositionController_Cursor", 0), // property Cursor
			/* 12 */ imports.NewTable("TCoreWebView2CompositionController_SystemCursorID", 0), // property SystemCursorID
			/* 13 */ imports.NewTable("TCoreWebView2CompositionController_AutomationProvider", 0), // property AutomationProvider
		}
	})
	return coreWebView2CompositionControllerImport
}
