import { NgModule } from '@angular/core';
import { AutobuildComponent } from './autobuild/autobuild.component';
import { AddAutobuildComponent } from './autobuild/add/add.component';
import { AlbumAutobuildComponent } from './autobuild/album/album.component';
import { CallbackAutobuildComponent } from './autobuild/callback/callback.component';
import { MqttAutobuildComponent } from './autobuild/mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from './autobuild/upgrade/upgrade.component';
import { CmsAutobuildComponent } from './autobuild/cms/cms.component';
import { AutobuildService  } from '../../service/storybox/autobuild.service';
import { MessageHandlerService  } from '../../service/base/message-handler.service';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule  } from '@angular/platform-browser/animations';
import { ClarityModule  } from "@clr/angular";
import { HttpClientModule  }    from '@angular/common/http';
import { HttpModule  } from '@angular/http';
import { FormsModule }   from '@angular/forms';

@NgModule({
  declarations: [
    AutobuildComponent,
    AddAutobuildComponent,
    AlbumAutobuildComponent,
    CallbackAutobuildComponent,
    MqttAutobuildComponent,
    UpgradeAutobuildComponent,
    CmsAutobuildComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    HttpModule,
    FormsModule,
    ClarityModule,
  ],
  providers: [
    AutobuildService,
    MessageHandlerService
  ],
  exports: [
    AutobuildComponent,
  ]
})
export class StoryboxModule { 
}
