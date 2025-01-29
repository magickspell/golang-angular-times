import { Component, OnInit } from '@angular/core';
import { ScheduleService } from '../scheduler/schedule.service';
import { AuthService } from '../auth/auth.service';
import { ScheduleDTO } from '../scheduler/scheduler.component';
import { FormsModule } from '@angular/forms';
import { Observable } from 'rxjs/internal/Observable';

const SELECT_DAY: string = "выберите день";

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss'],
  imports: [FormsModule]
})
export class MainComponent implements OnInit {
  schedules: ScheduleDTO[] = [];
  hours: number = 0;
  minutes: number = 0;
  day: string = SELECT_DAY;

  constructor(private scheduleService: ScheduleService, public authService: AuthService) {}

  ngOnInit() {
    this.scheduleService.getSchedules().subscribe((data: any) => {
      this.schedules = data;
    });
  }

  checkTime(hour: number, minute: number, day: string) {
    if (day === SELECT_DAY) {
      window.alert("Не выбран день");
      return;
    }

    this.scheduleService.checkTime(hour, minute, day).subscribe((res) => {
      if (res?.allowed) {
        window.alert("Разрешено")
      } else {
        window.alert("Запрещено")
      }
    });
  }
}