package schedule

import "github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"

type Repository interface {
	Create(schedule ScheduleClass) error
	Delete(id string) error
	Update(schedule ScheduleClass) error
	FindById(id string) (*ScheduleClass, error)
	FindAll(paginator paginator.Pagination) (*paginator.PaginationResult, error)
	SyncSchedule(scheduleDto RoomScheduleDto) error
}
