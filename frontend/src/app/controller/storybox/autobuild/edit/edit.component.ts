import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'edit-autobuild',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.css']
})

export class EditAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  @Output() edit = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  newEdit(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    this.modelOpened = true;
  }

  submit(): void {
    this.autobuildService.edit(this.autobuild).subscribe(autobuild => {
      this.modelOpened = false;
      this.edit.emit(true);
    })
  }
}
