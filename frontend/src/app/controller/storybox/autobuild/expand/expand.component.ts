import { Component, OnInit, Input, Output, EventEmitter, ViewChild } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';
import { AlbumAutobuildComponent } from '../album/album.component';
import { CallbackAutobuildComponent } from '../callback/callback.component';
import { MqttAutobuildComponent } from '../mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from '../upgrade/upgrade.component';
import { CmsAutobuildComponent } from '../cms/cms.component';
import { ConfigAutobuildComponent } from '../config/config.component';
import { ScanAutobuildComponent } from '../scan/scan.component';

@Component({
  selector: 'expand-autobuild',
  templateUrl: './expand.component.html',
  styleUrls: ['./expand.component.css']
})

export class ExpandAutobuildComponent implements OnInit {
  @ViewChild(AlbumAutobuildComponent)
  albumAutobuild: AlbumAutobuildComponent;
  @ViewChild(CallbackAutobuildComponent)
  callbackAutobuild: CallbackAutobuildComponent;
  @ViewChild(MqttAutobuildComponent)
  mqttAutobuild: MqttAutobuildComponent;
  @ViewChild(UpgradeAutobuildComponent)
  upgradeAutobuild: UpgradeAutobuildComponent;
  @ViewChild(CmsAutobuildComponent)
  cmsAutobuild: CmsAutobuildComponent;
  @ViewChild(ScanAutobuildComponent)
  scanAutobuild: ScanAutobuildComponent;
  @ViewChild(ConfigAutobuildComponent)
  configAutobuild: ConfigAutobuildComponent;

  @Input() build:Autobuild;

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  openAlbumModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.albumAutobuild.newAlbum(autobuild);
  }

  openCallbackModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.callbackAutobuild.newCallback(autobuild);
  }

  openMqttModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.mqttAutobuild.newMqtt(autobuild);
  }

  openCmsModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.cmsAutobuild.newCms(autobuild);
  }

  openUpgradeModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.upgradeAutobuild.newUpgrade(autobuild);
  }

  openScanModel(ab: Autobuild): void {
    this.scanAutobuild.newScan(ab.AppId);
  }

  openConfigModel(ab: Autobuild): void {
    this.configAutobuild.newConfig(ab);
  }

  rollbackFinished(resource: string): void {
    if (resource != '') {
      this.autobuildService.one(this.build.Id).subscribe(
        res => {
          this.build = res;
        }
      ) 
    }
  }

  createFinished(resource: string): void {
    if (resource != '') {
      this.autobuildService.one(this.build.Id).subscribe(
        res => {
          this.build = res;
        }
      ) 
    }
  }
}
