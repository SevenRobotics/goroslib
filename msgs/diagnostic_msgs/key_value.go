package diagnostic_msgs

import (
	"github.com/aler9/goroslib/msg"
)

type KeyValue struct {
	msg.Package `ros:"diagnostic_msgs"`
	Key         string
	Value       string
}
