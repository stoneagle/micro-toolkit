import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'add-autobuild',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})

export class AddAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  newAdd(): void {
    this.autobuild = new Autobuild();
    this.modelOpened = true;
  }

  submit(): void {
    this.autobuildService.add(this.autobuild)
    .subscribe(autobuild => {
      this.modelOpened = false;
      this.create.emit(true);
    })
  }
}
