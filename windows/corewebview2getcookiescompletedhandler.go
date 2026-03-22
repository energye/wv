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

// ICoreWebView2GetCookiesCompletedHandler Parent: IInterfacedObject
type ICoreWebView2GetCookiesCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result ICoreWebView2CookieList) types.HRESULT // function
	AsIntfGetCookiesCompletedHandler() uintptr
}

type TCoreWebView2GetCookiesCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2GetCookiesCompletedHandler) Invoke(errorCode types.HRESULT, result ICoreWebView2CookieList) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2GetCookiesCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2GetCookiesCompletedHandler) AsIntfGetCookiesCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2GetCookiesCompletedHandler class constructor
func NewCoreWebView2GetCookiesCompletedHandler(events IWVBrowserBase) ICoreWebView2GetCookiesCompletedHandler {
	var getCookiesCompletedHandlerPtr uintptr // ICoreWebView2GetCookiesCompletedHandler
	r := coreWebView2GetCookiesCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&getCookiesCompletedHandlerPtr)))
	ret := AsCoreWebView2GetCookiesCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, getCookiesCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2GetCookiesCompletedHandlerOnce   base.Once
	coreWebView2GetCookiesCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2GetCookiesCompletedHandlerAPI() *imports.Imports {
	coreWebView2GetCookiesCompletedHandlerOnce.Do(func() {
		coreWebView2GetCookiesCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2GetCookiesCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2GetCookiesCompletedHandler_Create", 0), // constructor NewCoreWebView2GetCookiesCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2GetCookiesCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2GetCookiesCompletedHandlerImport
}
