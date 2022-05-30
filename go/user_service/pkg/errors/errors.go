package errors

/*
{
	"errorName": "Validation error"
	"errors": [
		{
	  		"field": "username",
	  		"userMessage": "username is already taken",
	  		"internalMessage": "repository.go: Register: username violates unique constraint"
		},
	]
}
*/

type Response struct {
	ErrorName string  `json:"errorName"`
	Errors    []Error `json:"errors"`
}

type Error struct {
	Field           string `json:"field"`
	UserMessage     string `json:"userMessage"`
	InternalMessage string `json:"internalMessage"`
}

// create http error response based on the type of error
func CreateResponse(e error) (int, Response) {
	var resp Response
	var err Error
	var statusCode int

	switch e.(type) {
	case *NotFound:
		statusCode = 404
		err.Field = "id"
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "authentication error"
		resp.Errors = append(resp.Errors, err)
	case *InvalidCredentials:
		statusCode = 409
		err.Field = "email"
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "validation error"
		resp.Errors = append(resp.Errors, err)
	case *RabbitError:
		statusCode = 500
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "rabbitmq internal error"
		resp.Errors = append(resp.Errors, err)
	case *InvalidUpdate:
		statusCode = 400
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "validation error"
		resp.Errors = append(resp.Errors, err)
	case *FileError:
		statusCode = 400
		err.Field = "avatar"
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "validation error"
		resp.Errors = append(resp.Errors, err)
	case *JwtAuthorization:
		statusCode = 403
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "authorization error"
		resp.Errors = append(resp.Errors, err)
	case *JwtBadRequest:
		statusCode = 400
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "authorization error"
		resp.Errors = append(resp.Errors, err)
	case *MethodNotAllowed:
		statusCode = 405
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "request error"
		resp.Errors = append(resp.Errors, err)
	default:
		statusCode = 500
		err.Field = ""
		err.UserMessage = e.Error()
		err.InternalMessage = e.Error()
		resp.ErrorName = "unkown"
		resp.Errors = append(resp.Errors, err)
	}
	return statusCode, resp
}
