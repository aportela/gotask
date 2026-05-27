import { axiosInstance } from "../../../api/client";

import type { AddRequest, NoteResponse, SearchResponse } from "../types/dto";

export const noteService = {
  async add(projectId: string, payload: AddRequest): Promise<NoteResponse> {
    const { data } = await axiosInstance.post<NoteResponse>(
      "/projects/" + projectId + "/notes/",
      payload,
    );
    return data;
  },
  async delete(projectId: string, noteId: string): Promise<void> {
    await axiosInstance.delete<void>(
      "/projects/" + projectId + "/notes/" + noteId,
    );
  },
  async search(projectId: string): Promise<SearchResponse> {
    const { data } = await axiosInstance.get<SearchResponse>(
      "/projects/" + projectId + "/notes",
    );
    return data;
  },
};
