package schema

const Users string = `
CREATE TABLE IF NOT EXISTS users (
	id varchar(50) not null PRIMARY KEY, 
	name varchar(100) not null
);
`

const Todo string = `
CREATE TABLE IF NOT EXISTS todo (
	id varchar(50) PRIMARY KEY,
	todo_text text not null,
	done boolean not null default false,
	user_id varchar(50) not null,
	FOREIGN KEY (user_id) REFERENCES users(id)
);`
