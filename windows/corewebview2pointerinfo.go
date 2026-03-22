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
)

// ICoreWebView2PointerInfo Parent: IObject
type ICoreWebView2PointerInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PointerInfo // property BaseIntf Getter
	// PointerKind
	//  Get the PointerKind of the pointer event. This corresponds to the
	//  pointerKind property of the POINTER_INFO struct. The values are defined by
	//  the POINTER_INPUT_KIND enum in the Windows SDK (winuser.h). Supports
	//  PT_PEN and PT_TOUCH.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerkind">See the ICoreWebView2PointerInfo article.</see>
	PointerKind() uint32         // property PointerKind Getter
	SetPointerKind(value uint32) // property PointerKind Setter
	// PointerId
	//  Get the PointerId of the pointer event. This corresponds to the pointerId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerid">See the ICoreWebView2PointerInfo article.</see>
	PointerId() uint32         // property PointerId Getter
	SetPointerId(value uint32) // property PointerId Setter
	// FrameId
	//  Get the FrameID of the pointer event. This corresponds to the frameId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_frameid">See the ICoreWebView2PointerInfo article.</see>
	FrameId() uint32         // property FrameId Getter
	SetFrameId(value uint32) // property FrameId Setter
	// PointerFlags
	//  Get the PointerFlags of the pointer event. This corresponds to the
	//  pointerFlags property of the POINTER_INFO struct. The values are defined
	//  by the POINTER_FLAGS constants in the Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerflags">See the ICoreWebView2PointerInfo article.</see>
	PointerFlags() uint32         // property PointerFlags Getter
	SetPointerFlags(value uint32) // property PointerFlags Setter
	// PointerDeviceRect
	//  Get the PointerDeviceRect of the sourceDevice property of the
	//  POINTER_INFO struct as defined in the Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerdevicerect">See the ICoreWebView2PointerInfo article.</see>
	PointerDeviceRect() types.TRect         // property PointerDeviceRect Getter
	SetPointerDeviceRect(value types.TRect) // property PointerDeviceRect Setter
	// DisplayRect
	//  Get the DisplayRect of the sourceDevice property of the POINTER_INFO
	//  struct as defined in the Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_displayrect">See the ICoreWebView2PointerInfo article.</see>
	DisplayRect() types.TRect         // property DisplayRect Getter
	SetDisplayRect(value types.TRect) // property DisplayRect Setter
	// PixelLocation
	//  Get the PixelLocation of the pointer event. This corresponds to the
	//  ptPixelLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocation">See the ICoreWebView2PointerInfo article.</see>
	PixelLocation() types.TPoint         // property PixelLocation Getter
	SetPixelLocation(value types.TPoint) // property PixelLocation Setter
	// HimetricLocation
	//  Get the HimetricLocation of the pointer event. This corresponds to the
	//  ptHimetricLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocation">See the ICoreWebView2PointerInfo article.</see>
	HimetricLocation() types.TPoint         // property HimetricLocation Getter
	SetHimetricLocation(value types.TPoint) // property HimetricLocation Setter
	// PixelLocationRaw
	//  Get the PixelLocationRaw of the pointer event. This corresponds to the
	//  ptPixelLocationRaw property of the POINTER_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocationraw">See the ICoreWebView2PointerInfo article.</see>
	PixelLocationRaw() types.TPoint         // property PixelLocationRaw Getter
	SetPixelLocationRaw(value types.TPoint) // property PixelLocationRaw Setter
	// HimetricLocationRaw
	//  Get the HimetricLocationRaw of the pointer event. This corresponds to the
	//  ptHimetricLocationRaw property of the POINTER_INFO struct as defined in
	//  the Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocationraw">See the ICoreWebView2PointerInfo article.</see>
	HimetricLocationRaw() types.TPoint         // property HimetricLocationRaw Getter
	SetHimetricLocationRaw(value types.TPoint) // property HimetricLocationRaw Setter
	// Time
	//  Get the Time of the pointer event. This corresponds to the dwTime property
	//  of the POINTER_INFO struct as defined in the Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_time">See the ICoreWebView2PointerInfo article.</see>
	Time() uint32         // property Time Getter
	SetTime(value uint32) // property Time Setter
	// HistoryCount
	//  Get the HistoryCount of the pointer event. This corresponds to the
	//  historyCount property of the POINTER_INFO struct as defined in the Windows
	//  SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_historycount">See the ICoreWebView2PointerInfo article.</see>
	HistoryCount() uint32         // property HistoryCount Getter
	SetHistoryCount(value uint32) // property HistoryCount Setter
	// InputData
	//  Get the InputData of the pointer event. This corresponds to the InputData
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_inputdata">See the ICoreWebView2PointerInfo article.</see>
	InputData() int32         // property InputData Getter
	SetInputData(value int32) // property InputData Setter
	// KeyStates
	//  Get the KeyStates of the pointer event. This corresponds to the
	//  dwKeyStates property of the POINTER_INFO struct as defined in the Windows
	//  SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_keystates">See the ICoreWebView2PointerInfo article.</see>
	KeyStates() uint32         // property KeyStates Getter
	SetKeyStates(value uint32) // property KeyStates Setter
	// PerformanceCount
	//  Get the PerformanceCount of the pointer event. This corresponds to the
	//  PerformanceCount property of the POINTER_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_performancecount">See the ICoreWebView2PointerInfo article.</see>
	PerformanceCount() uint64         // property PerformanceCount Getter
	SetPerformanceCount(value uint64) // property PerformanceCount Setter
	// ButtonChangeKind
	//  Get the ButtonChangeKind of the pointer event. This corresponds to the
	//  ButtonChangeKind property of the POINTER_INFO struct. The values are
	//  defined by the POINTER_BUTTON_CHANGE_KIND enum in the Windows SDK
	//  (winuser.h).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_buttonchangekind">See the ICoreWebView2PointerInfo article.</see>
	ButtonChangeKind() int32         // property ButtonChangeKind Getter
	SetButtonChangeKind(value int32) // property ButtonChangeKind Setter
	// PenFlags
	//  Get the PenFlags of the pointer event. This corresponds to the penFlags
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_FLAGS constants in the Windows SDK (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penflags">See the ICoreWebView2PointerInfo article.</see>
	PenFlags() uint32         // property PenFlags Getter
	SetPenFlags(value uint32) // property PenFlags Setter
	// PenMask
	//  Get the PenMask of the pointer event. This corresponds to the penMask
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_MASK constants in the Windows SDK (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penmask">See the ICoreWebView2PointerInfo article.</see>
	PenMask() uint32         // property PenMask Getter
	SetPenMask(value uint32) // property PenMask Setter
	// PenPressure
	//  Get the PenPressure of the pointer event. This corresponds to the pressure
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penpressure">See the ICoreWebView2PointerInfo article.</see>
	PenPressure() uint32         // property PenPressure Getter
	SetPenPressure(value uint32) // property PenPressure Setter
	// PenRotation
	//  Get the PenRotation of the pointer event. This corresponds to the rotation
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penrotation">See the ICoreWebView2PointerInfo article.</see>
	PenRotation() uint32         // property PenRotation Getter
	SetPenRotation(value uint32) // property PenRotation Setter
	// PenTiltX
	//  Get the PenTiltX of the pointer event. This corresponds to the tiltX
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentiltx">See the ICoreWebView2PointerInfo article.</see>
	PenTiltX() int32         // property PenTiltX Getter
	SetPenTiltX(value int32) // property PenTiltX Setter
	// PenTiltY
	//  Get the PenTiltY of the pointer event. This corresponds to the tiltY
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  (winuser.h).
	//  This is a Pen specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentilty">See the ICoreWebView2PointerInfo article.</see>
	PenTiltY() int32         // property PenTiltY Getter
	SetPenTiltY(value int32) // property PenTiltY Setter
	// TouchFlags
	//  Get the TouchFlags of the pointer event. This corresponds to the
	//  touchFlags property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_FLAGS constants in the Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchflags">See the ICoreWebView2PointerInfo article.</see>
	TouchFlags() uint32         // property TouchFlags Getter
	SetTouchFlags(value uint32) // property TouchFlags Setter
	// TouchMask
	//  Get the TouchMask of the pointer event. This corresponds to the
	//  touchMask property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_MASK constants in the Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchmask">See the ICoreWebView2PointerInfo article.</see>
	TouchMask() uint32         // property TouchMask Getter
	SetTouchMask(value uint32) // property TouchMask Setter
	// TouchContact
	//  Get the TouchContact of the pointer event. This corresponds to the
	//  rcContact property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontact">See the ICoreWebView2PointerInfo article.</see>
	TouchContact() types.TRect         // property TouchContact Getter
	SetTouchContact(value types.TRect) // property TouchContact Setter
	// TouchContactRaw
	//  Get the TouchContactRaw of the pointer event. This corresponds to the
	//  rcContactRaw property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontactraw">See the ICoreWebView2PointerInfo article.</see>
	TouchContactRaw() types.TRect         // property TouchContactRaw Getter
	SetTouchContactRaw(value types.TRect) // property TouchContactRaw Setter
	// TouchOrientation
	//  Get the TouchOrientation of the pointer event. This corresponds to the
	//  orientation property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchorientation">See the ICoreWebView2PointerInfo article.</see>
	TouchOrientation() uint32         // property TouchOrientation Getter
	SetTouchOrientation(value uint32) // property TouchOrientation Setter
	// TouchPressure
	//  Get the TouchPressure of the pointer event. This corresponds to the
	//  pressure property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK (winuser.h).
	//  This is a Touch specific attribute.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchpressure">See the ICoreWebView2PointerInfo article.</see>
	TouchPressure() uint32         // property TouchPressure Getter
	SetTouchPressure(value uint32) // property TouchPressure Setter
}

type TCoreWebView2PointerInfo struct {
	TObject
}

func (m *TCoreWebView2PointerInfo) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PointerInfoAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PointerInfo) BaseIntf() (result ICoreWebView2PointerInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PointerInfoAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PointerInfo(resultPtr)
	return
}

