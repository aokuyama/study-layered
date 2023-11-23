package prisma

import (
	"fmt"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/event/guest"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/prisma/db"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
)

type eventRepositoryPrisma struct {
	prisma *prisma
}

func NewEventRepositoryPrisma(client *prisma) *eventRepositoryPrisma {
	r := eventRepositoryPrisma{client}
	return &r
}

func (r *eventRepositoryPrisma) Create(e *event.Event) error {
	d := e.Path().Digest()
	en, err := e.Path().Encrypt()
	if err != nil {
		return err
	}

	_, err = r.prisma.client().Event.CreateOne(
		db.Event.ID.Set(e.ID().String()),
		db.Event.PathDigest.Set(d[:]),
		db.Event.Path.Set(en.Data),
		db.Event.PathIv.Set(en.Iv),
		db.Event.Name.Set(e.Name().String()),
		db.Event.Circle.Link(db.Circle.ID.Set(e.CircleID().String())),
	).Exec(r.prisma.ctx)

	return err
}

func (r *eventRepositoryPrisma) Find(i *event.EventID) (*event.Event, error) {
	return newEventModel(r.prisma.client().Event.FindUnique(
		db.Event.ID.Equals(i.String()),
	).
		With(db.Event.EventUser.Fetch().With(db.EventUser.User.Fetch())).
		Exec(r.prisma.ctx))
}

func (r *eventRepositoryPrisma) FindByPath(p *path.Path) (*event.Event, error) {
	d := p.Digest()
	return newEventModel(r.prisma.client().Event.FindUnique(
		db.Event.PathDigest.Equals(d[:]),
	).
		With(db.Event.EventUser.Fetch().With(db.EventUser.User.Fetch())).
		Exec(r.prisma.ctx))
}

func (r *eventRepositoryPrisma) Update(after *event.Event, before *event.Event) error {
	var txs []transaction.Param

	for _, afterGuest := range after.Guest().Items() {
		beforeGuest := before.Guest().IdenticalItem(&afterGuest)
		if beforeGuest == nil {
			// 追加
			txs = append(txs,
				r.prisma.client().EventUser.CreateOne(
					db.EventUser.Number.Set(int(*afterGuest.Number())),
					db.EventUser.Event.Link(db.Event.ID.Set(after.ID().String())),
					db.EventUser.User.Link(db.User.ID.Set(afterGuest.UserID().String())),
				).Tx(),
			)
			// ユーザー名の更新
			txs = append(txs,
				r.prisma.client().User.FindUnique(
					db.User.ID.Equals(afterGuest.UserID().String()),
				).Update(
					db.User.Name.Set(*afterGuest.Name()),
				).Tx(),
			)
		} else if !beforeGuest.Equals(&afterGuest) {
			// 更新
			txs = append(txs,
				r.prisma.client().EventUser.FindUnique(
					db.EventUser.EventIDUserID(
						db.EventUser.EventID.Equals(after.ID().String()),
						db.EventUser.UserID.Equals(afterGuest.UserID().String()),
					),
				).Update(
					db.EventUser.Number.Set(int(*afterGuest.Number())),
				).Tx(),
			)
			// ユーザー名の更新
			txs = append(txs,
				r.prisma.client().User.FindUnique(
					db.User.ID.Equals(afterGuest.UserID().String()),
				).Update(
					db.User.Name.Set(*afterGuest.Name()),
				).Tx(),
			)
		}
	}
	for _, beforeGuest := range before.Guest().Items() {
		if !after.Guest().ExistsIdentical(&beforeGuest) {
			// 削除
			txs = append(txs,
				r.prisma.client().EventUser.FindUnique(
					db.EventUser.EventIDUserID(
						db.EventUser.EventID.Equals(after.ID().String()),
						db.EventUser.UserID.Equals(beforeGuest.UserID().String()),
					),
				).Delete().Tx(),
			)
		}
	}

	if err := r.prisma.client().Prisma.Transaction(txs...).Exec(r.prisma.ctx); err != nil {
		panic(err)
	}
	return nil
}

func newEventModel(f *db.EventModel, err error) (*event.Event, error) {
	if err != nil {
		return nil, fmt.Errorf("%w %w", errs.ErrNotFound, err)
	}
	en := path.Encrypted{
		Data: f.Path,
		Iv:   f.PathIv,
	}
	p2, err := path.DecryptPath(&en)
	if err != nil {
		panic(err)
	}

	g := []guest.GuestInput{}
	for _, u := range f.EventUser() {
		g = append(g, guest.GuestInput{
			UserID: u.UserID,
			Name:   getUserName(u.User().Name()),
			Number: uint8(u.Number),
		})
	}

	e, err := event.NewEvent(&event.EventInput{
		ID:       f.ID,
		CircleID: f.CircleID,
		Name:     f.Name,
		Path:     p2.RawValue(), // 一度値オブジェクトにしたものを文字列にして詰め直しているのがイケてないが仕方ない
		Guest:    g,
	})

	if err != nil {
		panic(err)
	}
	return e, nil
}

func getUserName(value string, ok bool) string {
	if len(value) > 0 {
		return value
	}
	return "Guest"
}
