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

// IWKNavigationAction Root Interface
type IWKNavigationAction interface {
	IObject
	Data() WKNavigationAction            // function
	SourceFrame() WKFrameInfo            // function
	TargetFrame() WKFrameInfo            // function
	NavigationType() WKNavigationType    // function
	Request() NSURLRequest               // function
	ModifierFlags() NSEventModifierFlags // function
	ButtonNumber() int32                 // function
}

// TWKNavigationAction Root Object
type TWKNavigationAction struct {
	TObject
}

func NewWKNavigationAction(aData WKNavigationAction) IWKNavigationAction {
	r1 := wKNavigationActionImportAPI().SysCallN(1, uintptr(aData))
	return AsWKNavigationAction(r1)
}

func (m *TWKNavigationAction) Data() WKNavigationAction {
	r1 := wKNavigationActionImportAPI().SysCallN(2, m.Instance())
	return WKNavigationAction(r1)
}

func (m *TWKNavigationAction) SourceFrame() WKFrameInfo {
	r1 := wKNavigationActionImportAPI().SysCallN(6, m.Instance())
	return WKFrameInfo(r1)
}

func (m *TWKNavigationAction) TargetFrame() WKFrameInfo {
	r1 := wKNavigationActionImportAPI().SysCallN(7, m.Instance())
	return WKFrameInfo(r1)
}

func (m *TWKNavigationAction) NavigationType() WKNavigationType {
	r1 := wKNavigationActionImportAPI().SysCallN(4, m.Instance())
	return WKNavigationType(r1)
}

func (m *TWKNavigationAction) Request() NSURLRequest {
	r1 := wKNavigationActionImportAPI().SysCallN(5, m.Instance())
	return NSURLRequest(r1)
}

func (m *TWKNavigationAction) ModifierFlags() NSEventModifierFlags {
	r1 := wKNavigationActionImportAPI().SysCallN(3, m.Instance())
	return NSEventModifierFlags(r1)
}

func (m *TWKNavigationAction) ButtonNumber() int32 {
	r1 := wKNavigationActionImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

var (
	wKNavigationActionImport       *imports.Imports = nil
	wKNavigationActionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKNavigationAction_ButtonNumber", 0),
		/*1*/ imports.NewTable("WKNavigationAction_Create", 0),
		/*2*/ imports.NewTable("WKNavigationAction_Data", 0),
		/*3*/ imports.NewTable("WKNavigationAction_ModifierFlags", 0),
		/*4*/ imports.NewTable("WKNavigationAction_NavigationType", 0),
		/*5*/ imports.NewTable("WKNavigationAction_Request", 0),
		/*6*/ imports.NewTable("WKNavigationAction_SourceFrame", 0),
		/*7*/ imports.NewTable("WKNavigationAction_TargetFrame", 0),
	}
)

func wKNavigationActionImportAPI() *imports.Imports {
	if wKNavigationActionImport == nil {
		wKNavigationActionImport = NewDefaultImports()
		wKNavigationActionImport.SetImportTable(wKNavigationActionImportTables)
		wKNavigationActionImportTables = nil
	}
	return wKNavigationActionImport
}
