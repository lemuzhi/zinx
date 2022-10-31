package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/lemuzhi/zinx/utils"
	"github.com/lemuzhi/zinx/ziface"
)

type DataPack struct {}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包的头的长度方法
func (d *DataPack) GetHeadLen() uint32 {
	//Datalen uint32(4字节) + ID uint32(4字节)
	return 8
}

// Pack 封包方法
// |dataLen|msgID|data|
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将dataLen 写进dataBuff中
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}

	//将MsgId 写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}

	//将data数据 写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法 (将包的Head信息读出来) 之后再根据head信息里的data的长度，再进行一次读
func (d *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息，得到dataLen和MsgId
	msg := Message{}

	//读dataLen
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	//读MsgId
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}

	//判断dataLen是否已经超出了我们允许的最大包长度
	if(utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize) {
		return nil, errors.New("too Large msg data recv!")
	}

	return &msg, nil

}