import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { Mqtt } from '../../../../model/storybox/mqtt';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'mqtt-autobuild',
  templateUrl: './mqtt.component.html',
  styleUrls: ['./mqtt.component.css']
})

export class MqttAutobuildComponent implements OnInit {
	autobuild: Autobuild = new Autobuild;
  mqtts: Mqtt[] = [];
  modelOpened: boolean = false;
  showOpened: boolean = false;
  @Output() create = new EventEmitter<string>();
  @Output() roll = new EventEmitter<string>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newMqtt(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    if (autobuild.Mqtt == 0) {
      this.modelOpened = true;
    } else {
      this.autobuildService.getMqtt(this.autobuild)
      .subscribe(res => {
        this.mqtts = res;
        this.showOpened = true;
      })
    }
  }

  rollback(): void {
    this.autobuildService.rollbackMqtt(this.autobuild)
    .subscribe(res => {
      this.showOpened = false;
      this.roll.emit("mqtt");
    })
  }

  submit(): void {
    this.autobuildService.execMqtt(this.autobuild.Id)
    .subscribe(res => {
      this.modelOpened = false;
      this.create.emit("mqtt");
    })
  }
}
