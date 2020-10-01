package main

type config struct {
	Port                     string `envconfig:"PORT" default:"8080"`
	DataBaseConnectionString string `envconfig:"DB_CONNECTION_STRING" default:"-"`
	DomainName               string `envconfig:"DOMAIN_NAME" default:"-"` //for static file prefix
	FileRootPath             string `envconfig:"FILE_ROOT_PATH" default:"static"`
}
