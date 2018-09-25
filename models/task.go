package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID        uint      `json:"id" xorm:"'id' pk"`
	CreatedAt time.Time `json:"created_at" xorm:"'created_at'"`
	UpdatedAt time.Time `json:"updated_at" xorm:"'updated_at'"`
	Title     string    `json:"title" xorm:"'title'"`
}

func NewTask(title string, time time.Time) *Task {
	return &Task{
		Title:     title,
		CreatedAt: time,
		UpdatedAt: time,
	}
}

func (t *Task) GetTaskById() error {
	_, err := DBConnect().Get(t)
	return err
}

func GetAllTasks() []Task {
	var taskSlice []Task
	err := DBConnect().Find(&taskSlice)
	if err != nil {
		panic(err.Error())
	}

	return taskSlice
}

func (t *Task) Save() error {
	_, err := DBConnect().Insert(t)
	return err
}

func (t *Task) Update(title string) error {
	t.UpdatedAt = time.Now()
	t.Title = title
	_, err := DBConnect().ID(t.ID).Update(t)
	return err
}

func (t *Task) Delete() error {
	_, err := DBConnect().ID(t.ID).Delete(t)
	return err
}
