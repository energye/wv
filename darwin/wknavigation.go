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

	wvTypes "github.com/energye/wv/types/darwin"
)

// IWkNavigation Parent: lcl.IObject
type IWkNavigation interface {
	lcl.IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKNavigation // function
	// EffectiveContentMode
	//  The content mode WebKit uses to load the webpage.
	EffectiveContentMode() wvTypes.WKContentMode // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

type TWkNavigation struct {
	lcl.TObject
}

func (m *TWkNavigation) Data() wvTypes.WKNavigation {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationAPI().SysCallN(1, m.Instance())
	return wvTypes.WKNavigation(r)
}

func (m *TWkNavigation) EffectiveContentMode() wvTypes.WKContentMode {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationAPI().SysCallN(2, m.Instance())
	return wvTypes.WKContentMode(r)
}

func (m *TWkNavigation) Release() {
	if !m.IsValid() {
		return
	}
	wkNavigationAPI().SysCallN(3, m.Instance())
}

// NewNavigation class constructor
func NewNavigation(data wvTypes.WKNavigation) IWkNavigation {
	r := wkNavigationAPI().SysCallN(0, uintptr(data))
	return AsWkNavigation(r)
}

var (
	wkNavigationOnce   base.Once
	wkNavigationImport *imports.Imports = nil
)

func wkNavigationAPI() *imports.Imports {
	wkNavigationOnce.Do(func() {
		wkNavigationImport = api.NewDefaultImports()
		wkNavigationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigation_Create", 0), // constructor NewNavigation
			/* 1 */ imports.NewTable("TWkNavigation_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkNavigation_EffectiveContentMode", 0), // function EffectiveContentMode
			/* 3 */ imports.NewTable("TWkNavigation_Release", 0), // procedure Release
		}
	})
	return wkNavigationImport
}
