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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2ContextMenuItem Parent: IObject
type ICoreWebView2ContextMenuItem interface {
	IObject
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(browserComponent lcl.IComponent) bool // function
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuItem         // property BaseIntf Getter
	SetBaseIntf(value ICoreWebView2ContextMenuItem) // property BaseIntf Setter
	// Name
	//  Gets the unlocalized name for the `ContextMenuItem`. Use this to
	//  distinguish between context menu item types. This will be the English
	//  label of the menu item in lower camel case. For example, the "Save as"
	//  menu item will be "saveAs". Extension menu items will be "extension",
	//  custom menu items will be "custom" and spellcheck items will be
	//  "spellCheck".
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_name">See the ICoreWebView2ContextMenuItem article.</see>
	Name() string // property Name Getter
	// Label
	//  Gets the localized label for the `ContextMenuItem`. Will contain an
	//  ampersand for characters to be used as keyboard accelerator.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_label">See the ICoreWebView2ContextMenuItem article.</see>
	Label() string // property Label_ Getter
	// CommandId
	//  Gets the Command ID for the `ContextMenuItem`. Use this to report the
	//  `SelectedCommandId` in `ContextMenuRequested` event.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_commandid">See the ICoreWebView2ContextMenuItem article.</see>
	CommandId() int32 // property CommandId Getter
	// ShortcutKeyDescription
	//  Gets the localized keyboard shortcut for this ContextMenuItem. It will be
	//  the empty string if there is no keyboard shortcut. This is text intended
	//  to be displayed to the end user to show the keyboard shortcut. For example
	//  this property is Ctrl+Shift+I for the "Inspect" `ContextMenuItem`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_shortcutkeydescription">See the ICoreWebView2ContextMenuItem article.</see>
	ShortcutKeyDescription() string // property ShortcutKeyDescription Getter
	// Icon
	//  Gets the Icon for the `ContextMenuItem` in PNG, Bitmap or SVG formats in the form of an IStream.
	//  Stream will be rewound to the start of the image data.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_icon">See the ICoreWebView2ContextMenuItem article.</see>
	Icon() lcl.IStreamAdapter // property Icon Getter
	// Kind
	//  Gets the `ContextMenuItem` kind.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_kind">See the ICoreWebView2ContextMenuItem article.</see>
	Kind() wvTypes.TWVMenuItemKind // property Kind Getter
	// IsEnabled
	//  Gets the enabled property of the `ContextMenuItem`.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_isenabled">See the ICoreWebView2ContextMenuItem article.</see>
	IsEnabled() bool         // property IsEnabled Getter
	SetIsEnabled(value bool) // property IsEnabled Setter
	// IsChecked
	//  Gets the checked property of the `ContextMenuItem`, used if the kind is Check box or Radio.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_ischecked">See the ICoreWebView2ContextMenuItem article.</see>
	IsChecked() bool         // property IsChecked Getter
	SetIsChecked(value bool) // property IsChecked Setter
	// Children
	//  Gets the list of children menu items through a `ContextMenuItemCollection`
	//  if the kind is Submenu. If the kind is not submenu, will return null.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#get_children">See the ICoreWebView2ContextMenuItem article.</see>
	Children() ICoreWebView2ContextMenuItemCollection // property Children Getter
}

type TCoreWebView2ContextMenuItem struct {
	TObject
}

func (m *TCoreWebView2ContextMenuItem) AddAllBrowserEvents(browserComponent lcl.IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(browserComponent))
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItem) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItem) BaseIntf() (result ICoreWebView2ContextMenuItem) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuItemAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuItem(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuItem) SetBaseIntf(value ICoreWebView2ContextMenuItem) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoreWebView2ContextMenuItem) Name() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuItemAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuItem) Label() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuItemAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuItem) CommandId() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCoreWebView2ContextMenuItem) ShortcutKeyDescription() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuItemAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuItem) Icon() lcl.IStreamAdapter {
	if !m.IsValid() {
		return nil
	}
	var resultPtr uintptr
	coreWebView2ContextMenuItemAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	return lcl.AsStreamAdapter(resultPtr)
}

func (m *TCoreWebView2ContextMenuItem) Kind() wvTypes.TWVMenuItemKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(9, m.Instance())
	return wvTypes.TWVMenuItemKind(r)
}

func (m *TCoreWebView2ContextMenuItem) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItem) SetIsEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ContextMenuItem) IsChecked() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuItemAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuItem) SetIsChecked(value bool) {
	if !m.IsValid() {
		return
	}
	coreWebView2ContextMenuItemAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoreWebView2ContextMenuItem) Children() (result ICoreWebView2ContextMenuItemCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuItemAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuItemCollection(resultPtr)
	return
}

// NewCoreWebView2ContextMenuItem class constructor
func NewCoreWebView2ContextMenuItem(baseIntf ICoreWebView2ContextMenuItem) ICoreWebView2ContextMenuItem {
	r := coreWebView2ContextMenuItemAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ContextMenuItem(r)
}

var (
	coreWebView2ContextMenuItemOnce   base.Once
	coreWebView2ContextMenuItemImport *imports.Imports = nil
)

func coreWebView2ContextMenuItemAPI() *imports.Imports {
	coreWebView2ContextMenuItemOnce.Do(func() {
		coreWebView2ContextMenuItemImport = api.NewDefaultImports()
		coreWebView2ContextMenuItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ContextMenuItem_Create", 0), // constructor NewCoreWebView2ContextMenuItem
			/* 1 */ imports.NewTable("TCoreWebView2ContextMenuItem_AddAllBrowserEvents", 0), // function AddAllBrowserEvents
			/* 2 */ imports.NewTable("TCoreWebView2ContextMenuItem_Initialized", 0), // property Initialized
			/* 3 */ imports.NewTable("TCoreWebView2ContextMenuItem_BaseIntf", 0), // property BaseIntf
			/* 4 */ imports.NewTable("TCoreWebView2ContextMenuItem_Name", 0), // property Name
			/* 5 */ imports.NewTable("TCoreWebView2ContextMenuItem_Label_", 0), // property Label
			/* 6 */ imports.NewTable("TCoreWebView2ContextMenuItem_CommandId", 0), // property CommandId
			/* 7 */ imports.NewTable("TCoreWebView2ContextMenuItem_ShortcutKeyDescription", 0), // property ShortcutKeyDescription
			/* 8 */ imports.NewTable("TCoreWebView2ContextMenuItem_Icon", 0), // property Icon
			/* 9 */ imports.NewTable("TCoreWebView2ContextMenuItem_Kind", 0), // property Kind
			/* 10 */ imports.NewTable("TCoreWebView2ContextMenuItem_IsEnabled", 0), // property IsEnabled
			/* 11 */ imports.NewTable("TCoreWebView2ContextMenuItem_IsChecked", 0), // property IsChecked
			/* 12 */ imports.NewTable("TCoreWebView2ContextMenuItem_Children", 0), // property Children
		}
	})
	return coreWebView2ContextMenuItemImport
}
