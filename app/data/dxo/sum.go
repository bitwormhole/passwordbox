package dxo

import "github.com/starter-go/base/lang"

////////////////////////////////////////////////////////////////////////////////

type SHA256SUM [32]byte

func (d Digest) String() string {
	bin := d[:]
	hex := lang.HexFromBytes(bin)
	return hex.String()
}

func (d Digest) Bytes() []byte {
	return d[:]
}

////////////////////////////////////////////////////////////////////////////////

// Digest 是一个 sha-256 的 hash 值
type Digest = SHA256SUM

////////////////////////////////////////////////////////////////////////////////

type SUM = SHA256SUM

////////////////////////////////////////////////////////////////////////////////

type FingerPrint SHA256SUM

type CertificateFingerPrint FingerPrint

type PublicKeyFingerPrint FingerPrint

type BlockFingerPrint FingerPrint

////////////////////////////////////////////////////////////////////////////////

type BinaryDataSum SHA256SUM

type BinaryDataRef = BinaryDataSum

////////////////////////////////////////////////////////////////////////////////

func (fp BlockFingerPrint) String() string {
	sum := SHA256SUM(fp)
	return sum.String()
}

func (fp BlockFingerPrint) Bytes() []byte {
	sum := SHA256SUM(fp)
	return sum.Bytes()
}

func (fp BinaryDataSum) String() string {
	sum := SHA256SUM(fp)
	return sum.String()
}

func (fp BinaryDataSum) Bytes() []byte {
	sum := SHA256SUM(fp)
	return sum.Bytes()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
