package SHA256

// internal function
func rotr(x, n uint32) uint32 {
	return ((x >> n) | (x << (32 - n))) & 0xFFFFFFFF
}


func ch(x, y, z uint32) uint32 {
	return (x & y) ^ ((^x) & z)
}

func maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func sum0(x uint32) uint32 {
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func sum1(x uint32) uint32 {
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func sigma0(x uint32) uint32 {
	return rotr(x, 7) ^ rotr(x, 18) ^ (x >> 3)
}

func sigma1(x uint32) uint32 {
	return rotr(x, 17) ^ rotr(x, 19) ^ (x >> 10)
}