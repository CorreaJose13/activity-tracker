import { Component, ElementRef, HostListener, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-home',
  imports: [CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {

  @ViewChild('imageRef') imageRef!: ElementRef;
  @ViewChild('buttonContainer') buttonContainer!: ElementRef;

  userName: string | null = '';
  chatId: string | null = '';

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.userName = params['user'];
      this.chatId = params['chat_id'];
    });
  }

  trackHandler() {
    window.location.pathname = '/track';
  }

  reportHandler() {
    window.location.pathname = '/report';
  }

  wishlistHandler() {
    window.location.pathname = '/wishlist';
  }

  ngAfterViewInit() {
    this.adjustButtonWidth();
  }

  @HostListener('window:resize')
  adjustButtonWidth() {
    if (this.imageRef && this.buttonContainer) {
      this.buttonContainer.nativeElement.style.maxWidth = `${this.imageRef.nativeElement.clientWidth}px`;
    }
  }
}
