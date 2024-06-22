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

// ICoreWebView2PermissionSetting Parent: IObject
//
//	Provides a set of properties for a permission setting.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting">See the ICoreWebView2PermissionSetting article.</a>
type ICoreWebView2PermissionSetting interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionSetting // property
	// PermissionKind
	//  The kind of the permission setting. See `COREWEBVIEW2_PERMISSION_KIND` for
	//  more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionkind">See the ICoreWebView2PermissionSetting article.</a>
	PermissionKind() TWVPermissionKind // property
	// PermissionOrigin
	//  The origin of the permission setting.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionorigin">See the ICoreWebView2PermissionSetting article.</a>
	PermissionOrigin() string // property
	// PermissionState
	//  The state of the permission setting.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionstate">See the ICoreWebView2PermissionSetting article.</a>
	PermissionState() TWVPermissionState // property
}

// TCoreWebView2PermissionSetting Parent: TObject
//
//	Provides a set of properties for a permission setting.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting">See the ICoreWebView2PermissionSetting article.</a>
type TCoreWebView2PermissionSetting struct {
	TObject
}

func NewCoreWebView2PermissionSetting(aBaseIntf ICoreWebView2PermissionSetting) ICoreWebView2PermissionSetting {
	r1 := coreWebView2PermissionSettingImportAPI().SysCallN(1, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PermissionSetting(r1)
}

func (m *TCoreWebView2PermissionSetting) Initialized() bool {
	r1 := coreWebView2PermissionSettingImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PermissionSetting) BaseIntf() ICoreWebView2PermissionSetting {
	var resultCoreWebView2PermissionSetting uintptr
	coreWebView2PermissionSettingImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PermissionSetting)))
	return AsCoreWebView2PermissionSetting(resultCoreWebView2PermissionSetting)
}

func (m *TCoreWebView2PermissionSetting) PermissionKind() TWVPermissionKind {
	r1 := coreWebView2PermissionSettingImportAPI().SysCallN(3, m.Instance())
	return TWVPermissionKind(r1)
}

func (m *TCoreWebView2PermissionSetting) PermissionOrigin() string {
	r1 := coreWebView2PermissionSettingImportAPI().SysCallN(4, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2PermissionSetting) PermissionState() TWVPermissionState {
	r1 := coreWebView2PermissionSettingImportAPI().SysCallN(5, m.Instance())
	return TWVPermissionState(r1)
}

var (
	coreWebView2PermissionSettingImport       *imports.Imports = nil
	coreWebView2PermissionSettingImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2PermissionSetting_BaseIntf", 0),
		/*1*/ imports.NewTable("CoreWebView2PermissionSetting_Create", 0),
		/*2*/ imports.NewTable("CoreWebView2PermissionSetting_Initialized", 0),
		/*3*/ imports.NewTable("CoreWebView2PermissionSetting_PermissionKind", 0),
		/*4*/ imports.NewTable("CoreWebView2PermissionSetting_PermissionOrigin", 0),
		/*5*/ imports.NewTable("CoreWebView2PermissionSetting_PermissionState", 0),
	}
)

func coreWebView2PermissionSettingImportAPI() *imports.Imports {
	if coreWebView2PermissionSettingImport == nil {
		coreWebView2PermissionSettingImport = NewDefaultImports()
		coreWebView2PermissionSettingImport.SetImportTable(coreWebView2PermissionSettingImportTables)
		coreWebView2PermissionSettingImportTables = nil
	}
	return coreWebView2PermissionSettingImport
}
