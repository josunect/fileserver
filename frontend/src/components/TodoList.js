import React, { Component } from 'react'
import TodoItem from "./TodoItem";

import { connect } from "react-redux";
import { getTodos, findByText } from "../actions/todos";

class TodoList extends Component {

    constructor(props) {
        super(props);
        this.onChangeSearchText = this.onChangeSearchText.bind(this);
        this.refreshData = this.refreshData.bind(this);
        this.setActiveTodo = this.setActiveTodo.bind(this);
        this.findByText = this.findByText.bind(this);

        this.state = {
            currentTodo: null,
            currentIndex: -1,
            searchText: "",
        };
    }
    componentDidMount() {
        this.props.getTodos();
    }
    onChangeSearchText(e) {
        const searchText = e.target.value;
        this.setState({
            searchTitle: searchText,
        });
    }
    refreshData() {
        this.setState({
            currentTodo: null,
            currentIndex: -1,
        });
    }
    setActiveTodo(text, index) {
        this.setState({
            currentTodo: text,
            currentIndex: index,
        });
    }
    findByText() {
        this.refreshData();
        this.props.findByText(this.state.searchText);
    }

    render() {
        const { todos } = this.props;
        
        const listItems = todos.map((element, i) => {
            return (<TodoItem key={element.text} todo={element}/>);
        });

        return (
            <ul className="todoUl">
                {listItems}
            </ul>
        );
    }



};

const mapStateToProps = (state) => {
    return {
        todos: state.todos,
    };
};

export default connect(mapStateToProps, { getTodos, findByText })(TodoList);