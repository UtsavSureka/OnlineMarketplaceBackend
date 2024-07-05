package config

var (
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
	JWT_SECRET  string
)

func InitConfig() {
	DB_USER = "root"
	DB_PASSWORD = "root"
	DB_NAME = "Marketplace"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	JWT_SECRET = "HelloWorld"

}
