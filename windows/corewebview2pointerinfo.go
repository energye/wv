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
	. "github.com/energye/lcl/types"
)

// ICoreWebView2PointerInfo Parent: IObject
//
//	This mostly represents a combined win32
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO object. It takes fields
//	from all three and excludes some win32 specific data types like HWND and
//	HANDLE. Note, sourceDevice is taken out but we expect the PointerDeviceRect
//	and DisplayRect to cover the existing use cases of sourceDevice.
//	Another big difference is that any of the point or rect locations are
//	expected to be in WebView physical coordinates. That is, coordinates
//	relative to the WebView and no DPI scaling applied.
//	The PointerId, PointerFlags, ButtonChangeKind, PenFlags, PenMask, TouchFlags,
//	and TouchMask are all #defined flags or enums in the
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO structure. We define those
//	properties here as UINT32 or INT32 and expect the developer to know how to
//	populate those values based on the Windows definitions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo">See the ICoreWebView2PointerInfo article.</a>
type ICoreWebView2PointerInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PointerInfo // property
	// PointerKind
	//  Get the PointerKind of the pointer event. This corresponds to the
	//  pointerKind property of the POINTER_INFO struct. The values are defined by
	//  the POINTER_INPUT_KIND enum in the Windows SDK(winuser.h). Supports
	//  PT_PEN and PT_TOUCH.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerkind">See the ICoreWebView2PointerInfo article.</a>
	PointerKind() uint32 // property
	// SetPointerKind Set PointerKind
	SetPointerKind(AValue uint32) // property
	// PointerId
	//  Get the PointerId of the pointer event. This corresponds to the pointerId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerid">See the ICoreWebView2PointerInfo article.</a>
	PointerId() uint32 // property
	// SetPointerId Set PointerId
	SetPointerId(AValue uint32) // property
	// FrameId
	//  Get the FrameID of the pointer event. This corresponds to the frameId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_frameid">See the ICoreWebView2PointerInfo article.</a>
	FrameId() uint32 // property
	// SetFrameId Set FrameId
	SetFrameId(AValue uint32) // property
	// PointerFlags
	//  Get the PointerFlags of the pointer event. This corresponds to the
	//  pointerFlags property of the POINTER_INFO struct. The values are defined
	//  by the POINTER_FLAGS constants in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerflags">See the ICoreWebView2PointerInfo article.</a>
	PointerFlags() uint32 // property
	// SetPointerFlags Set PointerFlags
	SetPointerFlags(AValue uint32) // property
	// PointerDeviceRect
	//  Get the PointerDeviceRect of the sourceDevice property of the
	//  POINTER_INFO struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerdevicerect">See the ICoreWebView2PointerInfo article.</a>
	PointerDeviceRect() (resultRect TRect) // property
	// SetPointerDeviceRect Set PointerDeviceRect
	SetPointerDeviceRect(AValue *TRect) // property
	// DisplayRect
	//  Get the DisplayRect of the sourceDevice property of the POINTER_INFO
	//  struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_displayrect">See the ICoreWebView2PointerInfo article.</a>
	DisplayRect() (resultRect TRect) // property
	// SetDisplayRect Set DisplayRect
	SetDisplayRect(AValue *TRect) // property
	// PixelLocation
	//  Get the PixelLocation of the pointer event. This corresponds to the
	//  ptPixelLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocation">See the ICoreWebView2PointerInfo article.</a>
	PixelLocation() (resultPoint TPoint) // property
	// SetPixelLocation Set PixelLocation
	SetPixelLocation(AValue *TPoint) // property
	// HimetricLocation
	//  Get the HimetricLocation of the pointer event. This corresponds to the
	//  ptHimetricLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocation">See the ICoreWebView2PointerInfo article.</a>
	HimetricLocation() (resultPoint TPoint) // property
	// SetHimetricLocation Set HimetricLocation
	SetHimetricLocation(AValue *TPoint) // property
	// PixelLocationRaw
	//  Get the PixelLocationRaw of the pointer event. This corresponds to the
	//  ptPixelLocationRaw property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocationraw">See the ICoreWebView2PointerInfo article.</a>
	PixelLocationRaw() (resultPoint TPoint) // property
	// SetPixelLocationRaw Set PixelLocationRaw
	SetPixelLocationRaw(AValue *TPoint) // property
	// HimetricLocationRaw
	//  Get the HimetricLocationRaw of the pointer event. This corresponds to the
	//  ptHimetricLocationRaw property of the POINTER_INFO struct as defined in
	//  the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocationraw">See the ICoreWebView2PointerInfo article.</a>
	HimetricLocationRaw() (resultPoint TPoint) // property
	// SetHimetricLocationRaw Set HimetricLocationRaw
	SetHimetricLocationRaw(AValue *TPoint) // property
	// Time
	//  Get the Time of the pointer event. This corresponds to the dwTime property
	//  of the POINTER_INFO struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_time">See the ICoreWebView2PointerInfo article.</a>
	Time() uint32 // property
	// SetTime Set Time
	SetTime(AValue uint32) // property
	// HistoryCount
	//  Get the HistoryCount of the pointer event. This corresponds to the
	//  historyCount property of the POINTER_INFO struct as defined in the Windows
	//  SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_historycount">See the ICoreWebView2PointerInfo article.</a>
	HistoryCount() uint32 // property
	// SetHistoryCount Set HistoryCount
	SetHistoryCount(AValue uint32) // property
	// InputData
	//  Get the InputData of the pointer event. This corresponds to the InputData
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_inputdata">See the ICoreWebView2PointerInfo article.</a>
	InputData() int32 // property
	// SetInputData Set InputData
	SetInputData(AValue int32) // property
	// KeyStates
	//  Get the KeyStates of the pointer event. This corresponds to the
	//  dwKeyStates property of the POINTER_INFO struct as defined in the Windows
	//  SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_keystates">See the ICoreWebView2PointerInfo article.</a>
	KeyStates() uint32 // property
	// SetKeyStates Set KeyStates
	SetKeyStates(AValue uint32) // property
	// PerformanceCount
	//  Get the PerformanceCount of the pointer event. This corresponds to the
	//  PerformanceCount property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_performancecount">See the ICoreWebView2PointerInfo article.</a>
	PerformanceCount() (resultUint64 uint64) // property
	// SetPerformanceCount Set PerformanceCount
	SetPerformanceCount(AValue uint64) // property
	// ButtonChangeKind
	//  Get the ButtonChangeKind of the pointer event. This corresponds to the
	//  ButtonChangeKind property of the POINTER_INFO struct. The values are
	//  defined by the POINTER_BUTTON_CHANGE_KIND enum in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_buttonchangekind">See the ICoreWebView2PointerInfo article.</a>
	ButtonChangeKind() int32 // property
	// SetButtonChangeKind Set ButtonChangeKind
	SetButtonChangeKind(AValue int32) // property
	// PenFlags
	//  Get the PenFlags of the pointer event. This corresponds to the penFlags
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_FLAGS constants in the Windows SDK(winuser.h).
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penflags">See the ICoreWebView2PointerInfo article.</a>
	PenFlags() uint32 // property
	// SetPenFlags Set PenFlags
	SetPenFlags(AValue uint32) // property
	// PenMask
	//  Get the PenMask of the pointer event. This corresponds to the penMask
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_MASK constants in the Windows SDK(winuser.h).
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penmask">See the ICoreWebView2PointerInfo article.</a>
	PenMask() uint32 // property
	// SetPenMask Set PenMask
	SetPenMask(AValue uint32) // property
	// PenPressure
	//  Get the PenPressure of the pointer event. This corresponds to the pressure
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penpressure">See the ICoreWebView2PointerInfo article.</a>
	PenPressure() uint32 // property
	// SetPenPressure Set PenPressure
	SetPenPressure(AValue uint32) // property
	// PenRotation
	//  Get the PenRotation of the pointer event. This corresponds to the rotation
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penrotation">See the ICoreWebView2PointerInfo article.</a>
	PenRotation() uint32 // property
	// SetPenRotation Set PenRotation
	SetPenRotation(AValue uint32) // property
	// PenTiltX
	//  Get the PenTiltX of the pointer event. This corresponds to the tiltX
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentiltx">See the ICoreWebView2PointerInfo article.</a>
	PenTiltX() int32 // property
	// SetPenTiltX Set PenTiltX
	SetPenTiltX(AValue int32) // property
	// PenTiltY
	//  Get the PenTiltY of the pointer event. This corresponds to the tiltY
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentilty">See the ICoreWebView2PointerInfo article.</a>
	PenTiltY() int32 // property
	// SetPenTiltY Set PenTiltY
	SetPenTiltY(AValue int32) // property
	// TouchFlags
	//  Get the TouchFlags of the pointer event. This corresponds to the
	//  touchFlags property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_FLAGS constants in the Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchflags">See the ICoreWebView2PointerInfo article.</a>
	TouchFlags() uint32 // property
	// SetTouchFlags Set TouchFlags
	SetTouchFlags(AValue uint32) // property
	// TouchMask
	//  Get the TouchMask of the pointer event. This corresponds to the
	//  touchMask property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_MASK constants in the Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchmask">See the ICoreWebView2PointerInfo article.</a>
	TouchMask() uint32 // property
	// SetTouchMask Set TouchMask
	SetTouchMask(AValue uint32) // property
	// TouchContact
	//  Get the TouchContact of the pointer event. This corresponds to the
	//  rcContact property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontact">See the ICoreWebView2PointerInfo article.</a>
	TouchContact() (resultRect TRect) // property
	// SetTouchContact Set TouchContact
	SetTouchContact(AValue *TRect) // property
	// TouchContactRaw
	//  Get the TouchContactRaw of the pointer event. This corresponds to the
	//  rcContactRaw property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontactraw">See the ICoreWebView2PointerInfo article.</a>
	TouchContactRaw() (resultRect TRect) // property
	// SetTouchContactRaw Set TouchContactRaw
	SetTouchContactRaw(AValue *TRect) // property
	// TouchOrientation
	//  Get the TouchOrientation of the pointer event. This corresponds to the
	//  orientation property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchorientation">See the ICoreWebView2PointerInfo article.</a>
	TouchOrientation() uint32 // property
	// SetTouchOrientation Set TouchOrientation
	SetTouchOrientation(AValue uint32) // property
	// TouchPressure
	//  Get the TouchPressure of the pointer event. This corresponds to the
	//  pressure property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchpressure">See the ICoreWebView2PointerInfo article.</a>
	TouchPressure() uint32 // property
	// SetTouchPressure Set TouchPressure
	SetTouchPressure(AValue uint32) // property
}

