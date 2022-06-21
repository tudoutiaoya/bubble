package models

import (
	"bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	if err := dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}

func GetATodoByID(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateATodo(todo *Todo) (err error) {
	return dao.DB.Save(todo).Error
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
