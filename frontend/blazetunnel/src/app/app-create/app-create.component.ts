import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
import { FirebaseServiceService } from '../firebase-service.service';
import { GetAppComponent } from '../get-app/get-app.component';

@Component({
  selector: 'app-create',
  templateUrl: './app-create.component.html',
  styleUrls: ['./app-create.component.scss']
})
export class AppCreateComponent implements OnInit {

  @ViewChild('applicationList') applicationList: GetAppComponent;
  value = '';

  constructor(private fbService: FirebaseServiceService) { }

  ngOnInit(): void {
  }

  createApp(app_name) {
    this.applicationList._resetApplicationList()

    this.fbService.createApp(app_name)
      .then(_ => {

        this.value = ''
        console.log('boomer')
        this.applicationList.refreshApplications()

      })
      .catch(_ => {
        console.log('boomer')
        console.log(_)
      })
  }
}
