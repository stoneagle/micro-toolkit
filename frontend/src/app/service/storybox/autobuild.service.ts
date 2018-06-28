import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { Observable } from 'rxjs';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap  } from 'rxjs/operators';
import { Autobuild } from '../../model/storybox/autobuild';
import { Mqtt } from '../../model/storybox/mqtt';
import { Callback } from '../../model/storybox/callback';
import { Upgrade } from '../../model/storybox/upgrade';
import { Album } from '../../model/storybox/album'
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
export class AutobuildService {
  private autobuildUrl = '/autobuild';

  constructor(
    private http: HttpClient,
    private messageHandlerService: MessageHandlerService 
  ) { }

  rollbackMqtt(autobuild: Autobuild): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/mqtt`, httpOptions).pipe(
      tap(res => this.log('PROCESS.ROLLBACK', res)),
      catchError(this.handleError<Response>('rollbackMqtt'))
    );
  }

  rollbackUpgrade(autobuild: Autobuild): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/upgrade`, httpOptions).pipe(
      tap(res => this.log('PROCESS.ROLLBACK', res)),
      catchError(this.handleError<Response>('rollbackUpgrade'))
    );
  }

  rollbackAlbum(autobuild: Autobuild): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/album`, httpOptions).pipe(
      tap(res => this.log('PROCESS.ROLLBACK', res)),
      catchError(this.handleError<Response>('rollbackAlbum'))
    );
  }

  rollbackCallback(autobuild: Autobuild): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/callback`, httpOptions).pipe(
      tap(res => this.log('PROCESS.ROLLBACK', res)),
      catchError(this.handleError<Response>('rollbackCallback'))
    );
  }

  execMqtt(id: number): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${id}/mqtt`, httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('execMqtt'))
    );
  }

  execUpgrade(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/upgrade`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('execUpgrade'))
    );
  }

  execAlbum(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/album`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('execAlbum'))
    );
  }

  execCms(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/cms`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('execCms'))
    );
  }

  execCallback(autobuild: Autobuild): Observable<Response> {
    return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/callback`, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('execCallback'))
    );
  }

  delete(id: number): Observable<Response> {
    return this.http.delete<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${id}`, httpOptions).pipe(
      tap(res => this.log('PROCESS.DELETE', res)),
      catchError(this.handleError<Response>('deleteAutobuild'))
    );
  }

  getList(): Observable<Autobuild[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl, httpOptions).pipe(
      catchError(this.handleError<Response>('getAutobuilds')),
      map(res => {
        let abs:Autobuild[] = []; 
        if (res) {
          res.data.map(
            one => {
              abs.push(new Autobuild(one));
            }
          )
        }
        return abs; 
      }),
    );
  }

  getConfig(autobuild: Autobuild): Observable<Map<string, string>> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/config`, httpOptions).pipe(
      catchError(this.handleError<Response>('getConfigs')),
      map(res => {
        let configs:Map<string, string> = new Map(); 
        if (res) {
          for (let k in res.data["config"]) {
            configs.set(k, res.data["config"][k]);
          }
        }
        return configs; 
      })
    );
  }

  getMqtt(autobuild: Autobuild): Observable<Mqtt[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/mqtt`, httpOptions).pipe(
      catchError(this.handleError<Response>('getMqtts')),
      map(res => {
        let mqtts:Mqtt[] = []; 
        if (res) {
          res.data.map(
            one => {
              mqtts.push(new Mqtt(one));
            }
          )
        }
        return mqtts; 
      })
    );
  }

  getCallback(autobuild: Autobuild): Observable<Callback[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/callback`, httpOptions).pipe(
      catchError(this.handleError<Response>('getCallbacks')),
      map(res => {
        let cbs:Callback[] = []; 
        if (res) {
          res.data.map(
            one => {
              cbs.push(new Callback(one));
            }
          )
        }
        return cbs; 
      })
    );
  }

  getUpgrade(autobuild: Autobuild): Observable<Upgrade[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/upgrade`, httpOptions).pipe(
      catchError(this.handleError<Response>('getUpgrades')),
      map(res => {
        let ups:Upgrade[] = []; 
        if (res) {
          res.data.map(
            one => {
              ups.push(new Upgrade(one));
            }
          )
        }
        return ups; 
      })
    );
  }

  getAlbum(autobuild: Autobuild): Observable<Album[]> {
    return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${autobuild.Id}/album`, httpOptions).pipe(
      catchError(this.handleError<Response>('getAlbums')),
      map(res => {
        let als:Album[] = []; 
        if (res) {
          res.data.map(
            one => {
              als.push(new Album(one));
            }
          )
        }
        return als; 
      })
    );
  }

  one(id: number): Observable<Autobuild> {
      return this.http.get<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl + `/${id}/autobuild`, httpOptions).pipe(
      catchError(this.handleError<Response>('getAutobuild')),
      map(res => new Autobuild(res.data))
    );
  }

  add(autobuild: Autobuild): Observable<Autobuild> {
      return this.http.post<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.ADD', res)),
      catchError(this.handleError<Response>('addAutobuild')),
      map(res => new Autobuild(res.data))
    );
  }

  edit(autobuild: Autobuild): Observable<Autobuild> {
      return this.http.put<Response>(AppConfig.settings.apiServer.endpoint + this.autobuildUrl, JSON.stringify(autobuild), httpOptions).pipe(
      tap(res => this.log('PROCESS.UPDATE', res)),
      catchError(this.handleError<Response>('editAutobuild')),
      map(res => new Autobuild(res.data))
    );
  }

  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      // this.messageHandlerService.showError(`${operation} failed: ${error.message}`, '');
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
