import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap  } from 'rxjs/operators';
import { AppConfig } from '../app.config';
import { MessageHandlerService  } from '../base/message-handler.service';

@Injectable()
export class BaseService {
  currentUser: string = null;

  constructor(
    private http: HttpClient,
    private messageHandlerService: MessageHandlerService 
  ) { }

  clear(): void {
    this.currentUser = null;
  }

  getCurrentUser(): string {
    return this.currentUser;
  }

  login(username: string, password: string): Observable<Response> {
    let httpOptions = {
      headers: new HttpHeaders({ 
        'Authorization': "Basic " + btoa(`${username}:${password}`) , 
        'Content-Type': 'application/x-www-form-urlencoded',
        'Access-Control-Allow-Origin': '*'
      })
    };

    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + `/login`, httpOptions).pipe(
      catchError(this.handleError<Response>('login'))
    );
  }

  logout(): Observable<Response> {
    let httpOptions = {
      headers: new HttpHeaders({ 
        'Content-Type': 'application/x-www-form-urlencoded',
        'Access-Control-Allow-Origin': '*'
      })
    };

    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + `/logout`, httpOptions).pipe(
      catchError(this.handleError<Response>('logout'))
    );
  }

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      this.messageHandlerService.handleError(error);
      return of(result as T);
    }
  }
}
