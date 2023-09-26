package config

type Config struct {
	GrpcMysqlAddr string `mapstructure:"GRPC_MYSQL_ADDR"`
	GrpcMysqlPort string `mapstructure:"GRPC_MYSQL_PORT"`
	ServerAddr    string `mapstructure:"SERVER_ADDR"`
	ReviewAddr    string `mapstructure:"REVIEW_ADDR"`
	JaegerAddr    string `mapstructure:"JAEGER_ADDR"`
}
