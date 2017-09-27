package main
import(
log "github.com/sirupsen/logrus"
"gopkg.in/natefinch/lumberjack.v2"
	"os"  )
func init() {
		log.SetFormatter(&log.TextFormatter{})
		log.SetFormatter(&log.JSONFormatter{})
		//log.SetOutput(os.Stdout)
		file, err := os.OpenFile("LogDemo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(
				&lumberjack.Logger{
					Filename:   "./LogDemo.log",
					MaxSize:    1,  // megabytes after which new file is created
					MaxBackups: 3,  // number of backups
					MaxAge:     28, //days
				})
			log.Println(file.Name(),"is generated log file")		
		} else {
			log.Info("Failed to log to file, using default stderr")	
		}
}
func main(){
	log.Info("Main Function Started")
	SaveUser()
}
func SaveUser(){
	//log messages by specifying log leve
	log.Info("SaveUser Success") //to specify info messages
	log.Error("Error while saving user") //to specify error messages
	log.Debug("Debug")
	log.Warn("Warning")
}