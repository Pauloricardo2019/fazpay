package testing

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"kickoff/dto"
	"kickoff/internal/constants"
	"math"
)

var _ = Describe("Criar usuário", func() {
	apiClient := &ApiClient{}
	BeforeEach(func() {
		apiClient = NewApiClient()

		loginResponse, _ := apiClient.Login(&dto.LoginRequest{
			Login:    "admin_kickoff@digitalsys.tech",
			Password: "Test@123",
		})

		apiClient.SetBearerToken(loginResponse.Token)
	})

	When("Quando eu registrar um usuário", func() {

		Context("Com os campos preenchidos corretamente", func() {
			It("Deve retornar o usuário criado", func() {
				userCreated, err := apiClient.CreateUser(&dto.CreateUserRequest{
					ProfileType: "ADMIN",
					FullName:    "full_name",
					Email:       "test@gmail.com",
					Login:       "login",
					Password:    "password",
				})

				if err != nil {
					Fail(err.Message)
				}
				Expect(userCreated).ToNot(BeNil())
				Expect(userCreated.ID).To(BeNumerically(">", 0))
			})
		})

		Context("Com campos inválidos", func() {

			It("Digitar incorretamente o campo ProfileType, deverá retornar um erro de validação do campo ProfileType", func() {
				_, err := apiClient.CreateUser(&dto.CreateUserRequest{
					ProfileType: "ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_",
					FullName:    "full_name",
					Email:       "test@gmail.com",
					Login:       "login",
					Password:    "password",
				})

				Expect(err).ToNot(BeNil())
			})

			It("Digitar incorretamente o campo login, deverá retornar um erro na validação do campo login", func() {
				_, err := apiClient.CreateUser(&dto.CreateUserRequest{
					ProfileType: "ADMIN",
					FullName:    "full_name",
					Email:       "test.@@test@gmail.com",
					Login:       "login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_",
					Password:    "password",
				})

				Expect(err).ToNot(BeNil())
			})

			It("Digitar incorretamente o campo email, deverá retornar um erro na validação do email", func() {
				_, err := apiClient.CreateUser(&dto.CreateUserRequest{
					ProfileType: "ADMIN",
					FullName:    "full_name",
					Email:       "test.@@test@gmail.com",
					Login:       "login",
					Password:    "password",
				})

				Expect(err).ToNot(BeNil())
			})

		})

	})
})

var _ = Describe("Pegar usuário", func() {
	apiClient := &ApiClient{}
	var userID uint64
	BeforeEach(func() {
		apiClient = NewApiClient()

		loginResponse, _ := apiClient.Login(&dto.LoginRequest{
			Login:    "admin_kickoff@digitalsys.tech",
			Password: "Test@123",
		})
		apiClient.SetBearerToken(loginResponse.Token)

		userCreated, err := apiClient.CreateUser(&dto.CreateUserRequest{
			ProfileType: "ADMIN",
			FullName:    "full_name",
			Email:       "test@gmail.com",
			Login:       "login",
			Password:    "password",
		})

		if err != nil {
			Fail(err.Message)
		}
		Expect(userCreated).ToNot(BeNil())
		Expect(userCreated.ID).To(BeNumerically(">", 0))
		userID = userCreated.ID
	})

	When("Quando eu pegar um usuario pelo id", func() {

		Context("Passando um id existente na base", func() {
			It("Deve retornar o usuário vinculado ao id passado", func() {
				userFound, err := apiClient.GetUserByID(userID)
				if err != nil {
					Fail(err.Message)
				}
				Expect(userFound).ToNot(BeNil())
				Expect(userFound.ID).To(Equal(userID))
			})
		})

		Context("Passando um id que não existe na base", func() {

			It("Deve retornar um erro de usuário não encontrado", func() {

				_, err := apiClient.GetUserByID(math.MaxUint16)
				Expect(err).ToNot(BeNil())
				Expect(err.Message).To(Equal(constants.ErrorUserNotFound.Error()))

			})

		})

	})
})

