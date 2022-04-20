package casamento

import (
	"testing"

	"github.com/renanperrudgarcia/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaCasamentoSemConvidados(t *testing.T) {
	s := NewCasamento()
	_, err := s.Calcula(domain.Parametros{
		Homens:          0,
		Mulheres:        0,
		Criancas:        0,
		Acompanhamentos: true,
	})
	assert.Equal(t, err.Error(), "não Existem convidados para Casamento?")
	// if err.Error() != "Homens e mulheres devem ser maiores que zero" {
	// 	t.Errorf("Recebido %s esperado %s", err.Error(), "Homens e mulheres devem ser maiores que zero")
	// }
}

func TestCalculaCasamento(t *testing.T) {
	s := NewCasamento()
	casamento, err := s.Calcula(domain.Parametros{
		Homens:          10,
		Mulheres:        10,
		Criancas:        5,
		Acompanhamentos: true,
	})
	assert.Nil(t, err)

	esperado := Casamento{
		TotalConvidados: 25,
		TotalComida:     8500,
		TotalBolo:       2800,
		NaoAlcoolicas:   10000,
		Alcoolicas:      10000,
	}
	assert.Equal(t, casamento, esperado)

}
