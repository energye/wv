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
	"github.com/energye/lcl/types"

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2DownloadOperation Parent: IObject
type ICoreWebView2DownloadOperation interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// Cancel
	//  Cancels the download. If canceled, the default download dialog shows
	//  that the download was canceled. Host should set the `Cancel` property from
	//  `ICoreWebView2SDownloadStartingEventArgs` if the download should be
	//  canceled without displaying the default download dialog.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#cancel">See the ICoreWebView2DownloadOperation article.</see>
	Cancel() bool // function
	// Pause
	//  Pauses the download. If paused, the default download dialog shows that the
	//  download is paused. No effect if download is already paused. Pausing a
	//  download changes the state to `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED`
	//  with `InterruptReason` set to `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#pause">See the ICoreWebView2DownloadOperation article.</see>
	Pause() bool // function
	// Resume
	//  Resumes a paused download. May also resume a download that was interrupted
	//  for another reason, if `CanResume` returns true. Resuming a download changes
	//  the state from `COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED` to
	//  `COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#resume">See the ICoreWebView2DownloadOperation article.</see>
	Resume() bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2DownloadOperation // property BaseIntf Getter
	// DownloadID
	//  Custom ID used to identify this download operation.
	DownloadID() int32 // property DownloadID Getter
	// URI
	//  The URI of the download.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_uri">See the ICoreWebView2DownloadOperation article.</see>
	URI() string // property URI Getter
	// ContentDisposition
	//  The Content-Disposition header value from the download's HTTP response.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_contentdisposition">See the ICoreWebView2DownloadOperation article.</see>
	ContentDisposition() string // property ContentDisposition Getter
	// MimeType
	//  MIME type of the downloaded content.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_mimetype">See the ICoreWebView2DownloadOperation article.</see>
	MimeType() string // property MimeType Getter
	// TotalBytesToReceive
	//  The expected size of the download in total number of bytes based on the
	//  HTTP Content-Length header. Returns -1 if the size is unknown.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_totalbytestoreceive">See the ICoreWebView2DownloadOperation article.</see>
	TotalBytesToReceive() int64 // property TotalBytesToReceive Getter
	// BytesReceived
	//  The number of bytes that have been written to the download file.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_bytesreceived">See the ICoreWebView2DownloadOperation article.</see>
	BytesReceived() int64 // property BytesReceived Getter
	// EstimatedEndTime
	//  The estimated end time in [ISO 8601 Date and Time Format](https://www.iso.org/iso-8601-date-and-time-format.html).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_estimatedendtime">See the ICoreWebView2DownloadOperation article.</see>
	EstimatedEndTime() types.TDateTime // property EstimatedEndTime Getter
	// ResultFilePath
	//  The absolute path to the download file, including file name. Host can change
	//  this from `ICoreWebView2DownloadStartingEventArgs`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_resultfilepath">See the ICoreWebView2DownloadOperation article.</see>
	ResultFilePath() string // property ResultFilePath Getter
	// State
	//  The state of the download. A download can be in progress, interrupted, or
	//  completed. See `COREWEBVIEW2_DOWNLOAD_STATE` for descriptions of states.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_state">See the ICoreWebView2DownloadOperation article.</see>
	State() wvTypes.TWVDownloadState // property State Getter
	// InterruptReason
	//  The reason why connection with file host was broken.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_interruptreason">See the ICoreWebView2DownloadOperation article.</see>
	InterruptReason() wvTypes.TWVDownloadInterruptReason // property InterruptReason Getter
	// CanResume
	//  Returns true if an interrupted download can be resumed. Downloads with
	//  the following interrupt reasons may automatically resume without you
	//  calling any methods:
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH`,
	//  `COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT`.
	//  In these cases download progress may be restarted with `BytesReceived`
	//  reset to 0.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#get_canresume">See the ICoreWebView2DownloadOperation article.</see>
	CanResume() bool // property CanResume Getter
}

type TCoreWebView2DownloadOperation struct {
	TObject
}

func (m *TCoreWebView2DownloadOperation) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadOperation) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadOperation) Pause() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadOperation) Resume() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadOperation) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2DownloadOperation) BaseIntf() (result ICoreWebView2DownloadOperation) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2DownloadOperationAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2DownloadOperation(resultPtr)
	return
}

func (m *TCoreWebView2DownloadOperation) DownloadID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2DownloadOperation) URI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DownloadOperationAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2DownloadOperation) ContentDisposition() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DownloadOperationAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2DownloadOperation) MimeType() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DownloadOperationAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2DownloadOperation) TotalBytesToReceive() (result int64) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadOperationAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2DownloadOperation) BytesReceived() (result int64) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadOperationAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2DownloadOperation) EstimatedEndTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	coreWebView2DownloadOperationAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2DownloadOperation) ResultFilePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2DownloadOperationAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2DownloadOperation) State() wvTypes.TWVDownloadState {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(15, m.Instance())
	return wvTypes.TWVDownloadState(r)
}

func (m *TCoreWebView2DownloadOperation) InterruptReason() wvTypes.TWVDownloadInterruptReason {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(16, m.Instance())
	return wvTypes.TWVDownloadInterruptReason(r)
}

func (m *TCoreWebView2DownloadOperation) CanResume() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2DownloadOperationAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

// NewCoreWebView2DownloadOperation class constructor
func NewCoreWebView2DownloadOperation(baseIntf ICoreWebView2DownloadOperation, downloadID int32) ICoreWebView2DownloadOperation {
	r := coreWebView2DownloadOperationAPI().SysCallN(0, base.GetObjectUintptr(baseIntf), uintptr(downloadID))
	return AsCoreWebView2DownloadOperation(r)
}

var (
	coreWebView2DownloadOperationOnce   base.Once
	coreWebView2DownloadOperationImport *imports.Imports = nil
)

func coreWebView2DownloadOperationAPI() *imports.Imports {
	coreWebView2DownloadOperationOnce.Do(func() {
		coreWebView2DownloadOperationImport = api.NewDefaultImports()
		coreWebView2DownloadOperationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2DownloadOperation_Create", 0), // constructor NewCoreWebView2DownloadOperation
			/* 1 */ imports.NewTable("TCoreWebView2DownloadOperation_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2DownloadOperation_Cancel", 0), // function Cancel
			/* 3 */ imports.NewTable("TCoreWebView2DownloadOperation_Pause", 0), // function Pause
			/* 4 */ imports.NewTable("TCoreWebView2DownloadOperation_Resume", 0), // function Resume
			/* 5 */ imports.NewTable("TCoreWebView2DownloadOperation_Initialized", 0), // property Initialized
			/* 6 */ imports.NewTable("TCoreWebView2DownloadOperation_BaseIntf", 0), // property BaseIntf
			/* 7 */ imports.NewTable("TCoreWebView2DownloadOperation_DownloadID", 0), // property DownloadID
			/* 8 */ imports.NewTable("TCoreWebView2DownloadOperation_URI", 0), // property URI
			/* 9 */ imports.NewTable("TCoreWebView2DownloadOperation_ContentDisposition", 0), // property ContentDisposition
			/* 10 */ imports.NewTable("TCoreWebView2DownloadOperation_MimeType", 0), // property MimeType
			/* 11 */ imports.NewTable("TCoreWebView2DownloadOperation_TotalBytesToReceive", 0), // property TotalBytesToReceive
			/* 12 */ imports.NewTable("TCoreWebView2DownloadOperation_BytesReceived", 0), // property BytesReceived
			/* 13 */ imports.NewTable("TCoreWebView2DownloadOperation_EstimatedEndTime", 0), // property EstimatedEndTime
			/* 14 */ imports.NewTable("TCoreWebView2DownloadOperation_ResultFilePath", 0), // property ResultFilePath
			/* 15 */ imports.NewTable("TCoreWebView2DownloadOperation_State", 0), // property State
			/* 16 */ imports.NewTable("TCoreWebView2DownloadOperation_InterruptReason", 0), // property InterruptReason
			/* 17 */ imports.NewTable("TCoreWebView2DownloadOperation_CanResume", 0), // property CanResume
		}
	})
	return coreWebView2DownloadOperationImport
}
