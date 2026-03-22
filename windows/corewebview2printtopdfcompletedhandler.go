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

// ICoreWebView2PrintToPdfCompletedHandler Parent: IInterfacedObject
type ICoreWebView2PrintToPdfCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result int32) types.HRESULT // function
	AsIntfPrintToPdfCompletedHandler() uintptr
}

type TCoreWebView2PrintToPdfCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2PrintToPdfCompletedHandler) Invoke(errorCode types.HRESULT, result int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintToPdfCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), uintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2PrintToPdfCompletedHandler) AsIntfPrintToPdfCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2PrintToPdfCompletedHandler class constructor
func NewCoreWebView2PrintToPdfCompletedHandler(events IWVBrowserBase) ICoreWebView2PrintToPdfCompletedHandler {
	var printToPdfCompletedHandlerPtr uintptr // ICoreWebView2PrintToPdfCompletedHandler
	r := coreWebView2PrintToPdfCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&printToPdfCompletedHandlerPtr)))
	ret := AsCoreWebView2PrintToPdfCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printToPdfCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2PrintToPdfCompletedHandlerOnce   base.Once
	coreWebView2PrintToPdfCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2PrintToPdfCompletedHandlerAPI() *imports.Imports {
	coreWebView2PrintToPdfCompletedHandlerOnce.Do(func() {
		coreWebView2PrintToPdfCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2PrintToPdfCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PrintToPdfCompletedHandler_Create", 0), // constructor NewCoreWebView2PrintToPdfCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2PrintToPdfCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2PrintToPdfCompletedHandlerImport
}