var _ = Describe("Atualizar usuário", func() {
	apiClient := &ApiClient{}
	var userID uint64
	BeforeEach(func() {
		apiClient = NewApiClient()

		loginResponse, _ := apiClient.Login(&dto.LoginRequest{
			Login:    "admin_kickoff@digitalsys.tech",
			Password: "Test@123",
		})
		apiClient.SetBearerToken(loginResponse.Token)

		userCreated, err := apiClient.CreateUser(&dto.CreateUserRequest{
			ProfileType: "ADMIN",
			FullName:    "full_name",
			Email:       "test@gmail.com",
			Login:       "login",
			Password:    "password",
		})

		if err != nil {
			Fail(err.Message)
		}
		Expect(userCreated).ToNot(BeNil())
		Expect(userCreated.ID).To(BeNumerically(">", 0))
		userID = userCreated.ID
	})

	When("Quando eu atualizar um usuário", func() {

		Context("Preenchendo os campos de atualização corretamente", func() {
			It("Deverá atualizar o usuário e retornar um status code 204", func() {
				err := apiClient.UpdateUser(&dto.UpdateUserRequest{
					ProfileType: "TEST",
					FullName:    "full_name_test",
					Email:       "test3@gmail.com",
					Login:       "login_test",
				},
					userID,
				)
				if err != nil {
					Fail(err.Message)
				}
				Expect(err).To(BeNil())
			})
		})

		Context("Preenchendo os campos de atualização incorretamente", func() {
			It("Digitar o campo ProfileType incorretamente, deverá retornar um erro na validação do campo ", func() {
				err := apiClient.UpdateUser(&dto.UpdateUserRequest{
					ProfileType: "ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_ADMIN_",
					FullName:    "full_name_test",
					Email:       "test3@gmail.com",
					Login:       "login_test",
				},
					userID,
				)
				Expect(err).ToNot(BeNil())

			})

			It("Digitar incorretamente o campo login, deverá retornar um erro na validação do campo login", func() {
				err := apiClient.UpdateUser(&dto.UpdateUserRequest{
					ProfileType: "ADMIN",
					FullName:    "full_name",
					Email:       "test.@@test@gmail.com",
					Login:       "login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_login_",
				},

					userID,
				)

				Expect(err).ToNot(BeNil())
			})

			It("Digitar incorretamente o campo email, deverá retornar um erro na validação do email", func() {
				err := apiClient.UpdateUser(&dto.UpdateUserRequest{
					ProfileType: "ADMIN",
					FullName:    "full_name",
					Email:       "test.@@test@gmail.com",
					Login:       "login",
				},
					userID,
				)

				Expect(err).ToNot(BeNil())
			})

		})

		When("Quando eu atualizar um usuário e realizar o login com os novos dados.", func() {
			Context("Usuário atualizado com sucesso", func() {
				It("Deve retornar um status code de 204", func() {
					err := apiClient.UpdateUser(&dto.UpdateUserRequest{
						ProfileType: "User",
						FullName:    "full_name_test_login",
						Email:       "test_login@gmail.com",
						Login:       "login_test",
					},
						userID,
					)
					if err != nil {
						Fail(err.Message)
					}
					Expect(err).To(BeNil())
				})

				It("Deve retornar um token de acesso válido", func() {
					loginResponse, _ := apiClient.Login(&dto.LoginRequest{
						Login:    "test_login@gmail.com",
						Password: "Test@123",
					})
					Expect(loginResponse).ToNot(BeNil())
					Expect(loginResponse.Token).ToNot(BeEmpty())
					Expect(loginResponse.UserID).To(Equal(userID))
				})

			})
		})

	})
})

var _ = Describe("Excluir um usuário", func() {
	apiClient := &ApiClient{}
	var userID uint64
	BeforeEach(func() {
		apiClient = NewApiClient()

		loginResponse, _ := apiClient.Login(&dto.LoginRequest{
			Login:    "admin_kickoff@digitalsys.tech",
			Password: "Test@123",
		})
		apiClient.SetBearerToken(loginResponse.Token)

		userCreated, err := apiClient.CreateUser(&dto.CreateUserRequest{
			ProfileType: "ADMIN",
			FullName:    "full_name",
			Email:       "test@gmail.com",
			Login:       "login",
			Password:    "password",
		})

		if err != nil {
			Fail(err.Message)
		}
		Expect(userCreated).ToNot(BeNil())
		Expect(userCreated.ID).To(BeNumerically(">", 0))
		userID = userCreated.ID
	})

	When("Eu excluir um usuário do sistema", func() {

		Context("Passando o id correto", func() {
			It("Deve retornar um status 204", func() {
				err := apiClient.DeleteUser(userID)
				if err != nil {
					Fail(err.Message)
				}
				Expect(err).To(BeNil())
			})
		})

		Context("Passando um id aleatório", func() {
			It("Deve retornar um erro ao excluir o usuario", func() {
				err := apiClient.DeleteUser(math.MaxUint16)
				Expect(err).ToNot(BeNil())
				Expect(err.Message).To(Equal(constants.ErrorUserNotFound.Error()))
			})

		})

	})
})
