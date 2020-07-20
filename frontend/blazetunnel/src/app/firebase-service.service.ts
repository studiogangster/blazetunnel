import { Injectable } from '@angular/core';
import { AngularFirestore } from '@angular/fire/firestore';
import { AuthService } from './shared/services/auth.service';
import { AngularFireAuth } from '@angular/fire/auth';

@Injectable({
  providedIn: 'root'
})
export class FirebaseServiceService {

  constructor(private firestore: AngularFirestore, private authService: AngularFireAuth, private _auth: AuthService) { }


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
        .collection("app")
        .add({ uid: user.uid, app_name: app_name })
    })
  }

  createService(app_id, service_name) {

    return this.authService.currentUser.then(user => {

      return new Promise<any>((resolve, reject) => {
        this.firestore
          .collection("app")
          .doc(app_id)
          .collection('service')
          .add({ service_name: service_name })
      });
    })
  }

  getApps(uid) {
    return this.firestore.collection("app"
      , ref => { return ref.where("uid", '==', uid) }
    )
      .get()
  }

  getServices(app_id) {
    return this.firestore
      .collection("app")
      .doc(app_id)
      .collection('service')
      .get()

  }

  deleteApp(id) {
    return this.firestore.collection('app').doc(id).delete()
  }

  enableService(app_id, service_id, enabled) {
    return this.firestore
      .collection("app")
      .doc(app_id)
      .collection('service')
      .doc(service_id)
      .update(
        { enabled: enabled }
      )
  }


  deleteService(app_id, service_id) {
    return this.firestore.collection('app')
      .doc(app_id)
      .collection('service')
      .doc(service_id)
      .delete()
  }
}
