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

// IWkResponsePolicyDecision Root Interface
type IWkResponsePolicyDecision interface {
	IObject
	GetRequest() WebKitURIRequest   // function
	GetResponse() WebKitURIResponse // function
}

// TWkResponsePolicyDecision Root Object
type TWkResponsePolicyDecision struct {
	TObject
}

func NewWkResponsePolicyDecision(aDecision WebKitResponsePolicyDecision) IWkResponsePolicyDecision {
	r1 := wkResponsePolicyDecisionImportAPI().SysCallN(0, uintptr(aDecision))
	return AsWkResponsePolicyDecision(r1)
}

func (m *TWkResponsePolicyDecision) GetRequest() WebKitURIRequest {
	r1 := wkResponsePolicyDecisionImportAPI().SysCallN(1, m.Instance())
	return WebKitURIRequest(r1)
}

func (m *TWkResponsePolicyDecision) GetResponse() WebKitURIResponse {
	r1 := wkResponsePolicyDecisionImportAPI().SysCallN(2, m.Instance())
	return WebKitURIResponse(r1)
}

var (
	wkResponsePolicyDecisionImport       *imports.Imports = nil
	wkResponsePolicyDecisionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WkResponsePolicyDecision_Create", 0),
		/*1*/ imports.NewTable("WkResponsePolicyDecision_GetRequest", 0),
		/*2*/ imports.NewTable("WkResponsePolicyDecision_GetResponse", 0),
	}
)

func wkResponsePolicyDecisionImportAPI() *imports.Imports {
	if wkResponsePolicyDecisionImport == nil {
		wkResponsePolicyDecisionImport = NewDefaultImports()
		wkResponsePolicyDecisionImport.SetImportTable(wkResponsePolicyDecisionImportTables)
		wkResponsePolicyDecisionImportTables = nil
	}
	return wkResponsePolicyDecisionImport
}
