import { Component, signal } from '@angular/core';
import { AddTrackerComponent } from './modals/add-tracker/add-tracker.component';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-track',
  imports: [CommonModule, AddTrackerComponent],
  templateUrl: './track.component.html',
  styleUrl: './track.component.css'
})
export class TrackComponent {
   //Signals 
  ////TODO: Implement a request with the tracking options used by a user to assign it to this signal, for the moment it is assigned a dummy value.
  userTrackingOptions = signal<string[]>([]);
  showModal = signal<boolean>(false);
}
