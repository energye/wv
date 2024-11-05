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

// INSMutableURLRequest Root Interface
type INSMutableURLRequest interface {
	IObject
	SetHTTPMethod(aMethod string)                            // procedure
	SetAllHTTPHeaderFields(headerFieldsJSONString string)    // procedure
	SetValueForHTTPHeaderField(aValue string, aField string) // procedure
	AddValueForHTTPHeaderField(aValue string, aField string) // procedure
}

// TNSMutableURLRequest Root Object
type TNSMutableURLRequest struct {
	TObject
}

func NewNSMutableURLRequest() INSMutableURLRequest {
	r1 := nSMutableURLRequestImportAPI().SysCallN(1)
	return AsNSMutableURLRequest(r1)
}

func NewNSMutableURLRequest1(aData NSMutableURLRequest) INSMutableURLRequest {
	r1 := nSMutableURLRequestImportAPI().SysCallN(2, uintptr(aData))
	return AsNSMutableURLRequest(r1)
}

func (m *TNSMutableURLRequest) SetHTTPMethod(aMethod string) {
	nSMutableURLRequestImportAPI().SysCallN(4, m.Instance(), PascalStr(aMethod))
}

func (m *TNSMutableURLRequest) SetAllHTTPHeaderFields(headerFieldsJSONString string) {
	nSMutableURLRequestImportAPI().SysCallN(3, m.Instance(), PascalStr(headerFieldsJSONString))
}

func (m *TNSMutableURLRequest) SetValueForHTTPHeaderField(aValue string, aField string) {
	nSMutableURLRequestImportAPI().SysCallN(5, m.Instance(), PascalStr(aValue), PascalStr(aField))
}

func (m *TNSMutableURLRequest) AddValueForHTTPHeaderField(aValue string, aField string) {
	nSMutableURLRequestImportAPI().SysCallN(0, m.Instance(), PascalStr(aValue), PascalStr(aField))
}

var (
	nSMutableURLRequestImport       *imports.Imports = nil
	nSMutableURLRequestImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSMutableURLRequest_AddValueForHTTPHeaderField", 0),
		/*1*/ imports.NewTable("NSMutableURLRequest_Create", 0),
		/*2*/ imports.NewTable("NSMutableURLRequest_Create1", 0),
		/*3*/ imports.NewTable("NSMutableURLRequest_SetAllHTTPHeaderFields", 0),
		/*4*/ imports.NewTable("NSMutableURLRequest_SetHTTPMethod", 0),
		/*5*/ imports.NewTable("NSMutableURLRequest_SetValueForHTTPHeaderField", 0),
	}
)

func nSMutableURLRequestImportAPI() *imports.Imports {
	if nSMutableURLRequestImport == nil {
		nSMutableURLRequestImport = NewDefaultImports()
		nSMutableURLRequestImport.SetImportTable(nSMutableURLRequestImportTables)
		nSMutableURLRequestImportTables = nil
	}
	return nSMutableURLRequestImport
}
