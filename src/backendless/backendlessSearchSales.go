package backendless
import (
	"net/http"
	"github.com/spf13/viper"
	"io/ioutil"
	"fmt"
)

func BackendlessSearchAsinFromDateToDate(fromDate string, toDate string, asin string) (jsonResponse string, errorValue error) {
	// Construct backendless url with search string
	var backendlessUrl string
	backendlessUrl += viper.GetString("backendlessEndpoint")
	backendlessUrl += viper.GetString("backendlessDataPath")
	// where=ASIN = '
	backendlessUrl += "?where=ASIN%20%3D%20'"
	// asin'
	backendlessUrl += asin
	backendlessUrl +="'"
	// AND
	backendlessUrl += "%20AND%20"
	// created >= from date
	backendlessUrl += "created%20%3E%3D%20'"
	backendlessUrl += fromDate
	// AND created <= to date
	backendlessUrl += "'%20AND%20created%20%3C%3D%20'"
	backendlessUrl += toDate
	backendlessUrl += "'"

	// Create request for URL
	request, errorValue := http.NewRequest("Get", backendlessUrl, nil)
	// If error occurred return with error
	if (errorValue != nil) {
		return
	}
	// Add header parameters to request
	request.Header.Add("application-id", viper.GetString("backendlessApplicationId"))
	request.Header.Add("secret-key", viper.GetString("backendlessSecretKey"))
	request.Header.Add("application-type", viper.GetString("backendlessApplicationType"))
	// Do the request
	httpClient := http.Client{}
	response, errorValue := httpClient.Do(request)
	// Return with error
	if (errorValue != nil) {
		return
	}
	// Close request et the end of this function
	defer response.Body.Close()
	// Read the request
	contents, errorValue := ioutil.ReadAll(response.Body)
	jsonResponse = string(contents)
	fmt.Println(jsonResponse)

	return
}
