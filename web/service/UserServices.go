package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"jiuban/model"
	"jiuban/repo"
	"log"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
)

const HASHKEY = "JiuBan2022"

type UserSerivce interface {
	VerifyEmail(email string) bool
	HashKey(account, password string) (key, hashPassword string)
	Login(ctx iris.Context, email, password string) error
	Register(ctx iris.Context, email, password, name string) error
	NewUserId(ctx iris.Context) string
	ForgotPassword(ctx iris.Context, email string) error
	SendEmail(body, email string)
	UpdateSetPassword(ctx iris.Context, email, password string) error
}
type userService struct {
	userRepo repo.UserRepo
	node     *snowflake.Node
}

func NewUserService(userRepo repo.UserRepo, node *snowflake.Node) UserSerivce {
	return &userService{
		userRepo: userRepo,
		node:     node,
	}
}

func (s *userService) VerifyEmail(email string) bool {
	pattern := `s([1-4]){1}([1-8]){1}([0-9]){3}([0-5]){1}([0-8]){1}([0-9]){3}@nutc\.edu\.tw`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func hash(key1, key2 string) string {
	h := hmac.New(sha256.New, []byte(HASHKEY))
	data := strings.TrimSpace(key1) + "/axolotl/" + strings.TrimSpace(key2)

	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *userService) HashKey(account, password string) (key, hashPassword string) {
	hashPassword = hash(HASHKEY, password)
	key = hash(account, hashPassword)
	return
}

func (s *userService) Login(ctx iris.Context, email, password string) error {
	result, err := s.userRepo.Get(ctx, email)
	if err != nil {
		ctx.WriteString("沒有此email\n")
		return err
	}
	key, hashpassword := s.HashKey(email, password)

	if result.Key == key && result.Password == hashpassword {
		ctx.JSON("登入成功")
	} else {
		ctx.WriteString("登入失敗 密碼錯誤\n")
	}

	return err
}

func (s *userService) Register(ctx iris.Context, email, password, name string) error {
	user := new(model.User)
	user.Email = email
	user.Password = password
	user.Name = name
	user.CreatedAt = time.Now()
	user.Id = s.NewUserId(ctx)
	_, err := s.userRepo.Get(ctx, email)

	if err == nil {
		ctx.WriteString("已有此帳戶\n")
		return err
	}
	checkEmail := s.VerifyEmail(email)
	if checkEmail == true {
		s.userRepo.Create(ctx, user)
		key, hashpassword := s.HashKey(email, password)

		s.userRepo.UpdateNewPassword(ctx, email, key, hashpassword)
		return nil
	}
	return nil

}

func (s *userService) ForgotPassword(ctx iris.Context, email string) error {
	_, err := s.userRepo.Get(ctx, email)
	if err == nil {
		ctx.WriteString("已寄信")
		return nil
	}
	s.SendEmail("qqq", email)

	return nil
}

func (s *userService) SendEmail(body, email string) {
	from := "abcd42822@gmail.com"
	pass := "vefoatrtwqmbhwbw"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: [JiuBan] 忘記密碼重置信\n\n" +
		"使用者您好 以下為您重置密碼的申請，請點擊連結進到重置密碼的頁面" +
		"如果您未使用此功能，請忽略此信 謝謝" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func (s *userService) NewUserId(ctx iris.Context) string {

	return s.node.Generate().String()

}

func (s *userService) UpdateSetPassword(ctx iris.Context, email, password string) error {
	user := new(model.User)
	user.Email = email
	user.Password = password
	user.UpdatedAt = time.Now()

	key, hashpassword := s.HashKey(email, password)
	s.userRepo.Update(ctx, user)
	s.userRepo.UpdateNewPassword(ctx, email, key, hashpassword)
	return nil

}
