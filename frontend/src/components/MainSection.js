import React from "react";
import Footer from "./Footer";
import TodoList from "../components/TodoList";
import { style } from 'typestyle';

const newTodo = style ({
    padding: "10px",
    minWidth: "25em"
})

const MainSection = () => {

    return (
        <section className="main">
            <span>
                <input
                    className={newTodo}
                    type="text"
                    placeholder="What needs to be done?"
                />
            </span>
            <TodoList />
            <Footer />
        </section>
    );
};

export default MainSection;