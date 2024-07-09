package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntoBuffer(t *testing.T) {
	payload := make([]byte, 1<<24) // 길이가 16MB인 바이트 슬라이스를 생성하여 사용할 수 있도록 준비
	_, err := rand.Read(payload)   // Read 메서드는 주어진 바이트 슬라이스(payload 변수)에 랜덤한 데이터를 채운다.
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:") // TCP 프로토콜을 사용하는 서버용 리스너를 생성
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		// listener 변수가 나타내는 TCP 리스너에서 클라이언트의 연결 요청을 기다린다.
		// 연결 요청이 들어오면 이를 수락하고, 새로운 net.Conn 인터페이스를 반환한다.
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload) // TCP 연결을 통해 데이터를 전송
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String()) // TCP 연결을 통해 클라이언트가 서버에 연결하는 작업을 수행
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1<<19)

	// TCP 연결로부터 데이터를 읽고, 읽은 데이터를 처리하며, 연결을 닫는 과정
	for {
		n, err := conn.Read(buf) // conn 객체로부터 데이터를 읽고, 이를 buf에 저장
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break // err가 io.EOF인 경우, 이는 연결의 끝(End Of File)을 의미하며, 더 이상 읽을 데이터가 없다는 뜻
		}

		t.Logf("read %d bytes", n) // 데이터를 성공적으로 읽었을 때, 읽은 바이트 수를 로그로 남긴다.
	}

	conn.Close()
}
