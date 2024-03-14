package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/constants"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	serviceIntf "github.com/Pauloricardo2019/teste_fazpay/internal/service/interface"
	"regexp"
)

type UserService struct {
	userRepository repositoryIntf.UserRepository
	logger         loggerIntf.LoggerInterface
}

func NewUserService(userRepository repositoryIntf.UserRepository, logger loggerIntf.LoggerInterface) serviceIntf.UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

type rulePassword struct {
	Name  string
	rule  string
	Error error
}

var rulePass = []rulePassword{
	{
		Name:  "min 7 characters",
		rule:  ".{7,}",
		Error: constants.PassMin7Characters,
	},
	{
		Name:  "min 1 letter",
		rule:  "([a-z]{1,})",
		Error: constants.PassMin1Letter,
	},
	{
		Name:  "min 1 letter uppercase",
		rule:  "([A-Z]{1,})",
		Error: constants.PassMin1LetterUpper,
	},
	{
		Name:  "min 1 number",
		rule:  "([0-9]{1,})",
		Error: constants.PassMin1Number,
	},
	{
		Name:  "min 1 special character",
		rule:  "([!@#$&*]{1,})",
		Error: constants.PassMin1SpecialCharacter,
	},
}

func (u *UserService) validatePassword(password string) error {
	if password == "" {
		return constants.PassCantBeEmpty
	}

	for _, r := range rulePass {
		if !regexp.MustCompile(r.rule).MatchString(password) {
			return r.Error
		}
	}
	return nil
}

func (u *UserService) validateEmail(email string) error {
	verifier := emailverifier.NewVerifier()

	ret, err := verifier.Verify(email)
	if err != nil {
		return err
	}

	if !ret.Syntax.Valid {
		return constants.InvalidEmail
	}

	return nil
}

func (u *UserService) validateUser(ctx context.Context, user *model.User) error {
	u.logger.LoggerInfo(ctx, "validateUser", "service")
	if err := user.Validate(); err != nil {
		return err
	}

	if err := u.validatePassword(user.Password); err != nil {
		return err
	}

	if err := u.validateEmail(user.Email); err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (bool, *model.User, error) {
	u.logger.LoggerInfo(ctx, "GetByEmail", "service")
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	u.logger.LoggerInfo(ctx, "Create", "service")
	if err := u.validateUser(ctx, user); err != nil {
		u.logger.LoggerError(ctx, err, "service")
		return nil, err
	}

	user = u.encryptAndClearingPassword(ctx, user)

	return u.userRepository.Create(ctx, user)
}

func (u *UserService) GetById(ctx context.Context, id uint64) (bool, *model.User, error) {
	return u.userRepository.GetById(ctx, id)
}

func (u *UserService) Update(ctx context.Context, user *model.User) error {

	if err := u.validateUser(ctx, user); err != nil {
		return err
	}
	user = u.encryptAndClearingPassword(ctx, user)
	return u.userRepository.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, id uint64) error {
	return u.userRepository.Delete(ctx, id)
}

func (u *UserService) GetByEmailAndPassword(ctx context.Context, user *model.User) (bool, *model.User, error) {
	user = u.encryptAndClearingPassword(ctx, user)
	return u.userRepository.GetByEmailAndPassword(ctx, user)
}

func (u *UserService) encryptAndClearingPassword(ctx context.Context, usuario *model.User) *model.User {
	u.logger.LoggerInfo(ctx, "encryptAndClearingPassword", "service")
	sum := sha256.Sum256([]byte(usuario.Password))
	usuario.HashedPassword = fmt.Sprintf("%x", sum)
	usuario.Password = ""

	return usuario
}
