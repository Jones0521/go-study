package conf

// conf pkg 的全局对象
// global 全局配置对象
var global *Config

func c() *Config {
	if global == nil {
		panic("config required")
	}
	return global
}

func SetGlobalConfig(conf *Config) {
	global = conf
}

// 配置通过对象来进行映射
// 我们定义是, 配置对象的数据结构

type Config struct {
	App   *app
	Mysql *mysql
	Log   *log
}

type app struct {
	// restful-api
	Name string
	// 127.0.0.1, 0.0.0.0
	Host string `toml:"host"`
	// 8080
	Port string `toml:"port"`
	// 比较敏感的数据,入库的是加密的数据，加密的秘钥就是该配置
	Key string `toml:"key"`
}

// mysql 数据库配置
type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}
type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}
