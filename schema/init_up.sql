--User Table
CREATE TABLE users (
    id SERIAL,
    name VARCHAR(50),
    login VARCHAR(50),
    password VARCHAR(255),
    PRIMARY KEY (id)
);

--TodoList

--List Table
CREATE TABLE lists (
    id SERIAL,
    title VARCHAR(50),
    description TEXT,
    PRIMARY KEY (id)
);

--Todo Table
CREATE TABLE todos (
    id SERIAL,
    title VARCHAR(100),
    description TEXT,
    importance BOOLEAN,
    done BOOLEAN,
    PRIMARY KEY (id)
);

--Users with Lists relations Table
CREATE TABLE users_lists (
    userId INT,
    listId INT,
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (listId) REFERENCES lists (id) ON DELETE CASCADE,
    PRIMARY KEY (userId, listId)
);

--Lists with Todos relations Table
CREATE TABLE lists_todos (
    listId INT,
    todoId INT,
    FOREIGN KEY (listId) REFERENCES lists (id) ON DELETE CASCADE,
    FOREIGN KEY (todoId) REFERENCES todos (id) ON DELETE  CASCADE,
    PRIMARY KEY (listId, todoId)
);