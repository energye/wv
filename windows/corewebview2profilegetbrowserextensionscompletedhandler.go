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

// ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler0 Parent: lcl.IInterfacedObject
type ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler0 interface {
	lcl.IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result ICoreWebView2BrowserExtensionList) types.HRESULT // function
}

// ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler Parent: ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler0
type ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler interface {
	ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler0
	AsIntfProfileGetBrowserExtensionsCompletedHandler() uintptr
}

type TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler) Invoke(errorCode types.HRESULT, result ICoreWebView2BrowserExtensionList) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProfileGetBrowserExtensionsCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler) AsIntfProfileGetBrowserExtensionsCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2ProfileGetBrowserExtensionsCompletedHandler class constructor
func NewCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(events IWVBrowserBase) ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	var profileGetBrowserExtensionsCompletedHandlerPtr uintptr // ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler
	r := coreWebView2ProfileGetBrowserExtensionsCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&profileGetBrowserExtensionsCompletedHandlerPtr)))
	ret := AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, profileGetBrowserExtensionsCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2ProfileGetBrowserExtensionsCompletedHandlerOnce   base.Once
	coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2ProfileGetBrowserExtensionsCompletedHandlerAPI() *imports.Imports {
	coreWebView2ProfileGetBrowserExtensionsCompletedHandlerOnce.Do(func() {
		coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler_Create", 0), // constructor NewCoreWebView2ProfileGetBrowserExtensionsCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2ProfileGetBrowserExtensionsCompletedHandlerImport
}
