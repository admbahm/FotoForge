package hash

import (
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		blake3 string
		sha256 string
	}{
		{
			name:   "empty",
			input:  "",
			blake3: "af1349b9f5f9a1a6a0404dea36dcc9499bcb25c9adc112b7cc9a93cae41f3262",
			sha256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:   "project name",
			input:  "fotoforge",
			blake3: "1e0d58583d3186d6b9a66f55956ac73732e371317a16864e19a12e59137b85e5",
			sha256: "9247430fe82fc566fd6a8a23ac73c32a65321e945e9f3449b57dc68413b95c8a",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := Reader(strings.NewReader(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if got.BLAKE3 != test.blake3 {
				t.Errorf("BLAKE3 = %q, want %q", got.BLAKE3, test.blake3)
			}
			if got.SHA256 != test.sha256 {
				t.Errorf("SHA-256 = %q, want %q", got.SHA256, test.sha256)
			}
		})
	}
}
