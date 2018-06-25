import { Serializable } from '../app.serializable.model';

export class Autobuild extends Serializable {
  Id: number;
  AppId: string; 
  Desc: string; 
  CmsSourceApp: string;
  CmsExecTime: string;
  Mqtt: number;
  Callback: string;
  UpgradeName: string;
  UpgradeVcode: number;
  UpgradeVname: string;
  AlbumList: string;
  CreateAt: string;
  UpdateAt: string;
  DeleteAt: string;
}
