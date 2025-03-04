package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/renanperrudgarcia/curso-go/domain"
	"github.com/renanperrudgarcia/curso-go/domain/batizado"
	"github.com/renanperrudgarcia/curso-go/domain/casamento"
	"github.com/renanperrudgarcia/curso-go/domain/churrasco"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/churrasco", calculaFesta(churrasco.NewChurrasco()))
	r.Post("/batizado", calculaFesta(batizado.NewBatizado()))
	r.Post("/casamento", calculaFesta(casamento.NewCasamento()))
	http.ListenAndServe(":3000", r)
}

func calculaFesta(s domain.Festa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var param domain.Parametros
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ch, err := s.Calcula(param)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		j, err := ch.ToJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_, err = w.Write(j)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