// TCoreWebView2PointerInfo Parent: TObject
//
//	This mostly represents a combined win32
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO object. It takes fields
//	from all three and excludes some win32 specific data types like HWND and
//	HANDLE. Note, sourceDevice is taken out but we expect the PointerDeviceRect
//	and DisplayRect to cover the existing use cases of sourceDevice.
//	Another big difference is that any of the point or rect locations are
//	expected to be in WebView physical coordinates. That is, coordinates
//	relative to the WebView and no DPI scaling applied.
//	The PointerId, PointerFlags, ButtonChangeKind, PenFlags, PenMask, TouchFlags,
//	and TouchMask are all #defined flags or enums in the
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO structure. We define those
//	properties here as UINT32 or INT32 and expect the developer to know how to
//	populate those values based on the Windows definitions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo">See the ICoreWebView2PointerInfo article.</a>
type TCoreWebView2PointerInfo struct {
	TObject
}

func NewCoreWebView2PointerInfo(aBaseIntf ICoreWebView2PointerInfo) ICoreWebView2PointerInfo {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PointerInfo(r1)
}

func (m *TCoreWebView2PointerInfo) Initialized() bool {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PointerInfo) BaseIntf() ICoreWebView2PointerInfo {
	var resultCoreWebView2PointerInfo uintptr
	coreWebView2PointerInfoImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PointerInfo)))
	return AsCoreWebView2PointerInfo(resultCoreWebView2PointerInfo)
}

