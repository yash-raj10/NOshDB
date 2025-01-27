# yamlDB

**yamlDB** is a lightweight, API-based NoSQL database that stores data in YAML format. It is designed for simplicity and ease of use, allowing developers to quickly store, retrieve, and manage structured data without requiring a full-scale database server. The database can be interacted with using HTTP APIs.

## Features

- **File-Based Storage**: Data is stored locally in YAML files, making it easy to read and manage.
- **No External Dependencies**: No need for external servers or installations.
- **Platform Support**: Works seamlessly on Linux, macOS (Darwin), and Windows.
- **API-Driven**: Expose REST APIs to interact with the database.
- **Pre-Built Binaries**: Use ready-made binaries for Darwin, Linux, and Windows for quick and easy setup.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yash-raj10/yamlDB.git
   cd yamlDB
   ```

2. Build the binary for your platform:

   **Linux**:

   ```bash
   GOOS=linux GOARCH=amd64 go build -o yamlDBLI
   ```

   **macOS (Darwin)**:

   ```bash
   GOOS=darwin GOARCH=amd64 go build -o yamlDBDAR
   ```

   **Windows**:

   ```bash
   GOOS=windows GOARCH=amd64 go build -o yamlDBWIN.exe
   ```

3. Move the binary to a directory in your `PATH` (optional):

   ```bash
   mv yamlDB /usr/local/bin
   ```

4. Alternatively, download the pre-built binaries directly from the repository for quick use:

   - **macOS**: `yamldbDAR`
   - **Linux**: `yamldbLI`
   - **Windows**: `yamldbWIN.exe`

   Simply download and execute the appropriate binary for your system.

---

## Usage

### Starting the Database Server

Run the binary to start the database server locally:

```bash
./yamldb...
```

The server will start on `http://localhost:8080` by default. You can configure the port in the source code or through flags (if implemented).

---

### Stopping the Database Server

1. Find the process ID (PID) of the running instance:

   ```bash
   ps aux | grep yamldbLI  # For Linux
   ps aux | grep yamldbDAR  # For macOS
   ps aux | grep yamldbWIN  # For Windows

2. Use the kill command to stop the process:

   ```bash
   kill <PID>


### API Endpoints

#### 1. Health Check

Verify if the server is running:

```bash
GET /
```

Response:

```json
{
  "Update": "working"
}
```

#### 2. Add Data to a Collection

Endpoint: `POST /postData`

Headers:

- `database`: Name of the database.
- `collection`: Name of the collection.

Body:

```json
{
  "key": "value",
  "anotherKey": "anotherValue"
}
```

Response:

```json
{
  "message": "data saved successfully!",
  "Id": "<unique-id>"
}
```

#### 3. Retrieve All Data from a Collection

Endpoint: `GET /getData/:database/:collection`

Example:

```bash
GET /getData/myDatabase/myCollection
```

Response:

```json
[
  { "id": "123", "key": "value" },
  { "id": "124", "anotherKey": "anotherValue" }
]
```

#### 4. Retrieve Data by ID

Endpoint: `GET /getData/:database/:collection/:id`

Example:

```bash
GET /getData/myDatabase/myCollection/123
```

Response:

```json
{
  "id": "123",
  "key": "value"
}
```

#### 5. Delete Data by ID

Endpoint: `DELETE /delete/:database/:collection/:id`

Example:

```bash
DELETE /delete/myDatabase/myCollection/123
```

Response:

```json
{
  "message": "item deleted with ID 123"
}
```



## YAML Storage Structure

The data is stored in a directory structure like this:

```
Databases/
  myDatabase/
    myCollection.yaml
```

Each collection file contains an array of objects stored in YAML format:

```yaml
- id: "123"
  key: "value"
- id: "124"
  anotherKey: "anotherValue"
```

---

## License

This project is licensed under the [MIT License](LICENSE).

