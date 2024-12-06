import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-track',
  imports: [CommonModule],
  templateUrl: './track.component.html',
  styleUrl: './track.component.css'
})
export class TrackComponent {
  activities = [
    'Track',
    'Report',
    'Wishlist',
    'Goals'
  ];
  disabled = true
  name = 'b'
  person = {
    name: 'besc',
    age: 1,
    avatar: "https://w3schools.com/howto/img_avatar.png",
  };

  btnHandler() {
    alert("pepe")
  }
}
