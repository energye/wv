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

// IWkNSURL Parent: IObject
type IWkNSURL interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURL // function
	// AbsoluteString
	//  The URL string for the receiver as an absolute URL. (read-only)
	AbsoluteString() string // function
	// RelativeString
	//  A string representation of the relative portion of the URL. (read-only)
	RelativeString() string // function
	// BaseURL
	//  The base URL. (read-only)
	BaseURL() IWkNSURL // function
	// AbsoluteURL
	//  An absolute URL that refers to the same resource as the receiver. (read-only)
	AbsoluteURL() IWkNSURL // function
	// Scheme
	//  The scheme. (read-only)
	Scheme() string // function
	// Host
	//  The host, conforming to RFC 1808. (read-only)
	Host() string // function
	// Port
	//  The port, conforming to RFC 1808.
	Port() int32 // function
	// User
	//  The user name, conforming to RFC 1808.
	User() string // function
	// Password
	//  The password conforming to RFC 1808. (read-only)
	Password() string // function
	// Path
	//  The path, conforming to RFC 1808. (read-only)
	Path() string // function
	// Fragment
	//  The fragment identifier, conforming to RFC 1808. (read-only)
	Fragment() string // function
	// ParameterString
	//  The parameter string conforming to RFC 1808. (read-only)
	//  Deprecated
	ParameterString() string // function
	// Query
	//  The query string, conforming to RFC 1808.
	Query() string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSURL struct {
	TObject
}

func (m *TWkNSURL) Data() NSURL {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLAPI().SysCallN(1, m.Instance())
	return NSURL(r)
}

func (m *TWkNSURL) AbsoluteString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) RelativeString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) BaseURL() IWkNSURL {
	if !m.IsValid() {
		return nil
	}
	r := wkNSURLAPI().SysCallN(7, m.Instance())
	return AsWkNSURL(r)
}

func (m *TWkNSURL) AbsoluteURL() IWkNSURL {
	if !m.IsValid() {
		return nil
	}
	r := wkNSURLAPI().SysCallN(8, m.Instance())
	return AsWkNSURL(r)
}

func (m *TWkNSURL) Scheme() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Host() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Port() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TWkNSURL) User() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Password() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Path() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Fragment() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) ParameterString() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Query() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSURLAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSURL) Release() {
	if !m.IsValid() {
		return
	}
	wkNSURLAPI().SysCallN(18, m.Instance())
}

// URL  is static instance
var URL _URLClass

// _URLClass is class type defined by TWkNSURL
type _URLClass uintptr

// New
//
//	Creates and returns an NSURL object.
func (_URLClass) New() IWkNSURL {
	r := wkNSURLAPI().SysCallN(2)
	return AsWkNSURL(r)
}

// URLWithString
//
//	Creates and returns an NSURL object initialized with a provided URL string.
func (_URLClass) URLWithString(uRLString string) IWkNSURL {
	r := wkNSURLAPI().SysCallN(3, api.PasStr(uRLString))
	return AsWkNSURL(r)
}

// URLWithStringRelativeToURL
//
//	Creates and returns an NSURL object initialized with a base URL and a relative string.
func (_URLClass) URLWithStringRelativeToURL(uRLString string, baseURL NSURL) IWkNSURL {
	r := wkNSURLAPI().SysCallN(4, api.PasStr(uRLString), uintptr(baseURL))
	return AsWkNSURL(r)
}

// NewURL class constructor
func NewURL(data NSURL) IWkNSURL {
	r := wkNSURLAPI().SysCallN(0, uintptr(data))
	return AsWkNSURL(r)
}

var (
	wkNSURLOnce   base.Once
	wkNSURLImport *imports.Imports = nil
)

func wkNSURLAPI() *imports.Imports {
	wkNSURLOnce.Do(func() {
		wkNSURLImport = api.NewDefaultImports()
		wkNSURLImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSURL_Create", 0), // constructor NewURL
			/* 1 */ imports.NewTable("TWkNSURL_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSURL_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkNSURL_URLWithString", 0), // static function URLWithString
			/* 4 */ imports.NewTable("TWkNSURL_URLWithStringRelativeToURL", 0), // static function URLWithStringRelativeToURL
			/* 5 */ imports.NewTable("TWkNSURL_AbsoluteString", 0), // function AbsoluteString
			/* 6 */ imports.NewTable("TWkNSURL_RelativeString", 0), // function RelativeString
			/* 7 */ imports.NewTable("TWkNSURL_BaseURL", 0), // function BaseURL
			/* 8 */ imports.NewTable("TWkNSURL_AbsoluteURL", 0), // function AbsoluteURL
			/* 9 */ imports.NewTable("TWkNSURL_Scheme", 0), // function Scheme
			/* 10 */ imports.NewTable("TWkNSURL_Host", 0), // function Host
			/* 11 */ imports.NewTable("TWkNSURL_Port", 0), // function Port
			/* 12 */ imports.NewTable("TWkNSURL_User", 0), // function User
			/* 13 */ imports.NewTable("TWkNSURL_Password", 0), // function Password
			/* 14 */ imports.NewTable("TWkNSURL_Path", 0), // function Path
			/* 15 */ imports.NewTable("TWkNSURL_Fragment", 0), // function Fragment
			/* 16 */ imports.NewTable("TWkNSURL_ParameterString", 0), // function ParameterString
			/* 17 */ imports.NewTable("TWkNSURL_Query", 0), // function Query
			/* 18 */ imports.NewTable("TWkNSURL_Release", 0), // procedure Release
		}
	})
	return wkNSURLImport
}
