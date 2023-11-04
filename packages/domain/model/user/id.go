package user

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common"

type UserID struct {
	*common.UUID
}

func NewUserID(v string) (*UserID, error) {
	u, err := common.NewUUID(v)
	if err != nil {
		return nil, err
	}
	i := UserID{u}
	return &i, err
}
