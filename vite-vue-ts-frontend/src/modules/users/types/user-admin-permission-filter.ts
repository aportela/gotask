export const UserPermissionFilterValue = {
  Any: 0,
  OnlyAdministrators: 1,
  OnlyUsers: 2,
} as const;

export type UserPermissionFilter =
  (typeof UserPermissionFilterValue)[keyof typeof UserPermissionFilterValue];
