import { Component, OnInit } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'config-autobuild',
  templateUrl: './config.component.html',
  styleUrls: ['./config.component.css']
})

export class ConfigAutobuildComponent implements OnInit {
  configs: Map<string, string> = new Map();
  showOpened: boolean = false;
  autobuild: Autobuild;

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  newConfig(ab: Autobuild): void {
    this.autobuild = ab;
    this.autobuildService.getConfig(this.autobuild)
    .subscribe(res => {
      this.configs = res;
    });
    this.showOpened = true;
  }

  getKeys(map) {
    return Array.from(map.keys());
  }
}
