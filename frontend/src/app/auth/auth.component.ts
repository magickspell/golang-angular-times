import { Component } from '@angular/core';
import { AuthService } from './auth.service';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css'],
  imports: [ReactiveFormsModule]
})
export class AuthComponent {
  profileForm = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  constructor(public authService: AuthService) {
    const token: string | null = window.localStorage.getItem("auth-token");
    console.log(`token: ${token}`);
    // console.log(`isLoggedIn: ${this.authService.isLoggedIn().subscribe(val => val.valueOf())}`);
    console.log(`isLoggedIn: ${this.authService}`);
  }

  login() {
    if (!this.profileForm.value.email || !this.profileForm.value.password) {
      window.alert("empty fields!");
      return;
    }

    this.authService.login(this.profileForm.value.email, this.profileForm.value.password).subscribe((response) => {
      window.localStorage.setItem("auth-token", response.token);
      this.authService.setLoggedIn(true);
    });
  }

  logout() {
    this.authService.logout();
    window.localStorage.removeItem("auth-token");
  }



  handleSubmit() {
    alert(this.profileForm.value.email + ' | ' + this.profileForm.value.email);
  }
}