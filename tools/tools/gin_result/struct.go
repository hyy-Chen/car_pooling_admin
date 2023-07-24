/**
 @author : ikkk
 @date   : 2023/8/10
 @text   : nil
**/

package gin_result

import (
	"encoding/json"
)

type Response struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误描述
	Data    interface{} `json:"data"`    // 返回数据
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"message"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Message,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}
