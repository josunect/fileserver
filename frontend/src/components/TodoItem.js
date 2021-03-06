import React, { Component } from "react";
import TodoTextInput from "./TodoTextInput";
import { style } from 'typestyle';
import {deleteTodo, updateTodo} from "../actions/todos";
import {connect} from "react-redux";

const todoList = style ({
    listStyle: "none",
    padding: "10px"
})

class TodoItem extends Component {
    constructor(props) {
        super(props);
        this.state = {
            editing: false,
            element: props.todo
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.updateItem = this.updateItem.bind(this);
    }

    handleDoubleClick = () => {
        this.setState({ editing: true });
    }

    handleChange = (e) => {
        let item = this.state.element;
        item.completed = e.target.checked;
        this.updateItem(item.id, item)
    }

    handleSave = (e) => {
        if (e.which === 13) {
            let item = this.state.element;
            item.text = e.target.value;
            this.updateItem(item.id, item)
        }
    }

    updateItem = (id, item) => {
        this.props.updateTodo(item.id, item).then((data) => {
            this.setState({
                editing: false
            });
            console.log("SUCCESS");
            console.log(data);
        })
        .catch((ex) => {
                console.log("ERROR");
                console.log(ex);
        });
    }

    handleDelete = (e) => {
        const item = this.state.element;
        this.props.deleteTodo(item.id).then((data) => {
            console.log("SUCCESS");
            console.log(data);
        })
        .catch((ex) => {
            console.log("ERROR");
            console.log(ex);
        });
    }

    render() {
        const { todo } = this.props;

        let element;
        if (this.state.editing) {
            element = (
                <input type="text"
                    text={todo.text}
                    placeholder={todo.text}
                    onKeyDown={this.handleSave}
                />
            );
        } else {
            element = (
                <div className="view">
                    <input
                        className="toggle"
                        type="checkbox"
                        checked={todo.completed}
                        onChange={this.handleChange}
                    />
                    <label onDoubleClick={this.handleDoubleClick}>{todo.text}</label>
                    <button className="destroy" onClick={this.handleDelete} > X </button>
                </div>
            );
        }

        return (
            <li className={todoList} >
                {element}
            </li>
        );
    }
}

export default connect(null, { updateTodo, deleteTodo })(TodoItem);