package zpack

import (
	"encoding/json"
	"github.com/llyRayYoung/zinx/ziface"
)

type DataPackJson struct{}

type JsonMessage struct {
	MsgID uint32 `json:"id"`
	Data  string `json:"data"`
}

func NewDataPackJson() ziface.IDataPack {
	return &DataPackJson{}
}

func (dp *DataPackJson) GetHeadLen() uint32 {
	return 0
}

func (dp *DataPackJson) Pack(msg ziface.IMessage) ([]byte, error) {
	jsonMsg := &JsonMessage{
		MsgID: msg.GetMsgID(),
		Data:  string(msg.GetData()),
	}
	data, err := json.Marshal(jsonMsg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dp *DataPackJson) Unpack(binaryData []byte) (ziface.IMessage, error) {
	jsonMsg := &JsonMessage{}
	err := json.Unmarshal(binaryData, jsonMsg)
	if err != nil {
		return nil, err
	}
	return NewMsgPackage(jsonMsg.MsgID, []byte(jsonMsg.Data)), nil
}
