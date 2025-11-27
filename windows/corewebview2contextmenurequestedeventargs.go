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

// ICoreWebView2ContextMenuRequestedEventArgs Parent: lcl.IObject
type ICoreWebView2ContextMenuRequestedEventArgs interface {
	lcl.IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuRequestedEventArgs // property BaseIntf Getter
	// MenuItems
	//  Gets the collection of `ContextMenuItem` objects.
	//  See `ICoreWebView2ContextMenuItemCollection` for more details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_menuitems">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	MenuItems() ICoreWebView2ContextMenuItemCollection // property MenuItems Getter
	// ContextMenuTarget
	//  Gets the target information associated with the requested context menu.
	//  See `ICoreWebView2ContextMenuTarget` for more details.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_contextmenutarget">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	ContextMenuTarget() ICoreWebView2ContextMenuTarget // property ContextMenuTarget Getter
	// Location
	//  Gets the coordinates where the context menu request occurred in relation to the upper
	//  left corner of the WebView bounds.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_location">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	Location() types.TPoint // property Location Getter
	// SelectedCommandId
	//  Gets the selected CommandId.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_selectedcommandid">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	SelectedCommandId() int32         // property SelectedCommandId Getter
	SetSelectedCommandId(value int32) // property SelectedCommandId Setter
	// Handled
	//  Gets whether the `ContextMenuRequested` event is handled by host.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#get_handled">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	Handled() bool         // property Handled Getter
	SetHandled(value bool) // property Handled Setter
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event when the custom context menu is closed.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenurequestedeventargs#getdeferral">See the ICoreWebView2ContextMenuRequestedEventArgs article.</see>
	Deferral() ICoreWebView2Deferral // property Deferral Getter
}

type TCoreWebView2ContextMenuRequestedEventArgs struct {
	lcl.TObject
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) BaseIntf() (result ICoreWebView2ContextMenuRequestedEventArgs) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuRequestedEventArgs(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) MenuItems() (result ICoreWebView2ContextMenuItemCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuItemCollection(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) ContextMenuTarget() (result ICoreWebView2ContextMenuTarget) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuTarget(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Location() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SelectedCommandId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SetSelectedCommandId(value int32) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Handled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) SetHandled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ContextMenuRequestedEventArgs) Deferral() (result ICoreWebView2Deferral) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2Deferral(resultPtr)
	return
}

// NewCoreWebView2ContextMenuRequestedEventArgs class constructor
func NewCoreWebView2ContextMenuRequestedEventArgs(args ICoreWebView2ContextMenuRequestedEventArgs) ICoreWebView2ContextMenuRequestedEventArgs {
	r := coreWebView2ContextMenuRequestedEventArgsAPI().SysCallN(0, base.GetObjectUintptr(args))
	return AsCoreWebView2ContextMenuRequestedEventArgs(r)
}

var (
	coreWebView2ContextMenuRequestedEventArgsOnce   base.Once
	coreWebView2ContextMenuRequestedEventArgsImport *imports.Imports = nil
)

func coreWebView2ContextMenuRequestedEventArgsAPI() *imports.Imports {
	coreWebView2ContextMenuRequestedEventArgsOnce.Do(func() {
		coreWebView2ContextMenuRequestedEventArgsImport = api.NewDefaultImports()
		coreWebView2ContextMenuRequestedEventArgsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_Create", 0), // constructor NewCoreWebView2ContextMenuRequestedEventArgs
			/* 1 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_MenuItems", 0), // property MenuItems
			/* 4 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_ContextMenuTarget", 0), // property ContextMenuTarget
			/* 5 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_Location", 0), // property Location
			/* 6 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_SelectedCommandId", 0), // property SelectedCommandId
			/* 7 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_Handled", 0), // property Handled
			/* 8 */ imports.NewTable("TCoreWebView2ContextMenuRequestedEventArgs_Deferral", 0), // property Deferral
		}
	})
	return coreWebView2ContextMenuRequestedEventArgsImport
}
