import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { ChatRequest, GroqResponse } from '../models/chat.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  constructor(private api: ApiService) {}

  sendMessage(prompt: string): Observable<GroqResponse> {

    const body: ChatRequest = {
      prompt
    };

    return this.api.post<GroqResponse>('/chat', body);
  }

}
