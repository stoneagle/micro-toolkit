import { Component, OnInit, ViewChild } from '@angular/core';
import { Autobuild } from '../../../model/storybox/autobuild';
import { AddAutobuildComponent } from './add/add.component';
import { EditAutobuildComponent } from './edit/edit.component';
import { ExpandAutobuildComponent } from './expand/expand.component';
import { AutobuildService  } from '../../../service/storybox/autobuild.service';
import { MessageHandlerService  } from '../../../service/base/message-handler.service';

@Component({
  selector: 'app-autobuild',
  templateUrl: './autobuild.component.html',
  styleUrls: ['./autobuild.component.css']
})
export class AutobuildComponent implements OnInit {
  @ViewChild(AddAutobuildComponent)
  addAutobuild: AddAutobuildComponent;
  @ViewChild(EditAutobuildComponent)
  editAutobuild: EditAutobuildComponent;
  @ViewChild(ExpandAutobuildComponent)
  expandAutobuild: ExpandAutobuildComponent;

  autobuilds: Autobuild[] = [];

	pageSize: number = 10;
	totalCount: number = 0;
	currentPage: number = 1;

  constructor(
    private autobuildService: AutobuildService,
    private messageHandlerService: MessageHandlerService 
  ) { }

  ngOnInit() {
    this.pageSize = 10;
    this.refresh();
  }

  createFinished(resource: string): void {
    if (resource != '') {
      this.refresh();
    }
  }

  editFinished(resource: string): void {
    if (resource != '') {
      this.refresh();
    }
  }

  openAddModel(): void {
    this.addAutobuild.newAdd();
  }

  openEditModel(ab: Autobuild): void {
    let autobuild = Object.assign({}, ab);
    this.editAutobuild.newEdit(autobuild);
  }

  delete(ab: Autobuild): void {
    this.autobuildService.delete(ab.Id)
    .subscribe(autobuild => {
      this.refresh();
    })
  }

	load(state: any): void {
    if (state && state.page) {
      this.refreshAutobuilds(state.page.from, state.page.to + 1);
    }
  }

  refresh() {
    this.refreshAutobuilds(0, 10);
  }

  refreshAutobuilds(from: number, to: number): void {
    this.autobuildService.getList()
    .subscribe(res => {
			this.totalCount = res.length;
      this.autobuilds = res.slice(from, to);
    })
  }
}
