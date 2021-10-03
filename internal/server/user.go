package server

import (
	"golang.org/x/crypto/bcrypt"
	"html"
	"net"
	"net/http"
	"regexp"
	"strings"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/asaskevich/govalidator"


)

type createUserRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
	// Email    string `json:"email" binding:"required,email"`
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

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responeErr(ctx, err, 1)
		return
	}
	if govalidator.IsNull(req.UserName) {
		responeErr(ctx, nil, 1)
		return
	}
	// create new user
	params := db.CreateUserParams{}
	copier.Copy(&params, req)
	logrus.Debug("params", params)
	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		responseOK(ctx, user)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
