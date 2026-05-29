import type { ProjectPermissionResponse as ProjectPermissionDTO } from "../types/dto";
import { UserBase } from "../../users/models/user";
import { Role } from "../../roles/models/role";

export class ProjectPermission {
  id: string | null;
  user: UserBase;
  role: Role;

  constructor(data?: ProjectPermissionDTO) {
    this.id = data?.id ?? null;
    this.user = new UserBase(data?.user);
    this.role = new Role(data?.role);
  }

  toDTO(): ProjectPermissionDTO {
    return {
      id: this.id ?? "",
      user: this.user.toDTO(),
      role: this.role.toDTO(),
    };
  }
}
