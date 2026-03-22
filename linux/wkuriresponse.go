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

// IWkURIResponse Parent: IObject
type IWkURIResponse interface {
	IObject
	Data() wvTypes.WebKitURIResponse      // function
	Headers() wvTypes.PSoupMessageHeaders // function
	GetStatusCode() int32                 // function
	GetContentLength() uint64             // function
	GetMimeType() string                  // function
	GetSuggestedFilename() string         // function
	URI() string                          // property URI Getter
}

type TWkURIResponse struct {
	TObject
}

func (m *TWkURIResponse) Data() wvTypes.WebKitURIResponse {
	if !m.IsValid() {
		return 0
	}
	r := wkURIResponseAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitURIResponse(r)
}

func (m *TWkURIResponse) Headers() wvTypes.PSoupMessageHeaders {
	if !m.IsValid() {
		return 0
	}
	r := wkURIResponseAPI().SysCallN(2, m.Instance())
	return wvTypes.PSoupMessageHeaders(r)
}

func (m *TWkURIResponse) GetStatusCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := wkURIResponseAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TWkURIResponse) GetContentLength() (result uint64) {
	if !m.IsValid() {
		return
	}
	wkURIResponseAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TWkURIResponse) GetMimeType() string {
	if !m.IsValid() {
		return ""
	}
	r := wkURIResponseAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TWkURIResponse) GetSuggestedFilename() string {
	if !m.IsValid() {
		return ""
	}
	r := wkURIResponseAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TWkURIResponse) URI() string {
	if !m.IsValid() {
		return ""
	}
	r := wkURIResponseAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

// NewURIResponse class constructor
func NewURIResponse(uRIResponse wvTypes.WebKitURIResponse) IWkURIResponse {
	r := wkURIResponseAPI().SysCallN(0, uintptr(uRIResponse))
	return AsWkURIResponse(r)
}

var (
	wkURIResponseOnce   base.Once
	wkURIResponseImport *imports.Imports = nil
)

func wkURIResponseAPI() *imports.Imports {
	wkURIResponseOnce.Do(func() {
		wkURIResponseImport = api.NewDefaultImports()
		wkURIResponseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkURIResponse_Create", 0), // constructor NewURIResponse
			/* 1 */ imports.NewTable("TWkURIResponse_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkURIResponse_Headers", 0), // function Headers
			/* 3 */ imports.NewTable("TWkURIResponse_GetStatusCode", 0), // function GetStatusCode
			/* 4 */ imports.NewTable("TWkURIResponse_GetContentLength", 0), // function GetContentLength
			/* 5 */ imports.NewTable("TWkURIResponse_GetMimeType", 0), // function GetMimeType
			/* 6 */ imports.NewTable("TWkURIResponse_GetSuggestedFilename", 0), // function GetSuggestedFilename
			/* 7 */ imports.NewTable("TWkURIResponse_URI", 0), // property URI
		}
	})
	return wkURIResponseImport
}
