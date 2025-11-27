//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

// IWVCustomSchemeInfoArrayWrap Parent: lcl.IObject
type IWVCustomSchemeInfoArrayWrap interface {
	lcl.IObject
	WVCustomSchemeInfoArrayData() types.PWVCustomSchemeInfoArray // function
	Size() int32                                                 // function
	GetValue(index int32) TWVCustomSchemeInfo                    // function
	SetValue(index int32, value TWVCustomSchemeInfo)             // procedure
}

type TWVCustomSchemeInfoArrayWrap struct {
	lcl.TObject
}

func (m *TWVCustomSchemeInfoArrayWrap) WVCustomSchemeInfoArrayData() types.PWVCustomSchemeInfoArray {
	if !m.IsValid() {
		return 0
	}
	r := wVCustomSchemeInfoArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PWVCustomSchemeInfoArray(r)
}

func (m *TWVCustomSchemeInfoArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wVCustomSchemeInfoArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TWVCustomSchemeInfoArrayWrap) GetValue(index int32) (result TWVCustomSchemeInfo) {
	if !m.IsValid() {
		return
	}
	resultPtr := result.ToPas()
	wVCustomSchemeInfoArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)))
	result = resultPtr.ToGo()
	return
}

func (m *TWVCustomSchemeInfoArrayWrap) SetValue(index int32, value TWVCustomSchemeInfo) {
	if !m.IsValid() {
		return
	}
	valuePtr := value.ToPas()
	wVCustomSchemeInfoArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(valuePtr)))
}

// NewCustomSchemeInfoArrayWrapWithInt class constructor
func NewCustomSchemeInfoArrayWrapWithInt(size int32) IWVCustomSchemeInfoArrayWrap {
	r := wVCustomSchemeInfoArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsWVCustomSchemeInfoArrayWrap(r)
}

var (
	wVCustomSchemeInfoArrayWrapOnce   base.Once
	wVCustomSchemeInfoArrayWrapImport *imports.Imports = nil
)

func wVCustomSchemeInfoArrayWrapAPI() *imports.Imports {
	wVCustomSchemeInfoArrayWrapOnce.Do(func() {
		wVCustomSchemeInfoArrayWrapImport = api.NewDefaultImports()
		wVCustomSchemeInfoArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVCustomSchemeInfoArrayWrap_CreateWithInt", 0), // constructor NewCustomSchemeInfoArrayWrapWithInt
			/* 1 */ imports.NewTable("TWVCustomSchemeInfoArrayWrap_WVCustomSchemeInfoArrayData", 0), // function WVCustomSchemeInfoArrayData
			/* 2 */ imports.NewTable("TWVCustomSchemeInfoArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TWVCustomSchemeInfoArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TWVCustomSchemeInfoArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return wVCustomSchemeInfoArrayWrapImport
}
