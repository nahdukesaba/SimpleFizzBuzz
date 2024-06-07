package helpers

var (
	ErrorMessageBadRequest = "Bad Request"

	BadRequesetMessage  = map[string]interface{}{"error": "Bad Request"}
	SuccessMessage      = map[string]interface{}{"success": true}
	UnauthorizedMessage = map[string]interface{}{"error": "Unauthorized! Please login first"}
)
