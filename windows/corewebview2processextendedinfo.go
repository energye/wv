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

// ICoreWebView2ProcessExtendedInfo Parent: IObject
type ICoreWebView2ProcessExtendedInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ProcessExtendedInfo         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2ProcessExtendedInfo) // property BaseIntf Setter
	// ProcessInfo
	//  The process info of the current process.
	ProcessInfo() ICoreWebView2ProcessInfo // property ProcessInfo Getter
	// AssociatedFrameInfos
	//  The collection of associated `FrameInfo`s which are actively running
	//  (showing or hiding UI elements) in this renderer process. `AssociatedFrameInfos`
	//  will only be populated when this `ICoreWebView2ProcessExtendedInfo`
	//  corresponds to a renderer process. Non-renderer processes will always
	//  have an empty `AssociatedFrameInfos`. The `AssociatedFrameInfos` may
	//  also be empty for renderer processes that have no active frames.
	AssociatedFrameInfos() ICoreWebView2FrameInfoCollection // property AssociatedFrameInfos Getter
}

type TCoreWebView2ProcessExtendedInfo struct {
	TObject
}

func (m *TCoreWebView2ProcessExtendedInfo) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ProcessExtendedInfoAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ProcessExtendedInfo) BaseIntf() (result ICoreWebView2ProcessExtendedInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessExtendedInfoAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessExtendedInfo(resultPtr)
	return
}

func (m *TCoreWebView2ProcessExtendedInfo) SetBaseIntf(value ICoreWebView2ProcessExtendedInfo) {
	if !m.IsValid() {
		return
	}
	coreWebView2ProcessExtendedInfoAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ProcessExtendedInfo) ProcessInfo() (result ICoreWebView2ProcessInfo) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessExtendedInfoAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ProcessInfo(resultPtr)
	return
}

func (m *TCoreWebView2ProcessExtendedInfo) AssociatedFrameInfos() (result ICoreWebView2FrameInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ProcessExtendedInfoAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfoCollection(resultPtr)
	return
}

// NewCoreWebView2ProcessExtendedInfo class constructor
func NewCoreWebView2ProcessExtendedInfo(baseIntf ICoreWebView2ProcessExtendedInfo) ICoreWebView2ProcessExtendedInfo {
	r := coreWebView2ProcessExtendedInfoAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ProcessExtendedInfo(r)
}

var (
	coreWebView2ProcessExtendedInfoOnce   base.Once
	coreWebView2ProcessExtendedInfoImport *imports.Imports = nil
)

func coreWebView2ProcessExtendedInfoAPI() *imports.Imports {
	coreWebView2ProcessExtendedInfoOnce.Do(func() {
		coreWebView2ProcessExtendedInfoImport = api.NewDefaultImports()
		coreWebView2ProcessExtendedInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProcessExtendedInfo_Create", 0), // constructor NewCoreWebView2ProcessExtendedInfo
			/* 1 */ imports.NewTable("TCoreWebView2ProcessExtendedInfo_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ProcessExtendedInfo_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ProcessExtendedInfo_ProcessInfo", 0), // property ProcessInfo
			/* 4 */ imports.NewTable("TCoreWebView2ProcessExtendedInfo_AssociatedFrameInfos", 0), // property AssociatedFrameInfos
		}
	})
	return coreWebView2ProcessExtendedInfoImport
}
