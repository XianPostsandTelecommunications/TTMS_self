package config

import "time"

type AllConfig struct {
	Server    Server    `mapstructure:"Server"`
	App       App       `mapstructure:"App"`
	Log       Log       `mapstructure:"Log"`
	Mysql     Mysql     `mapstructure:"Mysql"`
	Redis     Redis     `mapstructure:"Redis"`
	SMTPInfo  SMTPInfo  `mapstructure:"SMTPInfo"`
	Rule      Rule      `mapstructure:"Rule"`
	Work      Work      `mapstructure:"Work"`
	Token     Token     `mapstructure:"Token"`
	AliyunOSS AliyunOSS `json:"aliyunOSS" mapstructure:"AliyunOSS"`
	Auto      Auto      `mapstructure:"Auto"`
}

// Server 用于存储服务器相关的配置信息
type Server struct {
	RunMode               string        `mapstructure:"RunMode"`               //运行模式
	Address               string        `mapstructure:"Address"`               //监听地址
	ReadTimeout           time.Duration `mapstructure:"ReadTimeout"`           //读取超时时间（接收请求）
	WriteTimeout          time.Duration `mapstructure:"WriteTimeout"`          //写入超时时间（发送响应）
	DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"` //存储默认上下文超时时间（处理请求最大时间，未完成则取消）
}

// APP 项目名称及版本
type App struct {
	Name    string `mapstructure:"Name"`
	Version string `mapstructure:"Version"`
}

// Log 日志库相关配置信息
type Log struct {
	Level         string `yaml:"Level"`         //日志级别
	LogSavePath   string `yaml:"LogSavePath"`   //  指定日志文件保存路径
	LowLevelFile  string `yaml:"LowLevelFile"`  //  低级别日志文件名
	LogFileExt    string `yaml:"LogFileExt"`    //	日志文件拓展名
	HighLevelFile string `yaml:"HighLevelFile"` //	高级别日志文件名
	MaxSize       int    `yaml:"MaxSize"`       // 单个文件最大大小（mb）
	MaxAge        int    `yaml:"MaxAge"`        // 日志文件最大存储时间
	MaxBackups    int    `yaml:"MaxBackups"`    // 保存的旧日志文件最大个数
	Compress      bool   `yaml:"Compress"`      //	指定是否压缩旧的日志文件
}

type Mysql struct {
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	DbName   string `mapstructure:"DbName"`
}

type Redis struct {
	Addr      string        `mapstructure:"Addr"`
	Password  string        `mapstructure:"Password"`
	PoolSize  int           `mapstructure:"PoolSize"`
	CacheTime time.Duration `mapstructure:"CacheTime"`
}

// SMTPInfo 简单邮件传输协议
type SMTPInfo struct {
	Host     string   `json:"host" mapstructure:"Host"`         //邮件服务器主机名
	Port     int      `json:"port" mapstructure:"Port"`         //邮件服务器端口号
	IsSSL    bool     `json:"isSSL" mapstructure:"IsSSL"`       //是否使用SSL加密
	UserName string   `json:"userName" mapstructure:"UserName"` //登录邮件服务器的用户名
	Password string   `json:"password" mapstructure:"Password"` //登录邮件服务器的密码
	From     string   `json:"from" mapstructure:"From"`         //发件人邮箱
	To       []string `json:"to" mapstructure:"To"`             //收件人邮箱
}

// Rule 用户配置
type Rule struct {
	DelUserTime          time.Duration `json:"delUserTime" mapstructure:"DelUserTime"`                   //删除用户的时间间隔
	DelCodeTime          time.Duration `json:"delCodeTime" mapstructure:"DelCodeTime"`                   //删除代码时间间隔
	DefaultAccountAvatar string        `json:"defaultAccountAvatar" mapstructure:"DefaultAccountAvatar"` //默认用户头像
	DefaultClientTimeout time.Duration `json:"defaultClientTimeout" mapstructure:"DefaultClientTimeout"` //默认客户端超时时间
	FileMaxSize          int64         `json:"fileMaxSize" mapstructure:"FileMaxSize"`                   //文件最大大小限制
	DefaultPagePerNum    int64         `json:"defaultPagePerNum" mapstructure:"DefaultPagePerNum"`       //默认每页显示条目数
	DefaultInsertDaraNum int           `json:"defaultInsertDaraNum" mapstructure:"DefaultInsertDaraNum"` //默认插入数据数量
	DefaultUserFavorPage int           `json:"defaultUserFavorPage" mapstructure:"DefaultUserFavorPage"` //默认用户喜欢的页面
	DefaultUserFavorSize int           `json:"defaultUserFavorSize" mapstructure:"DefaultUserFavorSize"` //默认用户喜欢的大小
	LockTicketTime       time.Duration `json:"lockTicketTime" mapstructure:"LockTicketTime"`             //锁定票务的时间间隔
	LockOrderTime        time.Duration `json:"lockOrderTime" mapstructure:"LockOrderTime"`               //锁定订单的时间间隔
}

// Work 任务工作
type Work struct {
	TaskChanCapacity   int `json:"taskChanCapacity" mapstructure:"TaskChanCapacity"`     //任务通道容量
	WorkerChanCapacity int `json:"workerChanCapacity" mapstructure:"WorkerChanCapacity"` //工作者通道容量
	WorkerNum          int `json:"workerNum" mapstructure:"WorkerNum"`                   //工作者数量
}

// Token 令牌
type Token struct {
	Key              string        ` mapstructure:"Key"`              //令牌的键值
	AccessTokenTime  time.Duration ` mapstructure:"AccessTokenTime"`  //访问令牌有效期
	RefreshTokenTime time.Duration ` mapstructure:"RefreshTokenTime"` //刷新令牌有效期
	AuthType         string        ` mapstructure:"AuthType"`         //认证类型
	AuthKey          string        ` mapstructure:"AuthKey"`          //认证密钥
}

// AliyunOSS 阿里云存储服务
type AliyunOSS struct {
	Endpoint        string `json:"endpoint" mapstructure:"Endpoint" ` //访问域名，IP地址
	AccessKeyID     string `json:"accessKeyID" mapstructure:"AccessKeyID"`
	AccessKeySecret string `json:"accessKeySecret" mapstructure:"AccessKeySecret"`
	BucketName      string `json:"bucketName" mapstructure:"BucketName"` //存储桶名称
	BucketUrl       string `json:"bucketUrl" mapstructure:"BucketUrl"`   //存储桶URL
	BasePath        string `json:"basePath" mapstructure:"BasePath"`     //存储桶中存储对象的基本路径
}

// Auto 自动操作相关配置
type Auto struct {
	AutoFlushReadCount2DBTime time.Duration `json:"autoFlushReadCount2DBTime" mapstructure:"AutoFlushReadCount2DBTime" ` //自动读取计数器刷新到数据库的时间间隔
	PeopleFavorToCacheTime    time.Duration `json:"peopleFavorToCacheTime" mapstructure:"PeopleFavorToCacheTime"`        //用户喜好缓存时间
	DeleteOutTimeTime         time.Duration `json:"deleteOutTimeTime" mapstructure:"DeleteOutTimeTime"`                  //删除过期数据的时间间隔
}
