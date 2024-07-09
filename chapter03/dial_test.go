package chapter03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// 1. 서버 역할

	done := make(chan struct{}) // 채널을 생성하여 고루틴의 완료를 통보한다.
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			conn, err := listener.Accept() // 새로운 TCP 연결을 수락한다.
			if err != nil {
				t.Log(err)
				return
			}

			go func(c net.Conn) { // 또 다른 고루틴이 생성된 연결에서 데이터를 읽는다.
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf) // 연결에서 데이터를 읽어 버퍼에 저장한다.
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n]) // 읽은 데이터를 로그에 출력한다.
				}
			}(conn)
		}
	}()

	// 2. 클라이언트 역할

	conn, err := net.Dial("tcp", listener.Addr().String()) // 리스너의 주소로 TCP 연결을 연다.
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()     // 클라이언트 연결을 닫는다.
	<-done           // Waits for the outer server goroutine to signal it's done.
	listener.Close() // 리스너를 닫는다.
	<-done           // Waits for the inner connection handling goroutine to signal it's done.
}

/*

Sequence of Events

1. Server Goroutine:
- The server goroutine starts and listens for connections.
- When a connection is accepted, a new goroutine is spawned to handle that connection.
- The server goroutine waits for the next connection or until it encounters an error.

2. Connection Handling Goroutine:
- Handles the connection by reading data.
- When the connection is closed or an error occurs, it signals completion by sending a value to 'done'.

*/
