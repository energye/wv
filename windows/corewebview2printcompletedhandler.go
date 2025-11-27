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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2PrintCompletedHandler0 Parent: lcl.IInterfacedObject
type ICoreWebView2PrintCompletedHandler0 interface {
	lcl.IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result wvTypes.COREWEBVIEW2_PRINT_STATUS) types.HRESULT // function
}

// ICoreWebView2PrintCompletedHandler Parent: ICoreWebView2PrintCompletedHandler0
type ICoreWebView2PrintCompletedHandler interface {
	ICoreWebView2PrintCompletedHandler0
	AsIntfPrintCompletedHandler() uintptr
}

type TCoreWebView2PrintCompletedHandler struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2PrintCompletedHandler) Invoke(errorCode types.HRESULT, result wvTypes.COREWEBVIEW2_PRINT_STATUS) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2PrintCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), uintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2PrintCompletedHandler) AsIntfPrintCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2PrintCompletedHandler class constructor
func NewCoreWebView2PrintCompletedHandler(events IWVBrowserBase) ICoreWebView2PrintCompletedHandler {
	var printCompletedHandlerPtr uintptr // ICoreWebView2PrintCompletedHandler
	r := coreWebView2PrintCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&printCompletedHandlerPtr)))
	ret := AsCoreWebView2PrintCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, printCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2PrintCompletedHandlerOnce   base.Once
	coreWebView2PrintCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2PrintCompletedHandlerAPI() *imports.Imports {
	coreWebView2PrintCompletedHandlerOnce.Do(func() {
		coreWebView2PrintCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2PrintCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2PrintCompletedHandler_Create", 0), // constructor NewCoreWebView2PrintCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2PrintCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2PrintCompletedHandlerImport
}
