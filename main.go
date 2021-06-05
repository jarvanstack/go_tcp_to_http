package main
import "net"
const address = ":9000"
func main() {
	listen, _ := net.Listen("tcp", address)
	for  {
		conn, _ := listen.Accept()
		defer conn.Close()
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		conn.Write([]byte("Content-Length: 2\r\n"))
		conn.Write([]byte("Content-Type:text/html:charset=UTF-8\r\n\r\n"))
		conn.Write([]byte("OK"))
	}

}
