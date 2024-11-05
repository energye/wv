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

// IWVProxySettings Parent: IObject
//
//	Class used by the TWVLoader.ProxySettigns property to configure
//	a custom proxy server using the following command line switches:
//	--no-proxy-server, --proxy-auto-detect, --proxy-bypass-list,
//	--proxy-pac-url and --proxy-server.
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
type IWVProxySettings interface {
	IObject
	NoProxyServer() bool          // property
	SetNoProxyServer(AValue bool) // property
	AutoDetect() bool             // property
	SetAutoDetect(AValue bool)    // property
	ByPassList() string           // property
	SetByPassList(AValue string)  // property
	PacUrl() string               // property
	SetPacUrl(AValue string)      // property
	Server() string               // property
	SetServer(AValue string)      // property
}

// TWVProxySettings Parent: TObject
//
//	Class used by the TWVLoader.ProxySettigns property to configure
//	a custom proxy server using the following command line switches:
//	--no-proxy-server, --proxy-auto-detect, --proxy-bypass-list,
//	--proxy-pac-url and --proxy-server.
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
type TWVProxySettings struct {
	TObject
}

func NewWVProxySettings() IWVProxySettings {
	r1 := wVProxySettingsImportAPI().SysCallN(2)
	return AsWVProxySettings(r1)
}

func (m *TWVProxySettings) NoProxyServer() bool {
	r1 := wVProxySettingsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVProxySettings) SetNoProxyServer(AValue bool) {
	wVProxySettingsImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVProxySettings) AutoDetect() bool {
	r1 := wVProxySettingsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVProxySettings) SetAutoDetect(AValue bool) {
	wVProxySettingsImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVProxySettings) ByPassList() string {
	r1 := wVProxySettingsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetByPassList(AValue string) {
	wVProxySettingsImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVProxySettings) PacUrl() string {
	r1 := wVProxySettingsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetPacUrl(AValue string) {
	wVProxySettingsImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVProxySettings) Server() string {
	r1 := wVProxySettingsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetServer(AValue string) {
	wVProxySettingsImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

var (
	wVProxySettingsImport       *imports.Imports = nil
	wVProxySettingsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WVProxySettings_AutoDetect", 0),
		/*1*/ imports.NewTable("WVProxySettings_ByPassList", 0),
		/*2*/ imports.NewTable("WVProxySettings_Create", 0),
		/*3*/ imports.NewTable("WVProxySettings_NoProxyServer", 0),
		/*4*/ imports.NewTable("WVProxySettings_PacUrl", 0),
		/*5*/ imports.NewTable("WVProxySettings_Server", 0),
	}
)

func wVProxySettingsImportAPI() *imports.Imports {
	if wVProxySettingsImport == nil {
		wVProxySettingsImport = NewDefaultImports()
		wVProxySettingsImport.SetImportTable(wVProxySettingsImportTables)
		wVProxySettingsImportTables = nil
	}
	return wVProxySettingsImport
}
