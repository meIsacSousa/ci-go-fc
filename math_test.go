package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado inválido: Expectativa: %d X Realidade: %d", 30, total)
	}

}
