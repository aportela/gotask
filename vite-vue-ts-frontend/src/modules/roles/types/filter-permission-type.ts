export const PermissionTypeFilterValue = {
  All: 0,
  Allowed: 1,
  NotAllowed: 2,
} as const;

export type PermissionTypeFilter =
  (typeof PermissionTypeFilterValue)[keyof typeof PermissionTypeFilterValue];
