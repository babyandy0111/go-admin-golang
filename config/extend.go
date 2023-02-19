package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap AMap // 这里配置對應配置文件的結構即可
}

type AMap struct {
	Key string
}
