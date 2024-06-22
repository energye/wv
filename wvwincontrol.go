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
	. "github.com/energye/lcl/types"
)

// IWVWinControl Parent: IWinControl
//
//	Parent control that will be used by TWVWindowParent in VCL and LCL applications to show the web contents.
type IWVWinControl interface {
	IWinControl
	// ChildWindowHandle
	//  Handle of the first child control created by the browser.
	ChildWindowHandle() THandle   // property
	DragKind() TDragKind          // property
	SetDragKind(AValue TDragKind) // property
	DragCursor() TCursor          // property
	SetDragCursor(AValue TCursor) // property
	DragMode() TDragMode          // property
	SetDragMode(AValue TDragMode) // property
	// CreateHandle
	//  Creates underlying screen object.
	CreateHandle() // procedure
	// InvalidateChildren
	//  Invalidates all child controls created by the browser.
	InvalidateChildren() // procedure
	// UpdateSize
	//  Updates the size of the child controls created by the browser.
	UpdateSize()                       // procedure
	SetOnDragDrop(fn TDragDropEvent)   // property event
	SetOnDragOver(fn TDragOverEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent) // property event
	SetOnEndDrag(fn TEndDragEvent)     // property event
}

// TWVWinControl Parent: TWinControl
//
//	Parent control that will be used by TWVWindowParent in VCL and LCL applications to show the web contents.
type TWVWinControl struct {
	TWinControl
	dragDropPtr  uintptr
	dragOverPtr  uintptr
	startDragPtr uintptr
	endDragPtr   uintptr
}

func NewWVWinControl(TheOwner IComponent) IWVWinControl {
	r1 := wVWinControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsWVWinControl(r1)
}

func (m *TWVWinControl) ChildWindowHandle() THandle {
	r1 := wVWinControlImportAPI().SysCallN(0, m.Instance())
	return THandle(r1)
}

func (m *TWVWinControl) DragKind() TDragKind {
	r1 := wVWinControlImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TWVWinControl) SetDragKind(AValue TDragKind) {
	wVWinControlImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVWinControl) DragCursor() TCursor {
	r1 := wVWinControlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TWVWinControl) SetDragCursor(AValue TCursor) {
	wVWinControlImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVWinControl) DragMode() TDragMode {
	r1 := wVWinControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TWVWinControl) SetDragMode(AValue TDragMode) {
	wVWinControlImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVWinControl) CreateHandle() {
	wVWinControlImportAPI().SysCallN(2, m.Instance())
}

func (m *TWVWinControl) InvalidateChildren() {
	wVWinControlImportAPI().SysCallN(6, m.Instance())
}

func (m *TWVWinControl) UpdateSize() {
	wVWinControlImportAPI().SysCallN(11, m.Instance())
}

func (m *TWVWinControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	wVWinControlImportAPI().SysCallN(7, m.Instance(), m.dragDropPtr)
}

func (m *TWVWinControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	wVWinControlImportAPI().SysCallN(8, m.Instance(), m.dragOverPtr)
}

func (m *TWVWinControl) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	wVWinControlImportAPI().SysCallN(10, m.Instance(), m.startDragPtr)
}

func (m *TWVWinControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	wVWinControlImportAPI().SysCallN(9, m.Instance(), m.endDragPtr)
}

var (
	wVWinControlImport       *imports.Imports = nil
	wVWinControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WVWinControl_ChildWindowHandle", 0),
		/*1*/ imports.NewTable("WVWinControl_Create", 0),
		/*2*/ imports.NewTable("WVWinControl_CreateHandle", 0),
		/*3*/ imports.NewTable("WVWinControl_DragCursor", 0),
		/*4*/ imports.NewTable("WVWinControl_DragKind", 0),
		/*5*/ imports.NewTable("WVWinControl_DragMode", 0),
		/*6*/ imports.NewTable("WVWinControl_InvalidateChildren", 0),
		/*7*/ imports.NewTable("WVWinControl_SetOnDragDrop", 0),
		/*8*/ imports.NewTable("WVWinControl_SetOnDragOver", 0),
		/*9*/ imports.NewTable("WVWinControl_SetOnEndDrag", 0),
		/*10*/ imports.NewTable("WVWinControl_SetOnStartDrag", 0),
		/*11*/ imports.NewTable("WVWinControl_UpdateSize", 0),
	}
)

func wVWinControlImportAPI() *imports.Imports {
	if wVWinControlImport == nil {
		wVWinControlImport = NewDefaultImports()
		wVWinControlImport.SetImportTable(wVWinControlImportTables)
		wVWinControlImportTables = nil
	}
	return wVWinControlImport
}
