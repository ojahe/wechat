package context

import (
	"github.com/ojahe/wechat/v2/credential"
	"github.com/ojahe/wechat/v2/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
