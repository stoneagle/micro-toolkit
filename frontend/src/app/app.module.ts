import { APP_INITIALIZER } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule  } from '@angular/platform-browser/animations';
import { FormsModule }   from '@angular/forms';
import { NgModule } from '@angular/core';
import { HttpClientModule  }    from '@angular/common/http';
import { HttpModule  } from '@angular/http';

import { ClarityModule  } from "@clr/angular";
import { AppComponent } from './app.component';
import { AutobuildComponent } from './controller/storybox/autobuild/autobuild.component';
import { AppRoutingModule } from './app-routing.module';
import { AppConfig } from './service/app.config';

export function initializeApp(appConfig: AppConfig) {
  return () => appConfig.load();
}

@NgModule({
  declarations: [
    AppComponent,
    AutobuildComponent,
  ],
  imports: [
    ClarityModule,
    AppRoutingModule,
    HttpClientModule,
    HttpModule,
    BrowserModule,
    FormsModule,
    BrowserAnimationsModule,
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
