import { Serializable } from '../app.serializable.model';

export class Callback extends Serializable {
  Id: number;
  AppId: number;
  State: number;
  CallbackUrl: string;
  CallbackState: string;
  Action: string;
  CreateAt: string;
  UpdateAt: string;
}
