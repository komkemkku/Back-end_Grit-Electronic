package model

import "time"

type Paginate struct {
	Page       int64
	Size       int64
	Total      int64
}

// type Paginated struct {
// 	Page      int64
// 	Size      int64
// 	Total     int64
// 	Search    string
// 	Status    string
// 	StartDate time.Time
// 	EndDate   time.Time
// }

// model/paginate.go หรือไฟล์ที่คุณประกาศ struct Paginate

type CreateUpdateUnixTimestamp struct {
	CreateUnixTimestamp
	UpdateUnixTimestamp
}

type CreateUnixTimestamp struct {
	CreatedAt int64 `json:"created_at" bun:",notnull,default:EXTRACT(EPOCH FROM NOW())"`
}

type UpdateUnixTimestamp struct {
	UpdatedAt int64 `json:"updated_at" bun:",notnull,default:EXTRACT(EPOCH FROM NOW())"`
}

// type SoftDelete struct {
// 	DeletedAt *time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
// }

type SoftDelete struct {
	DeletedAt int64 `json:"deleted_at" bun:",soft_delete,nullzero"`
}

func (t *CreateUnixTimestamp) SetCreated(ts int64) {
	t.CreatedAt = ts
}

func (t *CreateUnixTimestamp) SetCreatedNow() {
	t.SetCreated(time.Now().Unix())
}

func (t *UpdateUnixTimestamp) SetUpdate(ts int64) {
	t.UpdatedAt = ts
}

func (t *UpdateUnixTimestamp) SetUpdateNow() {
	t.SetUpdate(time.Now().Unix())
}
