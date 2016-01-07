package sale
import (
	"net/http"
	"backendless"
	"converter"
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
	if  errorValue := salesRestService.doHTTPRequest(responseWriter, request); errorValue != nil {
		responseWriter.Write([]byte(errorValue.Error()))
	}
}

func (salesRestService ServeSalesRestService) doHTTPRequest(responseWriter http.ResponseWriter, request *http.Request) ( returnError error) {
	// Get query parameters from URL
	fromDate := request.URL.Query().Get(FromDateKey)
	toDate := request.URL.Query().Get(ToDateKey)
	asin := request.URL.Query().Get(AsinKey)
	// Query backendless API for data from, to for a given ASIN -- Return with error if error occurs
	backendlessResponse, returnError := backendless.BackendlessSearchAsinFromDateToDate(fromDate, toDate, asin)
	if (returnError != nil) {
		return
	}
	// Convert received json into
	convertedJson, returnError := converter.ConvertJsonBackendlessToWebsiteFormat(backendlessResponse)
	// Write json as a response
	responseWriter.Write([]byte(convertedJson))
	return
}
