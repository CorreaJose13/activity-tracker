
import { Component } from '@angular/core';

@Component({
  selector: 'app-home',
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  trackHandler() {
    window.location.pathname = '/track';
  }

  reportHandler() {
    window.location.pathname = '/report';
  }

  wishlistHandler() {
    window.location.pathname = '/wishlist';
  }
}
