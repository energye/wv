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

// ICoreWebView2ContextMenuItemCollection Parent: lcl.IObject
type ICoreWebView2ContextMenuItemCollection interface {
	lcl.IObject
	// RemoveValueAtIndex
	//  Removes the `ContextMenuItem` at the specified index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	RemoveValueAtIndex(index uint32) bool // function
	// InsertValueAtIndex
	//  Inserts the `ContextMenuItem` at the specified index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#insertvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	InsertValueAtIndex(index uint32, value ICoreWebView2ContextMenuItem) bool // function
	// AppendValue
	//  Appends the aValue item at the end of the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#insertvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	AppendValue(value ICoreWebView2ContextMenuItem) bool // function
	// RemoveAllMenuItems
	//  Removes all items from the collection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	RemoveAllMenuItems() // procedure
	// RemoveMenuItemWithInt
	//  Removes the item with the commandId value specified in the paramaters.
	//  <param name="aCommandId">The commandId value of the item that has to be removed.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	RemoveMenuItemWithInt(commandId int32) // procedure
	// RemoveMenuItemWithString
	//  Removes the item with the label value specified in the paramaters.
	//  <param name="aLabel">The label value of the item that has to be removed.</param>
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#removevalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	RemoveMenuItemWithString(label string) // procedure
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuItemCollection // property BaseIntf Getter
	// Count
	//  Gets the number of `ContextMenuItem` objects contained in the `ContextMenuItemCollection`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#get_count">See the ICoreWebView2ContextMenuItemCollection article.</see>
	Count() uint32 // property Count Getter
	// Items
	//  Gets the `ContextMenuItem` at the specified index.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitemcollection#getvalueatindex">See the ICoreWebView2ContextMenuItemCollection article.</see>
	Items(idx uint32) ICoreWebView2ContextMenuItem // property Items Getter
}

type TCoreWebView2ContextMenuItemCollection struct {
	lcl.TObject
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(1, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, value ICoreWebView2ContextMenuItem) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(2, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItemCollection) AppendValue(value ICoreWebView2ContextMenuItem) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveAllMenuItems() {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemCollectionAPI().SysCallN(4, m.Instance())
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveMenuItemWithInt(commandId int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemCollectionAPI().SysCallN(5, m.Instance(), uintptr(commandId))
}

func (m *TCoreWebView2ContextMenuItemCollection) RemoveMenuItemWithString(label string) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemCollectionAPI().SysCallN(6, m.Instance(), api.PasStr(label))
}

func (m *TCoreWebView2ContextMenuItemCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItemCollection) BaseIntf() (result ICoreWebView2ContextMenuItemCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuItemCollectionAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuItemCollection(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuItemCollection) Count() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(9, m.Instance())
	return uint32(r)
}

func (m *TCoreWebView2ContextMenuItemCollection) Items(idx uint32) (result ICoreWebView2ContextMenuItem) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuItemCollectionAPI().SysCallN(10, m.Instance(), uintptr(idx), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuItem(resultPtr)
	return
}

// NewCoreWebView2ContextMenuItemCollection class constructor
func NewCoreWebView2ContextMenuItemCollection(baseIntf ICoreWebView2ContextMenuItemCollection) ICoreWebView2ContextMenuItemCollection {
	r := coreWebView2ContextMenuItemCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ContextMenuItemCollection(r)
}

var (
	coreWebView2ContextMenuItemCollectionOnce   base.Once
	coreWebView2ContextMenuItemCollectionImport *imports.Imports = nil
)

func coreWebView2ContextMenuItemCollectionAPI() *imports.Imports {
	coreWebView2ContextMenuItemCollectionOnce.Do(func() {
		coreWebView2ContextMenuItemCollectionImport = api.NewDefaultImports()
		coreWebView2ContextMenuItemCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_Create", 0), // constructor NewCoreWebView2ContextMenuItemCollection
			/* 1 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_RemoveValueAtIndex", 0), // function RemoveValueAtIndex
			/* 2 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_InsertValueAtIndex", 0), // function InsertValueAtIndex
			/* 3 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_AppendValue", 0), // function AppendValue
			/* 4 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_RemoveAllMenuItems", 0), // procedure RemoveAllMenuItems
			/* 5 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_RemoveMenuItemWithInt", 0), // procedure RemoveMenuItemWithInt
			/* 6 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_RemoveMenuItemWithString", 0), // procedure RemoveMenuItemWithString
			/* 7 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_Initialized", 0), // property Initialized
			/* 8 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_BaseIntf", 0), // property BaseIntf
			/* 9 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_Count", 0), // property Count
			/* 10 */ imports.NewTable("TCoreWebView2ContextMenuItemCollection_Items", 0), // property Items
		}
	})
	return coreWebView2ContextMenuItemCollectionImport
}
