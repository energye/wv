//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkContextMenu Parent: IObject
type IWkContextMenu interface {
	IObject
	Data() wvTypes.WebKitContextMenu                                // function
	GetItemsLength() int32                                          // function
	GetItemAtPosition(position int32) wvTypes.WebKitContextMenuItem // function
	Prepend(item wvTypes.WebKitContextMenuItem)                     // procedure
	Append(item wvTypes.WebKitContextMenuItem)                      // procedure
	Insert(item wvTypes.WebKitContextMenuItem, position int32)      // procedure
	Remove(item wvTypes.WebKitContextMenuItem)                      // procedure
	RemoveAll()                                                     // procedure
}

type TWkContextMenu struct {
	TObject
}

func (m *TWkContextMenu) Data() wvTypes.WebKitContextMenu {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitContextMenu(r)
}

func (m *TWkContextMenu) GetItemsLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TWkContextMenu) GetItemAtPosition(position int32) wvTypes.WebKitContextMenuItem {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuAPI().SysCallN(3, m.Instance(), uintptr(position))
	return wvTypes.WebKitContextMenuItem(r)
}

func (m *TWkContextMenu) Prepend(item wvTypes.WebKitContextMenuItem) {
	if !m.IsValid() {
		return
	}
	wkContextMenuAPI().SysCallN(5, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) Append(item wvTypes.WebKitContextMenuItem) {
	if !m.IsValid() {
		return
	}
	wkContextMenuAPI().SysCallN(6, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) Insert(item wvTypes.WebKitContextMenuItem, position int32) {
	if !m.IsValid() {
		return
	}
	wkContextMenuAPI().SysCallN(7, m.Instance(), uintptr(item), uintptr(position))
}

func (m *TWkContextMenu) Remove(item wvTypes.WebKitContextMenuItem) {
	if !m.IsValid() {
		return
	}
	wkContextMenuAPI().SysCallN(8, m.Instance(), uintptr(item))
}

func (m *TWkContextMenu) RemoveAll() {
	if !m.IsValid() {
		return
	}
	wkContextMenuAPI().SysCallN(9, m.Instance())
}

// ContextMenu  is static instance
var ContextMenu _ContextMenuClass

// _ContextMenuClass is class type defined by TWkContextMenu
type _ContextMenuClass uintptr

func (_ContextMenuClass) New() IWkContextMenu {
	r := wkContextMenuAPI().SysCallN(4)
	return AsWkContextMenu(r)
}

// NewContextMenu class constructor
func NewContextMenu(contextMenu wvTypes.WebKitContextMenu) IWkContextMenu {
	r := wkContextMenuAPI().SysCallN(0, uintptr(contextMenu))
	return AsWkContextMenu(r)
}

var (
	wkContextMenuOnce   base.Once
	wkContextMenuImport *imports.Imports = nil
)

func wkContextMenuAPI() *imports.Imports {
	wkContextMenuOnce.Do(func() {
		wkContextMenuImport = api.NewDefaultImports()
		wkContextMenuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkContextMenu_Create", 0), // constructor NewContextMenu
			/* 1 */ imports.NewTable("TWkContextMenu_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkContextMenu_GetItemsLength", 0), // function GetItemsLength
			/* 3 */ imports.NewTable("TWkContextMenu_GetItemAtPosition", 0), // function GetItemAtPosition
			/* 4 */ imports.NewTable("TWkContextMenu_New", 0), // static function New
			/* 5 */ imports.NewTable("TWkContextMenu_Prepend", 0), // procedure Prepend
			/* 6 */ imports.NewTable("TWkContextMenu_Append", 0), // procedure Append
			/* 7 */ imports.NewTable("TWkContextMenu_Insert", 0), // procedure Insert
			/* 8 */ imports.NewTable("TWkContextMenu_Remove", 0), // procedure Remove
			/* 9 */ imports.NewTable("TWkContextMenu_RemoveAll", 0), // procedure RemoveAll
		}
	})
	return wkContextMenuImport
}
