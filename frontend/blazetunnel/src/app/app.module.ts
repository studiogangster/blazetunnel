import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatGridListModule} from '@angular/material/grid-list';
import { environment } from "src/environments/environment";
import { AngularFireModule } from "@angular/fire";
import { AngularFirestoreModule } from "@angular/fire/firestore";
import { AuthService } from './shared/services/auth.service';
import { SigninComponent } from './signin/signin.component';
import { SignupComponent } from './signup/signup.component';
import { ForgotpasswordComponent } from './forgotpassword/forgotpassword.component';
import { VerifyemailComponent } from './verifyemail/verifyemail.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AppRoutingModule } from './shared/routing/app-routing.module';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatButtonModule} from '@angular/material/button';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatIconModule} from '@angular/material/icon';
import {MatListModule} from '@angular/material/list';
import {MatInputModule} from '@angular/material/input';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatChipsModule} from '@angular/material/chips';
import {MatCardModule} from '@angular/material/card';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { AppCreateComponent } from './app-create/app-create.component';
import { GetAppComponent } from './get-app/get-app.component';
import {MatSelectModule} from '@angular/material/select';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatTreeModule} from '@angular/material/tree';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import { FlexLayoutModule } from '@angular/flex-layout';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatMenuModule} from '@angular/material/menu';
import {MatDialogModule} from '@angular/material/dialog';
import { CreateServiceDialogComponent } from './dialogs/create-service-dialog/create-service-dialog.component';
import {MatDividerModule} from '@angular/material/divider';
import { CreateAppDialogComponent } from './dialogs/create-app-dialog/create-app-dialog.component';
import { ServiceDetailsDialogComponent } from './dialogs/service-details-dialog/service-details-dialog.component';
import {MatTabsModule} from '@angular/material/tabs';
import { HighlightModule } from 'ngx-highlightjs';
import { ServiceDetailsComponent } from './service-details/service-details.component';
import { HttpClientModule } from '@angular/common/http';
import { SafePipe } from './safe.pipe';
import { SpecialCharacterDirective } from './special-character.directive';

@NgModule({
  declarations: [
    AppComponent,
    SigninComponent,
    SignupComponent,
    ForgotpasswordComponent,
    VerifyemailComponent,
    DashboardComponent,
    AppCreateComponent,
    GetAppComponent,
    CreateServiceDialogComponent,
    CreateAppDialogComponent,
    ServiceDetailsDialogComponent,
    ServiceDetailsComponent,
    SafePipe,
    SpecialCharacterDirective
  ],
  imports: [
    HttpClientModule,
    HighlightModule,
    MatTabsModule,
    MatDividerModule,
    MatDialogModule,
    MatMenuModule,
    MatProgressSpinnerModule,
    FlexLayoutModule,
    MatSlideToggleModule,
    MatTreeModule,
    MatExpansionModule,
    MatSelectModule,
    MatGridListModule,
    FormsModule,
    ReactiveFormsModule,
    MatCardModule,
    MatChipsModule,
    MatAutocompleteModule,
    MatInputModule,
    MatListModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatSidenavModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    AngularFireModule.initializeApp(environment.firebaseConfig),
    AngularFirestoreModule
  ],
  providers:  [
    AuthService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
