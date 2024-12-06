import { Routes } from '@angular/router';

import { HomeComponent } from './pages/home/home.component';
import { TrackComponent } from './pages/track/track.component';
export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'track',
    component: TrackComponent,
  }
];
