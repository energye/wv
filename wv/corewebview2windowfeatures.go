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

// ICoreWebView2WindowFeatures Parent: IObject
//
//	The window features for a WebView popup window.  The fields match the
//	windowFeatures passed to window.open as specified in
//	[Window features](https://developer.mozilla.org/docs/Web/API/Window/open#Window_features)
//	on MDN.
//	There is no requirement for you to respect the values.  If your app does
//	not have corresponding UI features (for example, no toolbar) or if all
//	instance of WebView are opened in tabs and do not have distinct size or
//	positions, then your app does not respect the values.  You may want to
//	respect values, but perhaps only some apply to the UI of you app.
//	Accordingly, you may respect all, some, or none of the properties as
//	appropriate for your app.  For all numeric properties, if the value that is
//	passed to window.open is outside the range of an unsigned 32bit int, the
//	resulting value is the absolute value of the maximum for unsigned 32bit
//	integer.  If you are not able to parse the value an integer, it is
//	considered 0.  If the value is a floating point value, it is rounded down
//	to an integer.
//	In runtime versions 98 or later, the values of ShouldDisplayMenuBar,
//	ShouldDisplayStatus, ShouldDisplayToolbar, and ShouldDisplayScrollBars
//	will not directly depend on the equivalent fields in the windowFeatures
//	string.  Instead, they will all be false if the window is expected to be a
//	popup, and true if it is not.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures">See the ICoreWebView2WindowFeatures article.</a>
type ICoreWebView2WindowFeatures interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2WindowFeatures // property
	// HasPosition
	//  Specifies left and top values.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_hasposition">See the ICoreWebView2WindowFeatures article.</a>
	HasPosition() bool // property
	// HasSize
	//  Specifies height and width values.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_hassize">See the ICoreWebView2WindowFeatures article.</a>
	HasSize() bool // property
	// Left
	//  Specifies the left position of the window. If `HasPosition` is set to
	//  `FALSE`, this field is ignored.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_left">See the ICoreWebView2WindowFeatures article.</a>
	Left() uint32 // property
	// Top
	//  Specifies the top position of the window. If `HasPosition` is set to
	//  `FALSE`, this field is ignored.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_top">See the ICoreWebView2WindowFeatures article.</a>
	Top() uint32 // property
	// Width
	//  Specifies the width of the window. Minimum value is `100`. If `HasSize`
	//  is set to `FALSE`, this field is ignored.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_width">See the ICoreWebView2WindowFeatures article.</a>
	Width() uint32 // property
	// Height
	//  Specifies the height of the window. Minimum value is `100`. If
	//  `HasSize` is set to `FALSE`, this field is ignored.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_height">See the ICoreWebView2WindowFeatures article.</a>
	Height() uint32 // property
	// ShouldDisplayMenuBar
	//  Indicates that the menu bar is displayed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaymenubar">See the ICoreWebView2WindowFeatures article.</a>
	ShouldDisplayMenuBar() bool // property
	// ShouldDisplayStatus
	//  Indicates that the status bar is displayed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaystatus">See the ICoreWebView2WindowFeatures article.</a>
	ShouldDisplayStatus() bool // property
	// ShouldDisplayToolbar
	//  Indicates that the browser toolbar is displayed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplaytoolbar">See the ICoreWebView2WindowFeatures article.</a>
	ShouldDisplayToolbar() bool // property
	// ShouldDisplayScrollBars
	//  Indicates that the scroll bars are displayed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures#get_shoulddisplayscrollbars">See the ICoreWebView2WindowFeatures article.</a>
	ShouldDisplayScrollBars() bool                   // property
	CopyToRecord(aWindowFeatures *TWVWindowFeatures) // procedure
}

