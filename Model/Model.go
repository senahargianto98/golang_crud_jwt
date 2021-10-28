package Model

// Untuk Migrate
// type Artikel struct {
// 	ID      int    `gorm:"primaryKey;autoIncrement"`
// 	Title   string `gorm:"column:title" json:"title"`
// 	Kontent string `gorm:"column:kontent" json:"kontent"`
// }

// type User struct {
// 	ID       int    `gorm:"primaryKey;autoIncrement"`
// 	Username string `gorm:"column:username" json:"username"`
// 	Password string `gorm:"column:password" json:"password"`
// }

type Artikel struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Kontent string `json:"kontent"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserLogin struct {
	UserName  string
	FirstName string
	LastName  string
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
