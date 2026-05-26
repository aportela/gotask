interface UserInterface {
  id: string;
  name: string;
  password?: string;
  email: string;
  isSuperUser: boolean;
  createdAt: number;
  updatedAt: number | null;
  deletedAt: number | null;
}

class UserClass implements UserInterface {
  id: string;
  name: string;
  password?: string;
  email: string;
  isSuperUser: boolean;
  createdAt: number;
  updatedAt: number | null;
  deletedAt: number | null;

  constructor(item: UserInterface) {
    this.id = item.id;
    this.name = item.name;
    this.password = item.password;
    this.email = item.email;
    this.isSuperUser = item.isSuperUser;
    this.createdAt = item.createdAt;
    this.updatedAt = item.updatedAt;
    this.deletedAt = item.deletedAt;
  }
}

const maxNameLength = 32;
const maxEmailLength = 255;

export { type UserInterface, UserClass, maxNameLength, maxEmailLength };
