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

// ICoreWebView2PermissionSetting Parent: IObject
type ICoreWebView2PermissionSetting interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PermissionSetting // property BaseIntf Getter
	// PermissionKind
	//  The kind of the permission setting. See `COREWEBVIEW2_PERMISSION_KIND` for
	//  more details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionkind">See the ICoreWebView2PermissionSetting article.</see>
	PermissionKind() wvTypes.TWVPermissionKind // property PermissionKind Getter
	// PermissionOrigin
	//  The origin of the permission setting.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionorigin">See the ICoreWebView2PermissionSetting article.</see>
	PermissionOrigin() string // property PermissionOrigin Getter
	// PermissionState
	//  The state of the permission setting.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2permissionsetting#get_permissionstate">See the ICoreWebView2PermissionSetting article.</see>
	PermissionState() wvTypes.TWVPermissionState // property PermissionState Getter
}

type TCoreWebView2PermissionSetting struct {
	TObject
}

func (m *TCoreWebView2PermissionSetting) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2PermissionSettingAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2PermissionSetting) BaseIntf() (result ICoreWebView2PermissionSetting) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2PermissionSettingAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2PermissionSetting(resultPtr)
	return
}

func (m *TCoreWebView2PermissionSetting) PermissionKind() wvTypes.TWVPermissionKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PermissionSettingAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVPermissionKind(r)
}

func (m *TCoreWebView2PermissionSetting) PermissionOrigin() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2PermissionSettingAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2PermissionSetting) PermissionState() wvTypes.TWVPermissionState {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PermissionSettingAPI().SysCallN(5, m.Instance())
	return wvTypes.TWVPermissionState(r)
}

// NewCoreWebView2PermissionSetting class constructor
func NewCoreWebView2PermissionSetting(baseIntf ICoreWebView2PermissionSetting) ICoreWebView2PermissionSetting {
	r := coreWebView2PermissionSettingAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2PermissionSetting(r)
}

var (
	coreWebView2PermissionSettingOnce   base.Once
	coreWebView2PermissionSettingImport *imports.Imports = nil
)

func coreWebView2PermissionSettingAPI() *imports.Imports {
	coreWebView2PermissionSettingOnce.Do(func() {
		coreWebView2PermissionSettingImport = api.NewDefaultImports()
		coreWebView2PermissionSettingImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PermissionSetting_Create", 0), // constructor NewCoreWebView2PermissionSetting
			/* 1 */ imports.NewTable("TCoreWebView2PermissionSetting_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2PermissionSetting_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2PermissionSetting_PermissionKind", 0), // property PermissionKind
			/* 4 */ imports.NewTable("TCoreWebView2PermissionSetting_PermissionOrigin", 0), // property PermissionOrigin
			/* 5 */ imports.NewTable("TCoreWebView2PermissionSetting_PermissionState", 0), // property PermissionState
		}
	})
	return coreWebView2PermissionSettingImport
}
