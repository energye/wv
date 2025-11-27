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

// ICoreWebView2ScreenCaptureStartingEventArgs Parent: lcl.IObject
type ICoreWebView2ScreenCaptureStartingEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScreenCaptureStartingEventArgs // property BaseIntf Getter
	// Cancel
	//  The host may set this flag to cancel the screen capture. If canceled,
	//  the screen capture UI is not displayed regardless of the
	//  `Handled` property.
	//  On the script side, it will return with a NotAllowedError as Permission denied.
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// Handled
	//  By default, both the `ScreenCaptureStarting` event handlers on the
	//  `CoreWebView2Frame` and the `CoreWebView2` will be invoked, with the
	//  `CoreWebView2Frame` event handlers invoked first. The host may
	//  set this flag to `TRUE` within the `CoreWebView2Frame` event handlers
	//  to prevent the remaining `CoreWebView2` event handlers from being
	//  invoked. If the flag is set to `FALSE` within the `CoreWebView2Frame`
	//  event handlers, downstream handlers can update the `Cancel` property.
	//
	//  If a deferral is taken on the event args, then you must synchronously
	//  set `Handled` to TRUE prior to taking your deferral to prevent the
	//  `CoreWebView2`s event handlers from being invoked.
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// OriginalSourceFrameInfo
	//  The associated frame information that requests the screen capture
	//  permission. This can be used to get the frame source, name, frameId,
	//  and parent frame information.
	OriginalSourceFrameInfo() ICoreWebView2FrameInfo // property OriginalSourceFrameInfo Getter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this deferral to
	//  defer the decision to show the Screen Capture UI. getDisplayMedia()
	//  won't call its callbacks until the deferral is completed.
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2ScreenCaptureStartingEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) BaseIntf() (result ICoreWebView2ScreenCaptureStartingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ScreenCaptureStartingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) OriginalSourceFrameInfo() (result ICoreWebView2FrameInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfo(resultPtr)
	return
}

func (m *TCoreWebView2ScreenCaptureStartingEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2ScreenCaptureStartingEventArgs class constructor
func NewCoreWebView2ScreenCaptureStartingEventArgs(args ICoreWebView2ScreenCaptureStartingEventArgs) ICoreWebView2ScreenCaptureStartingEventArgs {
	r := coreWebView2ScreenCaptureStartingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ScreenCaptureStartingEventArgs(r)
}

var (
	coreWebView2ScreenCaptureStartingEventArgsOnce   base.Once
	coreWebView2ScreenCaptureStartingEventArgsImport *imports.Imports = nil
)

func coreWebView2ScreenCaptureStartingEventArgsAPI() *imports.Imports {
	coreWebView2ScreenCaptureStartingEventArgsOnce.Do(func() {
		coreWebView2ScreenCaptureStartingEventArgsImport = api.NewDefaultImports()
		coreWebView2ScreenCaptureStartingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_Create", 0), // constructor NewCoreWebView2ScreenCaptureStartingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_Cancel", 0), // property Cancel
			/* 4 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_Handled", 0), // property Handled
			/* 5 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_OriginalSourceFrameInfo", 0), // property OriginalSourceFrameInfo
			/* 6 */ imports.NewTable("TCoreWebView2ScreenCaptureStartingEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2ScreenCaptureStartingEventArgsImport
}
