import React, { Component } from 'react'
import TodoItem from "./TodoItem";


class TodoList extends Component {

    render() {
        const { todos } = this.props.list;

        const listItems = todos.todos.map((element, i) => {
            return (<TodoItem key={element.text} todo={element}/>);
        });

        return (
            <ul className="todoUl">
                {listItems}
            </ul>
        );
    }
};

export default TodoList;