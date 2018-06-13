import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap  } from 'rxjs/operators';
import { DeviceScan } from '../../model/storybox/scan'
import { Response } from '../../model/app.response.model';
import { AppConfig } from '../app.config';
import { MessageHandlerService  } from '../base/message-handler.service';

const httpOptions = {
  headers: new HttpHeaders({ 
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*'
  })
};

@Injectable()
export class ScanService {
  private uri = '/scan';

  constructor(
    private http: HttpClient,
    private messageHandlerService: MessageHandlerService 
  ) { }

  delete(id: number): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.uri + `/${id}`, httpOptions).pipe(
      tap(res => this.log('PROCESS.DELETE', res)),
      catchError(this.handleError<Response>('deleteScan'))
    );
  }

  list(appId: string): Observable<DeviceScan[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.uri + `/${appId}`, httpOptions).pipe(
      catchError(this.handleError<Response>('getScans')),
      map(res => {
        let ret:DeviceScan[] = []; 
        if (res) {
          res.data.map(
            one => {
              ret.push(new DeviceScan(one));
            }
          )
        }
        return ret; 
      }),
    );
  }

  add(scan: DeviceScan): Observable<DeviceScan> {
      return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.uri, JSON.stringify(scan), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('addScan')),
      map(res => new DeviceScan(res.data))
    );
  }

  update(scan: DeviceScan): Observable<DeviceScan> {
      return this.http.put<Response>(AppConfig.settings.apiServer.endpoint + this.uri, JSON.stringify(scan), httpOptions).pipe(
      tap(res => this.log('PROCESS.UPDATE', res)),
      catchError(this.handleError<Response>('updateScan')),
      map(res => new DeviceScan(res.data))
    );
  }

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      this.messageHandlerService.handleError(error);
      return of(result as T);
    }
  }

  private log (message: string, res: Response) {
    if (res.code != 0) {
      this.messageHandlerService.showWarning(res.desc)
    } else {
      this.messageHandlerService.showSuccess(message)
    }
  }
}
