import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { Callback } from '../../../../model/storybox/callback';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'callback-autobuild',
  templateUrl: './callback.component.html',
  styleUrls: ['./callback.component.css']
})

export class CallbackAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  callbacks: Callback[] = [];
  modelOpened: boolean = false;
  showOpened: boolean = false;
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newCallback(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    if (this.autobuild.Callback == '') {
      this.modelOpened = true;
    } else {
      this.autobuildService.getCallback(this.autobuild)
      .subscribe(res => {
        this.callbacks = res;
        this.showOpened = true;
      })
    }
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
