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

// ICoreWebView2ProfileAddBrowserExtensionCompletedHandler Parent: IInterfacedObject
type ICoreWebView2ProfileAddBrowserExtensionCompletedHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result ICoreWebView2BrowserExtension) types.HRESULT // function
	AsIntfProfileAddBrowserExtensionCompletedHandler() uintptr
}

type TCoreWebView2ProfileAddBrowserExtensionCompletedHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2ProfileAddBrowserExtensionCompletedHandler) Invoke(errorCode types.HRESULT, result ICoreWebView2BrowserExtension) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ProfileAddBrowserExtensionCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2ProfileAddBrowserExtensionCompletedHandler) AsIntfProfileAddBrowserExtensionCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2ProfileAddBrowserExtensionCompletedHandler class constructor
func NewCoreWebView2ProfileAddBrowserExtensionCompletedHandler(events IWVBrowserBase) ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	var profileAddBrowserExtensionCompletedHandlerPtr uintptr // ICoreWebView2ProfileAddBrowserExtensionCompletedHandler
	r := coreWebView2ProfileAddBrowserExtensionCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&profileAddBrowserExtensionCompletedHandlerPtr)))
	ret := AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, profileAddBrowserExtensionCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2ProfileAddBrowserExtensionCompletedHandlerOnce   base.Once
	coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2ProfileAddBrowserExtensionCompletedHandlerAPI() *imports.Imports {
	coreWebView2ProfileAddBrowserExtensionCompletedHandlerOnce.Do(func() {
		coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ProfileAddBrowserExtensionCompletedHandler_Create", 0), // constructor NewCoreWebView2ProfileAddBrowserExtensionCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2ProfileAddBrowserExtensionCompletedHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2ProfileAddBrowserExtensionCompletedHandlerImport
}
