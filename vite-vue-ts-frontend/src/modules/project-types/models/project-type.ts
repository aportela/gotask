import type { ProjectTypeResponse as ProjectTypeDTO } from "../types/dto";

export class ProjectType {
  id: string | null;
  name: string | null;
  hexColor: string | null;

  constructor(data?: ProjectTypeDTO) {
    this.id = data?.id ?? "";
    this.name = data?.name ?? "";
    this.hexColor = data?.hexColor ?? "";
  }

  toDTO(): ProjectTypeDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      hexColor: this.hexColor ?? "",
    };
  }
}

export const MAX_NAME_LENGTH = 32;
