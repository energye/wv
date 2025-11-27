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

// IWkNavigationAction Parent: lcl.IObject
type IWkNavigationAction interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKNavigationAction // function
	// SourceFrame
	//  The frame that requested the navigation.
	SourceFrame() wvTypes.WKFrameInfo // function
	// TargetFrame
	//  The frame in which to display the new content.
	TargetFrame() wvTypes.WKFrameInfo // function
	// NavigationType
	//  The type of action that triggered the navigation.
	NavigationType() wvTypes.WKNavigationType // function
	// Request
	//  The URL request object associated with the navigation action.
	Request() NSURLRequest // function
	// ModifierFlags
	//  The modifier keys that were pressed at the time of the navigation request.
	ModifierFlags() wvTypes.NSEventModifierFlags // function
	// ButtonNumber
	//  The number of the mouse button that caused the navigation request.
	ButtonNumber() int32 // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNavigationAction struct {
	lcl.TObject
}

func (m *TWkNavigationAction) Data() wvTypes.WKNavigationAction {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(1, m.Instance())
	return wvTypes.WKNavigationAction(r)
}

func (m *TWkNavigationAction) SourceFrame() wvTypes.WKFrameInfo {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(2, m.Instance())
	return wvTypes.WKFrameInfo(r)
}

func (m *TWkNavigationAction) TargetFrame() wvTypes.WKFrameInfo {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(3, m.Instance())
	return wvTypes.WKFrameInfo(r)
}

func (m *TWkNavigationAction) NavigationType() wvTypes.WKNavigationType {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(4, m.Instance())
	return wvTypes.WKNavigationType(r)
}

func (m *TWkNavigationAction) Request() NSURLRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(5, m.Instance())
	return NSURLRequest(r)
}

func (m *TWkNavigationAction) ModifierFlags() wvTypes.NSEventModifierFlags {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(6, m.Instance())
	return wvTypes.NSEventModifierFlags(r)
}

func (m *TWkNavigationAction) ButtonNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TWkNavigationAction) Release() {
	if !m.IsValid() {
		return
	}
	wkNavigationActionAPI().SysCallN(8, m.Instance())
}

// NewNavigationAction class constructor
func NewNavigationAction(data wvTypes.WKNavigationAction) IWkNavigationAction {
	r := wkNavigationActionAPI().SysCallN(0, uintptr(data))
	return AsWkNavigationAction(r)
}

var (
	wkNavigationActionOnce   base.Once
	wkNavigationActionImport *imports.Imports = nil
)

func wkNavigationActionAPI() *imports.Imports {
	wkNavigationActionOnce.Do(func() {
		wkNavigationActionImport = api.NewDefaultImports()
		wkNavigationActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigationAction_Create", 0), // constructor NewNavigationAction
			/* 1 */ imports.NewTable("TWkNavigationAction_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNavigationAction_SourceFrame", 0), // function SourceFrame
			/* 3 */ imports.NewTable("TWkNavigationAction_TargetFrame", 0), // function TargetFrame
			/* 4 */ imports.NewTable("TWkNavigationAction_NavigationType", 0), // function NavigationType
			/* 5 */ imports.NewTable("TWkNavigationAction_Request", 0), // function Request
			/* 6 */ imports.NewTable("TWkNavigationAction_ModifierFlags", 0), // function ModifierFlags
			/* 7 */ imports.NewTable("TWkNavigationAction_ButtonNumber", 0), // function ButtonNumber
			/* 8 */ imports.NewTable("TWkNavigationAction_Release", 0), // procedure Release
		}
	})
	return wkNavigationActionImport
}
