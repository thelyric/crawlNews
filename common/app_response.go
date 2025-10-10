package common

import (
	"fmt"
	"time"
)

type basedResponse struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	TimeStamp int64  `json:"timestamp"`
}

type successRes struct {
	basedResponse
	Data   any `json:"data,omitempty"`
	Paging any `json:"paging,omitempty"`
	Filter any `json:"filter,omitempty"`
}
type FailRes struct {
	basedResponse
	Message any `json:"message,omitempty"`
}

func (f *FailRes) Error() string {
	if f == nil {
		return ""
	}
	return fmt.Sprintf("[%s] %v", f.Code, f.Message)
}

func NewSuccessResponse(data, paging, filter any) *successRes {
	return &successRes{
		basedResponse: basedResponse{
			Success:   true,
			Code:      "OK",
			TimeStamp: time.Now().Unix(),
		},
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data any) *successRes {
	return NewSuccessResponse(data, nil, nil)
}

func MakeFailResponse(code string, message any) *FailRes {
	var msg any
	if message == nil {
		msg = genericMessageForCode(code)
	} else {
		if err, ok := message.(error); ok && err != nil {
			// Hide internal error details from user, return friendly message
			msg = genericMessageForCode(code)
		} else {
			// If caller provided a friendly message (string/any), keep it
			msg = message
		}
	}

	return &FailRes{
		basedResponse: basedResponse{
			Success:   false,
			Code:      code,
			TimeStamp: time.Now().Unix(),
		},
		Message: msg,
	}
}

func genericMessageForCode(code string) any {
	switch code {
	case "TRANSPORT_ERROR":
		return "Có lỗi kết nối. Vui lòng thử lại sau."
	case "STORAGE_ERROR":
		return "Có lỗi xử lý dữ liệu. Vui lòng thử lại sau."
	case "BIZ_ERROR":
		return "Yêu cầu không thể thực hiện. Vui lòng kiểm tra và thử lại."
	case "INTERNAL_ERROR":
		return "Lỗi hệ thống. Vui lòng thử lại sau."
	case "RECORD_NOTFOUND_ERROR":
		return "Không tìm thấy dữ liệu."
	default:
		return "Đã xảy ra lỗi. Vui lòng thử lại sau."
	}
}

func NewTransportErrorResponse(err error) *FailRes {
	return MakeFailResponse("TRANSPORT_ERROR", err)
}

func NewStorageErrorResponse(err error) *FailRes {
	return MakeFailResponse("STORAGE_ERROR", err)
}

func NewBizErrorResponse(err error) *FailRes {
	return MakeFailResponse("BIZ_ERROR", err)
}

func NewInternalErrorResponse(err error) *FailRes {
	return MakeFailResponse("INTERNAL_ERROR", err)
}

func NewRecordNotFoundResponse(err error) *FailRes {
	return MakeFailResponse("RECORD_NOTFOUND_ERROR", err)
}
