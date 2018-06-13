import { Serializable } from '../app.serializable.model';

export class DeviceScan extends Serializable {
  Id: number;
  AppId: string;
  ScanType: number;
  FuncType: number;
  InfoType: number;
  Info: string;
  CreateAt: string;
  UpdateAt: string;
}

