package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) UploadSystem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			json.NewEncoder(w).Encode(&map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
	case "POST":
		r.ParseMultipartForm(10 << 20)

		// Get handler for filename, size and headers
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		csvLines, err := csv.NewReader(file).ReadAll()

		err = h.services.UploadFile(csvLines)
		if err != nil {
			json.NewEncoder(w).Encode(&map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		fmt.Fprintf(w, "File has been uploaded\n")
	}
}
