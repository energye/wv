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

// IWkNavigationDelegate Parent: lcl.IObject
type IWkNavigationDelegate interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKNavigationDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNavigationDelegate struct {
	lcl.TObject
}

func (m *TWkNavigationDelegate) Data() wvTypes.WKNavigationDelegateProtocol {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationDelegateAPI().SysCallN(1, m.Instance())
	return wvTypes.WKNavigationDelegateProtocol(r)
}

func (m *TWkNavigationDelegate) Release() {
	if !m.IsValid() {
		return
	}
	wkNavigationDelegateAPI().SysCallN(2, m.Instance())
}

// NewNavigationDelegate class constructor
func NewNavigationDelegate(event IWkWebview) IWkNavigationDelegate {
	r := wkNavigationDelegateAPI().SysCallN(0, base.GetObjectUintptr(event))
	return AsWkNavigationDelegate(r)
}

var (
	wkNavigationDelegateOnce   base.Once
	wkNavigationDelegateImport *imports.Imports = nil
)

func wkNavigationDelegateAPI() *imports.Imports {
	wkNavigationDelegateOnce.Do(func() {
		wkNavigationDelegateImport = api.NewDefaultImports()
		wkNavigationDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigationDelegate_Create", 0), // constructor NewNavigationDelegate
			/* 1 */ imports.NewTable("TWkNavigationDelegate_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNavigationDelegate_Release", 0), // procedure Release
		}
	})
	return wkNavigationDelegateImport
}
