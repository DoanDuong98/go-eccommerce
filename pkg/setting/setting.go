package setting

type Config struct {
	Mysql MysqlSetting `mapstructure:"mysql"`
}

type MysqlSetting struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"dbname"`
	MaxIdle        int    `mapstructure:"maxIdleConns"`
	MaxOpen        int    `mapstructure:"maxOpenConns"`
	ConMaxLifetime int    `mapstructure:"maxLifetimeConns"`
}
