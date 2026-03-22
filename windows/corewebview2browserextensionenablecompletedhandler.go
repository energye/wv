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

// ICoreWebView2BrowserExtensionEnableCompletedHandler Parent: IInterfacedObject
type ICoreWebView2BrowserExtensionEnableCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT) types.HRESULT // function
	AsIntfBrowserExtensionEnableCompletedHandler() uintptr
}

type TCoreWebView2BrowserExtensionEnableCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2BrowserExtensionEnableCompletedHandler) Invoke(errorCode types.HRESULT) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2BrowserExtensionEnableCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode))
	return types.HRESULT(r)
}

func (m *TCoreWebView2BrowserExtensionEnableCompletedHandler) AsIntfBrowserExtensionEnableCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2BrowserExtensionEnableCompletedHandler class constructor
func NewCoreWebView2BrowserExtensionEnableCompletedHandler(events IWVBrowserBase, extensionID string) ICoreWebView2BrowserExtensionEnableCompletedHandler {
	var browserExtensionEnableCompletedHandlerPtr uintptr // ICoreWebView2BrowserExtensionEnableCompletedHandler
	r := coreWebView2BrowserExtensionEnableCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), api.PasStr(extensionID), uintptr(base.UnsafePointer(&browserExtensionEnableCompletedHandlerPtr)))
	ret := AsCoreWebView2BrowserExtensionEnableCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserExtensionEnableCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2BrowserExtensionEnableCompletedHandlerOnce   base.Once
	coreWebView2BrowserExtensionEnableCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2BrowserExtensionEnableCompletedHandlerAPI() *imports.Imports {
	coreWebView2BrowserExtensionEnableCompletedHandlerOnce.Do(func() {
		coreWebView2BrowserExtensionEnableCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2BrowserExtensionEnableCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BrowserExtensionEnableCompletedHandler_Create", 0), // constructor NewCoreWebView2BrowserExtensionEnableCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2BrowserExtensionEnableCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2BrowserExtensionEnableCompletedHandlerImport
}
