import { APP_INITIALIZER } from '@angular/core';
import { NgModule } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { TranslateModule, TranslateLoader } from '@ngx-translate/core';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';
import { HTTP_INTERCEPTORS  } from '@angular/common/http';
import { CustomInterceptor } from './service/base/custom.interceptor';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { AppConfig } from './service/app.config';
import { ClarityModule  } from "@clr/angular";
import { BaseModule } from './base/base.module';
import { StoryboxModule } from './controller/storybox/storybox.module';

export function initializeApp(appConfig: AppConfig) {
  return () => appConfig.load();
}

export function createTranslateLoader(http: HttpClient) {
  return new TranslateHttpLoader(http, './assets/i18n/', '.json');
}

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    ClarityModule,
    AppRoutingModule,
    StoryboxModule,
    BaseModule,
    TranslateModule.forRoot({
      loader: {
        provide: TranslateLoader,
        useFactory: (createTranslateLoader),
        deps: [HttpClient]
      }
    })
  ],
  providers: [
    AppConfig,
    { 
      provide: APP_INITIALIZER,
      useFactory: initializeApp,
      deps: [AppConfig], 
      multi: true 
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: CustomInterceptor ,
      multi: true
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
