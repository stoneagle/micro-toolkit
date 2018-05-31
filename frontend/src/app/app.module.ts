import { APP_INITIALIZER } from '@angular/core';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { AutobuildModule } from './controller/storybox/autobuild/autobuild.module';
import { AppRoutingModule } from './app-routing.module';
import { AppConfig } from './service/app.config';
import { ClarityModule  } from "@clr/angular";

export function initializeApp(appConfig: AppConfig) {
  return () => appConfig.load();
}

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    ClarityModule,
    AutobuildModule,
    AppRoutingModule,
  ],
  providers: [
    AppConfig,
    { 
      provide: APP_INITIALIZER,
      useFactory: initializeApp,
      deps: [AppConfig], 
      multi: true 
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
