package main

type Config struct {
	Host         string `yaml:"host" env-default:"127.0.0.1"`
	Port         string `yaml:"port" env-default:"8080"`
	Host_db_name string `yaml:"host-db-path" env-default:"host.db"`
}
