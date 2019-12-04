package config

type DbConfig struct {
	Host         string `toml:"host"`
	Port         uint   `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Charset      string `toml:"charset"`
	Database     string `toml:"database"`
	Timeout      int    `toml:"timeout" json:"timeout"`
	MaxOpenConns int    `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns" json:"max_idle_conns"`
	MaxConnTtl   int    `toml:"max_conn_ttl" json:"max_conn_ttl"`
	Debug        bool   `toml:"debug"`
}
