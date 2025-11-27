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

// IWkUserScript Parent: lcl.IObject
type IWkUserScript interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKUserScript // function
	// Source
	//  The script’s source code.
	Source() string // function
	// InjectionTime
	//  The time at which to inject the script into the webpage.
	InjectionTime() int64 // function
	// IsForMainFrameOnly
	//  A Boolean value that indicates whether to inject the script into the main frame or all frames.
	IsForMainFrameOnly() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkUserScript struct {
	lcl.TObject
}

func (m *TWkUserScript) Data() wvTypes.WKUserScript {
	if !m.IsValid() {
		return 0
	}
	r := wkUserScriptAPI().SysCallN(1, m.Instance())
	return wvTypes.WKUserScript(r)
}

func (m *TWkUserScript) Source() string {
	if !m.IsValid() {
		return ""
	}
	r := wkUserScriptAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TWkUserScript) InjectionTime() (result int64) {
	if !m.IsValid() {
		return
	}
	wkUserScriptAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkUserScript) IsForMainFrameOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := wkUserScriptAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkUserScript) Release() {
	if !m.IsValid() {
		return
	}
	wkUserScriptAPI().SysCallN(7, m.Instance())
}

// UserScript  is static instance
var UserScript _UserScriptClass

// _UserScriptClass is class type defined by TWkUserScript
type _UserScriptClass uintptr

// New
//
//	Creates and returns an WKUserScript object.
func (_UserScriptClass) New() IWkUserScript {
	r := wkUserScriptAPI().SysCallN(2)
	return AsWkUserScript(r)
}

// InitWithSourceInjectionTimeForMainFrameOnly
//
//	Creates a user script object that contains the specified source code and attributes.
func (_UserScriptClass) InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime int64, forMainFrameOnly bool) IWkUserScript {
	r := wkUserScriptAPI().SysCallN(6, api.PasStr(source), uintptr(base.UnsafePointer(&injectionTime)), api.PasBool(forMainFrameOnly))
	return AsWkUserScript(r)
}

// NewUserScript class constructor
func NewUserScript(data wvTypes.WKUserScript) IWkUserScript {
	r := wkUserScriptAPI().SysCallN(0, uintptr(data))
	return AsWkUserScript(r)
}

var (
	wkUserScriptOnce   base.Once
	wkUserScriptImport *imports.Imports = nil
)

func wkUserScriptAPI() *imports.Imports {
	wkUserScriptOnce.Do(func() {
		wkUserScriptImport = api.NewDefaultImports()
		wkUserScriptImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkUserScript_Create", 0), // constructor NewUserScript
			/* 1 */ imports.NewTable("TWkUserScript_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkUserScript_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkUserScript_Source", 0), // function Source
			/* 4 */ imports.NewTable("TWkUserScript_InjectionTime", 0), // function InjectionTime
			/* 5 */ imports.NewTable("TWkUserScript_IsForMainFrameOnly", 0), // function IsForMainFrameOnly
			/* 6 */ imports.NewTable("TWkUserScript_InitWithSourceInjectionTimeForMainFrameOnly", 0), // static function InitWithSourceInjectionTimeForMainFrameOnly
			/* 7 */ imports.NewTable("TWkUserScript_Release", 0), // procedure Release
		}
	})
	return wkUserScriptImport
}
