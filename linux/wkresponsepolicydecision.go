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

// IWkResponsePolicyDecision Parent: IWkPolicyDecision
type IWkResponsePolicyDecision interface {
	IWkPolicyDecision
	GetRequest() wvTypes.WebKitURIRequest   // function
	GetResponse() wvTypes.WebKitURIResponse // function
}

type TWkResponsePolicyDecision struct {
	TWkPolicyDecision
}

func (m *TWkResponsePolicyDecision) GetRequest() wvTypes.WebKitURIRequest {
	if !m.IsValid() {
		return 0
	}
	r := wkResponsePolicyDecisionAPI().SysCallN(1, m.Instance())
	return wvTypes.WebKitURIRequest(r)
}

func (m *TWkResponsePolicyDecision) GetResponse() wvTypes.WebKitURIResponse {
	if !m.IsValid() {
		return 0
	}
	r := wkResponsePolicyDecisionAPI().SysCallN(2, m.Instance())
	return wvTypes.WebKitURIResponse(r)
}

// NewResponsePolicyDecision class constructor
func NewResponsePolicyDecision(decision wvTypes.WebKitResponsePolicyDecision) IWkResponsePolicyDecision {
	r := wkResponsePolicyDecisionAPI().SysCallN(0, uintptr(decision))
	return AsWkResponsePolicyDecision(r)
}

var (
	wkResponsePolicyDecisionOnce   base.Once
	wkResponsePolicyDecisionImport *imports.Imports = nil
)

func wkResponsePolicyDecisionAPI() *imports.Imports {
	wkResponsePolicyDecisionOnce.Do(func() {
		wkResponsePolicyDecisionImport = api.NewDefaultImports()
		wkResponsePolicyDecisionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWkResponsePolicyDecision_Create", 0), // constructor NewResponsePolicyDecision
			/* 1 */ imports.NewTable("TWkResponsePolicyDecision_GetRequest", 0), // function GetRequest
			/* 2 */ imports.NewTable("TWkResponsePolicyDecision_GetResponse", 0), // function GetResponse
		}
	})
	return wkResponsePolicyDecisionImport
}
