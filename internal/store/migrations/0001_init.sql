CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    password_hash VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS quizzes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS games (
    id VARCHAR(40) PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    is_started BOOLEAN DEFAULT FALSE,
    quiz_id INT,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);

CREATE TABLE IF NOT EXISTS images(
    id SERIAL PRIMARY KEY,
    image_url VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS questions(
    id SERIAL PRIMARY KEY,
    text_content VARCHAR(200) NOT NULL,
    position INT NOT NULL,
    quiz_id INT,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id),
    image_id INT,
    FOREIGN KEY (image_id) REFERENCES images(id)
);

CREATE TABLE IF NOT EXISTS answers(
    id SERIAL PRIMARY KEY,
    text_content VARCHAR(200) NOT NULL,
    is_correct BOOLEAN,
    quiz_id INT,
    FOREIGN KEY(quiz_id) REFERENCES questions(id)
);