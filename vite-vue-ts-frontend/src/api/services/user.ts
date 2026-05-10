import { axiosInstance } from "../client";
import type {
  AddRequestInterface,
  AddResponseInterface,
  UpdateRequestInterface,
  SearchRequestInterface,
  UpdateResponseInterface,
  GetResponseInterface,
  SearchResponseInterface,
} from "../types/dto/user";

export const authService = {
  async add(payload: AddRequestInterface): Promise<AddResponseInterface> {
    const { data } = await axiosInstance.post<AddResponseInterface>(
      "/users",
      payload,
    );
    return data;
  },
  async update(
    payload: UpdateRequestInterface,
  ): Promise<UpdateResponseInterface> {
    const { data } = await axiosInstance.post<UpdateResponseInterface>(
      "/users/" + payload.user.id,
      payload,
    );
    return data;
  },
  async delete(id: string): Promise<void> {
    await axiosInstance.delete<void>("/users/" + id);
  },
  async unDelete(id: string): Promise<void> {
    const params = {
      deletedAt: null,
    };
    await axiosInstance.patch<void>("/users/" + id, params);
  },
  async get(id: string): Promise<GetResponseInterface> {
    const { data } = await axiosInstance.get<GetResponseInterface>(
      "/users/" + id,
    );
    return data;
  },
  async search(
    params?: SearchRequestInterface,
  ): Promise<SearchResponseInterface> {
    const { data } = await axiosInstance.get<SearchResponseInterface>(
      "/users",
      { params },
    );
    return data;
  },
};
