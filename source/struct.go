package SHA256

type SHA256 struct {
	states [8]uint32
	remain []byte
	size int
}

func Init() SHA256 {
	return SHA256{
		size: 0,
		remain: nil,
		states: [8]uint32{
			0x6a09e667,
			0xbb67ae85,
			0x3c6ef372,
			0xa54ff53a,
			0x510e527f,
			0x9b05688c,
			0x1f83d9ab,
			0x5be0cd19,
		},
	}
}

func (s * SHA256) Reset() {
	s.size = 0
	s.remain = nil
	s.states = [8]uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}
}
