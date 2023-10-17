tls-config-name := tls_config
tls-config := $(shell pwd)/$(tls-config-name)
tls-config-server = $(tls-config)/server
tls-config-client = $(tls-config)/client
.PHONY: ca server client clean

ca:
	echo $(tls-config)
	if [ -d $(tls-config) ]; then echo "文件夹已存在"; else mkdir $(tls-config); fi

	echo "创建CA密钥.."
	openssl genrsa -out $(tls-config)/CA.key 2048
	# ca
	# 生成ca证书 ca.key是私钥，服务管理员必须好好保管
	echo "生成CA证书"
	openssl req -days 365 -new -x509 -key $(tls-config)/CA.key \
		-subj "/C=CN/ST=GD/L=SZ/O=Hello/CN=Hello Root CA" -out $(tls-config)/CA.crt -utf8

ca_server:
	echo "创建服务端密钥.."
	if [ -d $(tls-config-server) ]; then echo "文件夹已存在"; else mkdir $(tls-config-server); fi
	 # 生成网站证书请求
     # CN 一定要是网站的域名, 否则会通不过安全验证
	echo "创建网站(服务端)证书.."
	openssl req -newkey rsa:2048 -nodes -keyout $(tls-config-server)/server.key \
    	-subj "/C=CN/ST=GD/L=SZ/O=Hello/CN=hellosvc.com" -addext "subjectAltName = DNS:hellosvc.com" -out $(tls-config-server)/server.csr -utf8
	# CA签名网站证书请求
        # 不是拿到 CA 的证书了就可以说自己是 CA 的, 最重要的是, 签名需要有 CA 密钥
        # 如果客户端（个人浏览器）信任 CA 的证书的话, 那么他也就会信任由 CA 签名的网站证书
        # 因此让浏览器信任 CA 的证书之后, 客户端就自然信任服务端了, 只要做单向认证的话, 到这一步证书这一类材料就已经准备好了
        # 但是双向认证就还要给客户端（个人的浏览器）准备一份证书
        # 让服务端可以知道客户端也是合法的。
        # 假如让服务端也信任 CA 的证书
        # 那 CA 签名的客户端证书也就能被信任了。
	echo "通过CA证书签名, "
	openssl x509 -days 365 -req -extfile <(printf "subjectAltName=DNS:hellosvc.com") -in $(tls-config-server)/server.csr \
		-CA $(tls-config)/CA.crt -CAkey $(tls-config)/CA.key -CAcreateserial -out $(tls-config-server)/server.crt

ca_client:
	echo "创建客户端密钥.."
	if [ -d $(tls-config-client) ]; then echo "文件夹已存在"; else mkdir $(tls-config-client); fi
	openssl genrsa -out $(tls-config-client)/client.key 2048


	echo "生成浏览器(客户端)证书请求.."
	openssl req -new -key $(tls-config-client)/client.key -subj "/C=CN/ST=GD/L=SZ/O=HelloClient/CN=client.com/emailAddress=client@hello.com" \
		-out $(tls-config-client)/client.csr -utf8


	echo "通过CA证书签名, 创建浏览器(客户端)证书.."
	openssl x509 -days 365 -req -CA $(tls-config)/CA.crt -CAkey $(tls-config)/CA.key -CAcreateserial \
		-in $(tls-config-client)/client.csr -out $(tls-config-client)/client.crt


clean:
	rm -rf $(tls-config)

