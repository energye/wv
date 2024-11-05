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

// IWkWebsitePolicies Root Interface
type IWkWebsitePolicies interface {
	IObject
	Data() WebKitWebsitePolicies             // function
	GetAutoplayPolicy() WebKitAutoplayPolicy // function
}

// TWkWebsitePolicies Root Object
type TWkWebsitePolicies struct {
	TObject
}

func NewWkWebsitePolicies() IWkWebsitePolicies {
	r1 := wkWebsitePoliciesImportAPI().SysCallN(0)
	return AsWkWebsitePolicies(r1)
}

func NewWkWebsitePolicies1(firstpolicyname string) IWkWebsitePolicies {
	r1 := wkWebsitePoliciesImportAPI().SysCallN(1, PascalStr(firstpolicyname))
	return AsWkWebsitePolicies(r1)
}

func (m *TWkWebsitePolicies) Data() WebKitWebsitePolicies {
	r1 := wkWebsitePoliciesImportAPI().SysCallN(2, m.Instance())
	return WebKitWebsitePolicies(r1)
}

func (m *TWkWebsitePolicies) GetAutoplayPolicy() WebKitAutoplayPolicy {
	r1 := wkWebsitePoliciesImportAPI().SysCallN(3, m.Instance())
	return WebKitAutoplayPolicy(r1)
}

var (
	wkWebsitePoliciesImport       *imports.Imports = nil
	wkWebsitePoliciesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkWebsitePolicies_Create", 0),
		/*1*/ imports.NewTable("WkWebsitePolicies_Create1", 0),
		/*2*/ imports.NewTable("WkWebsitePolicies_Data", 0),
		/*3*/ imports.NewTable("WkWebsitePolicies_GetAutoplayPolicy", 0),
	}
)

func wkWebsitePoliciesImportAPI() *imports.Imports {
	if wkWebsitePoliciesImport == nil {
		wkWebsitePoliciesImport = NewDefaultImports()
		wkWebsitePoliciesImport.SetImportTable(wkWebsitePoliciesImportTables)
		wkWebsitePoliciesImportTables = nil
	}
	return wkWebsitePoliciesImport
}
