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
	"github.com/energye/lcl/types"
)

// ICoreWebView2TrySuspendCompletedHandler0 Parent: lcl.IInterfacedObject
type ICoreWebView2TrySuspendCompletedHandler0 interface {
	lcl.IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result int32) types.HRESULT // function
}

// ICoreWebView2TrySuspendCompletedHandler Parent: ICoreWebView2TrySuspendCompletedHandler0
type ICoreWebView2TrySuspendCompletedHandler interface {
	ICoreWebView2TrySuspendCompletedHandler0
	AsIntfTrySuspendCompletedHandler() uintptr
}

type TCoreWebView2TrySuspendCompletedHandler struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2TrySuspendCompletedHandler) Invoke(errorCode types.HRESULT, result int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2TrySuspendCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), uintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2TrySuspendCompletedHandler) AsIntfTrySuspendCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2TrySuspendCompletedHandler class constructor
func NewCoreWebView2TrySuspendCompletedHandler(events IWVBrowserBase) ICoreWebView2TrySuspendCompletedHandler {
	var trySuspendCompletedHandlerPtr uintptr // ICoreWebView2TrySuspendCompletedHandler
	r := coreWebView2TrySuspendCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&trySuspendCompletedHandlerPtr)))
	ret := AsCoreWebView2TrySuspendCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, trySuspendCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2TrySuspendCompletedHandlerOnce   base.Once
	coreWebView2TrySuspendCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2TrySuspendCompletedHandlerAPI() *imports.Imports {
	coreWebView2TrySuspendCompletedHandlerOnce.Do(func() {
		coreWebView2TrySuspendCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2TrySuspendCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2TrySuspendCompletedHandler_Create", 0), // constructor NewCoreWebView2TrySuspendCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2TrySuspendCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2TrySuspendCompletedHandlerImport
}
