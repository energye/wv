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

// ICoreWebView2PrintToPdfStreamCompletedHandler Parent: IInterfacedObject
type ICoreWebView2PrintToPdfStreamCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result lcl.IStreamAdapter) types.HRESULT // function
	AsIntfPrintToPdfStreamCompletedHandler() uintptr
}

type TCoreWebView2PrintToPdfStreamCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2PrintToPdfStreamCompletedHandler) Invoke(errorCode types.HRESULT, result lcl.IStreamAdapter) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintToPdfStreamCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2PrintToPdfStreamCompletedHandler) AsIntfPrintToPdfStreamCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2PrintToPdfStreamCompletedHandler class constructor
func NewCoreWebView2PrintToPdfStreamCompletedHandler(events IWVBrowserBase) ICoreWebView2PrintToPdfStreamCompletedHandler {
	var printToPdfStreamCompletedHandlerPtr uintptr // ICoreWebView2PrintToPdfStreamCompletedHandler
	r := coreWebView2PrintToPdfStreamCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&printToPdfStreamCompletedHandlerPtr)))
	ret := AsCoreWebView2PrintToPdfStreamCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printToPdfStreamCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2PrintToPdfStreamCompletedHandlerOnce   base.Once
	coreWebView2PrintToPdfStreamCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2PrintToPdfStreamCompletedHandlerAPI() *imports.Imports {
	coreWebView2PrintToPdfStreamCompletedHandlerOnce.Do(func() {
		coreWebView2PrintToPdfStreamCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2PrintToPdfStreamCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PrintToPdfStreamCompletedHandler_Create", 0), // constructor NewCoreWebView2PrintToPdfStreamCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2PrintToPdfStreamCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2PrintToPdfStreamCompletedHandlerImport
}
