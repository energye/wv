//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkWebviewParent Parent: lcl.ICustomControl
type IWkWebviewParent interface {
	lcl.ICustomControl
	Resize()                              // procedure
	SetWebview(webview wvTypes.WKWebView) // procedure
}

type TWkWebviewParent struct {
	lcl.TCustomControl
}

func (m *TWkWebviewParent) Resize() {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(1, m.Instance())
}

func (m *TWkWebviewParent) SetWebview(webview wvTypes.WKWebView) {
	if !m.IsValid() {
		return
	}
	wkWebviewParentAPI().SysCallN(2, m.Instance(), uintptr(webview))
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
			/* 1 */ imports.NewTable("TWkWebviewParent_Resize", 0), // procedure Resize
			/* 2 */ imports.NewTable("TWkWebviewParent_SetWebview", 0), // procedure SetWebview
		}
	})
	return wkWebviewParentImport
}
