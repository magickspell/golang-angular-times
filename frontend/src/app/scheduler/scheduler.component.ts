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
    let isErr: boolean = false;
    for (const schedule of this.schedules) {
      console.log(`[schedule]`)
      console.log(schedule)
      if(!this.checkTime(schedule.Start, schedule.End)) {
        isErr = true;
        break;
      }
    }
    if (isErr) {
      window.alert("Неправильно заполнен график");
      return;
    }

    // todo проверить что бросается ошибка в случае неправильного адрес например
    this.scheduleService.updateSchedules(this.schedules).subscribe(() => {
      alert('График обновлен');
    });
  }

  private checkTime(start: string, end: string) {
    if (!start.includes(":") || !end.includes(":")) {
      return false;
    }

    const startArr: string[] = start.split(":");
    const endArr: string[] = end.split(":");
    const isNumber = /^[0-9]+$/;
    for (const t of [...startArr, ...endArr]) {
      if (t.length > 2 || t.length === 0 || !isNumber.test(t)) {
        return false;
      }
    }

    const hoursStart: number = Number(startArr[0]);
    const minutesStart: number = Number(startArr[1]);
    const hoursEnd: number = Number(endArr[0]);
    const minutesEnd: number = Number(endArr[1]);
    if (hoursStart > 24 || hoursStart < 0 || hoursEnd > 24 || hoursEnd < 0) {
      return false;
    }
    if (minutesStart > 59 || minutesStart < 0 || minutesEnd > 59 || minutesEnd < 0) {
      return false;
    }

    if (hoursStart > hoursEnd || (hoursStart === hoursEnd && minutesStart >= minutesEnd)) {
      return false;
    }

    return true;
  }
}