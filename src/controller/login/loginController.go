package login

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login handler")
	student1 := student{"santosh", "1"}
	fmt.Println(student1)

	jsonResponse, jsonError := json.Marshal(student1)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(jsonError, string(jsonResponse))

	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student1)

}
