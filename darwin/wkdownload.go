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

// IWkDownload Parent: IObject
type IWkDownload interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() wvTypes.WKDownload // function
	// OriginalRequest
	//  An object that represents the request that initiated the download.
	OriginalRequest() NSURLRequest // function
	// Progress
	//  An object that conveys ongoing progress to the user for a specified task.
	//  https://developer.apple.com/documentation/foundation/nsprogress?language=objc
	Progress() NSProgress // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
	// SetDelegate
	//  An object you use to track download progress and handle redirects, authentication challenges, and failures.
	SetDelegate(downloadDelegate IWkDownloadDelegate) // procedure
	// Cancel
	//  Cancels the download, and optionally captures data so that you can resume the download later.
	Cancel() // procedure
}

type TWkDownload struct {
	TObject
}

func (m *TWkDownload) Data() wvTypes.WKDownload {
	if !m.IsValid() {
		return 0
	}
	r := wkDownloadAPI().SysCallN(1, m.Instance())
	return wvTypes.WKDownload(r)
}

func (m *TWkDownload) OriginalRequest() NSURLRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkDownloadAPI().SysCallN(2, m.Instance())
	return NSURLRequest(r)
}

func (m *TWkDownload) Progress() NSProgress {
	if !m.IsValid() {
		return 0
	}
	r := wkDownloadAPI().SysCallN(3, m.Instance())
	return NSProgress(r)
}

func (m *TWkDownload) Release() {
	if !m.IsValid() {
		return
	}
	wkDownloadAPI().SysCallN(4, m.Instance())
}

func (m *TWkDownload) SetDelegate(downloadDelegate IWkDownloadDelegate) {
	if !m.IsValid() {
		return
	}
	wkDownloadAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(downloadDelegate))
}

func (m *TWkDownload) Cancel() {
	if !m.IsValid() {
		return
	}
	wkDownloadAPI().SysCallN(6, m.Instance())
}

// NewDownload class constructor
func NewDownload(data wvTypes.WKDownload, webViewDelegate IWkWebview) IWkDownload {
	r := wkDownloadAPI().SysCallN(0, uintptr(data), base.GetObjectUintptr(webViewDelegate))
	return AsWkDownload(r)
}

var (
	wkDownloadOnce   base.Once
	wkDownloadImport *imports.Imports = nil
)

func wkDownloadAPI() *imports.Imports {
	wkDownloadOnce.Do(func() {
		wkDownloadImport = api.NewDefaultImports()
		wkDownloadImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkDownload_Create", 0), // constructor NewDownload
			/* 1 */ imports.NewTable("TWkDownload_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkDownload_OriginalRequest", 0), // function OriginalRequest
			/* 3 */ imports.NewTable("TWkDownload_Progress", 0), // function Progress
			/* 4 */ imports.NewTable("TWkDownload_Release", 0), // procedure Release
			/* 5 */ imports.NewTable("TWkDownload_SetDelegate", 0), // procedure SetDelegate
			/* 6 */ imports.NewTable("TWkDownload_Cancel", 0), // procedure Cancel
		}
	})
	return wkDownloadImport
}
