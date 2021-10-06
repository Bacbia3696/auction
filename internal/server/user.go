package server

import (
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"html"
	"net/http"
	"strings"
	"time"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type createUserRequest struct {
	RoleId        int32      `json:"roleId" binding:"required"`
	UserName      string    `json:"userName" binding:"required"`
	Password      string    `json:"password" binding:"required,min=6"`
	FullName      string    `json:"fullName" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	Address       string    `json:"address" binding:"required"`
	Phone         string    `json:"phone" binding:"required"`
	BirthDate     time.Time `json:"birthDate" `
	IdCard        string    `json:"idCard" binding:"required"`
	IdCardDate    time.Time `json:"idCardDate" `
	IdCardAddress string    `json:"idCardAddress" binding:"required"`
	BankId        string    `json:"bankId" binding:"required"`
	BankOwner     string    `json:"bankOwner" binding:"required"`
	BankName      string    `json:"bankName" binding:"required"`
	OrganizationName string    `json:"organizationName"`
	OrganizationId        string    `json:"organizationId"`
	OrganizationDate     string    `json:"organizationDate"`
	OrganizationAddress      string    `json:"organizationAddress"`
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
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func (s *Server) RegisterUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	//check roleId
	if req.RoleId < 2 || req.RoleId > 4 {
		ResponseErrMsg(ctx, nil, "RoleId invalid", -1)
		return
	}
	// organization
	if req.RoleId == 4{
		if govalidator.IsNull(req.OrganizationName) ||  govalidator.IsNull(req.OrganizationId)||  govalidator.IsNull(req.OrganizationAddress){
			ResponseErrMsg(ctx, nil, "Input invalid",1)
			return
		}
	}

	//checkUsername
	checkUsername, err := s.store.GetByUserName(ctx, req.UserName)
	if err == nil {
		if (db.User{}) != checkUsername {
			ResponseErrMsg(ctx, nil, "Username already exists ", 403)
			return
		}
	}
	//checkIdCard
	checkIdCard, err := s.store.GetByIdCard(ctx, req.IdCard)
	if err == nil {
		if (db.User{}) != checkIdCard {
			ResponseErrMsg(ctx, nil, "IdCard already exists ", 403)
			return
		}
	}
	//checkEmail
	checkEmail, err := s.store.GetByEmail(ctx, req.Email)
	if err == nil {
		if (db.User{}) != checkEmail {
			ResponseErrMsg(ctx, nil, "Email already exists ", 403)
			return
		}
	}
	// create new user
	req.IdCardDate = time.Now()
	req.BirthDate = time.Now()
	params := db.CreateUserParams{}
	copier.Copy(&params, req)
	params.Password = HashPassword(req.Password)
	params.UserName = req.UserName
	params.FullName = req.FullName
	params.IDCard = req.IdCard
	params.IDCardAddress = req.IdCardAddress
	params.BankID = req.BankId
	params.BankName = req.BankName
	params.BankOwner = req.BankOwner
	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	//add user role
	userRoleParam := db.CreateUserRoleParams{UserID: user.ID, RoleID: req.RoleId}
	_, err = s.store.CreateUserRole(ctx, userRoleParam)
	if err != nil {
		ResponseErr(ctx, err, 1)
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
			if (db.User{}) == user {
				ResponseErrMsg(ctx, nil, "User not exists ", 1)
				return
			}
			if user.Status <0 {
				ResponseErrMsg(ctx, nil, "User blocked ", 1)
				return
			}
			if CheckPasswordHash(req.Password, user.Password) {
				token, _ := token.GenToken(user)
				ResponseOK(ctx, token)
			} else {
				ResponseErrMsg(ctx, nil, "Unauthorized", 401)
			}
		} else {
			logrus.Error(err)
			ResponseErrMsg(ctx, nil, "Unauthorized", 401)
		}
	}
}
