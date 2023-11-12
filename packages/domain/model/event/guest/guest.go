package guest

import "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

type Guest struct {
	userId user.UserID
	name   string
	number uint8
}
