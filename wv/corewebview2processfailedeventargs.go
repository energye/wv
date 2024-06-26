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

// ICoreWebView2ProcessFailedEventArgs Parent: IObject
//
//	Event args for the ProcessFailed event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs">See the ICoreWebView2ProcessFailedEventArgs article.</a>
type ICoreWebView2ProcessFailedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessFailedEventArgs // property
	// ProcessFailedKind
	//  The kind of process failure that has occurred. This is a combination of
	//  process kind(for example, browser, renderer, gpu) and failure(exit,
	//  unresponsiveness). Renderer processes are further divided in _main frame_
	//  renderer(`RenderProcessExited`, `RenderProcessUnresponsive`) and
	//  _subframe_ renderer(`FrameRenderProcessExited`). To learn about the
	//  conditions under which each failure kind occurs, see
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs#get_processfailedkind">See the ICoreWebView2ProcessFailedEventArgs article.</a>
	ProcessFailedKind() TWVProcessFailedKind // property
	// Reason
	//  The reason for the process failure. Some of the reasons are only
	//  applicable to specific values of
	//  `ICoreWebView2ProcessFailedEventArgs.ProcessFailedKind`, and the
	//  following `ProcessFailedKind` values always return the indicated reason
	//  value:
	//  <code>
	//  ProcessFailedKind | Reason
	//  ---|---
	//  COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED | COREWEBVIEW2_PROCESS_FAILED_REASON_UNEXPECTED
	//  COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE | COREWEBVIEW2_PROCESS_FAILED_REASON_UNRESPONSIVE
	//  </code>
	//  For other `ProcessFailedKind` values, the reason may be any of the reason
	//  values. To learn about what these values mean, see
	//  `COREWEBVIEW2_PROCESS_FAILED_REASON`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_reason">See the ICoreWebView2ProcessFailedEventArgs2 article.</a>
	Reason() TWVProcessFailedReason // property
	// ExtiCode
	//  The exit code of the failing process, for telemetry purposes. The exit
	//  code is always `STILL_ACTIVE`(`259`) when `ProcessFailedKind` is
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_exitcode">See the ICoreWebView2ProcessFailedEventArgs2 article.</a>
	ExtiCode() int32 // property
	// ProcessDescription
	//  Description of the process assigned by the WebView2 Runtime. This is a
	//  technical English term appropriate for logging or development purposes,
	//  and not localized for the end user. It applies to utility processes(for
	//  example, "Audio Service", "Video Capture") and plugin processes(for
	//  example, "Flash"). The returned `processDescription` is empty if the
	//  WebView2 Runtime did not assign a description to the process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_processdescription">See the ICoreWebView2ProcessFailedEventArgs2 article.</a>
	ProcessDescription() string // property
	// FrameInfosForFailedProcess
	//  The collection of `FrameInfo`s for frames in the `ICoreWebView2` that were
	//  being rendered by the failed process. The content in these frames is
	//  replaced with an error page.
	//  This is only available when `ProcessFailedKind` is
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED`;
	//  `frames` is `null` for all other process failure kinds, including the case
	//  in which the failed process was the renderer for the main frame and
	//  subframes within it, for which the failure kind is
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_frameinfosforfailedprocess">See the ICoreWebView2ProcessFailedEventArgs2 article.</a>
	FrameInfosForFailedProcess() ICoreWebView2FrameInfoCollection // property
}

// TCoreWebView2ProcessFailedEventArgs Parent: TObject
//
//	Event args for the ProcessFailed event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs">See the ICoreWebView2ProcessFailedEventArgs article.</a>
type TCoreWebView2ProcessFailedEventArgs struct {
	TObject
}

func NewCoreWebView2ProcessFailedEventArgs(aArgs ICoreWebView2ProcessFailedEventArgs) ICoreWebView2ProcessFailedEventArgs {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(1, GetObjectUintptr(aArgs))
	return AsCoreWebView2ProcessFailedEventArgs(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) Initialized() bool {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) BaseIntf() ICoreWebView2ProcessFailedEventArgs {
	var resultCoreWebView2ProcessFailedEventArgs uintptr
	coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessFailedEventArgs)))
	return AsCoreWebView2ProcessFailedEventArgs(resultCoreWebView2ProcessFailedEventArgs)
}

func (m *TCoreWebView2ProcessFailedEventArgs) ProcessFailedKind() TWVProcessFailedKind {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(6, m.Instance())
	return TWVProcessFailedKind(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) Reason() TWVProcessFailedReason {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(7, m.Instance())
	return TWVProcessFailedReason(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) ExtiCode() int32 {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) ProcessDescription() string {
	r1 := coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ProcessFailedEventArgs) FrameInfosForFailedProcess() ICoreWebView2FrameInfoCollection {
	var resultCoreWebView2FrameInfoCollection uintptr
	coreWebView2ProcessFailedEventArgsImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfoCollection)))
	return AsCoreWebView2FrameInfoCollection(resultCoreWebView2FrameInfoCollection)
}

var (
	coreWebView2ProcessFailedEventArgsImport       *imports.Imports = nil
	coreWebView2ProcessFailedEventArgsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_ExtiCode", 0),
		/*3*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_FrameInfosForFailedProcess", 0),
		/*4*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_Initialized", 0),
		/*5*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_ProcessDescription", 0),
		/*6*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_ProcessFailedKind", 0),
		/*7*/ imports.NewTable("CoreWebView2ProcessFailedEventArgs_Reason", 0),
	}
)

func coreWebView2ProcessFailedEventArgsImportAPI() *imports.Imports {
	if coreWebView2ProcessFailedEventArgsImport == nil {
		coreWebView2ProcessFailedEventArgsImport = NewDefaultImports()
		coreWebView2ProcessFailedEventArgsImport.SetImportTable(coreWebView2ProcessFailedEventArgsImportTables)
		coreWebView2ProcessFailedEventArgsImportTables = nil
	}
	return coreWebView2ProcessFailedEventArgsImport
}
