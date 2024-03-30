import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-websocket-listener',
  templateUrl: 'websocket-listener.component.html',
  styleUrls: ['./websocket-listener.component.css'],
})
export class WebsocketListenerComponent implements OnInit {
  // WebSocket object
  ws!: WebSocket;
  messages: string[] = [];

  constructor() {}

  ngOnInit(): void {
    // Connect to WebSocket server
    this.ws = new WebSocket('ws://localhost:5000/ws');

    // Define WebSocket event listeners
    this.ws.onopen = (event) => {
      console.log('WebSocket connection opened');
    };

    this.ws.onmessage = (event) => {
      console.log('Message received:', event.data);
      this.messages.push(event.data);
    };

    this.ws.onerror = (event) => {
      console.error('WebSocket error:', event);
    };

    this.ws.onclose = (event) => {
      console.log('WebSocket connection closed');
    };
  }

  // Method to send message over WebSocket
  sendMessage(meg?: string): void {
    let ele: any = document.getElementById('takeText');
    let message: string = ele ? ele.value : 'dummy message';
    if (this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(message);
    } else {
      console.error('WebSocket connection is not open');
    }
  }

  // Method to disconnect WebSocket
  disconnect(): void {
    this.ws.close();
  }
}
