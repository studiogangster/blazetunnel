import { Component, OnInit, Inject } from '@angular/core';
import { FirebaseServiceService } from 'src/app/firebase-service.service';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-service-details-dialog',
  templateUrl: './service-details-dialog.component.html',
  styleUrls: ['./service-details-dialog.component.scss']
})
export class ServiceDetailsDialogComponent implements OnInit {


  code = `
  version: "3"
  services:
        nginx:
          image: nginx:1.15-alpine
          restart: unless-stopped
          volumes:
              - ./nginx/data/nginx/app.conf:/etc/nginx/conf.d/app.conf
              - ./nginx/data/nginx/nginx.conf:/etc/nginx/nginx.conf
              - ./nginx/data/html:/usr/share/nginx/html
              - ./nginx/data/certbot/conf:/etc/letsencrypt
              - ./nginx/data/certbot/www:/var/www/certbot
          depends_on:
              - blazeserver
          ports:
              - "443:443"
              - "80:80"

  `
  constructor(
    private fbService: FirebaseServiceService,
    public dialogRef: MatDialogRef<ServiceDetailsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any) { }

  ngOnInit(): void {

  }

  onNoClick(): void {
    this.dialogRef.close();
  }



  create_service(app_id, service_name) {
    this.fbService.createService(app_id, service_name).then(result => {
      console.log('closing')
      this.dialogRef.close(true);

    }).catch(_ => {
      console.log('closing')

    })
  }



}
