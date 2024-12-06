import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';

@Component({
  selector: 'app-home',
  imports: [CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
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
