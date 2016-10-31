// +build amd64 arm64

package sha256

func block(dig *digest, p []byte) {
	switch {
	case avx2:
		blockAvx2Go(dig, p)
	case avx:
		blockAvxGo(dig, p)
	case ssse3:
		blockSsseGo(dig, p)
	case armSha:
		blockArmGo(dig, p)
	default:
		blockGeneric(dig, p)
	}
}
