package user

import (
	"context"
)

// Delete ...
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.chatRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
