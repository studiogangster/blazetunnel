import { Component } from '@angular/core';
import { FirebaseServiceService } from './firebase-service.service';

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
export class AppComponent {
  title = 'blazetunnel';
  showFiller = false;

  folders: Section[] = [
    {
      name: 'Dashboard',
      description: new Date('1/1/16'),
      route: ['/dashboard'],
    },

    {
      name: 'Applications',
      description: new Date('1/17/16'),
      route: ['/application'],

    },

    {
      name: 'Sign In',
      description: new Date('1/17/16'),
      route: ['/sign-in'],

    },

    
    {
      name: 'Sign Out',
      description: new Date('1/28/16'),
      route: ['/sign-out'],

    }
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

  constructor(private firebaseService: FirebaseServiceService) {

  }

  test(){
    this.firebaseService.registerDomain( { "test" : 123 } ).then( d=>{console.log(d)} ).catch(d=>{console.log(d)})
  }

}
