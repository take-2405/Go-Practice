package model

type Login struct {
	UserID   string `form:"UserID" json:"UserID" xml:"UserID"  binding:"required"`
	Password string `form:"Password" json:"Password" xml:"Password" binding:"required"`
}
