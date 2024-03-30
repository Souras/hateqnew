import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { WebsocketListenerComponent } from './components/websocket-listener.component';
import { RouterModule, Routes } from '@angular/router';

const appRoutes: Routes = [
  { path: 'ws', component: WebsocketListenerComponent },
];


@NgModule({
  declarations: [
    AppComponent,
    WebsocketListenerComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule.forRoot(
      appRoutes
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }


