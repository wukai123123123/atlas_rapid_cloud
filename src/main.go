package main

import (
	"AtlasRapidCloud/src/config"
	"AtlasRapidCloud/src/store/pg"
	"fmt"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("./config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", cfg)
	_, err = pg.OpenDB(cfg.Database.GetDSN(), cfg.Database.MaxOpenConns, cfg.Database.MaxIdleConns, time.Duration(cfg.Database.ConnMaxLifetime)*time.Second)
	if err != nil {
		panic(err)
	}
}
