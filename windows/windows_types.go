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
	"github.com/energye/lcl/types"
)

type TWVCustomSchemeInfo struct {
	SchemeName            string         // wvstring
	TreatAsSecure         types.LongBool // boolean
	AllowedDomains        string         // wvstring
	HasAuthorityComponent types.LongBool // boolean
}

type TWVWindowFeatures struct {
	HasPosition             types.LongBool // boolean
	HasSize                 types.LongBool // boolean
	Left                    uint32         // cardinal
	Top                     uint32         // cardinal
	Width                   uint32         // cardinal
	Height                  uint32         // cardinal
	ShouldDisplayMenuBar    types.LongBool // boolean
	ShouldDisplayStatus     types.LongBool // boolean
	ShouldDisplayToolbar    types.LongBool // boolean
	ShouldDisplayScrollBars types.LongBool // boolean
}

func (m *TWVCustomSchemeInfo) ToPas() *tWVCustomSchemeInfo {
	if m == nil {
		return nil
	}
	return &tWVCustomSchemeInfo{
		SchemeName:            api.PasStr(m.SchemeName),
		TreatAsSecure:         m.TreatAsSecure,
		AllowedDomains:        api.PasStr(m.AllowedDomains),
		HasAuthorityComponent: m.HasAuthorityComponent,
	}
}

type tWVCustomSchemeInfo struct {
	SchemeName            uintptr        // wvstring
	TreatAsSecure         types.LongBool // boolean
	AllowedDomains        uintptr        // wvstring
	HasAuthorityComponent types.LongBool // boolean
}

func (m *tWVCustomSchemeInfo) ToGo() TWVCustomSchemeInfo {
	if m == nil {
		return TWVCustomSchemeInfo{}
	}
	return TWVCustomSchemeInfo{
		SchemeName:            api.GoStr(m.SchemeName),
		TreatAsSecure:         m.TreatAsSecure,
		AllowedDomains:        api.GoStr(m.AllowedDomains),
		HasAuthorityComponent: m.HasAuthorityComponent,
	}
}

type ICoreWebView2PermissionRequestedEventArgs2 uintptr

// PPCoreWebView2CustomSchemeRegistration = ^ICoreWebView2CustomSchemeRegistration
type PPCoreWebView2CustomSchemeRegistration uintptr
