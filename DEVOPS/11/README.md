# Anecdote API

This is a simple RESTful API for managing anecdotes using Go and the Gin framework. The API allows you to create, read, update, patch, and delete anecdotes, as well as like and dislike them.

## Endpoints

### 1. Get All Anecdotes

Retrieve a list of all anecdotes.

curl -X GET http://localhost:8080/anecdotes

### 2. Get Anecdote by ID

Retrieve a specific anecdote by its ID.

curl -X GET http://localhost:8080/anecdotes/:id

### 3. Post a New Anecdote

Add a new anecdote.

curl -X POST http://localhost:8080/anecdotes 
-H "Content-Type: application/json" 
-d '{"id": "4", "title": "Новый анекдот", "author": "Автор", "text": "Текст анекдота", "likes": 0, "dislikes": 0}'
### 4. Update an Anecdote

Update an existing anecdote by its ID.

curl -X PUT http://localhost:8080/anecdotes/:id 

-H "Content-Type: application/json" 

-d '{"id": "1", "title": "Обновленный анекдот", "author": "Автор", "text": "Обновленный текст", "likes": 10, "dislikes": 0}'

### 5. Patch an Anecdote

Partially update an existing anecdote by its ID.

curl -X PATCH http://localhost:8080/anecdotes/:id 

-H "Content-Type: application/json" 

-d '{"title": "Частично обновленный анекдот"}'

### 6. Delete an Anecdote by ID

Delete an anecdote by its ID.

curl -X DELETE http://localhost:8080/anecdotes/:id

### 7. Like an Anecdote

Like an anecdote by its ID.

curl -X POST http://localhost:8080/anecdotes/:id/like

### 8. Dislike an Anecdote

Dislike an anecdote by its ID.

curl -X POST http://localhost:8080/anecdotes/:id/dislike

