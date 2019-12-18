package model

type Login struct{
Mail string 'form:"Mail" json:"Mail"'
Passward string 'aaaa:"aaa"'
}

type Login struct{
    UserID     string `form:"UserID" json:"UserID" xml:"UserID"  binding:"required"`
	Password   string `form:"Password" json:"Password" xml:"Password" binding:"required"`
}