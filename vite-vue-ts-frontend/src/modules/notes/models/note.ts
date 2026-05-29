import type { NoteResponse as NoteDTO } from "../types/dto";
import { UserBase } from "../../users/models/user";
import { IDate } from "../../../shared/types/idate";

export class Note {
  id: string | null;
  user: UserBase;
  createdAt: IDate | null;
  updatedAt: IDate | null;
  body: string;

  constructor(data?: NoteDTO) {
    this.id = data?.id ?? null;
    this.user = new UserBase(data?.user);
    this.createdAt = data?.createdAt ? new IDate(data.createdAt) : null;
    this.updatedAt = data?.updatedAt ? new IDate(data.updatedAt) : null;
    this.body = data?.body ?? "";
  }

  toDTO(): NoteDTO {
    return {
      id: this.id ?? "",
      user: this.user.toDTO(),
      createdAt: this.createdAt?.msTimestamp ?? 0,
      updatedAt: this.updatedAt?.msTimestamp ?? null,
      body: this.body ?? "",
    };
  }
}
