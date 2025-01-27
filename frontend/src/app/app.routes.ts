import { Routes } from '@angular/router';

import { HomeComponent } from './pages/home/home.component';
import { TrackComponent } from './pages/track/track.component';
import { ReportComponent } from './pages/report/report.component';
import { WishlistComponent } from './pages/wishlist/wishlist.component';
import { ResourcesComponent } from './pages/resources/resources.component';
export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'track',
    component: TrackComponent,
  },
  {
    path: 'report',
    component: ReportComponent,
  },
  {
    path: 'wishlist',
    component: WishlistComponent
  },
  {
    path: 'resources',
    component: ResourcesComponent,
  }
];
