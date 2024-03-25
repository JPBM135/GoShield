package handlers

import (
	"net/http"

	"jpbm135.go-shield/pkg/utils"
)

// var INFO_STRUCTURE = map[string]interface{}{
// 	"author":    "jpbm135",
// 	"version":   "1.0.0",
// 	"goVersion": "1.20",
// }

type Info struct {
	Author    string `json:"author"`
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	inf := Info{
		Author:    "jpbm135",
		Version:   "1.0.0",
		GoVersion: "1.22.1",
	}

	utils.WriteJSON(writer, inf)
}
