package jsonutil

type Response map[string]interface{}

func SuccessResponse(objectName string, object interface{}) Response {
	return Response{
		objectName: object,
	}
}
