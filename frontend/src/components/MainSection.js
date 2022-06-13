import React from "react";
import Footer from "./Footer";
import TodoList from "../components/TodoList";
import TodoTextInput from "../components/TodoTextInput";
import { style } from 'typestyle';

const newTodo = style ({
    padding: "10px",
    minWidth: "25em"
})

const MainSection = () => {

    return (
        <section className="main">
            <span>
                <TodoTextInput  />
            </span>
            <TodoList />
            <Footer />
        </section>
    );
};

export default MainSection;