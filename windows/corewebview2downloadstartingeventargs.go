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

// ICoreWebView2DownloadStartingEventArgs Parent: IObject
type ICoreWebView2DownloadStartingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DownloadStartingEventArgs // property BaseIntf Getter
	// DownloadOperation
	//  Returns the `ICoreWebView2DownloadOperation` for the download that
	//  has started.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_downloadoperation">See the ICoreWebView2DownloadStartingEventArgs article.</see>
	DownloadOperation() ICoreWebView2DownloadOperation // property DownloadOperation Getter
	// Cancel
	//  The host may set this flag to cancel the download. If canceled, the
	//  download save dialog is not displayed regardless of the
	//  `Handled` property.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_cancel">See the ICoreWebView2DownloadStartingEventArgs article.</see>
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// ResultFilePath
	//  The path to the file. If setting the path, the host should ensure that it
	//  is an absolute path, including the file name, and that the path does not
	//  point to an existing file. If the path points to an existing file, the
	//  file will be overwritten. If the directory does not exist, it is created.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_resultfilepath">See the ICoreWebView2DownloadStartingEventArgs article.</see>
	ResultFilePath() string         // property ResultFilePath Getter
	SetResultFilePath(value string) // property ResultFilePath Setter
	// Handled
	//  The host may set this flag to `TRUE` to hide the default download dialog
	//  for this download. The download will progress as normal if it is not
	//  canceled, there will just be no default UI shown. By default the value is
	//  `FALSE` and the default download dialog is shown.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_handled">See the ICoreWebView2DownloadStartingEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#getdeferral">See the ICoreWebView2DownloadStartingEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2DownloadStartingEventArgs struct {
	TObject
}

func (m *TCoreWebView2DownloadStartingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadStartingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadStartingEventArgs) BaseIntf() (result ICoreWebView2DownloadStartingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2DownloadStartingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2DownloadStartingEventArgs) DownloadOperation() (result ICoreWebView2DownloadOperation) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2DownloadOperation(resultPtr)
	return
}

func (m *TCoreWebView2DownloadStartingEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadStartingEventArgsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2DownloadStartingEventArgs) ResultFilePath() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2DownloadStartingEventArgsAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetResultFilePath(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2DownloadStartingEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadStartingEventArgsAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2DownloadStartingEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DownloadStartingEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2DownloadStartingEventArgs class constructor
func NewCoreWebView2DownloadStartingEventArgs(args ICoreWebView2DownloadStartingEventArgs) ICoreWebView2DownloadStartingEventArgs {
	r := coreWebView2DownloadStartingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2DownloadStartingEventArgs(r)
}

var (
	coreWebView2DownloadStartingEventArgsOnce   base.Once
	coreWebView2DownloadStartingEventArgsImport *imports.Imports = nil
)

func coreWebView2DownloadStartingEventArgsAPI() *imports.Imports {
	coreWebView2DownloadStartingEventArgsOnce.Do(func() {
		coreWebView2DownloadStartingEventArgsImport = api.NewDefaultImports()
		coreWebView2DownloadStartingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_Create", 0), // constructor NewCoreWebView2DownloadStartingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_DownloadOperation", 0), // property DownloadOperation
			/* 4 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_Cancel", 0), // property Cancel
			/* 5 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_ResultFilePath", 0), // property ResultFilePath
			/* 6 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_Handled", 0), // property Handled
			/* 7 */ imports.NewTable("TCoreWebView2DownloadStartingEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2DownloadStartingEventArgsImport
}
