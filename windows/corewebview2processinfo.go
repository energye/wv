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

// ICoreWebView2ProcessInfo Parent: IObject
type ICoreWebView2ProcessInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessInfo         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2ProcessInfo) // property BaseIntf Setter
	// Kind
	//  The kind of the process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</see>
	Kind() wvTypes.TWVProcessKind // property Kind Getter
	// KindStr
	//  The kind of the process in string format.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_kind">See the ICoreWebView2ProcessInfo article.</see>
	KindStr() string // property KindStr Getter
	// ProcessId
	//  The process id of the process.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2processinfo#get_processid">See the ICoreWebView2ProcessInfo article.</see>
	ProcessId() int32 // property ProcessId Getter
}

type TCoreWebView2ProcessInfo struct {
	TObject
}

func (m *TCoreWebView2ProcessInfo) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProcessInfoAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ProcessInfo) BaseIntf() (result ICoreWebView2ProcessInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessInfoAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfo(resultPtr)
	return
}

func (m *TCoreWebView2ProcessInfo) SetBaseIntf(value ICoreWebView2ProcessInfo) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProcessInfoAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ProcessInfo) Kind() wvTypes.TWVProcessKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessInfoAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVProcessKind(r)
}

func (m *TCoreWebView2ProcessInfo) KindStr() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ProcessInfoAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ProcessInfo) ProcessId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProcessInfoAPI().SysCallN(5, m.Instance())
	return int32(r)
}

// NewCoreWebView2ProcessInfo class constructor
func NewCoreWebView2ProcessInfo(baseIntf ICoreWebView2ProcessInfo) ICoreWebView2ProcessInfo {
	r := coreWebView2ProcessInfoAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ProcessInfo(r)
}

var (
	coreWebView2ProcessInfoOnce   base.Once
	coreWebView2ProcessInfoImport *imports.Imports = nil
)

func coreWebView2ProcessInfoAPI() *imports.Imports {
	coreWebView2ProcessInfoOnce.Do(func() {
		coreWebView2ProcessInfoImport = api.NewDefaultImports()
		coreWebView2ProcessInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProcessInfo_Create", 0), // constructor NewCoreWebView2ProcessInfo
			/* 1 */ imports.NewTable("TCoreWebView2ProcessInfo_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ProcessInfo_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ProcessInfo_Kind", 0), // property Kind
			/* 4 */ imports.NewTable("TCoreWebView2ProcessInfo_KindStr", 0), // property KindStr
			/* 5 */ imports.NewTable("TCoreWebView2ProcessInfo_ProcessId", 0), // property ProcessId
		}
	})
	return coreWebView2ProcessInfoImport
}
