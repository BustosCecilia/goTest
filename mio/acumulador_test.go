package mio

import (
	"fmt"
	"testing" // utilizo las caracteristicas de test de go
)

func TestAcumulado(t *testing.T) {

	lista := []int{6, 3, 1, 5, 7, 0, 2, 9, 8, 4}
	err := Acumular(lista)
	if err != nil {
		t.Error("No se esperaba un error")
	}
	if lista[9] != 45 {
		t.Error("no Se acumulo segun lo esperado")
	}

}

func TestAcumularListaNula(t *testing.T) {

	err := Acumular(nil)
	if err == nil {
		t.Error(fmt.Sprint("Se esperaba el error"))
	}
	errorEsperado := "la lista esta vacia"
	if err.Error() != errorEsperado {
		t.Error(fmt.Sprintf("Mensaje de error no esperado %s , mensaje de error recibido %s", errorEsperado, err.Error()))
	}
}

func BenchmarkAcumular10Elementos(b *testing.B) {
	lista := generarLista(10)
	for i := 0; i < b.N; i++ {
		Acumular(lista)
	}
}

func generarLista(n int) []int {
	var lista = make([]int, n)
	for i := 0; i < n; i++ {
		lista[i] = i
	}
	return lista
}
