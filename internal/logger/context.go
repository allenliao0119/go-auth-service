package logger

type ContextKey string

const (
	ContextKeyRequestID ContextKey = "request_id"
	ContextKeyLogger ContextKey = "logger"
)

var loggedContextKey = []ContextKey{
	ContextKeyRequestID,
}
