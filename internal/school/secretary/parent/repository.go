package parent

type Repository interface {
	Create(parent Parent) error
}
