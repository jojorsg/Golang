package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco (t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia 06", "Rodovia"},
		{"Praça DBA", "Tipo Inválido"},
		{"Estrada 17", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo Inválido"},
		{"AVENIDA DAS DORES", "Avenida"},
		{" ", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido '%s' é diferente do esperado '%s'", 
   				tipoDeEnderecoRecebido,
    			cenario.retornoEsperado,
			)
		}
	}


	// enderecoParaTeste := "Rua Paulista"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", 
	// 		tipoDeEnderecoEsperado,
	// 		tipoDeEnderecoRecebido)
	// } 
}