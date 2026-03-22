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

// ICoreWebView2FrameInfoCollection Parent: IObject
type ICoreWebView2FrameInfoCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property Initialized Getter
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameInfoCollection // property BaseIntf Getter
	// Iterator
	//  Gets an iterator over the collection of `FrameInfo`s.
	//  <see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollection#getiterator">See the ICoreWebView2FrameInfoCollection article.</see>
	Iterator() ICoreWebView2FrameInfoCollectionIterator // property Iterator Getter
}

type TCoreWebView2FrameInfoCollection struct {
	TObject
}

func (m *TCoreWebView2FrameInfoCollection) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r := coreWebView2FrameInfoCollectionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCoreWebView2FrameInfoCollection) BaseIntf() (result ICoreWebView2FrameInfoCollection) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoCollectionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfoCollection(resultPtr)
	return
}

func (m *TCoreWebView2FrameInfoCollection) Iterator() (result ICoreWebView2FrameInfoCollectionIterator) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	coreWebView2FrameInfoCollectionAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsCoreWebView2FrameInfoCollectionIterator(resultPtr)
	return
}

// NewCoreWebView2FrameInfoCollection class constructor
func NewCoreWebView2FrameInfoCollection(baseIntf ICoreWebView2FrameInfoCollection) ICoreWebView2FrameInfoCollection {
	r := coreWebView2FrameInfoCollectionAPI().SysCallN(0, base.GetObjectUintptr(baseIntf))
	return AsCoreWebView2FrameInfoCollection(r)
}

var (
	coreWebView2FrameInfoCollectionOnce   base.Once
	coreWebView2FrameInfoCollectionImport *imports.Imports = nil
)

func coreWebView2FrameInfoCollectionAPI() *imports.Imports {
	coreWebView2FrameInfoCollectionOnce.Do(func() {
		coreWebView2FrameInfoCollectionImport = api.NewDefaultImports()
		coreWebView2FrameInfoCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2FrameInfoCollection_Create", 0), // constructor NewCoreWebView2FrameInfoCollection
			/* 1 */ imports.NewTable("TCoreWebView2FrameInfoCollection_Initialized", 0), // property Initialized
			/* 2 */ imports.NewTable("TCoreWebView2FrameInfoCollection_BaseIntf", 0), // property BaseIntf
			/* 3 */ imports.NewTable("TCoreWebView2FrameInfoCollection_Iterator", 0), // property Iterator
		}
	})
	return coreWebView2FrameInfoCollectionImport
}
