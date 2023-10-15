//go:generate mockgen -source=$GOFILE -destination=.mock/$GOFILE
package circle

type CircleRepository interface {
	Save(*Circle) error
	Find(*CircleID) (*Circle, error)
}
