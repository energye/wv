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

// IWkPolicyDecision Parent: IObject
type IWkPolicyDecision interface {
	IObject
	Data() wvTypes.WebKitPolicyDecision                     // function
	Use()                                                   // procedure
	UseWithPolicies(policies wvTypes.WebKitWebsitePolicies) // procedure
	Ignore()                                                // procedure
	Download()                                              // procedure
}

type TWkPolicyDecision struct {
	TObject
}

func (m *TWkPolicyDecision) Data() wvTypes.WebKitPolicyDecision {
	if !m.IsValid() {
		return 0
	}
	r := wkPolicyDecisionAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitPolicyDecision(r)
}

func (m *TWkPolicyDecision) Use() {
	if !m.IsValid() {
		return
	}
	wkPolicyDecisionAPI().SysCallN(2, m.Instance())
}

func (m *TWkPolicyDecision) UseWithPolicies(policies wvTypes.WebKitWebsitePolicies) {
	if !m.IsValid() {
		return
	}
	wkPolicyDecisionAPI().SysCallN(3, m.Instance(), uintptr(policies))
}

func (m *TWkPolicyDecision) Ignore() {
	if !m.IsValid() {
		return
	}
	wkPolicyDecisionAPI().SysCallN(4, m.Instance())
}

func (m *TWkPolicyDecision) Download() {
	if !m.IsValid() {
		return
	}
	wkPolicyDecisionAPI().SysCallN(5, m.Instance())
}

// NewPolicyDecision class constructor
func NewPolicyDecision(decision wvTypes.WebKitPolicyDecision) IWkPolicyDecision {
	r := wkPolicyDecisionAPI().SysCallN(0, uintptr(decision))
	return AsWkPolicyDecision(r)
}

var (
	wkPolicyDecisionOnce   base.Once
	wkPolicyDecisionImport *imports.Imports = nil
)

func wkPolicyDecisionAPI() *imports.Imports {
	wkPolicyDecisionOnce.Do(func() {
		wkPolicyDecisionImport = api.NewDefaultImports()
		wkPolicyDecisionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkPolicyDecision_Create", 0), // constructor NewPolicyDecision
			/* 1 */ imports.NewTable("TWkPolicyDecision_Data", 0), // function Data
			/* 2 */ imports.NewTable("TWkPolicyDecision_Use", 0), // procedure Use
			/* 3 */ imports.NewTable("TWkPolicyDecision_UseWithPolicies", 0), // procedure UseWithPolicies
			/* 4 */ imports.NewTable("TWkPolicyDecision_Ignore", 0), // procedure Ignore
			/* 5 */ imports.NewTable("TWkPolicyDecision_Download", 0), // procedure Download
		}
	})
	return wkPolicyDecisionImport
}
