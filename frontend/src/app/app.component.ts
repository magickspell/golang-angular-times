import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { AuthService } from './auth/auth.service';
import { AuthComponent } from './auth/auth.component';
import { SchedulerComponent } from './scheduler/scheduler.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, AuthComponent, SchedulerComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'frontend';

  constructor(public authService: AuthService) {}
}