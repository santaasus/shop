package smtp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shop/notification_service/inner_layer/domain"
	"strings"
	"time"
)

func StartListenServ() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config *domain.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.SMTP.Source)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	reader := bufio.NewReader(conn)

	conn.Write([]byte("220 Mock SMTP server ready\r\n"))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}

		line = strings.TrimSpace(line)
		fmt.Println("Received:", line)

		if line == "QUIT" {
			conn.Write([]byte("221 Bye\r\n"))
			break
		}

		if isHasPrefix(line, "EHLO") || isHasPrefix(line, "HELO") {
			conn.Write([]byte("250-Hello\r\n250 OK\r\n"))
		} else if isHasPrefix(line, "RCPT TO") || isHasPrefix(line, "MAIL FROM") {
			conn.Write([]byte("250 OK\r\n"))
		} else if isHasPrefix(line, "DATA") {
			conn.Write([]byte("354 End data with <CR><LF>.<CR><LF>\r\n"))
		} else if line == "." {
			conn.Write([]byte("250 OK\r\n"))
		}
	}
}

func isHasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
