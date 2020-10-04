package std_msgs

import (
	"github.com/aler9/goroslib/msg"
)

type UInt8MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []uint8
}
