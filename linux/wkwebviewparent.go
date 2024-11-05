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
)

// IWkWebViewParent Parent: ICustomPanel
type IWkWebViewParent interface {
	ICustomPanel
	ScrolledWindow() TScrolledWindow          // property
	SetScrolledWindow(AValue TScrolledWindow) // property
	WebView() IWkWebview                      // property
	SetWebView(AValue IWkWebview)             // property
	FreeChild()                               // procedure
}

// TWkWebViewParent Parent: TCustomPanel
type TWkWebViewParent struct {
	TCustomPanel
}

func NewWkWebViewParent(aOwner IComponent) IWkWebViewParent {
	r1 := wkWebViewParentImportAPI().SysCallN(0, GetObjectUintptr(aOwner))
	return AsWkWebViewParent(r1)
}

func (m *TWkWebViewParent) ScrolledWindow() TScrolledWindow {
	r1 := wkWebViewParentImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TScrolledWindow(r1)
}

func (m *TWkWebViewParent) SetScrolledWindow(AValue TScrolledWindow) {
	wkWebViewParentImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TWkWebViewParent) WebView() IWkWebview {
	r1 := wkWebViewParentImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsWkWebview(r1)
}

func (m *TWkWebViewParent) SetWebView(AValue IWkWebview) {
	wkWebViewParentImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWkWebViewParent) FreeChild() {
	wkWebViewParentImportAPI().SysCallN(1, m.Instance())
}

var (
	wkWebViewParentImport       *imports.Imports = nil
	wkWebViewParentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebViewParent_Create", 0),
		/*1*/ imports.NewTable("WkWebViewParent_FreeChild", 0),
		/*2*/ imports.NewTable("WkWebViewParent_ScrolledWindow", 0),
		/*3*/ imports.NewTable("WkWebViewParent_WebView", 0),
	}
)

func wkWebViewParentImportAPI() *imports.Imports {
	if wkWebViewParentImport == nil {
		wkWebViewParentImport = NewDefaultImports()
		wkWebViewParentImport.SetImportTable(wkWebViewParentImportTables)
		wkWebViewParentImportTables = nil
	}
	return wkWebViewParentImport
}
