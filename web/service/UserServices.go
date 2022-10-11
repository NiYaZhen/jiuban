package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"jiuban/model"
	"jiuban/repo"
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
		ctx.WriteString("登入成功\n")
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

func (s *userService) NewUserId(ctx iris.Context) string {

	return s.node.Generate().String()

}
