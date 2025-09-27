package service

import (
	"awesomeProject/dao"
	"awesomeProject/models"
	"awesomeProject/types"
	"fmt"
	"strings"
)

type TaskService struct {
	taskDAO *dao.TaskDAO
	repo    *dao.EsRepo
}

func NewTaskService(taskDAO *dao.TaskDAO, repo *dao.EsRepo) *TaskService {
	return &TaskService{taskDAO: taskDAO, repo: repo}
}

func (s *TaskService) CreateTask(task *types.CreateTask) error {
	task1 := &models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      models.TaskStatus(task.Status),
		UserID:      task.UserID,
		DueDate:     task.DueDate,
		ReminderAt:  task.ReminderAt,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		Tags:        strings.Join(task.Tags, ""),
	}
	// 1. 保存到MySQL
	if err := s.taskDAO.Create(task1); err != nil {
		return err
	}

	// 2. 索引到ES
	if err := s.repo.IndexTask(task); err != nil {
		fmt.Println(err)
		// 记录错误但不阻止流程
		// TODO: 可以考虑使用消息队列重试
	}

	return nil
}
