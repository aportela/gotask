import { defineStore, acceptHMRUpdate } from "pinia";
import { createStorageEntry } from "../composables/localStorage";

const localStorageUserSettingsFluidLayout = createStorageEntry<boolean>(
  "userSettings.fluidLayout",
  false,
);

interface State {
  fluidLayout: boolean;
}

export const useUserSettingsStore = defineStore("userSettings", {
  state: (): State => ({
    fluidLayout: false,
  }),
  getters: {
    hasFluidLayout: (state): boolean => state.fluidLayout == true,
  },
  actions: {
    init(): void {
      this.fluidLayout = localStorageUserSettingsFluidLayout.get();
    },
    setFluidLayout(fluid: boolean): void {
      this.fluidLayout = fluid;
      localStorageUserSettingsFluidLayout.set(this.fluidLayout);
    },
    toggleFluidLayout(): void {
      this.setFluidLayout(!this.fluidLayout);
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(
    acceptHMRUpdate(useUserSettingsStore, import.meta.hot),
  );
}
