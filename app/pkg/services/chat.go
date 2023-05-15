package services

import (
	"fmt"
	"log"

	api_model "github.com/FurkanSamaraz/Dependency-Injection/app/pkg/model"
	api_structure "github.com/FurkanSamaraz/Dependency-Injection/app/structures"
	"gorm.io/gorm"
)

type ChatService struct {
	DB     *gorm.DB
	client *api_model.RedisModel
}

func NewChatService(database *gorm.DB, client *api_model.RedisModel) *ChatService {
	return &ChatService{
		DB:     database,
		client: client,
	}
}

type IChatInstance interface {
	VerifyContact(username string) *api_structure.Response
	ChatHistory(username1, username2, fromTS, toTS string) *api_structure.Response
	ContactList(username string) *api_structure.Response
}

func (r *ChatService) VerifyContact(username string) (*api_structure.Response, error) {
	// geçersiz kullanıcı adı ve şifre hatası verirse
	// geçerli kullanıcı ise yeni oturum oluştur
	res := &api_structure.Response{Status: true}

	status, err := r.client.IsUserExist(username)
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		res.Status = false
		res.Message = "invalid username"
	}

	return res, err
}

func (r *ChatService) ChatHistory(username1, username2 string) (*api_structure.Response, error) {
	// geçersiz kullanıcı adları hata verirse
	// geçerli kullanıcılar sohbetleri getirirse
	res := &api_structure.Response{}

	// Kullanıcının var olup olmadığını kontrol edin
	status1, err := r.client.IsUserExist(username1)
	if err != nil {
		log.Fatal(err)
	}
	if !status1 {
		res.Status = false
		res.Message = "invalid username"
	}

	status2, err := r.client.IsUserExist(username2)
	if err != nil {
		log.Fatal(err)
	}
	if !status2 {
		res.Status = false
		res.Message = "invalid username"
	}

	chats, err := r.client.FetchChatBetween(username1, username2)
	if err != nil {
		log.Println("error in fetch chat between", err)
		res.Message = "unable to fetch chat history. please try again later."
		return res, err
	}

	res.Status = true
	res.Data = chats
	return res, err
}

func (r *ChatService) ContactList(username string) (*api_structure.Response, error) {
	// kullanıcı adı geçersiz ise hata döndürür
	// geçerli kullanıcılar sohbetleri getirirse
	res := &api_structure.Response{}

	status, err := r.client.IsUserExist(username)
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		res.Status = false
		res.Message = "invalid username"
	}

	contactList, err := r.client.FetchContactList(username)
	if err != nil {
		log.Println("error in fetch contact list of username: ", username, err)
		res.Message = "unable to fetch contact list. please try again later."
		return res, err
	}
	fmt.Println(contactList)
	res.Status = true
	res.Data = contactList
	return res, err
}
