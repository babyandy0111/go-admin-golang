package models

type SysJob struct {
	JobId          int    `json:"jobId" gorm:"primaryKey;autoIncrement;comment:流水號"`
	JobName        string `json:"jobName" gorm:"size:255;comment:名稱"`
	JobGroup       string `json:"jobGroup" gorm:"size:255;comment:任務分组"`
	JobType        int    `json:"jobType" gorm:"size:1;comment:任務類型"`
	CronExpression string `json:"cronExpression" gorm:"size:255;comment:cron表達式"`
	InvokeTarget   string `json:"invokeTarget" gorm:"size:255;comment:調用目標"`
	Args           string `json:"args" gorm:"size:255;comment:目標參數"`
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;comment:執行策略"`
	Concurrent     int    `json:"concurrent" gorm:"size:1;comment:是否併發"`
	Status         int    `json:"status" gorm:"size:1;comment:狀態"`
	EntryId        int    `json:"entry_id" gorm:"size:11;comment:job啟動時返回的id"`
	ModelTime
	ControlBy
}

func (SysJob) TableName() string {
	return "sys_job"
}
