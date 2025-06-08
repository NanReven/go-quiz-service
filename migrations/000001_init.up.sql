CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE quizzes (
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TYPE question_type AS ENUM ('single_choice', 'multiple_choice', 'text');

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    question_type question_type NOT NULL
);

CREATE TABLE quiz_scores (
    id SERIAL PRIMARY KEY,
    quiz_id INTEGER NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    start_datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_datetime TIMESTAMP NULL,
    score INTEGER NOT NULL DEFAULT 0,
    CHECK (score >= 0)
);

CREATE TABLE quiz_questions (
    quiz_id INTEGER NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    PRIMARY KEY (quiz_id, question_id),
    position INTEGER NOT NULL DEFAULT 0,
    CHECK (position >= 0)
);

CREATE TABLE options (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE question_options (
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    option_id INTEGER NOT NULL REFERENCES options(id) ON DELETE CASCADE,
    PRIMARY KEY (question_id, option_id),
    is_correct BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE user_answers_choices (
    id SERIAL PRIMARY KEY,
    score_id INTEGER NOT NULL REFERENCES quiz_scores(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    option_id INTEGER NOT NULL REFERENCES options(id) ON DELETE CASCADE
);

CREATE TABLE user_answers_texts (
    id SERIAL PRIMARY KEY,
    score_id INTEGER NOT NULL REFERENCES quiz_scores(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    answer TEXT NOT NULL
);

CREATE TABLE correct_text_answers (
    question_id INTEGER PRIMARY KEY REFERENCES questions(id) ON DELETE CASCADE,
    answer TEXT NOT NULL
);