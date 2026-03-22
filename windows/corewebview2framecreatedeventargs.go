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

// ICoreWebView2FrameCreatedEventArgs Parent: IObject
type ICoreWebView2FrameCreatedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameCreatedEventArgs // property BaseIntf Getter
	// Frame
	//  The frame which was created.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs#get_frame">See the ICoreWebView2FrameCreatedEventArgs article.</see>
	Frame() ICoreWebView2Frame // property Frame Getter
}

type TCoreWebView2FrameCreatedEventArgs struct {
	TObject
}

func (m *TCoreWebView2FrameCreatedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameCreatedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameCreatedEventArgs) BaseIntf() (result ICoreWebView2FrameCreatedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameCreatedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameCreatedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2FrameCreatedEventArgs) Frame() (result ICoreWebView2Frame) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameCreatedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Frame(resultPtr)
	return
}

// NewCoreWebView2FrameCreatedEventArgs class constructor
func NewCoreWebView2FrameCreatedEventArgs(args ICoreWebView2FrameCreatedEventArgs) ICoreWebView2FrameCreatedEventArgs {
	r := coreWebView2FrameCreatedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2FrameCreatedEventArgs(r)
}

var (
	coreWebView2FrameCreatedEventArgsOnce   base.Once
	coreWebView2FrameCreatedEventArgsImport *imports.Imports = nil
)

func coreWebView2FrameCreatedEventArgsAPI() *imports.Imports {
	coreWebView2FrameCreatedEventArgsOnce.Do(func() {
		coreWebView2FrameCreatedEventArgsImport = api.NewDefaultImports()
		coreWebView2FrameCreatedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2FrameCreatedEventArgs_Create", 0), // constructor NewCoreWebView2FrameCreatedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2FrameCreatedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2FrameCreatedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2FrameCreatedEventArgs_Frame", 0), // property Frame
		}
	})
	return coreWebView2FrameCreatedEventArgsImport
}
