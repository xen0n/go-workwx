//+build sdkcodegen

package main

var errcodeNameMap = map[int64]string{
	// this must be renamed otherwise `ErrCode-1` is illegal ident
	-1: "ServiceUnavailable",
	// this is well-known enough to warrant a name
	0: "Success",
}
