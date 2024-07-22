## Login 

* Admin
curl -X POST http://localhost:4000/login -H 'Content-Type: application/json' -d '{"username": "admin","password": "adminpassword"}'
response : {
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUyMjY2NDAsInVzZXJfdHlwZSI6ImFkbWluIn0.iRIg5YLOz7su6M11agPP6ErciH1tKMcaO2OWAIvb1jY"
}

...........................
* User
curl -X POST http://localhost:4000/login -H 'Content-Type: application/json' -d '{"username": "regular","password": "regularpassword"}'
response : {
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUyMjc1MzcsInVzZXJfdHlwZSI6ImFkbWluIn0.Dg2dYZmYINBMOgfQrrPo84rIBhwhKY4Uj82HyvfgMRg"
}

............................

## View books  
curl -X GET http://localhost:4000/home -H 'Authorization: <YOUR_JWT_TOKEN>'

response : [
  {
    "name": "Book Name",
    "author": "Author",
    "publication_year": 0
  },
  {
    "name": "The Da Vinci Code",
    "author": "Dan Brown",
    "publication_year": 2003
  },
  {
    "name": "Think and Grow Rich",
    "author": "Napoleon Hill",
    "publication_year": 1937
  },
  {
    "name": "Harry Potter and the Half-Blood Prince",
    "author": "J.K. Rowling",
    "publication_year": 2005
  },
  {
    "name": "The Catcher in the Rye",
    "author": "J.D. Salinger",
    "publication_year": 1951
  },
  {
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publication_year": 1988
  },
  {
    "name": "Book Name",
    "author": "Author",
    "publication_year": 0
  },
  {
    "name": "Don Quixote",
    "author": "Miguel de Cervantes",
    "publication_year": 1605
  },
  {
    "name": "A Tale of Two Cities",
    "author": "Charles Dickens",
    "publication_year": 1859
  },
  {
    "name": "The Lord of the Rings",
    "author": "J.R.R. Tolkien",
    "publication_year": 1954
  },
  {
    "name": "The Little Prince",
    "author": "Antoine de Saint-Exupery",
    "publication_year": 1943
  },
  {
    "name": "Harry Potter and the Sorcerer’s Stone",
    "author": "J.K. Rowling",
    "publication_year": 1997
  },
  {
    "name": "And Then There Were None",
    "author": "Agatha Christie",
    "publication_year": 1939
  },
  {
    "name": "The Dream of the Red Chamber",
    "author": "Cao Xueqin",
    "publication_year": 1791
  },
  {
    "name": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "publication_year": 1937
  },
  {
    "name": "She: A History of Adventure",
    "author": "H. Rider Haggard",
    "publication_year": 1887
  },
  {
    "name": "The Lion, the Witch and the Wardrobe",
    "author": "C.S. Lewis",
    "publication_year": 1950
  },
  {
    "name": "my book",
    "author": "Nihal",
    "publication_year": 2024
  },
  {
    "name": "my book1",
    "author": "Nicks",
    "publication_year": 2024
  }
]

...........................

## Add Book (only admin)

curl -X POST http://localhost:4000/addBook -H 'Authorization: <YOUR_JWT_TOKEN>' -H 'Content-Type: application/json' -d '{"name": "New Book","author": "New Author","publication_year": 2022}'

..............................
## Delete Book(only admin)

curl -X POST http://localhost:4000/deleteBook -H 'Authorization: <YOUR_JWT_TOKEN>' -H 'Content-Type: application/json' -d 'New Book'
