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

// IWkCookie Root Interface
type IWkCookie interface {
	IObject
	Name() string                // property
	SetName(AValue string)       // property
	Value() string               // property
	SetValue(AValue string)      // property
	Domain() string              // property
	SetDomain(AValue string)     // property
	Path() string                // property
	SetPath(AValue string)       // property
	HttpOnly() bool              // property
	SetHttpOnly(AValue bool)     // property
	Secure() bool                // property
	SetSecure(AValue bool)       // property
	Expires() PSoupDate          // property
	SetExpires(AValue PSoupDate) // property
	Data() PSoupCookie           // function
	SetMaxAge(aValue int32)      // procedure
}

// TWkCookie Root Object
type TWkCookie struct {
	TObject
}

func NewWkCookie(aCookie PSoupCookie) IWkCookie {
	r1 := wkCookieImportAPI().SysCallN(0, uintptr(aCookie))
	return AsWkCookie(r1)
}

// WkCookieRef -> IWkCookie
var WkCookieRef wkCookie

// wkCookie TWkCookie Ref
type wkCookie uintptr

func (m *wkCookie) NewCookie(aName, aValue, aDomain, aPath string, aMaxAge int32) IWkCookie {
	r1 := wkCookieImportAPI().SysCallN(6, PascalStr(aName), PascalStr(aValue), PascalStr(aDomain), PascalStr(aPath), uintptr(aMaxAge))
	return AsWkCookie(r1)
}

func (m *TWkCookie) Name() string {
	r1 := wkCookieImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkCookie) SetName(AValue string) {
	wkCookieImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkCookie) Value() string {
	r1 := wkCookieImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkCookie) SetValue(AValue string) {
	wkCookieImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkCookie) Domain() string {
	r1 := wkCookieImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkCookie) SetDomain(AValue string) {
	wkCookieImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkCookie) Path() string {
	r1 := wkCookieImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWkCookie) SetPath(AValue string) {
	wkCookieImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWkCookie) HttpOnly() bool {
	r1 := wkCookieImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWkCookie) SetHttpOnly(AValue bool) {
	wkCookieImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWkCookie) Secure() bool {
	r1 := wkCookieImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWkCookie) SetSecure(AValue bool) {
	wkCookieImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWkCookie) Expires() PSoupDate {
	r1 := wkCookieImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return PSoupDate(r1)
}

func (m *TWkCookie) SetExpires(AValue PSoupDate) {
	wkCookieImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TWkCookie) Data() PSoupCookie {
	r1 := wkCookieImportAPI().SysCallN(1, m.Instance())
	return PSoupCookie(r1)
}

func (m *TWkCookie) SetMaxAge(aValue int32) {
	wkCookieImportAPI().SysCallN(9, m.Instance(), uintptr(aValue))
}

var (
	wkCookieImport       *imports.Imports = nil
	wkCookieImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkCookie_Create", 0),
		/*1*/ imports.NewTable("WkCookie_Data", 0),
		/*2*/ imports.NewTable("WkCookie_Domain", 0),
		/*3*/ imports.NewTable("WkCookie_Expires", 0),
		/*4*/ imports.NewTable("WkCookie_HttpOnly", 0),
		/*5*/ imports.NewTable("WkCookie_Name", 0),
		/*6*/ imports.NewTable("WkCookie_NewCookie", 0),
		/*7*/ imports.NewTable("WkCookie_Path", 0),
		/*8*/ imports.NewTable("WkCookie_Secure", 0),
		/*9*/ imports.NewTable("WkCookie_SetMaxAge", 0),
		/*10*/ imports.NewTable("WkCookie_Value", 0),
	}
)

func wkCookieImportAPI() *imports.Imports {
	if wkCookieImport == nil {
		wkCookieImport = NewDefaultImports()
		wkCookieImport.SetImportTable(wkCookieImportTables)
		wkCookieImportTables = nil
	}
	return wkCookieImport
}
