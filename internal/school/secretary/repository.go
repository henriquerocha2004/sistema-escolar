package secretary

import (
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ClassRoomRepository interface {
	Create(classRoom entities.ClassRoom) error
	Delete(id string) error
	Update(classRoom entities.ClassRoom) error
	FindAll(pagination common.Pagination) (*[]entities.ClassRoom, error)
	FindById(id string) (*entities.ClassRoom, error)
}

type RoomRepository interface {
	Create(room entities.Room) error
	Delete(id string) error
	Update(room entities.Room) error
	FindByCode(code string) (*entities.Room, error)
	FindAll(pagination common.Pagination) (*[]entities.Room, error)
	FindById(id string) (*entities.Room, error)
	SyncSchedule(scheduleDto dto.RoomScheduleDto) error
}

type ScheduleRoomRepository interface {
	Create(schedule entities.ScheduleClass) error
	Delete(id string) error
	Update(schedule entities.ScheduleClass) error
	FindById(id string) (*entities.ScheduleClass, error)
	FindAll(paginator common.Pagination) (*[]entities.ScheduleClass, error)
}

type SchoolYearRepository interface {
	Create(schoolYear entities.SchoolYear) error
	Delete(id string) error
	Update(schoolYear entities.SchoolYear) error
	FindById(id string) (*entities.SchoolYear, error)
	FindByYear(year string) (*entities.SchoolYear, error)
	FindAll(paginator common.Pagination) (*[]entities.SchoolYear, error)
}
