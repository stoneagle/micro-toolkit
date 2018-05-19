import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';

import { Observable  } from 'rxjs';
import { catchError, map, tap  } from 'rxjs/operators';

import { AutoBuild } from '../../model/storybox/autobuild';

@Injectable()
export class AutoBuildService {
  constructor() { }
}
