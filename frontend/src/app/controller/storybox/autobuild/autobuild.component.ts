import { Component, OnInit, ViewChild } from '@angular/core';
import { Autobuild } from '../../../model/storybox/autobuild';
import { AddAutobuildComponent } from './add/add.component';
import { AlbumAutobuildComponent } from './album/album.component';
import { CallbackAutobuildComponent } from './callback/callback.component';
import { MqttAutobuildComponent } from './mqtt/mqtt.component';
import { UpgradeAutobuildComponent } from './upgrade/upgrade.component';
import { AutobuildService  } from '../../../service/storybox/autobuild.service';

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

  autobuilds: Autobuild[] = [];

	pageSize: number = 10;
	totalCount: number = 0;
	currentPage: number = 1;

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
    this.pageSize = 10;
    this.refresh();
  }

  createAutobuild(created: boolean):void {
    if (created) {
      this.refresh();
    }
  }

  execUpgrade(created: boolean): void {
    if (created) {
      this.refresh();
    }
  }

  execAlbum(created: boolean): void {
    if (created) {
      this.refresh();
    }
  }

  execCallback(created: boolean): void {
    if (created) {
      this.refresh();
    }
  }

	execMqtt(created: boolean): void {
    if (created) {
      this.refresh();
    }
	}

  openAlbumModel(ab: Autobuild): void {
    if (ab.AlbumList == '') {
      this.albumAutobuild.newAlbum(ab);
    }
  }

  openCallbackModel(ab: Autobuild): void {
    if (ab.Callback == '') {
      this.callbackAutobuild.newCallback(ab);
    }
  }

  openMqttModel(ab: Autobuild): void {
    if (ab.Mqtt == 0) {
      this.mqttAutobuild.newMqtt(ab);
    }
  }

  openCmsModel(ab: Autobuild): void {
    // this.cmsModel = true;
  }

  openUpgradeModel(ab: Autobuild): void {
    if (ab.UpgradeName == '') {
      this.upgradeAutobuild.newUpgrade(ab);
    }
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
