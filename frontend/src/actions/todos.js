import {
    ADD_TODO,
    GET_TODOS,
    UPDATE_TODO,
    DELETE_TODO
} from "./types";

import TodoDataService from "../services/todo.service";

export const addTodo = (text, completed) => async (dispatch) => {
    try {
        const res = await TodoDataService.create({ text, completed });
        dispatch({
            type: ADD_TODO,
            payload: res.data,
        });
        return Promise.resolve(res.data);
    } catch (err) {
        return Promise.reject(err);
    }
};

export const getTodos = () => async (dispatch) => {
    try {
        const res = await TodoDataService.getAll();
        dispatch({
            type: GET_TODOS,
            payload: res.data,
        });
    } catch (err) {
        console.log(err);
    }
};
export const updateTodo = (id, data) => async (dispatch) => {
    try {
        const res = await TodoDataService.update(id, data);
        dispatch({
            type: UPDATE_TODO,
            payload: data,
        });
        return Promise.resolve(res.data);
    } catch (err) {
        return Promise.reject(err);
    }
};
export const deleteTodo = (id) => async (dispatch) => {
    try {
        await TodoDataService.delete(id);
        dispatch({
            type: DELETE_TODO,
            payload: { id },
        });
    } catch (err) {
        console.log(err);
    }
};

export const findByText = (text) => async (dispatch) => {
    try {
        const res = await TodoDataService.findByText(text);
        dispatch({
            type: GET_TODOS,
            payload: res.data,
        });
    } catch (err) {
        console.log(err);
    }
};
