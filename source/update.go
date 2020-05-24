package SHA256

import (
	"encoding/binary"
	"math"
)

func createWordsArray(data []byte) [64]uint32 {
	var result [64]uint32

	for idx := 0; idx < 16; idx++ {
		result[idx] = binary.BigEndian.Uint32(data[idx*4 : (idx*4) + 4])
	}

	for idx := 16; idx < 64; idx++ {
		result[idx] = result[idx-16] + sigma0(result[idx-15]) + result[idx-7] + sigma1(result[idx-2])
	}

	return result
}

func (s *SHA256) hashing(words [64]uint32) {
	var t1, t2 uint32

	var a = s.states[0]
	var b = s.states[1]
	var c = s.states[2]
	var d = s.states[3]
	var e = s.states[4]
	var f = s.states[5]
	var g = s.states[6]
	var h = s.states[7]

	for idx := 0; idx < 64; idx++ {
		t1 = h + sum1(e) + ch(e, f, g) + iterativeConstantes[idx] + words[idx]
		t2 = sum0(a) + maj(a, b, c)
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}

	s.states[0] += a
	s.states[1] += b
	s.states[2] += c
	s.states[3] += d
	s.states[4] += e
	s.states[5] += f
	s.states[6] += g
	s.states[7] += h
}

func (s *SHA256) Update(data []byte) {
	if s.remain != nil {
		data = append(s.remain, data...)
	}

	var lengthData = len(data)
	var numOfLoop = int(math.Floor(float64(lengthData) / 64))
	var idx = 0

	for ; idx < numOfLoop; idx++ {
		chunk := data[idx*64 : (idx*64) + 64]
		words := createWordsArray(chunk)
		s.hashing(words)
		s.size += 64
	}

	s.remain = data[idx*64 :]
}
