package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/donh/identity/pkg/auth"
	"github.com/donh/identity/pkg/crypto"
	"github.com/donh/identity/pkg/email"
	"github.com/donh/identity/pkg/models"
	"github.com/donh/identity/pkg/storage"
	"github.com/donh/identity/pkg/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// SetAPIRoutes set API routes
func SetAPIRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/common/now", nowTimestamp).Methods("GET")
	router.HandleFunc("/api/v1/common/uuid", _uuid).Methods("GET")
	router.HandleFunc("/api/v1/dids", didByAddress).Methods("GET").Queries("address", "{address}")
	router.HandleFunc("/api/v1/emails/send", verificationCode).Methods("POST")
	router.HandleFunc("/api/v1/emails/activation", verify).Methods("PUT")
	router.HandleFunc("/api/v1/emails/{email}/did", didByEmail).Methods("GET")
	router.HandleFunc("/api/v1/emails/{email}/resend", verificationResend).Methods("GET")
	router.HandleFunc("/api/v1/emails/{email}/status", verificationStatus).Methods("GET")
	router.HandleFunc("/api/v1/files/upload", upload).Methods("POST")
	router.HandleFunc("/api/v1/kyc", kycInsert).Methods("POST")
	router.HandleFunc("/api/v1/kyc/{did}/claim", kycClaim).Methods("GET").Queries("expiration", "{expiration}", "status", "{status}")
	router.HandleFunc("/api/v1/kyc/{did}/query", kycQuery).Methods("GET").Queries("type", "{type}")
	router.HandleFunc("/api/v1/kyc/{did}/erc725/claim", kycERC725Claim).Methods("GET")
	router.HandleFunc("/api/v1/publickeys", publicKeyFromAddress).Methods("GET").Queries("address", "{address}")
	router.HandleFunc("/api/v1/users", getUser).Methods("GET")
	router.HandleFunc("/api/v1/users", addUser).Methods("POST")
	router.HandleFunc("/api/v1/users", updateUser).Methods("PUT")
	router.HandleFunc("/api/v1/users", removeUser).Methods("DELETE")
	router.HandleFunc("/api/v1/users/authentication", didAuth).Queries("jwt", "{jwt}")
	router.HandleFunc("/api/v1/users/jwt", dbInsert).Methods("POST")
	router.HandleFunc("/api/v1/users/jwt", dbUpdate).Methods("PUT")
	router.HandleFunc("/api/v1/users/login", wsHandler)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	handler := cors.Default().Handler(router)
	handler = c.Handler(handler)
	port := util.Config().Port
	s := ":" + strconv.Itoa(port)
	log.Println("API server started. Listening on port:", port)
	log.Fatal(http.ListenAndServe(s, handler))
}

