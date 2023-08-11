package common

import "github.com/google/uuid"

type UUID string

func NewUUID(v string) (*UUID, error) {
	_, err := uuid.Parse(v)
	if err != nil {
		return nil, err
	}
	i := UUID(v)
	return &i, nil
}

func GenerateUUID() (*UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return NewUUID(id.String())
}

func (v *UUID) String() string {
	return string(*v)
}

func (v *UUID) Equals(c *UUID) bool {
	return v.String() == c.String()
}
