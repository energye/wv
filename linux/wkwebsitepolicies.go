//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"

	wvTypes "github.com/energye/wv/types/linux"
)

// IWkWebsitePolicies Parent: IObject
type IWkWebsitePolicies interface {
	IObject
	Data() wvTypes.WebKitWebsitePolicies             // function
	GetAutoplayPolicy() wvTypes.WebKitAutoplayPolicy // function
}

type TWkWebsitePolicies struct {
	TObject
}

func (m *TWkWebsitePolicies) Data() wvTypes.WebKitWebsitePolicies {
	if !m.IsValid() {
		return 0
	}
	r := wkWebsitePoliciesAPI().SysCallN(2, m.Instance())
	return wvTypes.WebKitWebsitePolicies(r)
}

func (m *TWkWebsitePolicies) GetAutoplayPolicy() wvTypes.WebKitAutoplayPolicy {
	if !m.IsValid() {
		return 0
	}
	r := wkWebsitePoliciesAPI().SysCallN(3, m.Instance())
	return wvTypes.WebKitAutoplayPolicy(r)
}

// NewWebsitePolicies class constructor
func NewWebsitePolicies() IWkWebsitePolicies {
	r := wkWebsitePoliciesAPI().SysCallN(0)
	return AsWkWebsitePolicies(r)
}

// NewWebsitePoliciesWithStr class constructor
func NewWebsitePoliciesWithStr(firstPolicyName string) IWkWebsitePolicies {
	r := wkWebsitePoliciesAPI().SysCallN(1, api.PasStr(firstPolicyName))
	return AsWkWebsitePolicies(r)
}

var (
	wkWebsitePoliciesOnce   base.Once
	wkWebsitePoliciesImport *imports.Imports = nil
)

func wkWebsitePoliciesAPI() *imports.Imports {
	wkWebsitePoliciesOnce.Do(func() {
		wkWebsitePoliciesImport = api.NewDefaultImports()
		wkWebsitePoliciesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkWebsitePolicies_Create", 0), // constructor NewWebsitePolicies
			/* 1 */ imports.NewTable("TWkWebsitePolicies_CreateWithStr", 0), // constructor NewWebsitePoliciesWithStr
			/* 2 */ imports.NewTable("TWkWebsitePolicies_Data", 0), // function Data
			/* 3 */ imports.NewTable("TWkWebsitePolicies_GetAutoplayPolicy", 0), // function GetAutoplayPolicy
		}
	})
	return wkWebsitePoliciesImport
}
