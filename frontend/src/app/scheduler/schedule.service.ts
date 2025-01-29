import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ScheduleService {

  constructor(private http: HttpClient) {}

  getSchedules() {
    return this.http.get('http://localhost:8080/schedules');
  }

  updateSchedules(schedules: any) {
    return this.http.post('http://localhost:8080/schedules', schedules);
  }
}