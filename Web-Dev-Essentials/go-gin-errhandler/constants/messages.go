package constants

var (
	UnauthorizedRequest    = "MSG_UNAUTH"
	ConfigInitFailed       = "CONFIG_INIT_FAILED"
	GrpcServerExitSuccess  = "GRPC_SERVER_EXITED_SUCCESSFULLY"
	ListenFail             = "FAILED_TO_LISTEN"
	ServeFail              = "FAILED_TO_SERVE"
	LogInitFailed          = "LOG_INIT_FAILED"
	InternalServerError    = "MSG_INTERNAL_SERVER_ERROR"
	InvalidToken           = "MSG_INVALID_TOKEN"
	InvalidJson            = "MSG_INVALID_JSON"
	DBError                = "MSG_DB_ERROR"
	ServiceUnavaibleError  = "MSG_SERVICE_UNAVAILABLE"
	UnknownError           = "MSG_UNKNOWN_ERROR_OCCURED"
	InvalidUsernamePasword = "MSG_INVALID_USERNNAME_PASSWORD"
	InvalidArgument        = "MSG_INVALID_ARGUMENT"
	Unimplemented          = "MSG_INVALID_UNIMPLEMENTED"
	InvalidApiPath         = "MSG_INVALID_API_PATH"
	MetadataRetrivalFail   = "MSG_FAILED_RETRIVAL_METADATA"
)

var (
	UniqueViolation     = "MSG_UNIQUE_KEY_VIOLATION"
	ForeignKeyViolation = "MSG_FOREIGN_KEY_VIOLATION"
	CheckViolation      = "MSG_CHECK_KEY_VIOLATION"
)

const (
	Bearer         = "Bearer "
	Authorization  = "authorization"
	AuthDetails    = "authdetails"
	APIPath        = "apipath"
	APITokenPrefix = "cb-auth-jwt-token-api"
)
