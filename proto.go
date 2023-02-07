package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

const Head = "####"

const _headLength = 8 // 数据头长度 Head的长度加数据的长度数据 等等

func IntToBytes(i int) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	//var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入包头
	err := binary.Write(pkg, binary.LittleEndian, []byte(Head))
	if err != nil {
		return nil, err
	}

	// 写入消息长度
	err = binary.Write(pkg, binary.LittleEndian, IntToBytes(len(message)))
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) ([]byte, error) {
	length, err := getDataLength(reader)
	if err != nil {
		return nil, err
	}

	//
	return decode(reader, length)
}

func decode(reader *bufio.Reader, length int) ([]byte, error) {

	var (
		body = make([]byte, length)
		err  error
	)

	// 判断是否大于缓冲区的数据
	if length <= reader.Size() {
		if _, err = reader.Read(body); err != nil {
			return nil, err
		}
	} else {
		var (
			count      int // 已读取长度
			residue    int // 未读取长度
			peekLength = reader.Size()
		)

		for {
			if n, err := reader.Read(body[count : count+peekLength]); err != nil {
				return nil, err
			} else {
				count += n
			}

			// 判断是否读取完成
			if residue = length - count; residue > 0 {
				if residue < reader.Size() {
					peekLength = residue
				}
				if _, err = reader.Peek(peekLength); err != nil {
					return nil, err
				}
			} else {
				break // 读取完成退出
			}
		}
	}
	return body[_headLength:], nil
}

func getDataLength(reader *bufio.Reader) (int, error) {
	// 获取数据长度
	peek, err := reader.Peek(_headLength)
	if err != nil {
		return 0, err
	}

	var length int32

	// 默认大端
	if err := binary.Read(bytes.NewBuffer(peek[_headLength-4:]), binary.BigEndian, &length); err != nil {
		return 0, err
	}

	// 小端
	if length < 0 {
		return int(binary.LittleEndian.Uint32(peek[_headLength-4:]) + 8), nil
	}

	return int(length + _headLength), nil
}
