import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { Upgrade } from '../../../../model/storybox/upgrade';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'upgrade-autobuild',
  templateUrl: './upgrade.component.html',
  styleUrls: ['./upgrade.component.css']
})

export class UpgradeAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  upgrades: Upgrade[] = [];
  modelOpened: boolean = false;
  showOpened: boolean = false;
  @Output() create = new EventEmitter<string>();
  @Output() roll = new EventEmitter<string>();

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  newUpgrade(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    if (this.autobuild.UpgradeName == '') {
      this.modelOpened = true;
    } else {
      this.autobuildService.getUpgrade(this.autobuild)
      .subscribe(res => {
        this.upgrades = res;
        this.showOpened = true;
      })
    }
  }

  rollback(): void {
    this.autobuildService.rollbackUpgrade(this.autobuild)
    .subscribe(res => {
      this.showOpened = false;
      this.roll.emit("upgrade");
    })
  }

  submit(): void {
    if (this.autobuild.UpgradeName != '' && this.autobuild.UpgradeVcode != 0) {
      this.autobuildService.execUpgrade(this.autobuild)
      .subscribe(res => {
        this.create.emit("upgrade");
        this.modelOpened = false;
      })
    }
  }
}
