# learngo-fakestoreapi
A Go microservice with REST apis and RDBMS integration


## setup from scratch

*** Folder structure ***
``` bash
fakestoreapi/
│
├── cmd/server/             → Entry point (main.go)
│
├── config/                 → Handles configuration loading (env, JSON, etc.)
│
├── internal/               → Application’s internal code (domain logic)
│   ├── api/                → Gin route definitions
│   ├── handler/            → Request handlers/controllers
│   ├── service/            → Business logic and service layer
│   ├── repository/         → Database interactions (MySQL)
│   ├── client/             → External API clients (like FakeStore)
│   ├── model/              → Domain/data models (DTOs, structs)
│   └── middleware/         → Custom Gin middleware (auth, logger, etc.)
│
├── pkg/                    → Shared reusable code (logger, http client, etc.)
│
├── test/                   → Test code (unit, integration)
│
├── go.mod
├── go.sum
├── Dockerfile              → For containerizing your app
└── README.md


```

*** Adding the dependencies ****
If you are adding any dependencies(libs) which is outside of go then you need to install it first for import of dependencies to work. There are two ways of doing that. Fo exaple in order to import gingonic web application dependency for go you can do it in either of the following way:

1. In your root directory run following command

``` bash
go get github.com/gin-gonic/gin

```
Thiw would install the latest version of gin and add an entry in go.mod with a direct and many indirect dependencies.

```
require github.com/gin-gonic/gin v1.10.0

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
    .
    .
    .
)

```

2. Other way is to add the require dependency yourself in go.mod and run the following command:
```
require github.com/gin-gonic/gin v1.10.0
```

```bash
    go mod tidy 
    -- or you can run 
    go build
    -- or 
    go run main.go
```
**Note:** When managing Go dependencies, you only need to specify direct dependencies in your `go.mod` file. Go will automatically handle the resolution and inclusion of indirect dependencies.

*   **Run `go mod tidy` Regularly:**  After adding or removing imports manually, always execute `go mod tidy`.
*   **Maintain Cleanliness:** `go mod tidy` ensures your `go.mod` and `go.sum` files remain clean and accurate.
*   **Dependency Management Actions:** The `go mod tidy` command performs the following actions:
    *   Downloads any modules that are not already present locally.
    *   Removes unused imports from your `go.mod` file.
    *   Synchronizes the `go.sum` file to reflect the current dependencies.


## Phase 2: Service layer

We would follow the clean architecture: handler → service → client

🧱 Step-by-Step Plan
✅ Step 2.1: Define the domain model (Product)
✅ Step 2.2: Create the external API client (FakeStoreClient)
✅ Step 2.3: Create a service layer with an interface
✅ Step 2.4: Hook it into the Gin route
✅ Step 2.5: Test it by hitting new endpoint

