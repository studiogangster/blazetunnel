import { Component, OnInit, Input } from '@angular/core';
import { AuthService } from '../shared/services/auth.service';
import { FirebaseServiceService } from '../firebase-service.service';
import { templateJitUrl, ThrowStmt } from '@angular/compiler';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-service-details',
  templateUrl: './service-details.component.html',
  styleUrls: ['./service-details.component.scss']
})
export class ServiceDetailsComponent implements OnInit {

  @Input('code') code = ''
  @Input('app_id') app_id = ''
  @Input('service_id') service_id = ''

  @Input('app_name') app_name = ''
  @Input('service_name') service_name = ''

  token = undefined

  docker_code = [

    `
    version: '3'
    services:

      blazetunnel_side_car:
          image: golang
          environment: 
              -  token=`, `
              -  tunnel=blazetunnel.meddler.xyz
          command: ./blazetunnel client --local {{mockserver:8000}} -i 3600
          working_dir: /go/src/github.com/rounak316/blazetunnel
          read_only: true
          depends_on:
              - mockserver
    `

  ]

  binary_code = [`./blazetunnel client --tunnel `, `-`, `.blazetunnel.meddler.xyz --local localhost:4200 -i 3600`, ' --token ']

  constructor(private fbService: FirebaseServiceService , private _snackBar: MatSnackBar) { }

  ngOnInit(): void {


    this.fbService.GetAuthToken(this.app_id, this.service_id).then(d => {
      this.token = d['auth_token']


    }).catch(err => {
      this.token = "default_token"
    })
  }

  copyMessage(val: string) {
    const selBox = document.createElement('textarea');
    selBox.style.position = 'fixed';
    selBox.style.left = '0';
    selBox.style.top = '0';
    selBox.style.opacity = '0';
    selBox.value = val;
    document.body.appendChild(selBox);
    selBox.focus();
    selBox.select();
    document.execCommand('copy');
    document.body.removeChild(selBox);

    this._snackBar.open("Copied to clipboard", 'Okay', {
      duration: 2000,
    });
  }

}
