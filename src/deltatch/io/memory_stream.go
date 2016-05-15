// Coyright alphaair 2016
// 这是一个简单的内存流

package io

import (
	"io"
)

type MemoryStream struct {
	//暂存数据流
	buffer   []byte
	Length   int
	Position int
}

// NewMemoryStream 以指定大小初始化一个内存数据流
func NewMemoryStream(size int) *MemoryStream {
	stream := new(MemoryStream)
	stream.buffer = make([]byte, size)

	return stream
}

// Read 读取内存数据流的数据到p
func (self *MemoryStream) Read(p []byte) (n int, err error) {

	if self.Length == 0 || self.Position+1 >= self.Length {
		return 0, io.EOF
	}

	cp := 0
	for i, _ := range p {
		self.Position++
		if self.Position >= self.Length {
			break
		}
		p[i] = self.buffer[self.Position]
		cp++
	}

	return cp, nil
}

// Write 将数据p写入到内存流中
func (self *MemoryStream) Write(p []byte) (n int, err error) {

	self.buffer = p
	self.Length = len(p)
	self.Position = -1

	return len(p), nil
}
