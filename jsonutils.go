package jsonutils

type JSONResponse map[string]interface{}

func SuccessResponse(objectName string, object interface{}) JSONResponse {
	return JSONResponse{
		objectName: object,
	}
}
