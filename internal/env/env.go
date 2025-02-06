package env

import (
	"os"
	"reflect"

	"github.com/freitzzz/gameboy-db-api/internal/logging"
	"github.com/joho/godotenv"
)

type env struct {
	ServerHost        *string `env:"server_host"`
	ServerPort        *string `env:"server_port"`
	ServerTLSCert     *string `env:"server_tls_crt_fp"`
	ServerTLSKey      *string `env:"server_tls_key_fp"`
	ServerVirtualPath *string `env:"server_virtual_path"`
	DBPath            *string `env:"db_path"`
}

func loadEnv() env {
	if err := godotenv.Load(); err == nil {
		logging.Warning("couldn't load .env file! (%v)", err)
	}

	env := env{}
	v := reflect.ValueOf(env)
	addr := reflect.ValueOf(&env).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		k := f.Tag.Get("env")
		if v, ok := os.LookupEnv(k); ok {
			addr.Field(i).Set(reflect.ValueOf(&v))
		}

	}

	return env
}

var Env env = loadEnv()
