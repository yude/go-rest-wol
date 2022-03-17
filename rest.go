// Rest API Implementations

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//restWakeUpWithComputerName - REST Handler for Processing URLS /api/computer/<computerName>
func restWakeUpWithComputerName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	computerName := vars["computerName"]

	var result WakeUpResponseObject
	result.Success = false

	// Ensure computerName is not empty
	if computerName == "" {
		result.Message = "ホスト名を指定してください。"
		result.ErrorObject = nil
		w.WriteHeader(http.StatusBadRequest)
		// Computername is empty
	} else {

		// Get Computer from List
		for _, c := range ComputerList {
			if c.Name == computerName {

				// We found the Computername
				if err := SendMagicPacket(c.Mac, c.BroadcastIPAddress, ""); err != nil {
					// We got an internal Error on SendMagicPacket
					w.WriteHeader(http.StatusInternalServerError)
					result.Success = false
					result.Message = "マジックパケットを送る処理の途中でエラーが発生しました。"
					result.ErrorObject = err
				} else {
					// Horray we send the WOL Packet succesfully
					result.Success = true
					result.Message = fmt.Sprintf("ホスト %s (MACアドレス: %s) を起動しました。", c.Name, c.Mac)
					result.ErrorObject = nil
				}
			}
		}

		if result.Success == false && result.ErrorObject == nil {
			// We could not find the Computername
			w.WriteHeader(http.StatusNotFound)
			result.Message = fmt.Sprintf("ホスト %s は登録されていません。", computerName)
		}
	}
	json.NewEncoder(w).Encode(result)
}
