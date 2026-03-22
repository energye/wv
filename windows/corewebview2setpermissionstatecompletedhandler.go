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

// ICoreWebView2SetPermissionStateCompletedHandler Parent: IInterfacedObject
type ICoreWebView2SetPermissionStateCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT) types.HRESULT // function
	AsIntfSetPermissionStateCompletedHandler() uintptr
}

type TCoreWebView2SetPermissionStateCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2SetPermissionStateCompletedHandler) Invoke(errorCode types.HRESULT) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2SetPermissionStateCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode))
	return types.HRESULT(r)
}

func (m *TCoreWebView2SetPermissionStateCompletedHandler) AsIntfSetPermissionStateCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2SetPermissionStateCompletedHandler class constructor
func NewCoreWebView2SetPermissionStateCompletedHandler(events IWVBrowserBase) ICoreWebView2SetPermissionStateCompletedHandler {
	var setPermissionStateCompletedHandlerPtr uintptr // ICoreWebView2SetPermissionStateCompletedHandler
	r := coreWebView2SetPermissionStateCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&setPermissionStateCompletedHandlerPtr)))
	ret := AsCoreWebView2SetPermissionStateCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, setPermissionStateCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2SetPermissionStateCompletedHandlerOnce   base.Once
	coreWebView2SetPermissionStateCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2SetPermissionStateCompletedHandlerAPI() *imports.Imports {
	coreWebView2SetPermissionStateCompletedHandlerOnce.Do(func() {
		coreWebView2SetPermissionStateCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2SetPermissionStateCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2SetPermissionStateCompletedHandler_Create", 0), // constructor NewCoreWebView2SetPermissionStateCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2SetPermissionStateCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2SetPermissionStateCompletedHandlerImport
}
