package rotas

import "net/http"

// todas as rotas da api
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r http.Request)
	RequerAutenticacao bool
}
