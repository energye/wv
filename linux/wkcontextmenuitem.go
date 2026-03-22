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

// IWkContextMenuItem Parent: IObject
type IWkContextMenuItem interface {
	IObject
	Data() wvTypes.WebKitContextMenuItem             // function
	GetSubmenu() wvTypes.WebKitContextMenu           // function
	GetAction() wvTypes.PWkAction                    // function
	GetStockAction() wvTypes.WebKitContextMenuAction // function
	IsSeparator() bool                               // function
	SetSubmenu(submenu wvTypes.WebKitContextMenu)    // procedure
}

type TWkContextMenuItem struct {
	TObject
}

func (m *TWkContextMenuItem) Data() wvTypes.WebKitContextMenuItem {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuItemAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitContextMenuItem(r)
}

func (m *TWkContextMenuItem) GetSubmenu() wvTypes.WebKitContextMenu {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuItemAPI().SysCallN(2, m.Instance())
	return wvTypes.WebKitContextMenu(r)
}

func (m *TWkContextMenuItem) GetAction() wvTypes.PWkAction {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuItemAPI().SysCallN(3, m.Instance())
	return wvTypes.PWkAction(r)
}

func (m *TWkContextMenuItem) GetStockAction() wvTypes.WebKitContextMenuAction {
	if !m.IsValid() {
		return 0
	}
	r := wkContextMenuItemAPI().SysCallN(4, m.Instance())
	return wvTypes.WebKitContextMenuAction(r)
}

func (m *TWkContextMenuItem) IsSeparator() bool {
	if !m.IsValid() {
		return false
	}
	r := wkContextMenuItemAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkContextMenuItem) SetSubmenu(submenu wvTypes.WebKitContextMenu) {
	if !m.IsValid() {
		return
	}
	wkContextMenuItemAPI().SysCallN(11, m.Instance(), uintptr(submenu))
}

// ContextMenuItem  is static instance
var ContextMenuItem _ContextMenuItemClass

// _ContextMenuItemClass is class type defined by TWkContextMenuItem
type _ContextMenuItemClass uintptr

func (_ContextMenuItemClass) NewFromAction(action wvTypes.PWkAction, label string, commandId int32) IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(6, uintptr(action), api.PasStr(label), uintptr(commandId))
	return AsWkContextMenuItem(r)
}

func (_ContextMenuItemClass) NewFromStockAction(action wvTypes.WebKitContextMenuAction) IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(7, uintptr(action))
	return AsWkContextMenuItem(r)
}

func (_ContextMenuItemClass) NewFromStockActionWithLabel(action wvTypes.WebKitContextMenuAction, label string) IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(8, uintptr(action), api.PasStr(label))
	return AsWkContextMenuItem(r)
}

func (_ContextMenuItemClass) NewWithSubmenu(label string, submenu wvTypes.WebKitContextMenu) IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(9, api.PasStr(label), uintptr(submenu))
	return AsWkContextMenuItem(r)
}

func (_ContextMenuItemClass) NewSeparator() IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(10)
	return AsWkContextMenuItem(r)
}

// NewContextMenuItem class constructor
func NewContextMenuItem(contextMenu wvTypes.WebKitContextMenuItem) IWkContextMenuItem {
	r := wkContextMenuItemAPI().SysCallN(0, uintptr(contextMenu))
	return AsWkContextMenuItem(r)
}

var (
	wkContextMenuItemOnce   base.Once
	wkContextMenuItemImport *imports.Imports = nil
)

func wkContextMenuItemAPI() *imports.Imports {
	wkContextMenuItemOnce.Do(func() {
		wkContextMenuItemImport = api.NewDefaultImports()
		wkContextMenuItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkContextMenuItem_Create", 0), // constructor NewContextMenuItem
			/* 1 */ imports.NewTable("TWkContextMenuItem_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkContextMenuItem_GetSubmenu", 0), // function GetSubmenu
			/* 3 */ imports.NewTable("TWkContextMenuItem_GetAction", 0), // function GetAction
			/* 4 */ imports.NewTable("TWkContextMenuItem_GetStockAction", 0), // function GetStockAction
			/* 5 */ imports.NewTable("TWkContextMenuItem_IsSeparator", 0), // function IsSeparator
			/* 6 */ imports.NewTable("TWkContextMenuItem_NewFromAction", 0), // static function NewFromAction
			/* 7 */ imports.NewTable("TWkContextMenuItem_NewFromStockAction", 0), // static function NewFromStockAction
			/* 8 */ imports.NewTable("TWkContextMenuItem_NewFromStockActionWithLabel", 0), // static function NewFromStockActionWithLabel
			/* 9 */ imports.NewTable("TWkContextMenuItem_NewWithSubmenu", 0), // static function NewWithSubmenu
			/* 10 */ imports.NewTable("TWkContextMenuItem_NewSeparator", 0), // static function NewSeparator
			/* 11 */ imports.NewTable("TWkContextMenuItem_SetSubmenu", 0), // procedure SetSubmenu
		}
	})
	return wkContextMenuItemImport
}
