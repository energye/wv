//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/lcl/lcl"

	"unsafe"
)

type unsafePointer = unsafe.Pointer

type IStrings = lcl.IStrings
type TStrings = lcl.TStrings

type ICustomPanel = lcl.ICustomPanel
type TCustomPanel = lcl.TCustomPanel

type IObject = lcl.IObject
type TObject = lcl.TObject

type IComponent = lcl.IComponent
type TComponent = lcl.TComponent

type IWinControl = lcl.IWinControl
type TWinControl = lcl.TWinControl

// IUnknown 根接口
//
//	UUID: '{00000000-0000-0000-C000-000000000046}'
type IUnknown = lcl.IUnknown
type Unknown = lcl.Unknown

// GetInstance As操作的简化。
//
// Simplification of As operation.
//
//go:noinline
func GetInstance(value interface{}) unsafePointer {
	return lcl.GetInstance(value)
}

// SetObjectInstance 设置对你指针实例值, 用于外部组件创建
func SetObjectInstance(object interface{}, instance unsafePointer) {
	if object == nil {
		return
	}
	switch object.(type) {
	case IObject:
		lcl.SetObjectInstance(object.(IObject), instance)
	}
}

// GetObjectUintptr 获取对象指针地址值
func GetObjectUintptr(object interface{}) uintptr {
	if object == nil {
		return 0
	}
	switch object.(type) {
	case IObject:
		return lcl.GetObjectUintptr(object.(IObject))
	}
	return 0
}

type TWkButtonEvent struct {
	EventType    TGdkEventType
	State        TGdkModifierType
	Button       int32
	X, Y         int32
	XRoot, YRoot int32
	Time         int32
}
