
## Deployment Documentation (Docker)

### Building Docker Image
```bash
docker build -t golang-memory-store .
```

### Running Docker Container
```bash
docker run -p 8080:8080 -d golang-memory-store
```

### Stopping Docker Container
```bash
docker stop <container_id>
```

### Removing Docker Container
```bash
docker rm <container_id>
```

### Environment Variables (For Production)
- `JWT_SECRET`: Secret Key for JWT authentication.
- `PERSISTENCE_FILE`: File name to save and load data.

### Example Docker Compose (Optional)
```yaml
version: '3'
services:
  in-memory-store:
    image: golang-memory-store
    ports:
      - "8080:8080"
    environment:
      - JWT_SECRET=supersecretkey
      - PERSISTENCE_FILE=data.json
```
