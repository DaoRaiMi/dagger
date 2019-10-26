package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"
)

const (
	EnvPrefix = "DAGGER_"
)

func init() {
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetBool(key string) bool {
	var valueBool bool
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetBool(key)
	}
	switch value {
	case "True":
		valueBool = true
	case "true":
		valueBool = true
	}
	return valueBool
}

func GetString(key string) string {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return strings.TrimSpace(viper.GetString(key))
	}
	return strings.TrimSpace(value)
}

func GetInt(key string) int {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetInt(key)
	}

	d, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0
	}
	return d
}

func GetUint32(key string) uint32 {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetUint32(key)
	}

	d, err := strconv.ParseUint(strings.TrimSpace(value), 10, 32)
	if err != nil {
		return 0
	}
	return uint32(d)
}

func GetUint64(key string) uint64 {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetUint64(key)
	}

	d, err := strconv.ParseUint(strings.TrimSpace(value), 10, 64)
	if err != nil {
		return 0
	}
	return d
}

func GetFloat32(key string) float64 {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetFloat64(key)
	}

	d, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0
	}
	return d
}

func GetTime(key string) time.Time {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetTime(key)
	}

	t, err := time.ParseInLocation(time.RFC3339, strings.TrimSpace(value), time.Local)
	if err != nil {
		return time.Time{}
	}
	return t
}

func GetDuration(key string) time.Duration {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetDuration(key)
	}

	d, err := time.ParseDuration(strings.TrimSpace(value))
	if err != nil {
		return time.Duration(0)
	}
	return d
}

func GetStringSlice(key, sep string) []string {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		return viper.GetStringSlice(key)
	}
	if sep == "" {
		sep = ","
	}
	valueSlice := strings.Split(strings.TrimSpace(value), sep)
	return valueSlice
}

func GetLogLevel(key string) zapcore.LevelEnabler {
	envKey := EnvPrefix + strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	value, exists := os.LookupEnv(envKey)
	if !exists {
		value = viper.GetString(key)
	}

	switch value {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
