package configs

import (
	"os"
)

func LoadConfigEnv() {
	//DB Env Variable
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_USER", "root")
	os.Setenv("DATABASE_PORT", "8001")
	os.Setenv("DATABASE_PASSWORD", "1")
	os.Setenv("DATABASE_NAME", "root")

	//JWT Env Variable
	os.Setenv("SECRET_KEY", "ultra_secret_key")
	os.Setenv("JWT_LIFE", "24")
}
