package value_objects

import "errors"

type Shift string

func (s *Shift) Validate() error {
	currentVal := string(*s)
	allowedShifts := []string{
		"morning",
		"afternoon",
		"nocturnal",
		"full-time",
	}

	allowed := false

	for _, shift := range allowedShifts {
		if currentVal == shift {
			allowed = true
		}
	}

	if !allowed {
		return errors.New("invalid shift provided")
	}

	return nil
}
