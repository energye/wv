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

// IWKUserScript Root Interface
//
//	A script that the web view injects into a webpage.
//	https://developer.apple.com/documentation/webkit/wkuserscript?language=objc
type IWKUserScript interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() WKUserScript // function
	// Source
	//  The script’s source code.
	Source() string // function
	// InjectionTime
	//  The time at which to inject the script into the webpage.
	InjectionTime() (resultInt64 int64) // function
	// IsForMainFrameOnly
	//  A Boolean value that indicates whether to inject the script into the main frame or all frames.
	IsForMainFrameOnly() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TWKUserScript Root Object
//
//	A script that the web view injects into a webpage.
//	https://developer.apple.com/documentation/webkit/wkuserscript?language=objc
type TWKUserScript struct {
	TObject
}

func NewWKUserScript(aData WKUserScript) IWKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(0, uintptr(aData))
	return AsWKUserScript(r1)
}

// WKUserScriptRef -> IWKUserScript
var WKUserScriptRef wKUserScript

// wKUserScript TWKUserScript Ref
type wKUserScript uintptr

// New
//
//	Creates and returns an WKUserScript object.
func (m *wKUserScript) New() IWKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(5)
	return AsWKUserScript(r1)
}

// InitWithSourceInjectionTimeForMainFrameOnly
//
//	Creates a user script object that contains the specified source code and attributes.
func (m *wKUserScript) InitWithSourceInjectionTimeForMainFrameOnly(aSource string, aInjectionTime int64, aForMainFrameOnly bool) IWKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(2, PascalStr(aSource), uintptr(unsafePointer(&aInjectionTime)), PascalBool(aForMainFrameOnly))
	return AsWKUserScript(r1)
}

func (m *TWKUserScript) Data() WKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(1, m.Instance())
	return WKUserScript(r1)
}

func (m *TWKUserScript) Source() string {
	r1 := wKUserScriptImportAPI().SysCallN(7, m.Instance())
	return GoStr(r1)
}

func (m *TWKUserScript) InjectionTime() (resultInt64 int64) {
	wKUserScriptImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TWKUserScript) IsForMainFrameOnly() bool {
	r1 := wKUserScriptImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TWKUserScript) Release() {
	wKUserScriptImportAPI().SysCallN(6, m.Instance())
}

var (
	wKUserScriptImport       *imports.Imports = nil
	wKUserScriptImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKUserScript_Create", 0),
		/*1*/ imports.NewTable("WKUserScript_Data", 0),
		/*2*/ imports.NewTable("WKUserScript_InitWithSourceInjectionTimeForMainFrameOnly", 0),
		/*3*/ imports.NewTable("WKUserScript_InjectionTime", 0),
		/*4*/ imports.NewTable("WKUserScript_IsForMainFrameOnly", 0),
		/*5*/ imports.NewTable("WKUserScript_New", 0),
		/*6*/ imports.NewTable("WKUserScript_Release", 0),
		/*7*/ imports.NewTable("WKUserScript_Source", 0),
	}
)

func wKUserScriptImportAPI() *imports.Imports {
	if wKUserScriptImport == nil {
		wKUserScriptImport = NewDefaultImports()
		wKUserScriptImport.SetImportTable(wKUserScriptImportTables)
		wKUserScriptImportTables = nil
	}
	return wKUserScriptImport
}
