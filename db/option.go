package db

type Option func(*Options)

//默认参数
const (
	DF_Host            string = "127.0.0.1"
	DF_DbName          string = "payv1"
	DF_Port            int    = 3306
	DF_User            string = "root"
	DF_Password        string = "123456"
	DF_MaxIdleConn     int    = 1000
	DF_MaxOpenConn     int    = 1024
	DF_ConnMaxLifeTime int    = 3600
	DF_Charset         string = "utf8mb4"
	DF_IsLog           bool   = true
)

type Options struct {
	Host            string `mapstructure:"host"`
	DbName          string `mapstructure:"db_name"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
	Charset         string `mapstructure:"charset"`
	IsLog           bool   `mapstructure:"is_log"`
}

func WithHost(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func WithDbName(dbName string) Option {
	return func(o *Options) {
		o.DbName = dbName
	}
}

func WithPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func WithUser(user string) Option {
	return func(o *Options) {
		o.User = user
	}
}

func WithPassword(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(o *Options) {
		o.MaxIdleConn = maxIdleConn
	}
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(o *Options) {
		o.MaxOpenConn = maxOpenConn
	}
}

func WithMaxLifeTime(connMaxLifeTime int) Option {
	return func(o *Options) {
		o.ConnMaxLifeTime = connMaxLifeTime
	}
}

func WithCharset(charset string) Option {
	return func(o *Options) {
		o.Charset = charset
	}
}

func WithIsLog(isLog bool) Option {
	return func(o *Options) {
		o.IsLog = isLog
	}
}
