import { NgModule } from '@angular/core';
import { AutobuildComponent } from './autobuild/autobuild.component';
import { AddAutobuildComponent } from './autobuild/add/add.component';
import { EditAutobuildComponent } from './autobuild/edit/edit.component';
import { ExpandAutobuildComponent } from './autobuild/expand/expand.component';
import { ConfigAutobuildComponent } from './autobuild/config/config.component';
import { AlbumAutobuildComponent } from './autobuild/album/album.component';
import { CallbackAutobuildComponent } from './autobuild/callback/callback.component';
import { MqttAutobuildComponent } from './autobuild/mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from './autobuild/upgrade/upgrade.component';
import { CmsAutobuildComponent } from './autobuild/cms/cms.component';
import { ScanAutobuildComponent } from './autobuild/scan/scan.component';
import { AutobuildService  } from '../../service/storybox/autobuild.service';
import { ScanService  } from '../../service/storybox/scan.service';
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
    EditAutobuildComponent,
    ExpandAutobuildComponent,
    AlbumAutobuildComponent,
    ConfigAutobuildComponent,
    CallbackAutobuildComponent,
    MqttAutobuildComponent,
    UpgradeAutobuildComponent,
    CmsAutobuildComponent,
    ScanAutobuildComponent,
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
    ScanService,
    MessageHandlerService
  ],
  exports: [
    AutobuildComponent,
  ]
})
export class StoryboxModule { 
}
