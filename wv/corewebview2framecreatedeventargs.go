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

// ICoreWebView2FrameCreatedEventArgs Parent: IObject
//
//	Event args for the FrameCreated events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs">See the ICoreWebView2FrameCreatedEventArgs article.</a>
type ICoreWebView2FrameCreatedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameCreatedEventArgs // property
	// Frame
	//  The frame which was created.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs#get_frame">See the ICoreWebView2FrameCreatedEventArgs article.</a>
	Frame() ICoreWebView2Frame // property
}

// TCoreWebView2FrameCreatedEventArgs Parent: TObject
//
//	Event args for the FrameCreated events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs">See the ICoreWebView2FrameCreatedEventArgs article.</a>
type TCoreWebView2FrameCreatedEventArgs struct {
	TObject
}

func NewCoreWebView2FrameCreatedEventArgs(aArgs ICoreWebView2FrameCreatedEventArgs) ICoreWebView2FrameCreatedEventArgs {
	r1 := coreWebView2FrameCreatedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2FrameCreatedEventArgs(r1)
}

func (m *TCoreWebView2FrameCreatedEventArgs) Initialized() bool {
	r1 := coreWebView2FrameCreatedEventArgsImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameCreatedEventArgs) BaseIntf() ICoreWebView2FrameCreatedEventArgs {
	var resultCoreWebView2FrameCreatedEventArgs uintptr
	coreWebView2FrameCreatedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameCreatedEventArgs)))
	return AsCoreWebView2FrameCreatedEventArgs(resultCoreWebView2FrameCreatedEventArgs)
}

func (m *TCoreWebView2FrameCreatedEventArgs) Frame() ICoreWebView2Frame {
	var resultCoreWebView2Frame uintptr
	coreWebView2FrameCreatedEventArgsImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Frame)))
	return AsCoreWebView2Frame(resultCoreWebView2Frame)
}

var (
	coreWebView2FrameCreatedEventArgsImport       *imports.Imports = nil
	coreWebView2FrameCreatedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2FrameCreatedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2FrameCreatedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2FrameCreatedEventArgs_Frame", 0),
		/*3*/ imports.NewTable("CoreWebView2FrameCreatedEventArgs_Initialized", 0),
	}
)

func coreWebView2FrameCreatedEventArgsImportAPI() *imports.Imports {
	if coreWebView2FrameCreatedEventArgsImport == nil {
		coreWebView2FrameCreatedEventArgsImport = NewDefaultImports()
		coreWebView2FrameCreatedEventArgsImport.SetImportTable(coreWebView2FrameCreatedEventArgsImportTables)
		coreWebView2FrameCreatedEventArgsImportTables = nil
	}
	return coreWebView2FrameCreatedEventArgsImport
}
