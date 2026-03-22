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

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkWindowFeatures Parent: IObject
type IWkWindowFeatures interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKWindowFeatures // function
	// MenuBarVisibility
	//  A Boolean value that indicates whether the webpage requests a visible menu bar.
	MenuBarVisibility() bool // function
	// StatusBarVisibility
	//  A Boolean value that indicates whether the webpage requested a visible status bar.
	StatusBarVisibility() bool // function
	// ToolbarsVisibility
	//  A Boolean value that indicates whether the webpage requested a visible toolbar.
	ToolbarsVisibility() bool // function
	// AllowsResizing
	//  A Boolean value that indicates whether to make the containing window window resizable.
	AllowsResizing() bool // function
	// X
	//  The requested x-coordinate of the containing window.
	X() int32 // function
	// Y
	//  The requested y-coordinate of the containing window.
	Y() int32 // function
	// Width
	//  The requested width of the containing window.
	Width() int32 // function
	// Height
	//  The requested height of the containing window.
	Height() int32 // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkWindowFeatures struct {
	TObject
}

func (m *TWkWindowFeatures) Data() wvTypes.WKWindowFeatures {
	if !m.IsValid() {
		return 0
	}
	r := wkWindowFeaturesAPI().SysCallN(1, m.Instance())
	return wvTypes.WKWindowFeatures(r)
}

func (m *TWkWindowFeatures) MenuBarVisibility() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWindowFeaturesAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWindowFeatures) StatusBarVisibility() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWindowFeaturesAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWindowFeatures) ToolbarsVisibility() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWindowFeaturesAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWindowFeatures) AllowsResizing() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWindowFeaturesAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWindowFeatures) X() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkWindowFeaturesAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TWkWindowFeatures) Y() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkWindowFeaturesAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TWkWindowFeatures) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkWindowFeaturesAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TWkWindowFeatures) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkWindowFeaturesAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TWkWindowFeatures) Release() {
	if !m.IsValid() {
		return
	}
	wkWindowFeaturesAPI().SysCallN(10, m.Instance())
}

// NewWindowFeatures class constructor
func NewWindowFeatures(data wvTypes.WKWindowFeatures) IWkWindowFeatures {
	r := wkWindowFeaturesAPI().SysCallN(0, uintptr(data))
	return AsWkWindowFeatures(r)
}

var (
	wkWindowFeaturesOnce   base.Once
	wkWindowFeaturesImport *imports.Imports = nil
)

func wkWindowFeaturesAPI() *imports.Imports {
	wkWindowFeaturesOnce.Do(func() {
		wkWindowFeaturesImport = api.NewDefaultImports()
		wkWindowFeaturesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWindowFeatures_Create", 0), // constructor NewWindowFeatures
			/* 1 */ imports.NewTable("TWkWindowFeatures_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkWindowFeatures_MenuBarVisibility", 0), // function MenuBarVisibility
			/* 3 */ imports.NewTable("TWkWindowFeatures_StatusBarVisibility", 0), // function StatusBarVisibility
			/* 4 */ imports.NewTable("TWkWindowFeatures_ToolbarsVisibility", 0), // function ToolbarsVisibility
			/* 5 */ imports.NewTable("TWkWindowFeatures_AllowsResizing", 0), // function AllowsResizing
			/* 6 */ imports.NewTable("TWkWindowFeatures_X", 0), // function X
			/* 7 */ imports.NewTable("TWkWindowFeatures_Y", 0), // function Y
			/* 8 */ imports.NewTable("TWkWindowFeatures_Width", 0), // function Width
			/* 9 */ imports.NewTable("TWkWindowFeatures_Height", 0), // function Height
			/* 10 */ imports.NewTable("TWkWindowFeatures_Release", 0), // procedure Release
		}
	})
	return wkWindowFeaturesImport
}
