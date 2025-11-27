//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// ICoreWebView2ObjectCollection Parent: ICoreWebView2ObjectCollectionView
type ICoreWebView2ObjectCollection interface {
	ICoreWebView2ObjectCollectionView
	// RemoveValueAtIndex
	//  Removes the object at the specified index.
	RemoveValueAtIndex(index uint32) bool // function
}

type TCoreWebView2ObjectCollection struct {
	TCoreWebView2ObjectCollectionView
}

func (m *TCoreWebView2ObjectCollection) RemoveValueAtIndex(index uint32) bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2ObjectCollectionAPI().SysCallN(1, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

// NewCoreWebView2ObjectCollection class constructor
func NewCoreWebView2ObjectCollection(baseIntf ICoreWebView2ObjectCollectionView) ICoreWebView2ObjectCollection {
	r := coreWebView2ObjectCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2ObjectCollection(r)
}

var (
	coreWebView2ObjectCollectionOnce   base.Once
	coreWebView2ObjectCollectionImport *imports.Imports = nil
)

func coreWebView2ObjectCollectionAPI() *imports.Imports {
	coreWebView2ObjectCollectionOnce.Do(func() {
		coreWebView2ObjectCollectionImport = api.NewDefaultImports()
		coreWebView2ObjectCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2ObjectCollection_Create", 0), // constructor NewCoreWebView2ObjectCollection
			/* 1 */ imports.NewTable("TCoreWebView2ObjectCollection_RemoveValueAtIndex", 0), // function RemoveValueAtIndex
		}
	})
	return coreWebView2ObjectCollectionImport
}
