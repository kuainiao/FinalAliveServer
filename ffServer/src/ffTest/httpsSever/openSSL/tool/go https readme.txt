http://qingwang.blog.51cto.com/505009/160626
http://www.tuicool.com/articles/aymYbmM


=========================================================================
go��http.ListenAndServeTLS��Ҫ�����ر������һ���Ƿ���˵�˽Կ�ļ�·��������һ���Ƿ���˵�����֤���ļ�·��������openssl���ߣ����ǿ����Լ��������˽Կ����ǩ��������֤�顣


openssl genrsa -out server.key 2048
�������ɷ����˽Կ�ļ�server.key������Ĳ���2048��λ��bit����˽Կ�ĳ��ȡ�
openssl���ɵ�˽Կ�а����˹�Կ����Ϣ


openssl rsa -in server.key -out server.key.public
���ǿ��Ը���˽Կ���ɹ�Կ


openssl req -new -x509 -key server.key -out server.crt -days 3650
����Ҳ���Ը���˽Կֱ��������ǩ��������֤��


x509: certificate signed by unknown authority
Ĭ��Ҳ��Ҫ�Է���˴�����������֤�����У��ģ����ͻ�����ʾ�����֤�����ɲ�֪��CAǩ���ġ�
tr := &http.Transport{ 
TLSClientConfig:    &tls.Config{InsecureSkipVerify: true}, 
} 
client := &http.Client{Transport: tr} 


====================================�ͻ�����֤�����=====================================
������self-signed(��ǩ��)֤����˵�����ն˲�û�������self-CA������֤�飬Ҳ����û��CA��Կ��Ҳ��û�а취������֤ ���ǩ��������֤��������Ҫ��дһ�����Զ�self-signed֤�����У��Ľ��ն˳���Ļ�����������Ҫ���ľ��ǽ���һ�������Լ��� CA���ø�CAǩ�����ǵ�server��֤�飬������CA���������֤����ͻ���һ��������



���������Լ���CA����Ҫ����һ��CA˽Կ��һ��CA������֤��:
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt



����server�˵�˽Կ����������֤�����󣬲������ǵ�ca˽Կǩ��server������֤��
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000


CA: 
˽Կ�ļ� ca.key 
����֤�� ca.crt


Server: 
˽Կ�ļ� server.key 
����֤�� server.crt


CSR��Cerificate Signing Request��Ӣ����д����֤�������ļ���Ҳ����֤������������������֤��ʱ��CSP(���ܷ����ṩ��)������˽Կ��ͬʱҲ����֤�������ļ���֤��������ֻҪ��CSR�ļ��ύ��֤��䷢������֤��䷢����ʹ�����֤��˽Կǩ����������֤�鹫Կ�ļ���Ҳ���ǰ䷢���û���֤�顣



client����Ҫ��֤server�˵�����֤�飬���client����ҪԤ�ȼ���ca.crt�������ڷ��������֤���У��
pool := x509.NewCertPool() 
caCertPath := "ca.crt"

caCrt, err := ioutil.ReadFile(caCertPath) 
if err != nil { 
fmt.Println("ReadFile err:", err) 
return 
} 
pool.AppendCertsFromPEM(caCrt)

tr := &http.Transport{ 
TLSClientConfig: &tls.Config{RootCAs: pool}, 
} 
client := &http.Client{Transport: tr} 



====================================�������֤�ͻ���=====================================
�����Ҫ�Կͻ�������֤�����У�飬���ȿͻ�����Ҫ�����Լ���֤��



���ɿͻ��˵�˽Կ��֤��
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=tonybai_cn" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000



�������֤
pool := x509.NewCertPool() 
caCertPath := "ca.crt"

caCrt, err := ioutil.ReadFile(caCertPath) 
if err != nil { 
fmt.Println("ReadFile err:", err) 
return 
} 
pool.AppendCertsFromPEM(caCrt)

