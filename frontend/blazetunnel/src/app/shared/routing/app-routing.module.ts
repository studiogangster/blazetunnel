import { NgModule } from '@angular/core';
// Required services for navigation
import { Routes, RouterModule } from '@angular/router';
import { SigninComponent } from 'src/app/signin/signin.component';
import { SignupComponent } from 'src/app/signup/signup.component';
import { ForgotpasswordComponent } from 'src/app/forgotpassword/forgotpassword.component';
import { VerifyemailComponent } from 'src/app/verifyemail/verifyemail.component';
import { DashboardComponent } from 'src/app/dashboard/dashboard.component';
import { AuthGuard } from '../guard/auth.guard';
import { AppCreateComponent } from 'src/app/app-create/app-create.component';


// Import all the components for which navigation service has to be activated 

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full'},
  { path: 'application', component: AppCreateComponent},
  { path: 'sign-in', component: SigninComponent},

  
  { path: 'register-user', component: SignupComponent},
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: 'forgot-password', component: ForgotpasswordComponent },
  { path: 'verify-email-address', component: VerifyemailComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }