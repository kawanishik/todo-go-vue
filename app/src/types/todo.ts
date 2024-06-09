export type Todo = {
  Id: number;
  Title: string;
  Description: string;
  Completed: boolean;
  Created_at: string;
  Updated_at: string;
};

export type AddTodo = {
  Title: string;
  Description: string;
};

export type EditTodo = {
  Id: number;
  Title: string;
  Description: string;
};