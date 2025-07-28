# `api`

This package holds the API service of the smartik marking system.

## Overview

The smartik API service is built with:
- **Go** - Programming language
- **Echo** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Database
- **Go-Nanoid** - Collision-resistant unique identifiers

## Prerequisites

> These services will be started automatically using docker when the development server is started. See the [air config](./.air.toml)

- Minio object storage
- PostgreSQL database

## Environment

| Name | Default | Description |
| :--- | :--- | :--- |
| GO_ENV | 'development' | The environment to be optimized for when run. |
| SERVER_URL | 'http://localhost:1323' | The full address of the machine the API will run on |
| PORT | '1323' | The port to listen for requests on |
| POSTGRES_URI | 'postgresql://root:123456@localhost:5432/postgres' | A uri of a running postgres database |
| MINIO_ENDPOINT_URL | 'localhost:9000' | The host & port number where minio will listen for connections |
| MINIO_ACCESS_ID | 'minioadmin' | An access ID used to programmatically access a running instance of Minio |
| MINIO_SECRET_KEY | 'minioadmin' | A secret key used to programmatically authenticate with a running instance of Minio |
| MINIO_STORAGE_BUCKET | 'smartik' | The name of the storage bucket where scripts will be stored |

## Port Mapping

When the development server is running, you will have access to:

| Service | Port | Default |
| :--- | :--- | :--- |
| **API** | The value you used for the `PORT` environment variable | `:1323` |
| **PostgrSQL** | - | `:5432` |
| **Minio** | - | `:9000`, `9001` | 

## Getting Started

For the complete development configuration, see the [getting started guide](../../docs/getting-started.md).

## API Reference

### Base URL

```
http://localhost:PORT/api/v1
```

> `PORT` will be replaced by the value you chose if you modified the variable, otherwise it will stick with the default `1323`.

##### **ANY `/health`**

**Response (200 OK)**
```json
{
  "status": "healthy",
  "time":   "time-of-request",              // In ISO format
  "uptime": "n Hours, n Minutes, n Seconds" // 'n' Being the repective value
}
```

##### **Any `/reference`**

**Response (200 OK)**
```json
{
  "message": "Available Routes Reference",
  "routes": [
    {
      "method": "GET",          // Method the endpoint accepts
      "path":   "/reference",   // The path of the endpoint
      "name":   "api_reference" // Name of the endpoint
    }
  ]
}
```

#### Students

##### **POST `/api/v1/students/create`**

**Request Body:**
```json
{
  "exam_number":  "string",
  "first_name":   "string",    // optional
  "last_name":    "string"     // optional 
}
```

**Response (201 Created):**
```json
{
  "message": "Student created successfully",
  "student": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "first_name": "John",
    "last_name": "Doe",
    "exam_number": "JD2025001"
  }
}
```

#### **GET `/api/v1/students`**

**Response (200 OK):**
```json
{
  "message": "Students retrieved successfully",
  "students": [
    {
      "id": "cmddih9m9000097hndiy6afpx",
      "CreatedAt": "2025-07-22T10:30:00Z",
      "UpdatedAt": "2025-07-22T10:30:00Z", 
      "first_name": "John",
      "last_name": "Doe",
      "exam_number": "JD2025001"
    }
  ]
}
```

#### **GET `/api/v1/students/{id}`**


**Path Parameters:**
- `id` (string) - The student's ID in the database

**Response (200 OK):**
```json
{
  "message": "Student retrieved successfully",
  "student": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "first_name": "John",
    "last_name": "Doe",
    "exam_number": "JD2025001"
  }
}
```

#### **PATCH `/api/v1/students/update/{id}`**

**Path Parameters:**
- `id` (string) - The student's ID in the database

**Request Body:**
```json
{
  "first_name":  "string",  // optional
  "last_name":   "string",  // optional
  "exam_number": "string"   // optional
}
```

**Response (200 OK):**
```json
{
  "message": "Student updated successfully",
  "student": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:35:00Z",
    "first_name": "Jane",
    "last_name": "Smith",
    "exam_number": "JD2025001"
  }
}
```

#### **DELETE `/api/v1/students/delete/{id}`**

**Path Parameters:**
- `id` (string) - The student's ID in the database

**Response (204 No Content):**

#### Errors

**Error Response (404 Not Found):**
```json
{
  "message": "Student not found"
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "message": "Failed to create student"
}
```

---

#### Subjects

##### **POST `/api/v1/subjects/create`**

