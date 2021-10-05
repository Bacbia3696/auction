package server

import (
	"golang.org/x/crypto/bcrypt"
	"html"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type createUserRequest struct {
	RoleId int `json:"roleId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Address string `json:"address" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	BirthDate time.Time `json:"birthDate" `
	IdCard string `json:"idCard" binding:"required"`
	IdCardDate time.Time `json:"idCardDate" `
	IdCardAddress string `json:"idCardAddress" binding:"required"`
	BankId string `json:"bankId" binding:"required"`
	BankOwner string `json:"bankOwner" binding:"required"`
	BankName string `json:"bankName" binding:"required"`
}

type UserLogin struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func Santize(data string) string{
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

func (s *Server) RegisterUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responeErr(ctx, err, 1)
		return
	}
	if govalidator.IsNull(req.UserName) {
		responeErr(ctx, nil, 1)
		return
	}
	//checkUsername
	checkUsername, err := s.store.GetByUserName(ctx, req.UserName)
	if err == nil {
		if (db.User{}) != checkUsername{
			responeErrMsg(ctx, nil,"Username already exists ", 403)
			return
		}
	}
	//checkIdCard
	checkIdCard, err := s.store.GetByIdCard(ctx, req.IdCard)
	if err == nil {
		if (db.User{}) != checkIdCard{
			responeErrMsg(ctx, nil,"IdCard already exists ", 403)
			return
		}
	}
	//checkEmail
	checkEmail, err := s.store.GetByEmail(ctx, req.Email)
	if err == nil {
		if (db.User{}) != checkEmail{
			responeErrMsg(ctx, nil,"Email already exists ", 403)
			return
		}
	}
	// create new user
	req.IdCardDate = time.Now()
	req.BirthDate = time.Now()
	params := db.CreateUserParams{}
	copier.Copy(&params, req)
	params.Password = HashPassword(req.Password)
	params.Username = req.UserName
	params.Fullname = req.FullName
	params.Idcard = req.IdCard
	params.Idcardaddress= req.IdCardAddress
	params.Bankid= req.BankId
	params.Bankname= req.BankName
	params.Bankowner= req.BankOwner
	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		responeErr(ctx, err,1)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (s *Server) LoginUser(ctx *gin.Context) {
	var req UserLogin
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user, err := s.store.GetByUserName(ctx, req.UserName)
		if err == nil {
			if CheckPasswordHash(req.Password,user.Password){
				token, _ := token.GenToken(user)
				responseOK(ctx, token)
			} else {
				responeErrMsg(ctx, nil,"Unauthorized", 401)
			}
		}else{
			responeErrMsg(ctx, nil,"Unauthorized", 401)
		}
	}
}