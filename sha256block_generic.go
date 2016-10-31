// +build 386,arm

package sha256

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
