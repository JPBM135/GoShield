package handlers

import (
	"crypto"
	"encoding/hex"
	"net/http"
	"strings"

	"jpbm135.go-shield/src/utils"
)

var algorithms = map[string]interface{}{
	"sha512": crypto.SHA512,
	"sha256": crypto.SHA256,
	"sha1":   crypto.SHA1,
	"md5":    crypto.MD5,
}

func POSTHashHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if request.ContentLength == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("No content"))
		return
	}

	algorithm := "sha512"

	queryParamAlgorithm := request.URL.Query().Get("algorithm")

	if queryParamAlgorithm != "" {
		_, found := algorithms[queryParamAlgorithm]

		if !found {

			mapKeys := utils.GetMapKeys(algorithms)
			message := "Invalid algorithm, supported algorithms are: " + strings.Join(mapKeys, ", ")

			utils.WriteError(writer, message, http.StatusBadRequest)
			return
		}

		algorithm = queryParamAlgorithm
	}

	body := make([]byte, request.ContentLength)
	_, readBodyErr := request.Body.Read(body)

	if readBodyErr != nil {
		utils.WriteError(writer, "Error reading body", http.StatusBadRequest)
		return
	}

	hash := algorithms[algorithm].(crypto.Hash).New()

	_, hashWriteErr := hash.Write(body)

	if hashWriteErr != nil {
		utils.WriteError(writer, "Error writing hash", http.StatusInternalServerError)
		return
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(hexHash))
}
