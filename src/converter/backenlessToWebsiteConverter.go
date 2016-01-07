package converter
import "encoding/json"

type JsonBackendlessSearchResponse struct {
	TotalObjects int
	Data []JsonBackendlessEntry
}

type JsonBackendlessEntry struct {
	ASIN string
	Units string
}

func ConvertJsonBackendlessToWebsiteFormat (jsonSearchResponse string) (websiteJson string, errorValue error) {
	// Create an on object
	var searchObject JsonBackendlessSearchResponse
	// Unmarshall json resonse into object, give address to unmarshall function
	errorValue = json.Unmarshal([]byte(jsonSearchResponse), &searchObject)
	jsonByteArray, errorValue := json.Marshal(searchObject)
	websiteJson = string(jsonByteArray)
	return
}
