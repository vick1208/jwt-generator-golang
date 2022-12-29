package productcontroller

import (
	"net/http"

	"github.com/vick1208/jwt-go/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":    1,
			"name":  "Kemeja",
			"stock": "1000",
		},
		{
			"id":    2,
			"name":  "Sepatu",
			"stock": "100",
		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
