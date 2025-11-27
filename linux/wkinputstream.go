//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkInputStream Parent: lcl.IObject
type IWkInputStream interface {
	lcl.IObject
	Data() wvTypes.PInputStream                                          // function
	Read(buf uintptr, bufSize int32, readCount *int32, errorMsg *string) // procedure
	Close()                                                              // procedure
}

type TWkInputStream struct {
	lcl.TObject
}

func (m *TWkInputStream) Data() wvTypes.PInputStream {
	if !m.IsValid() {
		return 0
	}
	r := wkInputStreamAPI().SysCallN(1, m.Instance())
	return wvTypes.PInputStream(r)
}

func (m *TWkInputStream) Read(buf uintptr, bufSize int32, readCount *int32, errorMsg *string) {
	if !m.IsValid() {
		return
	}
	readCountPtr := uintptr(*readCount)
	errorMsgPtr := api.PasStr(*errorMsg)
	wkInputStreamAPI().SysCallN(3, m.Instance(), uintptr(buf), uintptr(bufSize), uintptr(base.UnsafePointer(&readCountPtr)), uintptr(base.UnsafePointer(&errorMsgPtr)))
	*readCount = int32(readCountPtr)
	*errorMsg = api.GoStr(errorMsgPtr)
}

func (m *TWkInputStream) Close() {
	if !m.IsValid() {
		return
	}
	wkInputStreamAPI().SysCallN(4, m.Instance())
}

// InputStream  is static instance
var InputStream _InputStreamClass

// _InputStreamClass is class type defined by TWkInputStream
type _InputStreamClass uintptr

func (_InputStreamClass) New(buf uintptr, streamLength int64) IWkInputStream {
	r := wkInputStreamAPI().SysCallN(2, uintptr(buf), uintptr(base.UnsafePointer(&streamLength)))
	return AsWkInputStream(r)
}

// NewInputStream class constructor
func NewInputStream(inputStream wvTypes.PInputStream) IWkInputStream {
	r := wkInputStreamAPI().SysCallN(0, uintptr(inputStream))
	return AsWkInputStream(r)
}

var (
	wkInputStreamOnce   base.Once
	wkInputStreamImport *imports.Imports = nil
)

func wkInputStreamAPI() *imports.Imports {
	wkInputStreamOnce.Do(func() {
		wkInputStreamImport = api.NewDefaultImports()
		wkInputStreamImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkInputStream_Create", 0), // constructor NewInputStream
			/* 1 */ imports.NewTable("TWkInputStream_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkInputStream_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkInputStream_Read", 0), // procedure Read
			/* 4 */ imports.NewTable("TWkInputStream_Close", 0), // procedure Close
		}
	})
	return wkInputStreamImport
}
