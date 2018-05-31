import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'mqtt-autobuild',
  templateUrl: './mqtt.component.html',
  styleUrls: ['./mqtt.component.css']
})

export class MqttAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  modelOpened: boolean = false;
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newMqtt(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    this.modelOpened = true;
  }

  submit(): void {
    this.autobuildService.execMqtt(this.autobuild.Id)
    .subscribe(res => {
      this.modelOpened = false;
      this.create.emit(true);
    })
  }
}
