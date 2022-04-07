package services

import (
    "log"
    "io/ioutil"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/docgen"
)

var DOC_JSON_FILE_PATH = "/data/routes.json"

func ExportApiDocumentation(router chi.Router) {
    log.Printf("Generating documentation...")
    documentation := []byte(docgen.JSONRoutesDoc(router))

    log.Printf("Exporting documentation to %s\n", DOC_JSON_FILE_PATH)
    error := ioutil.WriteFile(DOC_JSON_FILE_PATH, documentation, 0644)
    if error != nil {
        log.Fatal(error)
    }
}
