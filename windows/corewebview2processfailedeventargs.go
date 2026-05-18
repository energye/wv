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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2ProcessFailedEventArgs Parent: IObject
type ICoreWebView2ProcessFailedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessFailedEventArgs // property BaseIntf Getter
	// ProcessFailedKind
	//  The kind of process failure that has occurred. This is a combination of
	//  process kind (for example, browser, renderer, gpu) and failure (exit,
	//  unresponsiveness). Renderer processes are further divided in _main frame_
	//  renderer (`RenderProcessExited`, `RenderProcessUnresponsive`) and
	//  _subframe_ renderer (`FrameRenderProcessExited`). To learn about the
	//  conditions under which each failure kind occurs, see
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs#get_processfailedkind">See the ICoreWebView2ProcessFailedEventArgs article.</see>
	ProcessFailedKind() wvTypes.TWVProcessFailedKind // property ProcessFailedKind Getter
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
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_reason">See the ICoreWebView2ProcessFailedEventArgs2 article.</see>
	Reason() wvTypes.TWVProcessFailedReason // property Reason Getter
	// ExtiCode
	//  The exit code of the failing process, for telemetry purposes. The exit
	//  code is always `STILL_ACTIVE` (`259`) when `ProcessFailedKind` is
	//  `COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_exitcode">See the ICoreWebView2ProcessFailedEventArgs2 article.</see>
	ExtiCode() int32 // property ExtiCode Getter
	// ProcessDescription
	//  Description of the process assigned by the WebView2 Runtime. This is a
	//  technical English term appropriate for logging or development purposes,
	//  and not localized for the end user. It applies to utility processes (for
	//  example, "Audio Service", "Video Capture") and plugin processes (for
	//  example, "Flash"). The returned `processDescription` is empty if the
	//  WebView2 Runtime did not assign a description to the process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_processdescription">See the ICoreWebView2ProcessFailedEventArgs2 article.</see>
	ProcessDescription() string // property ProcessDescription Getter
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
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2#get_frameinfosforfailedprocess">See the ICoreWebView2ProcessFailedEventArgs2 article.</see>
	FrameInfosForFailedProcess() ICoreWebView2FrameInfoCollection // property FrameInfosForFailedProcess Getter
	// FailureSourceModulePath
	//  This property is the full path of the module that caused the
	//  crash in cases of Windows Code Integrity failures.
	//  [Windows Code Integrity](/mem/intune/user-help/you-need-to-enable-code-integrity)
	//  is a feature that verifies the integrity and
	//  authenticity of dynamic-link libraries (DLLs)
	//  on Windows systems. It ensures that only trusted
	//  code can run on the system and prevents unauthorized or
	//  malicious modifications.
	//  When ProcessFailed occurred due to a failed Code Integrity check,
	//  this property returns the full path of the file that was prevented from
	//  loading on the system.
	//  The webview2 process which tried to load the DLL will fail with
	//  exit code STATUS_INVALID_IMAGE_HASH(-1073740760).
	//  A file can fail integrity check for various
	//  reasons, such as:
	//  <code>
	//  - It has an invalid or missing signature that does
	//  not match the publisher or signer of the file.
	//  - It has been tampered with or corrupted by malware or other software.
	//  - It has been blocked by an administrator or a security policy.
	//  </code>
	//  This property always will be the empty string if failure is not caused by
	//  STATUS_INVALID_IMAGE_HASH.
	FailureSourceModulePath() string // property FailureSourceModulePath Getter
}

type TCoreWebView2ProcessFailedEventArgs struct {
	TObject
}

func (m *TCoreWebView2ProcessFailedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProcessFailedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ProcessFailedEventArgs) BaseIntf() (result ICoreWebView2ProcessFailedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessFailedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessFailedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ProcessFailedEventArgs) ProcessFailedKind() wvTypes.TWVProcessFailedKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessFailedEventArgsAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVProcessFailedKind(r)
}

func (m *TCoreWebView2ProcessFailedEventArgs) Reason() wvTypes.TWVProcessFailedReason {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessFailedEventArgsAPI().SysCallN(4, m.Instance())
	return wvTypes.TWVProcessFailedReason(r)
}

func (m *TCoreWebView2ProcessFailedEventArgs) ExtiCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessFailedEventArgsAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2ProcessFailedEventArgs) ProcessDescription() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProcessFailedEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ProcessFailedEventArgs) FrameInfosForFailedProcess() (result ICoreWebView2FrameInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessFailedEventArgsAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfoCollection(resultPtr)
	return
}

func (m *TCoreWebView2ProcessFailedEventArgs) FailureSourceModulePath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProcessFailedEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

// NewCoreWebView2ProcessFailedEventArgs class constructor
func NewCoreWebView2ProcessFailedEventArgs(args ICoreWebView2ProcessFailedEventArgs) ICoreWebView2ProcessFailedEventArgs {
	r := coreWebView2ProcessFailedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ProcessFailedEventArgs(r)
}

var (
	coreWebView2ProcessFailedEventArgsOnce   base.Once
	coreWebView2ProcessFailedEventArgsImport *imports.Imports = nil
)

func coreWebView2ProcessFailedEventArgsAPI() *imports.Imports {
	coreWebView2ProcessFailedEventArgsOnce.Do(func() {
		coreWebView2ProcessFailedEventArgsImport = api.NewDefaultImports()
		coreWebView2ProcessFailedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_Create", 0), // constructor NewCoreWebView2ProcessFailedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_ProcessFailedKind", 0), // property ProcessFailedKind
			/* 4 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_Reason", 0), // property Reason
			/* 5 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_ExtiCode", 0), // property ExtiCode
			/* 6 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_ProcessDescription", 0), // property ProcessDescription
			/* 7 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_FrameInfosForFailedProcess", 0), // property FrameInfosForFailedProcess
			/* 8 */ imports.NewTable("TCoreWebView2ProcessFailedEventArgs_FailureSourceModulePath", 0), // property FailureSourceModulePath
		}
	})
	return coreWebView2ProcessFailedEventArgsImport
}
