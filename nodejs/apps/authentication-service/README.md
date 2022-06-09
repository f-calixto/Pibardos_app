![badge](https://img.shields.io/badge/microservice-authentication-informational?style=for-the-badge&logo=node.js)

# Authentication service

This service will handle user authentication and email & password update.

## Tech stack
We are using the following libraries:
- **jsonwebtoken** - create an user's access token
- **bcryptjs** - hash user's password
- **dotenv** - load environmental variables from a .env file
- **express** - web application framework
- **mongoose** - ODM for mongoDB
- **mongoose**-unique-validator - validator for mongoose unique fields
- **morgan** - logger
- **uuid** - unique identifier for user's document
- **validator** - mongoose fields validator
- **joi** - validator with schemas

## Environmental variables
| Variable               | Default value |
| ---------------------- |:-------------:|
| PORT                   | 3000          |
| MONGODB_URI_CONNECTION | undefined     |
| ACCESS_TOKEN_SECRET    | undefined     |
| ACCESS_TOKEN_EXP_TIME  | undefined     |
| BCRYPT_SALT_ROUNDS     | undefined     |

## Usage
To start the service you will need Docker and docker-compose installed in your computer and run the following command:
*This command will start node service, mongodb and rabbitmq in its respective containers.*

```bash
docker-compose -f docker-compose.dev.yml up --build
```

## Endpoints
- **POST** > /login
- **POST** > /register

## License
[MIT](https://choosealicense.com/licenses/mit/)
