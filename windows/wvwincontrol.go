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

// IWVWinControl Parent: lcl.IWinControl
type IWVWinControl interface {
	lcl.IWinControl
	// CreateHandle
	//  Creates underlying screen object.
	CreateHandle() // procedure
	// InvalidateChildren
	//  Invalidates all child controls created by the browser.
	InvalidateChildren() // procedure
	// UpdateSize
	//  Updates the size of the child controls created by the browser.
	UpdateSize() // procedure
	// ChildWindowHandle
	//  Handle of the first child control created by the browser.
	ChildWindowHandle() types.THandle  // property ChildWindowHandle Getter
	DragKind() types.TDragKind         // property DragKind Getter
	SetDragKind(value types.TDragKind) // property DragKind Setter
	DragCursor() types.TCursor         // property DragCursor Getter
	SetDragCursor(value types.TCursor) // property DragCursor Setter
	DragMode() types.TDragMode         // property DragMode Getter
	SetDragMode(value types.TDragMode) // property DragMode Setter
	SetOnDragDrop(fn TDragDropEvent)   // property event
	SetOnDragOver(fn TDragOverEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent) // property event
	SetOnEndDrag(fn TEndDragEvent)     // property event
}

type TWVWinControl struct {
	lcl.TWinControl
}

func (m *TWVWinControl) CreateHandle() {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(2, m.Instance())
}

func (m *TWVWinControl) InvalidateChildren() {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(3, m.Instance())
}

func (m *TWVWinControl) UpdateSize() {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(4, m.Instance())
}

func (m *TWVWinControl) ChildWindowHandle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := wVWinControlAPI().SysCallN(5, m.Instance())
	return types.THandle(r)
}

func (m *TWVWinControl) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := wVWinControlAPI().SysCallN(6, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TWVWinControl) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TWVWinControl) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := wVWinControlAPI().SysCallN(7, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TWVWinControl) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TWVWinControl) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := wVWinControlAPI().SysCallN(8, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TWVWinControl) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	wVWinControlAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TWVWinControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, wVWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVWinControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, wVWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVWinControl) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 11, wVWinControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWVWinControl) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, wVWinControlAPI(), api.MakeEventDataPtr(cb))
}

// NewWinControl class constructor
func NewWinControl(theOwner lcl.IComponent) IWVWinControl {
	r := wVWinControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsWVWinControl(r)
}

// NewWinControlParented class constructor
func NewWinControlParented(parentWindow types.HWND) IWVWinControl {
	r := wVWinControlAPI().SysCallN(1, uintptr(parentWindow))
	return AsWVWinControl(r)
}

var (
	wVWinControlOnce   base.Once
	wVWinControlImport *imports.Imports = nil
)

func wVWinControlAPI() *imports.Imports {
	wVWinControlOnce.Do(func() {
		wVWinControlImport = api.NewDefaultImports()
		wVWinControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVWinControl_Create", 0), // constructor NewWinControl
			/* 1 */ imports.NewTable("TWVWinControl_CreateParented", 0), // constructor NewWinControlParented
			/* 2 */ imports.NewTable("TWVWinControl_CreateHandle", 0), // procedure CreateHandle
			/* 3 */ imports.NewTable("TWVWinControl_InvalidateChildren", 0), // procedure InvalidateChildren
			/* 4 */ imports.NewTable("TWVWinControl_UpdateSize", 0), // procedure UpdateSize
			/* 5 */ imports.NewTable("TWVWinControl_ChildWindowHandle", 0), // property ChildWindowHandle
			/* 6 */ imports.NewTable("TWVWinControl_DragKind", 0), // property DragKind
			/* 7 */ imports.NewTable("TWVWinControl_DragCursor", 0), // property DragCursor
			/* 8 */ imports.NewTable("TWVWinControl_DragMode", 0), // property DragMode
			/* 9 */ imports.NewTable("TWVWinControl_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TWVWinControl_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TWVWinControl_OnStartDrag", 0), // event OnStartDrag
			/* 12 */ imports.NewTable("TWVWinControl_OnEndDrag", 0), // event OnEndDrag
		}
	})
	return wVWinControlImport
}
