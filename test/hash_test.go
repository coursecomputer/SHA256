package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/coursecomputer/SHA256/source"
)

// echo -n "toto" | shasum -a 256

func Test_hash(t *testing.T) {
	var sha256 = SHA256.Init()

	sha256.Update([]byte("toto"))
	hash := sha256.Digest()

	assert.Equal(t, "31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66", hash)
}

func Test_hash2(t *testing.T) {
	var sha256 = SHA256.Init()

	sha256.Update([]byte(""))
	hash := sha256.Digest()

	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", hash)
}

func Test_hash3(t *testing.T) {
	var sha256 = SHA256.Init()

	sha256.Update([]byte("wqbjdbjkwqbdqwbd2787321783   j3k232  2 22';.';.,;21  2"))
	hash := sha256.Digest()

	assert.Equal(t, "1324312f2f419acabb8928f0c7317f03e831ee376b901d47d0984784088765fb", hash)
}

func Test_hash4(t *testing.T) {
	var sha256 = SHA256.Init()

	sha256.Update([]byte("tvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjv"))
	hash := sha256.Digest()

	assert.Equal(t, "54a467caf57660e8bab26a7d1882aa2ec82a027570c6a9dd8783dc246c3d0eed", hash)
}

func Test_hash5(t *testing.T) {
	var sha256 = SHA256.Init()

	sha256.Update([]byte("tvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhje"))
	sha256.Update([]byte("qwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjvtvdhjeqwvhjdkwqdkjhefdvehwjvefhewjv"))
	hash := sha256.Digest()

	assert.Equal(t, "54a467caf57660e8bab26a7d1882aa2ec82a027570c6a9dd8783dc246c3d0eed", hash)
}