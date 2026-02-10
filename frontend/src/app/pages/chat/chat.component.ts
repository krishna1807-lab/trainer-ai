import { Component, ElementRef, ViewChild } from '@angular/core';
import { ChatService } from '../../service/chat.service';
import { DocumentService } from '../../service/document.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html'
})
export class ChatComponent {

  @ViewChild('scrollContainer') scrollContainer!: ElementRef;

  userInput = '';
  messages: any[] = [];

  constructor(
    private chatService: ChatService,
    private docService: DocumentService
  ) {}

  getCurrentTime() {
    return new Date().toLocaleTimeString([], {
      hour: '2-digit',
      minute: '2-digit'
    });
  }

  sendMessage() {

    console.log("SEND CLICKED");

    if (!this.userInput.trim()) return;

    const userMsg = {
      role: 'user',
      content: this.userInput,
      time: this.getCurrentTime()
    };

    this.messages.push(userMsg);

    const prompt = this.userInput;
    this.userInput = '';

    this.scrollToBottom();

    this.chatService.sendMessage(prompt).subscribe({
      next: (res: any) => {

        console.log("API RESPONSE:", res);

        const aiMsg = {
          role: 'assistant',
          // ⭐ FIXED HERE (Matches Your Go Backend)
          content: res?.response || "No AI response",
          time: this.getCurrentTime()
        };

        this.messages.push(aiMsg);
        this.scrollToBottom();

      },
      error: (err) => {

        console.error("CHAT ERROR:", err);

        this.messages.push({
          role: 'assistant',
          content: 'Something went wrong.',
          time: this.getCurrentTime()
        });

      }
    });

  }

  onFileSelected(event: any) {

    const file = event.target.files[0];
    if (!file) return;

    this.docService.uploadPDF(file).subscribe({
      next: () => {

        this.messages.push({
          role: 'assistant',
          content: 'Document uploaded successfully.',
          time: this.getCurrentTime()
        });

        this.scrollToBottom();

      },
      error: (err) => {
        console.error("UPLOAD ERROR:", err);
      }
    });

  }

  scrollToBottom() {
    setTimeout(() => {
      if (this.scrollContainer) {
        this.scrollContainer.nativeElement.scrollTop =
          this.scrollContainer.nativeElement.scrollHeight;
      }
    }, 100);
  }

}
