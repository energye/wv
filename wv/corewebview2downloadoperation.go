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

// ICoreWebView2DownloadOperation Parent: IObject
//
//	Represents a download operation. Gives access to the download's metadata
//	and supports a user canceling, pausing, or resuming the download.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation">See the ICoreWebView2DownloadOperation article.</a>
type ICoreWebView2DownloadOperation interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DownloadOperation // property
	// DownloadID
	//  Custom ID used to identify this download operation.
	DownloadID() int32 // property
	// URI
	//  The URI of the download.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_uri">See the ICoreWebView2DownloadOperation article.</a>
	URI() string // property
	// ContentDisposition
	//  The Content-Disposition header value from the download's HTTP response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_contentdisposition">See the ICoreWebView2DownloadOperation article.</a>
	ContentDisposition() string // property
	// MimeType
	//  MIME type of the downloaded content.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_mimetype">See the ICoreWebView2DownloadOperation article.</a>
	MimeType() string // property
	// TotalBytesToReceive
	//  The expected size of the download in total number of bytes based on the
	//  HTTP Content-Length header. Returns -1 if the size is unknown.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_totalbytestoreceive">See the ICoreWebView2DownloadOperation article.</a>
	TotalBytesToReceive() (resultInt64 int64) // property
	// BytesReceived
	//  The number of bytes that have been written to the download file.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_bytesreceived">See the ICoreWebView2DownloadOperation article.</a>
	BytesReceived() (resultInt64 int64) // property
	// EstimatedEndTime
	//  The estimated end time in [ISO 8601 Date and Time Format](https://www.iso.org/iso-8601-date-and-time-format.html).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_estimatedendtime">See the ICoreWebView2DownloadOperation article.</a>
	EstimatedEndTime() (resultDateTime float64) // property
	// ResultFilePath
	//  The absolute path to the download file, including file name. Host can change
	//  this from `ICoreWebView2DownloadStartingEventArgs`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_resultfilepath">See the ICoreWebView2DownloadOperation article.</a>
	ResultFilePath() string // property
	// State
	//  The state of the download. A download can be in progress, interrupted, or
	//  completed. See `COREWEBVIEW2_DOWNLOAD_STATE` for descriptions of states.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_state">See the ICoreWebView2DownloadOperation article.</a>
	State() TWVDownloadState // property
	// InterruptReason
	//  The reason why connection with file host was broken.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_interruptreason">See the ICoreWebView2DownloadOperation article.</a>
	InterruptReason() TWVDownloadInterruptReason // property
	// CanResume
	//  Returns true if an interrupted download can be resumed. Downloads with
	//  the following interrupt reasons may automatically resume without you
	//  calling any methods:
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT`.
	//  In these cases download progress may be restarted with `BytesReceived`
	//  reset to 0.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_canresume">See the ICoreWebView2DownloadOperation article.</a>
	CanResume() bool // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// Cancel
	//  Cancels the download. If canceled, the default download dialog shows
	//  that the download was canceled. Host should set the `Cancel` property from
	//  `ICoreWebView2SDownloadStartingEventArgs` if the download should be
	//  canceled without displaying the default download dialog.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#cancel">See the ICoreWebView2DownloadOperation article.</a>
	Cancel() bool // function
	// Pause
	//  Pauses the download. If paused, the default download dialog shows that the
	//  download is paused. No effect if download is already paused. Pausing a
	//  download changes the state to `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED`
	//  with `InterruptReason` set to `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#pause">See the ICoreWebView2DownloadOperation article.</a>
	Pause() bool // function
	// Resume
	//  Resumes a paused download. May also resume a download that was interrupted
	//  for another reason, if `CanResume` returns true. Resuming a download changes
	//  the state from `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED` to
	//  `COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#resume">See the ICoreWebView2DownloadOperation article.</a>
	Resume() bool // function
}

// TCoreWebView2DownloadOperation Parent: TObject
//
//	Represents a download operation. Gives access to the download's metadata
//	and supports a user canceling, pausing, or resuming the download.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation">See the ICoreWebView2DownloadOperation article.</a>
type TCoreWebView2DownloadOperation struct {
	TObject
}

