import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LayoutComponent } from './layout/layout/layout.component';
import { ChatComponent } from './pages/chat/chat.component';
import { DocumentComponent } from './pages/document/document.component';
import { FlashcardComponent } from './pages/flashcard/flashcard.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      { path: 'chat', component: ChatComponent },
      { path: 'documents', component: DocumentComponent },
      { path: 'flashcards', component:FlashcardComponent },
      { path: 'profile', component:ProfileComponent },
      { path: 'dashboard', component:DashboardComponent },
      { path: '', redirectTo: 'chat', pathMatch: 'full' }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
