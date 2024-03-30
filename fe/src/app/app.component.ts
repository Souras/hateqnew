import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { WebsocketListenerComponent } from './components/websocket-listener.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'ditto';

  constructor(private router:Router){

  }

  navigateToWS(){
    this.router.navigate(['\ws'])
  }
}
