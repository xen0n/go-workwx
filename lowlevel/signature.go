package lowlevel

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

func makeDevMsgSignature(paramValues ...string) string {
	tmp := make([]string, len(paramValues))
	for i, x := range paramValues {
		tmp[i] = x
	}

	sort.Strings(tmp)

	state := sha1.New()
	for _, x := range tmp {
		_, _ = state.Write([]byte(x))
	}

	result := state.Sum(nil)
	return fmt.Sprintf("%x", result)
}
