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

// IWkPolicyDecision Root Interface
type IWkPolicyDecision interface {
	IObject
	Data() WebKitPolicyDecision                     // function
	Use()                                           // procedure
	UseWithPolicies(policies WebKitWebsitePolicies) // procedure
	Ignore()                                        // procedure
	Download()                                      // procedure
}

// TWkPolicyDecision Root Object
type TWkPolicyDecision struct {
	TObject
}

func NewWkPolicyDecision(aDecision WebKitPolicyDecision) IWkPolicyDecision {
	r1 := wkPolicyDecisionImportAPI().SysCallN(0, uintptr(aDecision))
	return AsWkPolicyDecision(r1)
}

func (m *TWkPolicyDecision) Data() WebKitPolicyDecision {
	r1 := wkPolicyDecisionImportAPI().SysCallN(1, m.Instance())
	return WebKitPolicyDecision(r1)
}

func (m *TWkPolicyDecision) Use() {
	wkPolicyDecisionImportAPI().SysCallN(4, m.Instance())
}

func (m *TWkPolicyDecision) UseWithPolicies(policies WebKitWebsitePolicies) {
	wkPolicyDecisionImportAPI().SysCallN(5, m.Instance(), uintptr(policies))
}

func (m *TWkPolicyDecision) Ignore() {
	wkPolicyDecisionImportAPI().SysCallN(3, m.Instance())
}

func (m *TWkPolicyDecision) Download() {
	wkPolicyDecisionImportAPI().SysCallN(2, m.Instance())
}

var (
	wkPolicyDecisionImport       *imports.Imports = nil
	wkPolicyDecisionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkPolicyDecision_Create", 0),
		/*1*/ imports.NewTable("WkPolicyDecision_Data", 0),
		/*2*/ imports.NewTable("WkPolicyDecision_Download", 0),
		/*3*/ imports.NewTable("WkPolicyDecision_Ignore", 0),
		/*4*/ imports.NewTable("WkPolicyDecision_Use", 0),
		/*5*/ imports.NewTable("WkPolicyDecision_UseWithPolicies", 0),
	}
)

func wkPolicyDecisionImportAPI() *imports.Imports {
	if wkPolicyDecisionImport == nil {
		wkPolicyDecisionImport = NewDefaultImports()
		wkPolicyDecisionImport.SetImportTable(wkPolicyDecisionImportTables)
		wkPolicyDecisionImportTables = nil
	}
	return wkPolicyDecisionImport
}
