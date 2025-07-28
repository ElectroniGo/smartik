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

## Getting Started

For the complete development configuration, see the [getting started guide](../../docs/getting-started.md).

## API Reference

### Base URL

```
http://localhost:1323/api/v1
```

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