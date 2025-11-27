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

// IWkNSMutableURLRequest Parent: IWkNSURLRequest
type IWkNSMutableURLRequest interface {
	IWkNSURLRequest
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetHTTPMethod
	//  The HTTP request method.
	SetHTTPMethod(method string) // procedure
	// SetAllHTTPHeaderFields
	//  A dictionary containing all of the HTTP header fields for a request.
	SetAllHTTPHeaderFields(headerFieldsJSONString string) // procedure
	// SetValueForHTTPHeaderField
	//  Sets a value for the header field.
	SetValueForHTTPHeaderField(value string, field string) // procedure
	// AddValueForHTTPHeaderField
	//  Adds a value to the header field.
	AddValueForHTTPHeaderField(value string, field string) // procedure
}

type TWkNSMutableURLRequest struct {
	TWkNSURLRequest
}

func (m *TWkNSMutableURLRequest) Release() {
	if !m.IsValid() {
		return
	}
	wkNSMutableURLRequestAPI().SysCallN(2, m.Instance())
}

func (m *TWkNSMutableURLRequest) SetHTTPMethod(method string) {
	if !m.IsValid() {
		return
	}
	wkNSMutableURLRequestAPI().SysCallN(3, m.Instance(), api.PasStr(method))
}

func (m *TWkNSMutableURLRequest) SetAllHTTPHeaderFields(headerFieldsJSONString string) {
	if !m.IsValid() {
		return
	}
	wkNSMutableURLRequestAPI().SysCallN(4, m.Instance(), api.PasStr(headerFieldsJSONString))
}

func (m *TWkNSMutableURLRequest) SetValueForHTTPHeaderField(value string, field string) {
	if !m.IsValid() {
		return
	}
	wkNSMutableURLRequestAPI().SysCallN(5, m.Instance(), api.PasStr(value), api.PasStr(field))
}

func (m *TWkNSMutableURLRequest) AddValueForHTTPHeaderField(value string, field string) {
	if !m.IsValid() {
		return
	}
	wkNSMutableURLRequestAPI().SysCallN(6, m.Instance(), api.PasStr(value), api.PasStr(field))
}

// MutableURLRequest  is static instance
var MutableURLRequest _MutableURLRequestClass

// _MutableURLRequestClass is class type defined by TWkNSMutableURLRequest
type _MutableURLRequestClass uintptr

// NewToWkNSMutableURLRequest
//
//	Creates and returns an NSMutableURLRequest object.
func (_MutableURLRequestClass) NewToWkNSMutableURLRequest() IWkNSMutableURLRequest {
	r := wkNSMutableURLRequestAPI().SysCallN(1)
	return AsWkNSMutableURLRequest(r)
}

// NewMutableURLRequest class constructor
func NewMutableURLRequest(data NSMutableURLRequest) IWkNSMutableURLRequest {
	r := wkNSMutableURLRequestAPI().SysCallN(0, uintptr(data))
	return AsWkNSMutableURLRequest(r)
}

var (
	wkNSMutableURLRequestOnce   base.Once
	wkNSMutableURLRequestImport *imports.Imports = nil
)

func wkNSMutableURLRequestAPI() *imports.Imports {
	wkNSMutableURLRequestOnce.Do(func() {
		wkNSMutableURLRequestImport = api.NewDefaultImports()
		wkNSMutableURLRequestImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSMutableURLRequest_Create", 0), // constructor NewMutableURLRequest
			/* 1 */ imports.NewTable("TWkNSMutableURLRequest_NewToWkNSMutableURLRequest", 0), // static function NewToWkNSMutableURLRequest
			/* 2 */ imports.NewTable("TWkNSMutableURLRequest_Release", 0), // procedure Release
			/* 3 */ imports.NewTable("TWkNSMutableURLRequest_SetHTTPMethod", 0), // procedure SetHTTPMethod
			/* 4 */ imports.NewTable("TWkNSMutableURLRequest_SetAllHTTPHeaderFields", 0), // procedure SetAllHTTPHeaderFields
			/* 5 */ imports.NewTable("TWkNSMutableURLRequest_SetValueForHTTPHeaderField", 0), // procedure SetValueForHTTPHeaderField
			/* 6 */ imports.NewTable("TWkNSMutableURLRequest_AddValueForHTTPHeaderField", 0), // procedure AddValueForHTTPHeaderField
		}
	})
	return wkNSMutableURLRequestImport
}
