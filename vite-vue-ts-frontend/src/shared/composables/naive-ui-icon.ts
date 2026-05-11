import { h, type Component } from "vue";

import { NIcon } from "naive-ui";

export const renderIcon = (icon: Component) => {
  return (size = 32) =>
    () =>
      h(
        NIcon,
        { size },
        {
          default: () => h(icon),
        },
      );
};
