import { Component, OnInit, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FirebaseServiceService } from 'src/app/firebase-service.service';

@Component({
  selector: 'app-create-service-dialog',
  templateUrl: './create-service-dialog.component.html',
  styleUrls: ['./create-service-dialog.component.scss']
})
export class CreateServiceDialogComponent implements OnInit {



  constructor(
    private fbService: FirebaseServiceService,
    public dialogRef: MatDialogRef<CreateServiceDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any) { }

  ngOnInit(): void {

  }

  onNoClick(): void {
    this.dialogRef.close();
  }



  create_service(app_id, service_name) {
    this.fbService.createService(app_id, service_name).then(result => {
      console.log('closing')
      this.dialogRef.close( true );

    }).catch(_=>{
      console.log('closing')
      
    })
  }



}
