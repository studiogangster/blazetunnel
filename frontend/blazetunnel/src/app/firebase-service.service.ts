import { Injectable } from '@angular/core';
import { AngularFirestore } from '@angular/fire/firestore';
import { AuthService } from './shared/services/auth.service';
import { AngularFireAuth } from '@angular/fire/auth';
import { app } from 'firebase';
import { HttpClient } from '@angular/common/http';
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FirebaseServiceService {

  constructor(private http: HttpClient, private firestore: AngularFirestore, private authService: AngularFireAuth, private _auth: AuthService) { }

  // Get Auth Token 
  GetAuthToken(app_id: string, service_id: string) {

    return of(`
    {
      "status": false,
      "auth_token": "ThisIsAnAuthTokn"
  }
    `)

    // return this.authService.currentUser.then(user => {

    //   return user.getIdToken().then(token => {

    //     return this.http.post("http://localhost:90", { id_token: token, app_id: app_id, service_id: service_id }).subscribe()

    //   })
    // })
  }

  registerDomain(data) {
    return new Promise<any>((resolve, reject) => {
      this.firestore
        .collection("coffeeOrders")
        .add(data)
        .then(res => { }, err => reject(err));
    });
  }

  getDomains() {
    return this.firestore.collection("coffeeOrders").snapshotChanges();
  }

  createApp(app_name) {



    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .add({ uid: user.uid, app_name: app_name })
    })
  }

  createService(app_id, service_name) {

    return this.authService.currentUser.then(user => {

      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id)
        .collection('service')
        .add({ service_name: service_name })

    })
  }

  getApps() {

    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .get().toPromise()

    })


  }

  getServices(app_id) {


    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id)
        .collection('service')
        .get().toPromise()

    })


  }

  deleteApp(app_id) {

    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id).delete()

    })


  }

  enableApp(app_id, enabled) {

    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id)
        .update(
          { enabled: enabled }
        )

    })



  }


  enableService(app_id, service_id, enabled) {

    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id)
        .collection('service')
        .doc(service_id)
        .update(
          { enabled: enabled }
        )

    })



  }


  deleteService(app_id, service_id) {

    return this.authService.currentUser.then(user => {
      return this.firestore
        .collection("app").doc(user.uid)
        .collection("app")
        .doc(app_id)
        .collection('service')
        .doc(service_id)
        .delete()


    })

  }
}
