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

// IWkFrameInfo Parent: IObject
type IWkFrameInfo interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKFrameInfo // function
	// IsMainFrame
	//  A Boolean value indicating whether the frame is the web site's main frame or a subframe.
	IsMainFrame() bool // function
	// Request
	//  The frame’s current request.
	Request() NSURLRequest // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkFrameInfo struct {
	TObject
}

func (m *TWkFrameInfo) Data() wvTypes.WKFrameInfo {
	if !m.IsValid() {
		return 0
	}
	r := wkFrameInfoAPI().SysCallN(1, m.Instance())
	return wvTypes.WKFrameInfo(r)
}

func (m *TWkFrameInfo) IsMainFrame() bool {
	if !m.IsValid() {
		return false
	}
	r := wkFrameInfoAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TWkFrameInfo) Request() NSURLRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkFrameInfoAPI().SysCallN(3, m.Instance())
	return NSURLRequest(r)
}

func (m *TWkFrameInfo) Release() {
	if !m.IsValid() {
		return
	}
	wkFrameInfoAPI().SysCallN(4, m.Instance())
}

// NewFrameInfo class constructor
func NewFrameInfo(data wvTypes.WKFrameInfo) IWkFrameInfo {
	r := wkFrameInfoAPI().SysCallN(0, uintptr(data))
	return AsWkFrameInfo(r)
}

var (
	wkFrameInfoOnce   base.Once
	wkFrameInfoImport *imports.Imports = nil
)

func wkFrameInfoAPI() *imports.Imports {
	wkFrameInfoOnce.Do(func() {
		wkFrameInfoImport = api.NewDefaultImports()
		wkFrameInfoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkFrameInfo_Create", 0), // constructor NewFrameInfo
			/* 1 */ imports.NewTable("TWkFrameInfo_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkFrameInfo_IsMainFrame", 0), // function IsMainFrame
			/* 3 */ imports.NewTable("TWkFrameInfo_Request", 0), // function Request
			/* 4 */ imports.NewTable("TWkFrameInfo_Release", 0), // procedure Release
		}
	})
	return wkFrameInfoImport
}
