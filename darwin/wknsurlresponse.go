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
)

// IWkNSURLResponse Parent: lcl.IObject
type IWkNSURLResponse interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURLResponse // function
	// InitWithURLMIMETypeExpectedContentLengthTextEncodingName
	//  Creates an initialized NSURLResponse object with the URL, MIME type, length, and text encoding set to given values.
	InitWithURLMIMETypeExpectedContentLengthTextEncodingName(uRL NSURL, mIMEType string, length int32, name string) IWkNSURLResponse // function
	// URL
	//  The URL for the response.
	URL() NSURL // function
	// MIMEType
	//  The MIME type of the response.
	MIMEType() string // function
	// ExpectedContentLength
	//  The expected length of the response’s content.
	ExpectedContentLength() int64 // function
	// TextEncodingName
	//  The name of the text encoding provided by the response’s originating source.
	TextEncodingName() string // function
	// SuggestedFilename
	//  A suggested filename for the response data.
	SuggestedFilename() string // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNSURLResponse struct {
	lcl.TObject
}

func (m *TWkNSURLResponse) Data() NSURLResponse {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLResponseAPI().SysCallN(1, m.Instance())
	return NSURLResponse(r)
}

func (m *TWkNSURLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(uRL NSURL, mIMEType string, length int32, name string) IWkNSURLResponse {
	if !m.IsValid() {
		return nil
	}
	r := wkNSURLResponseAPI().SysCallN(3, m.Instance(), uintptr(uRL), api.PasStr(mIMEType), uintptr(length), api.PasStr(name))
	return AsWkNSURLResponse(r)
}

func (m *TWkNSURLResponse) URL() NSURL {
	if !m.IsValid() {
		return 0
	}
	r := wkNSURLResponseAPI().SysCallN(4, m.Instance())
	return NSURL(r)
}

func (m *TWkNSURLResponse) MIMEType() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLResponseAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSURLResponse) ExpectedContentLength() (result int64) {
	if !m.IsValid() {
		return
	}
	wkNSURLResponseAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkNSURLResponse) TextEncodingName() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLResponseAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSURLResponse) SuggestedFilename() string {
	if !m.IsValid() {
		return ""
	}
	r := wkNSURLResponseAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TWkNSURLResponse) Release() {
	if !m.IsValid() {
		return
	}
	wkNSURLResponseAPI().SysCallN(9, m.Instance())
}

// URLResponse  is static instance
var URLResponse _URLResponseClass

// _URLResponseClass is class type defined by TWkNSURLResponse
type _URLResponseClass uintptr

// New
//
//	Creates and returns an TWkNSURLResponse object.
func (_URLResponseClass) New() IWkNSURLResponse {
	r := wkNSURLResponseAPI().SysCallN(2)
	return AsWkNSURLResponse(r)
}

// NewURLResponse class constructor
func NewURLResponse(data NSURLResponse) IWkNSURLResponse {
	r := wkNSURLResponseAPI().SysCallN(0, uintptr(data))
	return AsWkNSURLResponse(r)
}

var (
	wkNSURLResponseOnce   base.Once
	wkNSURLResponseImport *imports.Imports = nil
)

func wkNSURLResponseAPI() *imports.Imports {
	wkNSURLResponseOnce.Do(func() {
		wkNSURLResponseImport = api.NewDefaultImports()
		wkNSURLResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNSURLResponse_Create", 0), // constructor NewURLResponse
			/* 1 */ imports.NewTable("TWkNSURLResponse_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNSURLResponse_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkNSURLResponse_InitWithURLMIMETypeExpectedContentLengthTextEncodingName", 0), // function InitWithURLMIMETypeExpectedContentLengthTextEncodingName
			/* 4 */ imports.NewTable("TWkNSURLResponse_URL", 0), // function URL
			/* 5 */ imports.NewTable("TWkNSURLResponse_MIMEType", 0), // function MIMEType
			/* 6 */ imports.NewTable("TWkNSURLResponse_ExpectedContentLength", 0), // function ExpectedContentLength
			/* 7 */ imports.NewTable("TWkNSURLResponse_TextEncodingName", 0), // function TextEncodingName
			/* 8 */ imports.NewTable("TWkNSURLResponse_SuggestedFilename", 0), // function SuggestedFilename
			/* 9 */ imports.NewTable("TWkNSURLResponse_Release", 0), // procedure Release
		}
	})
	return wkNSURLResponseImport
}
