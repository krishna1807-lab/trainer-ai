import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  // ⭐ Change when deploying
  private baseUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  post<T>(endpoint: string, body: any) {
    return this.http.post<T>(`${this.baseUrl}${endpoint}`, body);
  }

  upload<T>(endpoint: string, formData: FormData) {
    return this.http.post<T>(`${this.baseUrl}${endpoint}`, formData);
  }

}
