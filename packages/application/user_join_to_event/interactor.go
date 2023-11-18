package usecase

import (
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
)

type userJoinToEventInteractor struct {
	repository event.EventRepository
}

func New(r event.EventRepository) UserJoinToEventUsecase {
	u := userJoinToEventInteractor{r}
	return &u
}

func (u *userJoinToEventInteractor) Invoke(i *UserJoinToEventInput) (*UserJoinToEventOutput, error) {
	var err error
	guest, err := guest.NewGuest(&i.UserID, &i.Name, &i.Number)
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

	afterEvent := event.JoinGuest(guest)

	if afterEvent != nil {
		err := u.repository.Update(afterEvent, event)
		if err != nil {
			return nil, err
		}
	}

	out := UserJoinToEventOutput{*event, *guest}
	return &out, nil
}
