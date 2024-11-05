//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// ICoreWebView2WebResourceResponseViewGetContentCompletedHandler Parent: IObject
//
//	Receives the result of the ICoreWebView2WebResourceResponseView.GetContent method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler">See the ICoreWebView2WebResourceResponseViewGetContentCompletedHandler article.</a>
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler interface {
	IObject
	ResourceID() int32 // property
}

// TCoreWebView2WebResourceResponseViewGetContentCompletedHandler Parent: TObject
//
//	Receives the result of the ICoreWebView2WebResourceResponseView.GetContent method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseviewgetcontentcompletedhandler">See the ICoreWebView2WebResourceResponseViewGetContentCompletedHandler article.</a>
type TCoreWebView2WebResourceResponseViewGetContentCompletedHandler struct {
	TObject
}

func NewCoreWebView2WebResourceResponseViewGetContentCompletedHandler(aEvents IWVBrowserEvents, aResourceID int32) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	r1 := coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportAPI().SysCallN(0, GetObjectUintptr(aEvents), uintptr(aResourceID))
	return AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler(r1)
}

func (m *TCoreWebView2WebResourceResponseViewGetContentCompletedHandler) ResourceID() int32 {
	r1 := coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

var (
	coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport       *imports.Imports = nil
	coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoreWebView2WebResourceResponseViewGetContentCompletedHandler_Create", 0),
		/*1*/ imports.NewTable("CoreWebView2WebResourceResponseViewGetContentCompletedHandler_ResourceID", 0),
	}
)

func coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportAPI() *imports.Imports {
	if coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport == nil {
		coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport = NewDefaultImports()
		coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport.SetImportTable(coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportTables)
		coreWebView2WebResourceResponseViewGetContentCompletedHandlerImportTables = nil
	}
	return coreWebView2WebResourceResponseViewGetContentCompletedHandlerImport
}
