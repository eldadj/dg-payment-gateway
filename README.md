# Digital Colliers Payment-Gateway API
Payment gateway task for Senior Golang Software developer at Digital Colliers

# Notes
A. Database
1. PostgreSQL database. Dump file is included and executed during Docker run.
2. All SQL operations take place in application code. In most cases, it would be better to create views and stored procedures.

B. Deployment
1. Application can be deployed using Docker. Necessary Dockerfile and docker-compose files are included
2. Docker-compose also setups PostgreSQL as a dependency.
3. Environment settings are store in .env file

C. Development
1. JWT is use for Merchant authentication.
2. Middleware verifies authorization token for all endpoints except the initial authentication endpoint.
3. Custom error variables are used and stored in a single location
4. All endpoint request values are validated
5. All database operations are in performed within transactions

D. Choices Made
1. gorm ORM was used. In situations where sql statements are simple, raw sql mode is used.
2. beego MVC framework was used. It reads configuration settings from files store in conf folder
3. testify for tests. A parent TestSuite was created to enable reuse of functionality

E. Assumptions
1. Request/response data is application/json

F. Extra mile
1. Luhn check on credit card number
2. Client (merchant) authentication
3. Containerization
4. Application logging with logrus
