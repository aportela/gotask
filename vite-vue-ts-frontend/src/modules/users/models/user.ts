import type {
  UserResponse as UserDTO,
  UserBaseResponse as UserBaseDTO,
} from "../types/dto";
import { IDate } from "../../../shared/types/idate";

interface UserPermissions {
  isSuperUser: boolean;
}

export class UserBase {
  id: string | null;
  name: string | null;

  constructor(data?: UserBaseDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
  }

  toDTO(): UserBaseDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
    };
  }
}

export class User {
  id: string | null;
  name: string | null;
  email: string | null;
  password: string | null;
  permissions: UserPermissions;
  createdAt: IDate | null;
  updatedAt: IDate | null;
  deletedAt: IDate | null;

  constructor(data?: UserDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.email = data?.email ?? null;
    this.password = null;
    this.permissions = {
      isSuperUser: data?.permissions.isSuperUser ?? false,
    };
    this.createdAt = data?.createdAt ? new IDate(data.createdAt) : null;
    this.updatedAt = data?.updatedAt ? new IDate(data.updatedAt) : null;
    this.deletedAt = data?.deletedAt ? new IDate(data.deletedAt) : null;
  }

  toDTO(): UserDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      email: this.email ?? "",
      permissions: {
        isSuperUser: this.permissions?.isSuperUser ?? false,
      },
      createdAt: this.createdAt?.msTimestamp ?? 0,
      updatedAt: this.updatedAt?.msTimestamp ?? null,
      deletedAt: this.deletedAt?.msTimestamp ?? null,
    };
  }
}

export const MAX_NAME_LENGTH = 32;
export const MAX_EMAIL_LENGTH = 255;

export const MIN_PASSWORD_LENGTH = 4;
