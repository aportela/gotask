interface WorkspaceInterface {
  id: string;
  name: string;
}

class WorkspaceClass implements WorkspaceInterface {
  id: string;
  name: string;
  constructor(item: WorkspaceInterface) {
    this.id = item.id;
    this.name = item.name;
  }
}

const maxNameLength = 16;

export { type WorkspaceInterface, WorkspaceClass, maxNameLength };
