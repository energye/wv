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

// IWkNavigationPolicyDecision Parent: IWkPolicyDecision
type IWkNavigationPolicyDecision interface {
	IWkPolicyDecision
	GetNavigationAction() wvTypes.WebKitNavigationAction // function
}

type TWkNavigationPolicyDecision struct {
	TWkPolicyDecision
}

func (m *TWkNavigationPolicyDecision) GetNavigationAction() wvTypes.WebKitNavigationAction {
	if !m.IsValid() {
		return 0
	}
	r := wkNavigationPolicyDecisionAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitNavigationAction(r)
}

// NewNavigationPolicyDecision class constructor
func NewNavigationPolicyDecision(decision wvTypes.WebKitNavigationPolicyDecision) IWkNavigationPolicyDecision {
	r := wkNavigationPolicyDecisionAPI().SysCallN(0, uintptr(decision))
	return AsWkNavigationPolicyDecision(r)
}

var (
	wkNavigationPolicyDecisionOnce   base.Once
	wkNavigationPolicyDecisionImport *imports.Imports = nil
)

func wkNavigationPolicyDecisionAPI() *imports.Imports {
	wkNavigationPolicyDecisionOnce.Do(func() {
		wkNavigationPolicyDecisionImport = api.NewDefaultImports()
		wkNavigationPolicyDecisionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkNavigationPolicyDecision_Create", 0), // constructor NewNavigationPolicyDecision
			/* 1 */ imports.NewTable("TWkNavigationPolicyDecision_GetNavigationAction", 0), // function GetNavigationAction
		}
	})
	return wkNavigationPolicyDecisionImport
}
