import type { ProjectPriorityResponse as ProjectPriorityDTO } from "../types/dto";

export class ProjectPriority {
  id: string | null;
  name: string | null;
  hexColor: string | null;

  constructor(data?: ProjectPriorityDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.hexColor = data?.hexColor ?? null;
  }

  toDTO(): ProjectPriorityDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      hexColor: this.hexColor ?? "",
    };
  }
}

export const MAX_NAME_LENGTH = 32;
