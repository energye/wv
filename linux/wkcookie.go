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

// IWkCookie Parent: IObject
type IWkCookie interface {
	IObject
	Data() wvTypes.PSoupCookie          // function
	SetMaxAge(value int32)              // procedure
	Name() string                       // property Name Getter
	SetName(value string)               // property Name Setter
	Value() string                      // property Value Getter
	SetValue(value string)              // property Value Setter
	Domain() string                     // property Domain Getter
	SetDomain(value string)             // property Domain Setter
	Path() string                       // property Path Getter
	SetPath(value string)               // property Path Setter
	HttpOnly() bool                     // property HttpOnly Getter
	SetHttpOnly(value bool)             // property HttpOnly Setter
	Secure() bool                       // property Secure Getter
	SetSecure(value bool)               // property Secure Setter
	Expires() wvTypes.PSoupDate         // property Expires Getter
	SetExpires(value wvTypes.PSoupDate) // property Expires Setter
}

type TWkCookie struct {
	TObject
}

func (m *TWkCookie) Data() wvTypes.PSoupCookie {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieAPI().SysCallN(1, m.Instance())
	return wvTypes.PSoupCookie(r)
}

func (m *TWkCookie) SetMaxAge(value int32) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(3, m.Instance(), uintptr(value))
}

func (m *TWkCookie) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := wkCookieAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkCookie) SetName(value string) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkCookie) Value() string {
	if !m.IsValid() {
		return ""
	}
	r := wkCookieAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkCookie) SetValue(value string) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkCookie) Domain() string {
	if !m.IsValid() {
		return ""
	}
	r := wkCookieAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkCookie) SetDomain(value string) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkCookie) Path() string {
	if !m.IsValid() {
		return ""
	}
	r := wkCookieAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TWkCookie) SetPath(value string) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TWkCookie) HttpOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := wkCookieAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWkCookie) SetHttpOnly(value bool) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TWkCookie) Secure() bool {
	if !m.IsValid() {
		return false
	}
	r := wkCookieAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWkCookie) SetSecure(value bool) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TWkCookie) Expires() wvTypes.PSoupDate {
	if !m.IsValid() {
		return 0
	}
	r := wkCookieAPI().SysCallN(10, 0, m.Instance())
	return wvTypes.PSoupDate(r)
}

func (m *TWkCookie) SetExpires(value wvTypes.PSoupDate) {
	if !m.IsValid() {
		return
	}
	wkCookieAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

// Cookie  is static instance
var Cookie _CookieClass

// _CookieClass is class type defined by TWkCookie
type _CookieClass uintptr

func (_CookieClass) NewCookie(name string, value string, domain string, path string, maxAge int32) IWkCookie {
	r := wkCookieAPI().SysCallN(2, api.PasStr(name), api.PasStr(value), api.PasStr(domain), api.PasStr(path), uintptr(maxAge))
	return AsWkCookie(r)
}

// NewCookie class constructor
func NewCookie(cookie wvTypes.PSoupCookie) IWkCookie {
	r := wkCookieAPI().SysCallN(0, uintptr(cookie))
	return AsWkCookie(r)
}

var (
	wkCookieOnce   base.Once
	wkCookieImport *imports.Imports = nil
)

func wkCookieAPI() *imports.Imports {
	wkCookieOnce.Do(func() {
		wkCookieImport = api.NewDefaultImports()
		wkCookieImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkCookie_Create", 0), // constructor NewCookie
			/* 1 */ imports.NewTable("TWkCookie_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkCookie_NewCookie", 0), // static function NewCookie
			/* 3 */ imports.NewTable("TWkCookie_SetMaxAge", 0), // procedure SetMaxAge
			/* 4 */ imports.NewTable("TWkCookie_Name", 0), // property Name
			/* 5 */ imports.NewTable("TWkCookie_Value", 0), // property Value
			/* 6 */ imports.NewTable("TWkCookie_Domain", 0), // property Domain
			/* 7 */ imports.NewTable("TWkCookie_Path", 0), // property Path
			/* 8 */ imports.NewTable("TWkCookie_HttpOnly", 0), // property HttpOnly
			/* 9 */ imports.NewTable("TWkCookie_Secure", 0), // property Secure
			/* 10 */ imports.NewTable("TWkCookie_Expires", 0), // property Expires
		}
	})
	return wkCookieImport
}