**Request Body:**
```json
{
  "name": "string",
  "code": "string",
}
```

**Response (201 Created):**
```json
{
  "message": "Subject created successfully",
  "subject": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "name": "Calculus 1",
    "code": "CAL1502"
  }
}
```

#### **GET `/api/v1/subjects`**

**Response (200 OK):**
```json
{
  "message": "Subjects retrieved successfully",
  "subjects": [
    {
      "id": "cmddih9m9000097hndiy6afpx",
      "CreatedAt": "2025-07-22T10:30:00Z",
      "UpdatedAt": "2025-07-22T10:30:00Z",
      "name": "Calculus 1",
      "code": "CAL1502"
    }
  ]
}
```

#### **GET `/api/v1/subjects/{id}`**


**Path Parameters:**
- `id` (string) - The subject's ID in the database

**Response (200 OK):**
```json
{
  "message": "Subject retrieved successfully",
  "subject": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "name": "Calculus 1",
    "code": "CAL1502"
  }
}
```

#### **PATCH `/api/v1/subjects/update/{id}`**

**Path Parameters:**
- `id` (string) - The subject's ID in the database

**Request Body:**
```json
{
  "name": "string",  // optional
  "code": "string",  // optional
}
```

**Response (200 OK):**
```json
{
  "message": "Subject updated successfully",
  "subject": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "name": "Calculus 2",
    "code": "CAL2202"
  }
}
```

#### **DELETE `/api/v1/subjects/delete/{id}`**

**Path Parameters:**
- `id` (string) - The subject's ID in the database

**Response (204 No Content):**


**Error Response (404 Not Found):**
```json
{
  "message": "Subject not found"
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "message": "Failed to create subject"
}
```

---

#### Exams

##### **POST `/api/v1/exams/create`**

**Request Body:**
```json
{
  "date": "string", // Must be in ISO format
}
```

**Response (201 Created):**
```json
{
  "message": "Exam created successfully",
  "exam": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "date": "2025-07-22T10:30:00Z"
  }
}
```

#### **GET `/api/v1/exams`**

**Response (200 OK):**
```json
{
  "message": "Exams retrieved successfully",
  "exams": [
    {
      "id": "cmddih9m9000097hndiy6afpx",
      "CreatedAt": "2025-07-22T10:30:00Z",
      "UpdatedAt": "2025-07-22T10:30:00Z",
      "date": "2025-07-22T10:30:00Z"
    }
  ]
}
```

#### **GET `/api/v1/exams/{id}`**


**Path Parameters:**
- `id` (string) - The exam's ID in the database

**Response (200 OK):**
```json
{
  "message": "Exam retrieved successfully",
  "exam": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "date": "2025-07-22T10:30:00Z"
  }
}
```

#### **PATCH `/api/v1/exams/update/{id}`**

**Path Parameters:**
- `id` (string) - The exam's ID in the database

**Request Body:**
```json
{
  "date": "2025-07-22T10:30:00Z" // Must be a string in ISO 8601 format
}
```

**Response (200 OK):**
```json
{
  "message": "Exam updated successfully",
  "exam": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "date": "2025-09-22T10:30:00Z"
  }
}
```

#### **DELETE `/api/v1/exams/delete/{id}`**

**Path Parameters:**
- `id` (string) - The exam's ID in the database

**Response (204 No Content):**


**Error Response (404 Not Found):**
```json
{
  "message": "Exam not found"
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "message": "Failed to create exam"
}
```

---

#### Answer Scripts

##### **POST `/api/v1/scripts/upload`**

**Request:** Multipart form data with file upload

**Form Fields:**
- `answer_scripts` (file[]) - Array of answer script files to upload

**Response (200 OK):**
```json
{
  "message": "Answer scripts uploaded successfully",
  "count": 2,
  "answer_scripts": [
    {
      "id": "cmddih9m9000097hndiy6afpx",
      "CreatedAt": "2025-07-22T10:30:00Z",
      "UpdatedAt": "2025-07-22T10:30:00Z",
      "file_name": "exam_script_001.pdf",
      "file_url": null,
      "student_id": null,
      "subject_id": null,
      "exam_id": null,
      "total_marks": null,
      "obtained_marks": null,
      "scanned_exam_number": null,
      "confidence_score": null,
      "matched_at": null,
      "processing_status": "uploaded"
    }
  ]
}
```

**Response (207 Multi-Status):**
```json
{
  "message": "Some answer scripts failed to upload",
  "errors": {
    "count": 1,
    "file_name": {
      "filename": "corrupted_file.pdf",
      "error": "Failed to upload to MinIO storage"
    }
  }
}
```

