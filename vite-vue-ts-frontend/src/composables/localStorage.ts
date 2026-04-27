import { LOCAL_STORAGE_NAMESPACE } from "../constants";

export type StorageValue = string | number | boolean | null | object;

const createStorageEntry = <T extends StorageValue>(
  key: string,
  defaultValue: T,
) => ({
  get(): T {
    try {
      const stored = localStorage.getItem(LOCAL_STORAGE_NAMESPACE + key);
      return stored === null ? defaultValue : (stored as T);
    } catch (error) {
      console.error("Error accessing localStorage:", error);
      return defaultValue;
    }
  },
  set(value: T) {
    try {
      localStorage.set(
        LOCAL_STORAGE_NAMESPACE + key,
        value as unknown as StorageValue,
      );
    } catch (error) {
      console.error("Error accessing localStorage:", error);
    }
  },
  remove() {
    try {
      localStorage.remove(LOCAL_STORAGE_NAMESPACE + key);
    } catch (error) {
      console.error("Error accessing localStorage:", error);
    }
  },
});

export { createStorageEntry };
