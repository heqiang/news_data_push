package config

type ServerConfig struct {
	Name                string `mapstructure:"name" json:"name,omitempty"`
	*RedisConfig        `mapstructure:"redis" json:"*redis,omitempty"`
	*LogConfig          `mapstructure:"log" json:"*_log_config,omitempty"`
	PushHour            string `json:"push_hour"`             // 推送哪个小时的数据 ex 01 02 03 .....
	PushDay             string `json:"push_day"`              // 推送哪天的数据 ex:2023-02-03
	DayHours            string `json:"day_hours"`             // ex: 2023020301
	SourceNewsPath      string `json:"source_news_path"`      //原始新闻数据存放路径
	VideoPath           string `json:"video_path"`            //视频保存路径
	IsSupplementaryPush bool   `json:"is_supplementary_push"` //是否补推数据
	JsonSavePath        string `json:"json_save_path"`        //json保存路径
	FtpDataPath         string `json:"ftp_data_path"`         //ftp路径
	ZipPath             string `json:"zip_path"`              //压缩包保存路径
	DataPath            string `json:"data_path"`             //中间数据保存路径
}
type RedisConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	Db       int    `json:"db" mapstructure:"db"`
	SetKey   string `json:"setKey"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
