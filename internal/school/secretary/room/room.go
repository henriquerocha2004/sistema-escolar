package room

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
)

type Room struct {
	id          uuid.UUID
	code        string
	description string
	capacity    int
}

func New(code string, description string, capacity int) (*Room, error) {
	r := &Room{
		id: uuid.New(),
	}

	err := r.ChangeCode(code)
	if err != nil {
		return nil, err
	}

	r.ChangeDescription(description)
	err = r.ChangeCapacity(capacity)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func Load(id string, code string, description string, capacity int) (*Room, error) {
	r, err := New(code, description, capacity)
	if err != nil {
		return nil, err
	}

	_ = r.ChangeId(id)

	return r, err
}

func (r *Room) Code() string {
	return r.code
}

func (r *Room) Description() string {
	return r.description
}

func (r *Room) Capacity() int {
	return r.capacity
}

func (r *Room) Id() uuid.UUID {
	return r.id
}

func (r *Room) ChangeCapacity(capacity int) error {
	if capacity == 0 {
		return errors.New("capacity cannot be empty")
	}

	r.capacity = capacity

	return nil
}

func (r *Room) ChangeDescription(description string) {
	if description == "" {
		return
	}

	r.description = description
}

func (r *Room) ChangeCode(code string) error {
	if code == "" {
		return errors.New("code cannot be empty")
	}

	r.code = code

	return nil
}

func (r *Room) ChangeId(id string) error {
	if id == "" {
		return errors.New("room id cannot be empty")
	}

	room, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return errors.New("failed to change room id")
	}

	r.id = room

	return nil
}

func (r *Room) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id          string `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Capacity    int    `json:"capacity"`
	}{
		Id:          r.Id().String(),
		Code:        r.Code(),
		Description: r.Description(),
		Capacity:    r.Capacity(),
	})
}
