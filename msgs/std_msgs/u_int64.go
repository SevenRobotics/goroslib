package std_msgs

import (
	"github.com/aler9/goroslib/msg"
)

type UInt64 struct {
	msg.Package `ros:"std_msgs"`
	Data        uint64
}
