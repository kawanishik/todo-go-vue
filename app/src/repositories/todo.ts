import http from "@/http-common";
import type { Todo } from "@/types/todo";

class TodoRepository {
  getTodoList(): Promise<Todo[] | null> {
    return http.get("/index")
      .then((response) => {
        return response.data as Todo[];
      });
  }
}

export default new TodoRepository();