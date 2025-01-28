import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
@Component({
  selector: 'app-add-tracker',
  imports: [CommonModule],
  templateUrl: './add-tracker.component.html',
  standalone: true,
  styleUrl: './add-tracker.component.css'
})
export class AddTrackerComponent {

  @Input({ required: true }) inputAvailableOptions: string[] = [];
  @Output() close = new EventEmitter<void>();

  closeModal() {
    this.close.emit(); 
  }


}
