//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// GetGlobalWebView2Loader 获取全局 WebView2Loader 实例
// 一搬用于在初始时尝试获取
func GetGlobalWebView2Loader() IWVLoader {
	r := wvLoaderDefAPI().SysCallN(0)
	return AsWVLoader(r)
}

// SetGlobalWebView2Loader 设置全局 WebView2Loader 实例
// 用于在创建 IWVLoader 后设置
func SetGlobalWebView2Loader(loader IWVLoader) {
	wvLoaderDefAPI().SysCallN(1, loader.Instance())
}

// DestroyGlobalWebView2Loader 销毁 全局 WebView2Loader 实例
// 用于在应用退出后使用
func DestroyGlobalWebView2Loader() {
	wvLoaderDefAPI().SysCallN(2)
}

var (
	wvLoaderDefOnce   base.Once
	wvLoaderDefImport *imports.Imports = nil
)

func wvLoaderDefAPI() *imports.Imports {
	wvLoaderDefOnce.Do(func() {
		wvLoaderDefImport = api.NewDefaultImports()
		wvLoaderDefImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("GetGlobalWebView2Loader", 0),
			/* 1 */ imports.NewTable("SetGlobalWebView2Loader", 0),
			/* 2 */ imports.NewTable("DestroyGlobalWebView2Loader", 0),
		}
	})
	return wvLoaderDefImport
}
