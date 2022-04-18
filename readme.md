# 1. Banking Microservice

#### Design Decisions

    1. Hexagonal Architecture.
    2. REST Api.
    3. Dependency Injection.
    4. Database Transactions.
    5. Structured Error Library.
    6. Concurrent Server with Graceful Shutdown.
    7. Distributed Logging.

#### Tools Used

    1. Go
    2. Postgres
    3. Docker

#### Libraries Used

    1. Pgxpool (Concurrency Safe Postgres Driver)
    2. gin-gonic (Web Frameowrk)
    3. Validator (JSON Validatin Middleware)

#### Routes

    1. GetCustomerById:         GET     /customers/:id
    2. GetCustomersByStatus:    GET     /customers?status=
    3. CreateCustomer           POST    /customers
    4. UpdateCustomer           PUT     /customers/:id
    
    To Be Implented
    5.  NewAccount:             POST    /customers/:id/accounts
    6.. NewTransaction:         POST    /customers/:id/account/:account_id

#### Data Exchange Format

    1. JSON (REST)
    2. CSV (Database bulk insert)

#### Models

    1. Customer:
        - id: bigserial
        - name: string
        - date_of_birth: date
        - city: string
        - zipcode: string
        - status: small int

    2. Accounts:
        - id: int
        - customer_id: int
        - opening_date: timestamp
        - amount: decimal
        - account_type: string

    4. Users:
        - id: bigserial
        - username: string
        - password: string
        - role: string

    3. Transactions:
        - id: bigserial
        - account_id: int
        - type: string
        - amount: decimal
        - transacation_date: timestamp


#### Services

    1. Customer:
        - Get All Customers By Status
        - Get Customer By Id

    2. Account:
        - New Account
        - Get Account By account_id
        - Get all accounts by account_id and customer_id

    3. Transaction:
        - New Transaction


# 2. Auth Microservice(YET TO BE IMPLEMENTED)


#### Auth Policy

    1. Role based access control.
    2. JWT based authentication.
    3. JWT based authorization.
    4. JWT based token verification.
    5. JWT based token refresh(TODO).

#### JWT Auth Process

    1. (user -> auth-server) login request.
    2. (auth-server -> user) token in response.
    3. (user -> banking-server) request resource with token.
    4. (banking server -> auth server) verify the token.
    5. (auth-server -> banking-server) token verification response.
    6. (bankng-server -> user) resource response.

#### Routes

    1. GetAllCustomers:     GET     /customers
    2. GetCustomer:         GET     /customers/:id
    3  NewAccount:          POST    /customers/:id/accounts
    4. NewTransaction:      POST    /customers/:id/account/:account_id


#### RBAC

    1. Role: admin  -> All.
    2. Role: user   -> GetCustomer & NewTransaction.

#### Verification Process

    1. Validity of the token(include expiry time and signature).
    2. Verify if the role has access to the resource.
    3. vefify if the resource being accessd by same user.



### Overall Project Status - Ongoing
