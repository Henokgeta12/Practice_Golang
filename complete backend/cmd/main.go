package main

import(
	"complete backend\cmd\api"
	"log"
	"complete backend\cmd\db"
	"github.com/spf13/viper"
)

func main(){

	viper.SetConfigName(".env")

	viper.AddConfigPath("..")

	viper.AutomaticEnv()


	if err : viper.ReadInConfig(); err != nill {
		fmt.printf("error reading file  %s",err)
	}
	db ,err =db.newstorage(CFG{
		publichost : viper.getstring("PUBLICHOST")	
		user : viper.getstring("USER")
		port : viper.getstring("PORT")
		password :viper.getstring("DBPASSEWORD")
	})
	
	server := api.NewApiServer(":8080",nil)

	if err := server.RUN() err != nill{
		log.Fatal(err)
	}

}