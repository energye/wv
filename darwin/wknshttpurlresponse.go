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

// IWkNSHTTPURLResponse Parent: IWkNSURLResponse
type IWkNSHTTPURLResponse interface {
	IWkNSURLResponse
	// InitWithURLStatusCodeHTTPVersionHeaderFields
	//  Initializes an HTTP URL response object with a status code, protocol version, and response headers.
	InitWithURLStatusCodeHTTPVersionHeaderFields(url NSURL, statusCode int32, hTTPVersion string, headerFieldsJSONString string) IWkNSHTTPURLResponse // function
	// StatusCode
	//  The response’s HTTP status code.
	StatusCode() int32 // function
	// AllHeaderFields
	//  All HTTP header fields of the response.
	AllHeaderFields() string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSHTTPURLResponse struct {
	TWkNSURLResponse
}

func (m *TWkNSHTTPURLResponse) InitWithURLStatusCodeHTTPVersionHeaderFields(url NSURL, statusCode int32, hTTPVersion string, headerFieldsJSONString string) IWkNSHTTPURLResponse {
	if !m.IsValid() {
		return nil
	}
	r := wkNSHTTPURLResponseAPI().SysCallN(2, m.Instance(), uintptr(url), uintptr(statusCode), api.PasStr(hTTPVersion), api.PasStr(headerFieldsJSONString))
	return AsWkNSHTTPURLResponse(r)
}

func (m *TWkNSHTTPURLResponse) StatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkNSHTTPURLResponseAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TWkNSHTTPURLResponse) AllHeaderFields() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkNSHTTPURLResponseAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkNSHTTPURLResponse) Release() {
	if !m.IsValid() {
		return
	}
	wkNSHTTPURLResponseAPI().SysCallN(6, m.Instance())
}

// HTTPURLResponse  is static instance
var HTTPURLResponse _HTTPURLResponseClass

// _HTTPURLResponseClass is class type defined by TWkNSHTTPURLResponse
type _HTTPURLResponseClass uintptr

// NewToWkNSHTTPURLResponse
//
//	Creates and returns an TWkNSHTTPURLResponse object.
func (_HTTPURLResponseClass) NewToWkNSHTTPURLResponse() IWkNSHTTPURLResponse {
	r := wkNSHTTPURLResponseAPI().SysCallN(1)
	return AsWkNSHTTPURLResponse(r)
}

// LocalizedStringForStatusCode
//
//	Returns a localized string corresponding to a specified HTTP status code.
//	statusCode: https://www.ietf.org/rfc/rfc2616.txt
func (_HTTPURLResponseClass) LocalizedStringForStatusCode(statusCode int32) string {
	r := wkNSHTTPURLResponseAPI().SysCallN(5, uintptr(statusCode))
	return api.GoStr(r)
}

// NewHTTPURLResponse class constructor
func NewHTTPURLResponse(data NSHTTPURLResponse) IWkNSHTTPURLResponse {
	r := wkNSHTTPURLResponseAPI().SysCallN(0, uintptr(data))
	return AsWkNSHTTPURLResponse(r)
}

var (
	wkNSHTTPURLResponseOnce   base.Once
	wkNSHTTPURLResponseImport *imports.Imports = nil
)

func wkNSHTTPURLResponseAPI() *imports.Imports {
	wkNSHTTPURLResponseOnce.Do(func() {
		wkNSHTTPURLResponseImport = api.NewDefaultImports()
		wkNSHTTPURLResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSHTTPURLResponse_Create", 0), // constructor NewHTTPURLResponse
			/* 1 */ imports.NewTable("TWkNSHTTPURLResponse_NewToWkNSHTTPURLResponse", 0), // static function NewToWkNSHTTPURLResponse
			/* 2 */ imports.NewTable("TWkNSHTTPURLResponse_InitWithURLStatusCodeHTTPVersionHeaderFields", 0), // function InitWithURLStatusCodeHTTPVersionHeaderFields
			/* 3 */ imports.NewTable("TWkNSHTTPURLResponse_StatusCode", 0), // function StatusCode
			/* 4 */ imports.NewTable("TWkNSHTTPURLResponse_AllHeaderFields", 0), // function AllHeaderFields
			/* 5 */ imports.NewTable("TWkNSHTTPURLResponse_LocalizedStringForStatusCode", 0), // static function LocalizedStringForStatusCode
			/* 6 */ imports.NewTable("TWkNSHTTPURLResponse_Release", 0), // procedure Release
		}
	})
	return wkNSHTTPURLResponseImport
}
