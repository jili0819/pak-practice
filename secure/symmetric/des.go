package symmetric

// 注：
// ecb/ctr:不需要初始化向量（go接口中的iv可以理解为随机数种子, iv的长度 == 明文分组的长度）
// cbc/cfb/ofb:需要初始化向量（des/3des->8字节，aes->16字节，加解密向量相同）
// 如果使用ecb/cbc分组模式需要对密文分组进行填充,cfb/ofb/ctr都不需要对分组密文填充

// 密文先分组，再加密
// des/3des按8字节分组
// 分组密码模式(ecb/cbc模式需要填充 )

// 1、创建des/3des/aes密码接口
// crypto/des   (des/3des)
// crypto/aes   (aes)
// NewCipher(key []byte)(cipher.Block,error)

// 2、如果使用cbc/ecb分组模式需要对密文分组进行填充,cfb/ofb/ctr都不需要对分组密文填充

// 3、创建分组模式接口对象
// cbc:NewCBCEncrypter(b block,iv []byte) BlockMode 加密
// cfb:NewCFBEncrypter(b Block, iv []byte) Stream 加密
