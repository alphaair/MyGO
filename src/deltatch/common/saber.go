package common

import "fmt"
import "bufio"
import "unicode/utf8"
import "os"

// ReadKey 从当前标准输入流中读取一个字节
func ReadKey(msg string) byte {

	if len(msg) > 0 {
		fmt.Printf(msg)
	}

	render := bufio.NewReader(os.Stdin)
	key, _ := render.ReadByte()

	return key
}

// ReadLine 从当前OS标准流中读取一行字符
func ReadLine(msg string) string {
	if len(msg) > 0 {
		fmt.Printf(msg)
	}

	render := bufio.NewReader(os.Stdin)
	buffer, _, _ := render.ReadLine()

	return string(buffer)
}

// AppendForByte 将src的切片追回到dst中
func AppendForByte(dst *[]byte, src []byte) {
	for _, e := range src {
		*dst = append(*dst, e)
	}
}

// TextToUtf8Bytes 将text转换成Utf8的字节形式,本方法可以通过硬转转实现，仅为练习使用
func TextToUtf8Bytes(text string) []byte {
	//==[]byte(text)
	var buffer []byte
	for len(text) > 0 {
		ru, size := utf8.DecodeRuneInString(text)
		text = text[size:]

		bf := make([]byte, 3)
		size = utf8.EncodeRune(bf, ru)
		AppendForByte(&buffer, bf[:size])
	}

	return buffer
}

// BytesToUft8Text 将UTF8编码的字节数组还原成相应的字符串
func BytesToUft8Text(buffer []byte) string {
	var tmp []rune

	for len(buffer) > 0 {
		r, size := utf8.DecodeRune(buffer)
		tmp = append(tmp, r)
		buffer = buffer[size:]
	}

	return string(tmp)
}
