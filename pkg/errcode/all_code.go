package errcode

var (
	InvalidParams = NewApiError(400, "invalid params")
	NotFound      = NewApiError(404, "not found")
	InternalError = NewApiError(500, "internal error")

	UserNotExists     = NewApiError(4001, "user not exists")
	UserAlreadyExists = NewApiError(4002, "user already exists")
	UserNotActive     = NewApiError(4003, "user not active")
	PasswordWrong     = NewApiError(4005, "password wrong")

	RoleNotExists     = NewApiError(4011, "role not exists")
	RoleAlreadyExists = NewApiError(4012, "role already exists")
	HaveRelativeUser  = NewApiError(4012, "role have relative user")

	TokenGenerate = NewApiError(4101, "generate token failed")
	TokenInvalid  = NewApiError(4102, "token invalid")

	PermissionDeny = NewApiError(4300, "permission deny")

	RateLimitExceeded = NewApiError(4400, "rate limit exceeded")

	UploadFailed          = NewApiError(4500, "file upload failed")
	UploadExtNotSupported = NewApiError(4501, "file suffix is not supported")
	UploadExcessMaxSize   = NewApiError(4502, "file size exceeded maximum limit")
)
