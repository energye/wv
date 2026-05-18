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

// IWVProxySettings Parent: IObject
type IWVProxySettings interface {
	IObject
	NoProxyServer() bool         // property NoProxyServer Getter
	SetNoProxyServer(value bool) // property NoProxyServer Setter
	AutoDetect() bool            // property AutoDetect Getter
	SetAutoDetect(value bool)    // property AutoDetect Setter
	ByPassList() string          // property ByPassList Getter
	SetByPassList(value string)  // property ByPassList Setter
	PacUrl() string              // property PacUrl Getter
	SetPacUrl(value string)      // property PacUrl Setter
	Server() string              // property Server Getter
	SetServer(value string)      // property Server Setter
}

type TWVProxySettings struct {
	TObject
}

func (m *TWVProxySettings) NoProxyServer() bool {
	if !m.IsValid() {
		return false
	}
	r := wVProxySettingsAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVProxySettings) SetNoProxyServer(value bool) {
	if !m.IsValid() {
		return
	}
	wVProxySettingsAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVProxySettings) AutoDetect() bool {
	if !m.IsValid() {
		return false
	}
	r := wVProxySettingsAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWVProxySettings) SetAutoDetect(value bool) {
	if !m.IsValid() {
		return
	}
	wVProxySettingsAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TWVProxySettings) ByPassList() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVProxySettingsAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVProxySettings) SetByPassList(value string) {
	if !m.IsValid() {
		return
	}
	wVProxySettingsAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVProxySettings) PacUrl() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVProxySettingsAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVProxySettings) SetPacUrl(value string) {
	if !m.IsValid() {
		return
	}
	wVProxySettingsAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TWVProxySettings) Server() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	wVProxySettingsAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TWVProxySettings) SetServer(value string) {
	if !m.IsValid() {
		return
	}
	wVProxySettingsAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

// NewProxySettings class constructor
func NewProxySettings() IWVProxySettings {
	r := wVProxySettingsAPI().SysCallN(0)
	return AsWVProxySettings(r)
}

var (
	wVProxySettingsOnce   base.Once
	wVProxySettingsImport *imports.Imports = nil
)

func wVProxySettingsAPI() *imports.Imports {
	wVProxySettingsOnce.Do(func() {
		wVProxySettingsImport = api.NewDefaultImports()
		wVProxySettingsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWVProxySettings_Create", 0), // constructor NewProxySettings
			/* 1 */ imports.NewTable("TWVProxySettings_NoProxyServer", 0), // property NoProxyServer
			/* 2 */ imports.NewTable("TWVProxySettings_AutoDetect", 0), // property AutoDetect
			/* 3 */ imports.NewTable("TWVProxySettings_ByPassList", 0), // property ByPassList
			/* 4 */ imports.NewTable("TWVProxySettings_PacUrl", 0), // property PacUrl
			/* 5 */ imports.NewTable("TWVProxySettings_Server", 0), // property Server
		}
	})
	return wVProxySettingsImport
}
