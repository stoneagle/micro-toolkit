import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap  } from 'rxjs/operators';
import { Autobuild } from '../../model/storybox/autobuild';
import { Response } from '../../model/app.response.model';
import { AppConfig } from '../app.config';

const httpOptions = {
  headers: new HttpHeaders({ 
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*'
  })
};

@Injectable()
export class AutobuildService {
  private autobuildUrl = '/autobuild';

  constructor(
    private http: HttpClient,
  ) { }

  execMqtt(id: number): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${id}/mqtt`, httpOptions).pipe(
      tap(pools => this.log(`exec mqtt`)),
      catchError(this.handleError<Response>('execMqtt'))
    );
  }

  execUpgrade(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/upgrade`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(pools => this.log(`exec upgrade`)),
      catchError(this.handleError<Response>('execUpgrade'))
    );
  }

  execAlbum(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/album`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(pools => this.log(`exec album`)),
      catchError(this.handleError<Response>('execAlbum'))
    );
  }

  execCallback(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/callback`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(pools => this.log(`exec callback`)),
      catchError(this.handleError<Response>('execCallback'))
    );
  }

  delete(id: number): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${id}`, httpOptions).pipe(
      tap(pools => this.log(`delete autobuild`)),
      catchError(this.handleError<Response>('deleteAutobuild'))
    );
  }

  getList(): Observable<Autobuild[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl, httpOptions).pipe(
      tap(pools => this.log(`get autobuilds list`)),
      catchError(this.handleError<Response>('getAutobuilds')),
      map(res => {
        let abs:Autobuild[] = []; 
        res.data.map(
          one => {
            abs.push(new Autobuild(one));
          }
        )
        return abs; 
      })
    );
  }

  add(autobuild: Autobuild): Observable<Autobuild> {
      return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl, JSON.stringify(autobuild), httpOptions).pipe(
      tap((res: Response) => this.log(`added autobuild w/ id=${res.data.id}`)),
      catchError(this.handleError<Response>('addAutobuild')),
      map(res => new Autobuild(res.data))
    );
  }

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error);
      this.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    }
  }

  private log(message: string) {
    // this.messageService.add('PoolService:' + message);
  }
}
