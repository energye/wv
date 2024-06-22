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

// ICoreWebView2FrameInfo Parent: IObject
//
//	Provides a set of properties for a frame in the ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo">See the ICoreWebView2FrameInfo article.</a>
type ICoreWebView2FrameInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfo // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2FrameInfo) // property
	// Name
	//  The name attribute of the frame, as in `<iframe name="frame-name" ...>`.
	//  The returned string is empty when the frame has no name attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo#get_name">See the ICoreWebView2FrameInfo article.</a>
	Name() string // property
	// Source
	//  The URI of the document in the frame.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo#get_source">See the ICoreWebView2FrameInfo article.</a>
	Source() string // property
	// ParentFrameInfo
	//  This parent frame's `FrameInfo`. `ParentFrameInfo` will only be
	//  populated when obtained via calling
	//  `ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed` will
	//  always have a `null` `ParentFrameInfo`. This property is also `null` for the
	//  main frame in the WebView2 which has no parent frame.
	//  Note that this `ParentFrameInfo` could be out of date as it's a snapshot.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_parentframeinfo">See the ICoreWebView2FrameInfo2 article.</a>
	ParentFrameInfo() ICoreWebView2FrameInfo // property
	// FrameId
	//  The unique identifier of the frame associated with the current `FrameInfo`.
	//  It's the same kind of ID as with the `FrameId` in `ICoreWebView2` and via
	//  `ICoreWebView2Frame`. `FrameId` will only be populated(non-zero) when obtained
	//  calling `CoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed` will
	//  always have an invalid frame Id 0.
	//  Note that this `FrameId` could be out of date as it's a snapshot.
	//  If there's WebView2 created or destroyed or `FrameCreated/FrameDestroyed` events
	//  after the asynchronous call `CoreWebView2Environment.GetProcessExtendedInfos`
	//  starts, you may want to call asynchronous method again to get the updated `FrameInfo`s.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_frameid">See the ICoreWebView2FrameInfo2 article.</a>
	FrameId() uint32 // property
	// FrameKind
	//  The frame kind of the frame. `FrameKind` will only be populated when
	//  obtained calling `ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos`.
	//  `ICoreWebView2FrameInfo` objects obtained via `ICoreWebView2.ProcessFailed`
	//  will always have the default value `COREWEBVIEW2_FRAME_KIND_UNKNOWN`.
	//  Note that this `FrameKind` could be out of date as it's a snapshot.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_framekind">See the ICoreWebView2FrameInfo2 article.</a>
	FrameKind() TWVFrameKind // property
	// FrameKindStr
	//  The frame kind of the frame in string format.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2#get_framekind">See the ICoreWebView2FrameInfo2 article.</a>
	FrameKindStr() string // property
}

// TCoreWebView2FrameInfo Parent: TObject
//
//	Provides a set of properties for a frame in the ICoreWebView2.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo">See the ICoreWebView2FrameInfo article.</a>
type TCoreWebView2FrameInfo struct {
	TObject
}

func NewCoreWebView2FrameInfo(aBaseIntf ICoreWebView2FrameInfo) ICoreWebView2FrameInfo {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2FrameInfo(r1)
}

func (m *TCoreWebView2FrameInfo) Initialized() bool {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameInfo) BaseIntf() ICoreWebView2FrameInfo {
	var resultCoreWebView2FrameInfo uintptr
	coreWebView2FrameInfoImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2FrameInfo)))
	return AsCoreWebView2FrameInfo(resultCoreWebView2FrameInfo)
}

func (m *TCoreWebView2FrameInfo) SetBaseIntf(AValue ICoreWebView2FrameInfo) {
	coreWebView2FrameInfoImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2FrameInfo) Name() string {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2FrameInfo) Source() string {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2FrameInfo) ParentFrameInfo() ICoreWebView2FrameInfo {
	var resultCoreWebView2FrameInfo uintptr
	coreWebView2FrameInfoImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfo)))
	return AsCoreWebView2FrameInfo(resultCoreWebView2FrameInfo)
}

func (m *TCoreWebView2FrameInfo) FrameId() uint32 {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(2, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2FrameInfo) FrameKind() TWVFrameKind {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(3, m.Instance())
	return TWVFrameKind(r1)
}

func (m *TCoreWebView2FrameInfo) FrameKindStr() string {
	r1 := coreWebView2FrameInfoImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

var (
	coreWebView2FrameInfoImport       *imports.Imports = nil
	coreWebView2FrameInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2FrameInfo_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2FrameInfo_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2FrameInfo_FrameId", 0),
		/*3*/ imports.NewTable("CoreWebView2FrameInfo_FrameKind", 0),
		/*4*/ imports.NewTable("CoreWebView2FrameInfo_FrameKindStr", 0),
		/*5*/ imports.NewTable("CoreWebView2FrameInfo_Initialized", 0),
		/*6*/ imports.NewTable("CoreWebView2FrameInfo_Name", 0),
		/*7*/ imports.NewTable("CoreWebView2FrameInfo_ParentFrameInfo", 0),
		/*8*/ imports.NewTable("CoreWebView2FrameInfo_Source", 0),
	}
)

func coreWebView2FrameInfoImportAPI() *imports.Imports {
	if coreWebView2FrameInfoImport == nil {
		coreWebView2FrameInfoImport = NewDefaultImports()
		coreWebView2FrameInfoImport.SetImportTable(coreWebView2FrameInfoImportTables)
		coreWebView2FrameInfoImportTables = nil
	}
	return coreWebView2FrameInfoImport
}
