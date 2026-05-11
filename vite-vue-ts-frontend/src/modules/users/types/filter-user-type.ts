export const UserTypeFilterValue = {
  None: 0,
  OnlyAdministrators: 1,
  OnlyUsers: 2,
} as const;

export type UserTypeFilter =
  (typeof UserTypeFilterValue)[keyof typeof UserTypeFilterValue];
