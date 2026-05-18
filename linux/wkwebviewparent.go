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
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkWebviewParent Parent: ICustomPanel
type IWkWebviewParent interface {
	ICustomPanel
	UpdateSize(width int32, height int32)            // procedure
	FreeChild()                                      // procedure
	Webview() IWkWebview                             // property Webview Getter
	SetWebview(value IWkWebview)                     // property Webview Setter
	ScrolledWindow() wvTypes.TScrolledWindow         // property ScrolledWindow Getter
	SetScrolledWindow(value wvTypes.TScrolledWindow) // property ScrolledWindow Setter
}

type TWkWebviewParent struct {
	TCustomPanel
}

func (m *TWkWebviewParent) UpdateSize(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(1, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TWkWebviewParent) FreeChild() {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(2, m.Instance())
}

func (m *TWkWebviewParent) Webview() IWkWebview {
	if !m.IsValid() {
		return nil
	}
	r := wkWebviewParentAPI().SysCallN(3, 0, m.Instance())
	return AsWkWebview(r)
}

func (m *TWkWebviewParent) SetWebview(value IWkWebview) {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TWkWebviewParent) ScrolledWindow() wvTypes.TScrolledWindow {
	if !m.IsValid() {
		return 0
	}
	r := wkWebviewParentAPI().SysCallN(4, 0, m.Instance())
	return wvTypes.TScrolledWindow(r)
}

func (m *TWkWebviewParent) SetScrolledWindow(value wvTypes.TScrolledWindow) {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

// NewWebviewParent class constructor
func NewWebviewParent(owner lcl.IComponent) IWkWebviewParent {
	r := wkWebviewParentAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsWkWebviewParent(r)
}

var (
	wkWebviewParentOnce   base.Once
	wkWebviewParentImport *imports.Imports = nil
)

func wkWebviewParentAPI() *imports.Imports {
	wkWebviewParentOnce.Do(func() {
		wkWebviewParentImport = api.NewDefaultImports()
		wkWebviewParentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebviewParent_Create", 0), // constructor NewWebviewParent
			/* 1 */ imports.NewTable("TWkWebviewParent_UpdateSize", 0), // procedure UpdateSize
			/* 2 */ imports.NewTable("TWkWebviewParent_FreeChild", 0), // procedure FreeChild
			/* 3 */ imports.NewTable("TWkWebviewParent_Webview", 0), // property Webview
			/* 4 */ imports.NewTable("TWkWebviewParent_ScrolledWindow", 0), // property ScrolledWindow
		}
	})
	return wkWebviewParentImport
}
