package fnet

import (
	"github.com/bwmarrin/snowflake"
	"github.com/futurez/freedom/fconf"
	"github.com/futurez/freedom/flog"
	"sync"
)

const (
	INVALID_CONN_ID = 0 //非法conn_id
)

var connIdOnce sync.Once
var connIdNode *snowflake.Node

func generateConnId() int64 {
	connIdOnce.Do(func() {
		var err error
		connIdNode, err = snowflake.NewNode(int64(fconf.Conf.ServerType))
		if err != nil {
			flog.Panic("new connIdNode err = ", err.Error())
		}
	})
	return connIdNode.Generate().Int64()
}
