package classroom

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type ClassRoom struct {
	id              uuid.UUID
	active          bool
	vacancyQuantity int
	shift           string
	openDate        time.Time
	occupiedVacancy int
	status          string
	level           string
	identification  string
	schoolYearId    uuid.UUID
	roomId          uuid.NullUUID
	scheduleId      uuid.UUID
	localization    string
	typeClass       string
}

func New(vacancyQuantity int,
	shift string,
	level string,
	identification string,
	schoolYearId string,
	roomId string,
	scheduleId string,
	localization string,
	typeClass string) (*ClassRoom, error) {
	cr := &ClassRoom{
		id:       uuid.New(),
		openDate: time.Now(),
		status:   "open",
		active:   true,
	}
	err := cr.ChangeVacancyQuantity(vacancyQuantity)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeShift(shift)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeLevel(level)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeIdentification(identification)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeSchoolYearId(schoolYearId)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeRoomId(roomId)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeScheduleId(scheduleId)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeLocalization(localization)
	if err != nil {
		return nil, err
	}

	err = cr.ChangeClassType(typeClass)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

func Load(
	id string,
	active bool,
	status string,
	vacanciesOccupied int,
	vacancyQuantity int,
	openDate string,
	shift string,
	level string,
	identification string,
	schoolYearId string,
	roomId string,
	scheduleId string,
	localization string,
	typeClass string,
) (*ClassRoom, error) {

	classRoom, err := New(
		vacancyQuantity,
		shift,
		level,
		identification,
		schoolYearId,
		roomId,
		scheduleId,
		localization,
		typeClass,
	)

	if err != nil {
		return nil, err
	}

	classRoom.status = status
	classRoom.occupiedVacancy = vacanciesOccupied
	classRoom.active = active
	err = classRoom.ChangeId(id)
	if err != nil {
		return nil, err
	}

	opDate, err := time.Parse("2006-01-02", openDate)
	if err != nil {
		return nil, err
	}

	classRoom.openDate = opDate

	return classRoom, nil
}

func (cr *ClassRoom) ChangeClassType(classType string) error {
	allowedTypes := []string{"in_person", "remote"}
	match := false

	for _, typeAllowed := range allowedTypes {
		if typeAllowed == classType {
			match = true
		}
	}

	if !match {
		return errors.New("invalid class type provided")
	}

	cr.typeClass = classType

	return nil
}

func (cr *ClassRoom) ChangeLocalization(localization string) error {
	if localization == "" {
		return errors.New("localization cannot be empty")
	}

	cr.localization = localization

	return nil
}

func (cr *ClassRoom) ChangeScheduleId(scheduleId string) error {
	if scheduleId == "" {
		return errors.New("schedule id cannot be empty")
	}

	schedule, err := uuid.Parse(scheduleId)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change schedule id")
	}

	cr.scheduleId = schedule

	return nil
}

func (cr *ClassRoom) ChangeRoomId(roomId string) error {
	if roomId == "" {
		return nil
	}

	room, err := uuid.Parse(roomId)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change room id")
	}

	cr.roomId = uuid.NullUUID{
		UUID:  room,
		Valid: true,
	}

	return nil
}

func (cr *ClassRoom) ChangeSchoolYearId(schoolYearId string) error {
	if schoolYearId == "" {
		return errors.New("schoolYear id cannot be empty")
	}

	schoolYear, err := uuid.Parse(schoolYearId)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change school year id")
	}

	cr.schoolYearId = schoolYear

	return nil
}

func (cr *ClassRoom) ChangeIdentification(identification string) error {

	if identification == "" {
		return errors.New("identification cannot be empty")
	}

	cr.identification = identification

	return nil
}

func (cr *ClassRoom) ChangeLevel(level string) error {

	if level == "" {
		return errors.New("level cannot be empty")
	}

	cr.level = level

	return nil
}

