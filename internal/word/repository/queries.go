package repository

const createWord = `insert into words (id, text) VALUES ($1, $2)`

const suggestionsWord = `SELECT * FROM words WHERE text LIKE $1`

const selectArray = `SELECT description, url FROM videos WHERE vector_ids && $1`
