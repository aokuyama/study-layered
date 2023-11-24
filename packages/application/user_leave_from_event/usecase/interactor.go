package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
)

type userLeaveFromEventInteractor struct {
	repository event.EventRepository
}

func New(r event.EventRepository) UserLeaveFromEventUsecase {
	u := userLeaveFromEventInteractor{r}
	return &u
}

func (u *userLeaveFromEventInteractor) Invoke(i *UserLeaveFromEventInput) (*UserLeaveFromEventOutput, error) {
	var err error
	uid, err := user.NewUserID(i.UserID)
	if err != nil {
		return nil, err
	}

	eid, err := event.NewEventID(i.EventID)
	if err != nil {
		return nil, err
	}

	event, err := u.repository.Find(eid)
	if err != nil {
		return nil, err
	}

	afterEvent := event.RemoveGuest(uid)
	if afterEvent == nil {
		return nil, errs.NewNotFound("guest not found")
	}

	err = u.repository.Update(afterEvent, event)
	if err != nil {
		return nil, err
	}

	out := UserLeaveFromEventOutput{*afterEvent}
	return &out, nil
}
