import http from "@/http-common";
import type { Todo, AddTodo, EditTodo } from "@/types/todo";

class TodoRepository {
  getTodoList(): Promise<Todo[] | null> {
    return http.get("/index")
      .then((response) => {
        return response.data as Todo[];
      });
  }

  async addTodo(todo: AddTodo) {
    await http.post("/create", todo)
    .then((res) => {
      return res;
    })
    .catch((error) => {
      console.error(error);
    });
    return Promise.resolve();
  }

  async editTodo(todo: EditTodo) {
    await http.put("/edit", todo)
    .then((res) => {
      return res;
    })
    .catch((error) => {
      console.error(error);
    });
    return Promise.resolve();
  }
}

export default new TodoRepository();