import { Serializable } from '../app.serializable.model';

export class Mqtt extends Serializable {
  Id: number;
  ChannelType: string;
  Production: string;
  Service: string;
  Params: string;
  Status: number;
  CreateAt: string;
  UpdateAt: string;
}