s := &http.Server{ 
Addr:    ":8081", 
Handler: &myhandler{}, 
TLSConfig: &tls.Config{ 
ClientCAs:  pool, 
ClientAuth: tls.RequireAndVerifyClientCert , 
}, 
}

err = s.ListenAndServeTLS("server.crt", "server.key") 


�ͻ���
pool := x509.NewCertPool() 
caCertPath := "ca.crt"

caCrt, err := ioutil.ReadFile(caCertPath) 
if err != nil { 
fmt.Println("ReadFile err:", err) 
return 
} 
pool.AppendCertsFromPEM(caCrt)

cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key") 
if err != nil { 
fmt.Println("Loadx509keypair err:", err) 
return 
}

tr := &http.Transport{ 
TLSClientConfig: &tls.Config{ 
RootCAs:      pool, 
Certificates: []tls.Certificate{cliCrt}, 
}, 
} 
client := &http.Client{Transport: tr} 


������server�˵Ĵ�����־�������ƺ���client�˵�client.crt�ļ�������ĳЩ������
$go run server.go 
2015/04/30 22:13:33 http: TLS handshake error from 127.0.0.1:53542: 
tls: client's certificate's extended key usage doesn't permit it to be 
used for client authentication

$go run client.go 
Get error: Get https://localhost:8081: remote error: handshake failure


�������crypto/tls/handshake_server.go��
k := false 
for _, ku := range certs[0].ExtKeyUsage { 
if ku == x509.ExtKeyUsageClientAuth { 
ok = true 
break 
} 
} 
if !ok { 
c.sendAlert(alertHandshakeFailure) 
return nil, errors.New("tls: client's certificate's extended key usage doesn't permit it to be used for client authentication") 
}

�����ж���֤���е�ExtKeyUsage��ϢӦ�ð���clientAuth������openssl��������ϣ��˽⵽��CAǩ��������֤���а����Ķ���һЩbasic����Ϣ������û��ExtKeyUsage����Ϣ


�鿴һ�µ�ǰclient.crt������
openssl x509 -text -in client.crt -noout 


golang��tlsҪУ��ExtKeyUsage�����������Ҫ��������client.crt����������ʱָ��extKeyUsage


�����ļ�client.ext���ļ�����
extendedKeyUsage=clientAuth


�ؽ�client.crt
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000



�ٴβ��ԣ�ok


=========================================================================
(1)openssl genrsa -out rootCA.key 2048 
(2)openssl req -x509 -new -nodes -key rootCA.key -subj "/CN=*.tunnel.tonybai.com" -days 5000 -out rootCA.pem

(3)openssl genrsa -out device.key 2048 
(4)openssl req -new -key device.key -subj "/CN=*.tunnel.tonybai.com" -out device.csr
(5)openssl x509 -req -in device.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out device.crt -days 5000

(6)cp rootCA.pem assets/client/tls/ngrokroot.crt 
(7)cp device.crt assets/server/tls/snakeoil.crt 
(8)cp device.key assets/server/tls/snakeoil.key

�Լ��ngrok���񣬿ͻ���Ҫ��֤�����֤�飬������Ҫ�Լ���CA����˲���(1)�Ͳ���(2)��������CA�Լ��������Ϣ�� 
����(1) ������CA�Լ���˽Կ rootCA.key 
����(2)������CA�Լ���˽Կ������ǩ��������֤�飬��֤�������CA�Լ��Ĺ�Կ��

����(3)~(5)����������ngrok����˵�˽Կ������֤�飨����CAǩ������ 
����(3)������ngrok�����˽Կ�� 
����(4)������Certificate Sign Request��CSR��֤��ǩ������ 
����(5)����CA���Լ���CA˽Կ�Է�����ύ��csr����ǩ�������õ�����˵�����֤��device.crt��

����(6)������CA������֤��ͬ�ͻ���һ�����������ڿͻ��˶Է���˵�����֤�����У�顣 
����(7)�Ͳ���(8)��������˵�����֤���˽Կͬ�����һ��������