export const UserAdminPermissionFilterValue = {
  All: 0,
  OnlyAdministrators: 1,
  OnlyUsers: 2,
} as const;

export type UserAdminPermissionFilter =
  (typeof UserAdminPermissionFilterValue)[keyof typeof UserAdminPermissionFilterValue];
