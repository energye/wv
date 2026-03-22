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
)

// ICoreWebView2WindowFeatures Parent: IObject
type ICoreWebView2WindowFeatures interface {
	IObject
	CopyToRecord(windowFeatures *TWVWindowFeatures) // procedure
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WindowFeatures // property BaseIntf Getter
	// HasPosition
	//  Specifies left and top values.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_hasposition">See the ICoreWebView2WindowFeatures article.</see>
	HasPosition() bool // property HasPosition Getter
	// HasSize
	//  Specifies height and width values.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_hassize">See the ICoreWebView2WindowFeatures article.</see>
	HasSize() bool // property HasSize Getter
	// Left
	//  Specifies the left position of the window. If `HasPosition` is set to
	//  `FALSE`, this field is ignored.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_left">See the ICoreWebView2WindowFeatures article.</see>
	Left() uint32 // property Left Getter
	// Top
	//  Specifies the top position of the window. If `HasPosition` is set to
	//  `FALSE`, this field is ignored.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_top">See the ICoreWebView2WindowFeatures article.</see>
	Top() uint32 // property Top Getter
	// Width
	//  Specifies the width of the window. Minimum value is `100`. If `HasSize`
	//  is set to `FALSE`, this field is ignored.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_width">See the ICoreWebView2WindowFeatures article.</see>
	Width() uint32 // property Width Getter
	// Height
	//  Specifies the height of the window. Minimum value is `100`. If
	//  `HasSize` is set to `FALSE`, this field is ignored.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_height">See the ICoreWebView2WindowFeatures article.</see>
	Height() uint32 // property Height Getter
	// ShouldDisplayMenuBar
	//  Indicates that the menu bar is displayed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaymenubar">See the ICoreWebView2WindowFeatures article.</see>
	ShouldDisplayMenuBar() bool // property ShouldDisplayMenuBar Getter
	// ShouldDisplayStatus
	//  Indicates that the status bar is displayed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaystatus">See the ICoreWebView2WindowFeatures article.</see>
	ShouldDisplayStatus() bool // property ShouldDisplayStatus Getter
	// ShouldDisplayToolbar
	//  Indicates that the browser toolbar is displayed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaytoolbar">See the ICoreWebView2WindowFeatures article.</see>
	ShouldDisplayToolbar() bool // property ShouldDisplayToolbar Getter
	// ShouldDisplayScrollBars
	//  Indicates that the scroll bars are displayed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplayscrollbars">See the ICoreWebView2WindowFeatures article.</see>
	ShouldDisplayScrollBars() bool // property ShouldDisplayScrollBars Getter
}

type TCoreWebView2WindowFeatures struct {
	TObject
}

func (m *TCoreWebView2WindowFeatures) CopyToRecord(windowFeatures *TWVWindowFeatures) {
	if !m.IsValid() {
		return
	}
	coreWebView2WindowFeaturesAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(windowFeatures)))
}

func (m *TCoreWebView2WindowFeatures) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) BaseIntf() (result ICoreWebView2WindowFeatures) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2WindowFeaturesAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WindowFeatures(resultPtr)
	return
}

func (m *TCoreWebView2WindowFeatures) HasPosition() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) HasSize() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) Left() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(6, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2WindowFeatures) Top() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(7, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2WindowFeatures) Width() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(8, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2WindowFeatures) Height() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(9, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayMenuBar() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayStatus() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayToolbar() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayScrollBars() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2WindowFeaturesAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2WindowFeatures class constructor
func NewCoreWebView2WindowFeatures(baseIntf ICoreWebView2WindowFeatures) ICoreWebView2WindowFeatures {
	r := coreWebView2WindowFeaturesAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2WindowFeatures(r)
}

var (
	coreWebView2WindowFeaturesOnce   base.Once
	coreWebView2WindowFeaturesImport *imports.Imports = nil
)

func coreWebView2WindowFeaturesAPI() *imports.Imports {
	coreWebView2WindowFeaturesOnce.Do(func() {
		coreWebView2WindowFeaturesImport = api.NewDefaultImports()
		coreWebView2WindowFeaturesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WindowFeatures_Create", 0), // constructor NewCoreWebView2WindowFeatures
			/* 1 */ imports.NewTable("TCoreWebView2WindowFeatures_CopyToRecord", 0), // procedure CopyToRecord
			/* 2 */ imports.NewTable("TCoreWebView2WindowFeatures_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2WindowFeatures_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2WindowFeatures_HasPosition", 0), // property HasPosition
			/* 5 */ imports.NewTable("TCoreWebView2WindowFeatures_HasSize", 0), // property HasSize
			/* 6 */ imports.NewTable("TCoreWebView2WindowFeatures_Left", 0), // property Left
			/* 7 */ imports.NewTable("TCoreWebView2WindowFeatures_Top", 0), // property Top
			/* 8 */ imports.NewTable("TCoreWebView2WindowFeatures_Width", 0), // property Width
			/* 9 */ imports.NewTable("TCoreWebView2WindowFeatures_Height", 0), // property Height
			/* 10 */ imports.NewTable("TCoreWebView2WindowFeatures_ShouldDisplayMenuBar", 0), // property ShouldDisplayMenuBar
			/* 11 */ imports.NewTable("TCoreWebView2WindowFeatures_ShouldDisplayStatus", 0), // property ShouldDisplayStatus
			/* 12 */ imports.NewTable("TCoreWebView2WindowFeatures_ShouldDisplayToolbar", 0), // property ShouldDisplayToolbar
			/* 13 */ imports.NewTable("TCoreWebView2WindowFeatures_ShouldDisplayScrollBars", 0), // property ShouldDisplayScrollBars
		}
	})
	return coreWebView2WindowFeaturesImport
}
