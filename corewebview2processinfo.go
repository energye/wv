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

// ICoreWebView2ProcessInfo Parent: IObject
//
//	Provides a set of properties for a process in the ICoreWebView2Environment.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo">See the ICoreWebView2ProcessInfo article.</a>
type ICoreWebView2ProcessInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessInfo // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ProcessInfo) // property
	// Kind
	//  The kind of the process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</a>
	Kind() TWVProcessKind // property
	// KindStr
	//  The kind of the process in string format.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</a>
	KindStr() string // property
	// ProcessId
	//  The process id of the process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_processid">See the ICoreWebView2ProcessInfo article.</a>
	ProcessId() int32 // property
}

// TCoreWebView2ProcessInfo Parent: TObject
//
//	Provides a set of properties for a process in the ICoreWebView2Environment.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo">See the ICoreWebView2ProcessInfo article.</a>
type TCoreWebView2ProcessInfo struct {
	TObject
}

func NewCoreWebView2ProcessInfo(aBaseIntf ICoreWebView2ProcessInfo) ICoreWebView2ProcessInfo {
	r1 := coreWebView2ProcessInfoImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ProcessInfo(r1)
}

func (m *TCoreWebView2ProcessInfo) Initialized() bool {
	r1 := coreWebView2ProcessInfoImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ProcessInfo) BaseIntf() ICoreWebView2ProcessInfo {
	var resultCoreWebView2ProcessInfo uintptr
	coreWebView2ProcessInfoImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ProcessInfo)))
	return AsCoreWebView2ProcessInfo(resultCoreWebView2ProcessInfo)
}

func (m *TCoreWebView2ProcessInfo) SetBaseIntf(AValue ICoreWebView2ProcessInfo) {
	coreWebView2ProcessInfoImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ProcessInfo) Kind() TWVProcessKind {
	r1 := coreWebView2ProcessInfoImportAPI().SysCallN(3, m.Instance())
	return TWVProcessKind(r1)
}

func (m *TCoreWebView2ProcessInfo) KindStr() string {
	r1 := coreWebView2ProcessInfoImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ProcessInfo) ProcessId() int32 {
	r1 := coreWebView2ProcessInfoImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

var (
	coreWebView2ProcessInfoImport       *imports.Imports = nil
	coreWebView2ProcessInfoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2ProcessInfo_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2ProcessInfo_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2ProcessInfo_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2ProcessInfo_Kind", 0),
		/*4*/ imports.NewTable("CoreWebView2ProcessInfo_KindStr", 0),
		/*5*/ imports.NewTable("CoreWebView2ProcessInfo_ProcessId", 0),
	}
)

func coreWebView2ProcessInfoImportAPI() *imports.Imports {
	if coreWebView2ProcessInfoImport == nil {
		coreWebView2ProcessInfoImport = NewDefaultImports()
		coreWebView2ProcessInfoImport.SetImportTable(coreWebView2ProcessInfoImportTables)
		coreWebView2ProcessInfoImportTables = nil
	}
	return coreWebView2ProcessInfoImport
}
