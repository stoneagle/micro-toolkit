import { Component, OnInit } from '@angular/core';
import { AutoBuild } from '../../../model/storybox/autobuild';
import { AutoBuildService  } from '../../../service/storybox/autobuild.service';

@Component({
  selector: 'app-autobuild',
  templateUrl: './autobuild.component.html',
  styleUrls: ['./autobuild.component.css']
})
export class AutobuildComponent implements OnInit {

  autobuilds: AutoBuild[];

  constructor(private autoBuildService: AutoBuildService) { }

  ngOnInit() {
    this.getBuilds();
  }

  getBuilds(): void {
    // this.autoBuildService().getAutoBuilds().subscribe(autobuilds => this.autobuilds = autobuilds)
  }

}
