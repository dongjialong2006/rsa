# RSA

公钥私钥生成器

## 说明
该模块用于生成rsa的公钥和私钥，可根据用户数的key生成固定的rsa秘钥。

## 环境配置

- 本项目编译之前，需用户自行安装`Golang`、`github`和`glide`工具；
- 在编译时，本项目会先从`github`上拉去以来的项目，编译时间视具体网络环境而定；
- 要求`linux`或`macos`系统即可。

## 引用

- import "github.com/dongjialong2006/rsa"


## 使用

{
	import "fmt"
	import "github.com/dongjialong2006/rsa"
	
	func main() {
		priv, pub, _ := rsa.MakeSSHKeyPair("", 2048)
		fmt.Println(pub)
		
		priv, pub, _ := rsa.MakeSSHKeyPair("test", 2048)
		fmt.Println(pub)
	}
}

