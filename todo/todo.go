package todoapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

const filename = "todo.json"

type Items struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

var list []Items

func (item Items) Add(s string) {
	l := Items{
		Task:        s,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	fmt.Println(l)
	list = append(list, l)

}

func (item Items) Delete(index int) error {
	if index < 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list = append(list[:index-1], list[index:]...)
	return nil
}

func (item Items) Complete(index int) error {
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list[index-1].Done = true
	list[index-1].CompletedAt = time.Now()
	return nil
}

func List() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "\033[34m#\033[0m"},
			{Align: simpletable.AlignCenter, Text: "\033[34mTask\033[0m"},
			{Align: simpletable.AlignCenter, Text: "\033[34mDone?\033[0m"},
			{Align: simpletable.AlignCenter, Text: "\033[34mCreatedAt\033[0m"},
			{Align: simpletable.AlignCenter, Text: "\033[34mCompletedAt\033[0m"},
		},
	}
	var cells [][]*simpletable.Cell
	for index, item := range list {
		index++
		done := "\033[35mno\033[0m"
		task := fmt.Sprintf("\033[35m%s\033[0m", item.Task)
		if item.Done {
			task= "✅ \033[32m"+item.Task+"\033[0m"
			done = "\033[32myes\033[0m"
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format("2006-01-02 15:04:05")},
			{Text: item.CompletedAt.Format("2006-01-02 15:04:05")},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("\033[31mYou have %d TODO\033[0m", CountItems())},
	}}
	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func Save() error {

	t, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return err
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
	return ioutil.WriteFile(filename, t, 0644)

}

func Load() error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &list)
	return err
}

func CountItems() int {
	count := 0
	for _, item := range list {
		if !item.Done {
			count++
		}
	}
	return count
}

func (item Items) InComplete(index int) error {
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list[index-1].Done = false
	list[index-1].CompletedAt = time.Time{}
	return nil
}
