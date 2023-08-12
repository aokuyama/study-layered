package circle

type CircleRepository interface {
	Save(*Circle) (*Circle, error)
	LoadByID(*CircleID) (*Circle, error)
}
