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
	"github.com/energye/lcl/types"
)

// ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Parent: IInterfacedObject
type ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result ICoreWebView2PermissionSettingCollectionView) types.HRESULT // function
	AsIntfGetNonDefaultPermissionSettingsCompletedHandler() uintptr
}

type TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) Invoke(errorCode types.HRESULT, result ICoreWebView2PermissionSettingCollectionView) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) AsIntfGetNonDefaultPermissionSettingsCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler class constructor
func NewCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(events IWVBrowserBase) ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	var getNonDefaultPermissionSettingsCompletedHandlerPtr uintptr // ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler
	r := coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&getNonDefaultPermissionSettingsCompletedHandlerPtr)))
	ret := AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, getNonDefaultPermissionSettingsCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerOnce   base.Once
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerAPI() *imports.Imports {
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerOnce.Do(func() {
		coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler_Create", 0), // constructor NewCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2GetNonDefaultPermissionSettingsCompletedHandlerImport
}
