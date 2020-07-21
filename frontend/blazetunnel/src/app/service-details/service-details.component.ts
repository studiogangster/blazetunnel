import { Component, OnInit, Input } from '@angular/core';
import { AuthService } from '../shared/services/auth.service';
import { FirebaseServiceService } from '../firebase-service.service';

@Component({
  selector: 'app-service-details',
  templateUrl: './service-details.component.html',
  styleUrls: ['./service-details.component.scss']
})
export class ServiceDetailsComponent implements OnInit {

  @Input('code') code = ''
  @Input('app_id') app_id = ''
  @Input('service_id') service_id = ''


  constructor(private fbService: FirebaseServiceService) { }

  ngOnInit(): void {

    this.fbService.GetAuthToken(this.app_id, this.service_id).then()
  }

}
