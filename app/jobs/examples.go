package jobs

import (
	"fmt"
	"time"
)

// 需要将定義的struct 新增到字典中；
// 字典 key 可以配置到 自动任務 調用目標 中；
func InitJob() {
	jobList = map[string]JobsExec{
		"ExamplesOne": ExamplesOne{},
		// ...
	}
}

// 新新增的job 必须按照以下格式定義，并实现Exec函數
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 傳入參數是 string 所以 arg.(string)；請根據對應的類型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}

	return nil
}
