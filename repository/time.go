package repository

import "time"

type timeRepository struct{}

type TimeRepository interface {
	Now(time time.Time) time.Time
}

func NewTimeRepository() TimeRepository {
	return &timeRepository{}
}

func (c *timeRepository) Now(now time.Time) time.Time {
	if now.IsZero() {
		return time.Now()
	}
	return now
}
