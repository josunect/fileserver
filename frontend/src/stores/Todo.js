export const TodoStore = {
    todos: [{
        text: 'Comer',
        completed: true
    }, {
        text: 'Hacer ejercicio',
        completed: false
    }],
    visibilityFilter: 'SHOW_COMPLETED',
    id: "myTodoId"
};

export const addTodo = text =>
    TodoStore.set(
        ({ todos }) => ({
            todos: [
                ...todos,
                {
                    id: todos.reduce((maxId, todo) => Math.max(todo.id, maxId), -1) + 1,
                    completed: false,
                    text
                }
            ]
        }),
        "Add todo"
    );

export const deleteTodo = id =>
    TodoStore.set(
        ({ todos }) => ({
            todos: todos.filter(item => item.id !== id)
        }),
        "Delete todo"
    );

export const editTodo = (id, text) =>
    TodoStore.set(
        ({ todos }) => ({
            todos: todos.map(todo => (todo.id === id ? { ...todo, text } : todo))
        }),
        "Edit todo"
    );

export const completeTodo = id =>
    TodoStore.set(
        ({ todos }) => ({
            todos: todos.map(todo =>
                todo.id === id ? { ...todo, completed: !todo.completed } : todo
            )
        }),
        "Complete todo"
    );
