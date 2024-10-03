package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r http.Request) {

	w.Write([]byte("Criando usuario"))
}
func BuscarUsuarios(w http.ResponseWriter, r http.Request) {

	w.Write([]byte("buscando todos os usuarios"))
}
func BuscarUsuario(w http.ResponseWriter, r http.Request) {

	w.Write([]byte("buscando um usuario"))
}
func AtualizarUsuario(w http.ResponseWriter, r http.Request) {

	w.Write([]byte("atualizando usuario"))
}
func DeletarUsuario(w http.ResponseWriter, r http.Request) {

	w.Write([]byte("deletando usuario"))
}
