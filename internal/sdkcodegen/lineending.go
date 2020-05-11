//+build sdkcodegen

package main

import (
	"bytes"
)

// normalizeEOL will convert Windows (CRLF) and Mac (CR) EOLs to UNIX (LF)
//
// NOTE: this is https://github.com/go-gitea/gitea/blob/dc8036dcc680abab52b342d18181a5ee42f40318/modules/util/util.go#L68-L102
//
// Gitea is licensed under MIT, so is this function:
//
// Copyright (c) 2016 The Gitea Authors
// Copyright (c) 2015 The Gogs Authors
//
// We're MIT-licensed too, so referencing it won't be a problem!
func normalizeEOL(input []byte) []byte {
	var right, left, pos int
	if right = bytes.IndexByte(input, '\r'); right == -1 {
		return input
	}
	length := len(input)
	tmp := make([]byte, length)

	// We know that left < length because otherwise right would be -1 from IndexByte.
	copy(tmp[pos:pos+right], input[left:left+right])
	pos += right
	tmp[pos] = '\n'
	left += right + 1
	pos++

	for left < length {
		if input[left] == '\n' {
			left++
		}

		right = bytes.IndexByte(input[left:], '\r')
		if right == -1 {
			copy(tmp[pos:], input[left:])
			pos += length - left
			break
		}
		copy(tmp[pos:pos+right], input[left:left+right])
		pos += right
		tmp[pos] = '\n'
		left += right + 1
		pos++
	}
	return tmp[:pos]
}
