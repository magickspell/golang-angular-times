import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError, Observable, throwError } from 'rxjs';

export const AUTH_TOKEN = "auth-token";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private loggedIn = false;

  constructor(private http: HttpClient) {
    const token: string | null = window.localStorage.getItem("auth-token");
    this.loggedIn = token ? true : false;
  }

  isLoggedIn() {
    return this.loggedIn;
  }

  setLoggedIn(value: boolean) {
    this.loggedIn = value;
  }

  login(email: string, password: string): Observable<{token: string}> {
    return this.http.post('http://localhost:8080/login', { email, password }).pipe(
      catchError(error => {
        console.warn('Error:', error);
        window.alert("Ошибка авторизации");
        return throwError(error);
      })
    ) as unknown as Observable<{token: string}>;
  }

  logout() {
    this.loggedIn = false;
  }
}