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

// IWkWebviewParent Parent: ICustomControl
type IWkWebviewParent interface {
	ICustomControl
	UpdateBounds()                 // procedure
	Resize()                       // procedure
	SetWebview(aWebview WkWebview) // procedure
}

// TWkWebviewParent Parent: TCustomControl
type TWkWebviewParent struct {
	TCustomControl
}

func NewWkWebviewParent(aOwner IComponent) IWkWebviewParent {
	r1 := wkWebviewParentImportAPI().SysCallN(0, GetObjectUintptr(aOwner))
	return AsWkWebviewParent(r1)
}

func (m *TWkWebviewParent) UpdateBounds() {
	wkWebviewParentImportAPI().SysCallN(3, m.Instance())
}

func (m *TWkWebviewParent) Resize() {
	wkWebviewParentImportAPI().SysCallN(1, m.Instance())
}

func (m *TWkWebviewParent) SetWebview(aWebview WkWebview) {
	wkWebviewParentImportAPI().SysCallN(2, m.Instance(), uintptr(aWebview))
}

var (
	wkWebviewParentImport       *imports.Imports = nil
	wkWebviewParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebviewParent_Create", 0),
		/*1*/ imports.NewTable("WkWebviewParent_Resize", 0),
		/*2*/ imports.NewTable("WkWebviewParent_SetWebview", 0),
		/*3*/ imports.NewTable("WkWebviewParent_UpdateBounds", 0),
	}
)

func wkWebviewParentImportAPI() *imports.Imports {
	if wkWebviewParentImport == nil {
		wkWebviewParentImport = NewDefaultImports()
		wkWebviewParentImport.SetImportTable(wkWebviewParentImportTables)
		wkWebviewParentImportTables = nil
	}
	return wkWebviewParentImport
}
