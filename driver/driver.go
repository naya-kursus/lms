package driver

import (
	"flag"
	"os"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
)

type Driver struct {
	Config    *gowok.Config
	SQL       *gowok.SQL
	Validator *gowok.Validator
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	pkg.PrepareRuntime()
	configPath := flag.String("config", "config.yaml", "")
	flag.Parse()

	conf := gowok.Must(
		gowok.NewConfig(os.OpenFile(*configPath, os.O_RDONLY, 600)),
	)
	sqlDB := gowok.Must(
		gowok.NewSQL(conf.Databases),
	)
	val := gowok.NewValidator()

	d = &Driver{
		Config:    conf,
		SQL:       &sqlDB,
		Validator: val,
	}
	return d
}
