import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { catchError, tap, throwError } from 'rxjs';
import { ScheduleDTO } from './scheduler.component';

@Injectable({
  providedIn: 'root'
})
export class ScheduleService {

  constructor(private http: HttpClient) {}

  getSchedules() {
    return this.http.get('http://localhost:8080/schedules');
  }

  updateSchedules(schedules: ScheduleDTO[]) {
    console.log(`[schedules]`);
    console.log(schedules);
    return this.http.post('http://localhost:8080/schedules', schedules);
  }

  checkTime(hour: number, minute: number, day: string): Observable<{allowed?: boolean}>  {
    let hourStr: string = '';
    let minuteStr: string = '';

    if (hour <= 9)  {
      hourStr = `0${hour}`
    } else {
      hourStr = `${hour}`
    }
    if (minute <= 9)  {
      minuteStr = `0${minute}`
    } else {
      minuteStr = `${minute}`
    }

    const request: string = JSON.stringify({
      hour: hourStr,
      minute: minuteStr,
      day: day
    })

    // todo решить что оставлять
    // return this.http.post('http://localhost:8080/check1', request) as unknown as Observable<{allowed?: boolean}>;
    return this.http.post<{allowed?: boolean}>('http://localhost:8080/check', request).pipe(
      tap(response => {
        console.log('Response:', response);
      }),
      catchError(error => {
        console.warn('Error:', error);
        window.alert("Произошла ошибка при выполнении запроса");
        return throwError(error);
      })
    );
  }
}