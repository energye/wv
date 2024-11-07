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
type IWKUserScript interface {
	IObject
	Data() WKUserScript                 // function
	Source() string                     // function
	InjectionTime() (resultInt64 int64) // function
	IsForMainFrameOnly() bool           // function
}

// TWKUserScript Root Object
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

func (m *wKUserScript) New() IWKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(5)
	return AsWKUserScript(r1)
}

func (m *wKUserScript) InitWithSourceInjectionTimeForMainFrameOnly(aSource string, aInjectionTime int64, aForMainFrameOnly bool) IWKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(2, PascalStr(aSource), uintptr(unsafePointer(&aInjectionTime)), PascalBool(aForMainFrameOnly))
	return AsWKUserScript(r1)
}

func (m *TWKUserScript) Data() WKUserScript {
	r1 := wKUserScriptImportAPI().SysCallN(1, m.Instance())
	return WKUserScript(r1)
}

func (m *TWKUserScript) Source() string {
	r1 := wKUserScriptImportAPI().SysCallN(6, m.Instance())
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

var (
	wKUserScriptImport       *imports.Imports = nil
	wKUserScriptImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WKUserScript_Create", 0),
		/*1*/ imports.NewTable("WKUserScript_Data", 0),
		/*2*/ imports.NewTable("WKUserScript_InitWithSourceInjectionTimeForMainFrameOnly", 0),
		/*3*/ imports.NewTable("WKUserScript_InjectionTime", 0),
		/*4*/ imports.NewTable("WKUserScript_IsForMainFrameOnly", 0),
		/*5*/ imports.NewTable("WKUserScript_New", 0),
		/*6*/ imports.NewTable("WKUserScript_Source", 0),
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
