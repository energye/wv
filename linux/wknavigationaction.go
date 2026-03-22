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

// IWkNavigationAction Parent: IObject
type IWkNavigationAction interface {
	IObject
	Data() wvTypes.WebKitNavigationAction // function
	GetRequest() wvTypes.WebKitURIRequest // function
}

type TWkNavigationAction struct {
	TObject
}

func (m *TWkNavigationAction) Data() wvTypes.WebKitNavigationAction {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitNavigationAction(r)
}

func (m *TWkNavigationAction) GetRequest() wvTypes.WebKitURIRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationActionAPI().SysCallN(2, m.Instance())
	return wvTypes.WebKitURIRequest(r)
}

// NewNavigationAction class constructor
func NewNavigationAction(navigationAction wvTypes.WebKitNavigationAction) IWkNavigationAction {
	r := wkNavigationActionAPI().SysCallN(0, uintptr(navigationAction))
	return AsWkNavigationAction(r)
}

var (
	wkNavigationActionOnce   base.Once
	wkNavigationActionImport *imports.Imports = nil
)

func wkNavigationActionAPI() *imports.Imports {
	wkNavigationActionOnce.Do(func() {
		wkNavigationActionImport = api.NewDefaultImports()
		wkNavigationActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigationAction_Create", 0), // constructor NewNavigationAction
			/* 1 */ imports.NewTable("TWkNavigationAction_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNavigationAction_GetRequest", 0), // function GetRequest
		}
	})
	return wkNavigationActionImport
}
