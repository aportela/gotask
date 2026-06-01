import type { UserBaseResponse } from "../../users/types/dto";
import type { RoleResponse } from "../../roles/types/dto";

export type AddRequest = {
  user: UserBaseResponse;
  role: RoleResponse;
};

export type ProjectPermissionResponse = {
  id: string;
  user: UserBaseResponse;
  role: RoleResponse;
};

export type SearchResponse = {
  projectAttachments: ProjectPermissionResponse[];
};
