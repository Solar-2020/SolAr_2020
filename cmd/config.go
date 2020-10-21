package main

type config struct {
	Port string `envconfig:"PORT" default:"8099"`
	//DataBaseConnectionString      string `envconfig:"DB_CONNECTION_STRING" default:"-"`
	PostsDataBaseConnectionString  string `envconfig:"POSTS_DB_CONNECTION_STRING" default:"-"`
	UploadDataBaseConnectionString string `envconfig:"UPLOAD_DB_CONNECTION_STRING" default:"-"`
	//UserDataBaseConnectionString   string `envconfig:"USER_DB_CONNECTION_STRING" default:"-"`
	DomainName                     string `envconfig:"DOMAIN_NAME" default:"solar.ru"` //for static file prefix
	PhotoPath                      string `envconfig:"PHOTO_PATH" default:"/storage/photos"`
	FilePath                       string `envconfig:"FILE_PATH" default:"/storage/files"`
	InterviewService			   string `envconfig:"INTERVIEW_SERVICE" default:"localhost:8099"`
}