func (m *TCoreWebView2PointerInfo) PointerKind() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(3, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPointerKind(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PointerId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(4, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPointerId(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) FrameId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(5, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetFrameId(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PointerFlags() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(6, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPointerFlags(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PointerDeviceRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPointerDeviceRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) DisplayRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetDisplayRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) PixelLocation() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocation(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(9, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocation() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocation(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) PixelLocationRaw() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocationRaw(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocationRaw() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocationRaw(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(12, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) Time() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(13, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetTime(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) HistoryCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(14, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetHistoryCount(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) InputData() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PointerInfo) SetInputData(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) KeyStates() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(16, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetKeyStates(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PerformanceCount() (result uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(17, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPerformanceCount(value uint64) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(17, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) ButtonChangeKind() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PointerInfo) SetButtonChangeKind(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenFlags() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(19, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenFlags(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenMask() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(20, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenMask(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenPressure() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(21, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenPressure(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenRotation() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(22, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenRotation(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenTiltX() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltX(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) PenTiltY() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltY(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) TouchFlags() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(25, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetTouchFlags(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) TouchMask() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(26, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetTouchMask(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) TouchContact() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(27, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContact(value types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(27, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) TouchContactRaw() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(28, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContactRaw(value types.TRect) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(28, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCoreWebView2PointerInfo) TouchOrientation() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(29, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetTouchOrientation(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2PointerInfo) TouchPressure() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PointerInfoAPI().SysCallN(30, 0, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2PointerInfo) SetTouchPressure(value uint32) {
	if !m.IsValid() {
		return
	}
	coreWebView2PointerInfoAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

// NewCoreWebView2PointerInfo class constructor
func NewCoreWebView2PointerInfo(baseIntf ICoreWebView2PointerInfo) ICoreWebView2PointerInfo {
	r := coreWebView2PointerInfoAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2PointerInfo(r)
}

var (
	coreWebView2PointerInfoOnce   base.Once
	coreWebView2PointerInfoImport *imports.Imports = nil
)

func coreWebView2PointerInfoAPI() *imports.Imports {
	coreWebView2PointerInfoOnce.Do(func() {
		coreWebView2PointerInfoImport = api.NewDefaultImports()
		coreWebView2PointerInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PointerInfo_Create", 0), // constructor NewCoreWebView2PointerInfo
			/* 1 */ imports.NewTable("TCoreWebView2PointerInfo_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2PointerInfo_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2PointerInfo_PointerKind", 0), // property PointerKind
			/* 4 */ imports.NewTable("TCoreWebView2PointerInfo_PointerId", 0), // property PointerId
			/* 5 */ imports.NewTable("TCoreWebView2PointerInfo_FrameId", 0), // property FrameId
			/* 6 */ imports.NewTable("TCoreWebView2PointerInfo_PointerFlags", 0), // property PointerFlags
			/* 7 */ imports.NewTable("TCoreWebView2PointerInfo_PointerDeviceRect", 0), // property PointerDeviceRect
			/* 8 */ imports.NewTable("TCoreWebView2PointerInfo_DisplayRect", 0), // property DisplayRect
			/* 9 */ imports.NewTable("TCoreWebView2PointerInfo_PixelLocation", 0), // property PixelLocation
			/* 10 */ imports.NewTable("TCoreWebView2PointerInfo_HimetricLocation", 0), // property HimetricLocation
			/* 11 */ imports.NewTable("TCoreWebView2PointerInfo_PixelLocationRaw", 0), // property PixelLocationRaw
			/* 12 */ imports.NewTable("TCoreWebView2PointerInfo_HimetricLocationRaw", 0), // property HimetricLocationRaw
			/* 13 */ imports.NewTable("TCoreWebView2PointerInfo_Time", 0), // property Time
			/* 14 */ imports.NewTable("TCoreWebView2PointerInfo_HistoryCount", 0), // property HistoryCount
			/* 15 */ imports.NewTable("TCoreWebView2PointerInfo_InputData", 0), // property InputData
			/* 16 */ imports.NewTable("TCoreWebView2PointerInfo_KeyStates", 0), // property KeyStates
			/* 17 */ imports.NewTable("TCoreWebView2PointerInfo_PerformanceCount", 0), // property PerformanceCount
			/* 18 */ imports.NewTable("TCoreWebView2PointerInfo_ButtonChangeKind", 0), // property ButtonChangeKind
			/* 19 */ imports.NewTable("TCoreWebView2PointerInfo_PenFlags", 0), // property PenFlags
			/* 20 */ imports.NewTable("TCoreWebView2PointerInfo_PenMask", 0), // property PenMask
			/* 21 */ imports.NewTable("TCoreWebView2PointerInfo_PenPressure", 0), // property PenPressure
			/* 22 */ imports.NewTable("TCoreWebView2PointerInfo_PenRotation", 0), // property PenRotation
			/* 23 */ imports.NewTable("TCoreWebView2PointerInfo_PenTiltX", 0), // property PenTiltX
			/* 24 */ imports.NewTable("TCoreWebView2PointerInfo_PenTiltY", 0), // property PenTiltY
			/* 25 */ imports.NewTable("TCoreWebView2PointerInfo_TouchFlags", 0), // property TouchFlags
			/* 26 */ imports.NewTable("TCoreWebView2PointerInfo_TouchMask", 0), // property TouchMask
			/* 27 */ imports.NewTable("TCoreWebView2PointerInfo_TouchContact", 0), // property TouchContact
			/* 28 */ imports.NewTable("TCoreWebView2PointerInfo_TouchContactRaw", 0), // property TouchContactRaw
			/* 29 */ imports.NewTable("TCoreWebView2PointerInfo_TouchOrientation", 0), // property TouchOrientation
			/* 30 */ imports.NewTable("TCoreWebView2PointerInfo_TouchPressure", 0), // property TouchPressure
		}
	})
	return coreWebView2PointerInfoImport
}
