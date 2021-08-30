package configuration

import (
	"github.com/haproxytech/client-native/v2/models"

	"github.com/haproxytech/kubernetes-ingress/controller/utils"
)

// SetGlobal will set default values for Global config.
func SetGlobal(global *models.Global, env Env) {
	// Enforced values
	global.MasterWorker = true
	global.Pidfile = env.PIDFile
	global.Localpeer = "local"
	global.ServerStateBase = env.StateDir
	global.ServerStateFile = "global"
	global.RuntimeAPIs = append(global.RuntimeAPIs, &models.RuntimeAPI{
		Address:           &env.RuntimeSocket,
		ExposeFdListeners: true,
		Level:             "admin",
	})
	// Default values
	if global.Daemon == "" {
		global.Daemon = "enabled"
	}
	if global.StatsTimeout == nil {
		global.StatsTimeout = utils.PtrInt64(36000)
	}
	if global.TuneSslDefaultDhParam == 0 {
		global.TuneSslDefaultDhParam = 2048
	}
	if global.SslDefaultBindCiphers == "" {
		global.SslDefaultBindCiphers = "ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK"
	}
	if global.SslDefaultBindOptions == "" {
		global.SslDefaultBindOptions = "no-sslv3 no-tls-tickets no-tlsv10"
	}
}
