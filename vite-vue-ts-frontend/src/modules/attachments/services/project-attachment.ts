import { axiosInstance } from "../../../api/client";

import type { SearchResponse } from "../types/dto";

export const projectAttachmentService = {
  async deleteProjectAttachment(
    projectId: string,
    attachmentId: string,
  ): Promise<void> {
    await axiosInstance.delete<void>(
      "/projects/" + projectId + "/attachments/" + attachmentId,
    );
  },
  async getProjectAttachments(projectId: string): Promise<SearchResponse> {
    const { data } = await axiosInstance.get<SearchResponse>(
      "/projects/" + projectId + "/attachments",
    );
    return data;
  },
};
