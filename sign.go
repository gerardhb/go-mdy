package mdy

// Signer 明道云签名
type Signer interface {
	SetSign(string)

	SetAppKey(string)
}
