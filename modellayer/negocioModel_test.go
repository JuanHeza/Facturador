package modellayer

import (
	"log"
	"testing"
)

func TestNegocio_GenerateBearer(t *testing.T) {
	type args struct {
		period int
		tm     string
	}
	tests := []struct {
		name string
		ng   *Negocio
		args args
	}{
		{name: "test 1", ng: &Negocio{Clave: "evilPanda"}, args: args{period: 1, tm: "05/19/11, 10:47PM"}},
		{name: "test 1", ng: &Negocio{Clave: "evilPanda"}, args: args{period: 1, tm: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ng.GenerateBearer()
			log.Println(tt.ng.BearerToken)
		})
	}
}
