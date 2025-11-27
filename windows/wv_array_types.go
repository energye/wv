//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package windows

import "github.com/energye/lcl/base"

// IWVCustomSchemeRegistrationArray = array of ICoreWebView2CustomSchemeRegistration
type IWVCustomSchemeRegistrationArray interface {
	base.IArraySlice[ICoreWebView2CustomSchemeRegistration]
}

// NewWVCustomSchemeRegistrationArray 初始化 ICefBinaryValue 数组结构
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewWVCustomSchemeRegistrationArray(count int, instance uintptr) IWVCustomSchemeRegistrationArray {
	return base.NewInterfaceArraySlice[ICoreWebView2CustomSchemeRegistration](count, instance, AsCoreWebView2CustomSchemeRegistration)
}
