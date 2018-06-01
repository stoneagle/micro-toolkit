import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Autobuild } from '../../../../model/storybox/autobuild';
import { Album } from '../../../../model/storybox/album';
import { AutobuildService  } from '../../../../service/storybox/autobuild.service';

@Component({
  selector: 'album-autobuild',
  templateUrl: './album.component.html',
  styleUrls: ['./album.component.css']
})

export class AlbumAutobuildComponent implements OnInit {
  modelOpened: boolean = false;
  showOpened: boolean = false;
  albums: Album[] = [];
  autobuild: Autobuild = new Autobuild();
  @Output() create = new EventEmitter<string>();
  @Output() roll = new EventEmitter<string>();

  constructor(
    private autobuildService: AutobuildService,
  ) { }

  ngOnInit() {
  }

  newAlbum(autobuild: Autobuild): void {
    this.autobuild = autobuild;
    if (this.autobuild.AlbumList == '') {
      this.modelOpened = true;
    } else {
      this.autobuildService.getAlbum(this.autobuild)
      .subscribe(res => {
        this.albums = res;
        this.showOpened = true;
      })
    }
  }

  rollback(): void {
    this.autobuildService.rollbackAlbum(this.autobuild)
    .subscribe(res => {
      this.showOpened = false;
      this.roll.emit("album");
    })
  }

  submit(): void {
    if (this.autobuild.AlbumList != '') {
      this.autobuildService.execAlbum(this.autobuild)
      .subscribe(res => {
        this.modelOpened = false;
        this.create.emit("album");
      })
    }
  }
}
