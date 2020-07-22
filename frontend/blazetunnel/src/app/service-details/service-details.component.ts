import { Component, OnInit, Input } from '@angular/core';
import { AuthService } from '../shared/services/auth.service';
import { FirebaseServiceService } from '../firebase-service.service';
import { templateJitUrl, ThrowStmt } from '@angular/compiler';

@Component({
  selector: 'app-service-details',
  templateUrl: './service-details.component.html',
  styleUrls: ['./service-details.component.scss']
})
export class ServiceDetailsComponent implements OnInit {

  @Input('code') code = ''
  @Input('app_id') app_id = ''
  @Input('service_id') service_id = ''

  token = undefined

  docker_code = [
    
    `
    blazetunnel_side_car:
        image: golang
        environment: 
            -  token=`,`
            -  tunnel=blazetunnel.meddler.xyz
        command: ./blazetunnel client --local {{mockserver:8000}} -i 3600
        working_dir: /go/src/github.com/rounak316/blazetunnel
        read_only: true
        depends_on:
            - mockserver
    `

  ]

  constructor(private fbService: FirebaseServiceService) { }

  ngOnInit(): void {




    this.fbService.GetAuthToken(this.app_id, this.service_id).then(d => {
      this.token = d['auth_token']


    }).catch()
  }

}
