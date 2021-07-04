package model

const ServiceTaskFile = `package service

import (
	"app/model"
	"app/public"
	"app/schema"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// TaskCreate 任务创建业务逻辑
func TaskCreate(s schema.TaskCreate) (*model.Task, error) {
	if !errors.Is(public.Db.Table(model.TaskTableName).Select("id").
		Where("title = ?", s.Title).First(&model.Task{}).Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("任务（%s）已存在", s.Title)
	}
	var task = &model.Task{Title: s.Title, Status: "TODO"}
	if err := public.Db.Model(task).Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// TaskList 任务列表查询业务逻辑
func TaskList(s schema.TaskList) (data schema.QueryListSchema) {
	tx := public.Db.Table(model.TaskTableName)
	tx.Count(&data.TotalCount)
	switch s.Title {
	case "":
		data.FilterCount = data.TotalCount
	default:
		tx.Where("title like ?", public.LikeQueryJoin(s.Title))
	}
	var records []model.Task
	tx.Order("update_time desc").Offset(s.Size * (s.Page - 1)).Limit(s.Size).Find(&records)
	data.Records = records
	return data
}

// TaskStart 任务开始业务逻辑
func TaskStart(s schema.TaskIDSchema) (*model.Task, error) {
	var task = model.Task{ID: s.ID}
	if errors.Is(public.Db.Model(&task).First(&task).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not fount")
	}
	switch task.Status {
	case "DOING", "DONE":
		return nil, fmt.Errorf("当前任务不能开始")
	}
	task.StartTime = &model.DBTime{Time: time.Now()}
	task.Status = "DOING"
	if err := public.Db.Model(&task).Updates(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// TaskDone 任务完成业务逻辑
func TaskDone(s schema.TaskIDSchema) (*model.Task, error) {
	var task = model.Task{ID: s.ID}
	if errors.Is(public.Db.Model(&task).First(&task).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not fount")
	}
	switch task.Status {
	case "TODO", "DONE":
		return nil, fmt.Errorf("当前任务不能完成")
	}
	task.FinishTime = &model.DBTime{Time: time.Now()}
	task.Status = "DONE"
	if err := public.Db.Model(&task).Updates(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// TaskDelete 删除任务业务逻辑
func TaskDelete(s schema.TaskIDSchema) error {
	var task = model.Task{ID: s.ID}
	if errors.Is(public.Db.Model(&task).Select("id").First(&task).Error, gorm.ErrRecordNotFound) {
		return errors.New("not fount")
	}
	return public.Db.Model(&task).Delete(&task).Error
}
`
