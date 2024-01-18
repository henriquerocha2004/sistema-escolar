package registration

import "github.com/google/uuid"

type Repository interface {
	Create(registration Registration) error
	SearchStudentAlreadyRegistered(studentId uuid.UUID, classRoomId uuid.UUID) (string, error)
}
