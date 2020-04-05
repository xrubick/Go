package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//封包和解包

//encode将消息编码

//Encode ...
func Encode(msg string) ([]byte, error) {
	//读取消息长度，转换成int32类型，占4个字节(4*8bit)
	length := int32(len(msg))
	var pkg = new(bytes.Buffer)

	//写入消息头信息
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	//写入消息实体信息
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

//decode将消息解码

//Decode ...
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息长度前4个字节
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//Buffered返回缓冲区中现存的可读字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	//读取消息

	pkg := make([]byte, int(4+length))
	_, err = reader.Read(pkg)
	if err != nil {
		return "", err
	}
	return string(pkg[4:]), nil

}
