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

// ICoreWebView2BrowserExtensionRemoveCompletedHandler Parent: IInterfacedObject
type ICoreWebView2BrowserExtensionRemoveCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT) types.HRESULT // function
	AsIntfBrowserExtensionRemoveCompletedHandler() uintptr
}

type TCoreWebView2BrowserExtensionRemoveCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2BrowserExtensionRemoveCompletedHandler) Invoke(errorCode types.HRESULT) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2BrowserExtensionRemoveCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode))
	return types.HRESULT(r)
}

func (m *TCoreWebView2BrowserExtensionRemoveCompletedHandler) AsIntfBrowserExtensionRemoveCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2BrowserExtensionRemoveCompletedHandler class constructor
func NewCoreWebView2BrowserExtensionRemoveCompletedHandler(events IWVBrowserBase, extensionID string) ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	var browserExtensionRemoveCompletedHandlerPtr uintptr // ICoreWebView2BrowserExtensionRemoveCompletedHandler
	r := coreWebView2BrowserExtensionRemoveCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), api.PasStr(extensionID), uintptr(base.UnsafePointer(&browserExtensionRemoveCompletedHandlerPtr)))
	ret := AsCoreWebView2BrowserExtensionRemoveCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, browserExtensionRemoveCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2BrowserExtensionRemoveCompletedHandlerOnce   base.Once
	coreWebView2BrowserExtensionRemoveCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2BrowserExtensionRemoveCompletedHandlerAPI() *imports.Imports {
	coreWebView2BrowserExtensionRemoveCompletedHandlerOnce.Do(func() {
		coreWebView2BrowserExtensionRemoveCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2BrowserExtensionRemoveCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2BrowserExtensionRemoveCompletedHandler_Create", 0), // constructor NewCoreWebView2BrowserExtensionRemoveCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2BrowserExtensionRemoveCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2BrowserExtensionRemoveCompletedHandlerImport
}
