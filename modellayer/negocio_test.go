package modellayer

import (
	"log"
	"testing"
	"time"
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
		{name: "test 1", ng: &Negocio{NegocioID: 12, Clave: "evilPanda"}, args: args{period: 1, tm: "05/19/11, 10:47PM"}},
		{name: "test 1", ng: &Negocio{NegocioID: 12, Clave: "evilPanda"}, args: args{period: 1, tm: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tm time.Time
			if tt.args.tm != "" {
				tm, _ = time.Parse("01/02/06, 03:04PM", tt.args.tm)
			}
			tt.ng.GenerateBearer(tm, tt.args.period)
			log.Println(tt.ng.BearerToken)
		})
	}
}
