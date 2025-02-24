package setting

type Config struct {
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Setting ServerSetting `mapstructure:"server"`
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

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type LoggerSetting struct {
	LogLevel   string `mapstructure:"log_level"`
	LogFile    string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type ServerSetting struct {
	Port		string `mapstructure:"port"`
	Mode 	   	string `mapstructure:"mode"`
}
