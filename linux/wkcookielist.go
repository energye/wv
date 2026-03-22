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

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkCookieList Parent: IObject
type IWkCookieList interface {
	IObject
	Data() wvTypes.PList                       // function
	Length() int32                             // function
	GetCookie(index int32) wvTypes.PSoupCookie // function
}

type TWkCookieList struct {
	TObject
}

func (m *TWkCookieList) Data() wvTypes.PList {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieListAPI().SysCallN(1, m.Instance())
	return wvTypes.PList(r)
}

func (m *TWkCookieList) Length() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieListAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TWkCookieList) GetCookie(index int32) wvTypes.PSoupCookie {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieListAPI().SysCallN(3, m.Instance(), uintptr(index))
	return wvTypes.PSoupCookie(r)
}

// NewCookieList class constructor
func NewCookieList(list wvTypes.PList) IWkCookieList {
	r := wkCookieListAPI().SysCallN(0, uintptr(list))
	return AsWkCookieList(r)
}

var (
	wkCookieListOnce   base.Once
	wkCookieListImport *imports.Imports = nil
)

func wkCookieListAPI() *imports.Imports {
	wkCookieListOnce.Do(func() {
		wkCookieListImport = api.NewDefaultImports()
		wkCookieListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkCookieList_Create", 0), // constructor NewCookieList
			/* 1 */ imports.NewTable("TWkCookieList_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkCookieList_Length", 0), // function Length
			/* 3 */ imports.NewTable("TWkCookieList_GetCookie", 0), // function GetCookie
		}
	})
	return wkCookieListImport
}
