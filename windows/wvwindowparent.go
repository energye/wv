//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// IWVWindowParent Parent: IWVWinControl
//
//	Parent control used by VCL and LCL applications to show the web contents.
type IWVWindowParent interface {
	IWVWinControl
	// Browser
	//  Browser associated to this control to show web contents.
	Browser() IWVBrowserBase // property
	// SetBrowser Set Browser
	SetBrowser(AValue IWVBrowserBase) // property
}

// TWVWindowParent Parent: TWVWinControl
//
//	Parent control used by VCL and LCL applications to show the web contents.
type TWVWindowParent struct {
	TWVWinControl
}

func NewWVWindowParent(AOwner IComponent) IWVWindowParent {
	r1 := wVWindowParentImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsWVWindowParent(r1)
}

func (m *TWVWindowParent) Browser() IWVBrowserBase {
	var resultWVBrowserBase uintptr
	wVWindowParentImportAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(unsafePointer(&resultWVBrowserBase)))
	return AsWVBrowserBase(resultWVBrowserBase)
}

func (m *TWVWindowParent) SetBrowser(AValue IWVBrowserBase) {
	wVWindowParentImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

var (
	wVWindowParentImport       *imports.Imports = nil
	wVWindowParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WVWindowParent_Browser", 0),
		/*1*/ imports.NewTable("WVWindowParent_Create", 0),
	}
)

func wVWindowParentImportAPI() *imports.Imports {
	if wVWindowParentImport == nil {
		wVWindowParentImport = NewDefaultImports()
		wVWindowParentImport.SetImportTable(wVWindowParentImportTables)
		wVWindowParentImportTables = nil
	}
	return wVWindowParentImport
}
