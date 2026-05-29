import type { ProjectTypeResponse as ProjectTypeDTO } from "../types/dto";

export class ProjectType {
  id: string | null;
  name: string | null;
  hexColor: string | null;

  constructor(data?: ProjectTypeDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.hexColor = data?.hexColor ?? null;
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
