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

// IWkURIRequest Parent: IObject
type IWkURIRequest interface {
	IObject
	Data() wvTypes.WebKitURIRequest       // function
	Method() string                       // function
	Headers() wvTypes.PSoupMessageHeaders // function
	URI() string                          // property URI Getter
	SetURI(value string)                  // property URI Setter
}

type TWkURIRequest struct {
	TObject
}

func (m *TWkURIRequest) Data() wvTypes.WebKitURIRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkURIRequestAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitURIRequest(r)
}

func (m *TWkURIRequest) Method() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURIRequestAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURIRequest) Headers() wvTypes.PSoupMessageHeaders {
	if !m.IsValid() {
		return 0
	}
	r := wkURIRequestAPI().SysCallN(3, m.Instance())
	return wvTypes.PSoupMessageHeaders(r)
}

func (m *TWkURIRequest) URI() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wkURIRequestAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWkURIRequest) SetURI(value string) {
	if !m.IsValid() {
		return
	}
	wkURIRequestAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

// URIRequest  is static instance
var URIRequest _URIRequestClass

// _URIRequestClass is class type defined by TWkURIRequest
type _URIRequestClass uintptr

func (_URIRequestClass) NewURIRequest(uri string) IWkURIRequest {
	r := wkURIRequestAPI().SysCallN(4, api.PasStr(uri))
	return AsWkURIRequest(r)
}

// NewURIRequest class constructor
func NewURIRequest(uRIRequest wvTypes.WebKitURIRequest) IWkURIRequest {
	r := wkURIRequestAPI().SysCallN(0, uintptr(uRIRequest))
	return AsWkURIRequest(r)
}

var (
	wkURIRequestOnce   base.Once
	wkURIRequestImport *imports.Imports = nil
)

func wkURIRequestAPI() *imports.Imports {
	wkURIRequestOnce.Do(func() {
		wkURIRequestImport = api.NewDefaultImports()
		wkURIRequestImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkURIRequest_Create", 0), // constructor NewURIRequest
			/* 1 */ imports.NewTable("TWkURIRequest_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkURIRequest_Method", 0), // function Method
			/* 3 */ imports.NewTable("TWkURIRequest_Headers", 0), // function Headers
			/* 4 */ imports.NewTable("TWkURIRequest_NewURIRequest", 0), // static function NewURIRequest
			/* 5 */ imports.NewTable("TWkURIRequest_URI", 0), // property URI
		}
	})
	return wkURIRequestImport
}
