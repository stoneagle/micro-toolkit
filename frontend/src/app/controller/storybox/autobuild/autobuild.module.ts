import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { AutobuildComponent } from './autobuild.component';
import { AddAutobuildComponent } from './add/add.component';
import { AlbumAutobuildComponent } from './album/album.component';
import { CallbackAutobuildComponent } from './callback/callback.component';
import { MqttAutobuildComponent } from './mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from './upgrade/upgrade.component';
import { CmsAutobuildComponent } from './cms/cms.component';
import { AutobuildService  } from '../../../service/storybox/autobuild.service';

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
  ],
  exports: [
    AutobuildComponent,
  ]
})
export class AutobuildModule { 
}
