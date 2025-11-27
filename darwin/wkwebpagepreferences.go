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

// IWkWebpagePreferences Parent: lcl.IObject
type IWkWebpagePreferences interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKWebpagePreferences // function
	// PreferredContentMode
	//  Returns The content mode for the web view to use when it loads and renders a webpage.
	PreferredContentMode() wvTypes.WKContentMode // function
	// AllowsContentJavaScript
	//  Returns A Boolean value that indicates whether JavaScript from web content is allowed to run.
	AllowsContentJavaScript() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetPreferredContentMode
	//  Sets The content mode for the web view to use when it loads and renders a webpage.
	SetPreferredContentMode(value wvTypes.WKContentMode) // procedure
	// SetAllowsContentJavaScript
	//  Sets A Boolean value that indicates whether JavaScript from web content is allowed to run.
	SetAllowsContentJavaScript(value bool) // procedure
}

type TWkWebpagePreferences struct {
	lcl.TObject
}

func (m *TWkWebpagePreferences) Data() wvTypes.WKWebpagePreferences {
	if !m.IsValid() {
		return 0
	}
	r := wkWebpagePreferencesAPI().SysCallN(1, m.Instance())
	return wvTypes.WKWebpagePreferences(r)
}

func (m *TWkWebpagePreferences) PreferredContentMode() wvTypes.WKContentMode {
	if !m.IsValid() {
		return 0
	}
	r := wkWebpagePreferencesAPI().SysCallN(3, m.Instance())
	return wvTypes.WKContentMode(r)
}

func (m *TWkWebpagePreferences) AllowsContentJavaScript() bool {
	if !m.IsValid() {
		return false
	}
	r := wkWebpagePreferencesAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TWkWebpagePreferences) Release() {
	if !m.IsValid() {
		return
	}
	wkWebpagePreferencesAPI().SysCallN(5, m.Instance())
}

func (m *TWkWebpagePreferences) SetPreferredContentMode(value wvTypes.WKContentMode) {
	if !m.IsValid() {
		return
	}
	wkWebpagePreferencesAPI().SysCallN(6, m.Instance(), uintptr(value))
}

func (m *TWkWebpagePreferences) SetAllowsContentJavaScript(value bool) {
	if !m.IsValid() {
		return
	}
	wkWebpagePreferencesAPI().SysCallN(7, m.Instance(), api.PasBool(value))
}

// WebpagePreferences  is static instance
var WebpagePreferences _WebpagePreferencesClass

// _WebpagePreferencesClass is class type defined by TWkWebpagePreferences
type _WebpagePreferencesClass uintptr

// New
//
//	Creates and returns an WKUserScript object.
func (_WebpagePreferencesClass) New() IWkWebpagePreferences {
	r := wkWebpagePreferencesAPI().SysCallN(2)
	return AsWkWebpagePreferences(r)
}

// NewWebpagePreferences class constructor
func NewWebpagePreferences(data wvTypes.WKWebpagePreferences) IWkWebpagePreferences {
	r := wkWebpagePreferencesAPI().SysCallN(0, uintptr(data))
	return AsWkWebpagePreferences(r)
}

var (
	wkWebpagePreferencesOnce   base.Once
	wkWebpagePreferencesImport *imports.Imports = nil
)

func wkWebpagePreferencesAPI() *imports.Imports {
	wkWebpagePreferencesOnce.Do(func() {
		wkWebpagePreferencesImport = api.NewDefaultImports()
		wkWebpagePreferencesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebpagePreferences_Create", 0), // constructor NewWebpagePreferences
			/* 1 */ imports.NewTable("TWkWebpagePreferences_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkWebpagePreferences_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkWebpagePreferences_PreferredContentMode", 0), // function PreferredContentMode
			/* 4 */ imports.NewTable("TWkWebpagePreferences_AllowsContentJavaScript", 0), // function AllowsContentJavaScript
			/* 5 */ imports.NewTable("TWkWebpagePreferences_Release", 0), // procedure Release
			/* 6 */ imports.NewTable("TWkWebpagePreferences_SetPreferredContentMode", 0), // procedure SetPreferredContentMode
			/* 7 */ imports.NewTable("TWkWebpagePreferences_SetAllowsContentJavaScript", 0), // procedure SetAllowsContentJavaScript
		}
	})
	return wkWebpagePreferencesImport
}
