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

// ICoreWebView2SaveFileSecurityCheckStartingEventArgs Parent: IObject
type ICoreWebView2SaveFileSecurityCheckStartingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SaveFileSecurityCheckStartingEventArgs // property BaseIntf Getter
	// CancelSave
	//  Set if cancel the upcoming save/download. `TRUE` means the action
	//  will be cancelled before validations in default policy.
	//  The default value is `FALSE`.
	CancelSave() bool         // property CancelSave Getter
	SetCancelSave(value bool) // property CancelSave Setter
	// DocumentOriginUri
	//  Get the document origin URI of this file save operation.
	DocumentOriginUri() string // property DocumentOriginUri Getter
	// FileExtension
	//  Get the extension of file to be saved.
	//  The file extension is the extension portion of the FilePath,
	//  preserving original case.
	//  Only final extension with period "." will be provided. For example,
	//  "*.tar.gz" is a double extension, where the ".gz" will be its
	//  final extension.
	//  File extension can be empty, if the file name has no extension
	//  at all.
	FileExtension() string // property FileExtension Getter
	// FilePath
	//  Get the full path of file to be saved. This includes the
	//  file name and extension.
	//  This method doesn't provide path validation, the returned
	//  string may longer than MAX_PATH.
	FilePath() string // property FilePath Getter
	// SuppressDefaultPolicy
	//  Set if the default policy checking and security warning will be
	//  suppressed. `TRUE` means it will be suppressed.
	//  The default value is `FALSE`.
	SuppressDefaultPolicy() bool         // property SuppressDefaultPolicy Getter
	SetSuppressDefaultPolicy(value bool) // property SuppressDefaultPolicy Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to complete
	//  the SaveFileSecurityCheckStartingEvent.
	//  The default policy checking and any default UI will be blocked temporarily,
	//  saving file to local won't start, until the deferral is completed.
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2SaveFileSecurityCheckStartingEventArgs struct {
	TObject
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) BaseIntf() (result ICoreWebView2SaveFileSecurityCheckStartingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2SaveFileSecurityCheckStartingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) CancelSave() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) SetCancelSave(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) DocumentOriginUri() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(4, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) FileExtension() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) FilePath() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) SuppressDefaultPolicy() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) SetSuppressDefaultPolicy(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2SaveFileSecurityCheckStartingEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2SaveFileSecurityCheckStartingEventArgs class constructor
func NewCoreWebView2SaveFileSecurityCheckStartingEventArgs(args ICoreWebView2SaveFileSecurityCheckStartingEventArgs) ICoreWebView2SaveFileSecurityCheckStartingEventArgs {
	r := coreWebView2SaveFileSecurityCheckStartingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2SaveFileSecurityCheckStartingEventArgs(r)
}

var (
	coreWebView2SaveFileSecurityCheckStartingEventArgsOnce   base.Once
	coreWebView2SaveFileSecurityCheckStartingEventArgsImport *imports.Imports = nil
)

func coreWebView2SaveFileSecurityCheckStartingEventArgsAPI() *imports.Imports {
	coreWebView2SaveFileSecurityCheckStartingEventArgsOnce.Do(func() {
		coreWebView2SaveFileSecurityCheckStartingEventArgsImport = api.NewDefaultImports()
		coreWebView2SaveFileSecurityCheckStartingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_Create", 0), // constructor NewCoreWebView2SaveFileSecurityCheckStartingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_CancelSave", 0), // property CancelSave
			/* 4 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_DocumentOriginUri", 0), // property DocumentOriginUri
			/* 5 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_FileExtension", 0), // property FileExtension
			/* 6 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_FilePath", 0), // property FilePath
			/* 7 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_SuppressDefaultPolicy", 0), // property SuppressDefaultPolicy
			/* 8 */ imports.NewTable("TCoreWebView2SaveFileSecurityCheckStartingEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2SaveFileSecurityCheckStartingEventArgsImport
}
