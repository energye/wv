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

// INSURLAuthenticationChallenge Root Interface
type INSURLAuthenticationChallenge interface {
	IObject
	Data() NSURLAuthenticationChallenge // function
}

// TNSURLAuthenticationChallenge Root Object
type TNSURLAuthenticationChallenge struct {
	TObject
}

func NewNSURLAuthenticationChallenge(aData NSURLAuthenticationChallenge) INSURLAuthenticationChallenge {
	r1 := nSURLAuthenticationChallengeImportAPI().SysCallN(0, uintptr(aData))
	return AsNSURLAuthenticationChallenge(r1)
}

func (m *TNSURLAuthenticationChallenge) Data() NSURLAuthenticationChallenge {
	r1 := nSURLAuthenticationChallengeImportAPI().SysCallN(1, m.Instance())
	return NSURLAuthenticationChallenge(r1)
}

var (
	nSURLAuthenticationChallengeImport       *imports.Imports = nil
	nSURLAuthenticationChallengeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLAuthenticationChallenge_Create", 0),
		/*1*/ imports.NewTable("NSURLAuthenticationChallenge_Data", 0),
	}
)

func nSURLAuthenticationChallengeImportAPI() *imports.Imports {
	if nSURLAuthenticationChallengeImport == nil {
		nSURLAuthenticationChallengeImport = NewDefaultImports()
		nSURLAuthenticationChallengeImport.SetImportTable(nSURLAuthenticationChallengeImportTables)
		nSURLAuthenticationChallengeImportTables = nil
	}
	return nSURLAuthenticationChallengeImport
}