func (cr *ClassRoom) ChangeShift(shift string) error {
	allowedShifts := []string{"morning", "afternoon", "nocturnal", "full-time"}

	match := false

	for _, shiftAllowed := range allowedShifts {
		if shiftAllowed == shift {
			match = true
		}
	}

	if !match {
		return errors.New("invalid shift provided")
	}

	cr.shift = shift

	return nil
}

func (cr *ClassRoom) ChangeStatus(status string) error {
	allowedStatuses := []string{"open", "closed", "cancelled"}

	match := false

	for _, statusAllowed := range allowedStatuses {
		if statusAllowed == status {
			match = true
		}
	}

	if !match {
		return errors.New("invalid status provided")
	}

	cr.status = status

	return nil
}

func (cr *ClassRoom) ChangeVacancyQuantity(quantity int) error {
	if quantity == 0 {
		return errors.New("vacancy quantity cannot be empty")
	}

	cr.vacancyQuantity = quantity

	return nil
}

func (cr *ClassRoom) SetOccupiedVacancies(quantity int) error {

	err := cr.checkStatus()
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return nil
	}

	remainingVacancies := cr.vacancyQuantity - cr.occupiedVacancy

	if quantity > remainingVacancies {
		return errors.New("number of available vacancies is less than the number of vacancies requested")
	}

	if quantity > cr.vacancyQuantity {
		return errors.New("number of places for this class is less than the number requested")
	}

	cr.occupiedVacancy += quantity

	return nil
}

func (cr *ClassRoom) OccupiedVacancies() int {
	return cr.occupiedVacancy
}

func (cr *ClassRoom) checkStatus() error {
	if cr.status == "closed" {
		return errors.New("class Room is closed")
	}

	return nil
}

func (cr *ClassRoom) ChangeId(id string) error {
	if id == "" {
		return errors.New("class room id cannot be empty")
	}

	crId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change classroom id")
	}

	cr.id = crId

	return nil
}

func (cr *ClassRoom) Id() uuid.UUID {
	return cr.id
}

func (cr *ClassRoom) VacancyQuantity() int {
	return cr.vacancyQuantity
}

func (cr *ClassRoom) Shift() string {
	return cr.shift
}

func (cr *ClassRoom) OpenDate() time.Time {
	return cr.openDate
}

func (cr *ClassRoom) Status() string {
	return cr.status
}

func (cr *ClassRoom) Level() string {
	return cr.level
}

func (cr *ClassRoom) Identification() string {
	return cr.identification
}

func (cr *ClassRoom) SchoolYearId() uuid.UUID {
	return cr.schoolYearId
}

func (cr *ClassRoom) RoomId() uuid.NullUUID {
	return cr.roomId
}

func (cr *ClassRoom) ScheduleId() uuid.UUID {
	return cr.scheduleId
}

func (cr *ClassRoom) Localization() string {
	return cr.localization
}

func (cr *ClassRoom) TypeClass() string {
	return cr.typeClass
}

func (cr *ClassRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id              string `json:"id"`
		VacancyQuantity int    `json:"vacancy_quantity"`
		Shift           string `json:"shift"`
		OpenDate        string `json:"open_date"`
		OccupiedVacancy int    `json:"occupied_vacancy"`
		Status          string `json:"status"`
		Level           string `json:"level"`
		Identification  string `json:"identification"`
		SchoolYearId    string `json:"school_year_id"`
		RoomId          string `json:"room_id"`
		ScheduleId      string `json:"schedule_id"`
		Localization    string `json:"localization"`
		TypeClass       string `json:"type"`
	}{
		Id:              cr.id.String(),
		VacancyQuantity: cr.VacancyQuantity(),
		Shift:           cr.Shift(),
		OpenDate:        cr.OpenDate().Format("2006-01-02"),
		OccupiedVacancy: cr.OccupiedVacancies(),
		Status:          cr.Status(),
		Level:           cr.Level(),
		Identification:  cr.Identification(),
		SchoolYearId:    cr.SchoolYearId().String(),
		RoomId:          cr.RoomId().UUID.String(),
		ScheduleId:      cr.ScheduleId().String(),
		Localization:    cr.Localization(),
		TypeClass:       cr.TypeClass(),
	})
}
