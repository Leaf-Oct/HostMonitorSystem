package main

type Config struct {
	host         string `yaml:"host" env-default:"127.0.0.1"`
	port         string `yaml:"port" env-default:"8080"`
	host_db_name string `yaml:"host-db-path" env-default:"host.db"`
}
