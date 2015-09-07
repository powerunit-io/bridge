package main

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections/mysql"
	"github.com/powerunit-io/platform/logging"
)

// NewDb -
func NewDb(n string, conf map[string]interface{}, logger *logging.Logger) (mysql.Manager, error) {
	cnf, err := config.NewConfigManager(n, conf)

	if err != nil {
		logger.Error("Failed to configure mysql configuration manager for (manager: %s) (error: %s)", n, err)
		return nil, err
	}

	cnf.Set("name", n)

	mysqlconn, err := mysql.NewConnection(logger, cnf)

	if err != nil {
		logger.Error("Failed to initiate mysql connection for (manager: %s) (error: %s)", n, err)
		return nil, err
	}

	return mysqlconn, nil
}
