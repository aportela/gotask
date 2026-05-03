import { type NotificationType, useNotification } from "naive-ui";

export const useNotify = () => {
  const notification = useNotification();

  const notify = (type: NotificationType, content: string, meta: string) => {
    notification[type]({
      content,
      meta,
      duration: 2500,
      keepAliveOnHover: true,
    });
  };

  return { notify };
};
