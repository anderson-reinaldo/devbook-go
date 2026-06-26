package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type Erro struct {
	Erro string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
	}
}

func HandleStatusCode(w http.ResponseWriter, r *http.Response) {
	var erro Erro

	if r == nil || r.Body == nil {
		JSON(w, http.StatusInternalServerError, Erro{
			Erro: "resposta inválida",
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&erro); err != nil {
		JSON(w, http.StatusInternalServerError, Erro{
			Erro: "erro ao decodificar resposta",
		})
		return
	}
	fmt.Println(erro)
	JSON(w, r.StatusCode, erro)
}