func (m *TCoreWebView2PointerInfo) PointerKind() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerKind(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerId() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerId(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) FrameId() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetFrameId(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerFlags() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerFlags(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerDeviceRect() (resultRect TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(20, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPointerDeviceRect(AValue *TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(20, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) DisplayRect() (resultRect TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetDisplayRect(AValue *TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) PixelLocation() (resultPoint TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(18, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocation(AValue *TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(18, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocation() (resultPoint TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(5, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocation(AValue *TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(5, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) PixelLocationRaw() (resultPoint TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(19, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocationRaw(AValue *TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(19, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocationRaw() (resultPoint TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(6, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocationRaw(AValue *TPoint) {
	coreWebView2PointerInfoImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) Time() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTime(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) HistoryCount() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetHistoryCount(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) InputData() int32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetInputData(AValue int32) {
	coreWebView2PointerInfoImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) KeyStates() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetKeyStates(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PerformanceCount() (resultUint64 uint64) {
	coreWebView2PointerInfoImportAPI().SysCallN(17, 0, m.Instance(), uintptr(unsafePointer(&resultUint64)), uintptr(unsafePointer(&resultUint64)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPerformanceCount(AValue uint64) {
	coreWebView2PointerInfoImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) ButtonChangeKind() int32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetButtonChangeKind(AValue int32) {
	coreWebView2PointerInfoImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenFlags() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenFlags(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenMask() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenMask(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenPressure() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenPressure(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenRotation() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenRotation(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenTiltX() int32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltX(AValue int32) {
	coreWebView2PointerInfoImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenTiltY() int32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltY(AValue int32) {
	coreWebView2PointerInfoImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchFlags() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchFlags(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchMask() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchMask(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchContact() (resultRect TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(25, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContact(AValue *TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(25, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) TouchContactRaw() (resultRect TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(26, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContactRaw(AValue *TRect) {
	coreWebView2PointerInfoImportAPI().SysCallN(26, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) TouchOrientation() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchOrientation(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchPressure() uint32 {
	r1 := coreWebView2PointerInfoImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchPressure(AValue uint32) {
	coreWebView2PointerInfoImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

var (
	coreWebView2PointerInfoImport       *imports.Imports = nil
	coreWebView2PointerInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PointerInfo_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2PointerInfo_ButtonChangeKind", 0),
		/*2*/ imports.NewTable("CoreWebView2PointerInfo_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2PointerInfo_DisplayRect", 0),
		/*4*/ imports.NewTable("CoreWebView2PointerInfo_FrameId", 0),
		/*5*/ imports.NewTable("CoreWebView2PointerInfo_HimetricLocation", 0),
		/*6*/ imports.NewTable("CoreWebView2PointerInfo_HimetricLocationRaw", 0),
		/*7*/ imports.NewTable("CoreWebView2PointerInfo_HistoryCount", 0),
		/*8*/ imports.NewTable("CoreWebView2PointerInfo_Initialized", 0),
		/*9*/ imports.NewTable("CoreWebView2PointerInfo_InputData", 0),
		/*10*/ imports.NewTable("CoreWebView2PointerInfo_KeyStates", 0),
		/*11*/ imports.NewTable("CoreWebView2PointerInfo_PenFlags", 0),
		/*12*/ imports.NewTable("CoreWebView2PointerInfo_PenMask", 0),
		/*13*/ imports.NewTable("CoreWebView2PointerInfo_PenPressure", 0),
		/*14*/ imports.NewTable("CoreWebView2PointerInfo_PenRotation", 0),
		/*15*/ imports.NewTable("CoreWebView2PointerInfo_PenTiltX", 0),
		/*16*/ imports.NewTable("CoreWebView2PointerInfo_PenTiltY", 0),
		/*17*/ imports.NewTable("CoreWebView2PointerInfo_PerformanceCount", 0),
		/*18*/ imports.NewTable("CoreWebView2PointerInfo_PixelLocation", 0),
		/*19*/ imports.NewTable("CoreWebView2PointerInfo_PixelLocationRaw", 0),
		/*20*/ imports.NewTable("CoreWebView2PointerInfo_PointerDeviceRect", 0),
		/*21*/ imports.NewTable("CoreWebView2PointerInfo_PointerFlags", 0),
		/*22*/ imports.NewTable("CoreWebView2PointerInfo_PointerId", 0),
		/*23*/ imports.NewTable("CoreWebView2PointerInfo_PointerKind", 0),
		/*24*/ imports.NewTable("CoreWebView2PointerInfo_Time", 0),
		/*25*/ imports.NewTable("CoreWebView2PointerInfo_TouchContact", 0),
		/*26*/ imports.NewTable("CoreWebView2PointerInfo_TouchContactRaw", 0),
		/*27*/ imports.NewTable("CoreWebView2PointerInfo_TouchFlags", 0),
		/*28*/ imports.NewTable("CoreWebView2PointerInfo_TouchMask", 0),
		/*29*/ imports.NewTable("CoreWebView2PointerInfo_TouchOrientation", 0),
		/*30*/ imports.NewTable("CoreWebView2PointerInfo_TouchPressure", 0),
	}
)

func coreWebView2PointerInfoImportAPI() *imports.Imports {
	if coreWebView2PointerInfoImport == nil {
		coreWebView2PointerInfoImport = NewDefaultImports()
		coreWebView2PointerInfoImport.SetImportTable(coreWebView2PointerInfoImportTables)
		coreWebView2PointerInfoImportTables = nil
	}
	return coreWebView2PointerInfoImport
}
