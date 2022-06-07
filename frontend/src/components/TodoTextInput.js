import React, { Component } from 'react'


export default class TodoTextInput extends Component {

    state = {
        text: this.props.text || ''
    }
    newTodo;

    handleSubmit = e => {
        const text = e.target.value.trim()
        if (e.which === 13) {
            this.props.onSave(text)
            if (this.props.newTodo) {
                this.setState({ text: '' })
            }
        }
    }

    handleChange = e => {
        this.setState({ text: e.target.value })
    }

    handleBlur = e => {
        if (!this.props.newTodo) {
            this.props.onSave(e.target.value)
        }
    }

    render() {
        return (
            <input className={"todoInput"}
                   type="text"
                   placeholder={this.props.placeholder}
                   autoFocus="true"
                   value={this.state.text}
                   onBlur={this.handleBlur}
                   onChange={this.handleChange}
                   onKeyDown={this.handleSubmit} />
        )
    }
}