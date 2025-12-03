package conf

type Vfs struct {
	Enabled               bool     `ini:"enabled" mapstructure:"enabled" json:"enabled" yaml:"enabled"`                                                                 // 是否启用同步
	AccessToken           string   `ini:"access_token" mapstructure:"access_token" json:"access_token" yaml:"access_token"`                                             // 访问令牌
	RefreshToken          string   `ini:"refresh_token" mapstructure:"refresh_token" json:"refresh_token" yaml:"refresh_token"`                                         // 刷新令牌
	MountDirID            string   `ini:"mount_dir_id" mapstructure:"mount_dir_id" json:"mount_dir_id" yaml:"mount_dir_id"`                                             // 挂载目录id
	MountDirPath          string   `ini:"mount_dir_path" mapstructure:"mount_dir_path" json:"mount_dir_path" yaml:"mount_dir_path"`                                     // 挂载目录路径
	ForwardFileExt        []string `ini:"forward_file_ext,omitempty" mapstructure:"forward_file_ext" json:"forward_file_ext" yaml:"forward_file_ext"`                   // 转发文件后缀列表
	CacheTime             int      `ini:"cache_time" mapstructure:"cache_time" json:"cache_time" yaml:"cache_time"`                                                     // 缓存时间(秒)
	SyncFrequency         int      `ini:"sync_frequency" mapstructure:"sync_frequency" json:"sync_frequency" yaml:"sync_frequency"`                                     // 同步频率(秒)
	FullSyncFrequency     int      `ini:"full_sync_frequency" mapstructure:"full_sync_frequency" json:"full_sync_frequency" yaml:"full_sync_frequency"`                 // 全量同步频率(秒)
	RefreshTokenFrequency int      `ini:"refresh_token_frequency" mapstructure:"refresh_token_frequency" json:"refresh_token_frequency" yaml:"refresh_token_frequency"` // 刷新token频率(秒)
}
