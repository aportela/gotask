import type { ProjectResponse as ProjectDTO } from "../types/dto";
import { ProjectType } from "../../project-types/models/project-type";
import { ProjectPriority } from "../../project-priorities/models/project-priority";
import { ProjectStatus } from "../../project-statuses/models/project-status";
import { UserBase } from "../../users/models/user";
import { IDate } from "../../../shared/types/idate";

export class Project {
  id: string;
  key: string;
  summary: string;
  type: ProjectType;
  priority: ProjectPriority;
  status: ProjectStatus;
  createdAt: IDate;
  createdBy: UserBase;

  constructor(data: ProjectDTO) {
    this.id = data.id;
    this.key = data.key;
    this.summary = data.summary;
    this.type = new ProjectType(data.type);
    this.priority = new ProjectPriority(data.priority);
    this.status = new ProjectStatus(data.status);
    this.createdAt = new IDate(data.createdAt);
    this.createdBy = new UserBase(data.createdBy);
  }

  toDTO(): ProjectDTO {
    return {
      id: this.id,
      key: this.key,
      summary: this.summary,
      type: this.type.toDTO(),
      priority: this.priority.toDTO(),
      status: this.status.toDTO(),
      createdAt: this.createdAt.msTimestamp,
      createdBy: this.createdBy.toDTO(),
    };
  }
}

export const maxKeyLength = 8;
export const maxSummaryLength = 128;
