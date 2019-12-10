package handler

//type Login struct {
//    User     string `form:"user" json:"user" binding:"required"`
//    Password string `form:"password" json:"password" binding:"required"`
//}

type Registration struct {
    User     string `form:"user" json:"user" binding:"required"`
    Email    string `form:"email" json:"email" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
