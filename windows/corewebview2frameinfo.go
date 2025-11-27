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

// ICoreWebView2FrameInfo Parent: lcl.IObject
type ICoreWebView2FrameInfo interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfo         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2FrameInfo) // property BaseIntf Setter
	// Name
	//  The value of iframe's window.name property. The default value equals to
	//  iframe html tag declaring it, as in `<iframe name="frame-name">...</iframe>`.
	//  The returned string is empty when the frame has no name attribute and
	//  no assigned value for window.name.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo#get_name">See the ICoreWebView2FrameInfo article.</see>
	Name() string // property Name Getter
	// Source
	//  The URI of the document in the frame.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo#get_source">See the ICoreWebView2FrameInfo article.</see>
	Source() string // property Source Getter
	// ParentFrameInfo
	//  This parent frame's `FrameInfo`. `ParentFrameInfo` will only be
	//  populated when obtained via calling
	//  `ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed` will
	//  always have a `null` `ParentFrameInfo`. This property is also `null` for the
	//  main frame in the WebView2 which has no parent frame.
	//  Note that this `ParentFrameInfo` could be out of date as it's a snapshot.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_parentframeinfo">See the ICoreWebView2FrameInfo2 article.</see>
	ParentFrameInfo() ICoreWebView2FrameInfo // property ParentFrameInfo Getter
	// FrameId
	//  The unique identifier of the frame associated with the current `FrameInfo`.
	//  It's the same kind of ID as with the `FrameId` in `ICoreWebView2` and via
	//  `ICoreWebView2Frame`. `FrameId` will only be populated (non-zero) when obtained
	//  calling `CoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed` will
	//  always have an invalid frame Id 0.
	//  Note that this `FrameId` could be out of date as it's a snapshot.
	//  If there's WebView2 created or destroyed or `FrameCreated/FrameDestroyed` events
	//  after the asynchronous call `CoreWebView2Environment.GetProcessExtendedInfos`
	//  starts, you may want to call asynchronous method again to get the updated `FrameInfo`s.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_frameid">See the ICoreWebView2FrameInfo2 article.</see>
	FrameId() uint32 // property FrameId Getter
	// FrameKind
	//  The frame kind of the frame. `FrameKind` will only be populated when
	//  obtained calling `ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed`
	//  will always have the default value `COREWEBVIEW2_FRAME_KIND_UNKNOWN`.
	//  Note that this `FrameKind` could be out of date as it's a snapshot.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_framekind">See the ICoreWebView2FrameInfo2 article.</see>
	FrameKind() wvTypes.TWVFrameKind // property FrameKind Getter
	// FrameKindStr
	//  The frame kind of the frame in string format.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_framekind">See the ICoreWebView2FrameInfo2 article.</see>
	FrameKindStr() string // property FrameKindStr Getter
}

type TCoreWebView2FrameInfo struct {
	lcl.TObject
}

func (m *TCoreWebView2FrameInfo) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameInfoAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameInfo) BaseIntf() (result ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfo(resultPtr)
	return
}

func (m *TCoreWebView2FrameInfo) SetBaseIntf(value ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	coreWebView2FrameInfoAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2FrameInfo) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2FrameInfoAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2FrameInfo) Source() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2FrameInfoAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2FrameInfo) ParentFrameInfo() (result ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfo(resultPtr)
	return
}

func (m *TCoreWebView2FrameInfo) FrameId() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2FrameInfoAPI().SysCallN(6, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2FrameInfo) FrameKind() wvTypes.TWVFrameKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2FrameInfoAPI().SysCallN(7, m.Instance())
	return wvTypes.TWVFrameKind(r)
}

func (m *TCoreWebView2FrameInfo) FrameKindStr() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2FrameInfoAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

// NewCoreWebView2FrameInfo class constructor
func NewCoreWebView2FrameInfo(baseIntf ICoreWebView2FrameInfo) ICoreWebView2FrameInfo {
	r := coreWebView2FrameInfoAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2FrameInfo(r)
}

var (
	coreWebView2FrameInfoOnce   base.Once
	coreWebView2FrameInfoImport *imports.Imports = nil
)

func coreWebView2FrameInfoAPI() *imports.Imports {
	coreWebView2FrameInfoOnce.Do(func() {
		coreWebView2FrameInfoImport = api.NewDefaultImports()
		coreWebView2FrameInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2FrameInfo_Create", 0), // constructor NewCoreWebView2FrameInfo
			/* 1 */ imports.NewTable("TCoreWebView2FrameInfo_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2FrameInfo_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2FrameInfo_Name", 0), // property Name
			/* 4 */ imports.NewTable("TCoreWebView2FrameInfo_Source", 0), // property Source
			/* 5 */ imports.NewTable("TCoreWebView2FrameInfo_ParentFrameInfo", 0), // property ParentFrameInfo
			/* 6 */ imports.NewTable("TCoreWebView2FrameInfo_FrameId", 0), // property FrameId
			/* 7 */ imports.NewTable("TCoreWebView2FrameInfo_FrameKind", 0), // property FrameKind
			/* 8 */ imports.NewTable("TCoreWebView2FrameInfo_FrameKindStr", 0), // property FrameKindStr
		}
	})
	return coreWebView2FrameInfoImport
}
