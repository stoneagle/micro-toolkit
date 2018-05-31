import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'upgrade-autobuild',
  templateUrl: './upgrade.component.html',
  styleUrls: ['./upgrade.component.css']
})

export class UpgradeAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newUpgrade(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    this.modelOpened = true;
  }

  submit(): void {
    if (this.autobuild.UpgradeName != '' && this.autobuild.UpgradeVcode != 0) {
      this.autobuildService.execUpgrade(this.autobuild)
      .subscribe(res => {
        this.create.emit(true);
        this.modelOpened = false;
      })
    }
  }
}
