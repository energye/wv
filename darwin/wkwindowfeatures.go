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

// IWKWindowFeatures Root Interface
type IWKWindowFeatures interface {
	IObject
	Data() WKWindowFeatures    // function
	MenuBarVisibility() bool   // function
	StatusBarVisibility() bool // function
	ToolbarsVisibility() bool  // function
	AllowsResizing() bool      // function
	X() int32                  // function
	Y() int32                  // function
	Width() int32              // function
	Height() int32             // function
}

// TWKWindowFeatures Root Object
type TWKWindowFeatures struct {
	TObject
}

func NewWKWindowFeatures(aData WKWindowFeatures) IWKWindowFeatures {
	r1 := wKWindowFeaturesImportAPI().SysCallN(1, uintptr(aData))
	return AsWKWindowFeatures(r1)
}

func (m *TWKWindowFeatures) Data() WKWindowFeatures {
	r1 := wKWindowFeaturesImportAPI().SysCallN(2, m.Instance())
	return WKWindowFeatures(r1)
}

func (m *TWKWindowFeatures) MenuBarVisibility() bool {
	r1 := wKWindowFeaturesImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TWKWindowFeatures) StatusBarVisibility() bool {
	r1 := wKWindowFeaturesImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TWKWindowFeatures) ToolbarsVisibility() bool {
	r1 := wKWindowFeaturesImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TWKWindowFeatures) AllowsResizing() bool {
	r1 := wKWindowFeaturesImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TWKWindowFeatures) X() int32 {
	r1 := wKWindowFeaturesImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func (m *TWKWindowFeatures) Y() int32 {
	r1 := wKWindowFeaturesImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TWKWindowFeatures) Width() int32 {
	r1 := wKWindowFeaturesImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TWKWindowFeatures) Height() int32 {
	r1 := wKWindowFeaturesImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

var (
	wKWindowFeaturesImport       *imports.Imports = nil
	wKWindowFeaturesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKWindowFeatures_AllowsResizing", 0),
		/*1*/ imports.NewTable("WKWindowFeatures_Create", 0),
		/*2*/ imports.NewTable("WKWindowFeatures_Data", 0),
		/*3*/ imports.NewTable("WKWindowFeatures_Height", 0),
		/*4*/ imports.NewTable("WKWindowFeatures_MenuBarVisibility", 0),
		/*5*/ imports.NewTable("WKWindowFeatures_StatusBarVisibility", 0),
		/*6*/ imports.NewTable("WKWindowFeatures_ToolbarsVisibility", 0),
		/*7*/ imports.NewTable("WKWindowFeatures_Width", 0),
		/*8*/ imports.NewTable("WKWindowFeatures_X", 0),
		/*9*/ imports.NewTable("WKWindowFeatures_Y", 0),
	}
)

func wKWindowFeaturesImportAPI() *imports.Imports {
	if wKWindowFeaturesImport == nil {
		wKWindowFeaturesImport = NewDefaultImports()
		wKWindowFeaturesImport.SetImportTable(wKWindowFeaturesImportTables)
		wKWindowFeaturesImportTables = nil
	}
	return wKWindowFeaturesImport
}
