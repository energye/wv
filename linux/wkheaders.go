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

// IWkHeaders Root Interface
type IWkHeaders interface {
	IObject
	Data() PSoupMessageHeaders           // function
	List() IStrings                      // function
	Append(aName string, aValue string)  // procedure
	Clear()                              // procedure
	Remove(aName string)                 // procedure
	Replace(aName string, aValue string) // procedure
}

// TWkHeaders Root Object
type TWkHeaders struct {
	TObject
}

func NewWkHeaders(aMessageHeaders PSoupMessageHeaders) IWkHeaders {
	r1 := wkHeadersImportAPI().SysCallN(2, uintptr(aMessageHeaders))
	return AsWkHeaders(r1)
}

// WkHeadersRef -> IWkHeaders
var WkHeadersRef wkHeaders

// wkHeaders TWkHeaders Ref
type wkHeaders uintptr

func (m *wkHeaders) NewMessageHeadersType(type_ TSoupMessageHeadersType) IWkHeaders {
	r1 := wkHeadersImportAPI().SysCallN(5, uintptr(type_))
	return AsWkHeaders(r1)
}

func (m *TWkHeaders) Data() PSoupMessageHeaders {
	r1 := wkHeadersImportAPI().SysCallN(3, m.Instance())
	return PSoupMessageHeaders(r1)
}

func (m *TWkHeaders) List() IStrings {
	var resultStrings uintptr
	wkHeadersImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultStrings)))
	return AsStrings(resultStrings)
}

func (m *TWkHeaders) Append(aName string, aValue string) {
	wkHeadersImportAPI().SysCallN(0, m.Instance(), PascalStr(aName), PascalStr(aValue))
}

func (m *TWkHeaders) Clear() {
	wkHeadersImportAPI().SysCallN(1, m.Instance())
}

func (m *TWkHeaders) Remove(aName string) {
	wkHeadersImportAPI().SysCallN(6, m.Instance(), PascalStr(aName))
}

func (m *TWkHeaders) Replace(aName string, aValue string) {
	wkHeadersImportAPI().SysCallN(7, m.Instance(), PascalStr(aName), PascalStr(aValue))
}

var (
	wkHeadersImport       *imports.Imports = nil
	wkHeadersImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkHeaders_Append", 0),
		/*1*/ imports.NewTable("WkHeaders_Clear", 0),
		/*2*/ imports.NewTable("WkHeaders_Create", 0),
		/*3*/ imports.NewTable("WkHeaders_Data", 0),
		/*4*/ imports.NewTable("WkHeaders_List", 0),
		/*5*/ imports.NewTable("WkHeaders_NewMessageHeadersType", 0),
		/*6*/ imports.NewTable("WkHeaders_Remove", 0),
		/*7*/ imports.NewTable("WkHeaders_Replace", 0),
	}
)

func wkHeadersImportAPI() *imports.Imports {
	if wkHeadersImport == nil {
		wkHeadersImport = NewDefaultImports()
		wkHeadersImport.SetImportTable(wkHeadersImportTables)
		wkHeadersImportTables = nil
	}
	return wkHeadersImport
}
