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

// IWkNavigationPolicyDecision Root Interface
type IWkNavigationPolicyDecision interface {
	IObject
	GetNavigationAction() WebKitNavigationAction // function
}

// TWkNavigationPolicyDecision Root Object
type TWkNavigationPolicyDecision struct {
	TObject
}

func NewWkNavigationPolicyDecision(aDecision WebKitNavigationPolicyDecision) IWkNavigationPolicyDecision {
	r1 := wkNavigationPolicyDecisionImportAPI().SysCallN(0, uintptr(aDecision))
	return AsWkNavigationPolicyDecision(r1)
}

func (m *TWkNavigationPolicyDecision) GetNavigationAction() WebKitNavigationAction {
	r1 := wkNavigationPolicyDecisionImportAPI().SysCallN(1, m.Instance())
	return WebKitNavigationAction(r1)
}

var (
	wkNavigationPolicyDecisionImport       *imports.Imports = nil
	wkNavigationPolicyDecisionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkNavigationPolicyDecision_Create", 0),
		/*1*/ imports.NewTable("WkNavigationPolicyDecision_GetNavigationAction", 0),
	}
)

func wkNavigationPolicyDecisionImportAPI() *imports.Imports {
	if wkNavigationPolicyDecisionImport == nil {
		wkNavigationPolicyDecisionImport = NewDefaultImports()
		wkNavigationPolicyDecisionImport.SetImportTable(wkNavigationPolicyDecisionImportTables)
		wkNavigationPolicyDecisionImportTables = nil
	}
	return wkNavigationPolicyDecisionImport
}
