import { Serializable } from '../app.serializable.model';

export class Upgrade extends Serializable {
  Id: number;
  Name: string;
  Vcode: string;
  Vname: string;
  CreateAt: string;
  UpdateAt: string;
  DeleteAt: string;
}
