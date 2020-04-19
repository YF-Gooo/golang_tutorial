package protocol

import (
	"bytes"
	"encoding/binary"

	proto "github.com/golang/protobuf/proto"
)

const (
	ConstHeader         = "www.01happy.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

//封包
func Packet(message []byte) []byte {
	//利用proto把二进制打成压缩二进制
	body, _ := proto.Marshal(&Data{Value: message})
	return append(append([]byte(ConstHeader), IntToBytes(len(body))...), body...)
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstSaveDataLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
			//利用proto把二进制打成压缩二进制
			res := &Data{}
			proto.Unmarshal(data, res)
			readerChannel <- res.Value
			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
