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

// ICoreWebView2ClearBrowsingDataCompletedHandler0 Parent: lcl.IInterfacedObject
type ICoreWebView2ClearBrowsingDataCompletedHandler0 interface {
	lcl.IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT) types.HRESULT // function
}

// ICoreWebView2ClearBrowsingDataCompletedHandler Parent: ICoreWebView2ClearBrowsingDataCompletedHandler0
type ICoreWebView2ClearBrowsingDataCompletedHandler interface {
	ICoreWebView2ClearBrowsingDataCompletedHandler0
	AsIntfClearBrowsingDataCompletedHandler() uintptr
}

type TCoreWebView2ClearBrowsingDataCompletedHandler struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2ClearBrowsingDataCompletedHandler) Invoke(errorCode types.HRESULT) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ClearBrowsingDataCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode))
	return types.HRESULT(r)
}

func (m *TCoreWebView2ClearBrowsingDataCompletedHandler) AsIntfClearBrowsingDataCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2ClearBrowsingDataCompletedHandler class constructor
func NewCoreWebView2ClearBrowsingDataCompletedHandler(events IWVBrowserBase) ICoreWebView2ClearBrowsingDataCompletedHandler {
	var clearBrowsingDataCompletedHandlerPtr uintptr // ICoreWebView2ClearBrowsingDataCompletedHandler
	r := coreWebView2ClearBrowsingDataCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&clearBrowsingDataCompletedHandlerPtr)))
	ret := AsCoreWebView2ClearBrowsingDataCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, clearBrowsingDataCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2ClearBrowsingDataCompletedHandlerOnce   base.Once
	coreWebView2ClearBrowsingDataCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2ClearBrowsingDataCompletedHandlerAPI() *imports.Imports {
	coreWebView2ClearBrowsingDataCompletedHandlerOnce.Do(func() {
		coreWebView2ClearBrowsingDataCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2ClearBrowsingDataCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ClearBrowsingDataCompletedHandler_Create", 0), // constructor NewCoreWebView2ClearBrowsingDataCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2ClearBrowsingDataCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2ClearBrowsingDataCompletedHandlerImport
}
