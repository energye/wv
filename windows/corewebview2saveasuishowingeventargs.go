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

// ICoreWebView2SaveAsUIShowingEventArgs Parent: IObject
type ICoreWebView2SaveAsUIShowingEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SaveAsUIShowingEventArgs // property BaseIntf Getter
	// ContentMimeType
	//  Get the Mime type of content to be saved.
	ContentMimeType() string // property ContentMimeType Getter
	// Cancel
	//  Sets the `Cancel` property. Set this property to `TRUE` to cancel the Save As action
	//  and prevent the download from starting. ShowSaveAsUI returns
	//  `COREWEBVIEW2_SAVE_AS_UI_RESULT_CANCELLED` in this case. The default value is `FALSE`.
	Cancel() bool         // property Cancel Getter
	SetCancel(value bool) // property Cancel Setter
	// SuppressDefaultDialog
	//  Sets the `SuppressDefaultDialog` property, which indicates whether the system
	//  default dialog is suppressed. When `SuppressDefaultDialog` is `FALSE`, the default
	//  Save As dialog is shown and the values assigned through `SaveAsFilePath`, `AllowReplace`
	//  and `Kind` are ignored when the event args invoke completed.
	//
	//  Set `SuppressDefaultDialog` to `TRUE` to perform a silent Save As. When
	//  `SuppressDefaultDialog` is `TRUE`, the system dialog is skipped and the
	//  `SaveAsFilePath`, `AllowReplace` and `Kind` values are used.
	//
	//  The default value is FALSE.
	SuppressDefaultDialog() bool         // property SuppressDefaultDialog Getter
	SetSuppressDefaultDialog(value bool) // property SuppressDefaultDialog Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. This will defer showing the
	//  default Save As dialog and performing the Save As operation.
	Deferral() ICoreWebView2Deferral // property Deferral Getter
	// SaveAsFilePath
	//  Set the `SaveAsFilePath` property for Save As. `SaveAsFilePath` is an absolute path
	//  of the location. It includes the file name and extension. If `SaveAsFilePath` is not
	//  valid (for example, the root drive does not exist), Save As is denied and
	//  `COREWEBVIEW2_SAVE_AS_INVALID_PATH` is returned.
	//
	//  If the associated download completes successfully, a target file is saved at
	//  this location. If the Kind property is `COREWEBVIEW2_SAVE_AS_KIND_COMPLETE`,
	//  there will be an additional directory with resources files.
	//
	//  The default value is a system suggested path, based on users' local environment.
	SaveAsFilePath() string         // property SaveAsFilePath Getter
	SetSaveAsFilePath(value string) // property SaveAsFilePath Setter
	// AllowReplace
	//  `AllowReplace` allows user to control what happens when a file already
	//  exists in the file path to which the Save As operation is saving.
	//  Setting this property to `TRUE` allows existing files to be replaced.
	//  Setting this property to `FALSE` will not replace existing files and will return
	//  `COREWEBVIEW2_SAVE_AS_UI_RESULT_FILE_ALREADY_EXISTS`.
	//
	//  The default value is `FALSE`.
	AllowReplace() bool         // property AllowReplace Getter
	SetAllowReplace(value bool) // property AllowReplace Setter
	// Kind
	//  Sets the `Kind` property to save documents of different kinds. See the
	//  `COREWEBVIEW2_SAVE_AS_KIND` enum for a description of the different options.
	//  If the kind is not allowed for the current document, ShowSaveAsUI returns
	//  `COREWEBVIEW2_SAVE_AS_UI_RESULT_KIND_NOT_SUPPORTED`.
	//
	//  The default value is `COREWEBVIEW2_SAVE_AS_KIND_DEFAULT`.
	Kind() wvTypes.TWVSaveAsKind         // property Kind Getter
	SetKind(value wvTypes.TWVSaveAsKind) // property Kind Setter
}

type TCoreWebView2SaveAsUIShowingEventArgs struct {
	TObject
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) BaseIntf() (result ICoreWebView2SaveAsUIShowingEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2SaveAsUIShowingEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) ContentMimeType() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SuppressDefaultDialog() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SetSuppressDefaultDialog(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SaveAsFilePath() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SetSaveAsFilePath(value string) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) AllowReplace() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SetAllowReplace(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) Kind() wvTypes.TWVSaveAsKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(9, 0, m.Instance())
	return wvTypes.TWVSaveAsKind(r)
}

func (m *TCoreWebView2SaveAsUIShowingEventArgs) SetKind(value wvTypes.TWVSaveAsKind) {
	if !m.IsValid() {
		return
	}
	coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

// NewCoreWebView2SaveAsUIShowingEventArgs class constructor
func NewCoreWebView2SaveAsUIShowingEventArgs(args ICoreWebView2SaveAsUIShowingEventArgs) ICoreWebView2SaveAsUIShowingEventArgs {
	r := coreWebView2SaveAsUIShowingEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2SaveAsUIShowingEventArgs(r)
}

var (
	coreWebView2SaveAsUIShowingEventArgsOnce   base.Once
	coreWebView2SaveAsUIShowingEventArgsImport *imports.Imports = nil
)

func coreWebView2SaveAsUIShowingEventArgsAPI() *imports.Imports {
	coreWebView2SaveAsUIShowingEventArgsOnce.Do(func() {
		coreWebView2SaveAsUIShowingEventArgsImport = api.NewDefaultImports()
		coreWebView2SaveAsUIShowingEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_Create", 0), // constructor NewCoreWebView2SaveAsUIShowingEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_ContentMimeType", 0), // property ContentMimeType
			/* 4 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_Cancel", 0), // property Cancel
			/* 5 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_SuppressDefaultDialog", 0), // property SuppressDefaultDialog
			/* 6 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_Deferral", 0), // property Deferral
			/* 7 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_SaveAsFilePath", 0), // property SaveAsFilePath
			/* 8 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_AllowReplace", 0), // property AllowReplace
			/* 9 */ imports.NewTable("TCoreWebView2SaveAsUIShowingEventArgs_Kind", 0), // property Kind
		}
	})
	return coreWebView2SaveAsUIShowingEventArgsImport
}
