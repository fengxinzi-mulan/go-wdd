package conf

type Server struct {
	Vfs Vfs `ini:"vfs" mapstructure:"vfs" json:"vfs" yaml:"vfs"`
}
