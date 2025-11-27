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
)

// IWVBrowser Parent: IWVBrowserBase
type IWVBrowser interface {
	IWVBrowserBase
	AsIntfWVBrowserEvents() uintptr
}

type TWVBrowser struct {
	TWVBrowserBase
}

func (m *TWVBrowser) AsIntfWVBrowserEvents() uintptr {
	return m.GetIntfPointer(0)
}

// NewBrowser class constructor
func NewBrowser(owner lcl.IComponent) IWVBrowser {
	var wVBrowserEventsPtr uintptr // IWVBrowserEvents
	r := wVBrowserAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&wVBrowserEventsPtr)))
	ret := AsWVBrowser(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, wVBrowserEventsPtr)
	}
	return ret
}

var (
	wVBrowserOnce   base.Once
	wVBrowserImport *imports.Imports = nil
)

func wVBrowserAPI() *imports.Imports {
	wVBrowserOnce.Do(func() {
		wVBrowserImport = api.NewDefaultImports()
		wVBrowserImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVBrowser_Create", 0), // constructor NewBrowser
		}
	})
	return wVBrowserImport
}
