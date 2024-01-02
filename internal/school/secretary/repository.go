package secretary

import (
	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/value_objects"
)

type ClassRoomRepository interface {
	Create(classRoom entities.ClassRoom) error
	Delete(id string) error
	Update(classRoom entities.ClassRoom) error
	FindAll(pagination common.Pagination) (*common.ClassRoomPaginationResult, error)
	FindById(id string) (*entities.ClassRoom, error)
	FindByIdLock(id string) (*entities.ClassRoom, error)
}

type RoomRepository interface {
	Create(room entities.Room) error
	Delete(id string) error
	Update(room entities.Room) error
	FindByCode(code string) (*entities.Room, error)
	FindAll(pagination common.Pagination) (*common.RoomPaginationResult, error)
	FindById(id string) (*entities.Room, error)
	SyncSchedule(scheduleDto dto.RoomScheduleDto) error
}

type ScheduleRoomRepository interface {
	Create(schedule entities.ScheduleClass) error
	Delete(id string) error
	Update(schedule entities.ScheduleClass) error
	FindById(id string) (*entities.ScheduleClass, error)
	FindAll(paginator common.Pagination) (*common.SchedulePaginationResult, error)
}

type SchoolYearRepository interface {
	Create(schoolYear entities.SchoolYear) error
	Delete(id string) error
	Update(schoolYear entities.SchoolYear) error
	FindById(id string) (*entities.SchoolYear, error)
	FindByYear(year string) (*entities.SchoolYear, error)
	FindAll(paginator common.Pagination) (*common.SchoolYearPaginationResult, error)
}

type RegistrationRepository interface {
	Create(registration entities.Registration) error
	SearchStudentAlreadyRegistered(studentId uuid.UUID, classRoomId uuid.UUID) (string, error)
}

type StudentRepository interface {
	Create(student entities.Student) error
}

type ParentRepository interface {
	Create(parent entities.Parent) error
}

type AddressRepository interface {
	Create(address value_objects.Address) error
}

type PhoneRepository interface {
	Create(phone value_objects.Phone) error
}
