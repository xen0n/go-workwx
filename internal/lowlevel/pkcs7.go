package lowlevel

func pkcs7Unpad(x []byte) []byte {
	// last byte is number of suffix bytes to remove
	n := int(x[len(x)-1])
	return x[:len(x)-n]
}
