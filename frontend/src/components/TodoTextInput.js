import React, { Component } from 'react'
import { addTodo } from "../actions/todos";
import {connect} from "react-redux";

class TodoTextInput extends Component {

    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.newTodo = this.newTodo.bind(this);
        this.state = {
            id: null,
            text: "",
            completed: false,
        };
    }

    handleSubmit = e => {

        if (e.which === 13) {
            const {text, completed} = this.state;
            this.props.addTodo(text, completed).then((data) => {
                    this.setState({
                        id: data.id,
                        text: data.text,
                        completed: data.completed,
                    });
                    console.log("SUCCESS");
                    console.log(data);
            })
            .catch((ex) => {
                console.log("ERROR");
                console.log(ex);
           });
        }
    }

    newTodo() {
        this.setState({
           id: null,
           text: "",
           completed: false
        });
    }
    handleChange = e => {
        this.setState({ text: e.target.value })
    }

    render() {
        return (
            <input className={"newTodo"}
                   type="text"
                   placeholder={this.props.placeholder}
                   autoFocus="true"
                   value={this.state.text}
                   onChange={this.handleChange}
                   onKeyDown={this.handleSubmit} />
        )
    }
}

export default connect(null, { addTodo })(TodoTextInput);