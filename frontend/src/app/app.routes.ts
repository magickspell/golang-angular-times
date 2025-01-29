import { Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import { AuthComponent } from './auth/auth.component';
import { SchedulerComponent } from './scheduler/scheduler.component';

export const routes: Routes = [
  { path: '', component: MainComponent },
  { path: 'scheduler', component: SchedulerComponent },
  { path: 'login', component: AuthComponent },
];
