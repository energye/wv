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

// IWkUIDelegate Parent: lcl.IObject
type IWkUIDelegate interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKUIDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkUIDelegate struct {
	lcl.TObject
}

func (m *TWkUIDelegate) Data() wvTypes.WKUIDelegateProtocol {
	if !m.IsValid() {
		return 0
	}
	r := wkUIDelegateAPI().SysCallN(1, m.Instance())
	return wvTypes.WKUIDelegateProtocol(r)
}

func (m *TWkUIDelegate) Release() {
	if !m.IsValid() {
		return
	}
	wkUIDelegateAPI().SysCallN(2, m.Instance())
}

// NewUIDelegate class constructor
func NewUIDelegate(event IWkWebview) IWkUIDelegate {
	r := wkUIDelegateAPI().SysCallN(0, base.GetObjectUintptr(event))
	return AsWkUIDelegate(r)
}

var (
	wkUIDelegateOnce   base.Once
	wkUIDelegateImport *imports.Imports = nil
)

func wkUIDelegateAPI() *imports.Imports {
	wkUIDelegateOnce.Do(func() {
		wkUIDelegateImport = api.NewDefaultImports()
		wkUIDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkUIDelegate_Create", 0), // constructor NewUIDelegate
			/* 1 */ imports.NewTable("TWkUIDelegate_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkUIDelegate_Release", 0), // procedure Release
		}
	})
	return wkUIDelegateImport
}
