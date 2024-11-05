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

// INSURLCredential Root Interface
type INSURLCredential interface {
	IObject
	Data() NSURLCredential                   // function
	Persistence() NSURLCredentialPersistence // function
}

// TNSURLCredential Root Object
type TNSURLCredential struct {
	TObject
}

func NewNSURLCredential(aData NSURLCredential) INSURLCredential {
	r1 := nSURLCredentialImportAPI().SysCallN(0, uintptr(aData))
	return AsNSURLCredential(r1)
}

func (m *TNSURLCredential) Data() NSURLCredential {
	r1 := nSURLCredentialImportAPI().SysCallN(1, m.Instance())
	return NSURLCredential(r1)
}

func (m *TNSURLCredential) Persistence() NSURLCredentialPersistence {
	r1 := nSURLCredentialImportAPI().SysCallN(2, m.Instance())
	return NSURLCredentialPersistence(r1)
}

var (
	nSURLCredentialImport       *imports.Imports = nil
	nSURLCredentialImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLCredential_Create", 0),
		/*1*/ imports.NewTable("NSURLCredential_Data", 0),
		/*2*/ imports.NewTable("NSURLCredential_Persistence", 0),
	}
)

func nSURLCredentialImportAPI() *imports.Imports {
	if nSURLCredentialImport == nil {
		nSURLCredentialImport = NewDefaultImports()
		nSURLCredentialImport.SetImportTable(nSURLCredentialImportTables)
		nSURLCredentialImportTables = nil
	}
	return nSURLCredentialImport
}
