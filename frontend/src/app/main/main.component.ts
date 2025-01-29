import { Component, OnInit } from '@angular/core';
import { ScheduleService } from '../scheduler/schedule.service';
import { AuthService } from '../auth/auth.service';

export type Schedule = {
  Day: string,
  Start: string,
  End: string
}

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  schedules: Schedule[] = [];
  isLoggedIn = false;

  constructor(private scheduleService: ScheduleService, public authService: AuthService) {}

  ngOnInit() {
    this.scheduleService.getSchedules().subscribe((data: any) => {
      this.schedules = data;
    });

    this.isLoggedIn = this.authService.isLoggedIn();
  }

  saveSchedules() {
    this.scheduleService.updateSchedules(this.schedules).subscribe(() => {
      alert('График обновлен');
    });
  }
}