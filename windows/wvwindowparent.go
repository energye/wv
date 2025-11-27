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

// IWVWindowParent Parent: IWVWinControl
type IWVWindowParent interface {
	IWVWinControl
	// Browser
	//  Browser associated to this control to show web contents.
	Browser() IWVBrowserBase         // property Browser Getter
	SetBrowser(value IWVBrowserBase) // property Browser Setter
}

type TWVWindowParent struct {
	TWVWinControl
}

func (m *TWVWindowParent) Browser() IWVBrowserBase {
	if !m.IsValid() {
		return nil
	}
	r := wVWindowParentAPI().SysCallN(1, 0, m.Instance())
	return AsWVBrowserBase(r)
}

func (m *TWVWindowParent) SetBrowser(value IWVBrowserBase) {
	if !m.IsValid() {
		return
	}
	wVWindowParentAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewWindowParent class constructor
func NewWindowParent(owner lcl.IComponent) IWVWindowParent {
	r := wVWindowParentAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsWVWindowParent(r)
}

var (
	wVWindowParentOnce   base.Once
	wVWindowParentImport *imports.Imports = nil
)

func wVWindowParentAPI() *imports.Imports {
	wVWindowParentOnce.Do(func() {
		wVWindowParentImport = api.NewDefaultImports()
		wVWindowParentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVWindowParent_Create", 0), // constructor NewWindowParent
			/* 1 */ imports.NewTable("TWVWindowParent_Browser", 0), // property Browser
		}
	})
	return wVWindowParentImport
}
