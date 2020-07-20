import { Component, OnInit } from '@angular/core';
import { FirebaseServiceService } from '../firebase-service.service';
import { BehaviorSubject, Subject } from 'rxjs';
import { AuthService } from '../shared/services/auth.service';
import { AngularFireAuth } from '@angular/fire/auth';

import { filter } from 'rxjs/operators';
import { app } from 'firebase';
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'app-get',
  templateUrl: './get-app.component.html',
  styleUrls: ['./get-app.component.scss']
})
export class GetAppComponent implements OnInit {


  serviceList = new BehaviorSubject(null);
  serviceList$ = this.serviceList.asObservable()

  applications = undefined;
  constructor(private fbService: FirebaseServiceService, private authService: AuthService) { }


  refreshApplications() {

    this.authService.userDataSubject.pipe(filter(user => { return user })).subscribe(user => {

      this.applications = undefined;

      this.fbService.getApps(user.uid).subscribe(data => {

        this.applications = []

        data.forEach(d => {
          let data = d.data()
          data.id = d.id
          this.applications.push(data)

        })

      })
    })
  }

  ngOnInit(): void {
    this.refreshApplications()

  }


  delete(app_id) {
    this._resetApplicationList()

    this.fbService.deleteApp(app_id).then(result => {

      this.refreshApplications()
    })

  }

  getServices(app_id) {
    return this.fbService.getServices(app_id)
  }

  _resetApplicationList() {
    this.applications = null

  }
  _resetServiceList() {
    this.serviceList.next(null)

  }

  opened(app_id) {
    this._resetServiceList()
    this.getServices(app_id)
      .subscribe(data => {
        let _data = []
        data.forEach(_d => {
          let __d = _d.data()
          __d.id = _d.id
          _data.push(__d)
        })
        this.serviceList.next(_data)
      })

  }

  create_service(app_id, service_name) {
    this.fbService.createService(app_id, service_name).then(result => {
      this.opened(app_id)

    })
  }


  closed(app_id) {
    this._resetServiceList()
  }

  enableService(app_id, service_id, enabled) {
    this.fbService.enableService(app_id, service_id, enabled).then(result => {
      this.opened(app_id)
    })

  }

  deleteService(app_id, service_id) {
    this._resetServiceList()
    this.fbService.deleteService(app_id, service_id).then(result => {
      this.opened(app_id)
    })

  }
}
