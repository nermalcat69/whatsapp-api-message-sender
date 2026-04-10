package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func parseVCF(data string) []Contact {
	var contacts []Contact
	var current Contact

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")

		if strings.HasPrefix(line, "BEGIN:VCARD") {
			current = Contact{}
		} else if strings.HasPrefix(line, "FN:") {
			current.Name = strings.TrimPrefix(line, "FN:")
		} else if strings.HasPrefix(line, "N:") && current.Name == "" {
			// Fallback: parse N: field (Last;First;Middle;Prefix;Suffix)
			parts := strings.Split(strings.TrimPrefix(line, "N:"), ";")
			var nameParts []string
			if len(parts) > 1 && parts[1] != "" {
				nameParts = append(nameParts, parts[1])
			}
			if parts[0] != "" {
				nameParts = append(nameParts, parts[0])
			}
			current.Name = strings.Join(nameParts, " ")
		} else if strings.Contains(line, "TEL") {
			// Handle TEL;TYPE=...: or TEL:
			idx := strings.LastIndex(line, ":")
			if idx != -1 {
				phone := line[idx+1:]
				phone = strings.ReplaceAll(phone, " ", "")
				phone = strings.ReplaceAll(phone, "-", "")
				phone = strings.ReplaceAll(phone, "(", "")
				phone = strings.ReplaceAll(phone, ")", "")
				if phone != "" && current.Phone == "" {
					current.Phone = phone
				}
			}
		} else if strings.HasPrefix(line, "END:VCARD") {
			if current.Phone != "" {
				if current.Name == "" {
					current.Name = current.Phone
				}
				contacts = append(contacts, current)
			}
		}
	}
	return contacts
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/parse-vcf", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, "failed to parse form", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("vcf")
		if err != nil {
			http.Error(w, "failed to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "failed to read file", http.StatusInternalServerError)
			return
		}

		contacts := parseVCF(string(data))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6979"
	}
	fmt.Printf("Server running at http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
