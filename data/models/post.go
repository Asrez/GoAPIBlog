package models


type Post struct {
	BaseModel
	Title string `gorm:"type=string;size:256;not null";`
	Content string `gorm:"type=text;not null"`;
	User   User `gorm:"foreignKey:AuthorId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	AuthorId uint
}