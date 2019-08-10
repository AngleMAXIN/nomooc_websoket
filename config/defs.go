package config

// MysqlConfig mysql的相关配置
type MysqlConfig struct {
	Port   int
	Host   string
	User   string
	Passwd string
	DB     string
	Loc    string
}

// RedisConfig redis的相关配置
type RedisConfig struct {
	DialTime  int
	ReadTime  int
	WriteTime int
	PoolSize  int
	Port      int
	DB        int
	Host      string
}

// ServerConfig 监听服务的相关配置
type ServerConfig struct {
	Host             string
	MaxConnLimit     int
	CurrentConnLimit int
}

// Config 全局总配置
type Config struct {
	Mysql  MysqlConfig
	Redis  RedisConfig
	Server ServerConfig
}

// GetMysqlConfig 获得数据库的配置
func (c Config) GetMysqlConfig() MysqlConfig {
	return c.Mysql
}

// GetRedisConfig 获得Redis的配置
func (c Config) GetRedisConfig() RedisConfig {
	return c.Redis
}

// GetServerConfig 获得服务的配置
func (c Config) GetServerConfig() ServerConfig {
	return c.Server
}
