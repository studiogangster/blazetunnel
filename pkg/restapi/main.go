package restapi

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	// "blazetunnel/pkg/restapi"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

type TokenRequest struct {
	Id_token   string `json:"id_token"`
	App_id     string `json:"app_id"`
	Service_id string `json:"service_id"`
}

func StartRestApiServer() {

	flag.Parse()
	log.Println("wdad")
	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")
	ctx.Response.Header.Set("Content-Type", "application/json")

	tokenRequest := &TokenRequest{}
	err := json.Unmarshal(ctx.PostBody(), &tokenRequest)
	if err != nil {
		ctx.Response.SetBody([]byte(`{"status": false, "message": "Invalid Request Data"}`))
		return
	}

	id_token, err := getUid(tokenRequest.Id_token)

	if err != nil {
		log.Println("Eror", err)

		ctx.Response.SetBody([]byte(`{"status": false, "message": "Unauthorised Request"}`))
		return
	}

	tokenRequest.Id_token = id_token

	response, err := findService(tokenRequest.Id_token, tokenRequest.App_id, tokenRequest.Service_id)

	if err != nil {

		log.Println("Eror", err)

		ctx.Response.SetBody([]byte(`{"status": false, "message": "No data found"}`))
		return
	}

	token := []byte(fmt.Sprintf(`{"status": true, "auth_token":"%s"}`, response.Token))
	// token := []byte("{ \"auth_token\" : \"" + response.Token + "\" } ")

	ctx.Response.SetBody(token)
}
