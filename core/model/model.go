package model

const ModelBaseFile = `package model

import (
	"app/public"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// DBTime 自定义数据库时间类型
type DBTime struct {
	time.Time
}

func (t *DBTime) UnmarshalJSON(data []byte) error {
	var timestr = public.BytesToStr(data)
	if timestr == "null" {
		return nil
	}
	t1, err := time.Parse(public.SecondTimeFormat, strings.Trim(timestr, "\""))
	*t = DBTime{t1}
	return err
}

func (t DBTime) MarshalJSON() ([]byte, error) {
	return public.StrToBytes(fmt.Sprintf("\"%s\"", t.Format(public.SecondTimeFormat))), nil
}

func (t DBTime) Value() (driver.Value, error) {
	var zero time.Time
	if t.Time.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return t.Format(public.SecondTimeFormat), nil
}

func (t *DBTime) Scan(v interface{}) error {
	switch value := v.(type) {
	case string:
		n, _ := time.Parse(public.SecondTimeFormat, value)
		*t = DBTime{n}
	case time.Time:
		*t = DBTime{value}
	default:
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	return nil
}

// BaseModel 定义基础模型
type BaseModel struct {
	CreateTime DBTime ` + "`" + `gorm:"column:create_time;not null;autoCreateTime" json:"create_time"` + "`" + `
	UpdateTime DBTime ` + "`" + `gorm:"column:update_time;not null;autoUpdateTime" json:"update_time"` + "`" + `
}
`

const ModelTaskFile = `package model

// Task 任务
type Task struct {
	ID         uint    ` + "`" + `gorm:"primarykey" json:"id"` + "`" + `
	Title      string  ` + "`" + `gorm:"type:varchar;size:100;unique;uniqueIndex;not null" json:"title"` + "`" + ` // 任务名称
	Status     string  ` + "`" + `gorm:"type:varchar;size:20;index;default:'PENDING'" json:"status"` + "`" + `    // 任务状态
	StartTime  *DBTime ` + "`" + `gorm:"index"` + "`" + `                                                          // 开始时间
	FinishTime *DBTime ` + "`" + `gorm:"index"` + "`" + `                                                          // 结束时间

	BaseModel
}

const TaskTableName = "task"
`
