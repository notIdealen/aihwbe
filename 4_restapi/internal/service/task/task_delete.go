package service_task

import "context"

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {

	err := s.DB.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
