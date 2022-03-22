package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var Service *impl = &impl{}

type impl struct {
	// 可以更换自己熟悉的, logrus,标准库log,zap
	// mcube log模块是包装 zap
	log logger.Logger
	// 依赖数据库
	db *sql.DB
}

func (i *impl) Init() error {
	i.log = zap.L().Named("Host")
	return nil
}
