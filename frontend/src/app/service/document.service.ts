import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DocumentService {

  constructor(private api: ApiService) {}

  uploadPDF(file: File): Observable<any> {

    const formData = new FormData();
    formData.append('file', file);

    return this.api.upload('/upload-doc', formData);
  }

}
