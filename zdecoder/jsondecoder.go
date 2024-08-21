package zdecoder

import (
	"encoding/json"
	"github.com/llyRayYoung/zinx/ziface"
	"github.com/llyRayYoung/zinx/zpack"
)

type JsonDecoder struct {
	Tag   uint32 //T
	Value []byte //V
}

func NewJsonDecoder() ziface.IDecoder {
	return &JsonDecoder{}
}

func (j *JsonDecoder) GetLengthField() *ziface.LengthField {
	return nil
}

func (j *JsonDecoder) decode(data []byte) (*JsonDecoder, error) {
	jsonMsg := &zpack.JsonMessage{}
	err := json.Unmarshal(data, jsonMsg)
	if err != nil {
		return nil, err
	}
	return &JsonDecoder{
		Tag:   jsonMsg.MsgID,
		Value: []byte(jsonMsg.Data),
	}, nil
}
func (j *JsonDecoder) Intercept(chain ziface.IChain) ziface.IcResp {
	iMessage := chain.GetIMessage()
	if iMessage == nil {
		// Go to the next layer in the chain of responsibility
		return chain.ProceedWithIMessage(iMessage, nil)
	}
	data := iMessage.GetData()

	jsonData, err := j.decode(data)
	if err != nil {
		return chain.ProceedWithIMessage(iMessage, nil)
	}

	iMessage.SetMsgID(jsonData.Tag)
	iMessage.SetData(jsonData.Value)
	iMessage.SetDataLen(0)

	//6. Pass the decoded data to the next layer.
	// (将解码后的数据进入下一层)
	return chain.ProceedWithIMessage(iMessage, *jsonData)
}
