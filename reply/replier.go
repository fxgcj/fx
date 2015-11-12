package reply

import (
	"github.com/fxgcj/fx/lib"
)

type Replier interface {
	Reply(...interface{}) string
}

var (
	log           = lib.GetLogger()
	tulingReplier *TulingReply
)

func init() {
	if tulingReplier == nil {
		tulingReplier = &TulingReply{}
	}
}

func GetReplier() Replier {
	return tulingReplier
}
