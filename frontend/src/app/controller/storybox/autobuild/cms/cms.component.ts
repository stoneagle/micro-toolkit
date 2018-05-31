import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'cms-autobuild',
  templateUrl: './cms.component.html',
  styleUrls: ['./cms.component.css']
})

export class CmsAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  alertOpened: boolean = false;
  @Output() create = new EventEmitter<string>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newCms(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    if (this.autobuild.CmsSourceApp == '') {
      this.modelOpened = true;
    }
  }

  submit(): void {
    if (this.autobuild.CmsSourceApp != '' && this.alertOpened == false) {
		  this.alertOpened = true;
      this.autobuildService.execCms(this.autobuild)
      .subscribe(res => {
        this.modelOpened = false;
		  	this.alertOpened = false;
        this.create.emit("cms");
      })
    }
  }
}
