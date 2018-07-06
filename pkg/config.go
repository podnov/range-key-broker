package pkg

import (
	"github.com/spf13/viper"
)

const (
	ConfigKeyRangeValueRange = "range"
	ConfigKeyRedisAddress    = "redis_address"
	ConfigKeyRedisPassword   = "redis_password"
)

func init() {
	viper.SetEnvPrefix("RANGE_VALUE_BROKER")
	viper.AutomaticEnv()
}
