import { Component, OnInit, ViewChild } from '@angular/core';
import { Autobuild } from '../../../model/storybox/autobuild';
import { AddAutobuildComponent } from './add/add.component';
import { AlbumAutobuildComponent } from './album/album.component';
import { CallbackAutobuildComponent } from './callback/callback.component';
import { MqttAutobuildComponent } from './mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from './upgrade/upgrade.component';
import { CmsAutobuildComponent } from './cms/cms.component';
import { AutobuildService  } from '../../../service/storybox/autobuild.service';
import { MessageHandlerService  } from '../../../service/base/message-handler.service';

@Component({
  selector: 'app-autobuild',
  templateUrl: './autobuild.component.html',
  styleUrls: ['./autobuild.component.css']
})
export class AutobuildComponent implements OnInit {
  @ViewChild(AddAutobuildComponent)
  addAutobuild: AddAutobuildComponent;
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

  autobuilds: Autobuild[] = [];

	pageSize: number = 10;
	totalCount: number = 0;
	currentPage: number = 1;

  constructor(
    private autobuildService: AutobuildService,
    private messageHandlerService: MessageHandlerService 
  ) { }

  ngOnInit() {
    this.pageSize = 10;
    this.refresh();
  }

  createFinished(resource: string): void {
    if (resource != '') {
      this.refresh();
    }
  }

  rollbackFinished(resource: string): void {
    if (resource != '') {
      this.refresh();
    }
  }

  openAlbumModel(ab: Autobuild): void {
    this.albumAutobuild.newAlbum(ab);
  }

  openCallbackModel(ab: Autobuild): void {
    this.callbackAutobuild.newCallback(ab);
  }

  openMqttModel(ab: Autobuild): void {
    this.mqttAutobuild.newMqtt(ab);
  }

  openCmsModel(ab: Autobuild): void {
    this.cmsAutobuild.newCms(ab);
  }

  openUpgradeModel(ab: Autobuild): void {
    this.upgradeAutobuild.newUpgrade(ab);
  }

  openAddModel(): void {
    this.addAutobuild.newAdd();
  }

  delete(ab: Autobuild): void {
    this.autobuildService.delete(ab.Id)
    .subscribe(autobuild => {
      this.refresh();
    })
  }

	load(state: any): void {
    if (state && state.page) {
      this.refreshAutobuilds(state.page.from, state.page.to + 1);
    }
  }

  refresh() {
    this.currentPage = 1;
    this.refreshAutobuilds(0, 10);
  }

  refreshAutobuilds(from: number, to: number): void {
    this.autobuildService.getList()
    .subscribe(res => {
			this.totalCount = res.length;
      this.autobuilds = res.slice(from, to);
    })
  }
}
