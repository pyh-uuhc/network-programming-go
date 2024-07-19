package main

import (
	"errors"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1 // 1
	StringType                  // 2

	MaxPayloadSize uint32 = 10 << 20 // 10 메가바이트
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

// Payload 인터페이스는 페이로드의 동작을 정의하며, fmt.Stringer, io.ReaderFrom, io.WriterTo 인터페이스를 포함
type Payload interface {
	fmt.Stringer   // fmt.Stringer 인터페이스를 포함, 이는 String() 메서드를 요구
	io.ReaderFrom  // io.ReaderFrom 인터페이스를 포함, 이는 ReadFrom(r io.Reader) (n int64, err error) 메서드를 요구
	io.WriterTo    // io.WriterTo 인터페이스를 포함, 이는 WriteTo(w io.Writer) (n int64, err error) 메서드를 요구
	Bytes() []byte // Payload 인터페이스는 Bytes() 메서드를 추가로 요구, 이는 페이로드를 바이트 슬라이스로 반환
}
