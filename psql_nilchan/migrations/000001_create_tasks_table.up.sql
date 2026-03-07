CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	description VARCHAR(1000) NOT NULL,
	completed BOOLEAN NOT NULL,
	created_at TIMESTAMP NOT NULL,
	completed_at TIMESTAMP
);
