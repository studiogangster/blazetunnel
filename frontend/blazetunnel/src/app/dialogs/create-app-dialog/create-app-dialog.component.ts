import { Component, OnInit, Inject } from '@angular/core';
import { FirebaseServiceService } from 'src/app/firebase-service.service';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-create-app-dialog',
  templateUrl: './create-app-dialog.component.html',
  styleUrls: ['./create-app-dialog.component.scss']
})
export class CreateAppDialogComponent implements OnInit {


  constructor(
    private fbService: FirebaseServiceService,
    public dialogRef: MatDialogRef<CreateAppDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any



  ) { }


  ngOnInit(): void {
  }



  create_app(app_name) {

    this.fbService.createApp(app_name)
      .then(_ => {

        this.dialogRef.close(true);

      })
      .catch(_ => {
        console.log('boomer')
        console.log(_)
      })
  }
}
