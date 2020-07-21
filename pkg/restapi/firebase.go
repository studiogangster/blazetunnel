package restapi

import (
	"log"

	common "blazetunnel/common"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"

	"google.golang.org/api/option"
)

type ServiceData struct {
	Appname     string
	ServiceName string
	Domain      string
	Token       string
	Endpoint    string
}

var projectPath = "/secret/firebase.json"

func getUid(id_token string) (string, error) {
	opt := option.WithCredentialsFile(projectPath)

	ctx := context.TODO()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {

	}
	log.Println("debuggger", app, err)

	auth, _ := app.Auth(ctx)
	token, err := auth.VerifyIDToken(ctx, id_token)

	if err != nil {
		return "", err
	}

	uid := token.UID
	return uid, nil

}

func findService(uid string, appId string, serviceId string) (*ServiceData, error) {
	opt := option.WithCredentialsFile(projectPath)

	ctx := context.TODO()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	log.Println(app, err)

	firestore, _ := app.Firestore(ctx)
	// snapshot, err := firestore.Collection("app").Doc(uid).Collection("app").Doc(appId).Collection("service").Doc(serviceId).Get(ctx)

	ref := firestore.Collection("app").Doc(uid).Collection("app").Doc(appId)
	application, err := ref.Get(ctx)

	if err != nil {
		return nil, err
	}

	service, err := ref.Collection("service").Doc(serviceId).Get(ctx)

	if err != nil {
		return nil, err
	}

	applicationData := application.Data()
	serviceData := service.Data()
	log.Println(applicationData, serviceData)

	response := &ServiceData{
		Appname:     applicationData["app_name"].(string),
		ServiceName: serviceData["service_name"].(string),
		Domain:      "meddler.xyz",
	}

	token := common.GenerateAuthToken(response.Appname + "-" + response.ServiceName + "." + response.Domain)
	response.Token = token
	response.Endpoint = response.Appname + "-" + response.ServiceName + "." + response.Domain

	log.Println("token", token)
	return response, nil

}