func NewCoreWebView2DownloadOperation(aBaseIntf ICoreWebView2DownloadOperation, aDownloadID int32) ICoreWebView2DownloadOperation {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(6, GetObjectUintptr(aBaseIntf), uintptr(aDownloadID))
	return AsCoreWebView2DownloadOperation(r1)
}

func (m *TCoreWebView2DownloadOperation) Initialized() bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) BaseIntf() ICoreWebView2DownloadOperation {
	var resultCoreWebView2DownloadOperation uintptr
	coreWebView2DownloadOperationImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2DownloadOperation)))
	return AsCoreWebView2DownloadOperation(resultCoreWebView2DownloadOperation)
}

func (m *TCoreWebView2DownloadOperation) DownloadID() int32 {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2DownloadOperation) URI() string {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(17, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) ContentDisposition() string {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) MimeType() string {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(11, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) TotalBytesToReceive() (resultInt64 int64) {
	coreWebView2DownloadOperationImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCoreWebView2DownloadOperation) BytesReceived() (resultInt64 int64) {
	coreWebView2DownloadOperationImportAPI().SysCallN(2, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCoreWebView2DownloadOperation) EstimatedEndTime() (resultDateTime float64) {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(8, m.Instance())
	return *(*float64)(unsafePointer(&r1))
}

func (m *TCoreWebView2DownloadOperation) ResultFilePath() string {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(13, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2DownloadOperation) State() TWVDownloadState {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(15, m.Instance())
	return TWVDownloadState(r1)
}

func (m *TCoreWebView2DownloadOperation) InterruptReason() TWVDownloadInterruptReason {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(10, m.Instance())
	return TWVDownloadInterruptReason(r1)
}

func (m *TCoreWebView2DownloadOperation) CanResume() bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Cancel() bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Pause() bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2DownloadOperation) Resume() bool {
	r1 := coreWebView2DownloadOperationImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

var (
	coreWebView2DownloadOperationImport       *imports.Imports = nil
	coreWebView2DownloadOperationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2DownloadOperation_AddAllBrowserEvents", 0),
		/*1*/ imports.NewTable("CoreWebView2DownloadOperation_BaseIntf", 0),
		/*2*/ imports.NewTable("CoreWebView2DownloadOperation_BytesReceived", 0),
		/*3*/ imports.NewTable("CoreWebView2DownloadOperation_CanResume", 0),
		/*4*/ imports.NewTable("CoreWebView2DownloadOperation_Cancel", 0),
		/*5*/ imports.NewTable("CoreWebView2DownloadOperation_ContentDisposition", 0),
		/*6*/ imports.NewTable("CoreWebView2DownloadOperation_Create", 0),
		/*7*/ imports.NewTable("CoreWebView2DownloadOperation_DownloadID", 0),
		/*8*/ imports.NewTable("CoreWebView2DownloadOperation_EstimatedEndTime", 0),
		/*9*/ imports.NewTable("CoreWebView2DownloadOperation_Initialized", 0),
		/*10*/ imports.NewTable("CoreWebView2DownloadOperation_InterruptReason", 0),
		/*11*/ imports.NewTable("CoreWebView2DownloadOperation_MimeType", 0),
		/*12*/ imports.NewTable("CoreWebView2DownloadOperation_Pause", 0),
		/*13*/ imports.NewTable("CoreWebView2DownloadOperation_ResultFilePath", 0),
		/*14*/ imports.NewTable("CoreWebView2DownloadOperation_Resume", 0),
		/*15*/ imports.NewTable("CoreWebView2DownloadOperation_State", 0),
		/*16*/ imports.NewTable("CoreWebView2DownloadOperation_TotalBytesToReceive", 0),
		/*17*/ imports.NewTable("CoreWebView2DownloadOperation_URI", 0),
	}
)

func coreWebView2DownloadOperationImportAPI() *imports.Imports {
	if coreWebView2DownloadOperationImport == nil {
		coreWebView2DownloadOperationImport = NewDefaultImports()
		coreWebView2DownloadOperationImport.SetImportTable(coreWebView2DownloadOperationImportTables)
		coreWebView2DownloadOperationImportTables = nil
	}
	return coreWebView2DownloadOperationImport
}
