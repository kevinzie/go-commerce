package models

type Profiles struct {
	ID       uint    `db:"id" json:"-"`
	UserId   uint    `db:"user_id" json:"-"`
	Phone    uint    `db:"phone" json:"phone"`
	Nicename string  `db:"nicename" json:"nicename"`
	Avatar   *string `db:"avatar" json:"avatar"`
	Company  string  `db:"company" json:"company"`
}
