package znet

type Message struct {
	Id 	uint32 //消息的ID
	DataLen uint32 //消息的长度
	Data []byte //消息的内容
}


// GetMsgId 获取消息的ID
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// GetMsgLen 获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// GetMsgData 获取消息的内容
func (m *Message) GetMsgData() []byte {
	return m.Data
}

// SetMsgId 设置消息的ID
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

// SetMsgLen 设置消息的长度
func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

// SetDataLen 设置消息的内容
func (m *Message) SetDataLen(data []byte) {
	m.Data = data
}