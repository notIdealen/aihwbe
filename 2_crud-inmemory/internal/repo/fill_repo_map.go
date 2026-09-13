package repo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/notIdealen/aihwbe.git/internal/domain"
)

func (r *MemoryMapDB) Fill() {
	path := os.Getenv("DATA_FOR_REPO")
	fileName := fmt.Sprintf("%s/%s.json", path, "tasks")
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0222)
	if err != nil {
		panic("Bad fill repo")
	}
	defer file.Close()

	var tasks []domain.Task
	dec := json.NewDecoder(file)

	err = dec.Decode(&tasks)
	if err != nil {
		panic("Decode invalid")
	}

	for _, task := range tasks {
		r.db[task.ID] = task
	}
}
