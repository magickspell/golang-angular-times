import { Component } from '@angular/core';
import { AUTH_TOKEN, AuthService } from './auth.service';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss'],
  imports: [ReactiveFormsModule]
})
export class AuthComponent {
  profileForm = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  constructor(public authService: AuthService) {}

  login() {
    if (!this.profileForm.value.email || !this.profileForm.value.password) {
      window.alert("Введите логин и пароль");
      return;
    }

    this.authService.login(this.profileForm.value.email, this.profileForm.value.password).subscribe((response) => {
      window.localStorage.setItem(AUTH_TOKEN, response.token);
      this.authService.setLoggedIn(true);
    });
  }

  logout() {
    this.authService.logout();
    window.localStorage.removeItem(AUTH_TOKEN);
  }



  handleSubmit() {
    alert(this.profileForm.value.email + ' | ' + this.profileForm.value.email);
  }
}