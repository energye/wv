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
)

// ICoreWebView2File Parent: lcl.IObject
type ICoreWebView2File interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2File         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2File) // property BaseIntf Setter
	// Path
	//  Get the absolute file path.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2file#get_path">See the ICoreWebView2File article.</see>
	Path() string // property Path Getter
}

type TCoreWebView2File struct {
	lcl.TObject
}

func (m *TCoreWebView2File) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FileAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2File) BaseIntf() (result ICoreWebView2File) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FileAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2File(resultPtr)
	return
}

func (m *TCoreWebView2File) SetBaseIntf(value ICoreWebView2File) {
	if !m.IsValid() {
		return
	}
	coreWebView2FileAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2File) Path() string {
	if !m.IsValid() {
		return ""
	}
	r := coreWebView2FileAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

// NewCoreWebView2File class constructor
func NewCoreWebView2File(baseIntf ICoreWebView2File) ICoreWebView2File {
	r := coreWebView2FileAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2File(r)
}

var (
	coreWebView2FileOnce   base.Once
	coreWebView2FileImport *imports.Imports = nil
)

func coreWebView2FileAPI() *imports.Imports {
	coreWebView2FileOnce.Do(func() {
		coreWebView2FileImport = api.NewDefaultImports()
		coreWebView2FileImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2File_Create", 0), // constructor NewCoreWebView2File
			/* 1 */ imports.NewTable("TCoreWebView2File_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2File_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2File_Path", 0), // property Path
		}
	})
	return coreWebView2FileImport
}
