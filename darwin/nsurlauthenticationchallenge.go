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
//
//	A challenge from a server requiring authentication from the client.
//	https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge?language=objc
//	No Impl
type INSURLAuthenticationChallenge interface {
	IObject
	// Data
	//  Returns the object implemented by this class.
	Data() NSURLAuthenticationChallenge // function
	// Release
	//  Freeing the class and the objects it implements.
	Release() // procedure
}

// TNSURLAuthenticationChallenge Root Object
//
//	A challenge from a server requiring authentication from the client.
//	https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge?language=objc
//	No Impl
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

func (m *TNSURLAuthenticationChallenge) Release() {
	nSURLAuthenticationChallengeImportAPI().SysCallN(2, m.Instance())
}

var (
	nSURLAuthenticationChallengeImport       *imports.Imports = nil
	nSURLAuthenticationChallengeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("NSURLAuthenticationChallenge_Create", 0),
		/*1*/ imports.NewTable("NSURLAuthenticationChallenge_Data", 0),
		/*2*/ imports.NewTable("NSURLAuthenticationChallenge_Release", 0),
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
