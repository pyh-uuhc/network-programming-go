// 간단한 TCP 서버를 설정하고, 클라이언트가 연결되면 미리 정의된 메시지를 전송

package main

import (
	"net"
	"testing"
)

const payload = "The bigger the interface, the weaker the abstraction."

func TestScanner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:") // net.Listen 함수를 사용하여 TCP 리스너를 생성
	if err != nil {                                  // 리스너 생성에 실패하면, t.Fatal(err)을 호출하여 테스트를 실패로 처리하고 종료
		t.Fatal(err)
	}

	go func() { // 고루틴을 시작하여 비동기로 실행
		conn, err := listener.Accept() // listener.Accept를 호출하여 클라이언트 연결을 수락
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write([]byte(payload)) // 연결된 클라이언트에 payload 데이터를 전송, conn.Write는 전송된 바이트 수와 오류를 반환
		if err != nil {
			t.Error(err)
		}
	}()
}
