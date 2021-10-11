package server

import (
	"database/sql"
	"fmt"
	"html"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	RoleId              int32                   `form:"roleId" binding:"required"`
	UserName            string                  `form:"userName" binding:"required"`
	Password            string                  `form:"password" binding:"required,min=6"`
	FullName            string                  `form:"fullName" binding:"required"`
	Email               string                  `form:"email" binding:"required,email"`
	Address             string                  `form:"address" binding:"required"`
	Phone               string                  `form:"phone" binding:"required"`
	BirthDate           time.Time               `form:"birthDate" binding:"required"`
	IdCard              string                  `form:"idCard" binding:"required"`
	IdCardDate          time.Time               `form:"idCardDate" binding:"required"`
	IdCardAddress       string                  `form:"idCardAddress" binding:"required"`
	BankId              string                  `form:"bankId" binding:"required"`
	BankOwner           string                  `form:"bankOwner" binding:"required"`
	BankName            string                  `form:"bankName" binding:"required"`
	OrganizationName    string                  `form:"organizationName"`
	OrganizationId      string                  `form:"organizationId"`
	OrganizationDate    string                  `form:"organizationDate"`
	OrganizationAddress string                  `form:"organizationAddress"`
	Images              []*multipart.FileHeader `form:"images" binding:"required"`
}
type UserLogin struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logrus.Infoln("error hash password", err)
	}
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
	if err := ctx.ShouldBind(&req); err != nil {
		ResponseErr(ctx, err, 1)
		return
	}
	//check roleId
	if req.RoleId < 2 || req.RoleId > 4 {
		ResponseErrMsg(ctx, nil, "RoleId invalid", -1)
		return
	}
	// organization
	if req.RoleId == 4 {
		if govalidator.IsNull(req.OrganizationName) || govalidator.IsNull(req.OrganizationId) || govalidator.IsNull(req.OrganizationAddress) {
			ResponseErrMsg(ctx, nil, "Input invalid", 1)
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
	// req.IdCardDate = time.Now()
	// req.BirthDate = time.Now()
	params := db.CreateUserParams{
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
	params.Password = HashPassword(req.Password)
	params.UserName = req.UserName
	params.FullName = req.FullName
	params.IDCard = req.IdCard
	params.IDCardAddress = req.IdCardAddress
	params.BankID = req.BankId
	params.BankName = req.BankName
	params.BankOwner = req.BankOwner
	if req.RoleId == 4 {
		params.OrganizationID = sql.NullString{
			String: req.OrganizationId,
			Valid:  true,
		}
		params.OrganizationName = sql.NullString{
			String: req.OrganizationName,
			Valid:  true,
		}
		params.OrganizationAddress = sql.NullString{
			String: req.OrganizationAddress,
			Valid:  true,
		}
	}
	params.Birthdate = sql.NullTime{
		Time:  req.BirthDate,
		Valid: true,
	}
	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		logrus.Infoln("error create user", err)
		ResponseErr(ctx, err, http.StatusInternalServerError)
		return
	}
	//add user role
	userRoleParam := db.CreateUserRoleParams{UserID: user.ID, RoleID: req.RoleId}
	_, err = s.store.CreateUserRole(ctx, userRoleParam)
	if err != nil {
		ResponseErr(ctx, err, 1)
		return
	}

	//handle img
	forms, err := ctx.MultipartForm()
	if err != nil {
		logrus.Infoln("error parse MultipartForm", err)
		ResponseErr(ctx, err, http.StatusInternalServerError)
		return
	}
	images := forms.File["images"]

	for i := 0; i < len(images); i++ {
		logrus.Infoln("images", images[i].Filename)
		fileNames := strings.Split(images[i].Filename, ".")
		if len(fileNames) < 2 {
			logrus.Infoln("file invalid")
			ResponseErrMsg(ctx, nil, "Images input invalid ", 403)
			return
		}
		fileName := fmt.Sprintf("static/img/%d_%s", user.ID, RandStringRunes(8)+"."+fileNames[1])
		err = ctx.SaveUploadedFile(images[i], fileName)
		if err != nil {
			logrus.Infoln("error save image", err)
			ResponseErr(ctx, err, http.StatusInternalServerError)
			return
		}
		_, err = s.store.CreateUserImage(ctx, db.CreateUserImageParams{
			UserID: user.ID,
			Url:    fileName,
			Type:   1,
		})
		if err != nil {
			logrus.Infoln("error save user image", err)
			ResponseErr(ctx, err, http.StatusInternalServerError)
			return
		}
	}
	ResponseOK(ctx, user)
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
			if user.Status < 0 {
				ResponseErrMsg(ctx, nil, "User blocked ", 1)
				return
			}
			if CheckPasswordHash(req.Password, user.Password) {
				token, err := token.GenToken(user)
				if err != nil {
					logrus.Infoln("error GenToken", err)
					ResponseErr(ctx, err, http.StatusInternalServerError)
					return
				}
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

type ChangePassword struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (s *Server) ChangePassword(ctx *gin.Context) {
	var req ChangePassword
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Claims)
	userId := authPayload.Id
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user, err := s.store.GetById(ctx, userId)
		if err == nil {
			if CheckPasswordHash(req.Password, user.Password) {
				newPassword := HashPassword(req.NewPassword)
				_, err = s.store.UpdatePassword(ctx, db.UpdatePasswordParams{
					Password: newPassword,
					ID:       userId,
				})
				if err == nil {
					ResponseOK(ctx, nil)
				}
			} else {
				ResponseErrMsg(ctx, nil, "Password invalid", 401)
			}
		} else {
			logrus.Error(err)
			ResponseErrMsg(ctx, nil, "Password invalid ", 401)
		}
	}
}
