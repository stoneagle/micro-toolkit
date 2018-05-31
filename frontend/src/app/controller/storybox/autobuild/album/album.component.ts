import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'album-autobuild',
  templateUrl: './album.component.html',
  styleUrls: ['./album.component.css']
})

export class AlbumAutobuildComponent implements OnInit {
  modelOpened: boolean = false;
  autobuild: Autobuild = new Autobuild();
  @Output() create = new EventEmitter<boolean>();

  constructor(
    private autobuildService: AutobuildService
  ) { }

  ngOnInit() {
  }

  newAlbum(autobuild: Autobuild): void {
    this.modelOpened = true;
    this.autobuild = autobuild;
  }

  submit(): void {
    if (this.autobuild.AlbumList != '') {
      this.autobuildService.execAlbum(this.autobuild)
      .subscribe(res => {
        this.modelOpened = false;
        this.create.emit(true);
      })
    }
  }
}
