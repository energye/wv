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

// ICoreWebView2WebResourceResponseViewGetContentCompletedHandler0 Parent: lcl.IInterfacedObject
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler0 interface {
	lcl.IInterfacedObject
	// Invoke
	//  Provides the result of the corresponding asynchronous method.
	Invoke(errorCode types.HRESULT, result lcl.IStreamAdapter) types.HRESULT // function
}

// ICoreWebView2WebResourceResponseViewGetContentCompletedHandler Parent: ICoreWebView2WebResourceResponseViewGetContentCompletedHandler0
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler interface {
	ICoreWebView2WebResourceResponseViewGetContentCompletedHandler0
	AsIntfWebResourceResponseViewGetContentCompletedHandler() uintptr
}

type TCoreWebView2WebResourceResponseViewGetContentCompletedHandler struct {
	lcl.TInterfacedObject
}

func (m *TCoreWebView2WebResourceResponseViewGetContentCompletedHandler) Invoke(errorCode types.HRESULT, result lcl.IStreamAdapter) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceResponseViewGetContentCompletedHandlerAPI().SysCallN(1, m.Instance(), uintptr(errorCode), base.GetObjectUintptr(result))
	return types.HRESULT(r)
}

func (m *TCoreWebView2WebResourceResponseViewGetContentCompletedHandler) ResourceID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2WebResourceResponseViewGetContentCompletedHandlerAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2WebResourceResponseViewGetContentCompletedHandler) AsIntfWebResourceResponseViewGetContentCompletedHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2WebResourceResponseViewGetContentCompletedHandler class constructor
func NewCoreWebView2WebResourceResponseViewGetContentCompletedHandler(events IWVBrowserBase, resourceID int32) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	var webResourceResponseViewGetContentCompletedHandlerPtr uintptr // ICoreWebView2WebResourceResponseViewGetContentCompletedHandler
	r := coreWebView2WebResourceResponseViewGetContentCompletedHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(resourceID), uintptr(base.UnsafePointer(&webResourceResponseViewGetContentCompletedHandlerPtr)))
	ret := AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, webResourceResponseViewGetContentCompletedHandlerPtr)
	}
	return ret
}

var (
	coreWebView2WebResourceResponseViewGetContentCompletedHandlerOnce   base.Once
	coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport *imports.Imports = nil
)

func coreWebView2WebResourceResponseViewGetContentCompletedHandlerAPI() *imports.Imports {
	coreWebView2WebResourceResponseViewGetContentCompletedHandlerOnce.Do(func() {
		coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport = api.NewDefaultImports()
		coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2WebResourceResponseViewGetContentCompletedHandler_Create", 0), // constructor NewCoreWebView2WebResourceResponseViewGetContentCompletedHandler
			/* 1 */ imports.NewTable("TCoreWebView2WebResourceResponseViewGetContentCompletedHandler_Invoke", 0), // function Invoke
			/* 2 */ imports.NewTable("TCoreWebView2WebResourceResponseViewGetContentCompletedHandler_ResourceID", 0), // property ResourceID
		}
	})
	return coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport
}
