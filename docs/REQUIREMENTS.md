# **Task: REST API with Go (Gin) \- Books Management System**

## **Objective:**

Build and host a **REST API** using **Go (Gin framework)** to manage a **"Book" entity**, incorporating **Swagger documentation**, **Kafka for event streaming**, and **Redis for caching**. The API should be hosted so that **Swagger UI** is accessible via a URL. Provide the GitHub repository link along with the Swagger file.

---

## **API Requirements**

* **Gin Framework** for routing.  
* **Swagger Documentation** using `swaggo/swag`.  
* **Kafka Integration** to publish book-related events (e.g., book creation, update, deletion).  
* **Redis Caching** to optimize book retrieval performance.  
* **Database**: Use **PostgreSQL** (or an in-memory database like SQLite for simplicity).  
* **Hosting**: Deploy on **Heroku, AWS, or any cloud provider**.  
* **Best Practices**: Use validation, logging, and proper error handling.

---

## **Book Model**

| Field | Type | Description |
| ----- | ----- | ----- |
| ID | integer | Auto-incremented identifier |
| Title | string | Book title |
| Author | string | Author's name |
| Year | integer | Year of publication |

---

## **API Endpoints**

1. **GET /books** \- Retrieve a list of all books (cached using Redis).  
2. **GET /books/{id}** \- Retrieve a specific book by its ID (cached using Redis).  
3. **POST /books** \- Create a new book and publish an event to Kafka.  
4. **PUT /books/{id}** \- Update an existing book’s details and publish an update event to Kafka.  
5. **DELETE /books/{id}** \- Delete a book by its ID and publish a deletion event to Kafka.

---

## **Additional Enhancements**

### **Kafka Integration**

* Every **POST, PUT, DELETE** request should publish an event to a Kafka topic (`book_events`).  
* Consumers can later be added to handle these events (e.g., logging, analytics, notifications).

### **Redis Caching**

* Cache book data to reduce database queries.  
* **GET /books** and **GET /books/{id}** should first check Redis before querying the database.  
* Expire cache entries when a **POST, PUT, DELETE** operation is performed.

### **Pagination**

* Implement **pagination** for the `GET /books` endpoint using **limit and offset**.

### **Validation**

* Ensure the **title is not empty**, **author is not empty**, and **year is a valid number**.

### **Proper HTTP Status Codes**

* `200` for successful retrieval.  
* `201` for successful creation.  
* `400` for validation errors.  
* `404` if a book is not found.  
* `500` for internal server errors.

---

## **Setup and Deployment Steps**

### **1\. Project Initialization**

* Initialize a Go project.  
* Install **Gin, Swaggo, Kafka Go Client, Redis Client, GORM** for PostgreSQL.

### **2\. Implement API Endpoints**

* Set up controllers, services, and repository layers.  
* Use **GORM** for database operations.

### **3\. Add Kafka Producer & Consumer**

* **Producer**: Publish book events when changes occur.  
* **Consumer (Optional)**: Log received events.

### **4\. Integrate Redis Caching**

* Store frequently accessed books.  
* Implement cache invalidation on updates.

### **5\. Generate Swagger Documentation**

* Use `swag init` to generate Swagger docs.  
* Serve Swagger UI at `/swagger`.

### **6\. Deployment**

* Deploy API using **Heroku, AWS, or another cloud provider**.  
* Ensure the Swagger UI and API are publicly accessible.

### **7\. GitHub Repository**

* Push code to GitHub.  
* Include **README** with setup instructions.

---

## **Expected Deliverables**

1. **GitHub Repository Link**:

   * Provide a link to the GitHub repo containing the full project.  
2. **Swagger Documentation URL**:

   * Provide a public URL where Swagger UI is hosted.  
3. **Public API Endpoints**:

   * Ensure endpoints are accessible for testing.  
4. **Postman Collection / cURL Commands**:

   * Share API testing scripts.

---

## **Tools & Technologies**

* **Go** (Gin Framework)  
* **PostgreSQL** (Primary Database)  
* **Redis** (Caching)  
* **Kafka** (Event Streaming)  
* **Swagger** (API Documentation)  
* **Docker** (Optional for local development)