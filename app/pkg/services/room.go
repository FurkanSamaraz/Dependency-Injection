package services

import "github.com/FurkanSamaraz/Dependency-Injection/app/structures"

// Kullanıcıyı Odaya Ekle
func AddClientToRoom(room *structures.Room, username string, client *structures.Client) {
	room.AddClientToRoom(username, client)
}

// Kullanıcıyı Odadan Çıkar
func RemoveClientFromRoom(room *structures.Room, username string) {
	room.RemoveClientFromRoom(username)
}
