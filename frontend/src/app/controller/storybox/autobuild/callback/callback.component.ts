import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'callback-autobuild',
  templateUrl: './callback.component.html',
  styleUrls: ['./callback.component.css']
})

export class CallbackAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newCallback(autobuild: Autobuild): void {
    this.modelOpened = true;
    this.autobuild = autobuild;
  }

  submit(): void {
    if (this.autobuild.Callback != '') {
      this.autobuildService.execCallback(this.autobuild)
      .subscribe(res => {
        this.create.emit(true);
        this.modelOpened = false;
      })
    }
  }
}
