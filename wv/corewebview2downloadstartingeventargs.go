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

// ICoreWebView2DownloadStartingEventArgs Parent: IObject
//
//	Event args for the DownloadStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs">See the ICoreWebView2DownloadStartingEventArgs article.</a>
type ICoreWebView2DownloadStartingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DownloadStartingEventArgs // property
	// DownloadOperation
	//  Returns the `ICoreWebView2DownloadOperation` for the download that
	//  has started.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_downloadoperation">See the ICoreWebView2DownloadStartingEventArgs article.</a>
	DownloadOperation() ICoreWebView2DownloadOperation // property
	// Cancel
	//  The host may set this flag to cancel the download. If canceled, the
	//  download save dialog is not displayed regardless of the
	//  `Handled` property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_cancel">See the ICoreWebView2DownloadStartingEventArgs article.</a>
	Cancel() bool // property
	// SetCancel Set Cancel
	SetCancel(AValue bool) // property
	// ResultFilePath
	//  The path to the file. If setting the path, the host should ensure that it
	//  is an absolute path, including the file name, and that the path does not
	//  point to an existing file. If the path points to an existing file, the
	//  file will be overwritten. If the directory does not exist, it is created.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_resultfilepath">See the ICoreWebView2DownloadStartingEventArgs article.</a>
	ResultFilePath() string // property
	// SetResultFilePath Set ResultFilePath
	SetResultFilePath(AValue string) // property
	// Handled
	//  The host may set this flag to `TRUE` to hide the default download dialog
	//  for this download. The download will progress as normal if it is not
	//  canceled, there will just be no default UI shown. By default the value is
	//  `FALSE` and the default download dialog is shown.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#get_handled">See the ICoreWebView2DownloadStartingEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs#getdeferral">See the ICoreWebView2DownloadStartingEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
}

// TCoreWebView2DownloadStartingEventArgs Parent: TObject
//
//	Event args for the DownloadStarting event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadstartingeventargs">See the ICoreWebView2DownloadStartingEventArgs article.</a>
type TCoreWebView2DownloadStartingEventArgs struct {
	TObject
}

func NewCoreWebView2DownloadStartingEventArgs(aArgs ICoreWebView2DownloadStartingEventArgs) ICoreWebView2DownloadStartingEventArgs {
	r1 := coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(2, GetObjectUintptr(aArgs))
	return AsCoreWebView2DownloadStartingEventArgs(r1)
}

func (m *TCoreWebView2DownloadStartingEventArgs) Initialized() bool {
	r1 := coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadStartingEventArgs) BaseIntf() ICoreWebView2DownloadStartingEventArgs {
	var resultCoreWebView2DownloadStartingEventArgs uintptr
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DownloadStartingEventArgs)))
	return AsCoreWebView2DownloadStartingEventArgs(resultCoreWebView2DownloadStartingEventArgs)
}

func (m *TCoreWebView2DownloadStartingEventArgs) DownloadOperation() ICoreWebView2DownloadOperation {
	var resultCoreWebView2DownloadOperation uintptr
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DownloadOperation)))
	return AsCoreWebView2DownloadOperation(resultCoreWebView2DownloadOperation)
}

func (m *TCoreWebView2DownloadStartingEventArgs) Cancel() bool {
	r1 := coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetCancel(AValue bool) {
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2DownloadStartingEventArgs) ResultFilePath() string {
	r1 := coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetResultFilePath(AValue string) {
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2DownloadStartingEventArgs) Handled() bool {
	r1 := coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadStartingEventArgs) SetHandled(AValue bool) {
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2DownloadStartingEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	coreWebView2DownloadStartingEventArgsImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

var (
	coreWebView2DownloadStartingEventArgsImport       *imports.Imports = nil
	coreWebView2DownloadStartingEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_Cancel", 0),
		/*2*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_Create", 0),
		/*3*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_Deferral", 0),
		/*4*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_DownloadOperation", 0),
		/*5*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_Handled", 0),
		/*6*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_Initialized", 0),
		/*7*/ imports.NewTable("CoreWebView2DownloadStartingEventArgs_ResultFilePath", 0),
	}
)

func coreWebView2DownloadStartingEventArgsImportAPI() *imports.Imports {
	if coreWebView2DownloadStartingEventArgsImport == nil {
		coreWebView2DownloadStartingEventArgsImport = NewDefaultImports()
		coreWebView2DownloadStartingEventArgsImport.SetImportTable(coreWebView2DownloadStartingEventArgsImportTables)
		coreWebView2DownloadStartingEventArgsImportTables = nil
	}
	return coreWebView2DownloadStartingEventArgsImport
}
