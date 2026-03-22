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

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkNavigationResponse Parent: IObject
type IWkNavigationResponse interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKNavigationResponse // function
	// IsForMainFrame
	//  A Boolean value that indicates whether the response targets the web view’s main frame.
	IsForMainFrame() bool // function
	// Response
	//  The frame’s response.
	Response() NSURLResponse // function
	// CanShowMIMEType
	//  A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
	CanShowMIMEType() bool // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNavigationResponse struct {
	TObject
}

func (m *TWkNavigationResponse) Data() wvTypes.WKNavigationResponse {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationResponseAPI().SysCallN(1, m.Instance())
	return wvTypes.WKNavigationResponse(r)
}

func (m *TWkNavigationResponse) IsForMainFrame() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNavigationResponseAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNavigationResponse) Response() NSURLResponse {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationResponseAPI().SysCallN(4, m.Instance())
	return NSURLResponse(r)
}

func (m *TWkNavigationResponse) CanShowMIMEType() bool {
	if !m.IsValid() {
		return false
	}
	r := wkNavigationResponseAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TWkNavigationResponse) Release() {
	if !m.IsValid() {
		return
	}
	wkNavigationResponseAPI().SysCallN(6, m.Instance())
}

// NavigationResponse  is static instance
var NavigationResponse _NavigationResponseClass

// _NavigationResponseClass is class type defined by TWkNavigationResponse
type _NavigationResponseClass uintptr

// New
//
//	Creates and returns an WKNavigationResponse object.
func (_NavigationResponseClass) New() IWkNavigationResponse {
	r := wkNavigationResponseAPI().SysCallN(2)
	return AsWkNavigationResponse(r)
}

// NewNavigationResponse class constructor
func NewNavigationResponse(data wvTypes.WKNavigationResponse) IWkNavigationResponse {
	r := wkNavigationResponseAPI().SysCallN(0, uintptr(data))
	return AsWkNavigationResponse(r)
}

var (
	wkNavigationResponseOnce   base.Once
	wkNavigationResponseImport *imports.Imports = nil
)

func wkNavigationResponseAPI() *imports.Imports {
	wkNavigationResponseOnce.Do(func() {
		wkNavigationResponseImport = api.NewDefaultImports()
		wkNavigationResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigationResponse_Create", 0), // constructor NewNavigationResponse
			/* 1 */ imports.NewTable("TWkNavigationResponse_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNavigationResponse_New", 0), // static function New
			/* 3 */ imports.NewTable("TWkNavigationResponse_IsForMainFrame", 0), // function IsForMainFrame
			/* 4 */ imports.NewTable("TWkNavigationResponse_Response", 0), // function Response
			/* 5 */ imports.NewTable("TWkNavigationResponse_CanShowMIMEType", 0), // function CanShowMIMEType
			/* 6 */ imports.NewTable("TWkNavigationResponse_Release", 0), // procedure Release
		}
	})
	return wkNavigationResponseImport
}
