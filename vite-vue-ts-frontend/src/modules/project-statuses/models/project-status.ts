import type { ProjectStatusResponse as ProjectStatusDTO } from "../types/dto";

export class ProjectStatus {
  id: string | null;
  name: string | null;
  hexColor: string | null;

  constructor(data?: ProjectStatusDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.hexColor = data?.hexColor ?? null;
  }

  toDTO(): ProjectStatusDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      hexColor: this.hexColor ?? "",
    };
  }
}

export const MAX_NAME_LENGTH = 32;
