interface WorkspaceInterface {
  id: string;
  name: string;
  hexColor: string;
}

class WorkspaceClass implements WorkspaceInterface {
  id: string;
  name: string;
  hexColor: string;
  constructor(item: WorkspaceInterface) {
    this.id = item.id;
    this.name = item.name;
    this.hexColor = item.hexColor;
  }
}

const maxNameLength = 16;

export { type WorkspaceInterface, WorkspaceClass, maxNameLength };
