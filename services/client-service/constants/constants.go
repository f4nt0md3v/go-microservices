package constants

const TTL = 60

type InternalMessage struct {
	IsSuccess bool
	ErrorCode int32
	Details   map[string]interface{}
}