func setResponse(w http.ResponseWriter, result interface{}, status int, errorMessage string) {
	response := models.ResponseWrapper{
		Result: result,
		Status: status,
		Error:  errorMessage,
		Time:   time.Now(),
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func _uuid(w http.ResponseWriter, r *http.Request) {
	result, err := util.UUID()
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func parseQueryString(r *http.Request) (fields []string, filters map[string]string) {
	fields = []string{}
	filters = map[string]string{}
	vars := r.URL.Query()
	if val, ok := vars["fields"]; ok {
		s := strings.ToLower(val[0])
		if len(s) > 0 {
			fields = strings.Split(s, ",")
		}
	}
	if val, ok := vars["email"]; ok {
		filters["email"] = strings.ToLower(val[0])
	} else if val, ok := vars["address"]; ok {
		filters["address"] = strings.ToLower(val[0])
	}
	return fields, filters
}

func userHandler(action string, w http.ResponseWriter, r *http.Request) {
	var err error
	user := map[string]string{}
	if action == "read" {
		fields, filters := parseQueryString(r)
		user, err = storage.Query(fields, filters)
	} else {
		req := map[string]interface{}{}
		body, errParseBody := ioutil.ReadAll(r.Body)
		if errParseBody != nil {
			http.Error(w, errParseBody.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch action {
		case "create":
			err = storage.Insert(req)
		case "update":
			err = storage.Update(req)
		case "delete":
			err = storage.Delete(req)
		default:
		}
	}
	if err == nil {
		if action == "read" {
			setResponse(w, user, http.StatusOK, "")
		} else {
			setResponse(w, "Success", http.StatusOK, "")
		}
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userHandler("read", w, r)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	userHandler("create", w, r)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userHandler("update", w, r)
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	userHandler("delete", w, r)
}

func dbInsert(w http.ResponseWriter, r *http.Request) {
	req := map[string]interface{}{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &req)
	jwt := req["jwt"].(string)

	result, err := storage.DatabaseInsert(jwt)
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func dbUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	req := map[string]interface{}{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &req)
	jwt := req["jwt"].(string)

	result, err := storage.DatabaseUpdate(jwt)
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func didAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := auth.DIDAuth(vars["jwt"])
	if result && err == nil {
		setResponse(w, "Success", http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func didByAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := storage.DIDByAddress(vars["address"])
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func didByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := storage.DIDByEmail(vars["email"])
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func kycClaim(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intStatus := -1

	switch vars["status"] {
	case "approved":
		intStatus = 0
	case "rejected":
		intStatus = 1
	case "pending":
		intStatus = 2
	default:
	}

	result := crypto.KYCClaim(vars["did"], vars["expiration"], intStatus)
	if intStatus >= 0 {
		setResponse(w, result, http.StatusOK, "")
	} else {
		errorMessage := "Error: status should be one of the following: 'approved', 'rejected', or 'pending'."
		setResponse(w, "Failed", http.StatusBadRequest, errorMessage)
	}
}

func kycERC725Claim(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := auth.GetClaimSlice(vars["did"])
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func kycQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryType, _ := strconv.Atoi(vars["type"])
	result, err := storage.KYCClaimQuery(vars["did"], queryType)
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func kycInsert(w http.ResponseWriter, r *http.Request) {
	req := map[string]interface{}{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &req)
	claim := req["claim"].(string)
	signature := req["signature"].(string)
	result, err := storage.KYCClaimInsert(claim, signature)
	if err == nil {
		setResponse(w, result, http.StatusOK, "")
	} else {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}
}

func nowTimestamp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	setResponse(w, time.Now().Unix(), http.StatusOK, "")
}

func publicKeyFromAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := crypto.PublicKeyFromAddress(vars["address"])
	setResponse(w, result, http.StatusOK, "")
}

func verificationCode(w http.ResponseWriter, r *http.Request) {
	req := map[string]string{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &req)
	userEmail := req["email"]

	_, err := email.VerificationCode(userEmail)
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	}

	result, err := email.Send(userEmail)
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	} else {
		setResponse(w, result, http.StatusOK, "")
	}
}

func verificationResend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := email.Send(vars["email"])
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	} else {
		setResponse(w, result, http.StatusOK, "")
	}
}

func verificationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := email.Status(vars["email"])
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	} else {
		setResponse(w, result, http.StatusOK, "")
	}
}

func verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	req := map[string]interface{}{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &req)
	userEmail := req["email"].(string)
	code := req["code"].(string)

	result, err := email.Verify(userEmail, code)
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	} else {
		setResponse(w, result, http.StatusOK, "")
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		log.Println("disconnect !!")
		c.Close()
	}()
	mtype, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	auth.Clients[string(msg)] = c
	cnt := 0
	for {
		if _, ok := auth.Clients[string(msg)]; ok {
			Timeout := util.Config().Timeout
			if cnt > int(Timeout) {
				err = c.WriteMessage(mtype, []byte("Websocket Timeout Error"))
				if err != nil {
					log.Println("write:", err)
				}
				delete(auth.Clients, string(msg))
				break
			}
			cnt++
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			break
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	Buf := bytes.Buffer{}

	_ = r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, _ = io.Copy(&Buf, file)
	fmt.Fprintf(w, "%v", handler.Header)

	DID := r.Form.Get("did")

	result, err := storage.PictureUpdate(DID, Buf.Bytes())
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
	} else {
		setResponse(w, result, http.StatusOK, "")
	}
}
