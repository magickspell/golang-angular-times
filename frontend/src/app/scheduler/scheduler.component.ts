import { Component, OnInit } from '@angular/core';
import { ScheduleService } from './schedule.service';
import { AuthComponent } from '../auth/auth.component';
import { AuthService } from '../auth/auth.service';
import { FormsModule } from '@angular/forms';

export type ScheduleDTO = {
  Day: string,
  Start: string,
  End: string
}

@Component({
  selector: 'app-scheduler',
  templateUrl: './scheduler.component.html',
  styleUrls: ['./scheduler.component.scss'],
  imports: [AuthComponent, FormsModule]
})
export class SchedulerComponent implements OnInit {
  schedules: ScheduleDTO[] = [];
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