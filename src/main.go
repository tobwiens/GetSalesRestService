package main
import (
	"net/http"
	"log"
	"rest/sale"
	"github.com/spf13/viper"
)

func main() {

	// Initialize viper configuration management
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	var serveSalesRestService sale.ServeSalesRestService
	http.Handle("/", serveSalesRestService)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

