package main

type config struct {
	Port string `envconfig:"PORT" default:"8080"`
	//DataBaseConnectionString      string `envconfig:"DB_CONNECTION_STRING" default:"-"`
	PostsDataBaseConnectionString  string `envconfig:"POSTS_DB_CONNECTION_STRING" default:"-"`
	UploadDataBaseConnectionString string `envconfig:"POSTS_DB_CONNECTION_STRING" default:"-"`
	UserDataBaseConnectionString   string `envconfig:"POSTS_DB_CONNECTION_STRING" default:"-"`
	DomainName                     string `envconfig:"DOMAIN_NAME" default:"-"` //for static file prefix
	PhotoPath                      string `envconfig:"PHOTO_PATH" default:"static/photo"`
	FilePath                       string `envconfig:"FILE_PATH" default:"static/file"`
}
