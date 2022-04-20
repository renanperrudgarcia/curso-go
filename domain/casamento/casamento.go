package casamento

import (
	"encoding/json"
	"fmt"

	"github.com/renanperrudgarcia/curso-go/domain"
)

type Casamento struct {
	TotalConvidados int `json:"total-convidados"`
	TotalComida     int `json:"total-comida"`
	TotalBolo       int `json:"total-bolo"`
	NaoAlcoolicas   int `json:"nao-alcoolicas"`
	Alcoolicas      int `json:"alcoolicas"`
}

func (casamento Casamento) ToJSON() ([]byte, error) {
	return json.Marshal(casamento)
}

type Service struct{}

func NewCasamento() Service {
	return Service{}
}

func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	casamento := Casamento{}

	if p.Mulheres <= 0 || p.Homens <= 0 || p.Criancas <= 0 {
		return casamento, fmt.Errorf("não Existem convidados para Casamento?")
	}

	casamento.TotalConvidados = p.Mulheres + p.Homens + p.Criancas
	casamento.TotalComida = p.Mulheres*350 + p.Homens*400 + p.Criancas*200
	casamento.TotalBolo = p.Mulheres*100 + p.Homens*140 + p.Criancas*80
	casamento.NaoAlcoolicas = 400 * casamento.TotalConvidados
	casamento.Alcoolicas = 500 * (p.Mulheres + p.Homens)

	return casamento, nil

}
