package env

type env struct {
	DBUser     string `env:"POSTGRES_USER,required"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName     string `env:"POSTGRES_DB,required"`
	DBHost     string `env:"DB_HOST,required"`
}

var Env = new(env)
