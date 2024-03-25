package handlers

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"

	"jpbm135.go-shield/pkg/utils"
)

var algorithms = map[string]interface{}{
	"sha512": crypto.SHA512,
	"sha256": crypto.SHA256,
	"sha1":   crypto.SHA1,
	"md5":    crypto.MD5,
}

func POSTHashHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No content"))
		return
	}

	algorithm := "sha512"

	queryParamAlgorithm := r.URL.Query().Get("algorithm")

	if queryParamAlgorithm != "" {
		_, found := algorithms[queryParamAlgorithm]

		if !found {

			mapKeys := utils.GetMapKeys(algorithms)
			message := "Invalid algorithm, supported algorithms are: " + strings.Join(mapKeys, ", ")

			utils.WriteError(w, message, http.StatusBadRequest)
			return
		}

		algorithm = queryParamAlgorithm
	}

	//this line
	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("I am the error", err)
		utils.WriteError(w, "Error reading body", http.StatusBadRequest)
		return
	}

	hash := algorithms[algorithm].(crypto.Hash).New()

	_, hashWriteErr := hash.Write(bodyBytes)

	if hashWriteErr != nil {
		utils.WriteError(w, "Error writing hash", http.StatusInternalServerError)
		return
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hexHash))
}
