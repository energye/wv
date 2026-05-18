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

	wvTypes "github.com/energye/wv/types/windows"
)

// ICoreWebView2ContextMenuTarget Parent: IObject
type ICoreWebView2ContextMenuTarget interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ContextMenuTarget // property BaseIntf Getter
	// Kind
	//  Gets the kind of context that the user selected.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_kind">See the ICoreWebView2ContextMenuTarget article.</see>
	Kind() wvTypes.TWVMenuTargetKind // property Kind Getter
	// IsEditable
	//  Returns TRUE if the context menu is requested on an editable component.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_iseditable">See the ICoreWebView2ContextMenuTarget article.</see>
	IsEditable() bool // property IsEditable Getter
	// IsRequestedForMainFrame
	//  Returns TRUE if the context menu was requested on the main frame and
	//  FALSE if invoked on another frame.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_isrequestedformainframe">See the ICoreWebView2ContextMenuTarget article.</see>
	IsRequestedForMainFrame() bool // property IsRequestedForMainFrame Getter
	// PageUri
	//  Gets the uri of the page.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_pageuri">See the ICoreWebView2ContextMenuTarget article.</see>
	PageUri() string // property PageUri Getter
	// FrameUri
	//  Gets the uri of the frame. Will match the PageUri if `IsRequestedForMainFrame` is TRUE.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_frameuri">See the ICoreWebView2ContextMenuTarget article.</see>
	FrameUri() string // property FrameUri Getter
	// HasLinkUri
	//  Returns TRUE if the context menu is requested on HTML containing an anchor tag.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_haslinkuri">See the ICoreWebView2ContextMenuTarget article.</see>
	HasLinkUri() bool // property HasLinkUri Getter
	// LinkUri
	//  Gets the uri of the link (if `HasLinkUri` is TRUE, null otherwise).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_linkuri">See the ICoreWebView2ContextMenuTarget article.</see>
	LinkUri() string // property LinkUri Getter
	// HasLinkText
	//  Returns TRUE if the context menu is requested on text element that contains an anchor tag.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_haslinktext">See the ICoreWebView2ContextMenuTarget article.</see>
	HasLinkText() bool // property HasLinkText Getter
	// LinkText
	//  Gets the text of the link (if `HasLinkText` is TRUE, null otherwise).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_linktext">See the ICoreWebView2ContextMenuTarget article.</see>
	LinkText() string // property LinkText Getter
	// HasSourceUri
	//  Returns TRUE if the context menu is requested on HTML containing a source uri.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_hassourceuri">See the ICoreWebView2ContextMenuTarget article.</see>
	HasSourceUri() bool // property HasSourceUri Getter
	// SourceUri
	//  Gets the active source uri of element (if `HasSourceUri` is TRUE, null otherwise).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_sourceuri">See the ICoreWebView2ContextMenuTarget article.</see>
	SourceUri() string // property SourceUri Getter
	// HasSelection
	//  Returns TRUE if the context menu is requested on a selection.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_hasselection">See the ICoreWebView2ContextMenuTarget article.</see>
	HasSelection() bool // property HasSelection Getter
	// SelectionText
	//  Gets the selected text (if `HasSelection` is TRUE, null otherwise).
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget#get_selectiontext">See the ICoreWebView2ContextMenuTarget article.</see>
	SelectionText() string // property SelectionText Getter
}

type TCoreWebView2ContextMenuTarget struct {
	TObject
}

func (m *TCoreWebView2ContextMenuTarget) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) BaseIntf() (result ICoreWebView2ContextMenuTarget) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2ContextMenuTargetAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2ContextMenuTarget(resultPtr)
	return
}

func (m *TCoreWebView2ContextMenuTarget) Kind() wvTypes.TWVMenuTargetKind {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(3, m.Instance())
	return wvTypes.TWVMenuTargetKind(r)
}

func (m *TCoreWebView2ContextMenuTarget) IsEditable() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) IsRequestedForMainFrame() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) PageUri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuTarget) FrameUri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuTarget) HasLinkUri() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) LinkUri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuTarget) HasLinkText() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) LinkText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuTarget) HasSourceUri() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) SourceUri() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TCoreWebView2ContextMenuTarget) HasSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ContextMenuTargetAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2ContextMenuTarget) SelectionText() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	coreWebView2ContextMenuTargetAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

// NewCoreWebView2ContextMenuTarget class constructor
func NewCoreWebView2ContextMenuTarget(baseIntf ICoreWebView2ContextMenuTarget) ICoreWebView2ContextMenuTarget {
	r := coreWebView2ContextMenuTargetAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ContextMenuTarget(r)
}

var (
	coreWebView2ContextMenuTargetOnce   base.Once
	coreWebView2ContextMenuTargetImport *imports.Imports = nil
)

func coreWebView2ContextMenuTargetAPI() *imports.Imports {
	coreWebView2ContextMenuTargetOnce.Do(func() {
		coreWebView2ContextMenuTargetImport = api.NewDefaultImports()
		coreWebView2ContextMenuTargetImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ContextMenuTarget_Create", 0), // constructor NewCoreWebView2ContextMenuTarget
			/* 1 */ imports.NewTable("TCoreWebView2ContextMenuTarget_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2ContextMenuTarget_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2ContextMenuTarget_Kind", 0), // property Kind
			/* 4 */ imports.NewTable("TCoreWebView2ContextMenuTarget_IsEditable", 0), // property IsEditable
			/* 5 */ imports.NewTable("TCoreWebView2ContextMenuTarget_IsRequestedForMainFrame", 0), // property IsRequestedForMainFrame
			/* 6 */ imports.NewTable("TCoreWebView2ContextMenuTarget_PageUri", 0), // property PageUri
			/* 7 */ imports.NewTable("TCoreWebView2ContextMenuTarget_FrameUri", 0), // property FrameUri
			/* 8 */ imports.NewTable("TCoreWebView2ContextMenuTarget_HasLinkUri", 0), // property HasLinkUri
			/* 9 */ imports.NewTable("TCoreWebView2ContextMenuTarget_LinkUri", 0), // property LinkUri
			/* 10 */ imports.NewTable("TCoreWebView2ContextMenuTarget_HasLinkText", 0), // property HasLinkText
			/* 11 */ imports.NewTable("TCoreWebView2ContextMenuTarget_LinkText", 0), // property LinkText
			/* 12 */ imports.NewTable("TCoreWebView2ContextMenuTarget_HasSourceUri", 0), // property HasSourceUri
			/* 13 */ imports.NewTable("TCoreWebView2ContextMenuTarget_SourceUri", 0), // property SourceUri
			/* 14 */ imports.NewTable("TCoreWebView2ContextMenuTarget_HasSelection", 0), // property HasSelection
			/* 15 */ imports.NewTable("TCoreWebView2ContextMenuTarget_SelectionText", 0), // property SelectionText
		}
	})
	return coreWebView2ContextMenuTargetImport
}
