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

// IWkInputStream Root Interface
type IWkInputStream interface {
	IObject
	Data() PInputStream                                                  // function
	Read(buf uintptr, bufSize int32, readCount *int32, errorMsg *string) // procedure
	Close()                                                              // procedure
}

// TWkInputStream Root Object
type TWkInputStream struct {
	TObject
}

func NewWkInputStream(aInputStream PInputStream) IWkInputStream {
	r1 := wkInputStreamImportAPI().SysCallN(1, uintptr(aInputStream))
	return AsWkInputStream(r1)
}

// WkInputStreamRef -> IWkInputStream
var WkInputStreamRef wkInputStream

// wkInputStream TWkInputStream Ref
type wkInputStream uintptr

func (m *wkInputStream) New(aBuf uintptr, aStreamLength int64) IWkInputStream {
	r1 := wkInputStreamImportAPI().SysCallN(3, uintptr(aBuf), uintptr(unsafePointer(&aStreamLength)))
	return AsWkInputStream(r1)
}

func (m *TWkInputStream) Data() PInputStream {
	r1 := wkInputStreamImportAPI().SysCallN(2, m.Instance())
	return PInputStream(r1)
}

func (m *TWkInputStream) Read(buf uintptr, bufSize int32, readCount *int32, errorMsg *string) {
	var result2 uintptr
	var result3 uintptr
	wkInputStreamImportAPI().SysCallN(4, m.Instance(), uintptr(buf), uintptr(bufSize), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*readCount = int32(result2)
	*errorMsg = GoStr(result3)
}

func (m *TWkInputStream) Close() {
	wkInputStreamImportAPI().SysCallN(0, m.Instance())
}

var (
	wkInputStreamImport       *imports.Imports = nil
	wkInputStreamImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkInputStream_Close", 0),
		/*1*/ imports.NewTable("WkInputStream_Create", 0),
		/*2*/ imports.NewTable("WkInputStream_Data", 0),
		/*3*/ imports.NewTable("WkInputStream_New", 0),
		/*4*/ imports.NewTable("WkInputStream_Read", 0),
	}
)

func wkInputStreamImportAPI() *imports.Imports {
	if wkInputStreamImport == nil {
		wkInputStreamImport = NewDefaultImports()
		wkInputStreamImport.SetImportTable(wkInputStreamImportTables)
		wkInputStreamImportTables = nil
	}
	return wkInputStreamImport
}
