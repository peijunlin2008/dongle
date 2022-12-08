package dongle

import (
	"fmt"
)

var (
	// returns an invalid aes src error
	// 返回无效的 aes 明文错误
	invalidAesSrcError = func() error {
		return fmt.Errorf("aes: invalid src, the src is not full blocks")
	}
	// returns an invalid aes key error
	// 返回无效的 aes 密钥错误
	invalidAesKeyError = func() error {
		return fmt.Errorf("aes: invalid key, the key must be 16, 24 or 32 bytes")
	}
	// returns an invalid aes iv error、
	// 返回无效的 aes 偏移量错误
	invalidAesIVError = func() error {
		return fmt.Errorf("aes: invalid iv, the iv size must be 16 bytes")
	}
)

var (
	// returns an invalid des src error
	// 返回无效的 des 明文错误
	invalidDesSrcError = func() error {
		return fmt.Errorf("des: invalid src, the src is not full blocks")
	}
	// returns an invalid des key error
	// 返回无效的 des 密钥错误
	invalidDesKeyError = func() error {
		return fmt.Errorf("des: invalid key, the key must be 8 bytes")
	}
	// returns an invalid des iv error、
	// 返回无效的 des 偏移量错误
	invalidDesIVError = func() error {
		return fmt.Errorf("des: invalid iv, the iv size must be 8 bytes")
	}
)

var (
	// returns an invalid 3des src error
	// 返回无效的 3des 明文错误
	invalid3DesSrcError = func() error {
		return fmt.Errorf("3des: invalid src, the src is not full blocks")
	}
	// returns an invalid 3des key error
	// 返回无效的 3des 密钥错误
	invalid3DesKeyError = func() error {
		return fmt.Errorf("3des: invalids key, the key must be 24 bytes")
	}
	// returns an invalid 3des iv error、
	// 返回无效的 3des 偏移量错误
	invalid3DesIVError = func() error {
		return fmt.Errorf("3des: invalid iv, the iv size must be 8 bytes")
	}
)

var (
	// returns an invalid rsa public key error
	// 返回无效的 rsa 公钥错误
	invalidRsaPublicKeyError = func() error {
		return fmt.Errorf("rsa: invalid public key, please make sure the public key is valid")
	}
	// returns an invalid rsa private key error
	// 返回无效的 rsa 私钥错误
	invalidRsaPrivateKeyError = func() error {
		return fmt.Errorf("rsa: invalid private key, please make sure the private key is valid")
	}
	// returns an invalid rsa hash function error
	// 返回无效的 rsa 哈希函数错误
	invalidRsaHashError = func() error {
		return fmt.Errorf("rsa: invalid hash function, the hash function is unsupported")
	}
)

var (
	// returns an invalid tea src error
	// 返回无效的 tea 明文错误
	invalidTeaSrcError = func() error {
		return fmt.Errorf("tea: invalid src, the src must be 8 bytes")
	}
	// returns an invalid tea round error
	// 返回无效的 tea 迭代轮数错误
	invalidTeaRoundsError = func() error {
		return fmt.Errorf("tea: invalid rounds, the rounds must be even")
	}
	// returns an invalid tea key error
	// 返回无效的 tea 密钥错误
	invalidTeaKeyError = func() error {
		return fmt.Errorf("tea: invalid key, the key must be 16 bytes")
	}
)

var (
	// returns an invalid morse src error
	// 返回无效的 morse 明文错误
	invalidMorseSrcError = func() error {
		return fmt.Errorf("morse: invalid src, the src can't contain spaces")
	}
	// returns a morse decoding error
	// 返回 morse 解码错误
	morseDecodeError = func() error {
		return fmt.Errorf("morse: invalid decoding, the src can't be decoded")
	}
)

var (
	// returns an invalid bcrypt rounds error
	// 返回无效的 bcrypt 迭代轮数错误
	invalidBcryptRoundsError = func() error {
		return fmt.Errorf("bcrypt: invalid rounds, the rounds is outside allowed range (4,31)")
	}
)
var (
	// returns an invalid rc4 key error
	// 返回无效的 rc4 密钥错误
	invalidRc4KeyError = func() error {
		return fmt.Errorf("rc4: invalid key, the key at least 1 byte and at most 256 bytes")
	}
)
var (
	// returns an invalid hash size error
	// 返回无效的哈希大小错误
	invalidHashSizeError = func() error {
		return fmt.Errorf("hash: invalid size, the size is unsupported")
	}
)

var (
	// returns an invalid cipher mode error
	// 返回无效的分组模式错误
	invalidModeError = func(mode cipherMode) error {
		return fmt.Errorf("invalid cipher mode %q, currently only support CBC, ECB, CFB, OFB, CTR mode", mode)
	}
	// returns an invalid cipher padding error
	// 返回无效的填充方式错误
	invalidPaddingError = func(padding cipherPadding) error {
		return fmt.Errorf("invalid cipher padding %q, currently only support No, Zero, PKCS5, PKCS7 padding", padding)
	}
	// returns an invalid decoding error
	// 返回无效的解码方式错误
	invalidDecodingError = func(mode string) error {
		return fmt.Errorf("invalid decoding, the src can't be decoded by %s", mode)
	}
)
