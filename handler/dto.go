package handler

//make sure:> `form:"user" json:"user" binding:"required"`, is there.(binding, form). Also use mustBindJson is used.
type Registration struct {
    UserName string `form:"user" json:"user" binding:"required"`
    Email    string `form:"email" json:"email" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

type Login struct {
    Email     string `form:"email" json:"email" xml:"email"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type CreateAccountDetails struct {
    Amount string `form:"amount" json:"amount" xml:"amount" binding:"required"`
}


type ApiKeyHeader struct {
    ApiKey string `header:"ApiKey"`
}
