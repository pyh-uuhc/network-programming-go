package chapter03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0") // TCP 네트워크를 사용하는 서버 리스너를 생성, '127.0.0.1:0'은 로컬 호스트의 임의의 포트에 바인딩하도록 지정
	if err != nil {
		t.Fatal(err) // 't.Fatal' 함수는 테스트에서 중요한 에러가 발생했음을 나타내고, 테스트를 즉시 중지시키는 함수
	}
	defer func() { _ = listener.Close() }() // 'defer'를 사용하여 테스트 함수가 반환하기 직전에 'listener.Close()'를 호출하여 리소스 누수를 방지

	t.Logf("bound to %q", listener.Addr()) // 실제로 바인딩된 네트워크 주소를 가져온 후, 이를 포맷된 문자열로 출력
	// 't.Logf' 함수는 테스트 함수에서 포맷된 로그를 출력하는 함수
}
