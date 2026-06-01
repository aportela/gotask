import type { UserBaseResponse } from "../../users/types/dto";

export type AttachmentResponse = {
  id: string;
  createdBy: UserBaseResponse;
  createdAt: number;
  name: string;
  contentType: string;
  size: number;
};

export type SearchResponse = {
  attachments: AttachmentResponse[];
};
