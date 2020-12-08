// // Analysis of p
package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/verify-data-wranggle/models"
// )

// // DataDir base path of data directory
// const (
// 	DataDir = "data/"
// )

// func main() {
// 	internalErrorCodes := map[string][]models.Payment{}
// 	changedErrorCodesStat := map[string]int{}

// 	// TODO:- Step1 extract MozartId
// 	csv_file, err := os.Open(getFullPath("data.csv"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer csv_file.Close()

// 	r := csv.NewReader(csv_file)
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// Store the start time to analyze performance
// 	startTime := time.Now()
// 	// Sort in ascending order of time stamp
// 	// sort.Slice(records, func(x, y int) bool {
// 	// 	thetime1, err := time.Parse(records[x][10], time.RFC3339)

// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	thetime2, err := time.Parse(time.RFC3339, records[y][10])

// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	return thetime1.Unix() < thetime2.Unix()
// 	// })

// 	// 0 => context.response.data.gateway_response.status
// 	// 1 => context.response.data.status
// 	// 2 => context.response.error.description
// 	// 3 => context.response.error.gateway_error_code
// 	// 4 => context.response.error.gateway_error_description
// 	// 5 => context.response.error.gateway_status_code
// 	// 6 => context.response.error.internal_error_code
// 	// 7 => environment
// 	// 8 => mode
// 	// 9 => payment.payment_id
// 	// 10 => timestamp
// 	// Print all the field names
// 	for idx, fieldName := range records[0] {
// 		fmt.Println(idx, "=>", fieldName)
// 	}

// 	for _, record := range records {
// 		paymentID := record[9]

// 		// Check if paymentID is not blank and it's not added in Map
// 		if paymentID != "" && internalErrorCodes[paymentID] == nil {
// 			internalErrorCodes[paymentID] = []models.Payment{}
// 		}
// 		payment := models.Payment{
// 			GatewayResponseStatus: record[0],
// 			record[1],
// 			record[2],
// 			record[3],
// 			record[4],
// 			record[5],
// 			record[6],
// 			record[7],
// 			record[8],
// 			record[9],
// 			record[10],
// 		}

// 		internalErrorCodes[paymentID] = append(internalErrorCodes[paymentID], payment)

// 	}

// 	output, err := os.Create(getFullPath("output.json"))
// 	defer output.Close()
// 	output.WriteString("[\n")
// 	for paymentId, val := range internalErrorCodes {

// 		// fmt.Println("{\n\"", paymentId, "\": ", string(s), "\n}")
// 		// fmt.Println("Processing payment_id ", paymentId)
// 		// if length of array is >1 then only write this data to output file

// 		if len(val) < 2 {
// 			continue
// 		}

// 		// Create json struct
// 		jsonOutput := models.Log{paymentId, val}

// 		// Marshal it to JSON
// 		s, _ := json.MarshalIndent(jsonOutput, "", "\t")

// 		// Write formatted JSON String to file
// 		output.WriteString(string(s) + ",")

// 		n := len(val)
// 		initialStatus := val[n-1].GatewayResponseStatus
// 		initialInternalErrorCode := val[n-1].InternalErrorCode + "," + val[n-1].GatewayErrorCode

// 		if initialInternalErrorCode == "," {
// 			initialInternalErrorCode = "NO_ERROR_CODE"
// 		}
// 		for i := n - 2; i >= 0; i-- {

// 			// If status change then increase frequency
// 			if val[i].GatewayResponseStatus != initialStatus {
// 				if val, ok := changedErrorCodesStat[initialInternalErrorCode]; ok {
// 					changedErrorCodesStat[initialInternalErrorCode] = val + 1
// 				} else {
// 					changedErrorCodesStat[initialInternalErrorCode] = 0
// 				}
// 			}
// 		}

// 	}

// 	output.WriteString("\n]")

// 	errorStatFile, _ := os.Create(getFullPath("error_stat_output.json"))
// 	defer errorStatFile.Close()
// 	s, _ := json.MarshalIndent(changedErrorCodesStat, "", "\t")
// 	fmt.Println("Internal Error Code + Gateway Error Code frequency distribution :")
// 	fmt.Println(string(s))

// 	errorStatFile.WriteString(string(s))

// 	fmt.Println("Total processing time ", float64(time.Now().UnixNano()-startTime.UnixNano())/(1e+9), "s")

// 	plotDistribution(changedErrorCodesStat)
// }
