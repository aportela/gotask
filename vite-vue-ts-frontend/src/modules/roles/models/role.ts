import {
  type RoleResponse as RoleDTO,
  getDefaultPermissions,
} from "../types/dto";

type Permissions = {
  allowUpdateProject: boolean;
  allowDeleteProject: boolean;
  allowViewProject: boolean;
  allowAddTask: boolean;
  allowUpdateTask: boolean;
  allowDeleteTask: boolean;
  allowViewTask: boolean;
};

export class Role {
  id: string | null;
  name: string | null;
  permissions: Permissions;

  constructor(data?: RoleDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.permissions = data?.permissions ?? getDefaultPermissions();
  }

  toDTO(): RoleDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      permissions: {
        allowUpdateProject: this.permissions?.allowUpdateProject ?? false,
        allowDeleteProject: this.permissions?.allowDeleteProject ?? false,
        allowViewProject: this.permissions?.allowViewProject ?? false,
        allowAddTask: this.permissions?.allowAddTask ?? false,
        allowUpdateTask: this.permissions?.allowUpdateTask ?? false,
        allowDeleteTask: this.permissions?.allowDeleteTask ?? false,
        allowViewTask: this.permissions?.allowViewTask ?? false,
      },
    };
  }
}

export const MAX_NAME_LENGTH = 32;
