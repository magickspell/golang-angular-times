import { Component, OnInit } from '@angular/core';
import { ScheduleService } from './schedule.service';

@Component({
  selector: 'app-scheduler',
  templateUrl: './scheduler.component.html',
  styleUrls: ['./scheduler.component.css']
})
export class SchedulerComponent implements OnInit {
  schedules: any[] = [];

  constructor(private scheduleService: ScheduleService) {}

  ngOnInit() {
    this.scheduleService.getSchedules().subscribe((data: any) => {
      this.schedules = data;
    });
  }

  saveSchedules() {
    this.scheduleService.updateSchedules(this.schedules).subscribe(() => {
      alert('График обновлен');
    });
  }
}