// TCoreWebView2WindowFeatures Parent: TObject
//
//	The window features for a WebView popup window.  The fields match the
//	windowFeatures passed to window.open as specified in
//	[Window features](https://developer.mozilla.org/docs/Web/API/Window/open#Window_features)
//	on MDN.
//	There is no requirement for you to respect the values.  If your app does
//	not have corresponding UI features (for example, no toolbar) or if all
//	instance of WebView are opened in tabs and do not have distinct size or
//	positions, then your app does not respect the values.  You may want to
//	respect values, but perhaps only some apply to the UI of you app.
//	Accordingly, you may respect all, some, or none of the properties as
//	appropriate for your app.  For all numeric properties, if the value that is
//	passed to window.open is outside the range of an unsigned 32bit int, the
//	resulting value is the absolute value of the maximum for unsigned 32bit
//	integer.  If you are not able to parse the value an integer, it is
//	considered 0.  If the value is a floating point value, it is rounded down
//	to an integer.
//	In runtime versions 98 or later, the values of ShouldDisplayMenuBar,
//	ShouldDisplayStatus, ShouldDisplayToolbar, and ShouldDisplayScrollBars
//	will not directly depend on the equivalent fields in the windowFeatures
//	string.  Instead, they will all be false if the window is expected to be a
//	popup, and true if it is not.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2windowfeatures">See the ICoreWebView2WindowFeatures article.</a>
type TCoreWebView2WindowFeatures struct {
	TObject
}

func NewCoreWebView2WindowFeatures(aBaseIntf ICoreWebView2WindowFeatures) ICoreWebView2WindowFeatures {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(2, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2WindowFeatures(r1)
}

func (m *TCoreWebView2WindowFeatures) Initialized() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) BaseIntf() ICoreWebView2WindowFeatures {
	var resultCoreWebView2WindowFeatures uintptr
	coreWebView2WindowFeaturesImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WindowFeatures)))
	return AsCoreWebView2WindowFeatures(resultCoreWebView2WindowFeatures)
}

func (m *TCoreWebView2WindowFeatures) HasPosition() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) HasSize() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) Left() uint32 {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(7, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2WindowFeatures) Top() uint32 {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(12, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2WindowFeatures) Width() uint32 {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(13, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2WindowFeatures) Height() uint32 {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(5, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayMenuBar() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayStatus() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayToolbar() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(11, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) ShouldDisplayScrollBars() bool {
	r1 := coreWebView2WindowFeaturesImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2WindowFeatures) CopyToRecord(aWindowFeatures *TWVWindowFeatures) {
	var result0 uintptr
	coreWebView2WindowFeaturesImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&result0)))
	*aWindowFeatures = *(*TWVWindowFeatures)(unsafePointer(result0))
}

var (
	coreWebView2WindowFeaturesImport       *imports.Imports = nil
	coreWebView2WindowFeaturesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2WindowFeatures_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2WindowFeatures_CopyToRecord", 0),
		/*2*/ imports.NewTable("CoreWebView2WindowFeatures_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2WindowFeatures_HasPosition", 0),
		/*4*/ imports.NewTable("CoreWebView2WindowFeatures_HasSize", 0),
		/*5*/ imports.NewTable("CoreWebView2WindowFeatures_Height", 0),
		/*6*/ imports.NewTable("CoreWebView2WindowFeatures_Initialized", 0),
		/*7*/ imports.NewTable("CoreWebView2WindowFeatures_Left", 0),
		/*8*/ imports.NewTable("CoreWebView2WindowFeatures_ShouldDisplayMenuBar", 0),
		/*9*/ imports.NewTable("CoreWebView2WindowFeatures_ShouldDisplayScrollBars", 0),
		/*10*/ imports.NewTable("CoreWebView2WindowFeatures_ShouldDisplayStatus", 0),
		/*11*/ imports.NewTable("CoreWebView2WindowFeatures_ShouldDisplayToolbar", 0),
		/*12*/ imports.NewTable("CoreWebView2WindowFeatures_Top", 0),
		/*13*/ imports.NewTable("CoreWebView2WindowFeatures_Width", 0),
	}
)

func coreWebView2WindowFeaturesImportAPI() *imports.Imports {
	if coreWebView2WindowFeaturesImport == nil {
		coreWebView2WindowFeaturesImport = NewDefaultImports()
		coreWebView2WindowFeaturesImport.SetImportTable(coreWebView2WindowFeaturesImportTables)
		coreWebView2WindowFeaturesImportTables = nil
	}
	return coreWebView2WindowFeaturesImport
}
