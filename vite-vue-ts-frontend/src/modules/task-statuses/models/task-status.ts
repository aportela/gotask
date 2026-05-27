import type { TaskStatusResponse as TaskStatusDTO } from "../types/dto";

export class TaskStatus {
  id: string | null;
  name: string | null;
  hexColor: string | null;

  constructor(data?: TaskStatusDTO) {
    this.id = data?.id ?? null;
    this.name = data?.name ?? null;
    this.hexColor = data?.hexColor ?? null;
  }

  toDTO(): TaskStatusDTO {
    return {
      id: this.id ?? "",
      name: this.name ?? "",
      hexColor: this.hexColor ?? "",
    };
  }
}

export const MAX_NAME_LENGTH = 32;
