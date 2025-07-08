package utils


import(
	"net/http"
	"fmt"
	"encoding/json"
)

func parseJSON(r http.Request, payload any) error{
	if r.Body == nil{
		fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter , status int,v any) error{

	w.Header().add("content-type" , "applcation/json")
	w.WriteHeader(status)


	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
    WriteJSON(w, status, map[string]string{"error": err.Error()})
}
