import { Serializable } from '../app.serializable.model';

export class Album extends Serializable {
  Id: number;
  AppId: number;
  AlbumId: number;
  CreateAt: string;
  UpdateAt: string;
  DeleteAt: string;
}
