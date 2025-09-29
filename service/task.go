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
func (s *TaskService) GetAllTasks(userID uint) ([]types.CreateTask, error) {
	return s.repo.Search("", "", userID)
}

func (s *TaskService) GetTask(query string, status types.TaskStatus, userID uint) ([]types.CreateTask, error) {
	task, err := s.repo.Search(query, status, userID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(taskID uint, userID uint, new *types.CreateTask) error {
	task, err := s.taskDAO.GetByID(taskID)
	if err != nil {
		return fmt.Errorf("task not exist")
	}
	if task.UserID != userID {
		return fmt.Errorf("this task not belong to you")
	}
	task.Title = new.Title
	task.Description = new.Description
	task.Status = models.TaskStatus(new.Status)
	task.DueDate = new.DueDate
	task.ReminderAt = new.ReminderAt
	task.CreatedAt = new.CreatedAt
	task.UpdatedAt = new.UpdatedAt
	task.Tags = strings.Join(new.Tags, ",")

	if err := s.taskDAO.Update(task); err != nil {
		return err
	}

	new.ID = taskID
	new.UserID = userID
	if err := s.repo.IndexTask(new); err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(taskID uint, userID uint) error {
	task, err := s.taskDAO.GetByID(taskID)
	if err != nil {
		return fmt.Errorf("task not exist")
	}
	if task.UserID != userID {
		return fmt.Errorf("this task not belong to you")
	}
	if err := s.repo.DeleteTask(taskID); err != nil {
		return err
	}
	if err := s.taskDAO.DeleteByID(taskID); err != nil {
		return err
	}
	return nil
}
