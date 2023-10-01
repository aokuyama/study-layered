package owner

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type Owner struct {
	ID *OwnerID
}

func NewOwner(id *OwnerID) (*Owner, error) {
	c := Owner{id}
	return &c, nil
}

func GenerateOwner() (*Owner, error) {
	i, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	return NewOwner(i)
}

func (e *Owner) Identical(c *Owner) bool {
	return e.ID.Equals(c.ID)
}
