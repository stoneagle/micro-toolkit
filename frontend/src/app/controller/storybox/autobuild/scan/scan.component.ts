import { Component, OnInit } from '@angular/core';
import { DeviceScan } from '../../../../model/storybox/scan';
import { ScanService  } from '../../../../service/storybox/scan.service';

@Component({
  selector: 'scan-autobuild',
  templateUrl: './scan.component.html',
  styleUrls: ['./scan.component.css']
})

export class ScanAutobuildComponent implements OnInit {
  scan: DeviceScan = new DeviceScan();
  scans: DeviceScan[] = [];

  scanTypeConfig : Map<number, string> = new Map([[1, "二维码"], [2, "条形码"]]);
  funcTypeConfig : Map<number, string> = new Map([[1, "配网"]]);
  infoTypeConfig : Map<number, string> = new Map([[1, "文本"]]);

  appId: string;
  modelOpened: boolean = false;
  showOpened: boolean = false;

	pageSize: number = 5;
	totalCount: number = 0;
	currentPage: number = 1;

  constructor(
    private scanService: ScanService,
  ) { }

  ngOnInit() {
  }

  openCreate() {
    this.scan = new DeviceScan();
    this.scan.AppId = this.appId;
    this.modelOpened = true;
    this.showOpened = false;
  }

  openUpdate(scan: DeviceScan) {
    this.scan = scan;
    console.log(this.scan);
    this.modelOpened = true;
    this.showOpened = false;
  }

  closeCreate() {
    this.modelOpened = false;
    this.showOpened = true;
  }

  newScan(appId: string): void {
    this.appId = appId;
    this.pageSize = 5;
    this.refresh();
    this.showOpened = true;
  }

	load(state: any): void {
    if (state && state.page) {
      this.refreshScans(state.page.from, state.page.to + 1);
    }
  }

  refresh() {
    this.currentPage = 1;
    this.refreshScans(0, 10);
  }

  refreshScans(from: number, to: number): void {
    this.scanService.list(this.appId)
    .subscribe(res => {
      this.scans = res.slice(from, to);
			this.totalCount = res.length;
    })
  }

  delete(scan: DeviceScan): void {
    this.scanService.delete(scan.Id)
    .subscribe(res => {
      this.refresh();
    })
  }

  submit(): void {
    if (this.scan.Id == null) {
      this.scanService.add(this.scan)
      .subscribe(res => {
        this.modelOpened = false;
        this.showOpened = true;
        this.refresh();
      })
    } else {
      this.scanService.update(this.scan)
      .subscribe(res => {
        this.modelOpened = false;
        this.showOpened = true;
        this.refresh();
      })
    }
  }

  getKeys(map) {
    return Array.from(map.keys());
  }
}
