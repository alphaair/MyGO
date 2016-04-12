package common

import "fmt"
import "bufio"
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

// Append 将src的切片追回到dst中
func Append(dst, src []interface) {
    for _, e := range src {
        append(dst, e)
    }
}