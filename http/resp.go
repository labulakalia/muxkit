package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muxkit/muxkit/proto/status"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Error struct {
	HTTPCode int32       // http code http的状态码 默认为500
	Code     status.Code // 业务返回码
	Message  interface{} // 业务错误信息的详情字段
}

func (e Error) Error() string {
	return fmt.Sprintf("Error Code: %d Message: %v", e.Code, e.Message)
}

// http error
func Resp(w http.ResponseWriter, resp interface{}, err error) {
	var httpErr Error
	switch err := err.(type) {
	case Error:
		httpErr = err
		if httpErr.HTTPCode != 0 {
			httpErr.HTTPCode = http.StatusOK
		}
		if httpErr.Code != 0 {
			httpErr.Code = status.Code_UNKNOWN
		}
		if httpErr.Message == nil {
			httpErr.Message = fmt.Sprint(httpErr.Code)
		}
	case nil:
		httpErr = Error{
			HTTPCode: 200,
			Code:     status.Code_OK,
		}
	default:
		httpErr = Error{
			HTTPCode: http.StatusInternalServerError,
			Code:     status.Code_UNKNOWN,
			Message:  err.Error(),
		}
	}

	if resp != nil {
		data, err := Marshal(resp.(protoreflect.ProtoMessage))
		if err != nil {
			log.Println("marshal message failed", err)
		}
		_, _ = w.Write(data)
	}
	w.WriteHeader(int(httpErr.HTTPCode))
	w.Header().Add("X-Code", fmt.Sprint(httpErr.Code))
	w.Header().Add("X-Message", fmt.Sprint(httpErr.Message))
}

func InvalidArgument(message ...interface{}) error {
	err := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     status.Code_INVALID_ARGUMENT,
	}
	if len(message) > 0 {
		err.Message = message[0]
	}
	return err
}
