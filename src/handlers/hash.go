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
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if request.ContentLength == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("No content"))
		return
	}

	algorithm := "sha512"

	if request.URL.Query().Get("algorithm") != "" {
		_, found := algorithms[request.URL.Query().Get("algorithm")]

		if !found {

			mapKeys := utils.GetMapKeys(algorithms)
			message := "Invalid algorithm, supported algorithms are: " + strings.Join(mapKeys, ", ")

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(message))
			return
		}

		algorithm = request.URL.Query().Get("algorithm")
	}

	body := make([]byte, request.ContentLength)
	request.Body.Read(body)

	hash := algorithms[algorithm].(crypto.Hash).New()

	hash.Write(body)

	hexHash := hex.EncodeToString(hash.Sum(nil))
	writer.Write([]byte(hexHash))
}
