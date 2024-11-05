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

// IWkContextMenu Root Interface
type IWkContextMenu interface {
	IObject
	Data() WebKitContextMenu                                // function
	GetItemsLength() int32                                  // function
	GetItemAtPosition(position int32) WebKitContextMenuItem // function
	Prepend(item WebKitContextMenuItem)                     // procedure
	Append(item WebKitContextMenuItem)                      // procedure
	Insert(item WebKitContextMenuItem, position int32)      // procedure
	Remove(item WebKitContextMenuItem)                      // procedure
	RemoveAll()                                             // procedure
}

// TWkContextMenu Root Object
type TWkContextMenu struct {
	TObject
}

func NewWkContextMenu(aContextMenu WebKitContextMenu) IWkContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(1, uintptr(aContextMenu))
	return AsWkContextMenu(r1)
}

// WkContextMenuRef -> IWkContextMenu
var WkContextMenuRef wkContextMenu

// wkContextMenu TWkContextMenu Ref
type wkContextMenu uintptr

func (m *wkContextMenu) New() IWkContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(6)
	return AsWkContextMenu(r1)
}

func (m *TWkContextMenu) Data() WebKitContextMenu {
	r1 := wkContextMenuImportAPI().SysCallN(2, m.Instance())
	return WebKitContextMenu(r1)
}

func (m *TWkContextMenu) GetItemsLength() int32 {
	r1 := wkContextMenuImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TWkContextMenu) GetItemAtPosition(position int32) WebKitContextMenuItem {
	r1 := wkContextMenuImportAPI().SysCallN(3, m.Instance(), uintptr(position))
	return WebKitContextMenuItem(r1)
}

func (m *TWkContextMenu) Prepend(item WebKitContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(7, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) Append(item WebKitContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(0, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) Insert(item WebKitContextMenuItem, position int32) {
	wkContextMenuImportAPI().SysCallN(5, m.Instance(), uintptr(item), uintptr(position))
}

func (m *TWkContextMenu) Remove(item WebKitContextMenuItem) {
	wkContextMenuImportAPI().SysCallN(8, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) RemoveAll() {
	wkContextMenuImportAPI().SysCallN(9, m.Instance())
}

var (
	wkContextMenuImport       *imports.Imports = nil
	wkContextMenuImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkContextMenu_Append", 0),
		/*1*/ imports.NewTable("WkContextMenu_Create", 0),
		/*2*/ imports.NewTable("WkContextMenu_Data", 0),
		/*3*/ imports.NewTable("WkContextMenu_GetItemAtPosition", 0),
		/*4*/ imports.NewTable("WkContextMenu_GetItemsLength", 0),
		/*5*/ imports.NewTable("WkContextMenu_Insert", 0),
		/*6*/ imports.NewTable("WkContextMenu_New", 0),
		/*7*/ imports.NewTable("WkContextMenu_Prepend", 0),
		/*8*/ imports.NewTable("WkContextMenu_Remove", 0),
		/*9*/ imports.NewTable("WkContextMenu_RemoveAll", 0),
	}
)

func wkContextMenuImportAPI() *imports.Imports {
	if wkContextMenuImport == nil {
		wkContextMenuImport = NewDefaultImports()
		wkContextMenuImport.SetImportTable(wkContextMenuImportTables)
		wkContextMenuImportTables = nil
	}
	return wkContextMenuImport
}
