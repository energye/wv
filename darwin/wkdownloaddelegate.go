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

// IWkDownloadDelegate Parent: lcl.IObject
type IWkDownloadDelegate interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKDownloadDelegateProtocol // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkDownloadDelegate struct {
	lcl.TObject
}

func (m *TWkDownloadDelegate) Data() wvTypes.WKDownloadDelegateProtocol {
	if !m.IsValid() {
		return 0
	}
	r := wkDownloadDelegateAPI().SysCallN(1, m.Instance())
	return wvTypes.WKDownloadDelegateProtocol(r)
}

func (m *TWkDownloadDelegate) Release() {
	if !m.IsValid() {
		return
	}
	wkDownloadDelegateAPI().SysCallN(2, m.Instance())
}

// NewDownloadDelegate class constructor
func NewDownloadDelegate(event IWkWebview) IWkDownloadDelegate {
	r := wkDownloadDelegateAPI().SysCallN(0, base.GetObjectUintptr(event))
	return AsWkDownloadDelegate(r)
}

var (
	wkDownloadDelegateOnce   base.Once
	wkDownloadDelegateImport *imports.Imports = nil
)

func wkDownloadDelegateAPI() *imports.Imports {
	wkDownloadDelegateOnce.Do(func() {
		wkDownloadDelegateImport = api.NewDefaultImports()
		wkDownloadDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkDownloadDelegate_Create", 0), // constructor NewDownloadDelegate
			/* 1 */ imports.NewTable("TWkDownloadDelegate_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkDownloadDelegate_Release", 0), // procedure Release
		}
	})
	return wkDownloadDelegateImport
}
