import { Component } from '@angular/core';

interface Activity {
  title: string;
  description: string;
  time: string;
  icon: string;
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html'
})
export class DashboardComponent {

  stats = [
    {
      title: 'Total Documents',
      value: 24,
      icon: 'document'
    },
    {
      title: 'Total Flashcards',
      value: 132,
      icon: 'flashcard'
    },
    {
      title: 'Total Passed',
      value: 18,
      icon: 'passed'
    }
  ];

  recentActivities: Activity[] = [
    {
      title: 'Uploaded Angular Guide.pdf',
      description: 'Document processed successfully',
      time: '2 hours ago',
      icon: 'upload'
    },
    {
      title: 'Flashcards Generated',
      description: 'AI generated 25 flashcards',
      time: '5 hours ago',
      icon: 'flash'
    },
    {
      title: 'Quiz Completed',
      description: 'Score: 82%',
      time: 'Yesterday',
      icon: 'quiz'
    },
    {
      title: 'New Document Uploaded',
      description: 'RAG Architecture Notes.pdf',
      time: '2 days ago',
      icon: 'upload'
    }
  ];
}
