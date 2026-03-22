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
)

// IWkNSURLCredential Parent: IObject
type IWkNSURLCredential interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURLCredential // function
	// Persistence
	//  The credential’s persistence setting.
	Persistence() NSURLCredentialPersistence // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSURLCredential struct {
	TObject
}

func (m *TWkNSURLCredential) Data() NSURLCredential {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLCredentialAPI().SysCallN(1, m.Instance())
	return NSURLCredential(r)
}

func (m *TWkNSURLCredential) Persistence() NSURLCredentialPersistence {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLCredentialAPI().SysCallN(2, m.Instance())
	return NSURLCredentialPersistence(r)
}

func (m *TWkNSURLCredential) Release() {
	if !m.IsValid() {
		return
	}
	wkNSURLCredentialAPI().SysCallN(3, m.Instance())
}

// NewURLCredential class constructor
func NewURLCredential(data NSURLCredential) IWkNSURLCredential {
	r := wkNSURLCredentialAPI().SysCallN(0, uintptr(data))
	return AsWkNSURLCredential(r)
}

var (
	wkNSURLCredentialOnce   base.Once
	wkNSURLCredentialImport *imports.Imports = nil
)

func wkNSURLCredentialAPI() *imports.Imports {
	wkNSURLCredentialOnce.Do(func() {
		wkNSURLCredentialImport = api.NewDefaultImports()
		wkNSURLCredentialImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSURLCredential_Create", 0), // constructor NewURLCredential
			/* 1 */ imports.NewTable("TWkNSURLCredential_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSURLCredential_Persistence", 0), // function Persistence
			/* 3 */ imports.NewTable("TWkNSURLCredential_Release", 0), // procedure Release
		}
	})
	return wkNSURLCredentialImport
}
