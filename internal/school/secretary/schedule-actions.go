package secretary

import (
	"errors"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/common"
	"log"

	"github.com/google/uuid"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/entities"
)

type ScheduleActionsInterface interface {
	Create(dto dto.ScheduleRequestDto) error
	Delete(id string) error
	Update(id string, dto dto.ScheduleRequestDto) error
	FindOne(id string) (*entities.ScheduleClass, error)
	FindAll(dtoRequest dto.PaginatorRequest) (*common.SchedulePaginationResult, error)
}

type ScheduleClassActions struct {
	repository           ScheduleRoomRepository
	schoolYearRepository SchoolYearRepository
}

func NewScheduleClassActions(
	repository ScheduleRoomRepository,
	schoolYearRepo SchoolYearRepository) *ScheduleClassActions {
	return &ScheduleClassActions{
		repository:           repository,
		schoolYearRepository: schoolYearRepo,
	}
}

func (s *ScheduleClassActions) Create(dto dto.ScheduleRequestDto) error {
	schedule := entities.ScheduleClass{}
	err := schedule.FillFromDto(dto)
	if err != nil {
		return err
	}
	schedule.Id = uuid.New()

	schoolYear, err := s.schoolYearRepository.FindByYear(dto.SchoolYear)
	if err != nil || schoolYear == nil {
		log.Println(err)
		return errors.New("school year not found")
	}

	schedule.SchoolYear = schoolYear.Id

	err = s.repository.Create(schedule)
	if err != nil {
		log.Println(err)
		return errors.New("failed to create schedule")
	}

	return err
}

func (s *ScheduleClassActions) Delete(id string) error {
	err := s.repository.Delete(id)
	return err
}

func (s *ScheduleClassActions) Update(id string, dto dto.ScheduleRequestDto) error {
	schedule := entities.ScheduleClass{}
	err := schedule.FillFromDto(dto)
	if err != nil {
		return err
	}
	scheduleId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("id provided is invalid")
	}
	schedule.Id = scheduleId

	schoolYear, err := s.schoolYearRepository.FindByYear(dto.SchoolYear)
	if err != nil || schoolYear == nil {
		log.Println(err)
		return errors.New("school year not found")
	}

	schedule.SchoolYear = schoolYear.Id
	err = s.repository.Update(schedule)
	if err != nil {
		log.Println(err)
		return errors.New("failed to update schedule")
	}

	return nil
}

func (s *ScheduleClassActions) FindOne(id string) (*entities.ScheduleClass, error) {
	schedule, err := s.repository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get schedule")
	}

	return schedule, nil
}

func (s *ScheduleClassActions) FindAll(dtoRequest dto.PaginatorRequest) (*common.SchedulePaginationResult, error) {
	paginator := common.Pagination{}
	paginator.FillFromDto(dtoRequest)
	schedules, err := s.repository.FindAll(paginator)

	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to get schedules")
	}

	return schedules, nil
}
