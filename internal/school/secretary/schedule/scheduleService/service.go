package scheduleService

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schoolyear"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/shared/paginator"
	"log"
)

type ServiceScheduleInterface interface {
	Create(dto schedule.Request) error
	Delete(id string) error
	Update(id string, dto schedule.Request) error
	FindOne(id string) (*schedule.ScheduleClass, error)
	FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error)
	SyncSchedule(scheduleRoomDto schedule.RoomScheduleDto) error
}

type ServiceScheduleClass struct {
	repository           schedule.Repository
	schoolYearRepository schoolyear.Repository
}

func New(
	repository schedule.Repository,
	schoolYearRepo schoolyear.Repository) *ServiceScheduleClass {
	return &ServiceScheduleClass{
		repository:           repository,
		schoolYearRepository: schoolYearRepo,
	}
}

func (s *ServiceScheduleClass) Create(dto schedule.Request) error {

	scheduleClass, err := schedule.New(dto.Description, dto.InitialTime, dto.FinalTime, dto.SchoolYear)
	if err != nil {
		return err
	}

	err = s.repository.Create(*scheduleClass)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create schedule")
	}

	return err
}

func (s *ServiceScheduleClass) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}

func (s *ServiceScheduleClass) Update(id string, dto schedule.Request) error {

	scheduleClass, err := schedule.New(dto.Description, dto.InitialTime, dto.FinalTime, dto.SchoolYear)
	if err != nil {
		return err
	}

	err = scheduleClass.ChangeId(id)
	if err != nil {
		return err
	}

	err = s.repository.Update(*scheduleClass)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update schedule")
	}

	return nil
}

func (s *ServiceScheduleClass) FindOne(id string) (*schedule.ScheduleClass, error) {
	scheduleClass, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get schedule")
	}

	return scheduleClass, nil
}

func (s *ServiceScheduleClass) FindAll(dtoRequest paginator.PaginatorRequest) (*paginator.PaginationResult, error) {
	paginator := paginator.Pagination{}
	paginator.FillFromDto(dtoRequest)
	schedules, err := s.repository.FindAll(paginator)

	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get schedules")
	}

	return schedules, nil
}

func (r *ServiceScheduleClass) SyncSchedule(scheduleRoomDto schedule.RoomScheduleDto) error {
	err := r.repository.SyncSchedule(scheduleRoomDto)
	if err != nil {
		log.Println(err)
		return errors.New("failed at sync schedules with room")
	}

	return nil
}
