package server

import (
	"database/sql"
	"fmt"
	"html"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type createUserRequest struct {
	RoleId              int64                   `form:"roleId" binding:"required"`
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
	FrontImage          multipart.FileHeader    `form:"frontImage" binding:"required"`
	BackImage           multipart.FileHeader    `form:"backImage" binding:"required"`
	BusinessRegImage    multipart.FileHeader    `form:"businessRegImage"`
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
	res, err := s.registerUser(ctx)
	SendResponse(ctx, res, err)
}

func (s *Server) registerUser(ctx *gin.Context) (interface{}, *ServerError) {
	var req createUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		logrus.Error(err)
		return nil, ErrInvalidRequest.WithDevMsg(translateErr(err))
	}
	//check roleId
	if req.RoleId < 3 || req.RoleId > 4 {
		return nil, ErrInvalidRequest.WithDevMsg("RoleId invalid")
	}
	// organization
	if req.RoleId == 4 {
		if govalidator.IsNull(req.OrganizationName) || govalidator.IsNull(req.OrganizationId) || govalidator.IsNull(req.OrganizationAddress) {
			return nil, ErrInvalidRequest.WithDevMsg("Input invalid")
		}
	}

	//checkUsername
	_, err := s.store.GetByUserName(ctx, req.UserName)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, ErrGeneric
		}
	} else {
		return nil, ErrUserNameExisted
	}
	//TODO: checkIdCard
	//checkEmail
	_, err = s.store.GetByEmail(ctx, req.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, ErrGeneric
		}
	} else {
		return nil, ErrEmailExisted
	}
	// create new user
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
	var user db.User
	forms, err := ctx.MultipartForm()
	if err != nil {
		logrus.Infoln("error parse MultipartForm", err)
		return nil, ErrInvalidRequest.WithDevMsg("cannot parse MultipartForm")
	}
	frontImage := forms.File["frontImage"]
	backImage := forms.File["backImage"]
	businessRegImage := forms.File["businessRegImage"]
	images := forms.File["images"]

	err = s.store.ExecTx(ctx, func(q *db.Queries) (err error) {
		user, err = q.CreateUser(ctx, params)
		if err != nil {
			return
		}
		userRoleParam := db.CreateUserRoleParams{UserID: user.ID, RoleID: req.RoleId}
		_, err = q.CreateUserRole(ctx, userRoleParam)
		if err != nil {
			return
		}
		err = handleSaveImg(ctx, q, frontImage, user.ID, 1)
		if err != nil {
			return err
		}
		err = handleSaveImg(ctx, q, backImage, user.ID, 2)
		if err != nil {
			return err
		}
		err = handleSaveImg(ctx, q, businessRegImage, user.ID, 3)
		if err != nil {
			return err
		}
		return handleSaveImg(ctx, q, images, user.ID, 4)
	})
	if err != nil {
		logrus.Error(err)
		return nil, ErrGeneric
	}

	token, err := token.GenToken(user)
	if err != nil {
		logrus.Infoln("error GenToken", err)
		return nil, ErrGeneric
	}
	return token, nil
}

func handleSaveImg(ctx *gin.Context, q *db.Queries, images []*multipart.FileHeader, userId int64, typeId int32) error {
	for i := 0; i < len(images); i++ {
		fileName := fmt.Sprintf("static/img/%d_%s", userId, RandStringRunes(8)+filepath.Ext(images[i].Filename))
		err := ctx.SaveUploadedFile(images[i], fileName)
		if err != nil {
			return err
		}
		_, err = q.CreateUserImage(ctx, db.CreateUserImageParams{
			UserID: userId,
			Url:    fileName,
			Type:   typeId,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) LoginUser(ctx *gin.Context) {
	res, err := s.loginUser(ctx)
	SendResponse(ctx, res, err)
}

func (s *Server) loginUser(ctx *gin.Context) (interface{}, *ServerError) {
	var req UserLogin
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return nil, ErrInvalidRequest.WithDevMsg(err.Error())
	}
	user, err := s.store.GetByUserName(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidLogin
		}
		logrus.WithField("username", req.UserName).Error(err)
		return nil, ErrGeneric
	}
	if user.Status < 0 {
		return nil, ErrUserBlock
	}
	if CheckPasswordHash(req.Password, user.Password) {
		token, err := token.GenToken(user)
		if err != nil {
			logrus.WithField("username", req.UserName).Error(err)
			return nil, ErrGeneric
		}
		return token, nil
	}
	return nil, ErrInvalidLogin
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

func (s *Server) GetUserInfo(ctx *gin.Context) {
	userId := s.GetUserId(ctx)
	user, err := s.store.GetById(ctx, userId)
	if err == nil {
		images, _ := s.store.ListImage(ctx, userId)
		roleId, _ := s.store.GetRoleByUserId(ctx, userId)
		userInfo := UserInfo{
			User:   user,
			Images: images,
			RoleId: roleId,
		}
		ResponseOK(ctx, userInfo)
		return
	}
	logrus.Error("GetUserInfo ", err)
	ResponseErr(ctx, err, http.StatusInternalServerError)
	return
}