#### **GET `/api/v1/scripts`**

**Response (200 OK):**
```json
{
  "message": "Answer scripts retrieved successfully",
  "answer_scripts": [
    {
      "id": "cmddih9m9000097hndiy6afpx",
      "CreatedAt": "2025-07-22T10:30:00Z",
      "UpdatedAt": "2025-07-22T10:30:00Z",
      "file_name": "exam_script_001.pdf",
      "file_url": null,
      "student_id": "student_123",
      "subject_id": "subject_456",
      "exam_id": "exam_789",
      "total_marks": 100,
      "obtained_marks": 85,
      "scanned_exam_number": "12345",
      "confidence_score": 0.95,
      "matched_at": "2025-07-22T10:35:00Z",
      "processing_status": "uploaded"
    }
  ]
}
```

#### **GET `/api/v1/scripts/{id}`**

**Path Parameters:**
- `id` (string) - The answer script's ID in the database

**Response (200 OK):**
```json
{
  "message": "Answer script retrieved successfully",
  "answer_script": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:30:00Z",
    "file_name": "exam_script_001.pdf",
    "file_url": null,
    "student_id": "student_123",
    "subject_id": "subject_456",
    "exam_id": "exam_789",
    "total_marks": 100,
    "obtained_marks": 85,
    "scanned_exam_number": "12345",
    "confidence_score": 0.95,
    "matched_at": "2025-07-22T10:35:00Z",
    "processing_status": "uploaded"
  }
}
```

#### **GET `/api/v1/scripts/serve/{id}`**

**Path Parameters:**
- `id` (string) - The answer script's ID in the database

**Response (200 OK):**
Returns the actual file content with appropriate headers for inline viewing or download.

**Headers:**
- `Content-Type`: Original file MIME type
- `Content-Disposition`: `inline; filename="original_filename.pdf"`

#### **PATCH `/api/v1/scripts/update/{id}`**

**Path Parameters:**
- `id` (string) - The answer script's ID in the database

**Request Body:**
```json
{
  "file_name": "string",          // optional
  "file_url": "string",           // optional
  "student_id": "string",         // optional
  "subject_id": "string",         // optional
  "exam_id": "string",            // optional
  "total_marks": 100,             // optional
  "obtained_marks": 85,           // optional
  "scanned_exam_number": "string", // optional
  "confidence_score": 0.95,       // optional
  "matched_at": "2025-07-22T10:35:00Z" // optional, ISO format
}
```

**Response (200 OK):**
```json
{
  "message": "Answer script updated successfully",
  "answer_script": {
    "id": "cmddih9m9000097hndiy6afpx",
    "CreatedAt": "2025-07-22T10:30:00Z",
    "UpdatedAt": "2025-07-22T10:40:00Z",
    "file_name": "updated_script.pdf",
    "file_url": null,
    "student_id": "student_456",
    "subject_id": "subject_789",
    "exam_id": "exam_123",
    "total_marks": 100,
    "obtained_marks": 90,
    "scanned_exam_number": "54321",
    "confidence_score": 0.98,
    "matched_at": "2025-07-22T10:40:00Z",
    "processing_status": "uploaded"
  }
}
```

#### **DELETE `/api/v1/scripts/delete/{id}`**

**Path Parameters:**
- `id` (string) - The answer script's ID in the database

**Response (204 No Content):**

#### Errors

**Error Response (400 Bad Request):**
```json
{
  "message": "Invalid multipart form data",
  "error": "request Content-Type isn't multipart/form-data"
}
```

```json
{
  "message": "No answer scripts provided"
}
```

**Error Response (404 Not Found):**
```json
{
  "message": "Answer script not found"
}
```

**Error Response (500 Internal Server Error):**
```json
{
  "message": "Failed to save answer script record"
}
```

```json
{
  "message": "Failed to retrieve answer scripts"
}
```

```json
{
  "message": "Failed to retrieve answer script file"
}
```

```json
{
  "message": "Failed to update answer script"
}
```

```json
{
  "message": "Failed to delete answer script"
}
```

```json
{
  "message": "Failed to delete answer script file"
}
```


---

#### Shared Errors

##### **(400 Bad Request):**

###### Type conversion error:

```json
{
  "message": "Invalid input",
  "error": "Error message of how input is invalid"
}
```

###### Validation error

```json
{
  "message": "Invalid input",
  "errors": "Validation errors as string"
}
```