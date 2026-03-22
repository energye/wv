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

// ICoreWebView2CustomItemSelectedEventHandler Parent: IInterfacedObject
type ICoreWebView2CustomItemSelectedEventHandler interface {
	IInterfacedObject
	// Invoke
	//  Provides the event args for the corresponding event.
	Invoke(sender ICoreWebView2ContextMenuItem, args lcl.IUnknown) types.HRESULT // function
	AsIntfCustomItemSelectedEventHandler() uintptr
}

type TCoreWebView2CustomItemSelectedEventHandler struct {
	TInterfacedObject
}

func (m *TCoreWebView2CustomItemSelectedEventHandler) Invoke(sender ICoreWebView2ContextMenuItem, args lcl.IUnknown) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CustomItemSelectedEventHandlerAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(sender), base.GetObjectUintptr(args))
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomItemSelectedEventHandler) AsIntfCustomItemSelectedEventHandler() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2CustomItemSelectedEventHandler class constructor
func NewCoreWebView2CustomItemSelectedEventHandler(events IWVBrowserBase) ICoreWebView2CustomItemSelectedEventHandler {
	var customItemSelectedEventHandlerPtr uintptr // ICoreWebView2CustomItemSelectedEventHandler
	r := coreWebView2CustomItemSelectedEventHandlerAPI().SysCallN(0, base.GetObjectUintptr(events), uintptr(base.UnsafePointer(&customItemSelectedEventHandlerPtr)))
	ret := AsCoreWebView2CustomItemSelectedEventHandler(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customItemSelectedEventHandlerPtr)
	}
	return ret
}

var (
	coreWebView2CustomItemSelectedEventHandlerOnce   base.Once
	coreWebView2CustomItemSelectedEventHandlerImport *imports.Imports = nil
)

func coreWebView2CustomItemSelectedEventHandlerAPI() *imports.Imports {
	coreWebView2CustomItemSelectedEventHandlerOnce.Do(func() {
		coreWebView2CustomItemSelectedEventHandlerImport = api.NewDefaultImports()
		coreWebView2CustomItemSelectedEventHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2CustomItemSelectedEventHandler_Create", 0), // constructor NewCoreWebView2CustomItemSelectedEventHandler
			/* 1 */ imports.NewTable("TCoreWebView2CustomItemSelectedEventHandler_Invoke", 0), // function Invoke
		}
	})
	return coreWebView2CustomItemSelectedEventHandlerImport
}
