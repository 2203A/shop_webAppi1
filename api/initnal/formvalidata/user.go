package formvalidata

type User struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}
