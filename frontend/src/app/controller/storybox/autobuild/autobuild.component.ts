import { Component, OnInit } from '@angular/core';
import { Autobuild } from '../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../service/storybox/autobuild.service';

@Component({
  providers: [AutobuildService],
  selector: 'app-autobuild',
  templateUrl: './autobuild.component.html',
  styleUrls: ['./autobuild.component.css']
})
export class AutobuildComponent implements OnInit {
  autobuilds: Autobuild[] = [];
	addAutobuild: Autobuild = new Autobuild;
	mqttAutobuild: Autobuild = new Autobuild;
	callbackAutobuild: Autobuild = new Autobuild;
	upgradeAutobuild: Autobuild = new Autobuild;
	albumlistAutobuild: Autobuild = new Autobuild;

	pageSize: number = 10;
	totalCount: number = 0;
	currentPage: number = 1;

  addModel: boolean = false;
  cmsModel: boolean = false;
  mqttModel: boolean = false;
  callbackModel: boolean = false;
  upgradeModel: boolean = false;
  albumlistModel: boolean = false;

  mqttShowModel: boolean = false;
  callbackShowModel: boolean = false;
  upgradeShowModel: boolean = false;
  albumlistShowModel: boolean = false;

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
    this.pageSize = 10;
    this.refresh();
  }

  showAddModel(): void {
    this.addModel = true;
  }

  submitAdd(): void {
    this.addModel = false;
    this.autobuildService.add(this.addAutobuild)
    .subscribe(autobuild => {
      this.autobuilds.push(autobuild);
    })
  }

  delete(ab: Autobuild): void {
    this.autobuildService.delete(ab.Id)
    .subscribe(autobuild => {
      this.refresh();
    })
  }

  showCmsModel(ab: Autobuild): void {
    this.cmsModel = true;
  }

  showMqttModel(ab: Autobuild): void {
    if (ab.Mqtt == 0) {
      this.mqttAutobuild = ab;
      this.mqttModel = true;
    }
  }

	execMqtt(): void {
    this.autobuildService.execMqtt(this.mqttAutobuild.Id)
    .subscribe(res => {
      this.mqttModel = false;
      this.refresh();
    })
	}

  showCallbackModel(ab: Autobuild): void {
    if (ab.Callback == '') {
      this.callbackAutobuild = ab;
      this.callbackModel = true;
    }
  }

  execCallback(): void {
    if (this.callbackAutobuild.Callback != '') {
      this.autobuildService.execCallback(this.callbackAutobuild)
      .subscribe(res => {
        this.callbackModel = false;
        this.refresh();
      })
    }
  }

  showUpgradeModel(ab: Autobuild): void {
    if (ab.UpgradeName == '') {
      this.upgradeAutobuild = ab;
      this.upgradeModel = true;
    }
  }

  execUpgrade(): void {
    if (this.upgradeAutobuild.UpgradeName != '' && this.upgradeAutobuild.UpgradeVcode != 0) {
      this.autobuildService.execUpgrade(this.upgradeAutobuild)
      .subscribe(res => {
        this.upgradeModel = false;
        this.refresh();
      })
    }
  }

  showAlbumListModel(ab: Autobuild): void {
    if (ab.AlbumList == '') {
      this.albumlistAutobuild = ab;
      this.albumlistModel = true;
    }
  }

  execAlbum(): void {
    if (this.albumlistAutobuild.AlbumList != '') {
      this.autobuildService.execAlbum(this.albumlistAutobuild)
      .subscribe(res => {
        this.albumlistModel = false;
        this.refresh();
      })
    }
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
