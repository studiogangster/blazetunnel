import { Component, OnInit } from '@angular/core';
import { FirebaseServiceService } from './firebase-service.service';
import { MatDialog } from '@angular/material/dialog';
import { AuthService } from './shared/services/auth.service';
import { filter } from 'rxjs/operators';

export interface Section {
  name: string;
  route: string[];
  description: Date;
}



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'blazetunnel';
  showFiller = false;
  applications = undefined;

  folders: Section[] = [

    {
      name: 'Applications',
      description: new Date('1/17/16'),
      route: ['/application'],

    },

    {
      name: 'Profile',
      description: new Date('1/1/16'),
      route: ['/dashboard'],
    },


    {
      name: 'Sign In',
      description: new Date('1/17/16'),
      route: ['/sign-in'],

    },



  ];


  notes: Section[] = [
    {
      name: 'Vacation Itinerary',
      description: new Date('2/20/16'),
      route: ['/dashboard'],

    },
    {
      name: 'Kitchen Remodel',
      description: new Date('1/18/16'),
      route: ['/dashboard'],

    }
  ];



  constructor(
    private fbService: FirebaseServiceService, public authService: AuthService,

    public dialog: MatDialog
  ) {

    this.fbService.getApps().then()

  }
  ngOnInit(): void {

    this.refreshApplications()
  }

  refreshApplications() {

    this.authService.userDataSubject.pipe(filter(user => { return user })).subscribe(user => {


      this.applications = undefined;

      this.fbService.getApps().then(data => {

        this.applications = []

        data.forEach(d => {
          let data = d.data()
          data.id = d.id
          this.applications.push(data)

        })

      })
    })
  }

  test() {
    this.fbService.registerDomain({ "test": 123 }).then(d => { console.log(d) }).catch(d => { console.log(d) })
  }

}
