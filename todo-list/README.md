<div align="center"> <a href="https://genezio.com/"></a>
<img alt="genezio logo" src="https://github.com/genez-io/graphics/raw/HEAD/svg/Icon_Genezio_Black.svg" style="max-height: 50px;">

</div>

<div align="center">

[![deployed with: genezio](https://img.shields.io/badge/deployed_with-genezio-6742c1.svg?labelColor=62C353&style=flat)](https://github.com/genez-io/genezio)

[![Join our community](https://img.shields.io/discord/1024296197575422022?style=social&label=Join%20our%20community%20&logo=discord&labelColor=6A7EC2)](https://discord.gg/uc9H5YKjXv)
[![Follow @geneziodev](https://img.shields.io/twitter/url/https/twitter.com/geneziodev.svg?style=social&label=Follow%20%40geneziodev)](https://twitter.com/geneziodev)

</div>

# React todo application implemented with MongoDB and genezio

This is an example of a todo application that uses VueJS for the frontend application and Genezio for deploying and developing the backend.

## Prerequisites

- ✅ [Go](https://go.dev/)
- ✅ [npm](https://www.npmjs.com/)
- ✅ [genezio](https://genezio.com/)

1. Host a Mongo Database. Follow this [tutorial](https://genezio.com/docs/tutorials/connect-to-mongodb-atlas) to get a free tier database.
2. Create a `server/.env` file and add the following environment variables:

```env
MONGO_DB_URI=<your-mongo-uri>
```

## Project Structure

Inside the project folder, you will find the following files and folders:

```
├── server/
│   ├── go.mod
│   ├── go.sum
│   ├── backendService
│   |   |── backendService.go
├── client/
│   ├── src/
│   ├── public
│   ├── dist
│   ├── index.html
│   ├── node_modules
│   ├── package.json
│   ├── package-lock.json
│   ├── vite.config.ts
│   └── README.md
├── genezio.yaml
├── README.md
└── .genezioignore
```

Genezio looks for `genezio.yaml` to read the settings for deploying the project or for spinning a local dev server for testing.

The `backend` directory contains the implementation of the server side of the project.

The `frontend` directory contains a simple NodeJS application that talks with the genezio server.

To glue this two component together, an auto-generated SDK is installed in the `client/node_modules` folder.
This can be used by simply importing it into the frontend source code like any other dependency of your project.

## Run the project

### Clone this example

Clone the repository:

```
git clone https://github.com/Genez-io/genezio-examples
```

Navigate to the following directory:

```
cd ./genezio-examples/golang/todo-list
```

### Test your project locally

Test the project locally:

```
genezio local
```

Open a new terminal, navigate to the following directory, and run `npm run dev` to launch the Vue application:

```
cd ./client
npm install
npm run dev
```

### Deploy your project with genezio

If you wish to deploy your project to the Genezio infrastructure, follow these steps:

Log in to Genezio using the command genezio login:

```
genezio login
```

Deploy your project using the genezio deploy command from the `./genezio-examples/golang/todo-list` directory.

```
genezio deploy
```

## Commands

All commands are run from the root of the project, from a terminal:

| Command                  | Action                       |
| :----------------------- | :--------------------------- |
| `npm install -g genezio` | Installs genezio globally    |
| `genezio login`          | Logs in to genezio           |
| `genezio local`          | Starts a local server        |
| `genezio deploy`         | Deploys a production project |
| `genezio --help`         | Get help using genezio       |

## Want to learn more?

Check out:

- [Official genezio documentation](https://genezio.com/docs)
- [Web development tutorials](https://genezio.com/blog)
- [Discord channel](https://discord.gg/uc9H5YKjXv)

## Contact

If you need support or you have any questions, please join us in our [Discord channel](). We'd love to chat!

## Built With

- [Vue](https://vuejs.org/)
- [Genezio](https://genezio.com/)
- [Go](https://go.dev/)
- [Vite](https://vitejs.dev/)
