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

// IWkWebviewParent Parent: ICustomPanel
type IWkWebviewParent interface {
	ICustomPanel
	ScrolledWindow() TScrolledWindow          // property
	SetScrolledWindow(AValue TScrolledWindow) // property
	Resize()                                  // procedure
	SetWebview(aValue IWkWebview)             // procedure
	FreeChild()                               // procedure
}

// TWkWebviewParent Parent: TCustomPanel
type TWkWebviewParent struct {
	TCustomPanel
}

func NewWkWebviewParent(aOwner IComponent) IWkWebviewParent {
	r1 := wkWebviewParentImportAPI().SysCallN(0, GetObjectUintptr(aOwner))
	return AsWkWebviewParent(r1)
}

func (m *TWkWebviewParent) ScrolledWindow() TScrolledWindow {
	r1 := wkWebviewParentImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TScrolledWindow(r1)
}

func (m *TWkWebviewParent) SetScrolledWindow(AValue TScrolledWindow) {
	wkWebviewParentImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TWkWebviewParent) Resize() {
	wkWebviewParentImportAPI().SysCallN(2, m.Instance())
}

func (m *TWkWebviewParent) SetWebview(aValue IWkWebview) {
	wkWebviewParentImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(aValue))
}

func (m *TWkWebviewParent) FreeChild() {
	wkWebviewParentImportAPI().SysCallN(1, m.Instance())
}

var (
	wkWebviewParentImport       *imports.Imports = nil
	wkWebviewParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebviewParent_Create", 0),
		/*1*/ imports.NewTable("WkWebviewParent_FreeChild", 0),
		/*2*/ imports.NewTable("WkWebviewParent_Resize", 0),
		/*3*/ imports.NewTable("WkWebviewParent_ScrolledWindow", 0),
		/*4*/ imports.NewTable("WkWebviewParent_SetWebview", 0),
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
