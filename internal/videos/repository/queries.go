package repository

const selectArray = `SELECT description, url FROM videos WHERE related_vectors && $1`
