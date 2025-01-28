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

  //signals
  userTrackingOptions = signal<string[]>(['Shower', 'Run', 'Read']); 
  showModal = signal<boolean>(false);

  
}
