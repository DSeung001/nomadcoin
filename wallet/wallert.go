package wallet

1) we hash the msg.
	"i love you" -> hash(x) -> "hashed_message"

2) generate key pair
	KeyPair (privateKey, publicKey)
	(save privateKey to afile(wallet))

3) sign the hash
	("hashed_message" + privateKey) -> "signature"

4) verfiy
	("hashed_message" + "signature" + publicKey) -> true|false