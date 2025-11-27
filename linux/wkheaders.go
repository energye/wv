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

// IWkHeaders Parent: lcl.IObject
type IWkHeaders interface {
	lcl.IObject
	Data() wvTypes.PSoupMessageHeaders // function
	List() lcl.IStrings                // function
	Append(name string, value string)  // procedure
	Clear()                            // procedure
	Remove(name string)                // procedure
	Replace(name string, value string) // procedure
}

type TWkHeaders struct {
	lcl.TObject
}

func (m *TWkHeaders) Data() wvTypes.PSoupMessageHeaders {
	if !m.IsValid() {
		return 0
	}
	r := wkHeadersAPI().SysCallN(1, m.Instance())
	return wvTypes.PSoupMessageHeaders(r)
}

func (m *TWkHeaders) List() lcl.IStrings {
	if !m.IsValid() {
		return nil
	}
	r := wkHeadersAPI().SysCallN(2, m.Instance())
	return lcl.AsStrings(r)
}

func (m *TWkHeaders) Append(name string, value string) {
	if !m.IsValid() {
		return
	}
	wkHeadersAPI().SysCallN(4, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TWkHeaders) Clear() {
	if !m.IsValid() {
		return
	}
	wkHeadersAPI().SysCallN(5, m.Instance())
}

func (m *TWkHeaders) Remove(name string) {
	if !m.IsValid() {
		return
	}
	wkHeadersAPI().SysCallN(6, m.Instance(), api.PasStr(name))
}

func (m *TWkHeaders) Replace(name string, value string) {
	if !m.IsValid() {
		return
	}
	wkHeadersAPI().SysCallN(7, m.Instance(), api.PasStr(name), api.PasStr(value))
}

// Headers  is static instance
var Headers _HeadersClass

// _HeadersClass is class type defined by TWkHeaders
type _HeadersClass uintptr

func (_HeadersClass) NewMessageHeadersType(type_ wvTypes.TSoupMessageHeadersType) IWkHeaders {
	r := wkHeadersAPI().SysCallN(3, uintptr(type_))
	return AsWkHeaders(r)
}

// NewHeaders class constructor
func NewHeaders(messageHeaders wvTypes.PSoupMessageHeaders) IWkHeaders {
	r := wkHeadersAPI().SysCallN(0, uintptr(messageHeaders))
	return AsWkHeaders(r)
}

var (
	wkHeadersOnce   base.Once
	wkHeadersImport *imports.Imports = nil
)

func wkHeadersAPI() *imports.Imports {
	wkHeadersOnce.Do(func() {
		wkHeadersImport = api.NewDefaultImports()
		wkHeadersImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkHeaders_Create", 0), // constructor NewHeaders
			/* 1 */ imports.NewTable("TWkHeaders_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkHeaders_List", 0), // function List
			/* 3 */ imports.NewTable("TWkHeaders_NewMessageHeadersType", 0), // static function NewMessageHeadersType
			/* 4 */ imports.NewTable("TWkHeaders_Append", 0), // procedure Append
			/* 5 */ imports.NewTable("TWkHeaders_Clear", 0), // procedure Clear
			/* 6 */ imports.NewTable("TWkHeaders_Remove", 0), // procedure Remove
			/* 7 */ imports.NewTable("TWkHeaders_Replace", 0), // procedure Replace
		}
	})
	return wkHeadersImport
}
