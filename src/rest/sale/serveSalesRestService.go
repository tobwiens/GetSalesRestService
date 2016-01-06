package sale
import (
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"backendless"
)

const (
	FromDateKey string = "fromDate"
	ToDateKey string = "toDate"
	AsinKey string = "ASIN"
)
type ServeSalesRestService struct {
}

func (salesRestService ServeSalesRestService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// Execute doHTTPRequest method and handle errors if error != nil
	if statusValue, errorValue := salesRestService.doHTTPRequest(responseWriter, request); errorValue != nil {
		// Handle errors
		switch statusValue {
		case http.StatusNotFound:
			fmt.Fprintln(responseWriter, "Not found")
		}
	}
}

func (salesRestService ServeSalesRestService) doHTTPRequest(responseWriter http.ResponseWriter, request *http.Request) (status int, returnError error) {

	fmt.Fprintln(responseWriter, viper.Get("message"))
	status = http.StatusOK
	fromDate := request.URL.Query().Get(FromDateKey)
	toDate := request.URL.Query().Get(ToDateKey)
	asin := request.URL.Query().Get(AsinKey)
	fmt.Fprintln(responseWriter, "fromDate: "+fromDate)
	fmt.Fprintln(responseWriter, "toDate: "+toDate)
	fmt.Fprintln(responseWriter, "ASIN: "+asin)

	jsonResponse, returnError := backendless.BackendlessSearchAsinFromDateToDate(fromDate, toDate, asin)
	if (returnError != nil) {
		fmt.Println(returnError)
	}
	fmt.Println(jsonResponse)

	return
}
