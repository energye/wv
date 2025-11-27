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
)

// ICoreWebView2NewWindowRequestedEventArgs Parent: lcl.IObject
type ICoreWebView2NewWindowRequestedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NewWindowRequestedEventArgs // property BaseIntf Getter
	// URI
	//  The target uri of the new window requested.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_uri">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	URI() string // property URI Getter
	// NewWindow
	//  Gets the new window.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_newwindow">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	NewWindow() ICoreWebView2         // property NewWindow Getter
	SetNewWindow(value ICoreWebView2) // property NewWindow Setter
	// Handled
	//  Gets whether the `NewWindowRequested` event is handled by host.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_handled">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// IsUserInitiated
	//  `TRUE` when the new window request was initiated through a user gesture.
	//  Examples of user initiated requests are:
	//  - Selecting an anchor tag with target
	//  - Programmatic window open from a script that directly run as a result of
	//  user interaction such as via onclick handlers.
	//  Non-user initiated requests are programmatic window opens from a script
	//  that are not directly triggered by user interaction, such as those that
	//  run while loading a new page or via timers.
	//  The Microsoft Edge popup blocker is disabled for WebView so the app is
	//  able to use this flag to block non-user initiated popups.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_isuserinitiated">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	IsUserInitiated() bool // property IsUserInitiated Getter
	// Deferral
	//  Obtain an `ICoreWebView2Deferral` object and put the event into a
	//  deferred state. Use the `ICoreWebView2Deferral` object to complete the
	//  window open request at a later time. While this event is deferred the
	//  opener window returns a `WindowProxy` to an un-navigated window, which
	//  navigates when the deferral is complete.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#getdeferral">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
	// WindowFeatures
	//  Window features specified by the `window.open`. The features should be
	//  considered for positioning and sizing of new webview windows.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_windowfeatures">See the ICoreWebView2NewWindowRequestedEventArgs article.</see>
	WindowFeatures() ICoreWebView2WindowFeatures // property WindowFeatures Getter
	// Name
	//  Gets the name of the new window. This window can be created via `window.open(url, windowName)`,
	//  where the windowName parameter corresponds to `Name` property.
	//  If no windowName is passed to `window.open`, then the `Name` property
	//  will be set to an empty string. Additionally, if window is opened through other means,
	//  such as `<a target="windowName">...</a>` or `<iframe name="windowName">...</iframe>`,
	//  then the `Name` property will be set accordingly. In the case of target=_blank,
	//  the `Name` property will be an empty string.
	//  Opening a window via ctrl+clicking a link would result in the `Name` property
	//  being set to an empty string.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs2#get_name">See the ICoreWebView2NewWindowRequestedEventArgs2 article.</see>
	Name() string // property Name Getter
	// OriginalSourceFrameInfo
	//  The frame info of the frame where the new window request originated. The
	//  `OriginalSourceFrameInfo` is a snapshot of frame information at the time when the
	//  new window was requested. See `ICoreWebView2FrameInfo` for details on frame
	//  properties.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs3#get_originalsourceframeinfo">See the ICoreWebView2NewWindowRequestedEventArgs3 article.</see>
	OriginalSourceFrameInfo() ICoreWebView2FrameInfo // property OriginalSourceFrameInfo Getter
}

type TCoreWebView2NewWindowRequestedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) BaseIntf() (result ICoreWebView2NewWindowRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2NewWindowRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) URI() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) NewWindow() (result ICoreWebView2) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2(resultPtr)
	return
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) SetNewWindow(value ICoreWebView2) {
	if !m.IsValid() {
		return
	}
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) IsUserInitiated() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) WindowFeatures() (result ICoreWebView2WindowFeatures) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2WindowFeatures(resultPtr)
	return
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(9, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) OriginalSourceFrameInfo() (result ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfo(resultPtr)
	return
}

// NewCoreWebView2NewWindowRequestedEventArgs class constructor
func NewCoreWebView2NewWindowRequestedEventArgs(args ICoreWebView2NewWindowRequestedEventArgs) ICoreWebView2NewWindowRequestedEventArgs {
	r := coreWebView2NewWindowRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2NewWindowRequestedEventArgs(r)
}

var (
	coreWebView2NewWindowRequestedEventArgsOnce   base.Once
	coreWebView2NewWindowRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2NewWindowRequestedEventArgsAPI() *imports.Imports {
	coreWebView2NewWindowRequestedEventArgsOnce.Do(func() {
		coreWebView2NewWindowRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2NewWindowRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_Create", 0), // constructor NewCoreWebView2NewWindowRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_URI", 0), // property URI
			/* 4 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_NewWindow", 0), // property NewWindow
			/* 5 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_Handled", 0), // property Handled
			/* 6 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_IsUserInitiated", 0), // property IsUserInitiated
			/* 7 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_Deferral", 0), // property Deferral
			/* 8 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_WindowFeatures", 0), // property WindowFeatures
			/* 9 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_Name", 0), // property Name
			/* 10 */ imports.NewTable("TCoreWebView2NewWindowRequestedEventArgs_OriginalSourceFrameInfo", 0), // property OriginalSourceFrameInfo
		}
	})
	return coreWebView2NewWindowRequestedEventArgsImport
}
