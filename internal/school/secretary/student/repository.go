package student

type Repository interface {
	Create(student Student) error
}
