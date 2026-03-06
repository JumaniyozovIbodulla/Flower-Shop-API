package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	SmtpServer          = "smtp.gmail.com"
	SmtpPort            = "587"
	SmtpUsername        = "jumaniyozovibodulla07@gmail.com"
	SmtpPassword        = "pntm dene uuvh qavx"
	WORKER_TYPE         = "worker"
	EMPLOYEER_TYPE      = "empoyeer"
)

var SignedKey = []byte(`MAr3e9Qt72RK8vx559Pk7M4JCfpLKlDObASFp70`)