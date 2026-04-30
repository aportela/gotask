import { defineStore, acceptHMRUpdate } from "pinia";

interface State {
  loading: boolean;
}

export const useLoadingStore = defineStore("loadingStore", {
  state: (): State => ({
    loading: false,
  }),
  getters: {
    isLoading: (state): boolean => state.loading === true,
  },
  actions: {
    set(value: boolean): void {
      this.loading = value;
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useLoadingStore, import.meta.hot));
}
