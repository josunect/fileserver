import {
    ADD_TODO,
    GET_TODOS,
    UPDATE_TODO,
    DELETE_TODO
} from "../actions/types";

const initialState = [];

function todoReducer(todos = initialState, action) {
    const { type, payload } = action;
    switch (type) {
        case ADD_TODO:
            return [...todos, payload];
        case GET_TODOS:
            return payload;
        case UPDATE_TODO:
            return todos.map((tutorial) => {
                if (tutorial.id === payload.id) {
                    return {
                        ...tutorial,
                        ...payload,
                    };
                } else {
                    return tutorial;
                }
            });
        case DELETE_TODO:
            return todos.filter(({ id }) => id !== payload.id);
        default:
            return todos;
    }
};

export default todoReducer;
