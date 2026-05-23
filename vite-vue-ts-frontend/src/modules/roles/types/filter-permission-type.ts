export const RolePermissionFilterValue = {
  Any: 0,
  Allowed: 1,
  NotAllowed: 2,
} as const;

export type RolePermissionFilter =
  (typeof RolePermissionFilterValue)[keyof typeof RolePermissionFilterValue];
